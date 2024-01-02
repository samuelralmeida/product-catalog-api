-- +goose Up
-- +goose StatementBegin
CREATE TABLE products.manufacturers (
	id serial NOT NULL,
	"name" varchar(255) NULL,
	deleted_at timestamp NULL,
	CONSTRAINT manufacturer_pkey PRIMARY KEY (id)
);

ALTER TABLE products.products
ADD CONSTRAINT products_manufacturers_id_fkey
FOREIGN KEY (manufacturer_id)
REFERENCES products.manufacturers(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products.products
DROP CONSTRAINT products_manufacturers_id_fkey;

DROP TABLE products.manufacturers
-- +goose StatementEnd
