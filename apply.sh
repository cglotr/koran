kubectl apply -f kubernetes

kubectl set image deployments/deployment-client client=cglotr/koran-client:$(git rev-parse HEAD)
kubectl set image deployments/deployment-server server=cglotr/koran-server:$(git rev-parse HEAD)
