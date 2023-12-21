package main

import (
	"estudo/model"
	"fmt"
	"time"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua X",
		Numero: 15,
		Cidade: "cidade x",
	}
	pessoa := model.Pessoa{
		Nome:     "José",
		Endereco: endereco,
		DataDeNascimento: time.Date(1998, 06, 18, 0, 0, 0, 0, time.Local),
	}

	automovelMoto := model.Automovel{
		Ano: 2022,
		Placa: "XPTO",
		Modelo: "CG",
	}
	automovelCarro := model.Automovel{
		Ano: 2009,
		Placa: "GOO-0889",
		Modelo: "Uno",
	}
	carro := model.Carro{
		Automovel: automovelCarro,
		Portas: 4,
		Potencia: 1,
		ArCodicionado: false,
	}
	moto := model.Moto{
		Automovel: automovelMoto,
		Cilindrada: 125,
	}

	// idade := model.CalculaIdade(pessoa)
	

	fmt.Println("Cidade:", endereco.Cidade)
	fmt.Println("Rua:", endereco.Rua)
	fmt.Println("Numero:", endereco.Numero)
	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Data de Nascimanto:",pessoa.DataDeNascimento)
	pessoa.IdadeAtual()
	fmt.Println("Idade:", pessoa.Idade)
	fmt.Println("****** Moto ******")
	fmt.Println("Ano:",moto.Ano)
	fmt.Println("Modelo:",moto.Modelo)
	fmt.Println("Placa:",moto.Placa)
	fmt.Println("Cilindrada:",moto.Cilindrada)
	fmt.Println("****** Carro ******")
	fmt.Println("Ano:",carro.Ano)
	fmt.Println("Modelo:",carro.Modelo)
	fmt.Println("Placa:",carro.Placa)
	fmt.Println("Potência:",carro.Potencia)
	fmt.Println("Portas:",carro.Portas)
	fmt.Println("Ar Condicionado:",carro.ArCodicionado)
}
