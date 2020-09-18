FROM nginx:1.19.2

LABEL maintainer Mathias Vonende <mathias@vonende.de>
LABEL version="1.0"

# configure nginx docker
COPY ./mime.types /etc/nginx/mime.types
COPY ./nginx.conf /etc/nginx/nginx.conf

# copy bundled sources entrypoint
COPY ./dist/* /usr/share/nginx/html/

# expose port 8080
EXPOSE 8080

# start nginx not daemonized
CMD ["nginx", "-g", "daemon off;"]
