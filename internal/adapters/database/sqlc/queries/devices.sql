-- name: CreateDevice :one
INSERT INTO devices(id, created_at, updated_at, uid, serial)
VALUES($1, $2, $3, $4, $5
      )
    RETURNING *;

-- name: GetDeviceBySerial :one
SELECT * FROM devices WHERE serial = $1;

-- name: GetDeviceByUID :one
SELECT * FROM devices WHERE uid = $1;

-- name: GetDeviceByID :one
SELECT * FROM devices WHERE id = $1;

-- name: GetDevices :many
SELECT * FROM devices
ORDER BY created_at ASC;

-- name: GetPagedDevices :many
SELECT * FROM devices ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: GetDevicesCount :one
SELECT COUNT(*) AS total FROM devices;

-- name: DeleteDevice :exec
DELETE FROM devices WHERE id = $1;
