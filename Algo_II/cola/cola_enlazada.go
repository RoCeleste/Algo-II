package cola

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

type nodo[T any] struct {
	//
	siguiente *nodo[T]
	elem      T
}

func (cola *colaEnlazada[T]) crear_Nodo(valor T) *nodo[T] {
	Nodo := new(nodo[T])
	//
	Nodo.siguiente = nil
	Nodo.elem = valor
	return Nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	Nodo := cola.crear_Nodo(elem)
	if cola.EstaVacia() {
		cola.primero = Nodo
	} else {
		//
		cola.ultimo.siguiente = Nodo
	}
	cola.ultimo = Nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	} else if cola.primero == cola.ultimo {
		retorno := cola.ultimo.elem
		cola.primero = nil
		cola.ultimo = nil
		return retorno
	}
	retorno := cola.primero.elem
	cola.primero = cola.primero.siguiente
	//
	return retorno
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.elem
}
func (cola *colaEnlazada[T]) EstaVacia() bool {

	return cola.primero == nil
}
