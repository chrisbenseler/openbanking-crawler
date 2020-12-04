package interfaces

import (
	"openbankingcrawler/domain/branch"
)

//BranchInterface interface
type BranchInterface interface {
	GetFromInstitution(string) []branch.Entity
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
func (b *branchInterface) GetFromInstitution(id string) []branch.Entity {
	return b.branchService.FindByInstitution(id)
}
