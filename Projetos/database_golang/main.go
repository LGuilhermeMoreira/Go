package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

type BankAccount struct {
	id     string
	name   string
	amount float64
}

func newBank(name string, amount float64) *BankAccount {
	return &BankAccount{
		id:     uuid.New().String(),
		name:   name,
		amount: amount,
	}
}

func (b BankAccount) show() {
	fmt.Println(b.id, b.name, b.amount)
}

func insert(db *sql.DB, b BankAccount) error {
	stmt, err := db.Prepare("insert into bank_account(id,name,amount) values (? , ? , ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(b.id, b.name, b.amount)

	if err != nil {
		return err
	}

	return nil
}

func loseAllMoney(db *sql.DB, id string) error {
	stmt, err := db.Prepare("update bank_account set amount = 0.0 where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func loseAccount(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from bank_account where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil

}

func selectOneAccount(db *sql.DB, id string) (*BankAccount, error) {
	stmt, err := db.Prepare("select id,name,amount from bank_account where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var b BankAccount

	// seleciona uma linha e coloca na struct criada
	err = stmt.QueryRow(id).Scan(&b.id, &b.name, &b.amount)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func selectAllAccounts(db *sql.DB) ([]BankAccount, error) {
	rows, err := db.Query("select id,name,amount from bank_account")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []BankAccount

	for rows.Next() {
		var a BankAccount

		err = rows.Scan(&a.id, &a.name, &a.amount)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	return accounts, nil

}

func main() {
	// connection user:password@tcp(localhost:3306)/database_name
	db, err := sql.Open("mysql", "guigui:guigui@tcp(localhost:5423)/learnDatabase")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database connected")
	}

	defer db.Close()

	// for i := range 10 {
	// 	balance := 5000.0005 * float64(i) // Corrected data type for balance
	// 	err := insert(db, *newBank(fmt.Sprintf("%v %v", "guigui", i), balance))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	p, err := selectAllAccounts(db)

	if err != nil {
		panic(err)
	}

	if len(p) == 0 {
		panic(errors.New("nenhuma conta foi encontrada"))
	}

	for _, value := range p {
		value.show()
	}

}
