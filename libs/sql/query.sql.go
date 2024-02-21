// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package sql

import (
	"context"
)

const getDevices = `-- name: GetDevices :many
SELECT name,model,ip_addr,mac_addr from devices LIMIT 100
`

type GetDevicesRow struct {
	Name    string
	Model   string
	IpAddr  string
	MacAddr string
}

func (q *Queries) GetDevices(ctx context.Context) ([]GetDevicesRow, error) {
	rows, err := q.db.Query(ctx, getDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDevicesRow
	for rows.Next() {
		var i GetDevicesRow
		if err := rows.Scan(
			&i.Name,
			&i.Model,
			&i.IpAddr,
			&i.MacAddr,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchDevices = `-- name: SearchDevices :many
SELECT id, name, model, ip_addr, mac_addr from devices
WHERE name ILIKE $1
OR model ILIKE $2
OR ip_addr ILIKE $3
OR mac_addr ILIKE $4
LIMIT 100
`

type SearchDevicesParams struct {
	Name    string
	Model   string
	IpAddr  string
	MacAddr string
}

func (q *Queries) SearchDevices(ctx context.Context, arg SearchDevicesParams) ([]Device, error) {
	rows, err := q.db.Query(ctx, searchDevices,
		arg.Name,
		arg.Model,
		arg.IpAddr,
		arg.MacAddr,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Device
	for rows.Next() {
		var i Device
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Model,
			&i.IpAddr,
			&i.MacAddr,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}