package data

import (
	"awesomeProject/modelos"
	"errors"
)

var clientes []modelos.Cliente

//func SalvarCliente(cliente modelos.Cliente) {
//	clientes = append(clientes, cliente)
//}

func SalvarCliente(cliente modelos.Cliente) []modelos.Cliente {
	clientes = append(clientes, cliente)
	return clientes
}

func BuscarClientePorCpf(cpf string) (*modelos.Cliente, error) {
	for _, c := range clientes {
		if c.Cpf == cpf {
			return &c, nil
		}
	}
	return nil, errors.New("Cliente não encontrado!")
}
