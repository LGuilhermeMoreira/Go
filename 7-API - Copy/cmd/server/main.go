package main

import "API/fundamentos/configs"

func main() {
	// Carrega a configuração usando o caminho do arquivo .env
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(config.DBDriver)

}
