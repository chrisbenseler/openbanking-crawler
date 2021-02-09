package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/dtos"
	"openbankingcrawler/services"
)

//InstitutionInterface interface
type InstitutionInterface interface {
	ListAll() ([]dtos.Institution, common.CustomError)
	Create(string) (*dtos.Institution, common.CustomError)
	Delete(string) common.CustomError
	Get(string) (*dtos.Institution, common.CustomError)
	UpdateBranches(string) common.CustomError
	UpdateElectronicChannels(string) common.CustomError
	Update(string, string) (*dtos.Institution, common.CustomError)
	UpdatePersonalLoans(string) common.CustomError
	UpdatePersonalCreditCards(string) common.CustomError
	UpdatePersonalAccounts(string) common.CustomError
}

type institutionInterface struct {
	institutionService        institution.Service
	branchService             branch.Service
	electronicChannelService  electronicchannel.Service
	personalLoanService       personalloan.Service
	personalCreditCardService personalcreditcard.Service
	personalAccountService    personalaccount.Service
	crawler                   services.Crawler
}

//NewInstitution create a new interface for institutions
func NewInstitution(institutionService institution.Service, branchService branch.Service, electronicChannelService electronicchannel.Service, personalLoanService personalloan.Service, personalCreditCardService personalcreditcard.Service, personalAccountService personalaccount.Service, crawler services.Crawler) InstitutionInterface {

	return &institutionInterface{
		institutionService:        institutionService,
		branchService:             branchService,
		electronicChannelService:  electronicChannelService,
		personalLoanService:       personalLoanService,
		personalCreditCardService: personalCreditCardService,
		personalAccountService:    personalAccountService,
		crawler:                   crawler,
	}
}

func (i *institutionInterface) ListAll() ([]dtos.Institution, common.CustomError) {

	return i.institutionService.List()
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

	deleteError = i.electronicChannelService.DeleteAllFromInstitution(id)

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

	branches, crawlErr := i.crawler.Branches(institution.BaseURL, 1, []branch.Entity{})

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

//UpdateElectronicChannels update electronicChannels from institution
func (i *institutionInterface) UpdateElectronicChannels(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)

	if err != nil {
		return err
	}

	electronicChannels, crawlErr := i.crawler.ElectronicChannels(institution.BaseURL, 1, []electronicchannel.Entity{})

	if crawlErr != nil {
		return crawlErr
	}

	delErr := i.electronicChannelService.DeleteAllFromInstitution(id)

	if delErr != nil {
		return delErr
	}

	insertErr := i.electronicChannelService.InsertMany(*electronicChannels, id)
	if insertErr != nil {
		return insertErr
	}

	return nil

}

//UpdatePersonalLoans update personalLoans from institution
func (i *institutionInterface) UpdatePersonalLoans(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)

	if err != nil {
		return err
	}

	personalLoans, crawlErr := i.crawler.PersonalLoans(institution.BaseURL, 1, []personalloan.Entity{})

	if crawlErr != nil {
		return crawlErr
	}

	delErr := i.personalLoanService.DeleteAllFromInstitution(id)

	if delErr != nil {
		return delErr
	}

	insertErr := i.personalLoanService.InsertMany(*personalLoans, id)
	if insertErr != nil {
		return insertErr
	}

	return nil

}

//UpdatePersonalCreditCards update creditcards from institution
func (i *institutionInterface) UpdatePersonalCreditCards(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)

	if err != nil {
		return err
	}

	personalCreditCards, crawlErr := i.crawler.PersonalCreditCards(institution.BaseURL, 1, []personalcreditcard.Entity{})

	if crawlErr != nil {
		return crawlErr
	}

	delErr := i.personalCreditCardService.DeleteAllFromInstitution(id)

	if delErr != nil {
		return delErr
	}

	insertErr := i.personalCreditCardService.InsertMany(*personalCreditCards, id)
	if insertErr != nil {
		return insertErr
	}

	return nil

}

//UpdatePersonalAccounts update accounts from institution
func (i *institutionInterface) UpdatePersonalAccounts(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)

	if err != nil {
		return err
	}

	personalCreditCards, crawlErr := i.crawler.PersonalAccounts(institution.BaseURL, 1, []personalaccount.Entity{})

	if crawlErr != nil {
		return crawlErr
	}

	delErr := i.personalAccountService.DeleteAllFromInstitution(id)

	if delErr != nil {
		return delErr
	}

	insertErr := i.personalAccountService.InsertMany(*personalCreditCards, id)
	if insertErr != nil {
		return insertErr
	}

	return nil

}
