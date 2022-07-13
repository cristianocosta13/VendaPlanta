package modelos

import "fmt"

type Cliente struct {
	Nome, Telefone, Endereco, Cpf string
}

func (c Cliente) Cliente() {
	fmt.Print("Nome: ", c.Nome, "\nCPF: ", c.Cpf, "\nTelefone: ", c.Telefone)
}
