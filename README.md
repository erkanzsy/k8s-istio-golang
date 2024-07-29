* Pre requirements
    * Install go-task, visit [Taskfile.dev](https://taskfile.dev/)
        * `brew install go-task` 
    * Docker login
        * `task docker-login`
        * Search and change USERNAME with your docker hub username
    * Install kubectl
        * `task part-2-install-kubectl`
    * Install minikube
        * `task part-2-install-minikube`
    * Install istio
        * `task part-3-install-istioctl`

* Part-1
    * Implement two simple web services
        * `task part-1-app-v1-run`
        * `task part-1-app-v2-run`
    * Dockerize the services
    * Tag and push the images to docker hub
        * `task part-1-app-v1-push-tag`
        * `task part-1-app-v2-push-tag`

* Part-2
    * Start minikube
        * `task part-2-start-minikube`
    * Deploy app-v1 as a Pod
        * `task part-2-deploy-app-v1`
        * `task part-2-inside-pod-app-v1`
            * `apk add curl`
            * `curl localhost:3000`
    * Delete the pod
        * `task part-2-delete-pod-app-v1`
    * Deploy services and deployments
        * `task part-2-deployment-app-v1`
        * `task part-2-deployment-app-v2`
    * Expose the services
        * `task part-2-expose-app-v1`
        * `task part-2-expose-app-v2`
    * Delete the services and deployments
        * `task part-2-delete-all`

* Part-3
    * Create an HTTP Gateway to route traffic for specific hostnames
        * `task part-3-applying-gateway`
    * Create a Minikube tunnel to expose services locally
        * `task part-3-minikube-tunnel`
    * Istio App Deployment
        * Default
            * `task part-3-istio-app-deploy-v1`
            * `task part-3-istio-app-deploy-v2`
        * Traffic seperated by query params
            * `task part-3-istio-app-deploy-seperated-by-query-params`
        * Create a VirtualService to route and rewrite HTTP traffic based on URL paths and match conditions, with
          weighted traffic distribution
            * `task part-3-istio-app-deploy-seperated-by-path`
        * Response Header Manipulation
            * `task part-3-istio-app-deploy-modify-response-headers`
    * Delete all deployments and services
        * `task part-3-delete-all`