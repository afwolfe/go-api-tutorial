FROM golang:alpine as builder

COPY go.mod go.sum *.go /src/
WORKDIR /src

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o app

FROM alpine:3 as app
WORKDIR /app
COPY --from=builder /src/app ./

ENV PORT=8080
EXPOSE 8080
CMD [ "./app" ]