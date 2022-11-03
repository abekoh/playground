package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	lang := Language{
		Code: "ja",
		Name: "Japanese",
	}

	DB.Create(&lang)

	var cnt int64
	if err := DB.Model(&Language{Code: "ja"}).Count(&cnt).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	if cnt != 1 {
		t.Error("count must be 1")
	}
}
