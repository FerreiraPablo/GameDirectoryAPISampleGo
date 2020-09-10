/*
Game

Estructura de datos de un juego

Pablo Ferreira 20200910
*/

package game

type Game struct {
	Id          int `gorm:"primaryKey;autoIncrement;not null"`
	Name        string
	Description string
	Genre       string
	Platform    string
}
