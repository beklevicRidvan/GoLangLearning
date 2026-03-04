package main

import (
	"errors"
	"fmt"
)

func main() {
	var menu string = "\n1) Hesap oluştur\n2) Para yatır\n3) Para çek\n4) Transfer\n5) Hesap görüntüle\nq) Çıkış"

	fmt.Println("------- ATM CLİ SİMÜLASYONUNA HOŞGELDİNİZ ------")
	fmt.Println("-------------------------------------------------")
	b := Bank{}
	var accService AccountService = &b
	for {

		fmt.Println(menu)

		var choice string
		fmt.Print("Tercihinizi yapın:")
		fmt.Scanln(&choice)

		if choice == "q" {
			exitFunc()
			break
		}
		switch choice {
		case "1":
			myAccount, err := accService.CreateAccount()
			if err != nil {
				fmt.Println("Hata oluştu; ", err)
			} else {
				myAccount.showAccountInfo()
			}
		case "2":
			err := accService.Deposit()
			if err != nil {
				fmt.Println("Hata oluştu; ", err)
			}
		case "3":
			err := accService.Withdraw()
			if err != nil {
				fmt.Println("Hata oluştu; ", err)
			}
		case "4":
			err := accService.Transfer()
			if err != nil {
				fmt.Println("Hata oluştu; ", err)
			}
		case "5":
			myAccount, err := accService.GetAccount()
			if err != nil {
				fmt.Println("Hata oluştu; ", err)
			} else {
				myAccount.showAccountInfo()
			}

		}
	}

}

func exitFunc() {
	fmt.Println("Atm'yi kullandığınız için teşekkür ederiz !")
	for i := 3; i > 0; i-- {
		fmt.Println("İşleminiz sonlandırılıyor ....", i)
	}
}

type MyAccount struct {
	id           int
	owner        string
	balance      float64
	Transactions []Transaction
}

func (a MyAccount) showAccountInfo() {
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Hesap ID: %d\n", a.id)
	fmt.Printf("Hesap Sahibi: %s\n", a.owner)
	fmt.Printf("Bakiye: %.2f TL\n", a.balance)
	fmt.Println("-------------------------------------------------")

	if len(a.Transactions) == 0 {
		fmt.Println("Henüz işlem geçmişi bulunmuyor.")
		return
	}

	fmt.Println("İşlem Geçmişi:")
	for i, t := range a.Transactions {
		fmt.Printf("%d) [%s] %.2f TL - %s\n",
			i+1,
			t.transactionType,
			t.amount,
			t.description,
		)
	}
	fmt.Println("-------------------------------------------------")
}

type Transaction struct {
	transactionType string // (deposit / withdraw / transfer)
	amount          float64
	description     string
}

func createDepositTransaction(amount float64) *Transaction {
	return &Transaction{
		transactionType: "deposit",
		amount:          amount,
		description:     fmt.Sprintf("Hesaba %.2f TL yatırıldı.", amount),
	}
}
func createWithDrawTransaction(amount float64) *Transaction {
	return &Transaction{transactionType: "withdraw", amount: amount, description: fmt.Sprintf("Hesaptan %.2f TL çekildi.", amount)}
}
func createReceiverTransferTransaction(amount float64, senderID int) *Transaction {

	return &Transaction{transactionType: "transfer", amount: amount, description: fmt.Sprintf("%d numaralı hesaptan %.2f TL alındı.", senderID, amount)}
}
func createSenderTransferTransaction(amount float64, receiverID int) *Transaction {

	return &Transaction{transactionType: "transfer", amount: amount, description: fmt.Sprintf("%d numaralı hesaba %.2f TL gönderildi.", receiverID, amount)}
}

type Bank struct {
	accounts map[int]*MyAccount
	nextID   int
}

func (b *Bank) CreateAccount() (*MyAccount, error) {
	var owner string
	fmt.Print("Hesap adı girin:")
	fmt.Scanln(&owner)
	for _, v := range b.accounts {
		if v.owner == owner {
			return nil, errors.New("Bankamızda zaten böyle bir hesap bulunuyor")
		}
	}
	b.nextID += 1

	myAccount := &MyAccount{id: b.nextID, owner: owner}

	if len(b.accounts) > 0 {
		b.accounts[b.nextID] = myAccount
	} else {
		b.accounts = map[int]*MyAccount{
			b.nextID: myAccount,
		}
	}

	return myAccount, nil

}

func (b *Bank) Deposit() error {

	var id int
	fmt.Print("Para yatırmak istediğiniz banka hesap id'sini girin:")
	fmt.Scanln(&id)
	if myAccount, exists := b.accounts[id]; exists {
		var amount float64
		fmt.Print("Yatırmak istediğiniz tutarı girin: ")
		fmt.Scanln(&amount)
		if amount <= 0 {
			return errors.New("Lütfen 0 dan büyük bir tutar girin")
		} else {
			myAccount.balance += amount
			transaction := createDepositTransaction(amount)
			myAccount.Transactions = append(myAccount.Transactions, *transaction)
			return nil
		}

	} else {
		return errors.New("Böyle bir hesap bulunamadı")
	}

}

func (b *Bank) Withdraw() error {
	var id int
	fmt.Print("Para çekmek istediğiniz banka hesap id'sini girin:")
	fmt.Scanln(&id)

	if myAccount, exists := b.accounts[id]; exists {
		var amount float64

		fmt.Print("Çekmek istediğiniz tutarı girin: ")
		fmt.Scanln(&amount)
		if amount <= 0 {
			return errors.New("Lütfen 0 dan büyük bir tutar girin")
		} else if amount > myAccount.balance {
			message := fmt.Sprintf("Lütfen hesabınızdaki bakiye kadar para çekin  Bakiye : %.2f ", myAccount.balance)
			return errors.New(message)
		} else {
			myAccount.balance -= amount
			transaction := createWithDrawTransaction(amount)
			myAccount.Transactions = append(myAccount.Transactions, *transaction)
			return nil

		}
	} else {
		return errors.New("Böyle bir hesap bulunamadı")

	}

}

// fromId sender || toId receiver
func (b *Bank) Transfer() error {
	var fromID, toID int
	fmt.Print("Para gönderecek hesap id'sini girin:")
	fmt.Scanln(&fromID)
	fmt.Print("Parayı alacak hesap id'sini girin:")
	fmt.Scanln(&toID)

	if fromID == toID {
		return errors.New("Lütfen birbirinden farklı hesap idleri girin!!")
	}
	senderAccount, senderExists := b.accounts[fromID]

	receiverAccount, receiverExists := b.accounts[toID]

	if senderExists && receiverExists {
		var amount float64

		fmt.Print("Transfer etmek istediğiniz tutarı girin: ")
		fmt.Scanln(&amount)
		if amount <= 0 {
			return errors.New("Lütfen 0 dan büyük bir tutar girin")
		} else if amount > senderAccount.balance {
			var message string = fmt.Sprintf("Gönderen hesapta bu kadar bakiye yok, gönderen hesap bakiyesi %.2f", senderAccount.balance)
			return errors.New(message)
		}

		senderAccount.balance -= amount
		receiverAccount.balance += amount

		receiverTransaction := createReceiverTransferTransaction(amount, senderAccount.id)
		senderTransaction := createSenderTransferTransaction(amount, receiverAccount.id)

		senderAccount.Transactions = append(senderAccount.Transactions, *senderTransaction)
		receiverAccount.Transactions = append(receiverAccount.Transactions, *receiverTransaction)

		return nil

	} else {
		return errors.New("Girdiğiniz hesap id'sine ilişkili hesap bulunamadı")
	}

}

func (b Bank) GetAccount() (*MyAccount, error) {

	var id int
	fmt.Print("Bulmak istediğiniz banka hesap id'sini girin:")
	fmt.Scanln(&id)
	if myAccount, exists := b.accounts[id]; exists {
		return myAccount, nil
	} else {
		return nil, errors.New("Bankamızda böyle bir hesap bulamadık !!")
	}
}

type AccountService interface {
	CreateAccount() (*MyAccount, error)

	Deposit() error

	Withdraw() error

	Transfer() error

	GetAccount() (*MyAccount, error)
}
