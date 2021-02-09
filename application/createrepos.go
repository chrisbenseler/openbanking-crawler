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

//CreateRepositories CreateRepositories
func CreateRepositories(connection *bongo.Connection) (institution.Repository, branch.Repository, electronicchannel.Repository, personalloan.Repository, personalcreditcard.Repository, personalaccount.Repository, businessaccount.Repository) {
	institutionRepository := institution.NewRepository(connection.Collection("institution"))
	branchRepository := branch.NewRepository(connection.Collection("branch"))
	electronicChannelRepository := electronicchannel.NewRepository(connection.Collection("electronicChannel"))
	personalLoanRepository := personalloan.NewRepository(connection.Collection("personalLoan"))
	personalCreditCardRepository := personalcreditcard.NewRepository(connection.Collection("personalCreditCard"))
	personalAccountRepository := personalaccount.NewRepository(connection.Collection("personalAccount"))

	businessAccountRepository := businessaccount.NewRepository(connection.Collection("businessAccount"))

	return institutionRepository, branchRepository, electronicChannelRepository, personalLoanRepository, personalCreditCardRepository, personalAccountRepository, businessAccountRepository
}

//CreateServices create all services
func CreateServices(connection *bongo.Connection) (
	institution.Service, branch.Service, electronicchannel.Service, personalloan.Service, personalcreditcard.Service,
	personalaccount.Service, businessaccount.Service) {

	institutionRepository, branchRepository,
		electronicChannelRepository,
		personalLoanRepository,
		personalCreditCardRepository,
		personalAccountRepository,
		businessAccountRepository := CreateRepositories(connection)

	institutionService := institution.NewService(institutionRepository)
	branchService := branch.NewService(branchRepository)
	electronicChannelService := electronicchannel.NewService(electronicChannelRepository)
	personalLoanService := personalloan.NewService(personalLoanRepository)
	personalCreditCardService := personalcreditcard.NewService(personalCreditCardRepository)
	personalAccountService := personalaccount.NewService(personalAccountRepository)
	businessAccountService := businessaccount.NewService(businessAccountRepository)

	return institutionService, branchService, electronicChannelService,
		personalLoanService, personalCreditCardService,
		personalAccountService, businessAccountService
}
