# Docker configuration
In this documentation we illustrate how to create a web-host running on the docker.


## Run
```
go run server.go
```

## Build
```
go build
```

## Run Go Executable
```
./Your-Root-Folder-Name
```

## Build Docker

```
docker build -t go-server .
```

## Run Docker
```
docker run --publish 9143:9143 -t  go-server
```

## Check if your docker is running correctly
```
docker ps
```
## How to remove an image in the Docker
```
https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes
```
