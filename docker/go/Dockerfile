FROM golang:1.19-alpine

LABEL maintainer "morikawa"

ENV MYSQL_URL=mysql://root:root@tcp(mysql:3306)/go-queue

# golangci-lintをインストール golangのlinter
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1

# ワーキングディレクトリの設定
WORKDIR /var/www/src

# アップデートとgitのインストール
RUN apk update 
RUN apk add git
RUN go mod init go-queue
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest
RUN go install golang.org/x/tools/cmd/godoc@latest

CMD ["air", "-c", ".air.toml"]