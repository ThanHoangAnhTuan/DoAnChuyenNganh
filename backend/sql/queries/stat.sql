-- name: MonthlyEarnings :many
SELECT
	MONTH(FROM_UNIXTIME(ego.created_at / 1000)) AS month,
	COUNT(*) AS total_orders,
	CAST(SUM(ego.final_total) AS DECIMAL(15,2)) AS total_revenue
FROM
	`ecommerce_go_order` ego
JOIN `ecommerce_go_accommodation` ega ON
	ego.accommodation_id = ega.id
WHERE
	ega.manager_id = sqlc.arg('manager_id')
    
    AND ego.created_at >= sqlc.arg('start_time')
    AND ego.created_at <= sqlc.arg('end_time')
	AND ego.order_status = 'payment_success'
GROUP BY
	month
ORDER BY
	month;

-- name: DailyEarnings :many
SELECT
	DATE(FROM_UNIXTIME(ego.created_at / 1000)) AS day,
	COUNT(*) AS total_orders,
	CAST(SUM(ego.final_total) AS DECIMAL(15,2)) AS total_revenue
FROM
	`ecommerce_go_order` ego
JOIN `ecommerce_go_accommodation` ega ON
	ego.accommodation_id = ega.id
WHERE
	ega.manager_id = sqlc.arg('manager_id')
    
    AND ego.created_at >= sqlc.arg('start_time')
    AND ego.created_at <= sqlc.arg('end_time')
	AND ego.order_status = 'payment_success'
GROUP BY
	day
ORDER BY
	day;