package grafo

import (
	hash "grafo/diccionario"
	pila "grafo/pila"
	cola "grafo/cola"
	heap "grafo/cola_prioridad"
	"fmt"
	"strings"
)

// TDA Grafo (1ra parte)
//-----------------------------------------------------------------------------------------------------------------------
type Grafo[T comparable] struct {
	es_dirigido		bool
	vertices		[]T
	aristas			[]int
	adyacentes		hash.Diccionario[T, hash.Diccionario[T, int]]
}

func CrearGrafo[T comparable](es_dirigido bool) *Grafo[T] {
	g := new(Grafo[T])
	g.es_dirigido = es_dirigido
	g.vertices = make([]T, 0)
	g.aristas = make([]int, 0)
	g.adyacentes = hash.CrearHash[T, hash.Diccionario[T, int]]()
	return g
}

func (g *Grafo[T]) AgregarVertice(v T) {
	if !g.adyacentes.Pertenece(v) {
		g.vertices = append(g.vertices, v)
		g.adyacentes.Guardar(v, hash.CrearHash[T, int]())	
	}
}

func (g *Grafo[T]) AgregarArista(v T, w T, peso int) {
	g.adyacentes.Obtener(v).Guardar(w, peso)
	g.aristas = append(g.aristas, peso)
	if !g.es_dirigido {
		g.adyacentes.Obtener(w).Guardar(v, peso)
	}
}

func (g *Grafo[T]) ObtenerVertices() []T {
	return g.vertices
}

func (g *Grafo[T]) ObtenerAristas() []int {
	return g.aristas
}

func (g *Grafo[T]) Peso(v T, w T) int {
	return g.adyacentes.Obtener(v).Obtener(w)
}

func (g *Grafo[T]) ObtenerAdyacentes(v T) []T {
	lista := []T {}
	iter := g.adyacentes.Obtener(v).Iterador()
	for iter.HaySiguiente() {
		ad, _ := iter.VerActual()
		lista = append(lista, ad)
		iter.Siguiente()
	}
	return lista
}

// Funciones de grafo (2da parte)
// --------------------------------------------------------------------------------------------------

func BFS[T comparable](g *Grafo[T], origen T, visitados hash.Diccionario[T, int]) {
	cola := cola.CrearColaEnlazada[T]()
	cola.Encolar(origen)
	visitados.Guardar(origen, 0)
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, w := range g.ObtenerAdyacentes(v) {
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, 0)
				cola.Encolar(w)
			}
		}
	}
}

func Es_Conexo[T comparable](g *Grafo[T]) bool {
	visitados := hash.CrearHash[T, int]()
	cant := 0
	for _, v := range g.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			cant++
			if cant == 2 {
				return false
			}
			BFS(g, v, visitados)
		}
	}
	return true
}

func Tiene_ciclos[T comparable](g *Grafo[T]) bool {
	visitados := hash.CrearHash[T, int]()
	xadres := hash.CrearHash[T, T]()
	for _, v := range g.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			if !tiene_ciclos_util(g, v, visitados, xadres) {
				return false
			}
		}
	}
	return true
}

func tiene_ciclos_util[T comparable](g *Grafo[T], v T, visitados hash.Diccionario[T, int], xadres hash.Diccionario[T, T]) bool {
	visitados.Guardar(v, 0)
	for _, ad := range g.ObtenerAdyacentes(v) {
		if visitados.Pertenece(ad) {
			if ad != xadres.Obtener(v) {
				return true
			} 
		} else {
			xadres.Guardar(ad, v)
			return tiene_ciclos_util(g, ad, visitados, xadres)
		}

	}
	return false

}

func OrdenTopologico[T comparable](g *Grafo[T]) []T {
	visitados := hash.CrearHash[T, int]()
	Pila := pila.CrearPilaDinamica[T]()
	for _, v := range g.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, 0)
			orden_topologico_util(g, v, visitados, Pila)
		}
	}
	lista := []T{}
	for !Pila.EstaVacia() {
		lista = append(lista, Pila.Desapilar())
	}
	return lista
}

func orden_topologico_util[T comparable](g *Grafo[T], v T, visitados hash.Diccionario[T, int], Pila pila.Pila[T]) {
	for _, ad := range g.ObtenerAdyacentes(v) {
		if !visitados.Pertenece(ad) {
			visitados.Guardar(ad, 0)
			orden_topologico_util(g, ad, visitados, Pila)
		}
	}
	Pila.Apilar(v)
}

func invertir_lista[T comparable](lista []T) []T {
	for i, j := 0, len(lista)-1; i<j; i, j = i+1, j-1 {
      lista[i], lista[j] = lista[j], lista[i]
   }
   return lista
}

func main() {
	gr := CrearGrafo[string](false)
	gr.AgregarVertice("malena")
	gr.AgregarVertice("fiorella")
	gr.AgregarVertice("marcos")
	gr.AgregarVertice("valentina")
	gr.AgregarVertice("emiliano")
	gr.AgregarArista("fiorella", "malena", 5)
	gr.AgregarVertice("lorena")
	gr.AgregarArista("marcos", "lorena", 6)
	gr.AgregarArista("lorena", "valentina", 2)
	fmt.Println(gr.ObtenerVertices(), " ---> ", gr.ObtenerAristas())
	gr.AgregarArista("lorena", "malena", 3)
	for _, ad := range gr.ObtenerAdyacentes("lorena") {
		fmt.Println(ad)
	}
	gr.AgregarArista("malena", "fiorella", 1)
	gr.AgregarArista("fiorella", "marcos", 1)
	gr.AgregarArista("marcos", "valentina", 1)
	gr.AgregarArista("valentina", "emiliano", 1)
	gr.AgregarArista("emiliano", "malena", 1)
	gr.AgregarArista("emiliano", "valentina", 3)
	gr.AgregarArista("emiliano", "malena", 4)
	fmt.Println(Es_Conexo(gr))
	fmt.Println(OrdenTopologico(gr))
	l := invertir_lista(OrdenTopologico(gr))
	fmt.Println(l)
	fmt.Println(Dijkstra(gr, "malena", "valentina"))
}


func Dijkstra(g *Grafo[string], origen string, destino string) ([]string, int){
	xadres, _ := dijkstra_util(g, origen, destino)
	if xadres == nil {
		return nil, 0
	}
	camino := []string{}
	peso := 0
	actual := destino
	for actual != "" {								
		camino = append(camino, actual)
		if xadres.Obtener(actual) != "" {			
			peso += g.Peso(xadres.Obtener(actual), actual)
		}
		actual = xadres.Obtener(actual)
	}
	invertir_lista(camino)
	return camino, peso
}

func dijkstra_util(g *Grafo[string], origen string, destino string) (hash.Diccionario[string, string], hash.Diccionario[string, int]) {
	xadres := hash.CrearHash[string, string]()
	distancias := hash.CrearHash[string, int]()

	for _, v := range g.ObtenerVertices() {
		distancias.Guardar(v, 9000000000000)
		xadres.Guardar(v, "")						
	}
	xadres.Guardar(origen, "")						
	distancias.Guardar(origen, 0)						
	heap := heap.CrearHeap[string](strings.Compare)									
	heap.Encolar(origen)

	for !heap.EstaVacia() {
		v := heap.Desencolar()
		if v == destino {
			return xadres, distancias
		}

		for _, w := range g.ObtenerAdyacentes(v) {
			nueva_distancia := distancias.Obtener(v) + g.Peso(v, w)
			if nueva_distancia < distancias.Obtener(w) {
				distancias.Guardar(w, nueva_distancia)
				xadres.Guardar(w, v)
				heap.Encolar(w)
			}

		}
	}
	return nil, nil
}

// convertir todo en Dijkstra a T nuevamente
