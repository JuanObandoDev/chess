# Chess

Multiplayer simple chess project

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [mkcert](https://github.com/FiloSottile/mkcert)

## Setup

1. Run the following command to install the mkcert root CA and generate a new SSL certificate for the chess.localhost domain:

```bash
mkcert -install

mkcert -cert-file traefik/certs/local-cert.pem -key-file traefik/certs/local-key.pem "chess.localhost" "*.chess.localhost"
```

2. Create a Docker network for Traefik:

```bash
docker network create chess-dev-traefik-net
```

## Development

Start the project by running the following command:

```bash
docker-compose -f docker-compose.dev.yml up
```

you can access the client application in your web browser at [client.chess.localhost](https://client.chess.localhost).

Developers may also use (devcontainers)[https://code.visualstudio.com/docs/devcontainers/containers] to further streamline their development process. Devcontainers is a feature of Visual Studio Code that allows you to develop in a containerized environment.
