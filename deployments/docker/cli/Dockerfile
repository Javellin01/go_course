FROM golang:alpine AS CLI
ENV USER user
ENV HOME /home/${USER}

WORKDIR /go/src/service

RUN apk add --update sudo
RUN adduser -D $USER \
        && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER
USER $USER

RUN sudo apk add build-base
RUN sudo apk add protobuf-dev
RUN sudo apk add protoc

RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
RUN go get github.com/go-swagger/go-swagger/cmd/swagger
RUN go get github.com/joho/godotenv/cmd/godotenv
