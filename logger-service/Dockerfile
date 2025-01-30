FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY loggerApp /app
COPY config /app/config

CMD ["/app/loggerApp"]