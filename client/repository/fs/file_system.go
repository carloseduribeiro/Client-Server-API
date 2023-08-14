package fs

import (
	"github.com/carloseduribeiro/Client-Server-API/client/client"
	"os"
	"text/template"
)

const (
	defaultFileName  = "cotacao.txt"
	exchangeTemplate = "Dólar: {{.Bid}}"
)

type TextFileRepository struct {
	file *os.File
	temp *template.Template
}

func NewTextFileRepository() (*TextFileRepository, error) {
	f, err := os.Create(defaultFileName)
	if err != nil {
		return nil, err
	}
	t, err := template.New("cotacao").Parse(exchangeTemplate)
	if err != nil {
		return nil, err
	}
	return &TextFileRepository{file: f, temp: t}, nil
}

func (t TextFileRepository) SaveExchange(e *client.ExchangeDto) error {
	return t.temp.Execute(t.file, e)
}
