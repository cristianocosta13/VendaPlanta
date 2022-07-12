package ui

import (
	"awesomeProject/modelos"
	"awesomeProject/dados"
	"fmt"
)

func CadastrarPlantas() (*modelos.Planta, error) {
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

	planta := modelos.Planta{
		Especie:   especie,
		Cor:       cor,
		Porte:     porte,
		Preco:     preco,
		Estoque:   estoque,
		Descricao: descricao,
		Codigo: codigo,
	}
	if especie != "" && cor != "" && porte != "" && preco != 0 && descricao != "" && codigo != 0 {
		dados.SalvarPlanta(planta)
		return &planta, nil
	}
}
