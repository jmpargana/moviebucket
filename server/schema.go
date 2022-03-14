package main

const schema = `
CREATE TABLE IF NOT EXISTS movies (
	id INT PRIMARY KEY SERIAL,
	name VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS counts (
	id INT PRIMARY KEY, 
	count INT
);

INSERT INTO counts (id, count) VALUES (1, 0);
`
