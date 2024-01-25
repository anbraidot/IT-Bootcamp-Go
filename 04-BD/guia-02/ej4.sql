-- 1. Get all clients from Los Angeles
SELECT * FROM clients WHERE city = 'Los Angeles';

-- 2. Get client name and surname from clients from California that have a plan with speed greater than 100
SELECT name, surname FROM clients WHERE state = 'CA' AND plans_id IN (SELECT id FROM plans WHERE speed > 100);

-- 3. Get clients that born after 2000
SELECT * FROM clients WHERE birth_date > '2000-01-01';

-- 4. Get plans with speed greater than 100 and price less than 4500
SELECT * FROM plans WHERE speed > 100 AND price < 4500;

-- 5. Get clients name that your name starts with 'J'
SELECT * FROM clients WHERE name LIKE 'J%';

-- 6. Get the quantity of clients that have a plan with speed greater than 200
SELECT COUNT(*) as total FROM clients WHERE plans_id IN (SELECT id FROM plans WHERE speed > 200);

-- 7. Get the quantity of clients that live in California
SELECT COUNT(*) FROM clients WHERE state = 'CA';

-- 8. Get the price average of plans
SELECT AVG(price) FROM plans;

-- 9. Get the discount average of plans
SELECT AVG(discount) FROM plans;

-- 10. Get the client name that have the highest speed plan
SELECT name FROM clients WHERE plans_id IN (SELECT id FROM plans WHERE speed = (SELECT MAX(speed) FROM plans));