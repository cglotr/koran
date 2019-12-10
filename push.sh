docker push cglotr/koran-client:latest
docker push cglotr/koran-client:$(git rev-parse HEAD)

docker push cglotr/koran-server:latest
docker push cglotr/koran-server:$(git rev-parse HEAD)
