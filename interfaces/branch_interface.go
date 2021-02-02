package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/subentities"
)

//BranchInterface interface
type BranchInterface interface {
	GetFromInstitution(string, int) ([]branch.Entity, *subentities.Pagination, common.CustomError)
}

type branchInterface struct {
	branchService branch.Service
}

//NewBranch create a new interface for branches
func NewBranch(branchService branch.Service) BranchInterface {

	return &branchInterface{
		branchService: branchService,
	}
}

//Get get branches from institutution
func (b *branchInterface) GetFromInstitution(id string, page int) ([]branch.Entity, *subentities.Pagination, common.CustomError) {
	return b.branchService.FindByInstitution(id, page)
}
