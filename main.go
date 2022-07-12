package main

import (
	"awesomeProject/modelos"
	"fmt"
	"awesomeProject/ui"
	"awesomeProject/dados"
)

func main() {

	var opc = 0

	fmt.Println("\nSISTEMA DE VENDAS")

	for opc != 6 {
		fmt.Print("\nMENU:\n1. CADASTRAR VENDEDOR\n2. CADASTRAR CLIENTE\n3. CADASTRAR PLANTA" +
			"\n4. VENDER PLANTA\n5. EXIBIR PLANTAS\n6.Sair\n")
		fmt.Println("Insira sua opção de acordo com o número correspondente: ")
		fmt.Scan(&opc)

		if opc == 1 {
			_, erro := ui.CadastrarVendedor()
			if erro != nil {
				fmt.Print(erro)
			}
		} else if opc == 2 {
			_, erro := ui.CadastrarCliente()
			if erro != nil {
				fmt.Print(erro)
			}
		} else if opc == 3 {
			_, erro := ui.CadastrarPlantas()
			if erro != nil {
				fmt.Print(erro)
			}
		} else if opc == 4 {
			ui.VenderPlanta()
		} else if opc == 5 {
			var cod int
			fmt.Print("Qual o código da planta a ser exibida?")
			fmt.Scan(&cod)
			planta, erro := dados.BuscarPorCod(cod)
			if erro != nil{
				fmt.Print(erro)
			}else{
				planta.Exibir()
			}
		} else {
			fmt.Print("\nOpção inválida\n")
		}
	}
}
