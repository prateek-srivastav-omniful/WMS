CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL PRIMARY KEY,
    sku_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ALTER TABLE inventory 
-- ADD CONSTRAINT fk_inventory_sku FOREIGN KEY (sku_id) REFERENCES skus(id) ON DELETE CASCADE;
