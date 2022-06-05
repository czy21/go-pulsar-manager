FROM nginx:1.21.6

COPY ./web/build/ /usr/share/nginx/web/
COPY ./api/build/ /app/