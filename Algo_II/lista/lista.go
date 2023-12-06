package lista

type Lista[T any] interface {

	// EstaVacia devuelve un bool: true si no tiene ningun elemento (nil), false si tiene al menos un elemento
	// O(1)
	EstaVacia() bool

	// InsertarPrimero agrega el elemento T al inicio de la Lista
	// O(1)
	InsertarPrimero(T)

	// InsertarUltimo agrega el elemento T a al final de la Lista
	// O(1)
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la Lista y lo devuelve.
	// En caso de estar vacia, entra en panico con el mensaje "La lista esta vacia"
	// O(1)
	BorrarPrimero() T

	// VerPrimero obtiene el primer elemento de la Lista.
	// En caso de estar vacia, entra en panico con el mensaje "La lista esta vacia"
	// O(1)
	VerPrimero() T

	// VerUltimo obtiene el ultimo elemento de la Lista.
	// En caso de estar vacia, entra en panico con el mensaje "La lista esta vacia"
	// O(1)
	VerUltimo() T

	// Largo obtiene la cantidad de elementos de la Lista.
	// O(1)
	Largo() int

	// Iterar aplica la funcion recibida a todos los elementos de la Lista, hasta que visitar(T) devuelva false o Iterar haya recorrido toda la lista
	// En caso de estar vacia, entra en panico con el mensaje "La lista esta vacia"
	// O(n)
	Iterar(visitar func(T) bool)

	// Iterador devuelve una instancia de IteradorLista
	// O(1)
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el elemento actual del ciclo de iteracion
	// En caso de que el elemento actual sea nil (ya sea porque la lista esta vacia o porque se recorrieron todos los elementos), entra en panico con un mensaje correspondiente
	// O(1)
	VerActual() T

	// HaySiguiente indica Si existe un elemento proximo en el ciclo de iteracion.
	// O(1)
	HaySiguiente() bool

	// Siguiente avanza al siguiente elemento en el ciclo de iteracion
	// En caso de no haber un siguiente, entra en panico con el mensaje "El iterador termino de iterar"
	// O(1)
	Siguiente()

	// Insertar coloca el elemento T en la posicion entre el actual y el siguiente del ciclo de iteracion, esto NO devuelve una copia de la lista actualizada, sino que modifica la original.
	// O(1)
	Insertar(T)

	// Borrar elimina el elemento en la posicion actual del ciclo de iteracion, y lo devuelve
	// En caso de que el elemento actual sea nil (ya sea porque la lista esta vacia o porque se recorrieron todos los elementos), entra en panico con un mensaje correspondiente
	// O(1)
	Borrar() T
}
