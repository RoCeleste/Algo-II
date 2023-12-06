package cola_prioridad

const _CAP_INICIAL = 10
const _DUPLICAR = 2
const _TAZA_OCUPACION_CONTRACCION = 0.25
const _RAZON_CONTRACCION = 0.5

// tipos
// --------------------------------------------------------------------------------------------

type heap[T any] struct {
	tabla       []T
	comparacion func(T, T) int
	cantidad    int
}

// creadores
// --------------------------------------------------------------------------------------------

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {

	Heap := new(heap[T])
	Heap.tabla = make([]T, _CAP_INICIAL)
	Heap.comparacion = funcion_cmp
	return Heap

}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {

	Heap := new(heap[T])
	if len(arreglo) < _CAP_INICIAL {
		Heap.tabla = make([]T, _CAP_INICIAL)
	} else {
		Heap.tabla = make([]T, len(arreglo)*_DUPLICAR)
	}
	copy(Heap.tabla, arreglo)
	Heap.cantidad = len(arreglo)
	Heap.comparacion = funcion_cmp

	heapify(Heap.tabla, Heap.cantidad, Heap.comparacion)

	return Heap
}

// primitivas
// -----------------------------------------------------------------------------------------------

func (Heap *heap[T]) EstaVacia() bool {

	return Heap.cantidad == 0

}
func (Heap *heap[T]) Encolar(elemento T) {
	if (Heap.cantidad + 1) > cap(Heap.tabla) {
		Heap.redimensionar(cap(Heap.tabla) * _DUPLICAR)
	}

	Heap.tabla[Heap.cantidad] = elemento
	Heap.cantidad++
	upheap(Heap.cantidad-1, Heap.tabla, Heap.comparacion)

}

func heapify[T any](tabla []T, largo int, funcion_cmp func(T, T) int) {
	for i := largo / 2; i > -1; i-- {
		downheap(i, tabla, funcion_cmp, largo)
	}

}

func upheap[T any](indice int, tabla []T, funcion_cmp func(T, T) int) {
	if indice == 0 {
		return
	}

	for funcion_cmp(tabla[indice], tabla[obtener_xadre(indice)]) > 0 {
		swap(tabla, indice, obtener_xadre(indice))
		indice = obtener_xadre(indice)
	}
}

func (Heap *heap[T]) VerMax() T {

	if Heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return Heap.tabla[0]
}

func (Heap *heap[T]) Desencolar() T {

	if Heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	if (Heap.cantidad) <= int(float64(cap(Heap.tabla))*_TAZA_OCUPACION_CONTRACCION) && cap(Heap.tabla) > _CAP_INICIAL {
		Heap.redimensionar(int(float64(cap(Heap.tabla)) * _RAZON_CONTRACCION))
	}

	var valor_defecto T
	dato_a_devolver := Heap.tabla[0]
	swap(Heap.tabla, 0, Heap.cantidad-1)
	Heap.tabla[Heap.cantidad-1] = valor_defecto
	downheap(0, Heap.tabla, Heap.comparacion, Heap.cantidad-1)
	Heap.cantidad--

	return dato_a_devolver
}

func downheap[T any](indice int, tabla []T, funcion_cmp func(T, T) int, largo int) {

	izq := 2*indice + 1
	der := 2*indice + 2
	max := indice

	if izq < largo && funcion_cmp(tabla[indice], tabla[izq]) < 0 {
		max = izq
	}
	if der < largo && funcion_cmp(tabla[max], tabla[der]) < 0 {
		max = der
	}
	if max != indice {
		swap(tabla, indice, max)
		downheap(max, tabla, funcion_cmp, largo)
	}
}

func (Heap *heap[T]) Cantidad() int {
	return Heap.cantidad
}

// funciones auxiliares
// ---------------------------------------------------------------------------------------------------

func swap[T any](tabla []T, x, y int) {

	tabla[x], tabla[y] = tabla[y], tabla[x]
}

func obtener_xadre(indice int) int {
	return (indice - 1) / 2
}

func (Heap *heap[T]) redimensionar(capacidad_nueva int) {
	datos_redimensionados := make([]T, capacidad_nueva)
	copy(datos_redimensionados, Heap.tabla[:])
	Heap.tabla = datos_redimensionados
}

// heapsort
//----------------------------------------------------------------------------------------------------------

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

	heapify(elementos, len(elementos), funcion_cmp)

	posMax := 0
	ultimoRelativo := len(elementos) - 1

	for posMax != ultimoRelativo && len(elementos) != 0 {
		posMax := 0
		swap(elementos, posMax, ultimoRelativo)
		downheap(0, elementos, funcion_cmp, ultimoRelativo)
		ultimoRelativo--
	}

}
