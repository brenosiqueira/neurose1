FROM golang:1.7

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

EXPOSE 8080

COPY ./app/golang-foreground /usr/local/bin/
RUN chmod +x /usr/local/bin/golang-foreground

ENTRYPOINT ["/usr/local/bin/golang-foreground"]