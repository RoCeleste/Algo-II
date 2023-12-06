package main

import ( //todos los paquetes

	"bufio"
	"fmt"
	"math/rand"
	"os"
	err "rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"strconv"
	"strings"
	"rerepolez/cola"
)

const (
	PRESIDENTE      votos.TipoVoto = 0
	GOBERNADOR      votos.TipoVoto = 1
	INTENDENTE      votos.TipoVoto = 2
	COMANDO         int            = 0
	PARTIDO         int            = 0
	PADRON          int            = 1
	NUMERO_DNI      int            = 1
	TIPO_VOTO       int            = 1
	ALTERNATIVA     int            = 2
	CANT_PARAMETROS int            = 2
)

func quicksort(arr []int) { //sacado de internet, mas rapido que counting sort por que el rango del padron es muy alto
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr)
	quicksort(arr[:pivotIndex])
	quicksort(arr[pivotIndex+1:])
}

func partition(arr []int) int {
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]

	arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	return i
}

func output_final(partidos []votos.Partido, votos_impugnados int) {
	fmt.Fprintf(os.Stdout, "%s:\n", "Presidente")

	for _, partido := range partidos {
		str1, cant_votos := partido.ObtenerResultado(votos.PRESIDENTE)
		str2 := votos_en_str(cant_votos)
		fmt.Fprintf(os.Stdout, "%s%s\n", str1, str2)
	}
	fmt.Fprintf(os.Stdout, "\n")
	fmt.Fprintf(os.Stdout, "%s:\n", "Gobernador")

	for _, partido := range partidos {
		str1, cant_votos := partido.ObtenerResultado(votos.GOBERNADOR)
		str2 := votos_en_str(cant_votos)
		fmt.Fprintf(os.Stdout, "%s%s\n", str1, str2)
	}
	fmt.Fprintf(os.Stdout, "\n")
	fmt.Fprintf(os.Stdout, "%s:\n", "Intendente")

	for _, partido := range partidos {
		str1, cant_votos := partido.ObtenerResultado(votos.INTENDENTE)
		str2 := votos_en_str(cant_votos)
		fmt.Fprintf(os.Stdout, "%s%s\n", str1, str2)
	}
	fmt.Fprintf(os.Stdout, "\n")
	votos_imp_str := votos_en_str(votos_impugnados)
	fmt.Fprintf(os.Stdout, "Votos Impugnados: %s\n", votos_imp_str)
}

func votos_en_str(cant_voto int) string {
	if cant_voto == 1 {
		return "1 voto"
	}
	return fmt.Sprintf("%d votos", cant_voto)
}

func busquedaBinaria(numbers []votos.Votante, inicio, fin, numberToFind int) int {
	if inicio > fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if numbers[medio].LeerDNI() == numberToFind {
		return medio
	}
	if numbers[medio].LeerDNI() < numberToFind {
		return busquedaBinaria(numbers, medio+1, fin, numberToFind)
	} else {
		return busquedaBinaria(numbers, inicio, medio-1, numberToFind)
	}
}

func cargar_padron(archivo string) ([]votos.Votante, error) {
	var padron_final []votos.Votante
	var padron_int []int
	f, errores := os.Open(archivo)
	if errores != nil {
		return padron_final, errores
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		dni, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return padron_final, err
		}
		padron_int = append(padron_int, dni)
	}
	quicksort(padron_int)
	for _, dni := range padron_int {
		padron_final = append(padron_final, votos.CrearVotante(dni))
	}
	return padron_final, nil
}

func cargar_partidos(archivo string) ([]votos.Partido, int, error) {
	var todas_las_listas []votos.Partido
	var total_partidos int
	f, errores := os.Open(archivo)
	if errores != nil {
		return todas_las_listas, total_partidos, errores
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		linea := scanner.Text()
		fila := strings.Split(linea, ",")
		var candidatos [votos.CANT_VOTACION]string
		for i := range fila[1:] {
			candidatos[i] = fila[i+1]
		}
		if total_partidos == 0 {
			todas_las_listas = append(todas_las_listas, votos.CrearVotosEnBlanco())
		}
		partido := votos.CrearPartido(fila[0], candidatos)
		todas_las_listas = append(todas_las_listas, partido)
		total_partidos += 1
	}
	return todas_las_listas, total_partidos, nil
}

func encolarVotante(cola cola.Cola[votos.Votante], posible_dni string, padron []votos.Votante) error {
	posible_dni_int, errores := strconv.Atoi(posible_dni)
	if errores != nil {
		return err.DNIError{}
	}
	posicionDelVotanteEnPadron := busquedaBinaria(padron, 0, len(padron)-1, posible_dni_int)
	if posicionDelVotanteEnPadron == -1 {
		return err.DNIFueraPadron{}
	}
	cola.Encolar(padron[posicionDelVotanteEnPadron])
	return nil
}

func validar_tipo_voto(comando string) (votos.TipoVoto, error) {
	switch comando {
	case "Presidente":
		return PRESIDENTE, nil
	case "Gobernador":
		return GOBERNADOR, nil
	case "Intendente":
		return INTENDENTE, nil
	default:
		return -1, err.ErrorTipoVoto{}
	}
}

func validar_alternativa(alternativa string, cant_partidos int) (int, error) {
	alternativa_int, errores := strconv.Atoi(alternativa)
	if errores != nil || alternativa_int > cant_partidos || alternativa_int < 0 {
		return -1, err.ErrorAlternativaInvalida{}
	}
	return alternativa_int, nil
}

func verEstado(cola cola.Cola[votos.Votante], hay_votante *bool, votante_actual *votos.Votante) error {
	if cola.EstaVacia() && !*hay_votante {
		return err.FilaVacia{}
	}
	if !*hay_votante {
		*votante_actual = cola.Desencolar()
		*hay_votante = true
	}
	return nil
}

func ingresar(cola cola.Cola[votos.Votante], numero_dni string, padron []votos.Votante) {
	errores := encolarVotante(cola, numero_dni, padron)
	if errores != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errores)
		return
	}
	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}

func votar(hay_votante *bool, votante_actual votos.Votante, voto string, alternativa string, cant_partidos int) {
	tipo_voto, error_voto := validar_tipo_voto(voto)
	if error_voto != nil {
		fmt.Fprintf(os.Stdout, "%s\n", error_voto)
		return
	}
	//tengo que validar que la alternativa sea un entero
	pos_lista, error_alternativa := validar_alternativa(alternativa, cant_partidos)
	if error_alternativa != nil {
		fmt.Fprintf(os.Stdout, "%s\n", error_alternativa)
		return
	}
	err := votante_actual.Votar(tipo_voto, pos_lista)
	if err != nil { //caso dni fraudulento
		fmt.Fprintf(os.Stdout, "%s\n", err)
		votante_actual = nil
		*hay_votante = false
		return
	}
	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}

func deshacer(hay_votante *bool, votante_actual votos.Votante) {
	errores := votante_actual.Deshacer()
	if votante_actual.SePresento() {
		votante_actual = nil
		*hay_votante = false
		fmt.Fprintf(os.Stdout, "%s\n", errores)
		return
	}
	if errores != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errores)
		return
	}
	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}

func fin_votar(hay_votante *bool, votante_actual votos.Votante, partidos []votos.Partido, impugnados *int) {
	voto_final, errores := votante_actual.FinVoto()
	if errores != nil { //caso votante fraudulento
		fmt.Fprintf(os.Stdout, "%s\n", errores)
		votante_actual = nil
		*hay_votante = false
		return
	}
	es_impugnado := voto_final.Impugnado
	if !es_impugnado {
		votos := voto_final.VotoPorTipo
		partidos[votos[PRESIDENTE]].VotadoPara(PRESIDENTE)
		partidos[votos[GOBERNADOR]].VotadoPara(GOBERNADOR)
		partidos[votos[INTENDENTE]].VotadoPara(INTENDENTE)
	} else {
		*impugnados += 1
	}
	*hay_votante = false
	votante_actual = nil
	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}

//-------------------------------func principal----------------------------------

func main() {
	cola := cola.CrearColaEnlazada[votos.Votante]()
	params := os.Args[1:]
	votos_impugnados := 0
	var hay_votante bool
	var votante_actual votos.Votante

	if len(params) < CANT_PARAMETROS {
		fmt.Fprintf(os.Stdout, "%s\n", err.ErrorParametros{}.Error())
		return
	}

	if len(params) > CANT_PARAMETROS {
		fmt.Fprintf(os.Stdout, "%s\n", err.ErrorLeerArchivo{}.Error())
		return
	}
	padron, errores := cargar_padron(params[PADRON])
	if errores != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.ErrorLeerArchivo{}.Error())
		return
	}
	partidos, cant_partidos, errores := cargar_partidos(params[PARTIDO])
	if errores != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.ErrorLeerArchivo{}.Error())
		return
	}

	//--------------------------------parte del input--------------------------------------------
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inputComoArreglo := strings.Split(s.Text(), " ")

		if inputComoArreglo[COMANDO] == "ingresar" {
			ingresar(cola, inputComoArreglo[NUMERO_DNI], padron)
		}

		if inputComoArreglo[COMANDO] == "votar" {
			errs := verEstado(cola, &hay_votante, &votante_actual)
			if errs != nil {
				fmt.Fprintf(os.Stdout, "%s\n", errs)
				continue
			}
			votar(&hay_votante, votante_actual, inputComoArreglo[TIPO_VOTO], inputComoArreglo[ALTERNATIVA], cant_partidos)
		}

		if inputComoArreglo[COMANDO] == "deshacer" {
			errs := verEstado(cola, &hay_votante, &votante_actual)
			if errs != nil {
				fmt.Fprintf(os.Stdout, "%s\n", errs)
				continue
			}
			deshacer(&hay_votante, votante_actual)
		}

		if s.Text() == "fin-votar" {
			errs := verEstado(cola, &hay_votante, &votante_actual)
			if errs != nil {
				fmt.Fprintf(os.Stdout, "%s\n", errs)
				continue
			}
			fin_votar(&hay_votante, votante_actual, partidos, &votos_impugnados)
		}
	}
	// -------------------------------- parte del output--------------------------------------
	if !cola.EstaVacia() {
		fmt.Fprintf(os.Stdout, "%s\n", err.ErrorCiudadanosSinVotar{}.Error())
	}
	output_final(partidos, votos_impugnados)
}