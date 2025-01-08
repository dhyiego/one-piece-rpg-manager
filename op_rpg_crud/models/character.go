package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name          string `json:"name"`
	Image         string `json:"image"`
	Level         int    `json:"level"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	Coins         int    `json:"coins"`
	Crew          string `json:"crew"`
	FightingStyle string `json:"fighting_style"`
	Race          string `json:"race"`
	Profession    string `json:"profession"`
}
