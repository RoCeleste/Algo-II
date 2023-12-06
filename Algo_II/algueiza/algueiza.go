package main

import (
	err "algueiza/diseno_vuelos/errores"
	"algueiza/diseno_vuelos/vuelos"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	COMANDO                    int    = 0
	VUELOS                     int    = 1
	CSV                        int    = 1
	NUM_VUELO                  int    = 1
	K                          int    = 1
	ORIGEN                     int    = 1
	FECHA_INICIAL              int    = 1
	MODO                       int    = 2
	DESTINO                    int    = 2
	FECHA_FINAL                int    = 2
	LEN_PARAMETROS_INFO_VUELO  int    = 2
	DESDE                      int    = 3
	LEN_PARAMETROS_BORRAR      int    = 3
	FECHA                      int    = 3
	HASTA                      int    = 4
	LEN_PARAMETROS_VER_TABLERO int    = 5
	CARACTER_INCLUYA_HASTA     string = "$"
)

func ingresar_vuelos(vuelos_general vuelos.Sistema_Vuelos, archivo string) error {
	f, errores := os.Open(archivo)
	if errores != nil {
		return errores
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		info_vuelo := strings.Split(scanner.Text(), ",")
		nuevo_vuelo := vuelos.CrearVuelo(info_vuelo)
		vuelos_general.AgregarVuelo(*nuevo_vuelo)

	}
	return nil
}

func errores_input_ver_tablero(input []string) bool {
	cant, _ := strconv.Atoi(input[K])
	if cant <= 0 {
		return true
	}
	if input[MODO] != "asc" && input[MODO] != "desc" {
		return true
	}
	if strings.Compare(input[DESDE], input[HASTA]) > 0 {
		return true
	}
	return false
}

func printear_ver_tablero(cant int, modo string, vuelos []string) {
	// si el modo es "desc" recorro el arreglo, de atras para adelante, k veces o hasta el largo del array (lo que se ocurra primero)
	// si el modo es "asc" recorro el arreglo k veces o el largo del array
	if modo == "desc" {
		contador := 0
		for i := len(vuelos) - 1; i >= 0 && contador < cant; i-- {
			contador += 1
			fmt.Fprintf(os.Stdout, "%s\n", vuelos[i])
		}
	} else {
		contador := 0
		for i := 0; i < len(vuelos) && contador < cant; i++ {
			contador += 1
			fmt.Fprintf(os.Stdout, "%s\n", vuelos[i])
		}
	}
}

func main() {
	central := vuelos.CrearRegistroVuelos()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inputComoArreglo := strings.Split(s.Text(), " ")

		if inputComoArreglo[COMANDO] == "agregar_archivo" {
			errores := ingresar_vuelos(central, inputComoArreglo[CSV])
			if errores != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err.ErrorCargarVuelos{}.Error())
				continue
			}
			fmt.Fprintf(os.Stdout, "%s\n", "OK")
		}

		if inputComoArreglo[COMANDO] == "ver_tablero" {
			if len(inputComoArreglo) != LEN_PARAMETROS_VER_TABLERO {
				fmt.Fprintf(os.Stderr, "%s\n", err.ErrorTablero{}.Error())
				continue
			}
			if errores_input_ver_tablero(inputComoArreglo) {
				fmt.Fprintf(os.Stdout, "%s\n", "OK")
				continue
			}
			k, _ := strconv.Atoi(inputComoArreglo[K])
			vuelos := central.VuelosEnRango(inputComoArreglo[DESDE], inputComoArreglo[HASTA]+CARACTER_INCLUYA_HASTA)
			printear_ver_tablero(k, inputComoArreglo[MODO], vuelos)
			fmt.Fprintf(os.Stdout, "%s\n", "OK")
		}

		if inputComoArreglo[COMANDO] == "info_vuelo" {
			if len(inputComoArreglo) != LEN_PARAMETROS_INFO_VUELO {
				fmt.Fprintf(os.Stderr, "%s\n", err.ErrorVuelo{}.Error())
				continue
			}
			cadena_arr, errs := central.ObtenerInformacionVuelo(inputComoArreglo[NUM_VUELO])
			if errs != nil {
				fmt.Fprintf(os.Stderr, "%s\n", errs)
				continue
			}
			fmt.Fprintf(os.Stdout, "%s\n", cadena_arr)
			fmt.Fprintf(os.Stdout, "%s\n", "OK")

		}

		if inputComoArreglo[COMANDO] == "prioridad_vuelos" {
			k, _ := strconv.Atoi(inputComoArreglo[K])
			if k < 0 {
				fmt.Fprintf(os.Stderr, "%s\n", err.ErrorPrioridadVuelo{}.Error())
				continue
			}
			lista_ordenada_prioridad := central.ObtenerKMayorPrioridad(k)
			for _, vuelo := range lista_ordenada_prioridad {
				fmt.Fprintf(os.Stdout, "%s\n", vuelo)
			}
			fmt.Fprintf(os.Stdout, "%s\n", "OK")
		}

		if inputComoArreglo[COMANDO] == "siguiente_vuelo" {
			vuelo, errs := central.ObtenerSiguienteVuelo(inputComoArreglo[ORIGEN], inputComoArreglo[DESTINO], inputComoArreglo[FECHA])
			if errs != nil {
				fmt.Fprintf(os.Stdout, "%s\n", errs)
				fmt.Fprintf(os.Stdout, "%s\n", "OK")
				continue
			}
			fmt.Fprintf(os.Stdout, "%s\n", vuelo)
			fmt.Fprintf(os.Stdout, "%s\n", "OK")
		}

		if inputComoArreglo[COMANDO] == "borrar" {
			if len(inputComoArreglo) != LEN_PARAMETROS_BORRAR {
				fmt.Fprintf(os.Stderr, "%s\n", err.ErrorBorrarVuelo{}.Error())
				continue
			}
			vuelos_eliminados := central.BorrarVuelos(inputComoArreglo[FECHA_INICIAL], inputComoArreglo[FECHA_FINAL]+CARACTER_INCLUYA_HASTA)
			for _, vuelo_borrado := range vuelos_eliminados {
				fmt.Fprintf(os.Stdout, "%s\n", vuelo_borrado)
			}
			fmt.Fprintf(os.Stdout, "%s\n", "OK")
		}
	}
}