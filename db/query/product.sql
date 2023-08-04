-- name: Createproduct :one
INSERT INTO products(
    price,
    description,
    price
    ) VALUES (
        $1, $2, $3
    )
RETURNING *;