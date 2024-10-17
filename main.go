package main

import (
	"day-9/data"
	"fmt"
)

func main() {

	akun1 := data.BuatAkun("Ahmad", "Ahmad@example.com")
	data.DataAkun = append(data.DataAkun, akun1)

	akun2 := data.BuatAkun("Syarifuddin", "Ahmad@contoh.com")
	data.DataAkun = append(data.DataAkun, akun2)

	err := akun1.Saldo.TambahSaldo(-100)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil, Saldo %s Sekarang yaitu : %d\n", akun1.Nama, akun1.Saldo.Kas)
	}

	err = akun1.Saldo.TambahSaldo(1000)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil, Saldo %s Sekarang yaitu : %d\n", akun1.Nama, akun1.Saldo.Kas)
	}
	err = akun1.Saldo.KurangiSaldo(50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil, Saldo %s Sekarang yaitu : %d\n", akun1.Nama, akun1.Saldo.Kas)
	}

	err = akun2.Saldo.TambahSaldo(1000)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil, Saldo %s Sekarang yaitu : %d\n", akun2.Nama, akun2.Saldo.Kas)
	}

	err = akun2.Saldo.KurangiSaldo(10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Berhasil, Saldo %s Sekarang yaitu : %d\n", akun2.Nama, akun2.Saldo.Kas)
	}

	data.CariAkunBerdasarkanNama("Syarifuddin")
	data.CariAkunBerdasarkanNama("Ahmad")

	data.TampilkanSemuaAkun()

}
