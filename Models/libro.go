package models

import "errors"

type Libro struct {
	idLibro         int
	titulo          string
	descripcion     string
	isbn            string
	anioPublicacion int
	idioma          string
	estado          bool
}

func NewLibro(titulo, descripcion, isbn string, anioPublicacion int, idioma string) (*Libro, error) {

	if titulo == "" {
		return nil, errors.New("Título no puede estar vacío")
	}
	if descripcion == "" {
		return nil, errors.New("Descripción no puede estar vacía")
	}

	if anioPublicacion <= -9999 {
		return nil, errors.New("Año no válido")
	}
	if idioma == "" {
		return nil, errors.New("Idioma no puede estar vacío")
	}

	libro := &Libro{
		titulo:          titulo,
		descripcion:     descripcion,
		isbn:            isbn,
		anioPublicacion: anioPublicacion,
		idioma:          idioma,
		estado:          true,
	}
	return libro, nil
}

func (l *Libro) GetIDLibro() int {
	return l.idLibro
}

func (l *Libro) GetTitulo() string {
	return l.titulo
}

func (l *Libro) GetDescripcion() string {
	return l.descripcion
}

func (l *Libro) GetISBN() string {
	return l.isbn
}

func (l *Libro) GetAnioPublicacion() int {
	return l.anioPublicacion
}

func (l *Libro) GetIdioma() string {
	return l.idioma
}

func (l *Libro) GetEstado() bool {
	return l.estado
}

func (l *Libro) SetTitulo(titulo string) error {
	if titulo == "" {
		return errors.New("El título no puede estar vacío")
	}
	l.titulo = titulo
	return nil
}

func (l *Libro) SetDescripcion(descripcion string) error {
	if descripcion == "" {
		return errors.New("La descripción no puede estar vacía")
	}
	l.descripcion = descripcion
	return nil
}

func (l *Libro) SetISBN(isbn string) {
	l.isbn = isbn
}

func (l *Libro) SetAnioPublicacion(anio int) error {
	if anio <= -9999 {
		return errors.New("El año de publicación ingresado no es correcto")
	}
	l.anioPublicacion = anio
	return nil
}

func (l *Libro) SetIdioma(idioma string) error {
	if idioma == "" {
		return errors.New("El idioma no puede estar vacío")
	}
	l.idioma = idioma
	return nil
}

func (l *Libro) Activar() {
	l.estado = true
}

func (l *Libro) Desactivar() {
	l.estado = false
}
