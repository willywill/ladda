FROM golang:1.16.0-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
# TODO: Make this configurable
ENV PORT=8080
RUN go clean --modcache
RUN go build -v -o ladda-linux-amd64 ./cmd/main.go
CMD ["/app/ladda-linux-amd64"]
EXPOSE 8080