package collectorapi

import (
	"log"
	"os"
	"testing"

	"github.com/jimweng/bookingsystem/conf"
	. "github.com/jimweng/bookingsystem/logger"
	"github.com/jimweng/bookingsystem/plugins/inputs"
	"github.com/jimweng/bookingsystem/utils"
	"github.com/stretchr/testify/assert"
)

func TestCollectorApi(t *testing.T) {
	// initialize config
	configPath := "../../../conf/app.dev.ini"
	config, _ := conf.InitConfig(&configPath)
	log.Println(config)

	InitLog("./test.logs", "info")

	a := RespStruct{
		Url:        os.Getenv("requestUrl"),
		Client:     utils.RetriveHttpClient(),
		TempString: f,
	}

	err := a.Gather()
	assert.NoError(t, err)

	// if timeout this would stuck
	if err == nil {
		select {
		case c1 := <-inputs.ResultChannel:
			assert.Equal(t, "", c1)
		}
	}
}
