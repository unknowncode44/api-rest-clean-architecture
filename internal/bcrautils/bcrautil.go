package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/unknowncode44/api-rest-clean-architecture/config"
)

// definimos un tipo que contendra la fecha y el tipo de cambio
type (
	ExchangeObj struct {
		Date     string
		Exchange float64
	}
	// definimos un tipo que contendra el tipo de respuesta que obtendremos de la api del BCRA
	ResponseObj struct {
		Data  string  `json:"d"`
		Value float64 `json:"v"`
	}
)

// variables de uso en la funcion GetExchange
var (
	once             sync.Once
	exchangeInstance *ExchangeObj
)

// La funcion GetExchange se encargara de:
//
//	-Llamar a la API del BCRA
//	-Iterar en la respuesta y encontrar la fecha de la ultima cotizacion
//	-Devolver un objeto con la cotizacion y la fecha
func GetExchange(conf *config.Config) *ExchangeObj {
	once.Do(func() {
		now := time.Now() // definimos la fecha actual

		var date string

		date = now.Format("2006-01-02") // asignamos la fecha a la variable date

		// si es sabado o domingo tomamos la fecha del ultimo viernes
		if now.Weekday() == 6 {
			yesterday := now.AddDate(0, 0, -1)
			date = yesterday.Format("2006-01-02")
		}

		if now.Weekday() == 0 {
			friday := now.AddDate(0, 0, -2)
			date = friday.Format("2006-01-02")
		}

		// si es lunes antes de las 15hs es probable que todavia no este registrado el TC
		// usamos fecha y TC del viernes
		if now.Weekday() == 1 {
			if now.Hour() <= 15 {
				friday := now.AddDate(0, 0, -3)
				date = friday.Format("2006-01-02")
			}
		}

		// preparamos la llamada usando la funcion getData
		obj, err := getData(conf.Bcra.Url, conf.Bcra.Token, date)

		if err != nil {
			panic(err)
		}

		// la variable exchObj es un objeto con la fecha y el tipo de cambio
		var exchObj ExchangeObj
		exchObj.Date = obj.Data
		exchObj.Exchange = obj.Value

		// asingnamos la variable que contendra la instancia con el tipo de cambio
		exchangeInstance = &exchObj

		// imprimimos un mensaje con el tipo de cambio actual
		fmt.Printf("Fecha de tipo de cambio %v\n", obj.Data)
		fmt.Printf("Tipo de cambio AR$ %v = 1 USD \n", obj.Value)

	})
	return exchangeInstance

}

// la funcion getData se encarga de realizar una peticion GET a la api del BCRA
func getData(url string, token string, date string) (ResponseObj, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	// pasamos el token almacenado en nuestra configuracion
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	// realizamos la llamada
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Making request:", err)
		panic(err)
	}
	defer resp.Body.Close()

	body := resp.Body

	// parseamos la respuesta a json y obtenemos el array con cotizaciones
	decoder := json.NewDecoder(body)
	var responseArray []ResponseObj
	err = decoder.Decode(&responseArray)
	if err != nil {
		panic(err)
	}

	// utilizamos una busqueda binaria para encontrar la fecha de la ultima cotizacion vigente
	index, found := binarySearch(responseArray, date)

	if found {
		return responseArray[index], nil
	} else {
		// si no lo encontramos devolvemos el ultimo TC registrado en el array
		return responseArray[len(responseArray)-1], nil
	}

}

func less(items []ResponseObj, i, j int) bool {
	return items[i].Data < items[j].Data
}

// funcion para busqueda binaria
func binarySearch(items []ResponseObj, targetDate string) (int, bool) {
	sort.Slice(items, func(i, j int) bool { return less(items, i, j) })
	low := 0
	high := len(items) - 1
	for low <= high {
		mid := (low + high) / 2
		if items[mid].Data == targetDate {
			return mid, true
		} else if items[mid].Data < targetDate {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1, false
}
