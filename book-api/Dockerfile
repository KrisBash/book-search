FROM golang:1.9-alpine

EXPOSE 8081

WORKDIR /go/src/app
RUN apk --no-cache add curl git && \
    curl https://glide.sh/get | sh

COPY . .
RUN chmod 755 ./src/pkgs.sh && ./src/pkgs.sh
WORKDIR /go/src/app/src
RUN go build -o book-api

CMD /go/src/app/src/book-api

