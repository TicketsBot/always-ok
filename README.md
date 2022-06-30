# What is this?
A simple HTTP server that responds with a 200 to every request.

# Motivation
Health checks that only acknowledge 200 statuses as healthy.

# Building and running
```
go build main.go
HTTP_ADDR=:8080 ./main
```

The HTTP server binds to the address provided via the `HTTP_ADDR` environment variable.

# Docker Images
A docker image is available at `ghcr.io/ticketsbot/always-ok:latest`