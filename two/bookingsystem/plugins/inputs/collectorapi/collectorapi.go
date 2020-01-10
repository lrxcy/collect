package collectorapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jimweng/bookingsystem/plugins/inputs"
	"github.com/jimweng/bookingsystem/utils"
)

type RespStruct struct {
	Url        string
	TempString tmpStr
	Client     *utils.Client
}

type tmpStr func(string, string) string

var f = func(st string, et string) string {
	return fmt.Sprintf("StartTime=%v&EndTime=%v&WagersId=&IsUpdateTime=%v&key=%v", st, et, true, os.Getenv("signKey"))
}

func (r *RespStruct) Gather() error {
	var st, et string
	st = convertUnixUTCTimeStamp(startTime)
	et = convertUnixUTCTimeStamp(endTime)

	payloadbytes, _ := json.Marshal(map[string]interface{}{
		"StartTime":    st,
		"EndTime":      et,
		"IsUpdateTime": true,
		"Sign":         utils.GetMD5Hash(r.TempString(st, et)),
	})

	resp, err := r.Client.Post(r.Url, bytes.NewReader(payloadbytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if bodyBytes, err := ioutil.ReadAll(resp.Body); err != nil {
		return err

	} else {
		go writeToChannel(string(bodyBytes), inputs.ResultChannel)
	}

	return nil
}

var (
	// writeFlowChannel = make(chan string)
	endTime   = time.Now().Unix()
	startTime = endTime - 300 // 取的時間為現在時間回推5分鐘

	queryInterval = 60 // 可以跟定時收集的次數同樣時區，但是不能小於。否則會有收集不到數據的問題
)

func writeToChannel(s string, c chan string) {
	c <- s
}

// settleQueryTime 將原本初始化給出的時間向前推移
func settleQueryTime() (string, string) {
	endTime = endTime + int64(queryInterval)
	startTime = endTime - 300

	return fmt.Sprintf("%v", startTime), fmt.Sprintf("%v", endTime)
}

// convertUnixUTCTimeStamp 將unixtime轉換為可以送出的時間格式
func convertUnixUTCTimeStamp(i int64) string {
	return fmt.Sprintf(time.Unix(i, 0).UTC().Format("2006/01/02 15:04:05Z"))
}

func init() {
	inputs.AddInputPlugin("collectorapi", func() inputs.InputPlugin {
		return &RespStruct{
			Url:        os.Getenv("requestUrl"),
			Client:     utils.RetriveHttpClient(),
			TempString: f,
		}
	})
}
