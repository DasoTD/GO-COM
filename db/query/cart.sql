-- name: CreateCart :one
INSERT INTO carts (
  owner,
  product,
  quantity
) VALUES (
  $1, $2, $3
) RETURNING *;