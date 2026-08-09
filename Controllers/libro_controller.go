package controllers

//En el libro_controller se van a coordinar las operaciones que se pueden realizar con la clase Libro
//Por ejemplo se va a poder agregar libros, listar libros, editar libros, desactivar libro, buscar libro

import (
	"database/sql"
	"errors"

	models "github.com/angelbenalcazar/gestion_libros_electronicos/Models"
)

//Se va a utilizar una estructura para guardar temporalmente los libros que se van a crear

type LibroController struct {
	db *sql.DB
}

//Se puede utilizar un constructor con un controlador con lista vacía

func NewLibroController(db *sql.DB) *LibroController {
	return &LibroController{
		db: db,
	}
}

//Ahora para agregar un libro se usa el siguiente método
//libro *models.libro significa que la función recibe el objeto Libro del paquete models.

func (lc *LibroController) AgregarLibro(libro *models.Libro) error {
	if libro == nil {
		return errors.New("No se puede agregar un libro vacío")
	}

	consulta := `
		INSERT INTO Libro
		(titulo, descripcion, isbn, anioPublicacion, idioma, estado)
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6)
	`

	_, err := lc.db.Exec(
		consulta,
		libro.GetTitulo(),
		libro.GetDescripcion(),
		libro.GetISBN(),
		libro.GetAnioPublicacion(),
		libro.GetIdioma(),
		libro.GetEstado(),
	)

	if err != nil {
		return err
	}

	return nil
}

//Para listar la colección de libros

func (lc *LibroController) ListarLibros() ([]*models.Libro, error) {

	consulta := `
		SELECT
			idLibro,
			titulo,
			descripcion,
			COALESCE(isbn, ''),
			anioPublicacion,
			idioma,
			estado
		FROM Libro
	`

	filas, err := lc.db.Query(consulta)

	if err != nil {
		return nil, err
	}

	defer filas.Close()

	var libros []*models.Libro

	for filas.Next() {

		var idLibro int
		var titulo string
		var descripcion string
		var isbn string
		var anioPublicacion int
		var idioma string
		var estado bool

		err := filas.Scan(
			&idLibro,
			&titulo,
			&descripcion,
			&isbn,
			&anioPublicacion,
			&idioma,
			&estado,
		)

		if err != nil {
			return nil, err
		}

		libro := models.NewLibroDesdeBD(
			idLibro,
			titulo,
			descripcion,
			isbn,
			anioPublicacion,
			idioma,
			estado,
		)

		libros = append(libros, libro)
	}

	if err := filas.Err(); err != nil {
		return nil, err
	}

	return libros, nil
}
