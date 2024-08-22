package controllers

import (
	"first/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /
// @Summary Get list of ner results
// @Tags base
// @Accept json
// @Produce json
// @Success 200 {array} schemas.ReadNerResults
// @Failure 400 {object} schemas.HTTPError
// @Router /api/ner_results [get]
func NerResultsGetList(c *gin.Context) {
	var res = []db.NerResults{}

	if err := db.Conn.Find(&res).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @BasePath /
// @Summary Add one ner_result
// @Tags base
// @Accept json
// @Produce json
// @Param data body schemas.CreateNerResults true "data"
// @Success 200 {object} schemas.ReadNerResults
// @Failure 400 {object} schemas.HTTPError
// @Router /api/ner_results [post]
func NerResultsAddOne(c *gin.Context) {
	var model db.NerResults
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
// @Summary Get one ner_result
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Ner Result ID"
// @Success 200 {object} schemas.ReadNerResults
// @Failure 404 {object} schemas.HTTPError
// @Router /api/ner_results/{id} [get]
func NerResultsGetOne(c *gin.Context) {
	var model db.NerResults

	if err := db.Conn.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Not found."})
		return
	}

	c.JSON(http.StatusOK, model)
}

// @BasePath /
// @Summary Edit one ner_Result
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Ner Result ID"
// @Param data body schemas.UpdateNerResults true "data"
// @Success 200 {object} schemas.ReadNerResults
// @Failure 400 {object} schemas.HTTPError
// @Failure 404 {object} schemas.HTTPError
// @Router /api/ner_results/{id} [patch]
func NerResultsEditOne(c *gin.Context) {
	var model db.NerResults
	if err := db.Conn.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Not found."})
		return
	}

	var input db.NerResults
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
// @Summary Delete one ner result
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Ner Result ID"
// @Success 200 {object} schemas.ReadNerResults
// @Failure 400 {object} schemas.HTTPError
// @Failure 404 {object} schemas.HTTPError
// @Router /api/ner_results/{id} [delete]
func NerResultsDeleteOne(c *gin.Context) {
	var model db.NerResults
	if err := db.Conn.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"detail": "Not found."})
		return
	}

	if err := db.Conn.Delete(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}

// @BasePath /
// @Summary Add all ner results
// @Tags base
// @Accept json
// @Produce json
// @Param data body []schemas.CreateNerResults true "data"
// @Success 200 {object} schemas.HTTPError
// @Failure 400 {object} schemas.HTTPError
// @Router /api/ner_results_upload [post]
func NerResultsAddAll(c *gin.Context) {
	var modelSlice []db.NerResults
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
