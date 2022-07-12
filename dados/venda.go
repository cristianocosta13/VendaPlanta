package dados

import (
	"awesomeProject/modelos"
)

type VendaPlanta struct {
	Planta   modelos.Planta
	Vendedor modelos.Vendedor
	Cliente  modelos.Cliente
}
