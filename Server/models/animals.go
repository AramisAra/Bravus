package models

// Model for animal table
type Animals struct {
	Base
	Animal_Name   string `json:"animal_name"`
	Animal_Specie string `json:"animal_specie"`
	Animal_Age    string `json:"animal_age"`
}
