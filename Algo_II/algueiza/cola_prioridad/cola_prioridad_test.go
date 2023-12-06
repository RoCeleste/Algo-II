package cola_prioridad_test

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	TDAHeap "cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func arrayAleatorio(cant int) []int {
	var arr []int
	for i := 0; i < cant; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < cant*10; i++ {
		a := rand.Intn(cant - 1)
		b := rand.Intn(cant - 1)
		arr[a], arr[b] = arr[b], arr[a]
	}
	return arr

}

func compMayorAMenor(a int, b int) int {
	return a - b
}

func compMenorAMayor(a int, b int) int {
	return b - a
}

func Test_heap_vacio(t *testing.T) {
	heap := TDAHeap.CrearHeap(compMayorAMenor)
	require.True(t, heap.EstaVacia())
}

func Test_panic_de_operaciones_en_heap_vacio(t *testing.T) {
	heap := TDAHeap.CrearHeap(compMayorAMenor)
	require.Panics(t, func() { heap.Desencolar() })
	require.Panics(t, func() { heap.VerMax() })
}

func Test_encolar_un_elemento(t *testing.T) {
	heap := TDAHeap.CrearHeap(strings.Compare)
	heap.Encolar("Messi")
	require.Equal(t, "Messi", heap.VerMax())
	require.Equal(t, "Messi", heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func Test_encolar_varios_elementos(t *testing.T) {
	heap := TDAHeap.CrearHeap(compMayorAMenor)
	for i := 10; i > 0; i-- {
		heap.Encolar(i)
	}
	require.Equal(t, 10, heap.Desencolar())
	for i := 9; i > 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}
func Test_encolar_varios_elementos_heap_min(t *testing.T) {
	heap := TDAHeap.CrearHeap(compMenorAMayor)
	for i := 20; i > 0; i-- {
		heap.Encolar(i)
	}
	require.Equal(t, 1, heap.Desencolar())
	for i := 2; i <= 20; i++ {
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())

}

func Test_volumen_y_cantidad_correcta(t *testing.T) {
	heap := TDAHeap.CrearHeap(compMayorAMenor)
	for i := 1; i <= 1000000; i++ {
		heap.Encolar(i)
	}
	require.Equal(t, 1000000, heap.Cantidad())
	for i := 1000000; i >= 1; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())

}

func TestHeapCrearArrVacio(t *testing.T) {
	arreglo := []string{}
	heap := TDAHeap.CrearHeapArr(arreglo, strings.Compare)
	heap.Encolar("Messi")
	require.Equal(t, "Messi", heap.VerMax())
	heap.Desencolar()
	require.Panics(t, func() { heap.Desencolar() })
	require.Panics(t, func() { heap.VerMax() })

}

func TestHeapElementosRepetidos(t *testing.T) {
	heap := TDAHeap.CrearHeap(strings.Compare)
	heap.Encolar("Messi")
	heap.Encolar("Messi")
	require.Equal(t, "Messi", heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, "Messi", heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())
	require.Panics(t, func() { heap.Desencolar() })
	require.Panics(t, func() { heap.VerMax() })

}

func Test_modo_de_creacion_a_partir_de_arreglo(t *testing.T) {
	arreglo := []string{"Coca Cola", "Avon", "Nivea", "Terrabusi", "Marolio", "Philips", "Renault", "Maybelline", "Samsung", "MercadoLibre", "La Virginia", "Knorr"}
	heap := TDAHeap.CrearHeapArr(arreglo, strings.Compare)
	require.Equal(t, "Terrabusi", heap.Desencolar())
	require.Equal(t, "Samsung", heap.VerMax())
}

func Test_heapsort(t *testing.T) {
	arreglo := []string{"Coca Cola", "Avon", "Nivea", "Terrabusi", "Marolio", "Philips", "Renault", "Maybelline", "Samsung", "MercadoLibre", "La Virginia", "Knorr"}
	arrOrdenado := make([]string, len(arreglo))
	copy(arrOrdenado, arreglo)
	sort.Strings(arrOrdenado)
	TDAHeap.HeapSort(arreglo, strings.Compare)
	for i := range arreglo {
		require.Equal(t, arreglo[i], arrOrdenado[i])
	}
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	heapMax := TDAHeap.CrearHeap(compMayorAMenor)
	heapMin := TDAHeap.CrearHeap(compMenorAMayor)

	//genero arreglo ordenado de largo n
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	/* Genero arreglo desordenado de largo n*/
	arregloMezclado := arrayAleatorio(n)

	for _, valor := range arregloMezclado {
		heapMax.Encolar(valor)
		heapMin.Encolar(valor)
	}

	require.EqualValues(b, n, heapMax.Cantidad(), "La cantidad de elementos es incorrecta")
	require.EqualValues(b, n, heapMin.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	ok2 := true

	/* Verifica que borre los valores correctos */
	for i := n - 1; i > -1; i-- {
		ok = heapMax.Desencolar() == arr[i]
		if !ok {
			break
		}
	}

	for j := 0; j < n; j++ {
		ok2 = heapMin.Desencolar() == arr[j]
		if !ok2 {
			break
		}
	}
	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.True(b, ok2, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, heapMax.Cantidad())
	require.EqualValues(b, 0, heapMin.Cantidad())
}

func BenchmarkHeap(b *testing.B) {
	b.Log("Prueba de stress del heap. Prueba guardando distinta cantidad de elementos (muy grandes, enorden aleatorio), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad sea la adecuada. " +
		"Luego validamos que podemos encolar cada elemento generado, y que luego podemos desencolar sin problemas. " +
		"Se generan tanto para heap de máximo como para heap de mínimo.")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}
