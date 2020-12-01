package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/dtos"
	"openbankingcrawler/services"
)

//InstitutionInterface interface
type InstitutionInterface interface {
	Create(string) (*dtos.Institution, common.CustomError)
	Delete(string) common.CustomError
	Get(string) (*dtos.Institution, common.CustomError)
	UpdateBranches(string) common.CustomError
	Update(string, string) (*dtos.Institution, common.CustomError)
}

type institutionInterface struct {
	institutionService institution.Service
	branchService      branch.Service
	crawler            services.Crawler
}

//NewInstitution create a new interface for institutions
func NewInstitution(institutionService institution.Service, branchService branch.Service, crawler services.Crawler) InstitutionInterface {

	return &institutionInterface{
		institutionService: institutionService,
		branchService:      branchService,
		crawler:            crawler,
	}
}

func (i *institutionInterface) Create(name string) (*dtos.Institution, common.CustomError) {

	iDTO := dtos.Institution{Name: name}

	institution, err := i.institutionService.Create(iDTO)
	if err != nil {
		return nil, err
	}

	return institution, nil

}

//Update update an institution attributes
func (i *institutionInterface) Update(id string, baseURL string) (*dtos.Institution, common.CustomError) {

	institution, err := i.institutionService.Read(id)

	if err != nil {
		return nil, err
	}

	institution.BaseURL = baseURL

	return i.institutionService.Update(*institution)

}

//Delete delete an institution
func (i *institutionInterface) Delete(id string) common.CustomError {

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
	return i.institutionService.Read(id)
}

//UpdateBranches update branches from institution
func (i *institutionInterface) UpdateBranches(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)

	if err != nil {
		return err
	}

	branches, crawlErr := i.crawler.CrawlBranches(institution.ID)

	if crawlErr != nil {
		return crawlErr
	}

	delErr := i.branchService.DeleteAllFromInstitution(id)

	if delErr != nil {
		return delErr
	}

	insertErr := i.branchService.InsertMany(*branches, id)
	if insertErr != nil {
		return insertErr
	}

	return nil

}
