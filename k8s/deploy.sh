kubectl apply -f ping_deployment.yaml
kubectl apply -f pong_deployment.yaml
kubectl apply -f rabbit_deployment.yaml

kubectl apply -f ping_service.yaml
kubectl apply -f pong_service.yaml
kubectl apply -f rabbit_service.yaml