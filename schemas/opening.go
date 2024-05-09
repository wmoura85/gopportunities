package schemas

import ("gorm.io/gorm")

type Opening struct {
	gorm.Model
	Funcao string
	Empresa string
	Local string
	Remoto bool
	Link string
	Salario int64
}