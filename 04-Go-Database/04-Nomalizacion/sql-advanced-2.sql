USE movies_db;

-- 1. Agregar una película a la tabla movies.
INSERT INTO movies (title, rating, awards, release_date, length, genre_id)
VALUES ("La guerra de los mundos", 8.5, 2, "2005-12-12", 120, 1);

-- 2. Agregar un género a la tabla genres.
INSERT INTO genres (`name`, ranking, `active`) VALUES ("Bélicas", 13, 1);

-- 3. Asociar a la película del punto 1. genre el género creado en el punto 2.
UPDATE movies SET genre_id = 13 WHERE id = 23;

-- 4. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
UPDATE actors SET favorite_movie_id = 23 WHERE id = 1;

-- 5. Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_copy LIKE movies;
INSERT INTO movies_copy (SELECT * FROM movies);

-- 6. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM movies_copy WHERE awards < 5;

-- 7. Obtener la lista de todos los géneros que tengan al menos una película.
SELECT * FROM genres g
WHERE g.id IN (SELECT ge.id FROM genres ge INNER JOIN movies m ON ge.id = m.genre_id);

-- 8. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT * FROM actors a 
WHERE a.favorite_movie_id IN (SELECT m.id FROM movies m WHERE m.awards > 3);

-- 9. Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX idx_movies_name ON movies (title);

-- 10. Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;