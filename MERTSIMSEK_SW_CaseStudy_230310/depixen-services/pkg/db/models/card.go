package models

import (
    "github.com/go-pg/pg/v10"
    
)

type Card struct {
    tableName struct{} `pg:"tb_casestudy"`
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imageuri string `json:"imageuri"`
	Createddate string `json:"createddate"`
}

func CreateCard(db *pg.DB, req *Card) (*Card, error) {
    _, err := db.Model(req).Insert()
    if err != nil {
        return nil, err
    }

    card := &Card{}

    err = db.Model(card).
        Where("card.id = ?", req.ID).
        Select()

    return card, err
}
func GetTheLastCard(db *pg.DB) (*Card, error) {
    card := &Card{}

    err := db.Model(card).
    Order("id DESC").
    Limit(1).
    Select()

    return card, err
}

func GetCardById(db *pg.DB, cardID string) (*Card, error) {
    card := &Card{}

    err := db.Model(card).
        Where("card.id = ?", cardID).
        Select()

    return card, err
}

func GetCards(db *pg.DB) ([]*Card, error) {
    cards := make([]*Card, 0)

    err := db.Model(&cards).
        Select()

    return cards, err
}