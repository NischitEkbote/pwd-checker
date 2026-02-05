FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go helper.go ./

RUN go build -o pwd-checker .

ENTRYPOINT ["./pwd-checker"]
