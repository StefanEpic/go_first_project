package controllers

import (
	"first/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /
// @Summary Get list of reports
// @Tags base
// @Accept json
// @Produce json
// @Success 200 {array} schemas.ReadReports
// @Failure 400 {object} schemas.HTTPError
// @Router /api/reports [get]
func ReportsGetList(c *gin.Context) {
	var res = []db.Reports{}

	if err := db.Conn.Find(&res).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @BasePath /
// @Summary Add one report
// @Tags base
// @Accept json
// @Produce json
// @Param data body schemas.CreateReports true "data"
// @Success 200 {object} schemas.ReadReports
// @Failure 400 {object} schemas.HTTPError
// @Router /api/reports [post]
func ReportsAddOne(c *gin.Context) {
	var model db.Reports
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	if err := db.Conn.Create(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}

	c.JSON(http.StatusOK, model)
}

// @BasePath /
// @Summary Get one report
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Report ID"
// @Success 200 {object} schemas.ReadReports
// @Failure 404 {object} schemas.HTTPError
// @Router /api/reports/{id} [get]
func ReportsGetOne(c *gin.Context) {
	var model db.Reports

	if err := db.Conn.Preload("NerResults").Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Not found."})
		return
	}

	c.JSON(http.StatusOK, model)
}

// @BasePath /
// @Summary Edit one report
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Report ID"
// @Param data body schemas.UpdateReports true "data"
// @Success 200 {object} schemas.ReadReports
// @Failure 400 {object} schemas.HTTPError
// @Failure 404 {object} schemas.HTTPError
// @Router /api/reports/{id} [patch]
func ReportsEditOne(c *gin.Context) {
	var model db.Reports
	if err := db.Conn.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Not found."})
		return
	}

	var input db.Reports
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	if err := db.Conn.Model(&model).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}

	c.JSON(http.StatusOK, model)
}

// @BasePath /
// @Summary Delete one report
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Report ID"
// @Success 200 {object} schemas.ReadReports
// @Failure 400 {object} schemas.HTTPError
// @Failure 404 {object} schemas.HTTPError
// @Router /api/reports/{id} [delete]
func ReportsDeleteOne(c *gin.Context) {
	var model db.Reports
	if err := db.Conn.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Not found."})
		return
	}

	if err := db.Conn.Delete(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}

// @BasePath /
// @Summary Add all reports
// @Tags base
// @Accept json
// @Produce json
// @Param data body []schemas.CreateReports true "data"
// @Success 200 {object} schemas.HTTPError
// @Failure 400 {object} schemas.HTTPError
// @Router /api/reports_upload [post]
func ReportsAddAll(c *gin.Context) {
	var modelSlice []db.Reports
	if err := c.ShouldBindJSON(&modelSlice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	if err := db.Conn.CreateInBatches(&modelSlice, 1000).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}
