## Istio App Deployment - Default
```bash
kubectl apply -f virtualService.yaml

open -a "Google Chrome" "http://seperated-by-path.box/seperated"
open -a "Google Chrome" "http://seperated-by-path.box/v2"
open -a "Google Chrome" "http://seperated-by-path.box/v1"
```
