package votos

import (
	err "rerepolez/diseno_alumnos/errores"
	"rerepolez/pila"
)

type votanteImplementacion struct {
	dni           int
	voto          Voto
	se_presento   bool
	pila_de_votos pila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.se_presento = false
	// esto crea un voto en blanco, que es el default
	v := pila.CrearPilaDinamica[Voto]()
	v.Apilar(votoInicial())
	votante.pila_de_votos = v
	return votante
}

func votoInicial() Voto {
	voto := new(Voto)
	voto.VotoPorTipo = [CANT_VOTACION]int{0, 0, 0}
	voto.Impugnado = false
	return *voto
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.se_presento {
		error_a_devolver := new(err.ErrorVotanteFraudulento)
		error_a_devolver.Dni = votante.LeerDNI()
		return error_a_devolver
	}

	if alternativa == LISTA_IMPUGNA {
		votante.voto.Impugnado = true

	} else {
		votante.voto.VotoPorTipo[tipo] = alternativa
	}
	votante.pila_de_votos.Apilar(votante.voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.se_presento {
		error_a_devolver := new(err.ErrorVotanteFraudulento)
		error_a_devolver.Dni = votante.LeerDNI()
		return error_a_devolver
	}

	voto_actual := votante.pila_de_votos.Desapilar() // siempre va a haber un voto, si la pila queda vacia, es porque voto_actual era en blanco
	if votante.pila_de_votos.EstaVacia() {
		votante.pila_de_votos.Apilar(voto_actual)
		error_a_devolver := new(err.ErrorNoHayVotosAnteriores)
		return error_a_devolver
	} else {
		votante.voto = votante.pila_de_votos.VerTope() // la idea es que votante_voto sea siempre el tope de pila_de_votos, y el que se devuelva en fin-votar
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.se_presento {
		error_a_devolver := new(err.ErrorVotanteFraudulento)
		error_a_devolver.Dni = votante.LeerDNI()
		return votante.pila_de_votos.VerTope(), error_a_devolver
	}

	votante.se_presento = true
	return votante.voto, nil
}

func (votante votanteImplementacion) SePresento() bool {
	return votante.se_presento
}