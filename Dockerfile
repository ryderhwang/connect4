#  base image for Go
FROM golang:latest

RUN mkdir /connect4

# Set the Current Working Directory inside the container
WORKDIR /connect4

ADD . /connect4

RUN go build -o main

ENTRYPOINT ["/connect4/main", "sh"]
