-- +migrate Up
CREATE TABLE products (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          name TEXT NOT NULL CHECK (length(trim(name)) > 0 AND length(trim(name)) <= 255),
                          vendor TEXT NOT NULL CHECK (length(trim(vendor)) > 0 AND length(trim(vendor)) <= 255),
                          description TEXT CHECK (octet_length(description) <= 10000),
                          price NUMERIC(10,2) NOT NULL CHECK (price >= 0),
                          created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                          updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                          CONSTRAINT uq_products_name_vendor UNIQUE (name, vendor)
);

COMMENT ON COLUMN products.id IS 'Unique identifier for the product';
COMMENT ON COLUMN products.name IS 'Unique name of the product';
COMMENT ON COLUMN products.vendor IS 'Product vendor / manufacturer';
COMMENT ON COLUMN products.description IS 'Detailed description';
COMMENT ON COLUMN products.price IS 'Price in USD';
COMMENT ON COLUMN products.created_at IS 'Creation timestamp';
COMMENT ON COLUMN products.updated_at IS 'Last update timestamp';

CREATE TRIGGER products_set_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_updated_at();

-- +migrate Down
DROP TRIGGER IF EXISTS products_set_updated_at ON products;

DROP TABLE IF EXISTS products;
