[EN](README.md) | **[PT-BR](README_pt-BR.md)**

# Rate Limiter em GO

O projeto consiste em um middleware que controla a quantidade de requisições que um aplicativo pode fazer em um período de tempo.

## Requisitos para rodar o código

[Docker Compose](https://docs.docker.com/compose/install/)


## Instruções para rodar o código

Na pasta do projeto, execute:

```
docker-compose up -d
```

A API estará disponível em http://localhost:8080.

O rate limiter estará configurado para monitorar as requisições à rota `/api` que apenas retorna um *"Hello, world!"*

## Configurações

Todas as configurações podem ser alteradas no arquivo `.env`.

- `RATE_LIMITER_TIMEOUT`: Define o tempo em segundos em que o usuário ficará sem poder fazer mais requisições após ultrapassar o limite. *Padrão: 20*
- `RATE_LIMITER_TOKEN_LIMIT`: Define o número máximo de requisições que um token pode fazer no período de um segundo. *Padrão: 5*
- `RATE_LIMITER_IP_LIMIT`: O mesmo que `RATE_LIMITER_TOKEN_LIMIT`, mas para limitar por IP. *Padrão: 10*