-- name: Createproduct :one
INSERT INTO products(
    name,
    description,
    price
    ) VALUES (
        $1, $2, $3
    )
RETURNING *;