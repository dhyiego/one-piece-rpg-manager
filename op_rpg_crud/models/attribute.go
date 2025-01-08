package models

import "gorm.io/gorm"

type Attribute struct {
	gorm.Model
	CharacterID  uint `json:"character_id"`
	Strength     int  `json:"strength"`
	Combat       int  `json:"combat"`
	Agility      int  `json:"agility"`
	Accuracy     int  `json:"accuracy"`
	Vitality     int  `json:"vitality"`
	Intelligence int  `json:"intelligence"`
	Perception   int  `json:"perception"`
	Willpower    int  `json:"willpower"`
	Spirit       int  `json:"spirit"`
	Defense      int  `json:"defense" gorm:"-"`
	Charisma     int  `json:"charisma" gorm:"-"`
}
