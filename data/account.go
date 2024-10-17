package data

import "fmt"

type Account struct {
	Nama, Email string
	Saldo       Saldo
}

var DataAkun []*Account

func BuatAkun(nama string, email string) *Account {
	akun := Account{
		Nama:  nama,
		Email: email,
		Saldo: Saldo{Kas: 0},
	}
	fmt.Printf("Akun berhasil ditambahkan! [Nama: %s, Email: %s, Saldo: %d]\n", akun.Nama, akun.Email, akun.Saldo.Kas)

	return &akun
}

func CariAkunBerdasarkanNama(nama string) {

	for _, akun := range DataAkun {
		if akun.Nama == nama {
			fmt.Printf("Data Akun %s Yaitu %+v\n ", nama, akun)
		}
	}
}

func TampilkanSemuaAkun() {

	for i, akun := range DataAkun {
		fmt.Printf("Data akun ke %d Adalah %+v\n", i+1, akun)
	}
}
