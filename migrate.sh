if [ -z $1 ]
then
  echo "> error: mysql password not specified"
  exit 1
fi
MYSQL_POD_NAME=$(kubectl get pods -o name | grep deployment-mysql)
for migration in ./mysql/migrations/v*.sql; do
  [ -e $migration ] || continue
  echo "> applying $migration..."
  kubectl exec -i $MYSQL_POD_NAME -- mysql --database koran -p$1 -u root < $migration
done
