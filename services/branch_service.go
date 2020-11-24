package services

import (
	"encoding/json"
	"io/ioutil"
	"openbankingcrawler/domain/branch"
	"os"
)

//BranchService service
type BranchService interface {
	UpdateAll(string) error
	Crawl(string) (*[]branch.Entity, error)
	DeleteAllFromInstitution(string) error
}

type branchService struct {
	repository branch.Repository
}

//NewBranch create a new service for branches
func NewBranch(repository branch.Repository) BranchService {

	return &branchService{
		repository: repository,
	}
}

type branchJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				Branches []branch.Entity `json:"branches"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}

type branchesList struct {
	Branches []branch.Entity
}

//UpdateAll update branches from institution
func (s *branchService) UpdateAll(InstitutionID string) error {

	branches, crawlErr := s.Crawl(InstitutionID)

	if crawlErr != nil {
		return crawlErr
	}

	deleteErr := s.repository.DeleteMany(InstitutionID)

	if deleteErr != nil {
		return deleteErr
	}

	for _, branch := range *branches {
		branch.InstitutionID = InstitutionID
		s.repository.Save(branch)
	}

	return nil
}

//DeleteAllFromInstitution update branches from institution
func (s *branchService) DeleteAllFromInstitution(InstitutionID string) error {

	deleteErr := s.repository.DeleteMany(InstitutionID)

	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

//Crawl crawl branches from institution
func (s *branchService) Crawl(InstitutionID string) (*[]branch.Entity, error) {

	jsonFile, err := os.Open("./domain/branch/branches.json")

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	branchJSONData := &branchJSON{}

	jsonUnmarshallErr := json.Unmarshal(byteValue, &branchJSONData)

	if jsonUnmarshallErr != nil {
		return nil, jsonUnmarshallErr
	}

	companies := branchJSONData.Data.Brand.Companies[0]

	return &companies.Branches, nil

}
