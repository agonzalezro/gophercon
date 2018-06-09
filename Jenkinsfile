podTemplate(label: 'mypod', containers: [
    containerTemplate(name: 'docker', image: 'docker:stable-dind', ttyEnabled: true, alwaysPullImage: true, privileged: true,
      command: 'dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay'),
    containerTemplate(name: 'kubectl', image: 'lachlanevenson/k8s-kubectl:v1.10.5', ttyEnabled: true, alwaysPullImage: true, command: 'cat')
  ],
  volumes: [emptyDirVolume(memory: false, mountPath: '/var/lib/docker')]) {

  node('mypod') {
    stage('Checkout') {
      checkout scm
      shortCommit = sh(returnStdout: true, script: "git log -n 1 --pretty=format:'%h'").trim()
    }
    
    stage('Test') {
      container('docker') {
          sh 'docker build --target builder -t builder .'
          sh 'docker run --rm builder go test . -v'
      }
    }

    stage('Build and push') {
      container('docker') {
        docker.withRegistry('https://registry-1.docker.io/v2/', 'agonzalezro-hub-credentials') {
          dockerImage = 'agonzalezro/gophercon'
          built = docker.build(dockerImage)
          built.push(shortCommit)
          built.push('latest')
        }
      }
    }

    stage('Deploy to Kubernetes') {
      container('kubectl') {
        sh 'kubectl apply -f k8s/deployment.yaml'
        sh 'kubectl apply -f k8s/service.yaml'
        sh "kubectl set image deployment/gophercon-deployment gophercon=${dockerImage}:${shortCommit}"
      }
    }
  }
}