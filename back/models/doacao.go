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
	Codigo         int    `json:"codigo"`
	Nome           string `json:"nome"`
	CPF            string `json:"cpf"`
	Contato        string `json:"contato"`
	TipoDoSangue   string `json:"tipodosangue"`
	Rh             string `json:"rh"`
	TipoRHCorretos bool   `json:"tiporhcorretos"`
}

type Doacao struct {
	Codigo int
	Data   time.Time
	Hora   time.Duration
	Volume float64
}
