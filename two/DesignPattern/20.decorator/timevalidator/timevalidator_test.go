package timevalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeValidator(t *testing.T) {
	var c = &Info{
		Name:     "this is test name",
		Rep:      true,
		MinNum:   1,
		MaxNum:   6,
		ISize:    7,
		PSize:    3,
		IType:    IType1,
		DayTNo:   96,
		Frenc:    "15m",
		LType:    LType1,
		SortType: STypeDec,
		PRule: &Perid{
			BTime:  "2019-10-21 19:15:00",
			BValue: 9999,
			Frenc:  "15m",
			INum:   96,
		},
	}
	cc := WrapValidator(c)
	res := cc.checkformat()
	assert.Equal(t, true, res)
}

type timeForamter struct {
	timeStr string
}

type Info struct {
	Name     string
	Rep      bool
	SortType int
	MinNum   int
	MaxNum   int
	ISize    int
	PSize    int
	IType    int
	DayTNo   int
	Frenc    string
	LType    int
	PRule    *Perid
}

type Perid struct {
	TimePerid string
	Frenc     string
	INum      int
	BTime     string
	BValue    int64
}

func (i *Info) checkformat() bool {
	return i.Rep
}

func (i *Info) getIType() int {
	return i.IType
}

func (i *Info) getLType() int {
	return i.LType
}

func (i *Info) getSType() int {
	return i.SortType
}

func (i *Info) getIssue() int {
	return i.PRule.INum
}

func TestBaseValueIssueNo(t *testing.T) {
	str, err := baseValueIssueNo("15m", "2019-10-21 19:15:00", 96, 9999)
	assert.Nil(t, err)
	assert.Equal(t, "10001", str)

	assert.Equal(t, 5, len(str))
}
