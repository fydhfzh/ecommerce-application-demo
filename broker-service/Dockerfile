FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY brokerApp /app
COPY config /app/config

CMD ["/app/brokerApp"]