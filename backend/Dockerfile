FROM golang:latest AS builder
WORKDIR /app/backend
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/backend/main /

EXPOSE 8000
CMD ["./main"]
