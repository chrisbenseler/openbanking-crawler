package adapters

import (
	"openbankingcrawler/interfaces"

	"github.com/gin-gonic/gin"
)

//Controller controller interface
type Controller interface {
	GetInstitution(*gin.Context)
	UpdateInstitutionBranches(*gin.Context)
	CreateInstitution(*gin.Context)
}

type controller struct {
	institutionInterface interfaces.InstitutionInterface
}

//NewController create new controllers
func NewController(institutionInterface interfaces.InstitutionInterface) Controller {

	return &controller{
		institutionInterface: institutionInterface,
	}
}

//GetInstitution get an institution controller
func (ctrl *controller) GetInstitution(c *gin.Context) {
	id := c.Param("id")

	institution, err := ctrl.institutionInterface.Get(id)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, institution)
}

//UpdateInstitution update an institution controller
func (ctrl *controller) UpdateInstitutionBranches(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.institutionInterface.UpdateBranches(id)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, gin.H{})
}

//CreateInstitution create an institution controller
func (ctrl *controller) CreateInstitution(c *gin.Context) {

	type institutionPayload struct {
		Name string `json:"name"`
	}

	var payload institutionPayload

	c.BindJSON(&payload)

	institution, err := ctrl.institutionInterface.Create(payload.Name)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(201, institution)
}
