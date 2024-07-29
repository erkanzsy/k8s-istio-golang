## Istio App Deployment - Default
```bash
kubectl apply -f v1
kubectl apply -f v2

open -a "Google Chrome" "http://app-v1.box"
open -a "Google Chrome" "http://app-v2.box"

kubectl describe pods
```
