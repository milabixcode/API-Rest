-- Inserir dados na tabela Group
UPDATE "BillPaymentPOSCustomer".group g
SET external_group_id = eh.group_id
FROM "BillPaymentPOSCustomer".external_hierarchy as eh
JOIN
"BillPaymentPOSCustomer".customer as c ON c.id = eh.fk_customer_id
WHERE g.id = c.fk_group_id;

-- Inserir dados na tabela Company (exceções: Equatorial e Neoenergia se repetem para o mesmo group)
INSERT INTO "BillPaymentPOSCustomer".company (external_company_id, fk_group_id, name)
SELECT DISTINCT
eh.company_id,
g.id,
c.name
FROM
"BillPaymentPOSCustomer".external_hierarchy as eh
JOIN
"BillPaymentPOSCustomer"."group" as g ON g.external_group_id = eh.group_id
JOIN
"BillPaymentPOSCustomer".customer as c On c.id = eh.fk_customer_id
group by eh.company_id, g.id, c.name 
order by company_id;

-- Inserir dados na tabela Workspace (se repete para o mesmo company)
INSERT INTO "BillPaymentPOSCustomer".workspace (external_workspace_id, fk_company_id, name, document_number, document_type)
SELECT DISTINCT
eh.workspace_id,
co.id,
c.name,
c.document_number,
c.document_type
FROM
"BillPaymentPOSCustomer".external_hierarchy  as eh
JOIN
"BillPaymentPOSCustomer".company as co ON co.external_company_id = eh.company_id
JOIN
"BillPaymentPOSCustomer".customer as c ON c.id = eh.fk_customer_id and c."name" = co."name" 
group by
eh.workspace_id,
co.id,
c.name,
c.document_number,
c.document_type;


INSERT INTO "BillPaymentPOSCustomer".application (external_application_id, fk_workspace_id, fk_external_acquirer_info_id, name, slug, channel)
SELECT DISTINCT
eh.application_id,
w.id ,
acq.id,
c.name,
c.slug,
c.channel
FROM
"BillPaymentPOSCustomer".external_hierarchy as eh
JOIN 
"BillPaymentPOSCustomer".external_acquirer_info as acq ON acq.fk_customer_id = eh.fk_customer_id
join 
"BillPaymentPOSCustomer".workspace as w on w.external_workspace_id = eh.workspace_id
JOIN
"BillPaymentPOSCustomer".customer as c ON c.id = eh.fk_customer_id and c."name" = w."name" 
group by 
eh.application_id,
w.id,
acq.id,
c.name,
c.slug,
c.channel;