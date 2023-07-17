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
