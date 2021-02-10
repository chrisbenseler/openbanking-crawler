package application

import (
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/businessaccount"
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
	personalaccount.Repository, personalloan.Repository, personalfinancing.Repository, personalcreditcard.Repository, businessaccount.Repository) {

	personalLoanRepository := personalloan.NewRepository(connection.Collection("personalLoan"))
	personalCreditCardRepository := personalcreditcard.NewRepository(connection.Collection("personalCreditCard"))
	personalAccountRepository := personalaccount.NewRepository(connection.Collection("personalAccount"))
	personalFinancingRepository := personalfinancing.NewRepository(connection.Collection("personalFinancing"))

	businessAccountRepository := businessaccount.NewRepository(connection.Collection("businessAccount"))

	return personalAccountRepository, personalLoanRepository, personalFinancingRepository, personalCreditCardRepository, businessAccountRepository
}

//CreateProductsServicesServices create products services services
func CreateProductsServicesServices(connection *bongo.Connection) (
	personalaccount.Service, personalloan.Service, personalfinancing.Service, personalcreditcard.Service,
	businessaccount.Service) {

	personalAccountRepository,
		personalLoanRepository,
		personalFinancingRepository,
		personalCreditCardRepository,
		businessAccountRepository := createProductsServicesRepositories(connection)

	personalLoanService := personalloan.NewService(personalLoanRepository)
	personalCreditCardService := personalcreditcard.NewService(personalCreditCardRepository)
	personalAccountService := personalaccount.NewService(personalAccountRepository)
	businessAccountService := businessaccount.NewService(businessAccountRepository)
	personalFinanceService := personalfinancing.NewService(personalFinancingRepository)

	return personalAccountService, personalLoanService, personalFinanceService, personalCreditCardService,
		businessAccountService
}
