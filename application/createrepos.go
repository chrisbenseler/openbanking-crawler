package application

import (
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

	"github.com/go-bongo/bongo"
)

//createRepositories createRepositories
func createRepositories(connection *bongo.Connection) (institution.Repository, branch.Repository, electronicchannel.Repository) {
	institutionRepository := institution.NewRepository(connection.Collection("institution"))
	branchRepository := branch.NewRepository(connection.Collection("branch"))
	electronicChannelRepository := electronicchannel.NewRepository(connection.Collection("electronicChannel"))

	return institutionRepository, branchRepository, electronicChannelRepository
}

//CreateBasicServices create all services
func CreateBasicServices(connection *bongo.Connection) (
	institution.Service, branch.Service, electronicchannel.Service) {

	institutionRepository, branchRepository,
		electronicChannelRepository := createRepositories(connection)

	institutionService := institution.NewService(institutionRepository)
	branchService := branch.NewService(branchRepository)
	electronicChannelService := electronicchannel.NewService(electronicChannelRepository)

	return institutionService, branchService, electronicChannelService
}

//createProductsServicesRepositories create ProductsServices Repositories
func createProductsServicesRepositories(connection *bongo.Connection) (
	personalaccount.Repository, personalloan.Repository, personalfinancing.Repository, personalinvoicefinancing.Repository, personalcreditcard.Repository,
	businessaccount.Repository, businessloan.Repository, businessfinancing.Repository, businesscreditcard.Repository) {

	personalLoanRepository := personalloan.NewRepository(connection.Collection("personalLoan"))
	personalAccountRepository := personalaccount.NewRepository(connection.Collection("personalAccount"))
	personalFinancingRepository := personalfinancing.NewRepository(connection.Collection("personalFinancing"))
	personalInvoiceFinancingRepository := personalinvoicefinancing.NewRepository(connection.Collection("personalInvoiceFinancing"))
	personalCreditCardRepository := personalcreditcard.NewRepository(connection.Collection("personalCreditCard"))

	businessAccountRepository := businessaccount.NewRepository(connection.Collection("businessAccount"))
	businessLoanRepository := businessloan.NewRepository(connection.Collection("businessLoan"))
	businessFinancingRepository := businessfinancing.NewRepository(connection.Collection("businessFinancing"))
	businessCreditCardRepository := businesscreditcard.NewRepository(connection.Collection("businessCreditCard"))

	return personalAccountRepository, personalLoanRepository, personalFinancingRepository, personalInvoiceFinancingRepository, personalCreditCardRepository,
		businessAccountRepository, businessLoanRepository, businessFinancingRepository, businessCreditCardRepository
}

//CreateProductsServicesServices create products services services
func CreateProductsServicesServices(connection *bongo.Connection) (
	personalaccount.Service, personalloan.Service, personalfinancing.Service, personalinvoicefinancing.Service, personalcreditcard.Service,
	businessaccount.Service, businessloan.Service, businessfinancing.Service, businesscreditcard.Service) {

	personalAccountRepository,
		personalLoanRepository,
		personalFinancingRepository,
		personalInvoiceFinancingRepository,
		personalCreditCardRepository,
		businessAccountRepository,
		businessLoanRepository,
		businessFinancingRepository,
		businessCreditCardRepository := createProductsServicesRepositories(connection)

	personalAccountService := personalaccount.NewService(personalAccountRepository)
	personalLoanService := personalloan.NewService(personalLoanRepository)
	personalFinanceService := personalfinancing.NewService(personalFinancingRepository)
	personalInvoiceFinanceService := personalinvoicefinancing.NewService(personalInvoiceFinancingRepository)
	personalCreditCardService := personalcreditcard.NewService(personalCreditCardRepository)

	businessAccountService := businessaccount.NewService(businessAccountRepository)
	businessLoanService := businessloan.NewService(businessLoanRepository)
	businessFinancingService := businessfinancing.NewService(businessFinancingRepository)
	businessCreditCardService := businesscreditcard.NewService(businessCreditCardRepository)

	return personalAccountService, personalLoanService, personalFinanceService, personalInvoiceFinanceService, personalCreditCardService,
		businessAccountService, businessLoanService, businessFinancingService, businessCreditCardService
}

//businessCreditCardService businesscreditcard.Service
