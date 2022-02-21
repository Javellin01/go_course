FROM golang:latest AS app
WORKDIR /go/src/app
#COPY go.sum .
COPY go.mod .
RUN go mod download