package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	lenCreatePath = 11
	lenDeletePath = 5
	lenFiles      = 2
)

var (
	createpath = [lenCreatePath]string{
		"cmd",
		"internal",
		"bin",
		"pkg",
		"config",
		"pkg/etnity",
		"internal/infra",
		"internal/handler",
		"internal/utils",
		"internal/infra/database",
		"internal/infra/model",
	}
	deletePath = [lenDeletePath]string{
		"cmd",
		"internal",
		"bin",
		"pkg",
		"config",
	}
	files = [lenFiles]string{
		"go.mod",
		".air.toml",
	}
)

func generateDir() {
	for _, v := range createpath {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			if err := os.MkdirAll(v, 0755); err != nil {
				fmt.Printf("Erro ao criar o diretório %s: %v\n", v, err)
			} else {
				fmt.Printf("Diretório %s criado com sucesso.\n", v)
			}
		} else if err != nil {
			fmt.Printf("Erro ao verificar o diretório %s: %v\n", v, err)
		} else {
			fmt.Printf("Diretório %s já existe.\n", v)
		}
	}
}
func generateGoMod(mod string) {
	if err := exec.Command("go", "mod", "init", mod).Run(); err != nil {
		log.Println("Erro ao gerar go.mod", err)
	} 
	fmt.Println("go.mod criado com sucesso")
	
}

func generateAir() {
	if err := exec.Command("air", "init").Run(); err != nil {
		log.Println("Erro ao gerar .air.toml", err)
	} 
	fmt.Println(".air.tolm criado com sucesso")

}

func deleteAll() {
	if err := exec.Command("rm", "-rf", deletePath[0], deletePath[1], deletePath[2], deletePath[3], deletePath[4]).Run(); err != nil {
		log.Printf("Erro ao apagar %v\n", deletePath)
	} else {
		fmt.Printf("%v apagado com sucesso\n", deletePath)
	}

	if err := exec.Command("rm", files[0]).Run(); err != nil {
		log.Printf("Erro ao apagar %v\n", files[0])
	} else {
		fmt.Printf("%v apagado com sucesso\n", files[0])
	}

	if err := exec.Command("rm", files[1]).Run(); err != nil {
		log.Printf("Erro ao apagar %v\n", files[1])
	} else {
		fmt.Printf("%v apagado com sucesso\n", files[1])
	}

}

func main() {
	mod := flag.String("goMod", "", "flag destinada a gerar um go.mod")
	air := flag.Bool("air", false, "flag destinada a gerar um .air.toml")
	delete := flag.Bool("delete", false, "flag destinada a apagar os diretorios e arquvios criados")
	flag.Parse()

	//delete
	if *delete {
		deleteAll()
		return
	}

	generateDir()

	//generate go mod
	if *mod != "" {
		generateGoMod(*mod)
	}
	// generate air
	if *air {
		generateAir()
	}
}
