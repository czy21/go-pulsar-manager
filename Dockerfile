FROM nginx:1.21.6

RUN echo "/app/main" > /docker-entrypoint.d/01-api.sh
RUN chmod +x /docker-entrypoint.d/40-api.sh
COPY ./nginx-app.conf /etc/nginx/conf.d/default.conf
COPY ./web/build/ /usr/share/nginx/web/
COPY ./api/build/ /app/