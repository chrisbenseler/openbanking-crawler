package adapters

import (
	"openbankingcrawler/interfaces"
	"strconv"

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
	GetElectronicChannels(c *gin.Context)
	UpdateInstitutionElectronicChannels(c *gin.Context)
}

type controller struct {
	institutionInterface       interfaces.InstitutionInterface
	branchInterface            interfaces.BranchInterface
	electronicChannelInterface interfaces.ElectronicChannelInterface
}

//NewController create new controllers
func NewController(institutionInterface interfaces.InstitutionInterface,
	branchInterface interfaces.BranchInterface,
	electronicChannelInterface interfaces.ElectronicChannelInterface) Controller {

	return &controller{
		institutionInterface:       institutionInterface,
		branchInterface:            branchInterface,
		electronicChannelInterface: electronicChannelInterface,
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

	/*
		err := ctrl.institutionInterface.UpdateBranches(id)

		if err != nil {
			c.JSON(err.Status(), gin.H{"error": err.Message()})
			return
		}

	*/
	go ctrl.institutionInterface.UpdateBranches(id)

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

	page, errQuery := strconv.Atoi(c.Query("page"))

	if errQuery != nil {
		page = 1
	}

	branches, pagination, err := ctrl.branchInterface.GetFromInstitution(id, page)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, gin.H{"branches": branches, "pagination": pagination})

}

//UpdateInstitution update an institution electronic electronicChannels controller
func (ctrl *controller) UpdateInstitutionElectronicChannels(c *gin.Context) {
	id := c.Param("id")

	/*
		err := ctrl.institutionInterface.UpdateElectronicChannels(id)

		if err != nil {
			c.JSON(err.Status(), gin.H{"error": err.Message()})
			return
		}
	*/
	go ctrl.institutionInterface.UpdateElectronicChannels(id)

	c.JSON(200, gin.H{})
}

//GetElectronicChannels get electronicChannels from institution controller
func (ctrl *controller) GetElectronicChannels(c *gin.Context) {

	id := c.Param("id")

	page, errQuery := strconv.Atoi(c.Query("page"))

	if errQuery != nil {
		page = 1
	}

	electronicChannels, pagination, err := ctrl.electronicChannelInterface.GetFromInstitution(id, page)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(200, gin.H{"electronicChannels": electronicChannels, "pagination": pagination})

}
