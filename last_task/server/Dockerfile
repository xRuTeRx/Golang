FROM golang:1.16

WORKDIR /go/server
COPY . .
RUN go mod download
RUN go build -o ./go/server .

EXPOSE 8080
CMD ["./go/server"]



