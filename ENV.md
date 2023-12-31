# Available environment variable

## General

| Env name       | Default value | Available value         | Description                                                          |
|----------------|---------------|-------------------------|----------------------------------------------------------------------|
| ENVIRONMENT    | production    | production, development | This is env is for general purpose but affect the log format to json |
| APP_NAME       |               |                         | This is env is used in some log info and as jwt audience fields      |
| JWT_EXP_TIME   | 900           |                         | JWT expiration time in seconds                                       |
| JWT_SECRET_KEY |               |                         | MANDATORY                                                            |

## Database

| Env name         | default value | Available value | Description |
|------------------|---------------|-----------------|-------------|
| DB_HOST          | db-host       |                 |             |
| DB_PORT          | db-port       |                 |             |
| DB_USER          | db-user       |                 |             |
| DB_PASSWORD      | db-password   |                 |             |
| DB_SCHEMA        | db-schema     |                 |             |
| DB_CONN_RETRY    | 10            |                 |             |
| DB_MAX_OPEN_CONN | 150           |                 |             |

## Logger

| Env name      | default value | Available value          | Description                                       |
|---------------|---------------|--------------------------|---------------------------------------------------|
| LOGGER_LEVEL  | debug         | debug, info, warn, error | if provided wrong value, info level will be used  |
| LOGGER_FORMAT | json          | json, text               | if provided wrong value, json format will be used |
