FROM golang:latest

WORKDIR /flashcardApi
COPY . /flashcardApi

RUN cd /flashcardApi
RUN env GOOS=linux GOARCH=amd64 go  build -o card

CMD ["/flashcardApi/card"]