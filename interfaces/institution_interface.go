package interfaces

import (
	"openbankingcrawler/services"
)

//InstitutionInterface service
type InstitutionInterface interface {
	Create(string) error
	Delete(string) error
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
