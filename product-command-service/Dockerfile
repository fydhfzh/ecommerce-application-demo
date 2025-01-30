FROM alpine:latest

RUN mkdir /app


WORKDIR /app

COPY productCommandApp /app
COPY config /app/config

CMD [ "/app/productCommandApp" ]