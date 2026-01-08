-- name: SelectAllMigrations :many
SELECT * FROM migration
ORDER BY version ASC;
