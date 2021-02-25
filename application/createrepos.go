package application

import (
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/businessaccount"
	"openbankingcrawler/domain/businesscreditcard"
	"openbankingcrawler/domain/businessfinancing"
	"openbankingcrawler/domain/businessinvoicefinancing"
	"openbankingcrawler/domain/businessloan"
	"openbankingcrawler/domain/businessunarrangedaccountoverdraft"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalfinancing"
	"openbankingcrawler/domain/personalinvoicefinancing"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/domain/personalunarrangedaccountoverdraft"

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
	personalaccount.Repository, personalloan.Repository, personalfinancing.Repository,
	personalinvoicefinancing.Repository, personalcreditcard.Repository, personalunarrangedaccountoverdraft.Repository,
	businessaccount.Repository, businessloan.Repository, businessfinancing.Repository,
	businessinvoicefinancing.Repository, businesscreditcard.Repository, businessunarrangedaccountoverdraft.Repository) {

	personalLoanRepository := personalloan.NewRepository(connection.Collection("personalLoan"))
	personalAccountRepository := personalaccount.NewRepository(connection.Collection("personalAccount"))
	personalFinancingRepository := personalfinancing.NewRepository(connection.Collection("personalFinancing"))
	personalInvoiceFinancingRepository := personalinvoicefinancing.NewRepository(connection.Collection("personalInvoiceFinancing"))
	personalCreditCardRepository := personalcreditcard.NewRepository(connection.Collection("personalCreditCard"))
	personalUnarrangedAccountOverdraftsRepository := personalunarrangedaccountoverdraft.NewRepository(connection.Collection("personalUnarrangedAccountOverdrafts"))

	businessAccountRepository := businessaccount.NewRepository(connection.Collection("businessAccount"))
	businessLoanRepository := businessloan.NewRepository(connection.Collection("businessLoan"))
	businessFinancingRepository := businessfinancing.NewRepository(connection.Collection("businessFinancing"))
	businessInvoiceFinancingRepository := businessinvoicefinancing.NewRepository(connection.Collection("businessInvoiceFinancing"))
	businessCreditCardRepository := businesscreditcard.NewRepository(connection.Collection("businessCreditCard"))
	businessUnarrangedAccountOverdraftsRepository := businessunarrangedaccountoverdraft.NewRepository(connection.Collection("businessUnarrangedAccountOverdrafts"))

	return personalAccountRepository, personalLoanRepository, personalFinancingRepository,
		personalInvoiceFinancingRepository, personalCreditCardRepository, personalUnarrangedAccountOverdraftsRepository,
		businessAccountRepository, businessLoanRepository, businessFinancingRepository,
		businessInvoiceFinancingRepository, businessCreditCardRepository, businessUnarrangedAccountOverdraftsRepository
}

//CreateProductsServicesServices create products services services
func CreateProductsServicesServices(connection *bongo.Connection) (
	personalaccount.Service, personalloan.Service, personalfinancing.Service,
	personalinvoicefinancing.Service, personalcreditcard.Service, personalunarrangedaccountoverdraft.Service,
	businessaccount.Service, businessloan.Service, businessfinancing.Service,
	businessinvoicefinancing.Service, businesscreditcard.Service, businessunarrangedaccountoverdraft.Service) {

	personalAccountRepository,
		personalLoanRepository,
		personalFinancingRepository,
		personalInvoiceFinancingRepository,
		personalCreditCardRepository,
		personalUnarrangedAccountOverdraftsRepository,
		businessAccountRepository,
		businessLoanRepository,
		businessFinancingRepository,
		businessInvoiceFinancingRepository,
		businessCreditCardRepository,
		businessUnarrangedAccountOverdraftsRepository := createProductsServicesRepositories(connection)

	personalAccountService := personalaccount.NewService(personalAccountRepository)
	personalLoanService := personalloan.NewService(personalLoanRepository)
	personalFinanceService := personalfinancing.NewService(personalFinancingRepository)
	personalInvoiceFinanceService := personalinvoicefinancing.NewService(personalInvoiceFinancingRepository)
	personalCreditCardService := personalcreditcard.NewService(personalCreditCardRepository)
	personalUnarrangedAccountOverdraftService := personalunarrangedaccountoverdraft.NewService(personalUnarrangedAccountOverdraftsRepository)

	businessAccountService := businessaccount.NewService(businessAccountRepository)
	businessLoanService := businessloan.NewService(businessLoanRepository)
	businessFinancingService := businessfinancing.NewService(businessFinancingRepository)
	businessInvoiceFinancingService := businessinvoicefinancing.NewService(businessInvoiceFinancingRepository)
	businessCreditCardService := businesscreditcard.NewService(businessCreditCardRepository)
	businessUnarrangedAccountOverdraftService := businessunarrangedaccountoverdraft.NewService(businessUnarrangedAccountOverdraftsRepository)

	return personalAccountService, personalLoanService, personalFinanceService,
		personalInvoiceFinanceService, personalCreditCardService, personalUnarrangedAccountOverdraftService,
		businessAccountService, businessLoanService, businessFinancingService,
		businessInvoiceFinancingService, businessCreditCardService, businessUnarrangedAccountOverdraftService
}
