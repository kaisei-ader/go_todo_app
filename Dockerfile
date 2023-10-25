FROM golang:1.20-alpine

WORKDIR /go/src/app

# ソースを/go/src/appへコピー
COPY go.mod go.sum ./

COPY ./cmd/go-todo-app/main.go /go/src/app

ENV GO111MODULE=on

ENV TZ=Asia/Tokyo

ENV DOTENV_PATH=/go/src/app/.env

RUN go mod download

# AIRのインストール
RUN go install github.com/cosmtrek/air@latest

# 不要なpackageがあれば削除する
RUN go mod tidy


CMD ["air", "-c", ".air.toml"]

# RUN CGO_ENABLED=0 go build -o ./app ./cmd/go-todo-app/main.go

# FROM golang:1.18.0-alpine

# WORKDIR /go/src/app

# # ソースを/go/src/appへコピー
# ADD . /go/src/app

# ENV GO111MODULE=on

# ENV TZ=Asia/Tokyo

# ENV DOTENV_PATH=/go/src/app/.env

# RUN go mod download

# # AIRのインストール
# RUN go install github.com/cosmtrek/air@latest

# # 不要なpackageがあれば削除する
# RUN go mod tidy

# EXPOSE 8080

# CMD ["air", "-c", ".air.toml"]