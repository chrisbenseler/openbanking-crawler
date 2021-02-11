package personalinvoicefinancing

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/subentities"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

//Repository interface
type Repository interface {
	Save(Entity) error
	DeleteMany(string) error
	FindByInstitution(string, int) ([]Entity, *subentities.Pagination, common.CustomError)
}

type personalInvoiceFinancing struct {
	dao *bongo.Collection
}

//NewRepository create a new repository for personalLoan
func NewRepository(dao *bongo.Collection) Repository {
	return &personalInvoiceFinancing{
		dao: dao,
	}
}

//Save save an entity
func (r *personalInvoiceFinancing) Save(entity Entity) error {
	return r.dao.Save(&entity)
}

//DeleteMany delete all branches from an institution
func (r *personalInvoiceFinancing) DeleteMany(institutionID string) error {
	_, err := r.dao.Delete(bson.M{"institutionid": institutionID})
	return err
}

//FindByInstitution find all personalLoan from an institution
func (r *personalInvoiceFinancing) FindByInstitution(institutionID string, page int) ([]Entity, *subentities.Pagination, common.CustomError) {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	info, _ := results.Paginate(25, page)

	if results.Error != nil {
		return nil, nil, common.NewInternalServerError("Error on database", results.Error)
	}

	personalInvoices := make([]Entity, info.RecordsOnPage)

	for i := 0; i < info.RecordsOnPage; i++ {
		_ = results.Next(&personalInvoices[i])
	}

	pagination := subentities.Pagination{Total: info.TotalPages, Current: info.Current}

	return personalInvoices, &pagination, nil
}
