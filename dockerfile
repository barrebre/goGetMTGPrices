FROM golang:alpine
WORKDIR /go/src/github.com/barrebre/goGetMTGPrices
ADD . /go/src/github.com/barrebre/goGetMTGPrices
RUN go build -ldflags '-linkmode external' -o main .
CMD ["./main"]