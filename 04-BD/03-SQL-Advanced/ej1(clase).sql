USE movies_db;

-- 1. Mostrar el título y el nombre del género de todas las series.
SELECT m.title, g.name
FROM movies as m INNER JOIN genres g
WHERE m.genre_id = g.id;

-- 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT e.title, a.first_name, a.last_name
FROM episodes e INNER JOIN actor_episode ae ON e.id=ae.episode_id
INNER JOIN actors a ON ae.actor_id=a.id;

-- 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT s.title, COUNT(*) as seasons 
FROM series s INNER JOIN seasons se
ON s.id = se.serie_id
GROUP BY s.title;

-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT g.name, COUNT(*) as q_movies
FROM movies m INNER JOIN genres g
ON m.genre_id = g.id
GROUP BY g.name
HAVING q_movies >= 3;

-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias
--    y que estos no se repitan.
-- SELECT DISTINCT m.title, a.first_name, a.last_name
-- FROM movies m INNER JOIN actor_movie am ON m.id=am.movie_id 
-- INNER JOIN actors a ON am.actor_id=a.id
-- WHERE m.title="%La Guerra de las galaxias%";