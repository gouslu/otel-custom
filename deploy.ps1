kind load docker-image otel-custom:latest
kubectl delete -f ./otel-custom.yaml
kubectl apply -f ./otel-custom.yaml