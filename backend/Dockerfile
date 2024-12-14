FROM golang:latest AS builder
WORKDIR /app/backend
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
COPY ./backend .
RUN CGO_ENABLED=0 go build -o main .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/backend/main /
COPY --from=builder /app/backend/.env /

EXPOSE 8000
CMD ["./main"]
