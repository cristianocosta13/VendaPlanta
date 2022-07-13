package ui

import (
	"awesomeProject/data"
	"awesomeProject/modelos"
	"errors"
	"fmt"
)

func CadastrarCliente() (*modelos.Cliente, error) {
	var nome, telefone, endereco, cpf string

	fmt.Print("\nSobre o cliente, insira:\n")
	fmt.Print("Nome: ")
	fmt.Scan(&nome)
	fmt.Print("Telefone: ")
	fmt.Scan(&telefone)
	fmt.Print("CPF: ")
	fmt.Scan(&cpf)
	fmt.Print("Endereço: ")
	fmt.Scan(&endereco)

	cliente := modelos.Cliente{
		Nome:     nome,
		Endereco: endereco,
		Telefone: telefone,
		Cpf:      cpf,
	}
	if nome != "" && endereco != "" && telefone != "" && cpf != "" {
		data.SalvarCliente(cliente)
		return &cliente, nil
	}

	return nil, errors.New("Dados incompativeis ou não totalmente preenchidos")
}
