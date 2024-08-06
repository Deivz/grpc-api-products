package controllers

import (
	"src/server/models"
)

func Create(
	uuid string,
	name string,
	typeProduct string,
	price string,
	description string,
) {
	produto := models.Produtos{
		Uuid:        uuid,
		Name:        name,
		Type:        typeProduct,
		Price:       price,
		Description: description,
	}

	produto.Save()
}

func GetAll() []models.Produtos {

	produtos := models.Produtos{}

	return produtos.Find()
}
