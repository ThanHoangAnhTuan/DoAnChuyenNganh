// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: stat.sql

package database

import (
	"context"
	"time"
)

const dailyEarnings = `-- name: DailyEarnings :many
SELECT
	DATE(FROM_UNIXTIME(ego.created_at / 1000)) AS day,
	COUNT(*) AS total_orders,
	CAST(SUM(ego.final_total) AS DECIMAL(15,2)) AS total_revenue
FROM
	` + "`" + `ecommerce_go_order` + "`" + ` ego
JOIN ` + "`" + `ecommerce_go_accommodation` + "`" + ` ega ON
	ego.accommodation_id = ega.id
WHERE
	ega.manager_id = ?
    
    AND ego.created_at >= ?
    AND ego.created_at <= ?
	AND ego.order_status = 'payment_success'
GROUP BY
	day
ORDER BY
	day
`

type DailyEarningsParams struct {
	ManagerID string
	StartTime uint64
	EndTime   uint64
}

type DailyEarningsRow struct {
	Day          time.Time
	TotalOrders  int64
	TotalRevenue string
}

func (q *Queries) DailyEarnings(ctx context.Context, arg DailyEarningsParams) ([]DailyEarningsRow, error) {
	rows, err := q.db.QueryContext(ctx, dailyEarnings, arg.ManagerID, arg.StartTime, arg.EndTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DailyEarningsRow
	for rows.Next() {
		var i DailyEarningsRow
		if err := rows.Scan(&i.Day, &i.TotalOrders, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const monthlyEarnings = `-- name: MonthlyEarnings :many
SELECT
	MONTH(FROM_UNIXTIME(ego.created_at / 1000)) AS month,
	COUNT(*) AS total_orders,
	CAST(SUM(ego.final_total) AS DECIMAL(15,2)) AS total_revenue
FROM
	` + "`" + `ecommerce_go_order` + "`" + ` ego
JOIN ` + "`" + `ecommerce_go_accommodation` + "`" + ` ega ON
	ego.accommodation_id = ega.id
WHERE
	ega.manager_id = ?
    
    AND ego.created_at >= ?
    AND ego.created_at <= ?
	AND ego.order_status = 'payment_success'
GROUP BY
	month
ORDER BY
	month
`

type MonthlyEarningsParams struct {
	ManagerID string
	StartTime uint64
	EndTime   uint64
}

type MonthlyEarningsRow struct {
	Month        int32
	TotalOrders  int64
	TotalRevenue string
}

func (q *Queries) MonthlyEarnings(ctx context.Context, arg MonthlyEarningsParams) ([]MonthlyEarningsRow, error) {
	rows, err := q.db.QueryContext(ctx, monthlyEarnings, arg.ManagerID, arg.StartTime, arg.EndTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MonthlyEarningsRow
	for rows.Next() {
		var i MonthlyEarningsRow
		if err := rows.Scan(&i.Month, &i.TotalOrders, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
