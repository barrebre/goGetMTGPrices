FROM golang:alpine
WORKDIR /go/src/github.com/barrebre/goGetMTGPrices
ADD . /go/src/github.com/barrebre/goGetMTGPrices
RUN go build -o main .
CMD ["./main"]