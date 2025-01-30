FROM alpine:latest

RUN mkdir /app
RUN mkdir /app/logs

WORKDIR /app

COPY config /app/config
COPY authApp /app

CMD [ "/app/authApp" ]
