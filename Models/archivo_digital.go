package models

import "errors"

type ArchivoDigital struct {
	idArchivo int
	nombre    string
	ruta      string
	formato   string
	tamanio   float64
}

func NewArchivoDigital(nombre, ruta, formato string, tamanio float64) (*ArchivoDigital, error) {
	if nombre == "" {
		return nil, errors.New("Nombre no puede estar vacío")
	}

	if ruta == "" {
		return nil, errors.New("La ruta no puede estar vacío")
	}
	if formato == "" {
		return nil, errors.New("El formato no puede estar vacío")
	}

	if tamanio < 0 {
		return nil, errors.New("El tamaño de archivo no puede ser menor que 0")
	}

	archivodigital := &ArchivoDigital{
		nombre:  nombre,
		ruta:    ruta,
		formato: formato,
		tamanio: tamanio,
	}

	return archivodigital, nil
}

func (a *ArchivoDigital) GetIDArchivo() int {
	return a.idArchivo
}

func (a *ArchivoDigital) GetNombre() string {
	return a.nombre
}

func (a *ArchivoDigital) GetRuta() string {
	return a.ruta
}

func (a *ArchivoDigital) GetFormato() string {
	return a.formato
}

func (a *ArchivoDigital) GetTamanio() float64 {
	return a.tamanio
}

func (a *ArchivoDigital) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("El nombre no puede estar vacío")
	}

	a.nombre = nombre
	return nil
}

func (a *ArchivoDigital) SetRuta(ruta string) error {
	if ruta == "" {
		return errors.New("La ruta no puede estar vacía")
	}

	a.ruta = ruta
	return nil
}

func (a *ArchivoDigital) SetFormato(formato string) error {
	if formato == "" {
		return errors.New("El formato no puede estar vacío")
	}

	a.formato = formato
	return nil
}

func (a *ArchivoDigital) SetTamanio(tamanio float64) error {
	if tamanio < 0 {
		return errors.New("El tamaño no puede ser menor que 0")
	}

	a.tamanio = tamanio
	return nil
}
