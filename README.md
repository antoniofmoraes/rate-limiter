**[EN](README.md)** | [PT-BR](README_pt-BR.md)

# Rate Limiter in GO

The project consists of a middleware that controls the number of requests an application can make in a given period of time.

## Requirements to run the code

[Docker Compose](https://docs.docker.com/compose/install/)

## Instructions to run the code

In the project folder, run:
```
docker-compose up -d
```
The API will be available at http://localhost:8080.

The rate limiter will be configured to monitor requests to the `/api` route, which only returns a *"Hello, world!"*

## Configurations

All settings can be changed in the `.env` file.

- `RATE_LIMITER_TIMEOUT`: Defines the time in seconds that the user will be unable to make more requests after exceeding the limit. *Default: 20*
- `RATE_LIMITER_TOKEN_LIMIT`: Defines the maximum number of requests a token can make in one second. *Default: 5*
- `RATE_LIMITER_IP_LIMIT`: Same as `RATE_LIMITER_TOKEN_LIMIT`, but to limit by IP. *Default: 10*