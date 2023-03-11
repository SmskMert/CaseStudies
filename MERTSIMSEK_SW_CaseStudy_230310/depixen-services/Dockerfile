FROM golang

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /depixen ./cmd/server/main.go

EXPOSE 8080

CMD [ "/depixen" ]