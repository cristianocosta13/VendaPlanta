package main

import (
	"awesomeProject/data"
	"awesomeProject/ui"
	"fmt"
)

func main() {

	var opc = 0

	fmt.Println("\nSISTEMA DE VENDAS")

	for opc != 6 {
		fmt.Print("\nMENU:\n1. CADASTROS" +
			"\n2. VENDER PLANTA\n3. EXIBIÇÕES\n4. EXCLUSÕES\n5. CONSULTAR ESTOQUE\n6. SAIR\n")
		fmt.Println("Insira sua opção de acordo com o número correspondente: ")
		fmt.Scan(&opc)

		if opc == 1 {
			Cadastros()
		} else if opc == 2 {
			ui.VenderPlanta()
		} else if opc == 3 {
			Exibicoes()
		} else if opc == 4 {
			Exclusoes()
		} else if opc == 5 {
			fmt.Print("INDISPONÍVEL")
		} else {
			fmt.Print("\nOpção inválida\n")
		}
	}
}

func Cadastros() {
	var choice int
	fmt.Print("O que você deseja cadastrar? (VENDEDOR - 1 | CLIENTE - 2 | PLANTA - 3): ")
	fmt.Scan(&choice)
	if choice == 1 {
		_, erro := ui.CadastrarVendedor()
		if erro != nil {
			fmt.Print(erro)
		}
	} else if choice == 2 {
		_, erro := ui.CadastrarCliente()
		if erro != nil {
			fmt.Print(erro)
		}
	} else if choice == 3 {
		_, erro := ui.CadastrarPlantas()
		if erro != nil {
			fmt.Print(erro)
		}
	} else {
		fmt.Print("\nTipo de cadastro inválido\n")
	}
}

func Exibicoes() {
	var choice int
	fmt.Print("O que você deseja exibir? (VENDEDOR - 1 | CLIENTE - 2 | PLANTA - 3): ")
	fmt.Scan(&choice)
	if choice == 1 {
		var cpf string
		fmt.Print("Qual o CPF do vendedor a ser exibido?\n")
		fmt.Scan(&cpf)
		vendedor, erro := data.BuscarVendedorPorCpf(cpf)
		if erro != nil {
			fmt.Print(erro)
		} else {
			vendedor.Vendedor()
		}
	} else if choice == 2 {
		var cpf string
		fmt.Print("Qual o CPF do cliente a ser exibido?\n")
		fmt.Scan(&cpf)
		cliente, erro := data.BuscarClientePorCpf(cpf)
		if erro != nil {
			fmt.Print(erro)
		} else {
			cliente.Cliente()
		}
	} else if choice == 3 {
		var cod int
		fmt.Print("Qual o código da planta a ser exibida?\n")
		fmt.Scan(&cod)
		planta, erro := data.BuscarPorCod(cod)
		if erro != nil {
			fmt.Print(erro)
		} else {
			planta.Planta()
		}
	} else {
		fmt.Print("\nTipo de exibição inválida\n")
	}
}

func Exclusoes() {
	fmt.Print("INDISPONÍVEL")
}
