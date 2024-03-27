FROM golang:1.22.1-alpine3.19

WORKDIR /app 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY  . . 

RUN go build -v -o ./ ./cmd/main.go

EXPOSE 8080
CMD [ "./main" ]