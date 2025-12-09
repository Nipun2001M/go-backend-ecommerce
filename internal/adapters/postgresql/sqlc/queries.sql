-- name: ListProducts :many
SELECT * from products; 

-- name: FindProductById :one
SELECT * FROM products WHERE id=$1;