package model

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
}

type Moto struct {
	Automovel
	Cilindrada int
}

type Carro struct {
	Automovel
	Portas        int
	Potencia      int
	ArCodicionado bool
}
