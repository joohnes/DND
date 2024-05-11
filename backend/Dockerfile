FROM golang:latest AS backend
WORKDIR /app/backend
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
COPY ./backend .
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]
