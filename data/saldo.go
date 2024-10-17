package data

import (
	"errors"
)

type Saldo struct {
	Kas int
}

func (s *Saldo) TambahSaldo(jumlah int) error {
	if jumlah <= 0 {
		return errors.New("jumlah saldo yang ditambahkan harus lebih dari 0")
	}
	s.Kas += jumlah
	return nil
}

func (s *Saldo) KurangiSaldo(jumlah int) error {
	if jumlah <= 0 {
		return errors.New("jumlah saldo yang dikurangi harus lebih dari 0")
	}
	if s.Kas < jumlah {
		return errors.New("saldo tidak mencukupi")
	}
	s.Kas -= jumlah
	return nil
}
