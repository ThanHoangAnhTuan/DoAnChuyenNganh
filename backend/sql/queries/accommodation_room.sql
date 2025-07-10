-- name: CheckAccommodationTypeBelongToManager :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_accommodation` ega
            JOIN `ecommerce_go_accommodation_detail` egad ON ega.id = egad.accommodation_id
            JOIN `ecommerce_go_user_manager` egum ON egum.id = ega.manager_id
        WHERE
            egum.id = sqlc.arg ("managerID")
            AND egad.id = sqlc.arg ("AccommodationTypeID")
    );

-- name: CheckAccommodationRoomBelongsToManager :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            `ecommerce_go_accommodation` ega
            JOIN `ecommerce_go_accommodation_detail` egad ON ega.id = egad.accommodation_id
            JOIN `ecommerce_go_user_manager` egum ON egum.id = ega.manager_id
            JOIN `ecommerce_go_accommodation_room` egar ON egar.accommodation_type = egad.id
        WHERE
            egum.id = sqlc.arg ("managerID")
            AND egar.id = sqlc.arg ("AccommodationRoomID")
            AND egar.is_deleted = 0
    );

-- name: CreateAccommodationRoom :exec
INSERT INTO
    `ecommerce_go_accommodation_room` (
        `id`,
        `accommodation_type`,
        `name`,
        `status`,
        `is_deleted`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, 0, ?, ?);

-- name: GetAccommodationRooms :many
SELECT
    `id`,
    `accommodation_type`,
    `name`,
    `status`
FROM
    `ecommerce_go_accommodation_room`
WHERE
    `accommodation_type` = sqlc.arg ("accommodationTypeID")
    AND `is_deleted` = 0
ORDER BY `created_at` ASC;

-- name: UpdateAccommodationRooms :exec
UPDATE `ecommerce_go_accommodation_room`
SET
    `name` = sqlc.arg ("name"),
    `status` = sqlc.arg ("status"),
    `updated_at` = sqlc.arg ("updatedAt")
WHERE
    `id` = sqlc.arg ("id")
    AND `is_deleted` = 0;

-- name: DeleteAccommodationRoom :exec
UPDATE `ecommerce_go_accommodation_room`
SET
    `is_deleted` = 1
WHERE
    `id` = sqlc.arg ("id");

-- name: GetAccommodationRoomAvailable :one
SELECT
    *
FROM
    `ecommerce_go_accommodation_room` ar
WHERE
    ar.id NOT IN (
        SELECT
            egar.id
        FROM
            `ecommerce_go_order` ego
            JOIN `ecommerce_go_order_detail` egod ON ego.id = egod.order_id
            JOIN `ecommerce_go_accommodation_room` egar ON egar.id = egod.accommodation_room_id
        WHERE
            sqlc.arg ("check_out") > ego.checkin_date
            AND sqlc.arg ("check_in") < ego.checkout_date
            AND ego.order_status in ('payment_success', 'checked_in')
    )
    AND ar.accommodation_type = sqlc.arg ("accommodation_type_id")
    and ar.status in ('available')
LIMIT
    1;

-- name: CountAccommodationRoomAvailable :one
SELECT
    COUNT(ar.id)
FROM
    `ecommerce_go_accommodation_room` ar
WHERE
    ar.id NOT IN (
        SELECT
            egar.id
        FROM
            `ecommerce_go_order` ego
            JOIN `ecommerce_go_order_detail` egod ON ego.id = egod.order_id
            JOIN `ecommerce_go_accommodation_room` egar ON egar.id = egod.accommodation_room_id
        WHERE
            sqlc.arg ("check_out") > ego.checkin_date
            AND sqlc.arg ("check_in") < ego.checkout_date
            AND ego.order_status in ('payment_success', 'checked_in')
    )
    AND ar.accommodation_type = sqlc.arg ("accommodation_type_id")
    and ar.status in ('available') AND `is_deleted` = 0
LIMIT
    1;

-- name: BatchCountAccommodationRoomAvailable :many
SELECT
    ar.accommodation_type AS accommodation_type_id,
    COUNT(ar.id) AS available_count
FROM
    ecommerce_go_accommodation_room ar
WHERE
    ar.id NOT IN (
        SELECT
            egar.id
        FROM
            ecommerce_go_order ego
            JOIN ecommerce_go_order_detail egod ON ego.id = egod.order_id
            JOIN ecommerce_go_accommodation_room egar ON egar.id = egod.accommodation_room_id
        WHERE
            sqlc.arg('check_out') > ego.checkin_date
            AND sqlc.arg('check_in') < ego.checkout_date
            AND ego.order_status IN ('payment_success', 'checked_in')
    )
    AND ar.accommodation_type IN (sqlc.slice('ids'))
    AND ar.status = 'available'
    AND ar.is_deleted = 0
GROUP BY ar.accommodation_type;


-- name: CountAccommodationRoomAvailableByManager :one
SELECT
    COUNT(ar.id)
FROM
    `ecommerce_go_accommodation_room` ar
WHERE
    ar.accommodation_type = sqlc.arg ("accommodation_type_id")
    AND ar.status in ('available') AND `is_deleted` = 0;

-- name: BatchCountAccommodationRoomAvailableByManager :many
SELECT
    ar.accommodation_type as accommodation_type_id,
    COUNT(ar.id) as available_count
FROM
    `ecommerce_go_accommodation_room` ar
WHERE
    ar.accommodation_type IN (sqlc.slice("accommodation_type_ids"))
    AND ar.status IN ('available') 
    AND ar.is_deleted = 0
GROUP BY ar.accommodation_type;