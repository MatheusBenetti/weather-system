FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o temp-system .

FROM scratch
WORKDIR /app
COPY --from=builder /app/temp-system .
CMD ["./temp-system"]