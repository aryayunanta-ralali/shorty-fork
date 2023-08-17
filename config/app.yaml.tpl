app:
  name: {{ APP_NAME }}
  port: {{ APP_PORT }}
  timezone: {{ APP_TIMEZONE }}
  debug: {{ APP_DEBUG }}
  env: {{ APP_ENV }} # dev | stg | prod
  read_timeout_second: {{ APP_READ_TIMEOUT_SECOND }}
  write_timeout_second: {{ APP_WRITE_TIMEOUT_SECOND }}
  key: "{{ APP_KEY }}"
  default_lang: "{{ APP_DEFAULT_LANG }}"

apm:
  address: "{{ APM_ADDRESS }}"
  enable: {{ APM_ENABLE }}
  name: {{ APM_NAME }}

logger:
  url: "{{ LOGGER_URL }}"
  level: "{{ LOGGER_LEVEL }}"

redis:
  host: "{{ REDIS_HOST }}"
  db: {{ REDIS_DB }} # 0
  password: "{{ REDIS_PASSWORD }}"
  read_timeout_second: {{ REDIS_READ_TIMEOUT_SECOND }} # 1  second
  write_timeout_second: {{ REDIS_WRITE_TIMEOUT_SECOND }} # 1  second
  pool_size: {{ REDIS_POOL_SIZE }} # 100
  pool_timeout_second: {{ REDIS_POOL_TIMEOUT_SECOND }} # 100
  min_idle_conn: {{ REDIS_MIN_IDLE }} # 10
  idle_timeout_second: {{ REDIS_IDLE_TIMEOUT_SECOND }} # 240
  route_by_latency: {{ REDIS_ROUTE_BY_LATENCY }} # true
  idle_frequency_check: {{ REDIS_IDLE_FREQUENCY_CHECK }} # 1
  read_only: {{ REDIS_READ_ONLY }}
  route_randomly: {{ REDIS_ROUTE_RANDOMLY }}
  max_redirect: {{ REDIS_MAX_REDIRECT }} # set 3 for default redis
  cluster_mode: {{ REDIS_CLUSTER_MODE }}
  tls_enable: {{ REDIS_TLS_ENABLE }}
  insecure_skip_verify: {{ REDIS_INSECURE_SKIP_VERIFY }} # if tls_enable == true, this config use for tls insecure_skip_verify true or false

aws:
  access_key: "{{ AWS_ACCESS_KEY }}" # let this empty on production
  access_secret: "{{ AWS_ACCESS_SECRET }}" # let this empty on production
  profile: "{{ AWS_PROFILE }}"
  region: "{{ AWS_REGION }}"


write_db:
  host: "{{ AURORA_WRITE_HOST }}"
  port: {{ AURORA_WRITE_PORT }}
  name: "{{ AURORA_WRITE_DBNAME }}" # database name
  user: "{{ AURORA_WRITE_USERNAME }}" # database user
  pass: "{{ AURORA_WRITE_PASSWORD }}" # database password
  max_open: {{ AURORA_WRITE_MAXOPEN }}
  max_idle: {{ AURORA_WRITE_MAXIDLE }}
  timeout_second: {{ AURORA_WRITE_TIMEOUT }}
  life_time_ms: {{ AURORA_WRITE_LIFETIME }}
  charset: "{{ AURORA_WRITE_CHARSET }}"

read_db:
  host: "{{ AURORA_READ_HOST }}"
  port: {{ AURORA_READ_PORT }}
  name: "{{ AURORA_READ_DBNAME }}" # database name
  user: "{{ AURORA_READ_USERNAME }}" # database user
  pass: "{{ AURORA_READ_PASSWORD }}" # database password
  max_open: {{ AURORA_READ_MAXOPEN }}
  max_idle: {{ AURORA_READ_MAXIDLE }}
  timeout_second: {{ AURORA_READ_TIMEOUT }}
  life_time_ms: {{ AURORA_READ_LIFETIME }}
  charset: "{{ AURORA_READ_CHARSET }}"

kafka:
  brokers: "{{ KAFKA_BROKERS }}"
  version: "{{ KAFKA_VERSION }}" # kafka version
  client_id: "{{ KAFKA_CLIENT_ID }}"
  channel_buffer_size: {{ KAFKA_CHANNEL_BUFFER_SIZE }} # how many message in channel, default 256
  tls:
    ca_file: "{{ KAFKA_TLS_CA_FILE }}"
    cert_file: "{{ KAFKA_TLS_CERT_FILE }}"
    key_file: "{{ KAFKA_TLS_KEY_FILE }}"
    skip_verify: {{ KAFKA_TLS_SKIP_VERIFY }}
    enable: {{ KAFKA_TLS_ENABLE }}
  sasl:
    enable: {{ KAFKA_SASL_ENABLE }}
    mechanism: "{{ KAFKA_SASL_MECHANISM }}"
    user: "{{ KAFKA_SASL_USER }}"
    password: "{{ KAFKA_SASL_PASSWORD }}"
    version: {{ KAFKA_SASL_VERSION }} # version available : 0, 1
    handshake: {{ KAFKA_SASL_HANDSHAKE }}
  producer:
    timeout_second: {{ KAFKA_PRODUCER_TIMEOUT_SECOND }} # in second
    ack: {{ KAFKA_PRODUCER_ACK }} # -1 wait for all, 0 = NoResponse doesn't send any response, the TCP ACK is all you get, 1 = WaitForLocal waits for only the local commit to succeed before responding
    idem_potent: {{ KAFKA_PRODUCER_IDEM_POTENT }} # If enabled, the producer will ensure that exactly one copy of each message is written
    partition_strategy: "{{ KAFKA_PRODUCER_PARTITION_STRATEGY }}" # available strategy : hash, roundrobin, manual, random, reference
  consumer:
    rebalance_strategy: "{{ KAFKA_CONSUMER_REBALANCE_STRATEGY }}" # range, sticky, roundrobin, all of strategy only use for consumer group
    session_timeout_second: {{ KAFKA_CONSUMER_SESSION_TIMEOUT_SECOND }} # timeout for consumer group
    heartbeat_interval_ms: {{ KAFKA_CONSUMER_HEARTBEAT_INTERVAL_MS }}
    offset_initial: {{ KAFKA_CONSUMER_OFFSET_INITIAL }} # The initial offset to use if no offset was previously committed. Should be -1 = newest or  -2 = oldest
    auto_commit: {{ KAFKA_CONSUMER_AUTO_COMMIT }}
    isolation_level: {{ KAFKA_CONSUMER_ISOLATION_LEVEL }} # `0 = ReadUncommitted` (default) to consume and return all messages in message channel, 1 = ReadCommitted` to hide messages that are part of an aborted transaction
