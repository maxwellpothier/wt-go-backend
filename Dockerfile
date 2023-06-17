# Build the Go binary
FROM golang:1.19 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create the minimal runtime image for the Go binary
FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/main .

# Run the Go binary
CMD ["./main"]