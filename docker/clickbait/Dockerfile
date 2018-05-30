FROM golang:latest

ARG CONTEXT

RUN mkdir /app
ADD $CONTEXT/clickbait/main/ /app/
WORKDIR /app
RUN go get github.com/gin-gonic/gin
RUN go build -o main .
CMD ["/app/main"]

EXPOSE 8080