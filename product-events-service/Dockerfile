FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY productEventsApp /app
COPY config /app/config

CMD [ "/app/productEventsApp" ]