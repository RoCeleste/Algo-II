package vuelos

import (
	err "algueiza/diseno_vuelos/errores"
	"fmt"
	"strconv"
	"strings"
	TDAHeap "algueiza/cola_prioridad"
	TDADiccionario "algueiza/diccionario"
)

const (
	NADA          = ""
	N_VUELO   int = 0
	ORIGEN    int = 2
	DESTINO   int = 3
	PRIORIDAD int = 5
	FECHA     int = 6
)

//
//--------------------------TIPOS----------------------------------

type vuelosEnSistema struct {
	vuelos           TDADiccionario.Diccionario[string, vuelo]
	vuelos_ordenados TDADiccionario.DiccionarioOrdenado[string, vuelo]
}

type vuelo struct {
	numero_vuelo string
	origen       string
	destino      string
	prioridad    int
	fecha        string
	info_general string
}

//
//-------------------------CREADORES-----------------------------

func CrearRegistroVuelos() Sistema_Vuelos {
	registro := new(vuelosEnSistema)
	registro.vuelos = TDADiccionario.CrearHash[string, vuelo]()
	registro.vuelos_ordenados = TDADiccionario.CrearABB[string, vuelo](strings.Compare) // clave del tipo <fecha> - <n° vuelo>
	return registro
}

func CrearVuelo(info []string) *vuelo {
	vuelo := new(vuelo)
	vuelo.numero_vuelo = info[N_VUELO]
	vuelo.origen = info[ORIGEN]
	vuelo.destino = info[DESTINO]

	num_prioridad_int, _ := strconv.Atoi(info[PRIORIDAD])
	vuelo.prioridad = num_prioridad_int
	vuelo.fecha = info[FECHA]

	convertirStrValido(info)
	vuelo_info := strings.Join(info, " ")
	vuelo.info_general = vuelo_info
	return vuelo
}

//
//-------------------------------PRIMITIVAS---------------------------------------

func (registro *vuelosEnSistema) AgregarVuelo(vuelo vuelo) {

	if registro.vuelos.Pertenece(vuelo.numero_vuelo) {
		vuelo_a_borrar := registro.vuelos.Obtener(vuelo.numero_vuelo)
		registro.vuelos_ordenados.Borrar(fmt.Sprintf("%s - %s", vuelo_a_borrar.fecha, vuelo_a_borrar.numero_vuelo))
	}

	registro.vuelos.Guardar(vuelo.numero_vuelo, vuelo)
	registro.vuelos_ordenados.Guardar(fmt.Sprintf("%s - %s", vuelo.fecha, vuelo.numero_vuelo), vuelo)

}

func (registro *vuelosEnSistema) VuelosEnRango(desde string, hasta string) []string {
	var arreglo_vuelos_seleccionados []string

	vuelos := iteroabb(registro.vuelos_ordenados, desde, hasta)
	for _, vuelo := range vuelos {
		arreglo_vuelos_seleccionados = append(arreglo_vuelos_seleccionados, fmt.Sprintf("%s - %s", vuelo.fecha, vuelo.numero_vuelo))
	}
	return arreglo_vuelos_seleccionados

}

func (registro *vuelosEnSistema) ObtenerInformacionVuelo(numero_vuelo string) (string, error) {
	if !registro.vuelos.Pertenece(numero_vuelo) {
		error_a_devolver := new(err.ErrorVuelo)
		return NADA, error_a_devolver
	}
	return registro.vuelos.Obtener(numero_vuelo).info_general, nil
}

func (registro *vuelosEnSistema) ObtenerKMayorPrioridad(cant int) []string {
	var prioridad_vuelo []vuelo

	iter := registro.vuelos.Iterador()
	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		prioridad_vuelo = append(prioridad_vuelo, vuelo)
		iter.Siguiente()
	}

	heapify := TDAHeap.CrearHeapArr(prioridad_vuelo, comparacionPrioridadVuelo)
	var vuelos []string
	contador := 0

	//desencolo hasta que contador == cant o hasta que contador == largo del heap
	for contador < cant && contador < len(prioridad_vuelo) {
		vuelo := heapify.Desencolar()
		vuelos = append(vuelos, fmt.Sprintf("%d - %s", vuelo.prioridad, vuelo.numero_vuelo))
		contador += 1
	}
	return vuelos
}

func (registro *vuelosEnSistema) ObtenerSiguienteVuelo(origen string, destino string, fecha string) (string, error) {
	var vuelo_a_devolver vuelo
	registro.vuelos_ordenados.IterarRango(&fecha, nil, (func(clave string, dato vuelo) bool {

		if dato.origen == origen && dato.destino == destino {
			vuelo_a_devolver = dato
			return false
		}
		return true
	}))

	if vuelo_a_devolver.info_general == NADA {
		error_a_devolver := new(err.ErrorNoHaySiguiente)
		error_a_devolver.Origen = origen
		error_a_devolver.Destino = destino
		error_a_devolver.Fecha = fecha
		return vuelo_a_devolver.info_general, error_a_devolver
	}

	return vuelo_a_devolver.info_general, nil
}

func (registro *vuelosEnSistema) BorrarVuelos(desde string, hasta string) []string {
	var info_vuelos_borrados []string
	vuelos_a_borrar := iteroabb(registro.vuelos_ordenados, desde, hasta)

	for _, vuelo_eliminar := range vuelos_a_borrar {
		info_vuelos_borrados = append(info_vuelos_borrados, vuelo_eliminar.info_general)
		registro.vuelos.Borrar(vuelo_eliminar.numero_vuelo)
		registro.vuelos_ordenados.Borrar(fmt.Sprintf("%s - %s", vuelo_eliminar.fecha, vuelo_eliminar.numero_vuelo))

	}
	return info_vuelos_borrados
}

// -----------------------FUNCIONES AUXILIARES----------------------------------
func convertirStrValido(info []string) {
	//unico uso de esta funcion es sacarle los 0 (si es necesario) al numero de prioridad y al departure_delay
	// ej 06 ------> 6
	num_prioridad_int, _ := strconv.Atoi(info[5])
	info[5] = strconv.Itoa(num_prioridad_int)

	num_departure, _ := strconv.Atoi(info[7])
	info[7] = strconv.Itoa(num_departure)

}

func comparacionPrioridadVuelo(a vuelo, b vuelo) int {
	if a.prioridad == b.prioridad {
		return -strings.Compare(a.numero_vuelo, b.numero_vuelo)
	}
	return a.prioridad - b.prioridad
}

func iteroabb(abb TDADiccionario.DiccionarioOrdenado[string, vuelo], desde string, hasta string) []vuelo {
	var vuelos_a_borrar []vuelo
	iter := abb.IteradorRango(&desde, &hasta)

	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		vuelos_a_borrar = append(vuelos_a_borrar, vuelo)
		iter.Siguiente()
	}
	return vuelos_a_borrar
}