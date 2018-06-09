# gophercon 2018 demo

This repo contains all the software that is going to be used during the Gophercon 2018 demo of my talk.

## Prerequisites

- A Kubernetes cluster somewhere. For example, if you want one on gcloud:
 
      $ PROJECT_ID=gophercon-$RANDOM
      $ gcloud projects create $PROJECT_ID
      $ gcloud config set project $PROJECT_ID
      $ gcloud config set compute/zone europe-west1-b
      $ gcloud container clusters create gke # It will ask you to enable Kubernetes, do it

  If you are cheap (like me) you can create a cluster of preemptible machines instead:

      $ gcloud beta container clusters create gke-preemptible --preemptible
      $ gcloud config set container/cluster gke-preemptible
      $ gcloud beta container node-pools create preemptible-pool --preemptible
      $ gcloud beta container node-pools delete default-pool

  Now you can get the credentials to start using your cluster:

      $ gcloud container clusters get-credentials gke # or gke-preemptible

- Helm on it. If you are having RBAC problems: https://github.com/kubernetes/helm/blob/master/docs/rbac.md

      $ helm init

- Jenkins installed:
    
      $ helm install -f hack/jenkins-values.yaml stable/jenkins

  On Jenkins you will also need a "username and password" secret called: `agonzalezro-hub-credentials` (or change the `Jenkinsfile` if you want to use other ID) for having access to push images to Docker Hub.

- If you want to use the In-Cluster Config for `kubectl` without having to bother about permissions permissions you can give them all to the default account:

      $ kubectl create -f hack/do-not-do-this-at-home-rbac.yaml

## What else?

**Come to the talk and you will see!**
