# Postgres Notes

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
INSERT INTO users (age, email, first_name, last_name)
VALUES (30, 'bob@aol.com', 'Bob', 'Aol');
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