-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY id;

-- name: CreateAuthor :one
INSERT INTO authors (bio, birth_year)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
SET
    bio = coalesce(sqlc.narg('bio'), bio),
    birth_year = coalesce(sqlc.narg('birth_year'), birth_year)
WHERE id = sqlc.arg('id');

-- name: GetBioForAuthor :one
SELECT bio FROM authors
WHERE id = $1;

-- name: GetInfoForAuthor :one
SELECT bio, birth_year FROM authors
WHERE id = $1;

-- name: ListAuthorsByIDs :many
SELECT * FROM authors
WHERE id = ANY($1::int[]);