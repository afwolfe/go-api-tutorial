FROM golang:alpine as builder

COPY go.mod go.sum *.go /src/
WORKDIR /src

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o app

FROM alpine:3 as app
WORKDIR /app

RUN --mount=type=cache,target=/var/cache/apk \
    apk update && \
    apk add bash

COPY scripts/wait-for-it.sh ./
COPY --from=builder /src/app ./

ENV PORT=8080
EXPOSE 8080
CMD ["./wait-for-it.sh", "--host=db", "--port=3306", "--timeout=60", "--", "./app"]