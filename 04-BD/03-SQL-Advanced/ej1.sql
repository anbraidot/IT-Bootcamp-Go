--1. Listar los datos de los autores.
SELECT * FROM autor;

--2. Listar nombre y edad de los estudiantes
SELECT e.nombre, e.edad FROM estudiante e;

--3. ¿Qué estudiantes pertenecen a la carrera informática?
SELECT * FROM estudiante e WHERE e.carrera="Informática";

--4. ¿Qué autores son de nacionalidad francesa o italiana?
SELECT * FROM autor a WHERE a.nacionalidad="Francesa" OR a.nacionalidad="Italiana";

--5. ¿Qué libros no son del área de internet?
SELECT * FROM libro l WHERE l.area!="Internet";

--6. Listar los libros de la editorial Salamandra.
SELECT * FROM libro l WHERE l.editorial="Salamandra";

--7. Listar los datos de los estudiantes cuya edad es mayor al promedio.
SELECT * FROM estudiante e WHERE e.edad > (SELECT AVG(e.edad) FROM estudiante e);

--8. Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
SELECT e.nombre FROM estudiante e WHERE e.apellido LIKE "G%";

--9. Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
SELECT a.nombre FROM autor a
WHERE a.idAutor = (SELECT la.idAutor FROM libro l INNER JOIN libro_autor la
                    ON l.idLibro = la.idLibro);

--10. ¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT * FROM libros l
WHERE l.idLibro = (SELECT p.idLibro FROM prestamo p INNER JOIN estudiante e
                    ON p.idLector = e.idLector);

--11. Listar el nombre del estudiante de menor edad.
SELECT e.nombre FROM estudiante e
WHERE e.edad = (SELECT MIN(e.edad) FROM estudiante e);

--12. Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT e.nombre FROM estudiante e
WHERE e.idLector = (SELECT p.idLector FROM prestamo p INNER JOIN libro l
                    ON p.idLibro = l.idLibro
                    WHERE l.area LIKE "Base de Datos");

--13. Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT * FROM libro l
WHERE l.idLibro = (SELECT la.idLibro FROM libro_autor la INNER JOIN autor a
                    ON la.idAutor = a.idAutor
                    WHERE a.nombre LIKE "J.K. Rowling");

--14. Listar títulos de los libros que debían devolverse el 16/07/2021.
SELECT * FROM libro l INNER JOIN prestamo p
ON l.idLibro = p.idLibro
WHERE p.fechaDevolucion = "2021-07-16";