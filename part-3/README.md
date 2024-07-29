## Istio Install

```bash
brew install istioctl

kubectl label namespace default istio-injection=enabled
```

## Hostları Yapılandırma

```bash
127.0.0.1 app-v1.box
127.0.0.1 app-v2.box
127.0.0.1 seperated-by-query-params.box
127.0.0.1 seperated-by-path.box
127.0.0.1 response-headers.box
```

## Gateway Setup
```bash
istioctl install --set profile=default
kubectl apply -f gateway.yaml
```

## Expose Istio Traffic
```bash
minikube tunnel
```