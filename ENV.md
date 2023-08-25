# Available environment variable

## General

| Env name    | Default value | Available value         | Description                                                          |
|-------------|---------------|-------------------------|----------------------------------------------------------------------|
| ENVIRONMENT | development   | production, development | This is env is for general purpose but affect the log format to json |

## Database

| Env name         | default value | Available value | Description |
|------------------|---------------|-----------------|-------------|
| DB_HOST          | db-host       |                 |             |
| DB_PORT          | db-host       |                 |             |
| DB_USER          | db-host       |                 |             |
| DB_PASSWORD      | db-host       |                 |             |
| DB_SCHEMA        | db-host       |                 |             |
| DB_CONN_RETRY    | db-host       |                 |             |
| DB_MAX_OPEN_CONN | 150           |                 |             |

## Logger

| Env name     | default value | Available value          | Description                                      |
|--------------|---------------|--------------------------|--------------------------------------------------|
| LOGGER_LEVEL | debug         | debug, info, warn, error | if provided wrong value, info level will be used |
