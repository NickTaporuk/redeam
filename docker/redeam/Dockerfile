# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

RUN go install github.com/NickTaporuk/redeam

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/redeam

# Document that the service listens on port 8080.
EXPOSE 7777
