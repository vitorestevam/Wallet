package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/vitorestevam/wallet/db"
)

type PageData struct {
	Transactions []db.Transaction
}

func main() {

	database, _ := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("METHOD:", r.Method)

		switch r.Method {
		case http.MethodPost:
			value, _ := strconv.ParseFloat(r.FormValue("value"), 64)

			transaction := db.TransactionBase{
				Name:        r.FormValue("name"),
				Description: r.FormValue("description"),
				Value:       value,
				Kind:        db.Paying,
			}
			db.CreateTransaction(database, transaction)

		case http.MethodDelete:
			resp, _ := io.ReadAll(r.Body)
			id, _ := strconv.Atoi(string(resp))

			db.DeleteTransaction(database, id)
		}

		tmpl := template.Must(template.ParseFiles("layout.html"))
		data := PageData{
			Transactions: db.GetTransactions(database),
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8081", nil)

}
