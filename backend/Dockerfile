FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build


COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .


RUN go build -o main .


WORKDIR /dist

ENV env=prod

RUN cp /build/main .
COPY .env /dist
COPY .migrations /dist/migrations

EXPOSE 8080


CMD /dist/main -env $env