CREATE schema TimeMachine

CREATE TABLE TimeMachine.product (
    id SERIAL,
    productName VARCHAR(100) NOT NULL,
    productType VARCHAR(20) NOT NULL,
    supplierCost DECIMAL NOT NULL,
    salePrice DECIMAL NOT NULL,
    CONSTRAINT pk_product PRIMARY KEY (id),
    CONSTRAINT chk_product_supplierCost CHECK (supplierCost > 0),
    CONSTRAINT chk_product_salePrice CHECK (salePrice > 0),
    CONSTRAINT chk_product_productType CHECK(productType IN ('beer', 'laptop', 'watch'))
)

INSERT INTO TimeMachine.product(productName, productType, supplierCost, salePrice)
    VALUES('ThinkPad T Series', 'laptop', 800.00, 1000.99);
INSERT INTO TimeMachine.product(productName, productType, supplierCost, salePrice)
    VALUES('Seiko 5', 'watch', 45.00, 99.95);
INSERT INTO TimeMachine.product(productName, productType, supplierCost, salePrice)
    VALUES('Das Beer', 'beer', 3, 6.00);