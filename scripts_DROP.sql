
--Exclusão da coluna de fk_customer_id na tabela external_acquirer_info
ALTER TABLE "BillPaymentPOSCustomer".external_acquirer_info
DROP COLUMN fk_customer_id;

--Exclusão da antiga tabela de external_hierarchy
DROP INDEX IF EXISTS idx_external_hierarchy_deleted_at cascade;
DROP TABLE IF EXISTS "BillPaymentPOSCustomer".external_hierarchy;

--Exclusão da antiga tabela de customer
DROP INDEX IF EXISTS "BillPaymentPOSCustomer".idx_customer_slug;
DROP TABLE IF EXISTS "BillPaymentPOSCustomer".customer;