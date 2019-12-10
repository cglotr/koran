docker build -t cglotr/koran-client -t cglotr/koran-client:$(git rev-parse HEAD) ./client
docker build -t cglotr/koran-server -t cglotr/koran-server:$(git rev-parse HEAD) ./server
