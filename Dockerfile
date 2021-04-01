FROM golang:1.16.0-alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
ENV PORT=8080
RUN go clean --modcache
RUN go build -v -o ladda-linux-amd64 ./cmd/main.go
RUN ls && exit 1
ENTRYPOINT ["/app/ladda-linux-amd64"]
EXPOSE 8080