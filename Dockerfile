# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Set env variables
ENV PORT 8082

ADD . /go/src/github.com/shatvl/flatwindow

# Build the outyet command insi de the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/shatvl/flatwindow

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/flatwindow

# Document that the service listens on port 8082.
EXPOSE 8082

# docker build ./ -t flatwindow
# docker run --publish 8082:8082 --name flatwindow --rm flatwindow