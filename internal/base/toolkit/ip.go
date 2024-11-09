package toolkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type Location struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

// GetLocationByIP 根据ip查询地理位置
func GetLocationByIP(ip string) Location {

	if IsReservedIP(ip) {
		return Location{}
	}

	target := "http://ip-api.com/json/" + ip
	res, err := http.Get(target)
	if err != nil {
		return Location{}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error while closing response body: %v", err)
		}
	}(res.Body)

	location := Location{}
	if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
		fmt.Printf("error while decoding response body: %v", err)
		return Location{}
	}

	return location
}

// IsReservedIP 判断某个IPv4地址是否为保留IP
func IsReservedIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	reservedIPRanges := []struct {
		start string
		end   string
	}{
		{"0.0.0.0", "0.255.255.255"},
		{"10.0.0.0", "10.255.255.255"},
		{"100.64.0.0", "100.127.255.255"},
		{"127.0.0.0", "127.255.255.255"},
		{"169.254.0.0", "169.254.255.255"},
		{"172.16.0.0", "172.31.255.255"},
		{"192.0.0.0", "192.0.0.7"},
		{"192.0.2.0", "192.0.2.255"},
		{"192.88.99.0", "192.88.99.255"},
		{"192.168.0.0", "192.168.255.255"},
		{"198.18.0.0", "198.19.255.255"},
		{"198.51.100.0", "198.51.100.255"},
		{"203.0.113.0", "203.0.113.255"},
		{"224.0.0.0", "239.255.255.255"},
		{"240.0.0.0", "255.255.255.255"},
	}

	for _, r := range reservedIPRanges {
		start := net.ParseIP(r.start)
		end := net.ParseIP(r.end)
		if bytes.Compare(parsedIP, start) >= 0 && bytes.Compare(parsedIP, end) <= 0 {
			return true
		}
	}

	return false
}
