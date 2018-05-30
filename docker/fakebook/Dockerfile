FROM golang:latest

ARG CONTEXT

RUN mkdir /app
ADD $CONTEXT/fakebook/main/ /app/
WORKDIR /app
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get github.com/gin-gonic/contrib/sessions
RUN go get golang.org/x/crypto/bcrypt

RUN go build -o main .
CMD ["/app/main"]

EXPOSE 8081