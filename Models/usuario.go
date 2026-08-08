package models

import "errors"

type Usuario struct {
	idUsuario    int
	nombre       string
	apellido     string
	correo       string
	passwordHash string
	estado       bool
}

//Función para crear usuarios

func NewUsuario(nombre, apellido, correo, passwordHash string) (*Usuario, error) {
	if nombre == "" {
		return nil, errors.New("El nombre no puede estar vacío")
	}
	if apellido == "" {
		return nil, errors.New("El apellido no puede estar vacío")
	}
	if correo == "" {
		return nil, errors.New("El correo no puede estar vacío")
	}
	if passwordHash == "" {
		return nil, errors.New("La contraseña no puede estar vacía")
	}

	usuario := &Usuario{
		nombre:       nombre,
		apellido:     apellido,
		correo:       correo,
		passwordHash: passwordHash,
		estado:       true,
	}
	return usuario, nil
}

//Método Getter para consultar los atributos necesarios

func (u *Usuario) GetIDUsuario() int {
	return u.idUsuario
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetApellido() string {
	return u.apellido
}

func (u *Usuario) GetCorreo() string {
	return u.correo
}

func (u *Usuario) GetEstado() bool {
	return u.estado
}

func (u *Usuario) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("El nombre no puede estar vacío")
	}

	u.nombre = nombre
	return nil
}

func (u *Usuario) SetApellido(apellido string) error {
	if apellido == "" {
		return errors.New("El apellido no puede estar vacío")
	}

	u.apellido = apellido
	return nil
}

func (u *Usuario) SetCorreo(correo string) error {
	if correo == "" {
		return errors.New("El correo no puede estar vacío")
	}

	u.correo = correo
	return nil
}

func (u *Usuario) Activar() {
	u.estado = true
}

func (u *Usuario) Desactivar() {
	u.estado = false
}

func (u *Usuario) cambiarPassword(nuevoPassword string) error {
	if nuevoPassword == "" {
		return errors.New("La contraseña no puede estar vacía")
	}

	u.passwordHash = nuevoPassword
	return nil
}
