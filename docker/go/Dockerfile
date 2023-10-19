FROM golang:latest

COPY src/api/ /go/src/api/

WORKDIR /go/src/api/

RUN go mod tidy \
&& go get github.com/gin-gonic/gin@latest \
&& go get github.com/jinzhu/gorm@latest \
&& go get github.com/go-sql-driver/mysql@latest \
&& go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]