package controller

import (
	"net/http"

	"github.com/avtara/vip-management-system-api/dto"
	"github.com/avtara/vip-management-system-api/entity"
	"github.com/avtara/vip-management-system-api/helper"
	"github.com/avtara/vip-management-system-api/service"
	"github.com/gin-gonic/gin"
)

type AthleteController interface {
	AllAthlete(ctx *gin.Context)
	UpdateArrivedStatus(ctx *gin.Context)
	DetailAthlete(ctx *gin.Context)
	InsertAthlete(ctx *gin.Context)
}

type athleteController struct {
	athleteService service.AthleteService
}

//NewHospitalController create a new instances of BoookController
func NewAthleteController(athleteService service.AthleteService) AthleteController {
	return &athleteController{
		athleteService: athleteService,
	}
}

func (c *athleteController) AllAthlete(ctx *gin.Context) {
	var athletes []entity.AthleteJSON = c.athleteService.AllAthlete()
	res := helper.BuildResponse(true, "OK", athletes)
	ctx.JSON(http.StatusOK, res)
}

func (c *athleteController) UpdateArrivedStatus(ctx *gin.Context) {
	var updateArrivedDTO dto.UpdateArrivedDTO
	errDTO := ctx.ShouldBind(&updateArrivedDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	id := ctx.Param("id")
	c.athleteService.ChangeStatusArrived(updateArrivedDTO, id)
	response := helper.BuildResponseWithoutData(true, "OK!")
	ctx.JSON(http.StatusOK, response)
}

func (c *athleteController) DetailAthlete(ctx *gin.Context) {
	id := ctx.Param("id")
	var hospitals entity.AthleteJSON = c.athleteService.DetailAthlete(id)
	res := helper.BuildResponse(true, "OK", hospitals)
	ctx.JSON(http.StatusOK, res)
}

func (c *athleteController) InsertAthlete(ctx *gin.Context) {
	var insertDTO dto.InsertDTO
	if err := ctx.ShouldBindJSON(&insertDTO); err == nil {
		c.athleteService.InsertAthlete(insertDTO)
		response := helper.BuildResponseWithoutData(true, "OK!")
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

}
