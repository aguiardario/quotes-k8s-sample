FROM golang:1.13

WORKDIR /go/src/app

COPY . .

COPY go.mod ./
RUN go mod download

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

# VOLUME ["/go/src/app/files"]

CMD ["quotes"]