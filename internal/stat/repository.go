package stat

import (
	"go_pro_api/pkg/db"
	"gorm.io/datatypes"
	"time"
)

type StatRepository struct {
	DB *db.DB
}

func NewStatRepository(db *db.DB) *StatRepository {
	return &StatRepository{
		DB: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	repo.DB.Find(&stat, "link_id = ? and date = ?", linkId, datatypes.Date(time.Now()))
	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkId: linkId,
			Date:   datatypes.Date(time.Now()),
			Clicks: 1,
		})
	} else {
		stat.Clicks += 1
		repo.DB.Save(&stat)
	}

}

func (repo StatRepository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	repo.DB.Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}
