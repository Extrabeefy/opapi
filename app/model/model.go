package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Progression struct {
	gorm.Model
	Bossname       string `gorm:"unique" json:"bossname"`
	RaidFilter     string `json:"raidfilter"`
	ImgPath        string `json:"imgpath"`
	KillOrder      string `json:"killorder"`
	CreatedBy      string `json:"createdby"`
	KillDifficulty string `json:"killdifficulty"`
	Dead           bool   `json:"dead"`
}

func (p *Progression) Killed() {
	p.Dead = true
}

func (p *Progression) Revive() {
	p.Dead = false
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Progression{})
	return db
}
