FROM nginx:latest
EXPOSE 80
COPY build/ /usr/share/nginx/
RUN rm /etc/nginx/conf.d/*.conf
RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y dnsutils
COPY nginx.conf /etc/nginx/conf.d/book-site.conf
