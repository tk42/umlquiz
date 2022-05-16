###############################
# Full Builder container
###############################
FROM golang:alpine AS compiler

WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOPROXY=direct

RUN apk add git openssl openssh gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Compile
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app .

###############################
# Exec container
###############################
FROM alpine:latest

ENV TZ=Asia/Tokyo
WORKDIR /home
COPY --from=compiler /go/src/app /home
RUN mkdir -p /home/api
COPY ./api ./api

# install certificates
# These are needed if your application wants to communicate with the outside world.
# https://petomalina.medium.com/serverless-grpc-with-cloud-run-bab3622a47da
RUN apk update && apk add ca-certificates && update-ca-certificates && rm -rf /var/cache/apk/*

CMD ["sh", "-c", "sleep 4; ./app"]
