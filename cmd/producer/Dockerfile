FROM golang:1.23

WORKDIR /producer

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o producer ./main.go ./producer.go ./producer_utills.go ./model.go

CMD ["./producer"]
