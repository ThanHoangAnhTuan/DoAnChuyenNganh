-- name: MonthlyEarnings :many
SELECT
    DATE_FORMAT (FROM_UNIXTIME (o.created_at / 1000), '%Y-%m') AS revenue_month,
    CAST(SUM(p.total_price) AS DECIMAL(15, 2)) AS total_revenue
FROM
    `ecommerce_go_payment` p
    JOIN `ecommerce_go_order` o ON p.order_id = o.id
WHERE
    p.payment_status = 'success'
    AND o.user_id = ?
    AND o.created_at >= ?
    AND o.created_at < ?
GROUP BY
    revenue_month
ORDER BY
    revenue_month ASC;