-- name: SearchDevices :many
SELECT * from devices
WHERE name ILIKE $1
AND model ILIKE $2
AND ip_addr ILIKE $3
AND mac_addr ILIKE $4
LIMIT 100;

-- name: GetDevices :many
SELECT name,model,ip_addr,mac_addr from devices LIMIT 100;
