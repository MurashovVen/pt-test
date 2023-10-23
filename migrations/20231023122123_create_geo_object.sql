-- +goose Up
CREATE TABLE geo_objects
(
    id         uuid  NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    geo        point NOT NULL,
    created_at timestamp                  DEFAULT now(),
    updated_at timestamp                  DEFAULT now()
);

-- +goose Down
DROP TABLE geo_objects;