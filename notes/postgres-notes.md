# Postgres Notes

### Docker stuff...

go get https://docs.docker.com/get-docker/

```bash
docker version
# You should see a version output. This may be large.
docker compose version
# Again, you should see a version. This will likely be a shorter output.
```

Create a docker-compose.yml file wi

```bash
code docker-compose.yml
```
```bash
version: "3.9"

services:
  # Our Postgres database
  db: # The service will be named db.
    image: postgres # The postgres image will be used
    restart: always # Always try to restart if this stops running
    environment: # Provide environment variables
      POSTGRES_USER: baloo # POSTGRES_USER env var w/ value baloo
      POSTGRES_PASSWORD: junglebook
      POSTGRES_DB: lenslocked
    ports: # Expose ports so that apps not running via docker compose can connect to them.
      - 5432:5432 # format here is "port on our machine":"port on container"

  # Adminer provides a nice little web UI to connect to databases
  adminer:
    image: adminer
    restart: always
    environment:
      ADMINER_DESIGN: dracula # Pick a theme - https://github.com/vrana/adminer/tree/master/designs
    ports:
      - 3333:8080
```

Run Adminer
http://localhost:3333/


Start docker image

```bash
docker compose up [-d] 
```

Stop docker image

```bash
docker compose down
```


### Connect to Adminer
http://localhost:3333

### Connect to postgres on docker container

```bash
docker exec -it galere-db-1 /usr/bin/psql -U baloo -d lenslocked
```
and \q to quit connection

### Create users table

```sql
DROP TABLE IF EXISTS users;
```

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
```

### Insert records

```sql
INSERT INTO users (age, email, first_name, last_name) VALUES (30, 'bob@aol.com', 'Bob','Aol');
```
### Comments in SQL

```sql
-- comments in sql
```

### Update records

```sql
UPDATE users
SET first_name = 'Anonymous', last_name = 'Teenager'
WHERE age < 20 AND age > 12;
```

### Delete records
```sql
DELETE FROM users
WHERE id = 1;
```

# Migration stuff...

Need to clear out the existing database before running?
```
docker compose down
docker compose up -d
```

## goose...

command format...
goose <database-type> <"host=connect-string"> subcommand

### this builds the goose tables in the db if not there at status check...
```
goose postgres "host=localhost port=5432 user=baloo password=junglebook dbname=lenslocked sslmode=disable" status
```

### this applies the goose changes in the db...
```
goose postgres "host=localhost port=5432 user=baloo password=junglebook dbname=lenslocked sslmode=disable" up
```

### this undoes the goose changes in the db...
```
goose postgres "host=localhost port=5432 user=baloo password=junglebook dbname=lenslocked sslmode=disable" down
```
--- 

### Jon's goose workflow

My Workflow with Goose

My typical workflow is to keep the timestamps while developing a feature, then when I am ready to submit changes I will:

    goose down or goose reset
    git stash to stash my current changes
    git pull origin <main branch> --rebase to rebase my branch with any changes other devs have submitted.
    git stash apply to apply the changes i stashed
    goose fix to rename my migrations with new versions
    goose up and verify migrations work with new versions
    commit to main branch

Reviews may slow

Doing this will help ensure that you wait until the last possible minute to set a fixed version number and reduce the odds of conflicts.

Goose suggests a slightly different approach of hybrid versioning that can also be explored.
