if [ -z $1 ]
then
  echo "🚨 mysql password not specified"
  exit 1
fi
MYSQL_POD_NAME=$(kubectl get pods -o name | grep deployment-mysql)
for seed in ./mysql/seeds/*.sql; do
  [ -e $seed ] || continue
  echo "👩🏼‍🌾 seeding $seed..."
  kubectl exec -i $MYSQL_POD_NAME -- mysql -p$1 -u root --database koran --default-character-set=utf8mb4 < $seed
done
