package link

import (
	"go_pro_api/pkg/db"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	DataBase *db.DB
}

func NewRepository(db *db.DB) *LinkRepository {
	return &LinkRepository{
		DataBase: db,
	}
}

func (r *LinkRepository) Create(link *Link) (*Link, error) {
	result := r.DataBase.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (r *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := r.DataBase.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (r *LinkRepository) Update(link *Link) (*Link, error) {
	result := r.DataBase.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (r *LinkRepository) Delete(id uint) error {
	result := r.DataBase.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link
	result := r.DataBase.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (r *LinkRepository) Count() int64 {
	var count int64
	r.DataBase.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	return count
}

func (r *LinkRepository) GetAll(limit, offset int) []Link {
	var links []Link
	r.DataBase.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	return links
}
