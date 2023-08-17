#!/usr/bin/bash

# app
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_NAME --value shorty --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_ENV --value dev --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_TIMEZONE --value Asia/Jakarta --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_DEBUG --value true --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_PORT --value 8081 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_READ_TIMEOUT_SECOND --value 5 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_WRITE_TIMEOUT_SECOND --value 15 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/secret/shorty/api/APP_KEY --value 340c931fc23c3251079181762c046fbc9f8d40e39f7bd6a59cd1b0ef6605fff6 --type SecureString --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/APP_DEFAULT_LANG --value en --type String --overwrite


# logger
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/LOGGER_URL --value dsn-sentry --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/LOGGER_LEVEL --value warning --type String --overwrite


#redis
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_HOST --value redis-host --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_DB --value 1 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/secret/shorty/redis/REDIS_PASSWORD --value redisPassword --type SecureString --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_READ_TIMEOUT_SECOND  --value 2 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_WRITE_TIMEOUT_SECOND --value 2 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_POOL_SIZE --value 10 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_POOL_TIMEOUT_SECOND --value 10 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_MIN_IDLE --value 5 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_IDLE_TIMEOUT_SECOND --value 5 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_ROUTE_BY_LATENCY --value true --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/REDIS_IDLE_FREQUENCY_CHECK --value 1 --type String --overwrite

# aws account
aws ssm put-parameter --profile rll-dev --name /rll/dev/secret/shorty/iam/AWS_ACCESS_KEY --value awsAccesKey --type SecureString --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/secret/shorty/iam/AWS_ACCESS_SECRET --value awsSecretKey --type SecureString --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/AWS_PROFILE --value default --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/AWS_REGION --value ap-southeast-1 --type String --overwrite


# database Read
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_HOST --value dbHostRead --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_NAME --value dbNameRead --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_PORT --value 3308 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_CHARSET --value utf8 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_USERNAME --value dbUser --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/secret/shorty/database/DB_READ_PASSWORD --value dbPassword --type SecureString --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_TIMEOUT_SECOND --value 3 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_LIFETIME_MS --value default --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_MAX_OPEN --value 10 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_READ_MAX_IDLE --value 5 --type String --overwrite


# database Write
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_HOST --value dbHostWrite --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_NAME --value dbNameWrite --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_PORT --value 3308 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_CHARSET --value utf8 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_USERNAME --value dbUser --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/secret/shorty/database/DB_WRITE_PASSWORD --value dbPassword --type SecureString --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_TIMEOUT_SECOND --value 3 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_LIFETIME_MS --value default --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_MAX_OPEN --value 10 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/DB_WRITE_MAX_IDLE --value 5 --type String --overwrite


#kafka
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_BROKERS --value kafkbrokers2,kafkaBrokers2 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_VERSION --value 2.4.2 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_CLIENT_ID --value shorty --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_CHANNEL_BUFFER_SIZE --value 5 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_TLS_CA_FILE --value  --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_TLS_CERT_FILE --value  --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_TLS_KEY_FILE --value  --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_TLS_SKIP_VERIFY --value false --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_SASL_ENABLE --value false --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_SASL_MECHANISM --value  --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_SASL_USER --value --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/secret/shorty/kafka/KAFKA_SASL_PASSWORD --value kafkaSASLPaswd --type SecureString --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_SASL_VERSION --value 1 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_SASL_HANDSHAKE --value  --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_PRODUCER_TIMEOUT_SECOND --value 5 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_PRODUCER_ACK --value -1 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_PRODUCER_IDEM_POTENT --value true --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_PRODUCER_PARTITION_STRATEGY --value hash --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_CONSUMER_SESSION_TIMEOUT_SECOND --value 8 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_CONSUMER_HEARTBEAT_INTERVAL_MS --value 1500 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_CONSUMER_OFFSET_INITIAL --value -2 --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_CONSUMER_AUTO_COMMIT --value true --type String --overwrite
aws ssm put-parameter --profile rll-dev --name /rll/dev/general/shorty/KAFKA_CONSUMER_ISOLATION_LEVEL --value 1 --type String --overwrite