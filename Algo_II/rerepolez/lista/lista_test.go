package lista_test

import (
	TDALista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Lista_Vacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func Test_Volumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 1000000
	for i := 1; i <= tam; i++ {
		lista.InsertarPrimero(i)
	}
	for i := tam; i >= 1; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func Test_Invalidez_Lista_Vacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
}

func Test_se_Inserta_Al_Principio_Al_Crear_Iterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(4)
	require.Equal(t, 4, lista.VerPrimero())
}

func Test_se_Inserta_Al_Medio_Iterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(0)
	iter.Insertar(1)
	iter.Insertar(3)
	iter.Insertar(4)
	tam := 4
	for i := tam; i > tam/2; i-- { //itero hasta la mitad
		iter.Siguiente()
	}
	iter.Insertar(2)

	iter2 := lista.Iterador()
	for i := tam; i >= 0; i-- {
		require.Equal(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

func Test_se_Inserta_Al_Final_Al_Terminar_Iterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Insertar(2)
	iter.Insertar(3)
	iter.Insertar(4)
	tam := 4
	for i := tam; i >= 1; i-- {
		require.Equal(t, i, iter.VerActual())
		iter.Siguiente()
	}
	iter.Insertar(0)
	require.Equal(t, 0, lista.VerUltimo())

	iter2 := lista.Iterador()
	for i := tam; i >= 0; i-- {
		require.Equal(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

func Test_Eliminar_Al_Principio_Iterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.Equal(t, 1, iter.Borrar())
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())
}

func Test_Iterador_Borra_Al_Final(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	//itero hasta el final y borro el ultimo
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 3, iter.VerActual())
	require.Equal(t, 3, iter.Borrar())

	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	//itero y borro menos el ultimo
	iter2 := lista.Iterador()
	require.Equal(t, 1, iter2.VerActual())
	require.Equal(t, 1, iter2.Borrar())
	//pruebo q el iterador esta parado en el ultimo (y unico) elemento y que el primero de la lista == ultimo de la lista
	require.Equal(t, 2, iter2.VerActual())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
}

func Test_Eliminar_Al_Ultimo_Iterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Insertar(2)
	iter.Insertar(3)
	iter.Insertar(4)
	tam := 4
	for i := tam; i > 0; i-- {
		require.Equal(t, i, iter.VerActual())
		iter.Siguiente()
	}
	iter.Insertar(0)
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, 0, iter.Borrar())
	require.Equal(t, 1, lista.VerUltimo())
}

func Test_Se_Inserta_Despues_De_Borrar_Todo_Iterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(3)
	iter.Insertar(2)
	iter.Insertar(1)
	require.Equal(t, 1, iter.Borrar())
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 3, iter.Borrar())
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { iter.Borrar() })
	require.True(t, lista.EstaVacia())
	iter.Insertar(25)
	require.Equal(t, 25, lista.VerPrimero())
	require.Equal(t, 25, lista.VerUltimo())
	require.False(t, lista.EstaVacia())

}

func Test_Iterador_Borra_En_El_Medio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(77)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	tam := 5
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 77, iter.Borrar())
	require.Equal(t, 3, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 4, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

	iter2 := lista.Iterador()
	for i := 1; i > tam; i++ {
		require.Equal(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

func Test_Iterador_Borra_En_El_Medio_Luego_Itera_Con_Otro_Iterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	iter.Siguiente()
	require.Equal(t, 2, iter.Borrar())
	iter2 := lista.Iterador()
	require.Equal(t, 1, iter2.VerActual())
	iter2.Siguiente()
	require.Equal(t, 3, iter2.VerActual())
	iter2.Siguiente()
	require.False(t, iter2.HaySiguiente())
}

func Test_Insertar_Y_Borrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	tam := 3
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(25)
	require.Equal(t, 25, iter.VerActual())
	require.Equal(t, 25, iter.Borrar())
	require.Equal(t, 2, iter.VerActual())
	iter2 := lista.Iterador()
	for i := 1; i >= tam; i++ {
		require.Equal(t, i, iter2.VerActual())
		iter2.Siguiente()
	}

}
func Test_Suma_Toda_La_lista_Interno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(6)
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		return true
	})
	require.Equal(t, 14, contador)

}

func Test_5_Primeros_Lista_Mas_Grande_Interno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)

	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += 1
		require.Equal(t, contador, v)
		return *contador_ptr < 5
	})
	require.Equal(t, 5, contador)
}

func Test_5_Primeros_Lista_Mas_Chica_Interno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += 1
		require.Equal(t, contador, v)
		return *contador_ptr < 5
	})
	require.Equal(t, 3, contador)
}

func Test_Sumar_Pares_Interno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	suma := 0
	suma_ptr := &suma
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			*suma_ptr += v
		}
		return true
	})
	require.Equal(t, 10, suma)
}

func Test_Lista_Vacia_Iterador_Interno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		return true
	})
	require.Equal(t, 0, contador)
}
