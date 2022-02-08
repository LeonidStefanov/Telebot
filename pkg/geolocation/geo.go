package geolocation

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ipinfo/go/ipinfo"
)

type ApiIP struct {
	apiIP *ipinfo.Client
}

func NewIPinfo(timeout time.Duration) *ApiIP {

	return &ApiIP{
		apiIP: ipinfo.NewClient(&http.Client{
			Timeout: timeout,
		}),
	}

}

func (a *ApiIP) GetGeo(ip string) string {
	info, err := a.apiIP.GetGeo(net.ParseIP(ip))
	if err != nil {
		log.Fatal(err)
	}

	buf, err := json.Marshal(info)

	// fmt.Println(string(buf))

	fmt.Println(string(buf))

	return string(buf)

}
