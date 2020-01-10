package getproxyip

import (
	"log"

	"github.com/Aiicy/htmlquery"
)

//PLP get ip from proxylistplus.com
func PLP() []*IP {
	result := []*IP{}
	pollURL := "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1"
	doc, _ := htmlquery.LoadURL(pollURL)
	trNode, err := htmlquery.Find(doc, "//div[@class='hfeed site']//table[@class='bg']//tbody//tr")
	if err != nil {
		log.Printf("err: %v\n", err)
		// clog.Warn(err.Error())
	}
	for i := 3; i < len(trNode); i++ {
		tdNode, _ := htmlquery.Find(trNode[i], "//td")
		ip := htmlquery.InnerText(tdNode[1])
		port := htmlquery.InnerText(tdNode[2])
		Type := htmlquery.InnerText(tdNode[6])

		IP := NewIP()
		IP.Data = ip + ":" + port

		if Type == "yes" {
			IP.Type1 = "http"
			IP.Type2 = "https"

		} else if Type == "no" {
			IP.Type1 = "http"
		}

		// clog.Info("[PLP] ip.Data = %s,ip.Type = %s,%s", IP.Data, IP.Type1, IP.Type2)

		result = append(result, IP)
	}

	// clog.Info("PLP done.")
	return result
}

// IP struct
type IP struct {
	ID    int64  `xorm:"pk autoincr" json:"-"`
	Data  string `xorm:"NOT NULL" json:"ip"`
	Type1 string `xorm:"NOT NULL" json:"type1"`
	Type2 string `xorm:"NULL" json:"type2,omitempty"`
	Speed int64  `xorm:"NOT NULL" json:"speed,omitempty"`
}

// NewIP .
func NewIP() *IP {
	return &IP{Speed: 100}
}

// func main() {
// 	result := PLP()
// 	// fmt.Println(result)
// 	for i, j := range result {
// 		fmt.Printf("%v_%v\n", i, j.Data)
// 	}
// }
