package errores

import "fmt"

type ErrorCargarVuelos struct{}

func (e ErrorCargarVuelos) Error() string {
	return "Error en comando agregar_archivo"
}

type ErrorTablero struct{}

func (e ErrorTablero) Error() string {
	return "Error en comando ver_tablero"
}

type ErrorVuelo struct{}

func (e ErrorVuelo) Error() string {
	return "Error en comando info_vuelo"
}

type ErrorSigVuelo struct{}

func (e ErrorSigVuelo) Error() string {
	return "Error en comando siguiente_vuelo"
}

type ErrorPrioridadVuelo struct{}

func (e ErrorPrioridadVuelo) Error() string {
	return "Error en comando prioridad_vuelos"
}

type ErrorBorrarVuelo struct{}

func (e ErrorBorrarVuelo) Error() string {
	return "Error en comando borrar"
}

type ErrorNoHaySiguiente struct {
	Origen  string
	Destino string
	Fecha   string
}

func (e ErrorNoHaySiguiente) Error() string {
	return fmt.Sprintf("No hay vuelo registrado desde %s hacia %s desde %s", e.Origen, e.Destino, e.Fecha)
}