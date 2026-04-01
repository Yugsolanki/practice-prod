FROM golang:1.25-alpine

WORKDIR /app

COPY images ./images
COPY go.mod .
COPY index.htm .
COPY main.go .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]