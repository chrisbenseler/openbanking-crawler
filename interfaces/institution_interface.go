package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/dtos"
	"openbankingcrawler/services"
)

//InstitutionInterface service
type InstitutionInterface interface {
	Create(string) (*dtos.Institution, error)
	Delete(string) error
	Get(string) (*dtos.Institution, common.CustomError)
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

func (i *institutionInterface) Create(name string) (*dtos.Institution, error) {

	iDTO := dtos.Institution{Name: "teste"}

	institution, err := i.institutionService.Create(iDTO)
	if err != nil {
		return nil, err
	}

	return institution, nil

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

//Get get an institutuion
func (i *institutionInterface) Get(id string) (*dtos.Institution, common.CustomError) {
	return i.institutionService.Find(id)
}
