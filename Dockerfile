FROM golang:1.16-alpine
COPY go.mod ./
COPY go.sum ./
COPY cmd cmd
ENTRYPOINT [ "go","run","cmd/main.go" ]