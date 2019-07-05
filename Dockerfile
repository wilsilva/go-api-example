FROM golang:1.12-alpine
RUN apk add git
RUN go get github.com/tockins/realize
ADD . /application
WORKDIR /application
EXPOSE 8000
CMD ["go", "run", "cmd/api/main.go"]"