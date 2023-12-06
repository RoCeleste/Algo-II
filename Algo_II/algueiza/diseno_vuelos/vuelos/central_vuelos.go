package vuelos

// modela una base de datos de vuelos
type Sistema_Vuelos interface {

	//se agrega un vuelo a la base de datos
	AgregarVuelo(vuelo vuelo)

	//Devuelve un arreglo con todos los vuelos entre el desde y el hasta (inclusive) de la forma <fecha> - <n vuelo>
	VuelosEnRango(desde string, hasta string) []string

	//Devuelve toda la informacion del vuelo cuyo código de vuelo coincida con el que fue pasado por parámetro
	//si existe el vuelo devuelve como error a nil
	//caso contrario devuelve string vacio y el error correspondiente
	ObtenerInformacionVuelo(numero_vuelo string) (string, error)

	//Obtengo los K vuelos con mayor prioridad que hayan sido cargados en el sistema
	ObtenerKMayorPrioridad(K int) []string

	//Obtengo el siguiente vuelo de dada fecha que cumpla las misma condiciones de origen, destino
	//si existe devuelve toda informacion del vuelo y nil
	//caso contrario devuelve un string vacio y el error correspondiente
	ObtenerSiguienteVuelo(origen string, destino string, fecha string) (string, error)

	//Elimina del sistema todos los vuelos cuya fecha de despegue sea igual o mayor a desde e igual o menor a hasta
	//Devuelve dichos vuelos de la forma <fecha> - <n vuelo>
	BorrarVuelos(desde string, hasta string) []string
}