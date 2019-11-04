if [ -z $1 ]
then
  echo "> error: mysql password not specified"
  exit 1
fi
MYSQL_POD_NAME=$(kubectl get pods -o name | grep deployment-mysql)
kubectl exec -it $MYSQL_POD_NAME -- mysql -D koran -p$1 -u root
