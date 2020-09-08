FROM golang:1.15.1

LABEL maintainer Mathias Vonende <mathias@vonende.de>

# create container user and group
RUN useradd --user-group --system --create-home --no-log-init backend
USER backend

ADD httpserver.go .

# expose port 8080
EXPOSE 8081

# start nginx not daemonized
CMD ["go", "run", "httpserver.go"]
