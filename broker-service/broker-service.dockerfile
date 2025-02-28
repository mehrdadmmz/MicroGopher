# base go image
FROM golang:1.24-alpine as builder

RUN mkdir /app

# copy the source code into the container
COPY broker-service/ /app

# set the working directory
WORKDIR /app

# build the go app, we will not be using any C libraries
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api 

# just to make sure it is executable
RUN chmod +x /app/brokerApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]

