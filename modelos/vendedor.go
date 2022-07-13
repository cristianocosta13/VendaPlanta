package modelos

import (
	"fmt"
)

type Vendedor struct {
	Nome, Cpf, Telefone string
	Plantas             []Planta
}

func (v Vendedor) Vendedor() {
	fmt.Print("Nome: ", v.Nome, "\nCPF: ", v.Cpf, "\nTelefone: ", v.Telefone)
}
