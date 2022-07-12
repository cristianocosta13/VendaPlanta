package dados

import (
	"awesomeProject/modelos"
	"errors"
)

var plantas []modelos.Planta

//func SalvarPlanta(planta modelos.Planta) {
//	plantas = append(plantas, planta)
//}

func SalvarPlanta(planta modelos.Planta) []modelos.Planta {
	plantas = append(plantas, planta)
	return plantas
}

func BuscarPorCod(cod int) (*modelos.Planta, error) {
	for _, p := range plantas {
		if p.Codigo == cod {
			return &p, nil
		}
	}
	return nil, errors.New("Planta não encontrada!")
}
