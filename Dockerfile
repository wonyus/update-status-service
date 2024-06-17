# syntax=docker/dockerfile:1

FROM golang:1.22-alpine as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY contexts/* ./contexts/
COPY controllers/* ./controllers/
COPY pkg/client/* ./pkg/client/
COPY pkg/db/* ./pkg/db/
COPY utils/* ./utils/


# Build
RUN GOOS=linux go build -o /update-status

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose

FROM golang:1.22-alpine

# WORKDIR /app

COPY --from=builder /update-status /update-status

# RUN adduser -D -g '' appuser && chown -R appuser:appuser /app

# USER appuser

# Run
CMD ["/update-status"]