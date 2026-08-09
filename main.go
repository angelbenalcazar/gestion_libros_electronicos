package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	controllers "github.com/angelbenalcazar/gestion_libros_electronicos/Controllers"
	database "github.com/angelbenalcazar/gestion_libros_electronicos/Database"
	models "github.com/angelbenalcazar/gestion_libros_electronicos/Models"
)

func main() {

	//Para conectar con la base de datos SQL Server

	db, err := database.Conectar()

	if err != nil {
		fmt.Println("Error de conexión", err)
		return
	}

	defer db.Close()

	fmt.Println("Conexión exitosa con SQL Server")

	//Para llamar al constructor NewLibroController que se conecta a la base de datos y que se encuentra en el paquete controller

	libroController := controllers.NewLibroController(db)

	reader := bufio.NewReader(os.Stdin)

	//Creación del avance del menú interactivo

	for {
		fmt.Println("===SISTEMA DE GESTIÓN DE LIBROS===")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Salir")
		fmt.Println("Seleccione una opción")

		opcionTexto, _ := reader.ReadString('\n')
		opcionTexto = strings.TrimSpace(opcionTexto)

		opcion, err := strconv.Atoi(opcionTexto)

		if err != nil {
			fmt.Println("Debe ingresar una opción válida")
		}

		switch opcion {
		case 1:
			agregarLibro(reader, libroController)

		case 2:
			listarLibros(libroController)

		case 3:
			fmt.Println("Saliendo del sistema")
			return

		default:
			fmt.Println("Opción no válida")

		}

	}

}

//Agregar libros

func agregarLibro(
	reader *bufio.Reader,
	controller *controllers.LibroController,
) {
	fmt.Println("AGREGAR LIBRO")

	fmt.Println("Título: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	fmt.Println("Descripción: ")
	descripcion, _ := reader.ReadString('\n')
	descripcion = strings.TrimSpace(descripcion)

	fmt.Print("ISBN (opcional): ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	fmt.Print("Año de publicación: ")
	anioTexto, _ := reader.ReadString('\n')
	anioTexto = strings.TrimSpace(anioTexto)

	anio, err := strconv.Atoi(anioTexto)

	if err != nil {
		fmt.Println("Error: el año debe ser un número.")
		return
	}

	fmt.Println("Idioma: ")
	idioma, _ := reader.ReadString('\n')
	idioma = strings.TrimSpace(idioma)

	libro, err := models.NewLibro(
		titulo,
		descripcion,
		isbn,
		anio,
		idioma,
	)

	if err != nil {
		fmt.Println("Error al guardar un libro", err)
		return
	}

	err = controller.AgregarLibro(libro)

	if err != nil {
		fmt.Println("Error al guardar el libro", err)
		return
	}

	fmt.Println("Libro agregado correctamente")
}

//Listar libros

func listarLibros(controller *controllers.LibroController) {
	libros, err := controller.ListarLibros()

	if err != nil {
		fmt.Println("Error al listar libros", err)
	}

	if len(libros) == 0 {
		fmt.Println("\nNo existen libros registrados")
		return
	}
	fmt.Println("-LIBROS REGISTRADOS-")
	for _, libros := range libros {
		fmt.Printf("ID: %d. Título: %s |Año: %d | Idioma: %s\n", libros.GetIDLibro(), libros.GetTitulo(), libros.GetAnioPublicacion(), libros.GetIdioma())
	}
}
