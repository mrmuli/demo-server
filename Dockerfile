FROM golang:latest

COPY src/http-demo /app/http-demo

WORKDIR /app

EXPOSE 3030

ENTRYPOINT ["/app/http-demo"]