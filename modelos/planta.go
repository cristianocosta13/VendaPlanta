package modelos

import "fmt"

type Planta struct {
	Especie, Porte, Cor, Descricao string
	Preco                          float32
	Estoque, Codigo                int
}

func (p Planta) Planta() {
	fmt.Print("Espécie: ", p.Especie, "\nPorte: ", p.Porte, "\nCor: ", p.Cor, "\nPreço: ",
		p.Preco, "\nCódigo: ", p.Codigo, "\nDescrição: ", p.Descricao, "\nEstoque: ", p.Estoque)
}
