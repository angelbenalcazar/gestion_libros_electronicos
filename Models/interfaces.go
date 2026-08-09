package models

type GestorLibros interface {
	AgregarLibro(libro *Libro) error
}

type Activable interface {
	Activar()
	Desactivar()
	GetEstado()
}
