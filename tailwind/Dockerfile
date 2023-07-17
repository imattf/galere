# Start with an image with node installed
FROM node:latest

# Create the directories we need
RUN mkdir /tailwind

# Set /tailwind as the workdir. 
# A workdir is required for npm to work correctly
WORKDIR /tailwind

# Install tailwindcss and initialize
RUN npm init -y && \
    npm install tailwindcss && \
    npx tailwindcss init

# Run tailwindcss. 
# This will watch changes in /src/sytles.css and out to /dst/styles.css
CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css --watch --poll