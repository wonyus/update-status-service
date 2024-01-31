# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY controllers/* ./controllers/
COPY initials/* ./initials/
COPY utils/* ./utils/


# Build
RUN GOOS=linux go build -o /app.bin

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose

FROM golang:1.21-alpine

# WORKDIR /app

COPY --from=builder /app.bin /app.bin

EXPOSE 8080

# RUN adduser -D -g '' appuser && chown -R appuser:appuser /app

# USER appuser

# Run
CMD ["/app.bin"]