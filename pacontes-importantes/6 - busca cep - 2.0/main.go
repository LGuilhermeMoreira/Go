package main

// import de cada biblioteca
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// struct criada para "armazenar o cep"
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	// w -> response
	// r -> request

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//ex: localhost:8000/cep=----------
	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	cep, err := BuscaCEP(cepParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// faz com que retorne do formato JSON
	w.Header().Set("Content-Type", "application/json")

	// faz com que retorne com status OK
	w.WriteHeader(http.StatusOK)

	result, err := json.Marshal(cep)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// retorno da função associado ao "/"
	w.Write(result)

	// transformando a struct em JSON usando o encode e retornando
	//json.NewEncoder(w).Encode(cep)
}

func BuscaCEP(cep string) (*ViaCEP, error) {
	request, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return nil, err
	}

	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return nil, err
	}

	var C ViaCEP

	err = json.Unmarshal(body, &C)

	if err != nil {
		return nil, err
	}

	return &C, nil
}

func main() {
	http.HandleFunc("/", BuscaCEPHandler) // criando um end point

	// pode usar clousures
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Ola mundo"))
	// })
	http.ListenAndServe(":8000", nil) // abrindo || rodando um http servewr
}
