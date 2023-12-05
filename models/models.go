package models

import "gorm.io/gorm"

type Gebruiker struct {
	gorm.Model
	Gebruikersnaam string `json:"gebruikersnaam" gorm:"text;not null;default:null`
	Wachtwoord     string `json:"wachtwoord" gorm:"text;not null;default:null`
	Status         string `json:"status" gorm:"text;not null;default:null`
	Sleutel        string `json:"sleutel" gorm:"text;not null;default:null`
}

type Sleutel struct {
	gorm.Model
	Code          string `json:"code" gorm:"text;not null;default:null`
	IsBeschikbaar string `json:"isbeschikbaar" gorm:"text;not null;default:null`
}
