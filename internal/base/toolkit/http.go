package toolkit

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/mssola/user_agent"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// GetTitleByUrl 通过 Url 获取页面标题
func GetTitleByUrl(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", fmt.Errorf("invalid Url: %v", err)
	}

	resp, err := http.Get(parsedUrl.String())
	if err != nil {
		return "", fmt.Errorf("error while fetching Url: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error while closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return "", fmt.Errorf("error while parsing document: %v", err)
		}
		return doc.Find("title").Text(), nil
	}
	return "Error while fetching title.", nil
}

// GetRequestInfo 获取请求中的操作系统、浏览器、设备类型（PC/Mobile）、网络类型（WIFI/4G/5G）
func GetRequestInfo(c *fiber.Ctx) (string, string, string, string) {
	ua := user_agent.New(c.Get("User-Agent"))

	// 获取操作系统和浏览器信息
	os := ua.OS()
	name, version := ua.Browser()

	// 获取设备类型
	var deviceType string
	if ua.Mobile() {
		deviceType = "Mobile"
	} else {
		deviceType = "PC"
	}

	// 获取网络类型
	ip := c.IP()
	networkType := getNetworkType(ip)

	return os, fmt.Sprintf("%s %s", name, version), deviceType, networkType
}

// getNetworkType 根据IP地址获取网络类型
func getNetworkType(ip string) string {
	if strings.HasPrefix(ip, "192.168.") || strings.HasPrefix(ip, "10.") {
		return "WIFI"
	}
	return "Mobile"
}

// IsValidDomain 校验域名是否合法
// 支持 <ip>:<port> 格式
func IsValidDomain(domain string) bool {
	var domainRegex = regexp.MustCompile(
		`^((?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,}|(?:\d{1,3}\.){3}\d{1,3}|localhost)(:\d{1,5})?$`)
	return domainRegex.MatchString(domain)
}

// IsValidUrl 校验是否是合法的 URL
func IsValidUrl(rawUrl string) bool {
	_, err := url.ParseRequestURI(rawUrl)
	return err == nil
}
