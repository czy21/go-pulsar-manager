FROM nginx:1.21.6-alpine

COPY ./web/build/ /usr/share/nginx/web/
COPY ./api/build/ /app/