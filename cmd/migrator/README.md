# migrations example

To run migrations:

```shell
go run . db migrate
```

To rollback migrations:

```shell
go run . db rollback
```

To view status of migrations:

```shell
go run . db status
```

To create a Go migration:

```shell
go run . db create_go go_migration_name
```

To create a SQL migration:

```shell
go run . db create_sql sql_migration_name
```
