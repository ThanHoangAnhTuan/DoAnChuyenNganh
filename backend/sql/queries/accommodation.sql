-- name: CreateAccommodation :exec
INSERT INTO
    `ecommerce_go_accommodation` (
        `id`,
        `manager_id`,
        `name`,
        `city`,
        `provine`,
        `district`,
        `images`,
        `description`,
        `facilities`,
        `gg_map`,
        `property_surroundings`,
        `rules`,
        `created_at`,
        `updated_at`
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);