package mysqldb

import (
	"time"

	"github.com/fatih/structs"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	//`gorm:"column:id;AUTO_INCREMENT"` : 不保證ID非空，也不能確定id為唯一值
	ID        int    `gorm:"column:id;unsigned AUTO_INCREMENT;not null;primary_key"`
	Name      string `gor:"column:name";not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID        int    `gorm:"column:id;unsigned AUTO_INCREMENT;not null; primary_key"`
	Name      string `gorm:"column:name;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Users []User `gorm:"many2many:user_languages;"`
}

// CreateRecord is based on usr with multi-languages
func (db *mysqlDBObj) CreateRecord(usrname string, lang ...string) error {
	// use id ad primary key instead of user name
	var id int
	row := db.DB.Table("user").Where("name = ?", usrname).Select("id").Row()
	row.Scan(&id)

	user := User{
		ID:   id,
		Name: usrname,
	}

	language := make([]Language, len(lang))

	for i, j := range lang {

		// 修正: 如果不加這行，他會自動增加id導致會有重複的name出現;而非使用既有的ID...因為name不是primary_key
		var id int
		row := db.DB.Table("language").Where("name = ?", j).Select("id").Row()
		row.Scan(&id)

		language[i] = Language{ID: id, Name: j}
	}

	user.Languages = language
	return db.DB.Create(&user).Error
}

// QueryRecord is based on usr with multi-languages
func (db *mysqlDBObj) QueryRecord(usrname string) (*[]map[string]interface{}, error) {
	var languageIds []int

	var userId int
	row := db.DB.Table("user").Where("name = ?", usrname).Select("id").Row()
	row.Scan(&userId)

	relateRow, err := db.DB.Table("user_languages").Where("user_id = ?", userId).Select("language_id").Rows()
	if err != nil {
		return nil, err
	}

	defer relateRow.Close()
	for relateRow.Next() {
		var languageId int
		relateRow.Scan(&languageId)
		languageIds = append(languageIds, languageId)
	}

	languagesMap := make([]map[string]interface{}, 0)
	for _, j := range languageIds {
		var languageName string
		tmpRow := db.DB.Table("language").Where("id = ?", j).Select("name").Row()
		tmpRow.Scan(&languageName)

		tmpLanguage := &Language{ID: j, Name: languageName}
		languagesMap = append(languagesMap, structs.Map(tmpLanguage))
	}
	return &languagesMap, nil
}

func init() {
	tables = append(tables, &Language{})
}

// /*
//	TODO:
// 	Replace language implement with language interface{}
// */

// type LanguageImp interface {
// 	GetRawData() ([]byte, error)
// }

// func (l *Language) GetRawData() ([]byte, error) {
// 	return json.Marshal(l)
// }
