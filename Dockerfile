FROM golang:1.22

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GIN_MODE=release
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 80

CMD ["./main"]
