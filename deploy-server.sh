docker build -t cglotr/koran-server -t cglotr/koran-server:$(git rev-parse HEAD) ./server
docker push cglotr/koran-server:latest
docker push cglotr/koran-server:$(git rev-parse HEAD)
kubectl apply -f kubernetes/deployment-server.yml
kubectl set image deployments/deployment-server server=cglotr/koran-server:$(git rev-parse HEAD)
