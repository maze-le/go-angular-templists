# Go+Angular temperature lists

This project collects temperature values from a certain source, stores them persistently in a backend and finally presents them in a angular web application.

## Requirements

In order to build the project you need the following

- docker [www.docker.com](https://www.docker.com/)
- docker-compose [docs.docker.com/compose](https://docs.docker.com/compose/)

## Project organization

The project is organized in 2 seperate docker containers, bundled with docker-compose:

- A backend, written in golang. The source folder is: `src/backend`.
- a web frontend written in angular.js found in: `src/client`. The docker container is configured to start a small nginx webserver that delivers the web application.

## Build

...

## Run locally

...

