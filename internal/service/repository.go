package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetDB(sn, cl string, minDist, maxDist, minMass, maxMass *float64) ([]resStruct, error)
	PostDB(resStar *ResultStar) (ResultStar, error)
}

type resStruct struct {
	StarID        string  `gorm:"column:star_id"`
	Name          string  `gorm:"column:name"`
	Constellation string  `gorm:"column:constellation"`
	Distance      float64 `gorm:"column:distance"`
	Mass          float64 `gorm:"column:mass"`
}

type Repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{database: db}
}

func (r *Repo) GetDB(sn, cl string, minDist, maxDist, minMass, maxMass *float64) ([]resStruct, error) {
	var results []resStruct

	// Начинаем запрос
	db := r.database.Model(&StarInfo{})

	// Джойним с Star, чтобы получить имя
	db = db.Joins("JOIN stars ON stars.id = star_infos.star_id")

	// Фильтрация по name (из таблицы stars)
	if sn != "" {
		db = db.Where("stars.name ILIKE ?", "%"+sn+"%") // ILIKE — case-insensitive в PostgreSQL
	}

	// Фильтрация по constellation
	if cl != "" {
		db = db.Where("star_infos.constellation ILIKE ?", "%"+cl+"%")
	}

	// Фильтрация по distance
	if minDist != nil {
		db = db.Where("star_infos.distance >= ?", *minDist)
	}
	if maxDist != nil {
		db = db.Where("star_infos.distance <= ?", *maxDist)
	}

	// Фильтрация по mass
	if minMass != nil {
		db = db.Where("star_infos.mass >= ?", *minMass)
	}
	if maxMass != nil {
		db = db.Where("star_infos.mass <= ?", *maxMass)
	}

	if err := db.Select("stars.id as star_id, stars.name, star_infos.constellation, star_infos.distance, star_infos.mass").
		Find(&results).Error; err != nil {
		return []resStruct{}, err
	}

	return results, nil
}

func (r *Repo) PostDB(resStar *ResultStar) (ResultStar, error) {
	newID := uuid.NewString()
	res := Star{
		ID:   newID,
		Name: (*resStar).Name,
	}
	res1 := StarInfo{
		StarID:        res.ID,
		Constellation: (*resStar).Constl,
		Distance:      (*resStar).Dist,
		Mass:          (*resStar).Mass,
	}

	if err := r.database.Create(&res).Error; err != nil {
		return ResultStar{}, err

	}
	if err := r.database.Create(&res1).Error; err != nil {
		return ResultStar{}, err
	}

	return *resStar, nil
}
