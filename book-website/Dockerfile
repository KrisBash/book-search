FROM nginx:latest
EXPOSE 80
COPY build/ /usr/share/nginx/
RUN rm /etc/nginx/conf.d/*.conf
COPY nginx.conf /etc/nginx/conf.d/book-site.conf
