package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/GiovaniGitHub/multithreading/internal/entity"
)

// Get Address godoc
// @Summary      Get Address
// @Description  Get Address by Post Code
// @Tags         address
// @Accept       json
// @Produce      json
// @Param        cep     path       string    true    "Cep"
// @Success      200     {object}   entity.AnyAddress
// @Failure      400     string     "Bad Request"
// @Failure      404     string     "Not Found"
// @Failure      500     string     "Internal Server Error"
// @Router       /cep/{cep}	[get]
func GetAddress(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Path[len("/cep/"):]
	ch := make(chan entity.AnyAddress, 2)

	if !regexp.MustCompile(`^\d{5}-\d{3}$`).MatchString(cep) {
		cep = regexp.MustCompile(`[^\d]`).ReplaceAllString(cep, "")
		if len(cep) == 8 {
			cep = cep[:5] + "-" + cep[5:]
		}

	}
	go func() {
		resp, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
		if err != nil {
			fmt.Println("Erro ao fazer request para API 1:", err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro ao ler resposta da API 1:", err)
			return
		}

		var addressApiCep entity.AddressApiCep
		err = json.Unmarshal(body, &addressApiCep)
		if err != nil {
			fmt.Println("Erro ao fazer parse da resposta da API 1:", err)
			return
		}
		addressApiCep.Api = "https://cdn.apicep.com/file/apicep/"
		ch <- entity.AnyAddress(addressApiCep)
	}()

	go func() {
		resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Println("Erro ao fazer request para API 2:", err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro ao ler resposta da API 2:", err)
			return
		}

		var addressViaCep entity.AddressViaCep
		err = json.Unmarshal(body, &addressViaCep)
		if err != nil {
			fmt.Println("Erro ao fazer parse da resposta da API 2:", err)
			return
		}
		addressViaCep.Api = "http://viacep.com.br/ws/"
		ch <- entity.AnyAddress(addressViaCep)
	}()

	select {
	case res := <-ch:
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			log.Fatalf("Erro ao serializar os dados em JSON: %v", err)
		}
		log.Print(string(data))
		json.NewEncoder(w).Encode(res)

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout atingido. Nenhuma resposta recebida a tempo.")
		http.Error(w, "Timeout atingido. Nenhuma resposta recebida a tempo.", http.StatusRequestTimeout)
	}
}
