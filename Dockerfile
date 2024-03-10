FROM golang:1.22-alpine AS builder
ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /app
COPY . /app
RUN ls -al && go build -o main .

FROM alpine:3.13 AS main
RUN apk --no-cache add ca-certificates
RUN apk add bind-tools curl wget

COPY --from=builder /app/main /app/main
COPY ./assets/ /app/assets/
WORKDIR /app
CMD ["./main"]
