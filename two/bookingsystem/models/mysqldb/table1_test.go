package mysqldb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	os.Setenv("TestDB", "test")
	LoadMySQLDBConfig("mysql", "127.0.0.1", "3306", "root", "secret", true, 10, 10)
	StartMySQLDB()
	migrationA()

	ins := RetriveMySQLDBAccessObj()

	assert.NoError(t, ins.CreateRecord("tester1", "english", "madarian", "spanish"))

	assert.NoError(t, ins.CreateRecord("tester2", "english", "chinese", "spanish"))

	assert.NoError(t, ins.CreateRecord("tester3", "english", "chinese", "spanish"))

	// it should be update instead of create whereas create would cause duplicate entry for key primary ...
	assert.Error(t, ins.CreateRecord("tester2", "history"))

	langArr, err := ins.QueryRecord("tester2")
	assert.NoError(t, err)
	// assert.Equal(t, "", *langArr)
	for _, j := range *langArr {
		assert.NotEqual(t, "", j["Name"])
	}
}
