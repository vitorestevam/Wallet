package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	db  database
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.db, _ = NewDatabase(ctx, "mongodb://localhost:27018")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	fmt.Println(name)
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) PostTransaction(t Transaction) {
	err := a.db.InsertTransaction(&t)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *App) ListTransactions() []Transaction {
	transactions, err := a.db.GetTransactions(nil)
	if err != nil {
		fmt.Println(err)
		return make([]Transaction, 0)
	}

	return transactions
}

func (a *App) DeleteTransaction(t Transaction) {
	err := a.db.DeleteTransaction(t.ID.Hex())
	if err != nil {
		fmt.Println(err)
	}
}
