FROM golang:alpine
WORKDIR /build
COPY main.go .
RUN go build -o url-shortner main.go
CMD ["./server"]
EXPOSE 8080
# CMD ["go run ./url-shortner/cmd/server"]
