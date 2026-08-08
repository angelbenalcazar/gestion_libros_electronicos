package models

import "errors"

type Autor struct {
	idAutor      int
	nombre       string
	apellido     string
	biografia    string
	nacionalidad string
}

func NewAutor(nombre, apellido, biografia, nacionalidad string) (*Autor, error) {
	if nombre == "" {
		return nil, errors.New("Nombre no puede estar vacío")
	}
	if apellido == "" {
		return nil, errors.New("Apellido no puede estar vacía")
	}

	autor := &Autor{
		nombre:       nombre,
		apellido:     apellido,
		biografia:    biografia,
		nacionalidad: nacionalidad,
	}
	return autor, nil
}

func (a *Autor) GetIDAutor() int {
	return a.idAutor
}

func (a *Autor) GetNombre() string {
	return a.nombre
}

func (a *Autor) GetApellido() string {
	return a.apellido
}

func (a *Autor) GetBiografia() string {
	return a.biografia
}

func (a *Autor) GetNacionalidad() string {
	return a.nacionalidad
}

func (a *Autor) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("Nombre no puede estar vacío")
	}
	a.nombre = nombre
	return nil
}

func (a *Autor) SetApellido(apellido string) error {
	if apellido == "" {
		return errors.New("Apellido no puede estar vacío")
	}

	a.apellido = apellido
	return nil
}

func (a *Autor) SetBiografia(biografia string) {

	a.biografia = biografia
}

func (a *Autor) SetNacionalidad(nacionalidad string) {

	a.nacionalidad = nacionalidad
}
