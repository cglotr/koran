# koran

## Development

1. Install [Helm](https://github.com/helm/helm)
2. `helm install stable/nginx-ingress`
3. `kubectl create secret generic secret-mysql-password --from-literal=MYSQL_PASSWORD=$MYSQL_PASSWORD`
4. `kubectl apply -f kubernetes`
