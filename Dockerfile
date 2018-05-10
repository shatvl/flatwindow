# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

RUN mkdir -p /go/src/flatwindow-backend
WORKDIR /go/src/flatwindow-backend
COPY . /go/src/flatwindow-backend

# Download and install third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install

# Set the PORT environment variable
ENV PORT 8082

# Expose port 8082 to the Host so that outer-world can access your application
EXPOSE 8082

# Tell Docker what command to run when the container starts
CMD ["go-wrapper", "run"]

# docker build ./ -t flatwindow.
# docker run --publish 6060:8082 --name flatwindow --rm flatwindow