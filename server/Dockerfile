FROM golang:1.12.0-alpine3.9
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN apk update && apk add git

RUN go build -o main .

EXPOSE 9999

CMD ["go","run","main.go"]

