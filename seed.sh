if [ -z $1 ]
then
  echo "> error: mysql password not specified"
  exit 1
fi
MYSQL_POD_NAME=$(kubectl get pods -o name | grep deployment-mysql)
for seed in ./mysql/seeds/*.sql; do
  [ -e $seed ] || continue
  echo "👩🏼‍🌾 seeding $seed..."
  kubectl exec -i $MYSQL_POD_NAME -- mysql --database koran -p$1 -u root < $seed
done
