package models

import "gorm.io/gorm"

type Fruit struct {
	gorm.Model
	CharacterID uint   `json:"character_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Abilities   string `json:"abilities"`
}
