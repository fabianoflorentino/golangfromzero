# Database Migrations

This project uses [`golang-migrate/migrate`](https://github.com/golang-migrate/migrate) to handle database migrations. The migration CLI is installed automatically inside the development Docker container.

All migration scripts are located in the `database/migrations` directory.

## Running Commands via Docker Compose

Since the CLI is installed inside the `app` container, you should prefix your migration commands with `docker compose exec app`.

The application connects to the database via the `DATABASE_URL` environment variable, which is automatically populated from your `.env` file inside the container.

### 1. Create a New Migration

To create a new migration pair (`.up.sql` and `.down.sql`), run the following command:

```bash
docker compose exec app migrate create -ext sql -dir database/migrations -seq <migration_name>
```

For example, to create a migration for adding a posts table:
```bash
docker compose exec app migrate create -ext sql -dir database/migrations -seq create_posts_table
```

This will generate two files in `database/migrations/`:
- `<sequence>_create_posts_table.up.sql`
- `<sequence>_create_posts_table.down.sql`

Write your `CREATE`, `ALTER`, or `INSERT` statements in the `.up.sql` file, and the corresponding `DROP` or rollback statements in the `.down.sql` file.

### 2. Run Migrations (Up)

To apply all pending migrations and update your database schema to the latest version, run:

```bash
docker compose exec app sh -c 'migrate -path database/migrations -database "$DATABASE_URL" up'
```

### 3. Rollback Migrations (Down)

To rollback the last applied migration, you can step down:

```bash
docker compose exec app sh -c 'migrate -path database/migrations -database "$DATABASE_URL" down 1'
```

To rollback *all* migrations (drop all tables/state handled by migrations), run:
```bash
docker compose exec app sh -c 'migrate -path database/migrations -database "$DATABASE_URL" down'
```

### 4. Force a Version (Fixing Dirty State)

If a migration fails halfway (e.g. syntax error in SQL), your database will be flagged as "dirty". You must fix the SQL file, then force the migration version to the one that failed, and try again.

```bash
docker compose exec app sh -c 'migrate -path database/migrations -database "$DATABASE_URL" force <version_number>'
```
*(Replace `<version_number>` with the sequence number of the failed migration, e.g., `1`).*

## Troubleshooting

- **Check connection:** Ensure your `DATABASE_URL` matches the credentials in `.env`.
- **Docker running:** Your containers must be running (`docker compose up -d`) to use `docker compose exec`.
