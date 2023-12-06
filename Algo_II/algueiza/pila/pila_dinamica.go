package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos     []T
	cantidad  int
	capacidad int // necesito este campo porque buscar esta capacidad funciona en tiempo constante, si uso cap(), me corre en O(n)
}

const tam_inicial = 10
const expandir = "+"
const contraer = "-"
const razon_expansion = 150 / 100
const razon_contraccion = 66 / 100
const tasa_ocupacion_para_expansion = 80 / 100
const tasa_ocupacion_para_contraccion = 50 / 100

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, tam_inicial)
	pila.cantidad = 0
	return pila
}
func (pila *pilaDinamica[T]) redimensionar(cambio string) {
	var nuevo_arreglo []T
	switch cambio {

	case expandir:
		nuevo_arreglo = make([]T, int(pila.capacidad*razon_expansion))

	case contraer:
		nuevo_arreglo = make([]T, int(pila.capacidad*razon_contraccion))
	default:
		panic("valor de redimension invalido")
	}
	nuevo_arreglo = pila.datos[:]
	pila.datos = nuevo_arreglo
}

func (pila *pilaDinamica[T]) Apilar(elem T) {

	if (pila.cantidad + 1) > pila.capacidad*tasa_ocupacion_para_expansion {

		pila.redimensionar(expandir)
	}
	pila.datos = append(pila.datos, elem)
	pila.cantidad += 1

}
func (pila *pilaDinamica[T]) Desapilar() T {

	if pila.EstaVacia() {

		panic("La pila esta vacia")
	}
	if (pila.cantidad - 1) < pila.capacidad*tasa_ocupacion_para_contraccion {

		pila.redimensionar(contraer)
	}

	elem_desapilado := pila.VerTope()
	pila.datos = pila.datos[0 : len(pila.datos)-1]
	pila.cantidad -= 1
	return elem_desapilado

}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[len(pila.datos)-1]
}
func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}
