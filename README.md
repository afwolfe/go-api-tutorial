# Go API Tutorial

My implementation of the project from [YT: Tech with Tim's Go API Tutorial](https://www.youtube.com/watch?v=bj77B59nkTQ)

## Modifications

* Replaced the in-memory array of books with a MySQL backend.
* Builds and deploys with a database using Docker Compose.

## Usage

```bash
# Initialize the database first:
docker compose up db -d
mysql --user=root -ppassword --host=127.0.0.1 --port=3306 < initialize-database.sql
docker compose down
# Build and start the DB and API together:
docker compose up --build
```

The API will be available at http://localhost:8080/

## Resources

* [[YT] Tech with Tim: Go API Tutorial](https://www.youtube.com/watch?v=bj77B59nkTQ)
* [Go Docs: Tutorial: Accessing a relational database](https://go.dev/doc/tutorial/database-access)
