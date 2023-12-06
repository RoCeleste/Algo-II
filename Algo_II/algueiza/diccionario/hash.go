package diccionario

import (
	"fmt"
	"hash/fnv"
)

type tipoEstado int

const (
	VACIO tipoEstado = iota
	OCUPADO
	BORRADO
)

const _TAM_INICIAL = 25
const _RAZON_EXPANSION = 1.5
const _RAZON_CONTRACCION = 0.6
const _TASA_OCUPACION_PARA_EXPANSION = 0.7
const _TASA_OCUPACION_PARA_CONTRACCION = 0.3

// tipos
//-----------------------------------------------------------------

type elemento[K comparable, V any] struct {
	clave  K
	valor  V
	estado tipoEstado
}

type hash[K comparable, V any] struct {
	elementos         []elemento[K, V]
	capacidad_total   int
	cantidad          int
	cant_mas_borrados int
}

type iterador[K comparable, V any] struct {
	diccionario *hash[K, V]
	pos_actual  int
}

// metodos creadores
//-----------------------------------------------------------------

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	Hash := new(hash[K, V])
	Hash.cantidad = 0
	Hash.capacidad_total = _TAM_INICIAL
	Hash.elementos = crearArregloElementos[K, V](_TAM_INICIAL)
	return Hash
}

func crearArregloElementos[K comparable, V any](capacidad int) []elemento[K, V] {
	return make([]elemento[K, V], capacidad)
}

// primitivas de Diccionario
//------------------------------------------------------------------

func (diccionario *hash[K, V]) Guardar(clave K, valor V) {
	capNueva := float64(diccionario.capacidad_total) * _TASA_OCUPACION_PARA_EXPANSION
	if (diccionario.cant_mas_borrados + 1) > int(capNueva) {

		diccionario.redimensionar(int(float64(diccionario.capacidad_total) * _RAZON_EXPANSION))
	}

	elem := buscar_elemento(diccionario, clave)
	if diccionario.elementos[elem].estado == OCUPADO {
		diccionario.elementos[elem].valor = valor
	} else {
		diccionario.elementos[elem].clave = clave
		diccionario.elementos[elem].valor = valor
		diccionario.elementos[elem].estado = OCUPADO
		diccionario.cantidad += 1
		diccionario.cant_mas_borrados += 1
	}
}

func (diccionario *hash[K, V]) Pertenece(clave K) bool {
	elem := buscar_elemento(diccionario, clave)
	return diccionario.elementos[elem].estado == OCUPADO

}

func (diccionario *hash[K, V]) Obtener(clave K) V {
	elem := buscar_elemento(diccionario, clave)
	if diccionario.elementos[elem].estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	}

	return diccionario.elementos[elem].valor
}

func (diccionario *hash[K, V]) Borrar(clave K) V {
	capNueva := float64(diccionario.capacidad_total) * _TASA_OCUPACION_PARA_CONTRACCION
	if (diccionario.cant_mas_borrados-1) < int(capNueva) && diccionario.capacidad_total != _TAM_INICIAL {

		diccionario.redimensionar(int(float64(diccionario.capacidad_total) * _RAZON_CONTRACCION))
	}

	elem := buscar_elemento(diccionario, clave)
	if diccionario.elementos[elem].estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	} else {
		diccionario.elementos[elem].estado = BORRADO
		diccionario.cantidad -= 1
	}
	return diccionario.elementos[elem].valor
}

func (diccionario *hash[K, V]) Cantidad() int {
	return diccionario.cantidad
}

func (diccionario *hash[K, V]) Iterar(f func(clave K, valor V) bool) {
	for i := 0; i < diccionario.capacidad_total; i += 1 {
		if diccionario.elementos[i].estado == OCUPADO {
			if !f(diccionario.elementos[i].clave, diccionario.elementos[i].valor) {
				return
			}
		}
	}
}

func (dicc *hash[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterador[K, V])
	iter.diccionario = dicc
	for iter.pos_actual < iter.diccionario.capacidad_total && iter.diccionario.elementos[iter.pos_actual].estado != OCUPADO { //voy a la primera clave del arreglo o
		iter.pos_actual += 1
	}
	return iter
}

func (iter *iterador[K, V]) HaySiguiente() bool {
	return iter.pos_actual < iter.diccionario.capacidad_total
}

func (iter *iterador[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.diccionario.elementos[iter.pos_actual].clave, iter.diccionario.elementos[iter.pos_actual].valor
}

func (iter *iterador[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	for i := iter.pos_actual + 1; i < iter.diccionario.capacidad_total; i += 1 {
		if iter.diccionario.elementos[i].estado == OCUPADO {
			iter.pos_actual = i
			return
		}
	}
	iter.pos_actual = iter.diccionario.capacidad_total + 1

}

// funciones auxiliares
// -----------------------------------------------------------------
func buscar_elemento[K comparable, V any](diccionario *hash[K, V], clave K) int {
	ind := fnvHashing(diccionario, clave, diccionario.capacidad_total)
	for i := ind; i < diccionario.capacidad_total+ind; i += 1 {
		pos := i % diccionario.capacidad_total
		if (diccionario.elementos[pos].estado == OCUPADO && diccionario.elementos[pos].clave == clave) || diccionario.elementos[pos].estado == VACIO {
			return pos
		}
	}
	return -1
}

func fnvHashing[K comparable, V any](diccionario *hash[K, V], clave K, capacidad int) int { //funcion sacada de https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/hash/fnv/fnv.go
	bytes := []byte(fmt.Sprintf("%v", clave))
	hash := fnv.New32a()
	hash.Write(bytes)
	return int(hash.Sum32()) % capacidad
}
func (diccionario *hash[K, V]) redimensionar(nueva_capacidad int) {
	copia_elementos := diccionario.elementos
	diccionario.elementos = crearArregloElementos[K, V](nueva_capacidad)
	diccionario.cantidad = 0
	diccionario.capacidad_total = nueva_capacidad
	for _, elem := range copia_elementos {
		if elem.estado == OCUPADO {
			diccionario.Guardar(elem.clave, elem.valor)
		}
	}
	diccionario.cant_mas_borrados = diccionario.cantidad
}
