-- name: CreateMessage :one
INSERT INTO messages(id, created_at, payload, device_uid)
VALUES($1, $2, $3, $4
      )
    RETURNING *;

-- name: GetMessageByID :one
SELECT * FROM messages WHERE id = $1;

-- name: GetMessageByDeviceUID :one
SELECT * FROM messages WHERE device_uid = $1;

-- name: GetMessages :many
SELECT * FROM messages
ORDER BY created_at ASC;

-- name: GetPagedMessages :many
SELECT * FROM messages ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: GetMessagesCount :one
SELECT COUNT(*) AS total FROM messages;

-- name: DeleteMessage :exec
DELETE FROM messages WHERE id = $1;