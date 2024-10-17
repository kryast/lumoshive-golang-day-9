package main

import (
	"day-9/data"
	"fmt"
)

func main() {
	akun1 := data.BuatAkun("Ahmad", "Ahmad@example.com")

	err := akun1.Saldo.TambahSaldo(100)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Saldo berhasil ditambahkan! [Nama: %s, Saldo: %d]\n", akun1.Nama, akun1.Saldo.Kas)
	}

	err = akun1.Saldo.KurangiSaldo(50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Saldo berhasil dikurangi! [Nama: %s, Saldo: %d]\n", akun1.Nama, akun1.Saldo.Kas)
	}

}
