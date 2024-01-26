USE empresa_internet;

-- Insert 5 plans of internet into the table plans
INSERT INTO `plans` (`id`, `name`, `speed`, `price`, `discount`) VALUES
(1, 'Basic', 100, 1500.00, 0),
(2, 'Medium', 200, 2500.00, 0),
(3, 'Advanced', 300, 3500.00, 455.50),
(4, 'Premium', 400, 4500.00, 850.50),
(5, 'Ultra', 500, 5500.00, 1245.50);

-- Insert 10 clients into the table clients
INSERT INTO `clients` (`id`, `name`, `surname`, `dni`, `birth_date`, `state`, `city`, `plans_id`) VALUES
(1, 'John', 'Doe', '25678234', '1990-10-01', 'CA', 'Los Angeles', 1),
(2, 'Jane', 'Smith', '34123890', '1987-04-16', 'FL', 'Miami', 3),
(3, 'Bob', 'Jones', '12345678', '1980-01-01', 'TX', 'Houston', 5),
(4, 'Mary', 'Williams', '45234567', '2005-01-16', 'CA', 'Los Angeles', 2),
(5, 'Patricia', 'Brown', '34861342', '2000-12-05', 'FL', 'Miami', 4),
(6, 'Michael', 'Davis', '28765490', '1995-05-01', 'TX', 'Houston', 4),
(7, 'Linda', 'Miller', '2654890', '1936-04-29', 'CA', 'Los Angeles', 1),
(8, 'William', 'Wilson', '30987416', '1990-12-16', 'FL', 'Miami', 3),
(9, 'Elizabeth', 'Moore', '15639256', '1995-01-01', 'TX', 'Houston', 5),
(10, 'Barbara', 'Taylor', '41086329', '1998-01-16', 'CA', 'Los Angeles', 2);