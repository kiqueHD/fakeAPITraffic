package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type Result struct {
	Denominacion string  `json:"denominacion"`
	Estado       int     `json:"estado"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
}

type Response struct {
	Results []Result `json:"results"`
}

func randomDenominacion() string {
	calles := []string{
		"Calle", "Avenida", "Plaza", "Paseo", "Carretera",
		"Camino", "Ronda", "Travesía", "Bulevar", "Callejón",
		"Vía", "Pasaje", "Glorieta", "Sendero", "Autopista",
	}
	//nombes para las calles graciosos, un poco sin sentido
	nombres := []string{
		"Pepito", "Juanito", "Manolito Gafotas", "Paquito", "Jaimito",
		"Don Limpio", "de tu Suegra", "Chiquito de la calzada", "Torrente", "Mortadelo",
		"Filemón", "Doraemon", "Goku", "Patricio", "Gumball",
		"Ubuntu", "PHP", "Bash", "Spring boot", "Java", "Laravel", "golang",
		"Polimorfismo", "Herencia", "Encapsulamiento", "Abstracción",
		"Inner join", "Right join", "Left join",
	}

	// Generar una dirección aleatoria
	// Calle + nombre + número aleatorio entre 1 y 100
	direccion := fmt.Sprintf("%s %s %d",
		calles[rand.Intn(len(calles))],
		nombres[rand.Intn(len(nombres))],
		rand.Intn(100)+1,
	)

	return direccion
}

func randomLocationInSpain() (float64, float64) {
	minLat, maxLat := 36.0, 43.8
	minLon, maxLon := -9.3, 3.0

	lat := minLat + rand.Float64()*(maxLat-minLat)
	lon := minLon + rand.Float64()*(maxLon-minLon)

	return lon, lat
}

func generateRandomResults(count int) []Result {
	results := make([]Result, count)
	for i := 0; i < count; i++ {
		lon, lat := randomLocationInSpain()
		results[i] = Result{
			Denominacion: randomDenominacion(),
			Estado:       rand.Intn(10),
			Lon:          lon,
			Lat:          lat,
		}
	}
	return results
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Configurar CORS
	//w.Header().Set("Access-Control-Allow-Origin", "https://yourdomain.com")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	limit := 20
	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		// Validar el parámetro de límite y hacer que que se convierta a un número entero
		//condiciones que err sea nil (conversion a entero)y que el número sea mayor que 0
		if urlLimit, err := strconv.Atoi(limitParam); err == nil && urlLimit > 0 {
			//no mas de 10000 resultados
			if urlLimit > 100000 {
				urlLimit = 100000
			}
			limit = urlLimit
		}
	}

	results := generateRandomResults(limit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Results: results})
}

func main() {
	//endpoint de la API
	http.HandleFunc("/fakeTrafficAPI", apiHandler)
	//avisos por consola
	fmt.Println("Servidor iniciado en http://localhost:8081/fakeTrafficAPI")
	fmt.Println("Prueba con: http://localhost:8081/fakeTrafficAPI?limit=100")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
