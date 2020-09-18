FROM golang:1.15.1

LABEL maintainer Mathias Vonende <mathias@vonende.de>
LABEL version="1.0"

# create container user and group
RUN useradd --user-group --system --create-home --no-log-init backend
USER backend

RUN unset GOPATH
WORKDIR /home/backend/

# copy source files
COPY . .

# install dependencies
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/jinzhu/gorm
RUN go get -d -v github.com/julienschmidt/httprouter

# expose port 8082
EXPOSE 8082

# run backend
CMD ["go", "run", "main.go"]
