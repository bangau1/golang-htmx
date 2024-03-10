FROM golang:1.22 AS builder
ARG GOOS=linux
ARG GOARCH=amd64
# DISABLE CGO since we're targeting it to run on alpine
ENV CGO_ENABLED=0 

WORKDIR /app
COPY . /app
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod download
RUN make pre-build && go build -o main .

FROM alpine:3.13 AS main
RUN apk --no-cache add ca-certificates
RUN apk add bind-tools curl wget

COPY --from=builder /app/main /app/main
COPY ./assets/ /app/assets/
WORKDIR /app
CMD ["./main"]
