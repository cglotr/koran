docker build -t arikama/koran-server -t arikama/koran-server:$(git rev-parse HEAD) ./server
docker push arikama/koran-server:latest
docker push arikama/koran-server:$(git rev-parse HEAD)
kubectl apply -f kubernetes/deployment-server.yaml
kubectl set image deployments/deployment-server server=arikama/koran-server:$(git rev-parse HEAD)
