package diccionario

import (
	"algueiza/pila"
)

// tipos
// ----------------------------------------------------------------------------------------

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iteradorAbb[K comparable, V any] struct {
	pila          pila.Pila[*nodoAbb[K, V]]
	desde         *K
	hasta         *K
	f_comparacion func(K, K) int
}

// metodos creadores
// ---------------------------------------------------------------------------------------
func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	v := new(abb[K, V])
	v.cmp = funcion_cmp
	return v
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	b := new(nodoAbb[K, V])
	b.clave = clave
	b.dato = dato
	return b

}

// primitivas de abb
//-----------------------------------------------------------------------------------------

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	if abb.raiz == nil {
		abb.raiz = crearNodo(clave, valor)
		abb.cantidad++
		return
	}
	nodo_padre, _ := abb.buscar_nodo(abb.raiz, clave, nil)
	comparacion_nodos := abb.cmp(nodo_padre.clave, clave)
	if comparacion_nodos == 0 {
		nodo_padre.dato = valor
	}
	if comparacion_nodos < 0 {
		nodo_padre.derecho = crearNodo(clave, valor)
		abb.cantidad++
	}
	if comparacion_nodos > 0 {
		nodo_padre.izquierdo = crearNodo(clave, valor)
		abb.cantidad++
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	if abb.raiz == nil {
		return false
	}
	nodo, _ := abb.buscar_nodo(abb.raiz, clave, nil)
	return nodo.clave == clave
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}
	nodo, _ := abb.buscar_nodo(abb.raiz, clave, nil)
	if nodo.clave != clave {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (abb *abb[K, V]) buscarminimoderecha(nodo *nodoAbb[K, V]) K {
	if nodo.izquierdo == nil {
		return nodo.clave
	}
	return abb.buscarminimoderecha(nodo.izquierdo)
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}
	nodo, nodo_padre := abb.buscar_nodo(abb.raiz, clave, abb.raiz)
	if nodo == nil || nodo.clave != clave {
		panic("La clave no pertenece al diccionario")
	}
	a_devolver := nodo.dato
	if nodo.izquierdo != nil && nodo.derecho != nil {
		clave_remplazante := abb.buscarminimoderecha(nodo.derecho)
		dato_remplanzante := abb.Borrar(clave_remplazante)
		nodo.clave = clave_remplazante
		nodo.dato = dato_remplanzante
		return a_devolver
	}
	comparacion_nodos := abb.cmp(nodo.clave, nodo_padre.clave)
	if comparacion_nodos < 0 {
		if nodo.izquierdo == nil {
			nodo_padre.izquierdo = nodo.derecho
		}
		if nodo.derecho == nil {
			nodo_padre.izquierdo = nodo.izquierdo
		}
	}
	if comparacion_nodos == 0 { //quiero borrar la raiz
		if nodo.izquierdo == nil {
			abb.raiz = nodo.derecho
		}
		if nodo.derecho == nil {
			abb.raiz = nodo.izquierdo
		}
	}
	if comparacion_nodos > 0 {
		if nodo.izquierdo == nil {
			nodo_padre.derecho = nodo.derecho

		}
		if nodo.derecho == nil {
			nodo_padre.derecho = nodo.izquierdo
		}
	}
	abb.cantidad--
	return a_devolver
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (nodo *nodoAbb[K, V]) iterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool, f_comparacion func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || f_comparacion(*desde, nodo.clave) < 0 {
		if !nodo.izquierdo.iterarRango(desde, hasta, visitar, f_comparacion) {
			return false
		}
	}
	if (desde == nil || f_comparacion(*desde, nodo.clave) <= 0) && (hasta == nil || f_comparacion(*hasta, nodo.clave) >= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || f_comparacion(*hasta, nodo.clave) > 0 {
		if !nodo.derecho.iterarRango(desde, hasta, visitar, f_comparacion) {
			return false
		}
	}
	return true

}

func (abb abb[K, V]) Iterar(f func(clave K, valor V) bool) {
	abb.IterarRango(nil, nil, f)
}

func (abb abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarRango(desde, hasta, visitar, abb.cmp)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iteradorAbb[K, V])
	iter.desde = desde
	iter.hasta = hasta
	iter.f_comparacion = abb.cmp
	iter.pila = pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.irTodoIzquierdaRango(abb.raiz)
	return iter
}

func (iter *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()

}

func (iter *iteradorAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iteradorAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	iter.irTodoIzquierdaRango(nodo.derecho)
}

// funciones auxiliares
// ------------------------------------------------------------------------------------------
func (abb *abb[K, V]) buscar_nodo(nodo *nodoAbb[K, V], clave K, anterior *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	comparacion_nodos := abb.cmp(nodo.clave, clave)
	if comparacion_nodos < 0 {
		if nodo.derecho == nil {
			return nodo, anterior
		}
		return abb.buscar_nodo(nodo.derecho, clave, nodo)
	}
	if comparacion_nodos == 0 {
		return nodo, anterior
	}
	if nodo.izquierdo == nil {
		return nodo, anterior
	}
	return abb.buscar_nodo(nodo.izquierdo, clave, nodo)
}

func (iter *iteradorAbb[K, V]) irTodoIzquierdaRango(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if (iter.desde == nil || iter.f_comparacion(*iter.desde, nodo.clave) <= 0) && (iter.hasta == nil || iter.f_comparacion(*iter.hasta, nodo.clave) >= 0) {
		iter.pila.Apilar(nodo)
	}
	if iter.desde == nil || iter.f_comparacion(*iter.desde, nodo.clave) <= 0 {
		iter.irTodoIzquierdaRango(nodo.izquierdo)
	} else {
		iter.irTodoIzquierdaRango(nodo.derecho)
	}
}