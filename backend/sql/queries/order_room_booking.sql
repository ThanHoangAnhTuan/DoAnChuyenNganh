-- name: CreateOrderRoomBooking :exec
INSERT INTO `ecommerce_go_order_room_booking` (
    `id`,
    `order_detail_id`,
    `accommodation_room_id`,
    `booking_status`,
    `created_at`,
    `updated_at`
) VALUES (?, ?, ?, ?, ?, ?);

-- name: GetOrderRoomBookingsByOrderDetailID :many
SELECT * FROM `ecommerce_go_order_room_booking` 
WHERE `order_detail_id` = ? AND `is_deleted` = 0;

-- name: GetOrderRoomBookingsByOrderDetailIDWithRoomInfo :many
SELECT 
    orb.id,
    orb.order_detail_id,
    orb.accommodation_room_id,
    orb.booking_status,
    orb.created_at,
    orb.updated_at,
    ar.name as room_name
FROM `ecommerce_go_order_room_booking` orb
JOIN `ecommerce_go_accommodation_room` ar ON orb.accommodation_room_id = ar.id
WHERE orb.order_detail_id = ?
ORDER BY ar.name ASC;

-- name: GetOrderRoomBookingsByOrderID :many
SELECT orb.* FROM `ecommerce_go_order_room_booking` orb
JOIN `ecommerce_go_order_detail` od ON orb.order_detail_id = od.id
WHERE od.order_id = ? AND orb.is_deleted = 0;

-- name: UpdateOrderRoomBookingStatus :exec
UPDATE `ecommerce_go_order_room_booking` 
SET `booking_status` = ?, `updated_at` = ?
WHERE `id` = ?;

-- name: GetOrderRoomBookingByID :one
SELECT * FROM `ecommerce_go_order_room_booking` 
WHERE `id` = ? AND `is_deleted` = 0;

-- name: CancelOrderRoomBooking :exec
UPDATE `ecommerce_go_order_room_booking` 
SET `booking_status` = 'canceled', `updated_at` = ?
WHERE `order_detail_id` = ?;

-- name: GetRoomBookingsByAccommodationRoomID :many
SELECT * FROM `ecommerce_go_order_room_booking` 
WHERE `accommodation_room_id` = ? AND `booking_status` IN ('reserved', 'checked_in')
ORDER BY `created_at` DESC;
