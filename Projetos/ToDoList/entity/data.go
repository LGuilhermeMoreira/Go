package entity

import (
	"errors"
	"regexp"
	"strconv"
)

type Date struct {
	dia string
	mes string
	ano string
}

func DateToString(date Date) string {
	return date.ano + "-" + date.mes + "-" + date.dia
}

func NewDate(dia, mes, ano string) *Date {
	return &Date{
		dia: dia,
		mes: mes,
		ano: ano,
	}
}

func (d *Date) Show() string {
	return d.dia + "/" + d.mes + "/" + d.ano
}

func NewDateFormat(dataString string) (*Date, error) {
	// Utilizar expressão regular para validar o formato da data
	matched, err := regexp.MatchString(`^\d{2}/\d{2}/\d{4}$`, dataString)
	if err != nil || !matched {
		return nil, errors.New("formato de data inválido. Use o formato DD/MM/AAAA")
	}

	// Separar dia, mês e ano da string
	dia := dataString[0:2]
	mes := dataString[3:5]
	ano := dataString[6:10]

	date := &Date{dia, mes, ano}

	if err := validateDate(date); err != nil {
		return nil, err
	}

	return date, nil
}

func validateDate(date *Date) error {
	// Convertendo strings para inteiros para realizar as verificações
	ano, err := strconv.Atoi(date.ano)
	if err != nil {
		return errors.New("ano inválido")
	}

	mes, err := strconv.Atoi(date.mes)
	if err != nil || mes < 1 || mes > 12 {
		return errors.New("mês inválido")
	}

	dia, err := strconv.Atoi(date.dia)
	if err != nil || dia < 1 || dia > 31 {
		return errors.New("dia inválido")
	}

	// Verificar se o dia é válido para o mês
	if (mes == 4 || mes == 6 || mes == 9 || mes == 11) && dia > 30 {
		return errors.New("dia inválido para o mês")
	} else if mes == 2 {
		// Verificar se é fevereiro e se o dia é válido para o ano bissexto
		if (ano%4 != 0 || (ano%100 == 0 && ano%400 != 0)) && dia > 28 {
			return errors.New("dia inválido para fevereiro (não bissexto)")
		} else if dia > 29 {
			return errors.New("dia inválido para fevereiro (bissexto)")
		}
	}

	return nil
}
