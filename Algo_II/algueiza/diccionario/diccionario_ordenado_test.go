package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func arrayAleatorio() []int {
	var arr []int
	for i := 0; i < 1000; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < 10000; i++ {
		a := rand.Intn(999)
		b := rand.Intn(999)
		arr[a], arr[b] = arr[b], arr[a]
	}
	return arr

}

func buscarAbb(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func compararPrueba(entero1, entero2 int) int {
	return entero1 - entero2
}

func TestAbbVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararPrueba)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })
}

func TestAbbClaveDefault(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararPrueba)
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })

	abbStr := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abbStr.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbStr.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbStr.Borrar("") })
}

func TestAbbUnElement(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestAbbGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestAbbReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoAbbDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")
	var arr []int
	dic := TDADiccionario.CrearABB[int, int](compararPrueba)
	for i := 0; i < 1500; i++ {
		n := rand.Intn(1500)
		arr = append(arr, n)
		dic.Guardar(n, n)
	}
	for _, num := range arr {
		dic.Guardar(num, 2*num)
	}
	ok := true
	for _, num := range arr {
		ok = dic.Obtener(num) == 2*num
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestAbbBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestAbbReutlizacionDeBorrados(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestAbbConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compararPrueba)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestAbbClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestAbbValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestAbbCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestAbbIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.EqualValues(t, 0, buscarAbb(cs[0], claves))
	require.EqualValues(t, 1, buscarAbb(cs[1], claves))
	require.EqualValues(t, 2, buscarAbb(cs[2], claves))
	require.EqualValues(t, "Gato", cs[0])
	require.EqualValues(t, "Perro", cs[1])
	require.EqualValues(t, "Vaca", cs[2])
}

func TestAbbIteradorConRangoUnicoElemento(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	valor := "Gato"
	var hasta *string = &valor
	dic.IterarRango(nil, hasta, (func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	}))

	require.EqualValues(t, 1, cantidad)
	require.EqualValues(t, 0, buscarAbb(cs[0], claves))
	require.EqualValues(t, -1, buscarAbb(cs[1], claves))
	require.EqualValues(t, -1, buscarAbb(cs[2], claves))
	require.EqualValues(t, "Gato", cs[0])
	require.EqualValues(t, "", cs[1])
	require.EqualValues(t, "", cs[2])
}

func TestAbbIteradorConAlgunosElementos(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad
	ini := "Gato"
	valor := "Perro"
	var desde *string = &ini
	var hasta *string = &valor
	dic.IterarRango(desde, hasta, (func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	}))

	require.EqualValues(t, 2, cantidad)
	require.EqualValues(t, 0, buscarAbb(cs[0], claves))
	require.EqualValues(t, 1, buscarAbb(cs[1], claves))
	require.EqualValues(t, -1, buscarAbb(cs[2], claves))
	require.EqualValues(t, "Gato", cs[0])
	require.EqualValues(t, "Perro", cs[1])
	require.EqualValues(t, "", cs[2])
}

func TestAbbIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestAbbIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestAbbIteradorInternoCondicionCorte(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer despues de la condicion de corte")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	cantidad := 0
	ptrCantidad := &cantidad
	dic.Iterar(func(clave string, dato int) bool {
		*ptrCantidad += dato
		return clave < "Golondrina"
	})

	require.EqualValues(t, 15, cantidad)
}

func TestIterarAbbVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestAbbIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, primer_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarAbb(primero, claves))
	require.EqualValues(t, "miau", primer_valor)

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarAbb(segundo, claves))
	require.EqualValues(t, "guau", segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, tercer_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarAbb(tercero, claves))
	require.EqualValues(t, "moo", tercer_valor)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestAbbIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscarAbb(primero, claves))
	require.NotEqualValues(t, -1, buscarAbb(segundo, claves))
	require.NotEqualValues(t, -1, buscarAbb(tercero, claves))
}

func TestPruebaAbbIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}
func TestAbbIteradorCorteEnRango(t *testing.T) {
	t.Log("Prueba de iterador interno, para validar que siempre que" +
		"la función visitar de false, se corte la iteracion, o hasta que se itere entre el rango indicado")
	dic := TDADiccionario.CrearABB[int, int](compararPrueba)
	dic.Guardar(100, 100)
	dic.Guardar(50, 50)
	dic.Guardar(150, 150)
	dic.Guardar(25, 25)
	dic.Guardar(75, 75)
	dic.Guardar(125, 125)
	dic.Guardar(200, 200)
	dic.Guardar(10, 10)
	dic.Guardar(30, 30)
	dic.Guardar(60, 60)
	dic.Guardar(80, 80)
	dic.Guardar(110, 110)
	dic.Guardar(130, 130)
	dic.Guardar(190, 190)
	dic.Guardar(300, 300)

	ini := 60
	valor := 90
	var desde *int = &ini
	var hasta *int = &valor
	iter := dic.IteradorRango(desde, hasta)
	clave1, _ := iter.VerActual()
	require.EqualValues(t, clave1, 60)
	iter.Siguiente()

	clave2, _ := iter.VerActual()
	require.EqualValues(t, clave2, 75)
	iter.Siguiente()

	clave3, _ := iter.VerActual()
	require.EqualValues(t, clave3, 80)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestAbbIteradorSinHasta(t *testing.T) {
	t.Log("Prueba de iterador interno, para validar que siempre que" +
		"la función visitar de false, se corte la iteracion, o hasta que se itere entre el rango indicado")

	dic := TDADiccionario.CrearABB[int, int](compararPrueba)
	dic.Guardar(100, 100)
	dic.Guardar(50, 50)
	dic.Guardar(150, 150)
	dic.Guardar(75, 75)

	ini := 60
	var desde *int = &ini
	iter := dic.IteradorRango(desde, nil)
	clave1, _ := iter.VerActual()
	require.EqualValues(t, clave1, 75)
	iter.Siguiente()

	clave2, _ := iter.VerActual()
	require.EqualValues(t, clave2, 100)
	iter.Siguiente()

	clave3, _ := iter.VerActual()
	require.EqualValues(t, clave3, 150)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestAbbIteradorSinDesde(t *testing.T) {
	t.Log("Prueba de iterador interno, para validar que siempre que" +
		"la función visitar de false, se corte la iteracion, o hasta que se itere entre el rango indicado")

	dic := TDADiccionario.CrearABB[int, int](compararPrueba)
	dic.Guardar(100, 100)
	dic.Guardar(50, 50)
	dic.Guardar(150, 150)
	dic.Guardar(75, 75)

	valor := 100
	var hasta *int = &valor
	iter := dic.IteradorRango(nil, hasta)

	clave1, _ := iter.VerActual()
	require.EqualValues(t, clave1, 50)
	iter.Siguiente()

	clave2, _ := iter.VerActual()
	require.EqualValues(t, clave2, 75)
	iter.Siguiente()

	clave3, _ := iter.VerActual()
	require.EqualValues(t, clave3, 100)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestAbbVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compararPrueba)
	arr := arrayAleatorio()
	for _, valor := range arr {
		dic.Guardar(valor, valor)
	}
	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestEjecutarPruebaVolumenAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](compararPrueba)
	claves := arrayAleatorio()
	valores := claves
	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(t, 1000, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < 1000; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(t, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(t, 1000, dic.Cantidad(), "La cantidad de elementos es incorrecta")
}