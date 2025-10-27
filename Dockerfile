ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go get github.com/kreely/orm-gin/models
RUN go get github.com/kreely/orm-gin/controllers
RUN go get github.com/kreely/orm-gin/migrate
RUN go get github.com/kreely/orm-gin/initialisers

RUN go run migrate/migrate.go

RUN go build -v -o /run-app .


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
