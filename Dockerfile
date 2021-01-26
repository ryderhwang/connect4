FROM golang:latest
RUN mkdir /connect4
ADD . /connect4/
WORKDIR /connect4
RUN go build -o main .
CMD ["/connect4/main"]
