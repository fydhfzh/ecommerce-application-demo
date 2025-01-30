FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY productQueryApp /app
COPY config /app/config

CMD [ "/app/productQueryApp" ]