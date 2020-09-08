FROM nginx:1.19.2

LABEL maintainer Mathias Vonende <mathias@vonende.de>

# create container user and group
# RUN useradd --user-group --system --create-home --no-log-init webserver
# USER webserver

# configure nginx docker
COPY ./nginx.conf /etc/nginx/nginx.conf

# add html entrypoint
COPY ./index.html /usr/share/nginx/html

# expose port 8080
EXPOSE 8080

# start nginx not daemonized
CMD ["nginx", "-g", "daemon off;"]
