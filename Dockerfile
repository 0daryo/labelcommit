FROM golang:1.16-alpine
COPY cmd ./
ENTRYPOINT [ "go","run","cmd/main.go" ]