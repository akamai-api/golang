### Basic setup with docker-compose


## Instructions
* Make sure docker engine is installed ([instructions](https://docs.docker.com/engine/installation/))
* Make sure docker-compose is installed ([instructions](https://docs.docker.com/compose/install/))

* Start services/containers by running `docker-compose up`

Now we have 3 docker containers running.
* run arbitratry commands in your services using `docker-compose exec <service> <command>` e.g. 
```
docker-compose exec influxDB bash
```
* Recreate containers even if their configuration and image haven't changed using :  
```
docker-compose up --force-recreate
```

## Test stack

The stack can now be tested by running localtest.sh under `tests`
```
cd tests
./localtest.sh
```

Now visit local [grafana](http://localhost:3000) instance and login with admin/admin


## How to access the bash terminal in a certain container
 ```
 docker run -it server:latest bash
 ```
## How to rebuild images before recreating the containers
```
docker-compose up --force-recreate --build
```
