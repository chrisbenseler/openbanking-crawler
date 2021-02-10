package application

import (
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/businessaccount"
	"openbankingcrawler/domain/businessfinancing"
	"openbankingcrawler/domain/businessloan"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalfinancing"
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
	personalaccount.Repository, personalloan.Repository, personalfinancing.Repository, personalcreditcard.Repository,
	businessaccount.Repository, businessloan.Repository, businessfinancing.Repository) {

	personalLoanRepository := personalloan.NewRepository(connection.Collection("personalLoan"))
	personalCreditCardRepository := personalcreditcard.NewRepository(connection.Collection("personalCreditCard"))
	personalAccountRepository := personalaccount.NewRepository(connection.Collection("personalAccount"))
	personalFinancingRepository := personalfinancing.NewRepository(connection.Collection("personalFinancing"))

	businessAccountRepository := businessaccount.NewRepository(connection.Collection("businessAccount"))
	businessLoanRepository := businessloan.NewRepository(connection.Collection("businessLoan"))
	businessFinancingRepository := businessfinancing.NewRepository(connection.Collection("businessFinancing"))

	return personalAccountRepository, personalLoanRepository, personalFinancingRepository, personalCreditCardRepository,
		businessAccountRepository, businessLoanRepository, businessFinancingRepository
}

//CreateProductsServicesServices create products services services
func CreateProductsServicesServices(connection *bongo.Connection) (
	personalaccount.Service, personalloan.Service, personalfinancing.Service, personalcreditcard.Service,
	businessaccount.Service, businessloan.Service, businessfinancing.Service) {

	personalAccountRepository,
		personalLoanRepository,
		personalFinancingRepository,
		personalCreditCardRepository,
		businessAccountRepository,
		businessLoanRepository,
		businessFinancingRepository := createProductsServicesRepositories(connection)

	personalLoanService := personalloan.NewService(personalLoanRepository)
	personalCreditCardService := personalcreditcard.NewService(personalCreditCardRepository)
	personalAccountService := personalaccount.NewService(personalAccountRepository)
	personalFinanceService := personalfinancing.NewService(personalFinancingRepository)

	businessAccountService := businessaccount.NewService(businessAccountRepository)
	businessLoanService := businessloan.NewService(businessLoanRepository)
	businessFinancingService := businessfinancing.NewService(businessFinancingRepository)

	return personalAccountService, personalLoanService, personalFinanceService, personalCreditCardService,
		businessAccountService, businessLoanService, businessFinancingService
}
