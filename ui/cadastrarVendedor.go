package ui

import (
	"awesomeProject/modelos"
	"awesomeProject/dados"
	"fmt"
)

func CadastrarVendedor() (*modelos.Vendedor, error) {
	var nome, cpf, telefone string

	fmt.Print("\nSobre o vendedor, insira:\n")
	fmt.Print("Nome: ")
	fmt.Scan(&nome)
	fmt.Print("CPF: ")
	fmt.Scan(&cpf)
	fmt.Print("Telefone: ")
	fmt.Scan(&telefone)

	vendedor := modelos.Vendedor{
		Nome:     nome,
		Cpf:      cpf,
		Telefone: telefone,
	}
	if nome != "" && telefone != "" && cpf != "" {
		dados.SalvarVendedor(vendedor)
		return &vendedor, nil
	}
}
