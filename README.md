# ArchMark [![.github/workflows/ci.yml](https://github.com/sparkymat/archmark/actions/workflows/ci.yml/badge.svg)](https://github.com/sparkymat/archmark/actions/workflows/ci.yml)

ArchMark is a bookmark manager that archives the bookmarked page using [Monolith](https://github.com/Y2Z/monolith). It consists of the main web proccess as well as a worker that downloads the web page for archival.

## Installation

The simplest way to use ArchMark would be to use docker-compose. If you have [Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/) installed, you can follow the steps below to get up and running:

1. `mkdir archmark`
2. `cd archmark`
3. `curl https://raw.githubusercontent.com/sparkymat/archmark/main/docker-compose.prod.yml -o docker-compose.yml`
4. `mkdir -p data/db`
5. `mkdir -p data/faktory`
6. `mkdir -p data/archive`
7. `docker-compose pull`
8. `docker-compose up`

Alternatively, you can build and run from the code with:

1. `git clone https://github.com/sparkymat/archmark`
2. `cd archmark`
3. `docker-compose build`
4. `docker-compose up`

P.S: Don't forget to edit the `docker-compose.yml` file, and update the value of `JWT_SECRET` and `SESSION_SECRET` to something more secure. You can generate secrets using openssl like this:

```
openssl rand -hex 32
```

N.B: Everything in the download folder (configured using `DOWNLOAD_FOLDER`) will be served under `/uploads/` sub-path. 

## Configuration

Both the app and the worker are configured using environment variables.

| Variable                       | Description                                                                 |
| ------------------------------ | ----------------------------------------------------------------------------|
| `DISABLE_REGISTRATION`         | Set to 'true' if you want to disable registrations                          |
| `REVERSE_PROXY_AUTHENTICATION` | Enables reverse proxy authentication via forwarded headers                  |
| `PROXY_AUTH_USERNAME_HEADER`   | The request header where proxy passes the usernamae. Default: `Remote-User` |
| `PROXY_AUTH_NAME_HEADER`       | The request header where proxy passes the full name. Default: `Remote-Name` |
