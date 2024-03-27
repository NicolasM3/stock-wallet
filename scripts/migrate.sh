source .env

docker run --rm --network stock-wallet_stock-network -v $PWD/internal/migration:/flyway/sql flyway/flyway -url=jdbc:postgresql://$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=60 migrate

