FROM alpine:3.16
WORKDIR /app/

COPY ./web/build/ /app/dist/
COPY ./api/build/ /app/

RUN chmod +x /app/main

CMD ["main"]