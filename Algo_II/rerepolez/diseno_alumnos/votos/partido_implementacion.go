package votos

import (
	"fmt"
)

type partidoImplementacion struct {
	nombre          string
	candidatosLista [CANT_VOTACION]string
	votos           [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	nuevo_partido := new(partidoImplementacion)
	nuevo_partido.nombre = nombre
	nuevo_partido.candidatosLista = candidatos
	return nuevo_partido
}

func CrearVotosEnBlanco() Partido {
	partido_en_blanco := new(partidoEnBlanco)
	return partido_en_blanco
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo] += 1
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) (string, int) {
	return fmt.Sprintf("%s - %s: ", partido.nombre, partido.candidatosLista[tipo]), partido.votos[tipo]
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo] += 1
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) (string, int) {
	return "Votos en Blanco: ", blanco.votos[tipo]
}