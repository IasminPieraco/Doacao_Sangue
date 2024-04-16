package models

import (
	"time"
)

type RH int64
type TipoSanguineo int64

const (
	A  TipoSanguineo = 0
	B  TipoSanguineo = 1
	AB TipoSanguineo = 2
	O  TipoSanguineo = 3
)

const (
	POSITIVO RH = 1
	NEGATIVO RH = 0
)

type Doador struct {
	Codigo         int
	Nome           string
	CPF            string
	Contato        string
	TipoDoSangue   TipoSanguineo
	Rh             RH
	TipoRHCorretos bool
}

type Doacao struct {
	Codigo int
	Data   time.Time
	Hora   time.Duration
	Volume float64
}
