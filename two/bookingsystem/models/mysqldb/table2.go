package mysqldb

import (
	"encoding/json"
	"log"
	"time"
)

type DummyStruct struct {
	Code    int          `json:"Code"`
	Message string       `json:"Message"`
	Data    []DataStruct `json:"Data"`
}

// TODO: rename meta from `json` to `gorm`
type DataStruct struct {
	AgentID       string    `json:"AgentId"`
	UserAccount   string    `json:"UserAccount"`
	GameID        string    `json:"GameId"`
	WagersID      string    `gorm:"primary_key";json:"WagersId"`
	GameAccount   string    `json:"GameAccount"`
	GameWagersID  string    `json:"GameWagersId"`
	Bet           float64   `json:"Bet"`
	ValidBet      float64   `json:"ValidBet"`
	PayOff        float64   `json:"PayOff"`
	BetTime       time.Time `json:"BetTime"`
	BalanceTime   time.Time `json:"BalanceTime"`
	GameGroupType int       `json:"GameGroupType"`
	UpdateTime    time.Time `json:"UpdateTime"`
}

/*
	refer : convert map to struct
	https://stackoverflow.com/questions/26744873/converting-map-to-struct/26746461
*/

// CreateGameRecord 接收字串後再做json parse，將字串轉換為陣列
func (db *mysqlDBObj) CreateGameRecords(data string) error {
	dummystruct := &DummyStruct{}

	if err := json.Unmarshal([]byte(data), dummystruct); err != nil {
		return err
	}

	for _, j := range dummystruct.Data {
		if err := db.CreateGameRecord(&j); err != nil {
			return err
		}
	}
	return nil
}

func (db *mysqlDBObj) CreateGameRecord(ds *DataStruct) error {
	var recordUpdateTime time.Time
	row := db.DB.Table("data_struct").Where("wagers_id = ?", ds.WagersID).Select("update_time").Row()
	row.Scan(&recordUpdateTime)

	if recordUpdateTime.IsZero() {
		log.Printf("create new record %v ____ with time %v\n", ds.WagersID, ds.UpdateTime)
		return db.DB.Create(ds).Error

	} else {
		if afterTimeSpan(recordUpdateTime, ds.UpdateTime) {
			// 如果wagesid存在，而且UpdateTime又比較新。就更新，否則跳過
			log.Printf("Update wagerid %v ... with time %v\n", ds.WagersID, ds.UpdateTime)
			return db.DB.Model(&DataStruct{}).Where("wagers_id = ?", ds.WagersID).Update("update_time", ds.UpdateTime).Error
		}
		return nil
	}
}

// start time 表示比較基準/ check time 是被比較時間
func afterTimeSpan(start, check time.Time) bool {
	return check.After(start)
}

func init() {
	tables = append(tables, &DummyStruct{})
}
