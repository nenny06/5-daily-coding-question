package dailyquestions

import "fmt"

type Account struct {
    ID          int
    AccountName string
    Balance     float64
}

func (d *Account) DisplayAccount() {
    fmt.Println(d.AccountName)
    fmt.Println(d.Balance)
}

func main() {
    accounts := []Account{
        {ID: 1, AccountName: "Business Current", Balance: 50000},
        {ID: 2, AccountName: "Business Savings", Balance: 100000},
        {ID: 3, AccountName: "Petty Cash", Balance: 15000},
    }

    for _, account := range accounts {
        account.DisplayAccount()
    }
}