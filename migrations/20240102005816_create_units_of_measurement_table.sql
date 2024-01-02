-- +goose Up
-- +goose StatementBegin
CREATE TABLE products.units_of_measurement (
	"name" varchar(255) NULL,
	symbol varchar(255) NOT NULL,
	deleted_at timestamp NULL,
	CONSTRAINT units_of_measurement_pkey PRIMARY KEY (symbol)
);

ALTER TABLE products.products
ADD CONSTRAINT products_units_measurement_symbol_fkey
FOREIGN KEY (unit_of_measurement_symbol)
REFERENCES products.units_of_measurement(symbol);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products.products
DROP CONSTRAINT products_units_measurement_symbol_fkey;

DROP TABLE products.units_of_measurement;
-- +goose StatementEnd
