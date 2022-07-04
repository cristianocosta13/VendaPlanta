package main

import (
	"awesomeProject/modelos"
	"fmt"
)

func main() {

	var opc = 0
	var vendedor Vendedor
	var planta modelos.Planta
	var cliente Cliente
	fmt.Println("\nSISTEMA DE VENDAS")

	for opc != 6 {
		fmt.Print("\nMENU:\n1. CADASTRAR VENDEDOR\n2. CADASTRAR CLIENTE\n3. CADASTRAR PLANTA" +
			"\n4. VENDER PLANTA\n5. EXIBIR PLANTAS\n6.Sair\n")
		fmt.Println("Insira sua opção de acordo com o número correspondente: ")
		fmt.Scan(&opc)
		if opc == 1 {
			vendedor = cadastrarVendedor()
			fmt.Print("\nVendedor cadastrado com sucesso!\n")
		} else if opc == 2 {
			cliente = vendedor.cadastrarCliente()
			fmt.Print("\nCliente cadastrado com sucesso, agora ele pode efetuar compras!\n")
		} else if opc == 3 {
			planta = vendedor.cadastrarPlantas()
			fmt.Print("\nPlanta cadastrada com sucesso!\n")
		} else if opc == 4 {
			venderPlanta(planta, vendedor, cliente)
			//if planta != undefined && vendedor != undefined && cliente != undefined{}
		} else if opc == 5 {
			planta.Exibir()
		} else {
			fmt.Print("\nOpção inválida\n")
		}
	}
}

func cadastrarVendedor() Vendedor {
	var nome, cpf, telefone string

	fmt.Print("\nSobre o vendedor, insira:\n")
	fmt.Print("Nome: ")
	fmt.Scan(&nome)
	fmt.Print("CPF: ")
	fmt.Scan(&cpf)
	fmt.Print("Telefone: ")
	fmt.Scan(&telefone)

	return Vendedor{
		nome:     nome,
		cpf:      cpf,
		telefone: telefone,
	}
}
