# Project Setup

## Pre requirements
* **Install go-task**, visit [Taskfile.dev](https://taskfile.dev/)
    * `brew install go-task`
* **Docker login**
    * `task docker-login`
    * Search and change `USERNAME` with your Docker Hub username
* **Install kubectl**
    * `task part-2-install-kubectl`
* **Install minikube**
    * `task part-2-install-minikube`
* **Install istio**
    * `task part-3-install-istioctl`

## 1.0 Part
* **1.1 Implement two simple web services**
    * `task part-1-app-v1-run`
    * `task part-1-app-v2-run`
* **1.2 Dockerize the services**
* **1.3 Tag and push the images to Docker Hub**
    * `task part-1-app-v1-push-tag`
    * `task part-1-app-v2-push-tag`

## 2.0 Part
* **2.1 Start minikube**
    * `task part-2-start-minikube`
* **2.2 Deploy app-v1 as a Pod**
    * `task part-2-deploy-app-v1`
    * **2.2.1 Run commands inside the Pod**
        * `task part-2-inside-pod-app-v1`
            * `apk add curl`
            * `curl localhost:3000`
* **2.3 Delete the Pod**
    * `task part-2-delete-pod-app-v1`
* **2.4 Deploy services and deployments**
    * `task part-2-deployment-app-v1`
    * `task part-2-deployment-app-v2`
* **2.5 Expose the services**
    * `task part-2-expose-app-v1`
    * `task part-2-expose-app-v2`
* **2.6 Delete the services and deployments**
    * `task part-2-delete-all`

## 3.0 Part
* **3.1 Create an HTTP Gateway to route traffic for specific hostnames**
    * `task part-3-applying-gateway`
* **3.2 Create a Minikube tunnel to expose services locally**
    * `task part-3-minikube-tunnel`
* **3.3 Istio App Deployment**
    * **3.3.1 Default**
        * `task part-3-istio-app-deploy-v1`
        * `task part-3-istio-app-deploy-v2`
    * **3.3.2 Traffic separated by query params**
        * `task part-3-istio-app-deploy-seperated-by-query-params`
    * **3.3.3 Create a VirtualService to route and rewrite HTTP traffic based on URL paths and match conditions, with weighted traffic distribution**
        * `task part-3-istio-app-deploy-seperated-by-path`
    * **3.3.4 Response Header Manipulation**
        * `task part-3-istio-app-deploy-modify-response-headers`
* **3.4 Delete all deployments and services**
    * `task part-3-delete-all`