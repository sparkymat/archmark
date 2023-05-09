# ArchMark [![.github/workflows/ci.yml](https://github.com/sparkymat/archmark/actions/workflows/ci.yml/badge.svg)](https://github.com/sparkymat/archmark/actions/workflows/ci.yml)

ArchMark is a bookmark manager that archives the bookmarked page using [Monolith](https://github.com/Y2Z/monolith). It consists of the main web proccess as well as a worker that downloads the web page for archival.

## Features

- Allows links to be bookmarked and categorized
- Support local (password-based) login and registration, and reverse-proxy-based authentication (with forwarded headers)
- Caches a local copy of the bookmarked page (available via the 'cached' link below the original)
- Keeps deleted bookmarks in the "deleted" section for a period of time before deleting it (defaults to 48 hours); they can be restored before the timer hits
- Allows searching across content in the bookmarked links (a copy of the linked page is cached for search indexing)

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

| Variable                       | Description                                                                                   |
| ------------------------------ | --------------------------------------------------------------------------------------------- |
| `DISABLE_REGISTRATION`         | Set to 'true' if you want to disable registrations                                            |
| `REVERSE_PROXY_AUTHENTICATION` | Enables reverse proxy authentication via forwarded headers                                    |
| `PROXY_AUTH_USERNAME_HEADER`   | The request header where proxy passes the usernamae. Default: `Remote-User`                   |
| `PROXY_AUTH_NAME_HEADER`       | The request header where proxy passes the full name. Default: `Remote-Name`                   |
| `DELETE_TIMER_HOURS`           | Time (in hours) before deleted items are flushed from the recycle bin (**default**: 48 hours) |

## Screenshots

### Login page

![login](/docs/images/login.png)

### Bookmarks

![bookmarks](/docs/images/list.png)

### New bookmark

![new](/docs/images/new.png)

### Change category

![change_category](/docs/images/change_category.png)

### Deleted

![deleted](/docs/images/recycle.png)
