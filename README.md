# ArchMark

ArchMark is a bookmark manager that archives the bookmarked page using [Monolith](https://github.com/Y2Z/monolith). It consists of the main web proccess as well as a worker that downloads the web page for archival.

## Installation

The simplest way to use ArchMark would be to use docker-compose. If you have Docker and Docker Compose installed, you can follow the steps below to get up and running:

1. `git clone https://github.com/sparkymat/archmark`
2. `cd archmark`
3. `docker-compose up`

Alternatively, you can
1. `mkdir archmark`
2. `cd archmark`
3. `wget https://raw.githubusercontent.com/sparkymat/archmark/main/docker-compose.prod.yml -o docker-compose.yml`
4. `mkdir -p data/db`
5. `mkdir -p data/faktory`
6. `mkdir -p data/archive`
7. `docker-compose up`

## Screenshots

![list](https://raw.githubusercontent.com/sparkymat/archmark/master/docs/list.png)

![add](https://raw.githubusercontent.com/sparkymat/archmark/master/docs/add.png)

