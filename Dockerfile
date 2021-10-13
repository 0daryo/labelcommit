FROM golang:1.16-alpine
RUN go build -o /bin/app
ENTRYPOINT [ "/bin/app" ]