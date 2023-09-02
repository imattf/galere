## docker related items...

The --remove-orphans flag removes containers we may not have defined in our
docker-compose, but are still present from old configs.
```bash
docker compose down --remove-orphans
docker compose up -d
```

Start and force rebuild...
```bash
docker compose up --build 
```

Install tailwindcss locally...
```bash
npm install -D tailwindcss
npx tailwindcss init
code tailwind.config.js
```

To run/build tailwindcss locally...
```bash
npx tailwindcss -i ./styles.css -o ../assets/styles.css --watch
```

To install tailwind in Docker...
```bash
RUN npm init -y && \
    npm install tailwindcss && \
    npx tailwindcss init
```

To run/build tailwind in Docker...
```bash
CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css --watch --poll
```

To run tailwind in Docker manually...
```bash
docker compose run tailwind npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css
```
