-- +goose Up
-- +goose StatementBegin
CREATE TABLE products.products (
	id serial4 NOT NULL,
	gross_weight_g int4 NULL,
	height_mm int4 NULL,
	length_mm int4 NULL,
	net_weight_g int4 NULL,
	quantity int4 NOT NULL,
	width_mm int4 NULL,
	manufacturer_id int4 NOT NULL,
	brand varchar(255) NULL,
	description text NULL,
	gtin varchar(255) NULL,
	"name" varchar(255) NULL,
	ncm varchar(255) NULL,
	presentation varchar(255) NOT NULL,
	storage_condition varchar(255) NULL,
	unit_of_measurement_symbol varchar(255) NOT NULL,
	group_id int4 NULL,
	umbrella_item_id int4 NULL,
	drug_id int4 NULL,
	CONSTRAINT product_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products.products
-- +goose StatementEnd
