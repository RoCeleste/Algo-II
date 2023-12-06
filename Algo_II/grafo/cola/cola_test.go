package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "cola"
	"testing"
)

func Test_Cola_Vacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
}

func Test_Volumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	for i := 0.0; i < 1000000.0; i += 1.0 {
		cola.Encolar(i)
	}

	for i := 0.0; i < 1000000.0; i += 1.0 {
		cola.Desencolar()
	}

	require.True(t, cola.EstaVacia())
}

func Test_Invalidez_De_Primero_Y_Esta_Vacia_En_Cola_Recien_Creada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}

func Test_Mantiene_Invariante(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(4)
	cola.Encolar(532)
	require.Equal(t, 4, cola.VerPrimero())
	require.Equal(t, 4, cola.Desencolar())
	require.Equal(t, 532, cola.Desencolar())
	cola.Encolar(32)
	require.Equal(t, 32, cola.VerPrimero())
	require.False(t, cola.EstaVacia())
	require.Equal(t, 32, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func Test_Esta_Vacia_En_Cola_Recien_Creada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	require.True(t, cola.EstaVacia())
}

func Test_Invalidez_De_Primero_Y_Esta_Vacia_En_Cola_Apilada_Y_Desapilada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	for i := 0; i < 1000; i += 1 {
		cola.Encolar("elefante")
	}
	for i := 0; i < 1000; i += 1 {
		cola.Desencolar()
	}
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}
func Test_se_comporta_Igual_Recien_Creada_Y_Apilada_Y_Desapilada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Lana")
	cola.Encolar("Rana")
	cola.Encolar("Ana")

	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia())

	cola_2 := TDACola.CrearColaEnlazada[string]()

	require.True(t, cola_2.EstaVacia())
}
