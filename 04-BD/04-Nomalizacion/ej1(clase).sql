USE movies_db;

-- 1. Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y 
--    guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.
CREATE TEMPORARY TABLE twd LIKE episodes;

INSERT INTO twd (
	SELECT e.* FROM episodes e 
	WHERE e.season_id IN (SELECT se.id FROM series s INNER JOIN seasons se ON s.id = se.serie_id
						WHERE s.title LIKE "%The_Walking_Dead%"));

-- 2. Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.
SELECT * FROM twd t INNER JOIN seasons s ON t.season_id = s.id
WHERE s.`number` = 1;