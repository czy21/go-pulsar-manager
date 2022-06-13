FROM golang:1.18.3-bullseye
COPY ./web/build/ /app/dist/
COPY ./api/build/ /app/
RUN chmod +x /app/main
CMD ["/app/main"]