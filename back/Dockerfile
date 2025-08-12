FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download
RUN go mod verify

COPY . .

RUN go build main.go

EXPOSE 5000

CMD [ "./main" ]
