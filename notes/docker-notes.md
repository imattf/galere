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

Rebuild docker image
```bash
docker compose -f docker-compose.yaml -f docker-compose.production.yaml up --build
```

Start but remove server first
```bash
docker compose -f docker-compose.yml -f docker-compose.production.yml rm server
``````


The --remove-orphans flag removes containers we may not have defined in our
docker-compose, but are still present from old configs.
```bash
docker compose down --remove-orphans
docker compose up -d
```
