FROM golang:1.18.3-alpine
WORKDIR /app/
COPY ./web/build/ /app/dist/
COPY ./api/build/ /app/

CMD ["main"]