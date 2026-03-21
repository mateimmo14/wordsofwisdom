-- name: AddWisdom :one
INSERT INTO wisdoms (data, author)
values (
        $1,
        $2
       )
returning *;