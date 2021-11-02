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

P.S: Don't forget to edit the `docker-compose.yml` file, and update the value of `ADMIN_PASSWORD` to something more secure.

N.B: Everything in the download folder (configured using `DOWNLOAD_FOLDER`) will be served under `/b/` sub-path. 

## API

The app exposes an API which can be used to add bookmarks. The API requests are authenticated using API tokens that the user can manage from the API Tokens tab. 

After creating a user token, you can use it as follows (using curl as an example):

> ```curl -XPOST -H "Content-Type: application/json" -H "Authorization: Bearer <token-here>" -d '{"url": "url-to-add-here"}'  "http://localhost:8080/api/add"```

## Configuration

Both the app and the worker are configured using environment variables.

| Variable           | Description                                                    |
| ------------------ | -------------------------------------------------------------- |
| `DB_HOSTNAME`      | Hostname of the machine where PostgreSQL is running            |
| `DB_PORT`          | Port on which PostgreSQL is running                            |
| `DB_USERNAME`      | PostgreSQL user                                                |
| `DB_PASSWORD`      | PostgreSQL password                                            |
| `DB_DATABASE`      | PostgreSQL database name                                       |
| `DB_SSL_MODE`      | `true` if the connection to PostgreSQL should be done over SSL |
| `ADMIN_PASSWORD`   | Password for the `admin` user                                  |
| `MONOLITH_PATH`    | Full path to the Monolith binary                               |
| `DOWNLOAD_PATH`    | Full path to the folder where archived pages are to be stored  |
| `DEFAULT_LANGUAGE` | Language for the web app. Supported: en, ml                    |

## Screenshots

![newbookmark](https://raw.githubusercontent.com/sparkymat/archmark/master/docs/newbookmark.png)

![pendingbookmark](https://raw.githubusercontent.com/sparkymat/archmark/master/docs/pendingbookmark.png)

![bookmarks](https://raw.githubusercontent.com/sparkymat/archmark/master/docs/bookmarks.png)

![deleteconfirm](https://raw.githubusercontent.com/sparkymat/archmark/master/docs/deleteconfirm.png)

![tokens](https://raw.githubusercontent.com/sparkymat/archmark/master/docs/tokens.png)

