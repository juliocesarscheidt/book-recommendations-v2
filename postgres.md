# Postgres

```bash
docker-compose exec postgres sh -c "PG_PASSWORD=postgres psql -U postgres -p 5432 -h 127.0.0.1 --dbname=book_db"


CREATE EXTENSION hstore;

# connect
\c book_db

# list databases
\l

# list tables
\dt

SELECT * FROM "book";
TRUNCATE "book";

# list schemas
\dn

# list roles
\dg

# list tables, views, and sequences
\d

# list tablespaces
\db

# list views
\dv

# quit
\q

```
