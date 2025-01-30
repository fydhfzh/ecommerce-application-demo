FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY userApp /app
COPY config /app/config

CMD [ "/app/userApp" ]