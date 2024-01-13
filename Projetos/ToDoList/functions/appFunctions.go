package functions

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"todolist/data"
	"todolist/entity"
)

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func AppToDoList(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearTerminal()
		fmt.Printf("1 - Inserir tarefa\n2 - Mostrar tarefas\n3 - Alterar tarefas\n4 - Apagar tarefas\n5 - Exit\n\nresposta: ")

		scanner.Scan()
		clearTerminal()

		value := scanner.Text()

		switch value {
		case "1":
			menuInsert(db, scanner)
		case "2":
			menuRead(db, scanner)
		case "3":
			menuUpdate(db, scanner)
		case "4":
			menuDelete(db, scanner)
		case "5":
			return
		default:
			fmt.Println("Erro ao digitar")
		}
	}

}

func menuCreateTask(scanner *bufio.Scanner) *entity.Task {
	clearTerminal()
	fmt.Printf("Informe o conteudo da tarefa: ")
	scanner.Scan()
	conteudo := scanner.Text()

	fmt.Printf("Informe a data da tarefa seguindo um padrão (DD/MM/AAAA): ")
	scanner.Scan()

	data_value := scanner.Text()

	data, err := entity.NewDateFormat(data_value)

	if err != nil {
		log.Fatalf("err na criação de uma data na função mencreatetask %v", err)
	}

	return entity.NewTask(conteudo, *data)
}

func menuInsert(db *sql.DB, scanner *bufio.Scanner) {
	clearTerminal()
	fmt.Printf("1 - Cadastrar nova tarefa\n2 - Voltar para o menu principal\n\nresposta: ")
	scanner.Scan()

	value := scanner.Text()

menu:
	for {
		switch value {
		case "1":
			data.Insert(db, *menuCreateTask(scanner))
		case "2":
			break menu
		default:
			fmt.Printf("Digite uma resposta valida")
		}
	}
}

func menuUpdate(db *sql.DB, scanner *bufio.Scanner) {
	clearTerminal()
	fmt.Print("1 - Atualizar tarefa\n2 - Concluir Tarefa\n3 - Voltar para o menu principal\n\nresposta: ")

	scanner.Scan()

	value := scanner.Text()

menu:
	for {
		switch value {
		case "1":
			fmt.Printf("Informe o conteudo: ")
			scanner.Scan()

			conteudo := scanner.Text()
			ID := data.GetId(db, conteudo)

			fmt.Printf("Informe a data nesse padrão (DD/MM/AAAA): ")
			scanner.Scan()

			datanova := scanner.Text()

			data.UpdateByID(db, map[string]string{"conteudo": conteudo, "data": datanova}, ID)
		case "2":
			fmt.Printf("Informe o conteudo: ")
			scanner.Scan()

			conteudo := scanner.Text()

			data.ConcludeTask(db, conteudo)
		case "3":
			break menu
		default:
			fmt.Printf("Digite uma resposta valida")
		}
	}

}

func menuRead(db *sql.DB, scanner *bufio.Scanner) {
	tarefas := data.SelectAllTasks(db)

	clearTerminal()

	for _, value := range tarefas {
		fmt.Println("======================================")
		fmt.Printf("ID: %v\t%v\n", data.GetId(db, entity.ReadTaskConteudo(&value)), value.Show())
	}
	fmt.Println("======================================")

	scanner.Scan()
}

func menuDelete(db *sql.DB, scanner *bufio.Scanner) {
	clearTerminal()

	fmt.Print("1 - Excluir tarefa\n2 - Voltar para o menu principal\n\nresposta: ")

	scanner.Scan()

	decision := scanner.Text()

menu:
	for {
		clearTerminal()
		switch decision {
		case "1":
			fmt.Printf("Digite o conteudo da tarefa que deseja deletar: ")
			scanner.Scan()
			value := scanner.Text()
			data.DeleteByConteudo(db, value)
			break menu
		case "2":
			return
		default:
			fmt.Printf("Digite uma resposta valida")
		}
	}
}
