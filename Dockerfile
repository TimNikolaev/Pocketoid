FROM golang:1.24-alpine3.21 AS builder

COPY . /github.com/TimNikolaev/Pocketoid/
WORKDIR /github.com/TimNikolaev/Pocketoid/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/TimNikolaev/Pocketoid/bin/bot .
COPY --from=0 /github.com/TimNikolaev/Pocketoid/configs configs/

EXPOSE 8888

CMD [ "./bot" ]

