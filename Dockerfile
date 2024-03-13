FROM golang:1.22 AS builder

RUN apt-get update && apt-get install -y \
  make \
  && rm -rf /var/lib/apt/lists/*

COPY . /app/

WORKDIR /app
RUN go generate ./...
RUN make archmark

FROM debian:12-slim

RUN apt-get update && apt-get install -y \
  ca-certificates \
  ffmpeg \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/archmark /bin/archmark

WORKDIR /app
COPY public /app/public
COPY migrations /app/migrations

CMD ["/bin/archmark"]
