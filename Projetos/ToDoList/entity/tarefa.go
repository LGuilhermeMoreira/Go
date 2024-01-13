package entity

import (
	"strconv"
)

type Task struct {
	conteudo string
	date     Date
	status   bool
}

func NewTask(conteudo string, date Date) *Task {
	return &Task{
		conteudo: conteudo,
		date:     date,
		status:   false,
	}
}

func NewTaskWithStatusOK(conteudo string, date Date) *Task {
	return &Task{
		conteudo: conteudo,
		date:     date,
		status:   true,
	}
}

// altera o ID de uma Task
// func UpdateTaskID(Task *Task, id int) {
// 	Task.id = id
// }

// altera conteudo da Task
func UpdateTaskConteudo(Task *Task, conteudo string) {
	Task.conteudo = conteudo
}

// altera horario da Task
// func UpdateTaskHorario(Task *Task, horario string) {
// 	Task.horario = horario
// }

// altera date da Task
func UpdateTaskDate(Task *Task, date Date) {
	Task.date = date
}

// altera status da Task
func UpdateTaskStatus(Task *Task, status bool) {
	Task.status = status
}

// retorna valor ID da Task
// func ReadTaskID(Task *Task) int {
// 	return Task.id
// }

// retorna conteudo da Task
func ReadTaskConteudo(Task *Task) string {
	return Task.conteudo
}

// retorna Horario da Task
// func ReadTaskHorario(Task *Task) string {
// 	return Task.horario
// }

// retorna conteudo da Task
func ReadTaskDate(Task *Task) string {
	return DateToString(Task.date)
}

// retorna conteudo da Task
func ReadTaskStatus(Task *Task) bool {
	return Task.status
}

func (t *Task) Show() string {
	return "Conteudo: " + t.conteudo + "\t" + "Data: " + t.date.Show() + "\t" + "Concluido: " + strconv.FormatBool(t.status)
}
