FROM golang:1.20-alpine AS builder

RUN apk update && apk add make

COPY . /app/

WORKDIR /app
RUN go generate ./...
RUN make archmark

FROM alpine:3

COPY --from=builder /app/archmark /bin/archmark

CMD ["/bin/archmark"]
