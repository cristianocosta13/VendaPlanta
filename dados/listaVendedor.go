package dados

import (
	"awesomeProject/modelos"
	"errors"
)

var vendedores []modelos.Vendedor

//func SalvarVendedor(vendedor modelos.Vendedor) {
//	vendedores = append(vendedores, vendedor)
//}

func SalvarVendedor(vendedor modelos.Vendedor) []modelos.Vendedor {
	vendedores = append(vendedores, vendedor)
	return vendedores
}

func BuscarVendedorPorCpf(cpf string) (*modelos.Vendedor, error) {
	for _, v := range vendedores {
		if v.Cpf == cpf {
			return &v, nil
		}
	}
	return nil, errors.New("Planta não encontrada!")
}