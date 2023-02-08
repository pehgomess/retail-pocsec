FROM golang:1.19
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o retail-pocsec
RUN CGO_ENABLED=0 GOOS=linux go build -o retail-pocsec
CMD ["/app/retail-pocsec"]