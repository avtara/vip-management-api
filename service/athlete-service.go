package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/avtara/vip-management-system-api/dto"
	"github.com/avtara/vip-management-system-api/entity"
	"github.com/avtara/vip-management-system-api/repository"
	"github.com/mashingan/smapping"
)

//AthleteService is a contract about something that service can do
type AthleteService interface {
	AllAthlete() []entity.AthleteJSON
	ChangeStatusArrived(updateStatus dto.UpdateArrivedDTO, id string)
	DetailAthlete(userId string) entity.AthleteJSON
	InsertAthlete(order dto.InsertDTO)
}

type athleteService struct {
	athleteRepository repository.AthleteRepository
}

//NewHospitalService creates a new instance of AuthService
func NewAthleteService(athleteRepository repository.AthleteRepository) AthleteService {
	return &athleteService{
		athleteRepository: athleteRepository,
	}
}

func (service *athleteService) AllAthlete() []entity.AthleteJSON {
	var athletes []entity.Athlete = service.athleteRepository.AllAthlete()

	var result []entity.AthleteJSON
	for i := 0; i < len(athletes); i++ {
		if len(athletes[i].Attributes) != 0 {

			res := entity.AthleteJSON{
				ID:                athletes[i].ID,
				Name:              athletes[i].Name,
				Country_of_origin: athletes[i].Country_of_origin,
				Eta:               athletes[i].Eta,
				Arrived:           athletes[i].Arrived,
				Photo:             athletes[i].Photo,
				Attributes:        strings.Split(athletes[i].Attributes, ", "),
			}
			result = append(result, res)

		} else {
			var a []string = []string{"Tidak ada keterangan atribut!"}
			res := entity.AthleteJSON{
				ID:                athletes[i].ID,
				Name:              athletes[i].Name,
				Country_of_origin: athletes[i].Country_of_origin,
				Eta:               athletes[i].Eta,
				Arrived:           athletes[i].Arrived,
				Photo:             athletes[i].Photo,
				Attributes:        a,
			}
			result = append(result, res)

		}
	}
	return result
}

func (service *athleteService) ChangeStatusArrived(updateStatus dto.UpdateArrivedDTO, id string) {
	createOrder := entity.UpdateArrived{}
	err := smapping.FillStruct(&createOrder, smapping.MapFields(&updateStatus))
	fmt.Println(err)
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	service.athleteRepository.ChangeStatusArrived(id, updateStatus.Arrived)
}

func (service *athleteService) DetailAthlete(userId string) entity.AthleteJSON {
	var athlete entity.Athlete = service.athleteRepository.DetailAthlete(userId)
	if len(athlete.Attributes) != 0 {
		var result entity.AthleteJSON

		res := entity.AthleteJSON{
			ID:                athlete.ID,
			Name:              athlete.Name,
			Country_of_origin: athlete.Country_of_origin,
			Eta:               athlete.Eta,
			Arrived:           athlete.Arrived,
			Photo:             athlete.Photo,
			Attributes:        strings.Split(athlete.Attributes, ", "),
		}
		result = res

		return result
	} else {
		var result entity.AthleteJSON
		var a []string = []string{"Tidak ada keterangan atribut!"}
		res := entity.AthleteJSON{
			ID:                athlete.ID,
			Name:              athlete.Name,
			Country_of_origin: athlete.Country_of_origin,
			Eta:               athlete.Eta,
			Arrived:           athlete.Arrived,
			Photo:             athlete.Photo,
			Attributes:        a,
		}
		result = res

		return result
	}
}

func (service *athleteService) InsertAthlete(order dto.InsertDTO) {
	createOrder := entity.AthleteJSON{}
	err := smapping.FillStruct(&createOrder, smapping.MapFields(&order))
	fmt.Println(err)
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	service.athleteRepository.InsertAthlete(createOrder)
}
