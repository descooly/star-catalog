package service

// для базы данных
// Star — основная сущность звезды
type Star struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

// StarInfo — дополнительная информация о звезде
type StarInfo struct {
	StarID        string  `gorm:"primaryKey" json:"starID"` // будет внешним ключом к Star.ID
	Constellation string  `json:"constellation"`
	Distance      float64 `json:"distance"`
	Mass          float64 `json:"mass"`
}

// для принятия запроса и для отправки на сайт
type ResultStar struct {
	Name   string  `json:"name"`
	Constl string  `json:"constellation"`
	Dist   float64 `json:"distance"`
	Mass   float64 `json:"mass"`
}
