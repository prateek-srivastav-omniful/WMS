-- ALTER TABLE inventory 
-- ADD CONSTRAINT fk_inventory_sku FOREIGN KEY (sku_id) REFERENCES skus(id) ON DELETE CASCADE;

-- ALTER TABLE hubs 
-- ADD CONSTRAINT fk_hubs_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE;