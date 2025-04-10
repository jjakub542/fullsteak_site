# My website

Backend template for blog websites

## How to run locally

1. Install Go 1.22
2. Install and configure PostgreSQL (create user and db for project, grant permissions to user)
3. Clone repository
4. Add .env with PORT, APP_ENV, DB_USERNAME, DB_PASSWORD, DB_PORT, DB_HOST, DB_NAME
5. Run with commands below:

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

create mock articles
```bash
make articles
```

drop tables in postgresql
```bash
make drop
```

create admin user
```bash
make superuser
```

make tables in postgresql
```bash
make tables
```

clean up binary from the last build
```bash
make clean
```

## Useful commands:

Connect to psql as postgres:
```bash
sudo -u postgres psql
```

Then:
```sql
CREATE DATABASE fullsteak_db;
CREATE DATABASE fullsteak_db_test;
CREATE USER fullsteak_admin WITH PASSWORD '123';
GRANT ALL PRIVILEGES ON DATABASE fullsteak_db TO fullsteak_admin;
GRANT ALL PRIVILEGES ON DATABASE fullsteak_db_test TO fullsteak_admin;
ALTER USER fullsteak_admin WITH SUPERUSER;
```


Connect to my_website_db as my_website_admin:
```bash
psql -h localhost -d fullsteak_db -U fullsteak_admin -p 5432
```