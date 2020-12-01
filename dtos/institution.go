package dtos

import "openbankingcrawler/common"

//Institution institution DTO
type Institution struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Validate validate an institution DTO
func (i *Institution) Validate() common.CustomError {
	if len(i.Name) == 0 {
		return common.NewUnprocessableEntity("No name found")
	}

	return nil
}
