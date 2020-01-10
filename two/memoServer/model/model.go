package model

import (
	"fmt"
	"log"

	pagination "github.com/jimweng/memoServer/model/paginator"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

/* TODO: add extra category handle for MemoList*/
// meomo column: Title; Description; Start Date; End Date; Category[Pesonal, Business, Others]
type MemoList struct {
	gorm.Model
	Title       string `gorm:"not full;size:50"`   //;unique"` // 50 chars & unique
	Description string `gorm:"not full;size:1000"` // 1000 chars
	Category    int    // 0[Others], 1[Personal], 2[Business]
}

func InitDB() (*gorm.DB, error) {
	var err error
	if db, err = gorm.Open("sqlite3", "example.db"); err != nil {
		return nil, err
	}
	db.AutoMigrate(&MemoList{})

	return db, err

}

/* TODO: add extra category handle for Category */
// InsertData would create an new record for the memo list
func InsertData(title string, description string, category int) error {
	return db.Create(&MemoList{Title: title, Description: description, Category: category}).Error
}

// GetData would return the paginator info
func GetData(page int, limit int, index string, order string) *pagination.Paginator {
	var memoss []MemoList

	return pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{index + " " + order},
		ShowSQL: true,
	}, &memoss)
}

/* TODO: need to add query result handle before update data*/
// DeleteData would delete the specify id data
func DeleteData(id int) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	dbSQL := tx.Where("id = ?", id).Delete(&MemoList{})
	if err := dbSQL.Error; err != nil {
		tx.Rollback()
		return err
	}
	if dbSQL.RowsAffected < 1 {
		tx.Rollback()
		return fmt.Errorf("Invalid operations.")
	}
	return tx.Commit().Error
}

// BatchDeleteData would delete the specify id data
func BatchDeleteData(ids *[]int) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, id := range *ids {
		dbSQL := tx.Where("id = ?", id).Delete(&MemoList{})
		if err := dbSQL.Error; err != nil {
			tx.Rollback()
			return err
		}

		if dbSQL.RowsAffected < 1 {
			tx.Rollback()
			return fmt.Errorf("Invalid operation.")
		}
	}
	return tx.Commit().Error
}

/* TODO: need to add query result handle before update data*/
// UpdateData would update origin description via id
func UpdateData(id int, title string, description string, category int) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	dbSQL := db.Model(&MemoList{}).Where("id = ?", id).Update(&MemoList{Title: title, Description: description, Category: category})
	if err := dbSQL.Error; err != nil {
		tx.Rollback()
		return err
	}
	if dbSQL.RowsAffected < 1 {
		tx.Rollback()
		return fmt.Errorf("Invalid operation.")
	}
	return tx.Commit().Error
}

// BatchUpdateData would update origin description via ids
func BatchUpdateData(data *[]MemoList) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, j := range *data {
		dbSQL := db.Model(&MemoList{}).Where("id = ?", j.ID).Update(&MemoList{Title: j.Title, Description: j.Description, Category: j.Category})
		if err := dbSQL.Error; err != nil {
			tx.Rollback()
			return err
		}
		if dbSQL.RowsAffected < 1 {
			tx.Rollback()
			return fmt.Errorf("Invalid operation.")
		}
	}

	return tx.Commit().Error
}

/* insert sample data */
func InsertSampleData() {
	db.AutoMigrate(&MemoList{})

	db.Create(&MemoList{Title: "TestTitle1", Description: "1001 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle2", Description: "1002 words included", Category: 1})
	db.Create(&MemoList{Title: "TestTitle3", Description: "1003 words included", Category: 2})
	db.Create(&MemoList{Title: "TestTitle4", Description: "1004 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle5", Description: "1005 words included", Category: 1})
	db.Create(&MemoList{Title: "TestTitle6", Description: "1006 words included", Category: 2})
	db.Create(&MemoList{Title: "TestTitle7", Description: "1007 words included", Category: 0})
	db.Create(&MemoList{Title: "TestTitle8", Description: "1008 words included", Category: 1})
	db.Create(&MemoList{Title: "TestTitle9", Description: "1009 words included", Category: 2})
	log.Println("Insert OK!")
}
