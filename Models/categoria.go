package models

import "errors"

type Categoria struct {
	idCategoria int
	nombre      string
	descripcion string
}

func NewCategoria(nombre, descripcion string) (*Categoria, error) {
	if nombre == "" {
		return nil, errors.New("Categoría no puede quedar sin nombre")
	}

	categoria := &Categoria{
		nombre:      nombre,
		descripcion: descripcion,
	}
	return categoria, nil

}

func (c *Categoria) GetIDCategoria() int {
	return c.idCategoria
}

func (c *Categoria) GetNombre() string {
	return c.nombre
}

func (c *Categoria) GetDescripcion() string {
	return c.descripcion
}

func (c *Categoria) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("El nombre no puede quedar vacío")
	}
	c.nombre = nombre
	return nil
}

func (c *Categoria) SetDescripcion(descripcion string) {
	c.descripcion = descripcion
}
