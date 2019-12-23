echo "CLIENT_FIREBASE_API_KEY=$CLIENT_FIREBASE_API_KEY" >> client/.env

docker build -t cglotr/koran-client -t cglotr/koran-client:$(git rev-parse HEAD) ./client
docker build -t cglotr/koran-server -t cglotr/koran-server:$(git rev-parse HEAD) ./server
