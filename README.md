# Телеграм-бот «Печеньки для Паймон»
![Coverage](https://img.shields.io/badge/Coverage-58.3%25-yellow)

#### Version 0.2.4

## Run
Run project dev environment and all required migrations

```shell
make BOT_TOKEN=<set your telegram bot token> start
```

For «silent» (detached) start use

```shell
make BOT_TOKEN=<set your telegram bot token> silent-start
```

## Requirements

* Go 1.21
* PostgreSQL
* Docker & Docker Compose
