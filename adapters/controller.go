package adapters

import (
	"openbankingcrawler/interfaces"

	"github.com/gin-gonic/gin"
)

//Controller controller interface
type Controller interface {
	ListAllInstitutions(*gin.Context)
	GetInstitution(*gin.Context)
	UpdateInstitutionBranches(*gin.Context)
	CreateInstitution(*gin.Context)
	UpdateInstitution(*gin.Context)
	GetBranches(*gin.Context)
	GetChannels(c *gin.Context)
	UpdateInstitutionChannels(c *gin.Context)
}

type controller struct {
	institutionInterface interfaces.InstitutionInterface
	branchInterface      interfaces.BranchInterface
	channelInterface     interfaces.ChannelInterface
}

//NewController create new controllers
func NewController(institutionInterface interfaces.InstitutionInterface,
	branchInterface interfaces.BranchInterface,
	channelInterface interfaces.ChannelInterface) Controller {

	return &controller{
		institutionInterface: institutionInterface,
		branchInterface:      branchInterface,
		channelInterface:     channelInterface,
	}
}

//ListAllInstitutions list all institutions
func (ctrl *controller) ListAllInstitutions(c *gin.Context) {

	institutions, err := ctrl.institutionInterface.ListAll()

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, institutions)
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

	var payload InstitutionPayload

	c.BindJSON(&payload)

	institution, err := ctrl.institutionInterface.Create(payload.Name)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(201, institution)
}

//UpdateInstitution update an institution controller
func (ctrl *controller) UpdateInstitution(c *gin.Context) {

	id := c.Param("id")

	var payload InstitutionPayload

	c.BindJSON(&payload)

	institution, err := ctrl.institutionInterface.Update(id, payload.BaseURL)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, institution)
}

//GetBranches get branches from institution controller
func (ctrl *controller) GetBranches(c *gin.Context) {

	id := c.Param("id")

	branches, err := ctrl.branchInterface.GetFromInstitution(id)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, branches)

}

//UpdateInstitution update an institution electronic channels controller
func (ctrl *controller) UpdateInstitutionChannels(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.institutionInterface.UpdateChannels(id)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, gin.H{})
}

//GetChannels get channels from institution controller
func (ctrl *controller) GetChannels(c *gin.Context) {

	id := c.Param("id")

	channels, err := ctrl.channelInterface.GetFromInstitution(id)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, channels)

}
