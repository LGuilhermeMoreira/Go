package data

import (
	"database/sql"
	"fmt"
	"log"
	"todolist/entity"
)

func Insert(db *sql.DB, tarefas entity.Task) {
	sttm, err := db.Prepare("INSERT INTO tarefas(conteudo,data_termino,status_conclusao) VALUES (?,?,?);")

	if err != nil {
		log.Fatalf("Error na função inserir <db.Prepare> %v", err)
	}

	defer sttm.Close()

	_, err = sttm.Exec(
		entity.ReadTaskConteudo(&tarefas),
		entity.ReadTaskDate(&tarefas),
		entity.ReadTaskStatus(&tarefas),
	)

	if err != nil {
		log.Fatalf("Error na função inserir <sttm.Exec> %v", err)

	}
}

func DeleteByConteudo(db *sql.DB, conteudo string) {
	sttm, err := db.Prepare("DELETE FROM tarefas where conteudo = (?)")

	if err != nil {
		log.Fatalf("Error na função deletebyconteudo <db.Prepare> %v", err)
	}

	defer sttm.Close()

	_, err = sttm.Exec(conteudo)

	if err != nil {
		log.Fatalf("Error na função deletebyconteudo <db.Exec> %v", err)
	}
}

func DeleteByID(db *sql.DB, id int) {
	sttm, err := db.Prepare("DELETE FROM tarefas WHERE ID = (?)")

	if err != nil {
		log.Fatalf("Error na função deletebyid <db.Prepare> %v", err)
	}

	defer sttm.Close()

	_, err = sttm.Exec(id)

	if err != nil {
		log.Fatalf("Error na função deletebyid <db.Exec> %v", err)
	}

}

// criar as variações de update (conteudo,data)

func UpdateByID(db *sql.DB, novo map[string]string, id int) {
	sttm, err := db.Prepare("UPDATE tarefas SET conteudo = ?,data_termino = ? WHERE id = ?")

	if err != nil {
		log.Fatalf("Error na função updatebyid <db.Prepare> %v", err)
	}

	defer sttm.Close()

	_, err = sttm.Exec(novo["conteudo"], novo["data"], id)

	if err != nil {
		log.Fatalf("Error na função updatebyid <db.Exec> %v", err)
	}
}

// pegando o id relacionado a um conteudo
func GetId(db *sql.DB, conteudo string) int {
	sttm, err := db.Prepare("SELECT id FROM tarefas WHERE conteudo = ?")

	if err != nil {
		log.Fatalf("Error na função getid <db.Prepare> %v", err)
	}

	defer sttm.Close()

	var ID int

	err = sttm.QueryRow(conteudo).Scan(&ID)

	if err != nil {
		log.Fatalf("Error na função geid <sttm.queryroww> %v", err)
	}

	return ID
}

// criar a função concludeTask
func ConcludeTask(db *sql.DB, conteudo string) {
	sttm, err := db.Prepare("UPDATE tarefas SET status_conclusao = ? Where conteudo = ?")

	if err != nil {
		log.Fatalf("Error na função inserir <db.Prepare> %v", err)
	}

	defer sttm.Close()

	_, err = sttm.Exec(true, conteudo)

	if err != nil {
		log.Fatalf("Error na função inserir <db.Exec> %v", err)
	}
}

// select all

// função auxiliar

func stringTodDate(data string) *entity.Date {
	var dia, mes, ano string

	ano = string(data[0]) + string(data[1]) + string(data[2]) + string(data[3])

	mes = string(data[5]) + string(data[6])

	dia = string(data[8]) + string(data[9])

	return entity.NewDate(dia, mes, ano)
}

func SelectAllTasks(db *sql.DB) []entity.Task {
	sttm, err := db.Prepare("SELECT conteudo,data_termino,status_conclusao FROM tarefas")

	if err != nil {
		log.Fatalf("Error na função selectalltasks <db.Prepare> %v", err)
	}

	rows, err := sttm.Query()

	if err != nil {
		log.Fatalf("Error na função selectalltask <sttm.Query> %v", err)
	}

	var tarefas []entity.Task

	for rows.Next() {
		var cont, data string
		var status bool
		var task entity.Task
		rows.Scan(&cont, &data, &status)

		fmt.Print(data)

		if status {
			task = *entity.NewTaskWithStatusOK(cont, *stringTodDate(data))
		} else {
			task = *entity.NewTask(cont, *stringTodDate(data))
		}

		tarefas = append(tarefas, task)
	}

	return tarefas
}
