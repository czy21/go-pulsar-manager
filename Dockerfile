FROM nginx:1.21.6

RUN echo "/app/main /dev/null 2>&1 &" > /docker-entrypoint.d/01-api.sh
RUN chmod +x /docker-entrypoint.d/01-api.sh
COPY nginx/app.conf /etc/nginx/conf.d/default.conf

COPY ./web/build/dist.tar.gz /usr/share/nginx/web/
RUN tar -zxvf /usr/share/nginx/web/dist.tar.gz -C /usr/share/nginx/web/ && rm -rf /usr/share/nginx/web/dist.tar.gz
COPY ./api/build/ /app/