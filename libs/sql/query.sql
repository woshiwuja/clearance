-- name: SearchDevices :many
SELECT * from devices
WHERE id ILIKE $1
AND name ILIKE $2
AND model ILIKE $3
AND ip_addr ILIKE $4
AND mac_addr ILIKE $5
ORDER BY name
LIMIT 100;

-- name: GetDevices :many
SELECT * from devices LIMIT 100;

-- name: AddDevice :one 
INSERT INTO devices (
  id, name, model, ip_addr, mac_addr
) VALUES (
  $1,$2,$3,$4,$5
) returning id;
