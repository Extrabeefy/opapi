FROM golang

COPY . /go/src/github.com/opAPIProgression
WORKDIR /go/src/github.com/opAPIProgression
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 8080