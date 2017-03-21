# golang-todo

### Overview

Example of todo application with next technology stack:
- Go(1.8.0) programming language - https://golang.org/
- Chi router - https://github.com/pressly/chi
- glide package manager - https://github.com/Masterminds/glide
- RethinkDB(2.3.5) - https://www.rethinkdb.com/
- Docker(17.03.0-ce) - https://github.com/docker/docker
- Docker-Compose(1.11.2) - https://github.com/docker/compose

To run RethinkDB and golang application Git, Docker and Docker-Compose should be installed locally.
<br /> Docker install: https://docs.docker.com/engine/installation/
<br /> Docker-Compose install: https://docs.docker.com/compose/install/

### Usage
Open terminal and clone this repository by SSH(or HTTPS):
```
    git clone git@github.com:tokillamockingbird/golang-todo.git
```
Locate to golang-todo directory.
To start all needed services simply execute one command
```
    docker-compose up -d
```
After this you can check RethinkDB admin panel on 8080 port, open `localhost:8080` in browser to check it.
<br /> Golang application will be available on 8000 port, open `localhost:8000` to see hello message.
