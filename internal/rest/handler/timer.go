package handler

import (
	"net/http"
	"time"
	"whatisyourtime/internal/constants"
	"whatisyourtime/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetCurrentTime(c *gin.Context) {
	var req models.CurrentTimeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessageResponse{ErrorMessage: "invalid query datas"})
		logrus.Debugf("invalid query datas : %v", err)
		return
	}

	location, err := time.LoadLocation(req.Timezone)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessageResponse{ErrorMessage: "invalid query datas"})
		logrus.Debugf("time.LoadLocation() : %v", err)
		return
	}
	currentTime := time.Now().In(location)
	c.JSON(http.StatusOK, models.CurrentTimeResponse{
		Time:     currentTime.Format(constants.TimeLayout),
		Timezone: location.String(),
	})
}

func GetTimeInZone(c *gin.Context) {
	var req models.TargetTimeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessageResponse{ErrorMessage: "invalid query datas"})
		logrus.Debugf("invalid query datas : %v", err)
		return
	}

	if len(req.TargetInfo) == constants.Empty {
		c.JSON(http.StatusBadRequest, models.ErrorMessageResponse{ErrorMessage: "query data is empty"})
		return
	}

	// timezone load
	var response models.TargetTimeResponse
	var r models.TargetTimeInfo
	for _, target := range req.TargetInfo {
		location, err := time.LoadLocation(target.Timezone)
		if err != nil {
			continue
		}

		givenTime, err := time.Parse(constants.TimeLayout, target.Time)
		if err != nil {
			continue
		}

		r.Time = givenTime.Format(constants.TimeLayout)
		r.Timezone = location.String()
		response.TargetInfo = append(response.TargetInfo, r)
		r = models.TargetTimeInfo{}
	}
	c.JSON(http.StatusOK, response)
}

func GetWorldTimes(c *gin.Context) {
	var req models.WorldTimeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessageResponse{ErrorMessage: "invalid query datas"})
		logrus.Debugf("invalid query datas : %v", err)
		return
	}

	userLocation, err := time.LoadLocation(req.Timezone)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessageResponse{ErrorMessage: "invalid query datas"})
		logrus.Debugf("invalid query datas : %v", err)
		return
	}

	userTime := time.Now().In(userLocation)
	cities := map[string]string{
		"New York":    "America/New_York",
		"London":      "Europe/London",
		"Tokyo":       "Asia/Tokyo",
		"Seoul":       "Asia/Seoul",
		"Paris":       "Europe/Paris",
		"Los Angeles": "America/Los_Angeles",
	}

	worldTimes := make(map[string]string)
	for city, tz := range cities {
		location, err := time.LoadLocation(tz)
		if err != nil {
			logrus.Debugf("time.LoadLocation() error :%v", err)
			worldTimes[city] = "error loading timezone"
			continue
		}
		worldTimes[city] = time.Now().In(location).Format(constants.TimeLayout)
	}

	c.JSON(http.StatusOK, models.WorldTimeResponse{
		UserTime:     userTime.Format(constants.TimeLayout),
		UserTimezone: req.Timezone,
		CityTimes:    worldTimes,
	})
}
