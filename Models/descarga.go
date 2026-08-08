package models

import "time"

type Descarga struct {
	idDescarga    int
	fechaDescarga time.Time
}

func NewDescarga() *Descarga {
	descarga := &Descarga{
		fechaDescarga: time.Now(),
	}
	return descarga
}
func (d *Descarga) GetIdDescarga() int {
	return d.idDescarga
}

func (d *Descarga) GetFechaDescarga() time.Time {
	return d.fechaDescarga
}
