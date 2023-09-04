### Docker notes

Start docker image
```bash
docker compose up [-d] 
```

Stop docker image
```bash
docker compose down
```

Stop docker image & remove dropped orphans
```bash
docker compose down
```

Rebuild docker image with specific compose files
```bash
docker compose -f docker-compose.yml -f docker-compose.production.yml up --build
```

Remove specific image, server in this instance
```bash
docker compose -f docker-compose.yml -f docker-compose.production.yml rm server
``````


The --remove-orphans flag removes containers we may not have defined in our
docker-compose, but are still present from old configs.
```bash
docker compose down --remove-orphans
docker compose up -d
```


List Docker images
```
docker image ls
```


List Docker containers
```
docker container ls
```

Remove specific image
```
docker image rm -f [image name or IMAGEID]
```

Docker kill container...
```
docker kill [container name]
```


------

In Prod ubuntu, in order to stop container, had to first (was getting permission denied)
```
sudo aa-remove-unknown
sudo systemctl restart docker.service

```
source: https://superuser.com/questions/1447183/docker-container-not-stopping