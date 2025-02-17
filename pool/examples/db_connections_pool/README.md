# DB Connection Pool Demo

Hey! This is a simple demo showing why connection pools are awesome for database operations. We'll run some queries with and without a pool to see the difference in action.

## What you'll need

- Go 1.23.5+
- Docker with Docker Compose

## Getting Started

1. Fire up Postgres:
   ```bash
   docker compose up -d
   ```

2. The DB will be running at:
   - localhost:5432
   - user: postgres
   - password: postgres
   - database: postgres

## Running the Demo

The demo comes with some handy flags to play around with:

- `--with-pool`: Turn pooling on/off (defaults to on)
- `--pool-size`: How many connections to keep in the pool (default: 100)
- `--max-queries`: Number of queries to run at once (default: 102)

### Try these out:

1. Build the binary:
   ```bash
   go build .
   ```

2. Default run (with pool):
   ```bash
   ./dbpool
   ```

3. Without pool:
   ```bash
   ./dbpool --with-pool=false --max-queries=102
   ```

4. Mix it up:
   ```bash
   ./dbpool --pool-size=50 --max-queries=200
   ```

## What's happening?

We're running a bunch of `SELECT pg_sleep(1)` queries in parallel. Each query takes about a second to finish.

- No pool: DB reaches max connections limit, and we start getting connection errors. You can reduce the number of queries to less than 100 to see the error go away.
- With pool: Reuses connections from a pool, no matter how many queries you run, if you set the pool size to 100, you'll never max connection error.

