#### Project Run (Option 1)
```bash
kubectl apply -f v1/deployment.yaml
kubectl apply -f v1/service.yaml

kubectl apply -f v2/deployment.yaml
kubectl apply -f v2/service.yaml

minikube service list

minikube service app-v1-service
minikube service app-v2-service


kubectl get all -n default
  
kubectl delete pods --all -n default

kubectl delete deployments --all -n default
```


#### Project Run (Option 2)
```bash
cd 2-deployment
kubectl apply -f v1
kubectl apply -f v2
```