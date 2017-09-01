## Introduction
```
In this documentation we illustrate how to create a web-host running on the docker.
```

## Run
```
go run main.go
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
