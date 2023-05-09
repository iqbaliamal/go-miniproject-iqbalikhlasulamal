## Scholarship API

REST API application to manage and publish information of Scholarship. Written in Go with Fiber Framework and use PostgreSQL Databases.

## Notes

There are two branch in this repository:

- `master`: production branches.
- `develop`: development branches.

## How to use

1. Clone this repository.

2. Config the .env variables.

```sh
nano .env
```

3. Fill the database configurations inside the `.env` file.

4. Create a new database.

```sql
CREATE DATABASE miniproject;
```

5. Run the application. Make sure the database is online.

```sh
go run main.go
```
