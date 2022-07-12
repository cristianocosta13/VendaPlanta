package modelos

import "fmt"

type Cliente struct {
	Nome, Telefone, Endereco, Cpf string
}


func (Cliente) exibirCliente(Nome, Telefone, Endereco, Cpf string) {
	fmt.Print("Nome: ", Nome, "\nCPF: ", Cpf, "\nTelefone: ", Telefone)
}