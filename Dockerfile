FROM golang:alpine
WORKDIR /iguana_api

# add some necessary packages
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

# prevent the re-installation of vendors at every change in the source code
ADD . .

RUN go mod download && go mod verify

# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go install github.com/githubnemo/CompileDaemon@latest


ARG VERSION="4.14.1"

RUN set -x \
    && apk add --no-cache git \
    && git clone --branch "v${VERSION}" --depth 1 --single-branch https://github.com/golang-migrate/migrate /tmp/go-migrate

WORKDIR /tmp/go-migrate

RUN set -x \
    && CGO_ENABLED=0 go build -tags 'mysql' -ldflags="-s -w" -o ./migrate ./cmd/migrate \
    && ./migrate -version

RUN cp /tmp/go-migrate/migrate /usr/bin/migrate

WORKDIR /iguana_api

ENTRYPOINT CompileDaemon -command="./iguana_api"
# ENTRYPOINT [ "sh", "/entrypoint.sh" ]
