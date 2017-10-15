package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

type Motorista struct {
	Nome            string `json:"nome"`
	GrupoDesconto   string `json:"grupo_desconto"`
	TelefoneCliente string `json:"telefone_cliente"`
	EmailCliente    string `json:"email_cliente"`
	CPF             string `json:"cpf"`
}

func main() {

	file, err := xlsx.OpenFile("dados.xlsx")
	if err != nil {
		fmt.Println("deu ruim xlsx:", err)
		return
	}

	sheet := file.Sheets[0]
	var motoristas []Motorista
	for _, row := range sheet.Rows {
		if len(row.Cells) >= 5 {
			if row.Cells[2].String() == "" {
				return
			}
			nome := strings.ReplaceAll(row.Cells[0].String(), " ", "")
			grupoDesconto := row.Cells[1].String()
			telefoneCliente := "55" + strings.ReplaceAll(row.Cells[2].String(), " ", "")
			emailCliente := row.Cells[3].String()
			cpf := row.Cells[4].String()
			motorista := Motorista{
				Nome:            nome,
				GrupoDesconto:   grupoDesconto,
				TelefoneCliente: telefoneCliente,
				EmailCliente:    emailCliente,
				CPF:             cpf,
			}
			motoristas = append(motoristas, motorista)
		}
	}

	jsonFile, err := os.Create("motoristas.json")
	if err != nil {
		fmt.Println("err > json file:", err)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(motoristas)
	if err != nil {
		fmt.Println("erro ao salvar:", err)
		return
	}
	fmt.Println("salvo")
}
