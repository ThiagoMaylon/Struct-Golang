package model

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

func (p *Pessoa) IdadeAtual() {
	AnoNascimento := p.DataDeNascimento.Year()
	AnoAtual := time.Now().Year()
	p.Idade = AnoAtual - AnoNascimento
}

func CalculaIdade(p Pessoa) int {
	AnoNascimento := p.DataDeNascimento.Year()
	AnoAtual := time.Now().Year()
	return AnoAtual - AnoNascimento
}
