FROM alpine:latest

WORKDIR /

ADD main main

EXPOSE 8080

ENTRYPOINT ["./main"]