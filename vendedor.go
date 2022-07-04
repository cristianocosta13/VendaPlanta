package main

import (
	"awesomeProject/modelos"
	"fmt"
)

type Vendedor struct {
	nome, cpf, telefone string
}

func (Vendedor) exibirVendedor(nome, cpf, telefone string) {
	fmt.Print("Nome: ", nome, "\nCPF: ", cpf, "\nTelefone: ", telefone)
}

func (Vendedor) cadastrarPlantas() modelos.Planta {
	var especie, cor, porte, descricao string
	var estoque, codigo int
	var preco float32

	fmt.Print("\nSobre a planta, insira:\n")
	fmt.Print("Especie: ")
	fmt.Scan(&especie)
	fmt.Print("Cor: ")
	fmt.Scan(&cor)
	fmt.Print("Porte: ")
	fmt.Scan(&porte)
	fmt.Print("Preço: ")
	fmt.Scan(&preco)
	fmt.Print("Código: ")
	fmt.Scan(&codigo)
	fmt.Print("Estoque: ")
	fmt.Scan(&estoque)
	fmt.Print("Descrição: ")
	fmt.Scan(&descricao)

	return modelos.Planta{
		Especie:   especie,
		Cor:       cor,
		Porte:     porte,
		Preco:     preco,
		Estoque:   estoque,
		Descricao: descricao,
	}
}

func (Vendedor) cadastrarCliente() Cliente {
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

	return Cliente{
		nome:     nome,
		endereco: endereco,
		telefone: telefone,
		cpf:      cpf,
	}
}
