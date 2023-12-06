package lista

// tipos
//-----------------------------------------------------------------------------------------------------

type listaEnlazada[T any] struct {
	primero  *nodo[T]
	ultimo   *nodo[T]
	cantidad int
}

type nodo[T any] struct {
	elem    T
	proximo *nodo[T]
}

type iteradorLista[T any] struct {
	anterior *nodo[T]
	actual   *nodo[T]
	de_lista *listaEnlazada[T]
}

// creadores
//-----------------------------------------------------------------------------------------------------

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (lista *listaEnlazada[T]) crear_Nodo(valor T) *nodo[T] {
	Nodo := new(nodo[T])
	Nodo.elem = valor
	return Nodo
}

// metodos de listaEnlazada
// -----------------------------------------------------------------------------------------------------
func (lista *listaEnlazada[T]) EstaVacia() bool {

	return lista.cantidad == 0 && lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {

	Nodo := lista.crear_Nodo(elem)
	if lista.EstaVacia() {
		lista.ultimo = Nodo

	} else {
		Nodo.proximo = lista.primero

	}

	lista.primero = Nodo
	lista.cantidad++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {

	Nodo := lista.crear_Nodo(elem)
	if lista.EstaVacia() {
		lista.primero = Nodo

	} else {
		lista.ultimo.proximo = Nodo

	}

	lista.ultimo = Nodo
	lista.cantidad++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	a_devolver := lista.primero.elem
	lista.primero = lista.primero.proximo
	if lista.cantidad == 1 {
		lista.ultimo = lista.primero
	}
	lista.cantidad--
	return a_devolver
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.elem
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.elem
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.cantidad
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil && visitar(actual.elem) {
		actual = actual.proximo
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorLista[T])
	iter.actual = lista.primero
	iter.de_lista = lista
	return iter
}

// metodos de IteradorLista
//-----------------------------------------------------------------------------------------------------

func (iter *iteradorLista[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.elem
}

func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.anterior.proximo
}

func (iter *iteradorLista[T]) Insertar(elem T) {
	Nodo := iter.de_lista.crear_Nodo(elem)
	if !iter.HaySiguiente() { //caso quiero insertar ultimo elemento
		iter.de_lista.ultimo = Nodo
	}
	if iter.anterior == nil { //caso insertar en lista vacia || caso insertar al principio (lista no vacia)
		iter.de_lista.primero = Nodo
	} else { //caso no querer insertar en el primero
		iter.anterior.proximo = Nodo
	}

	Nodo.proximo = iter.actual
	iter.actual = Nodo
	iter.de_lista.cantidad++
}

func (iter *iteradorLista[T]) borrarPrimeroIterador() {
	iter.de_lista.primero = iter.de_lista.primero.proximo
	if iter.de_lista.cantidad == 1 {
		iter.de_lista.ultimo = iter.de_lista.primero
	}
	iter.actual = iter.actual.proximo
}

func (iter *iteradorLista[T]) borrarUltimoIterador() {
	iter.actual = iter.actual.proximo
	iter.de_lista.ultimo = iter.anterior
	iter.de_lista.ultimo.proximo = iter.actual
}

func (iter *iteradorLista[T]) Borrar() T {
	if !iter.HaySiguiente() { //iterador esta en el nil
		panic("El iterador termino de iterar")
	}
	a_devolver := iter.VerActual()
	if iter.actual == iter.de_lista.primero {
		iter.borrarPrimeroIterador()
	} else if iter.actual.proximo == nil {
		iter.borrarUltimoIterador()
	} else {
		iter.anterior.proximo = iter.actual.proximo
		iter.actual = iter.actual.proximo
	}

	iter.de_lista.cantidad--
	return a_devolver
}
