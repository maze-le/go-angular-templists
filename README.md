# Go+Angular temperature lists

This project collects temperature values from a open weather map, stores city references in a backend and finally presents them in a angular web application.

## Requirements

In order to build the project you need docker [www.docker.com](https://www.docker.com/) with docker-compose [docs.docker.com/compose](https://docs.docker.com/compose/)

## Project organization

The project is organized in 3 seperate docker containers, bundled with docker-compose:

- A postgersql database
- A backend, written in golang. The source folder is: `src/backend`.
- a web frontend written in angular.js found in: `src/client`.

## Configuration

In order to access the open weather map API you need an API key. Open the file `docker-compose.yml` and edit the line containing:

` OWM_ACCESS_KEY: <secret>`

and replace `<secret>` with your own API key.

## Build & Run

`docker-compose -f docker-compose.yml up`

## Implemetation Details

### DB

The database is a default postgres image from docker hub without any custom configuration.

### Backend 

The backend is a Model-Repository-Controller backend, implemented with `gorm` as ORM Layer to the database. It encapsulates transport logic in the controllers, found in: `src/controllers`, provides access to the (postgres) database with repositories in: `src/repositories` and describes the database entities with model-interfaces found in: `src/entities`. External services (open weather map) are implemented in: `src/services`.

The backend implements a custom error handler and a unified logging facilty found in: `src/middleware`. The controller and repo parts are CRUD services that implement the respective Index, Create, Read, Update and Delete opertaions.

#### init and shutdown

When the application starts, the database connection gets established. Then a signal handler is initialized to a ensure clean shutdown procedure on `SIGTERM` events: basically a database disconnect, that enforces writing or rolling back pending transactions. It then proceeds to start a http server on the port 8082.

