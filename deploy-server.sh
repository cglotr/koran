docker build -t cglotr/koran-server ./server
docker push cglotr/koran-server:latest
kubectl apply -f kubernetes/deployment-server.yml
