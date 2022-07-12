package modelos

import (
	"fmt"
)

type Vendedor struct {
	Nome, Cpf, Telefone string
	Plantas             []Planta
}

func (Vendedor) exibirVendedor(nome, cpf, telefone string) {
	fmt.Print("Nome: ", nome, "\nCPF: ", cpf, "\nTelefone: ", telefone)
}
