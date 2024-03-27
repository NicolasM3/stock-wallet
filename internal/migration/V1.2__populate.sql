ALTER TABLE stock
ALTER COLUMN created_at SET DEFAULT now(),
ALTER COLUMN last_update SET DEFAULT now();

INSERT INTO stock (code, name, current_price, deleted) VALUES ('AAPL', 'Apple Inc.', 123.45, false);
INSERT INTO stock (code, name, current_price, deleted) VALUES ('GOOGL', 'Alphabet Inc.', 234.56, false);
INSERT INTO stock (code, name, current_price, deleted) VALUES ('MSFT', 'Microsoft Corporation', 345.67, false);
INSERT INTO stock (code, name, current_price, deleted) VALUES ('AMZN', 'Amazon.com Inc.', 456.78, false);
INSERT INTO stock (code, name, current_price, deleted) VALUES ('FB', 'Facebook Inc.', 567.89, false);