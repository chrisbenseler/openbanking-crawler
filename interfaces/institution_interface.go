package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/businessaccount"
	"openbankingcrawler/domain/businesscreditcard"
	"openbankingcrawler/domain/businessfinancing"
	"openbankingcrawler/domain/businessloan"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalfinancing"
	"openbankingcrawler/domain/personalinvoicefinancing"
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
	UpdatePersonalAccounts(string) common.CustomError
	UpdatePersonalLoans(string) common.CustomError
	UpdatePersonalFinancings(string) common.CustomError
	UpdatePersonalInvoiceFinancings(string) common.CustomError
	UpdatePersonalCreditCards(string) common.CustomError
	UpdateBusinessAccounts(string) common.CustomError
	UpdateBusinessLoans(string) common.CustomError
	UpdateBusinessFinancings(string) common.CustomError
	UpdateBusinessCreditCards(string) common.CustomError
}

type institutionInterface struct {
	institutionService              institution.Service
	branchService                   branch.Service
	electronicChannelService        electronicchannel.Service
	personalAccountService          personalaccount.Service
	personalLoanService             personalloan.Service
	personalFinancingService        personalfinancing.Service
	personalInvoiceFinancingService personalinvoicefinancing.Service
	personalCreditCardService       personalcreditcard.Service
	businessAccountService          businessaccount.Service
	businessLoanService             businessloan.Service
	businessFinancingService        businessfinancing.Service
	businessCreditCardService       businesscreditcard.Service
	crawler                         services.Crawler
}

//NewInstitution create a new interface for institutions
func NewInstitution(institutionService institution.Service,
	branchService branch.Service,
	electronicChannelService electronicchannel.Service,
	personalAccountService personalaccount.Service,
	personalLoanService personalloan.Service,
	personalFinancingService personalfinancing.Service,
	personalInvoiceFinancingService personalinvoicefinancing.Service,
	personalCreditCardService personalcreditcard.Service,
	businessAccountService businessaccount.Service,
	businessLoanService businessloan.Service,
	businessFinancingService businessfinancing.Service,
	businessCreditCardService businesscreditcard.Service,
	crawler services.Crawler) InstitutionInterface {

	return &institutionInterface{
		institutionService:              institutionService,
		branchService:                   branchService,
		electronicChannelService:        electronicChannelService,
		personalAccountService:          personalAccountService,
		personalLoanService:             personalLoanService,
		personalFinancingService:        personalFinancingService,
		personalInvoiceFinancingService: personalInvoiceFinancingService,
		personalCreditCardService:       personalCreditCardService,
		businessAccountService:          businessAccountService,
		businessLoanService:             businessLoanService,
		businessFinancingService:        businessFinancingService,
		businessCreditCardService:       businessCreditCardService,
		crawler:                         crawler,
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

//UpdatePersonalAccounts update accounts from institution
func (i *institutionInterface) UpdatePersonalAccounts(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)
	if err != nil {
		return err
	}
	personalAccounts, crawlErr := i.crawler.PersonalAccounts(institution.BaseURL, 1, []personalaccount.Entity{})
	if crawlErr != nil {
		return crawlErr
	}
	delErr := i.personalAccountService.DeleteAllFromInstitution(id)
	if delErr != nil {
		return delErr
	}
	insertErr := i.personalAccountService.InsertMany(*personalAccounts, id)
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

//UpdatePersonalFinancings update personal financings from institution
func (i *institutionInterface) UpdatePersonalFinancings(id string) common.CustomError {
	institution, err := i.institutionService.Read(id)
	if err != nil {
		return err
	}
	personalFinancings, crawlErr := i.crawler.PersonalFinancings(institution.BaseURL, 1, []personalfinancing.Entity{})
	if crawlErr != nil {
		return crawlErr
	}
	delErr := i.personalFinancingService.DeleteAllFromInstitution(id)
	if delErr != nil {
		return delErr
	}
	insertErr := i.personalFinancingService.InsertMany(*personalFinancings, id)
	if insertErr != nil {
		return insertErr
	}
	return nil
}

//UpdatePersonalInvoiceFinancings update personal financings from institution
func (i *institutionInterface) UpdatePersonalInvoiceFinancings(id string) common.CustomError {
	institution, err := i.institutionService.Read(id)
	if err != nil {
		return err
	}
	personalInvoiceFinancings, crawlErr := i.crawler.PersonalInvoiceFinancings(institution.BaseURL, 1, []personalinvoicefinancing.Entity{})
	if crawlErr != nil {
		return crawlErr
	}
	delErr := i.personalInvoiceFinancingService.DeleteAllFromInstitution(id)
	if delErr != nil {
		return delErr
	}
	insertErr := i.personalInvoiceFinancingService.InsertMany(*personalInvoiceFinancings, id)
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

//UpdateBusinessAccounts update accounts from institution
func (i *institutionInterface) UpdateBusinessAccounts(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)
	if err != nil {
		return err
	}
	businessAccounts, crawlErr := i.crawler.BusinessAccounts(institution.BaseURL, 1, []businessaccount.Entity{})
	if crawlErr != nil {
		return crawlErr
	}
	delErr := i.businessAccountService.DeleteAllFromInstitution(id)
	if delErr != nil {
		return delErr
	}
	insertErr := i.businessAccountService.InsertMany(*businessAccounts, id)
	if insertErr != nil {
		return insertErr
	}
	return nil
}

//UpdateBusinessLoans update loans from institution
func (i *institutionInterface) UpdateBusinessLoans(id string) common.CustomError {

	institution, err := i.institutionService.Read(id)
	if err != nil {
		return err
	}
	businessLoans, crawlErr := i.crawler.BusinessLoans(institution.BaseURL, 1, []businessloan.Entity{})
	if crawlErr != nil {
		return crawlErr
	}
	delErr := i.businessLoanService.DeleteAllFromInstitution(id)
	if delErr != nil {
		return delErr
	}
	insertErr := i.businessLoanService.InsertMany(*businessLoans, id)
	if insertErr != nil {
		return insertErr
	}
	return nil
}

//UpdateBusinessFinancings update business financings from institution
func (i *institutionInterface) UpdateBusinessFinancings(id string) common.CustomError {
	institution, err := i.institutionService.Read(id)
	if err != nil {
		return err
	}
	businessFinancings, crawlErr := i.crawler.BusinessFinancings(institution.BaseURL, 1, []businessfinancing.Entity{})
	if crawlErr != nil {
		return crawlErr
	}
	delErr := i.businessFinancingService.DeleteAllFromInstitution(id)
	if delErr != nil {
		return delErr
	}
	insertErr := i.businessFinancingService.InsertMany(*businessFinancings, id)
	if insertErr != nil {
		return insertErr
	}
	return nil
}

//UpdateBusinessCreditCards update business credit cards from institution
func (i *institutionInterface) UpdateBusinessCreditCards(id string) common.CustomError {
	institution, err := i.institutionService.Read(id)
	if err != nil {
		return err
	}
	businessCreditCards, crawlErr := i.crawler.BusinessCreditCards(institution.BaseURL, 1, []businesscreditcard.Entity{})
	if crawlErr != nil {
		return crawlErr
	}
	delErr := i.businessCreditCardService.DeleteAllFromInstitution(id)
	if delErr != nil {
		return delErr
	}
	insertErr := i.businessCreditCardService.InsertMany(*businessCreditCards, id)
	if insertErr != nil {
		return insertErr
	}
	return nil
}
