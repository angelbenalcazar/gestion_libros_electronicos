package database

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func Conectar() (*sql.DB, error) {
	cadenaConexion := "server=MIGUEL\\SQLEXPRESS;database=GestionLibrosElectronicos"

	db, err := sql.Open("sqlserver", cadenaConexion)

	if err != nil {
		return nil, fmt.Errorf("Error al abrir la conexión: %w", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("Error al conectar con SQL Server: %w", err)
	}
	return db, nil
}
