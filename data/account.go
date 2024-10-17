package data

import "fmt"

type Account struct {
	Nama, Email string
	Saldo       Saldo
}

func BuatAkun(nama string, email string) Account {
	akun := Account{
		Nama:  nama,
		Email: email,
		Saldo: Saldo{Kas: 0},
	}
	fmt.Printf("Akun berhasil ditambahkan! [Nama: %s, Email: %s, Saldo: %d]\n", akun.Nama, akun.Email, akun.Saldo.Kas)

	return akun
}
