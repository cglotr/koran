docker build -t cglotr/koran-client -t cglotr/koran-client:$(git rev-parse HEAD) ./client
docker push cglotr/koran-client:latest
docker push cglotr/koran-client:$(git rev-parse HEAD)
kubectl apply -f kubernetes/deployment-client.yml
kubectl set image deployments/deployment-client client=cglotr/koran-client:$(git rev-parse HEAD)
