FROM tetafro/golang-gcc
WORKDIR /go/src/github.com/barrebre/goGetMTGPrices
ADD . /go/src/github.com/barrebre/goGetMTGPrices
RUN go build -ldflags '-linkmode external' -o main .
CMD ["./main"]