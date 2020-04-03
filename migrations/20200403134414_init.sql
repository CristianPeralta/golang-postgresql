-- +goose Up
-- SQL in this section is executed when the migration is applied.
    ALTER TABLE product_items_collections ADD COLUMN prod_sec_name text NOT NULL DEFAULT '';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
    ALTER TABLE product_items_collections DROP COLUMN prod_sec_name;