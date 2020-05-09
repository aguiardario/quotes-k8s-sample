FROM golang:1.13 as builder

WORKDIR /app
COPY ./sources/main.go .

COPY go.mod ./
RUN go mod download

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app ./main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/app .
CMD ["./app"]