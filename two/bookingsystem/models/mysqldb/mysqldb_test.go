package mysqldb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqlDBAccessObj(t *testing.T) {
	ins1 := RetriveMySQLDBAccessObj()
	ins2 := RetriveMySQLDBAccessObj()
	assert.Equal(t, ins1, ins2)
}
