## docker related items...

### The --remove-orphans flag removes containers we may not have defined in our
### docker-compose, but are still present from old configs.
docker compose down --remove-orphans
docker compose up -d

### start and force rebuild...
docker compose up --build 


### install tailwindcss locally...
npm install -D tailwindcss
npx tailwindcss init
code tailwind.config.js

### run/build tailwindcss locally...
npx tailwindcss -i ./styles.css -o ../assets/styles.css --watch

### install tailwind in Docker...
RUN npm init -y && \
    npm install tailwindcss && \
    npx tailwindcss init

### run/build tailwind in Docker...
CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css --watch --poll

### run tailwind in Docker manually...
docker compose run tailwind npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css
