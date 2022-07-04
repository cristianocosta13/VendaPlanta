package main

type Cliente struct {
	nome, telefone, endereco, cpf string
	valorpago                     float32
}

func (c Cliente) pagarCompra(desconto float32) float32 {
	return c.valorpago - desconto
}
