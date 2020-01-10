package mysqlsender

import (
	"github.com/jimweng/bookingsystem/models"
	"github.com/jimweng/bookingsystem/plugins/outputs"
)

type WriterStruct struct {
	MySqlClient models.MySqlImplement
}

func (w *WriterStruct) Write(records string) error {
	return w.MySqlClient.CreateGameRecords(records)
}

func init() {
	outputs.AddOutputPlugin("mysqlsender", func() outputs.OutputPlugin {
		return &WriterStruct{
			MySqlClient: models.RetriveMySqlDbAccessModel(),
		}
	})
}
