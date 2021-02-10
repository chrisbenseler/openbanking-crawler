package application

import (
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/businessaccount"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
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
func createProductsServicesRepositories(connection *bongo.Connection) (personalloan.Repository, personalcreditcard.Repository, personalaccount.Repository, businessaccount.Repository) {

	personalLoanRepository := personalloan.NewRepository(connection.Collection("personalLoan"))
	personalCreditCardRepository := personalcreditcard.NewRepository(connection.Collection("personalCreditCard"))
	personalAccountRepository := personalaccount.NewRepository(connection.Collection("personalAccount"))

	businessAccountRepository := businessaccount.NewRepository(connection.Collection("businessAccount"))

	return personalLoanRepository, personalCreditCardRepository, personalAccountRepository, businessAccountRepository
}

//CreateProductsServicesServices create products services services
func CreateProductsServicesServices(connection *bongo.Connection) (
	personalloan.Service, personalcreditcard.Service,
	personalaccount.Service, businessaccount.Service) {

	personalLoanRepository,
		personalCreditCardRepository,
		personalAccountRepository,
		businessAccountRepository := createProductsServicesRepositories(connection)

	personalLoanService := personalloan.NewService(personalLoanRepository)
	personalCreditCardService := personalcreditcard.NewService(personalCreditCardRepository)
	personalAccountService := personalaccount.NewService(personalAccountRepository)
	businessAccountService := businessaccount.NewService(businessAccountRepository)

	return personalLoanService, personalCreditCardService,
		personalAccountService, businessAccountService
}
