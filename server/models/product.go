package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Produtos struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

func (p *Produtos) Save() {
	filePath := "../database/produtos.txt"

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	productJSON, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	if _, err = writer.WriteString(string(productJSON) + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
}

func (p *Produtos) Find() []Produtos {
	filePath := "../database/produtos.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var products []Produtos
	for scanner.Scan() {
		var product Produtos
		err := json.Unmarshal([]byte(scanner.Text()), &product)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return nil
		}
		products = append(products, product)
	}
	return products
}
