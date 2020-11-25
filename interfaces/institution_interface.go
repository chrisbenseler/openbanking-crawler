package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/services"
)

//InstitutionInterface service
type InstitutionInterface interface {
	Create(string) error
	Delete(string) error
	Get(string) (*institution.Entity, common.CustomError)
}

type institutionInterface struct {
	institutionService services.InstitutionService
	branchService      services.BranchService
}

//NewInstitution create a new interface for institutions
func NewInstitution(institutionService services.InstitutionService, branchService services.BranchService) InstitutionInterface {

	return &institutionInterface{
		institutionService: institutionService,
		branchService:      branchService,
	}
}

func (i *institutionInterface) Create(name string) error {

	err := i.institutionService.Create(name)
	if err != nil {
		return err
	}

	return nil

}

func (i *institutionInterface) Delete(id string) error {

	err := i.institutionService.Delete(id)
	if err != nil {
		return err
	}

	deleteError := i.branchService.DeleteAllFromInstitution(id)

	if deleteError != nil {
		return deleteError
	}

	return nil
}

//GetByName get an institution by its name
func (i *institutionInterface) GetByName(name string) error {

	err := i.institutionService.Create(name)
	if err != nil {
		return err
	}

	return nil

}

//Get get an institutuion
func (i *institutionInterface) Get(id string) (*institution.Entity, common.CustomError) {
	return i.institutionService.Find(id)
}
