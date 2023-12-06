package pila_test

import (
	"github.com/stretchr/testify/require"
	"reflect"
	TDAPila "pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

}

func TestPilaDeCadenas(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("cadena")
	require.Equal(t, "cadena", pila.VerTope())
}
func TestApilarYDesapilarVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 1000000; i += 1 {
		pila.Apilar(i)
	}
	for i := 0; i < 1000000; i += 1 {
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}
func TestPilaConFloats(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	for i := 0.5; i < 10; i += 1 {
		pila.Apilar(i)
	}
	for pila.EstaVacia() == false {
		tope := pila.Desapilar()
		require.Equal(t, "float64", reflect.TypeOf(tope).String())
	}
}
func TestMetodosInvalidosEnPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}
func TestApilarYDesapilarHastaEstarVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 100000; i += 1 {
		pila.Apilar(i)
	}
	for i := 0; i < 100000; i += 1 {
		pila.Desapilar()
	}
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}
func TestseComportanIgual(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 1000; i += 1 {
		pila.Apilar(i)
	}
	for i := 0; i < 1000; i += 1 {
		pila.Desapilar()
	}
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())

	pila_2 := TDAPila.CrearPilaDinamica[int]()

	require.Panics(t, func() { pila_2.Desapilar() })
	require.Panics(t, func() { pila_2.VerTope() })
	require.True(t, pila_2.EstaVacia())
}
func TestseMantieneInvariante(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10; i += 1 {
		pila.Apilar(i)
	}
	for i := 0; i < 10; i += 1 {
		require.Equal(t, pila.Desapilar(), 10-i)
	}
}
