package mysqldb

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigration(t *testing.T) {
	// os.Setenv("TestDB", "test")
	LoadMySQLDBConfig("mysql", "127.0.0.1", "3306", "root", "secret", true, 10, 10)
	StartMySQLDB()

	r := RetriveMySQLDBAccessObj()
	r.migration(&DataStruct{})
}

func TestParser(t *testing.T) {

	d := &DummyStruct{}
	json.Unmarshal([]byte(DummpyTest), d)
	assert.Equal(t, 0, d.Code)
	assert.Equal(t, "Success", d.Message)
	for i, j := range d.Data {
		switch i {
		case 0:
			assert.Equal(t, "65", j.AgentID)
		default:
			continue
		}
	}

	// os.Setenv("TestDB", "test")
	LoadMySQLDBConfig("mysql", "127.0.0.1", "3306", "root", "secret", true, 10, 10)
	StartMySQLDB()

	r := RetriveMySQLDBAccessObj()
	r.CreateGameRecords(DummpyTest)
}

var (
	DummpyTest = `
	{
		"Code": 0,
		"Message": "Success",
		"Data": [
		  {
			"AgentId": "65",
			"UserAccount": "test",
			"GameId": "383",
			"WagersId": "10_25909_50-1575880311-44580244-2",
			"GameAccount": "gm1_00000000k2",
			"GameWagersId": "50-1575880311-44580244-2",
			"Bet": 5.00,
			"ValidBet": 2.85,
			"PayOff": 2.85,
			"BetTime": "2019-12-09T08:32:02Z",
			"BalanceTime": "2019-12-09T08:32:02Z",
			"GameGroupType": 3,
			"UpdateTime": "2019-12-09T08:34:35.619Z"
		  },
		  {
			"AgentId": "65",
			"UserAccount": "test",
			"GameId": "383",
			"WagersId": "10_25909_50-1575880311-44580244-2",
			"GameAccount": "gm1_00000000k2",
			"GameWagersId": "50-1575880311-44580244-2",
			"Bet": 5.00,
			"ValidBet": 2.85,
			"PayOff": 2.85,
			"BetTime": "2019-12-09T08:32:02Z",
			"BalanceTime": "2019-12-09T08:32:02Z",
			"GameGroupType": 3,
			"UpdateTime": "2019-12-09T08:34:55.624Z"
		  },
		  {
			"AgentId": "65",
			"UserAccount": "test",
			"GameId": "383",
			"WagersId": "10_25909_50-1575880290-44580104-2",
			"GameAccount": "gm1_00000000k2",
			"GameWagersId": "50-1575880290-44580104-2",
			"Bet": 15.00,
			"ValidBet": 2.85,
			"PayOff": 2.85,
			"BetTime": "2019-12-09T08:31:43Z",
			"BalanceTime": "2019-12-09T08:31:43Z",
			"GameGroupType": 3,
			"UpdateTime": "2019-12-09T08:33:40.659Z"
		  },
		  {
			"AgentId": "65",
			"UserAccount": "test",
			"GameId": "383",
			"WagersId": "10_25909_50-1575880242-44579773-3",
			"GameAccount": "gm1_00000000k2",
			"GameWagersId": "50-1575880242-44579773-3",
			"Bet": 5.00,
			"ValidBet": 5.00,
			"PayOff": -5.00,
			"BetTime": "2019-12-09T08:31:20Z",
			"BalanceTime": "2019-12-09T08:31:20Z",
			"GameGroupType": 3,
			"UpdateTime": "2019-12-09T08:33:40.661Z"
		  },
		  {
			"AgentId": "65",
			"UserAccount": "test",
			"GameId": "383",
			"WagersId": "10_25909_50-1575880290-44580104-2",
			"GameAccount": "gm1_00000000k2",
			"GameWagersId": "50-1575880290-44580104-2",
			"Bet": 15.00,
			"ValidBet": 2.85,
			"PayOff": 2.85,
			"BetTime": "2019-12-09T08:31:43Z",
			"BalanceTime": "2019-12-09T08:31:43Z",
			"GameGroupType": 3,
			"UpdateTime": "2019-12-09T08:33:20.625Z"
		  },
		  {
			"AgentId": "65",
			"UserAccount": "test",
			"GameId": "383",
			"WagersId": "10_25909_50-1575880242-44579773-3",
			"GameAccount": "gm1_00000000k2",
			"GameWagersId": "50-1575880242-44579773-3",
			"Bet": 5.00,
			"ValidBet": 5.00,
			"PayOff": -5.00,
			"BetTime": "2019-12-09T08:31:20Z",
			"BalanceTime": "2019-12-09T08:31:20Z",
			"GameGroupType": 3,
			"UpdateTime": "2019-12-09T08:33:20.645Z"
		  }
		]
	  }
	`
)
