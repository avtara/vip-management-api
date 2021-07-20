package repository

import (
	"github.com/avtara/vip-management-system-api/entity"
	"gorm.io/gorm"
)

//AthletelRepository is contract what userRepository can do to db
type AthleteRepository interface {
	AllAthlete() []entity.Athlete
	ChangeStatusArrived(userID string, status bool)
	DetailAthlete(userId string) entity.Athlete
	InsertAthlete(athlete entity.AthleteJSON)
}

type athleteConnection struct {
	connection *gorm.DB
}

//NewAthleteRepository is creates a new instance of UserRepository
func NewAthleteRepository(db *gorm.DB) AthleteRepository {
	return &athleteConnection{
		connection: db,
	}
}

func (db *athleteConnection) AllAthlete() []entity.Athlete {
	var result []entity.Athlete
	db.connection.Raw("SELECT ath.*, GROUP_CONCAT(att.name SEPARATOR ', ') AS 'attributes' FROM athletes ath LEFT JOIN attributes att ON ath.id = att.id GROUP BY ath.id").Scan(&result)
	return result
}

func (db *athleteConnection) ChangeStatusArrived(userID string, status bool) {
	var result entity.Athlete
	db.connection.Model(&result).Where("id = ?", userID).Update("arrived", status)
}

func (db *athleteConnection) DetailAthlete(userId string) entity.Athlete {
	var result entity.Athlete
	db.connection.Raw("SELECT ath.*, GROUP_CONCAT(att.name SEPARATOR ', ') AS 'attributes' FROM athletes ath LEFT JOIN attributes att ON ath.id = att.id WHERE ath.id = ? GROUP BY ath.id", userId).Scan(&result)
	return result
}

func (db *athleteConnection) InsertAthlete(athlete entity.AthleteJSON) {
	db.connection.Save(&athlete)
}
