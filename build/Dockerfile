FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./web ./web
COPY ./cmd ./cmd
COPY ./configs ./configs
COPY ./internal ./internal
COPY ./config.yaml ./config.yaml


ENV GIN_MODE=release

CMD ["go", "run", "cmd/timespend-server/main.go"]
