package link

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetLinkCacheValidTime(validDate time.Time) int64 {
	if validDate.IsZero() {
		return 0 // DEFAULT_CACHE_VALID_TIME
	}
	return validDate.Sub(time.Now()).Milliseconds()
}

func GetActualIp(r *http.Request) string {
	headers := []string{"X-Forwarded-For", "Proxy-Client-IP", "WL-Proxy-Client-IP", "HTTP_CLIENT_IP", "HTTP_X_FORWARDED_FOR"}
	for _, header := range headers {
		ip := r.Header.Get(header)
		if ip != "" && strings.ToLower(ip) != "unknown" {
			return ip
		}
	}
	return r.RemoteAddr
}

func GetOs(r *http.Request) string {
	userAgent := strings.ToLower(r.Header.Get("User-Agent"))
	switch {
	case strings.Contains(userAgent, "windows"):
		return "Windows"
	case strings.Contains(userAgent, "mac"):
		return "Mac OS"
	case strings.Contains(userAgent, "linux"):
		return "Linux"
	case strings.Contains(userAgent, "android"):
		return "Android"
	case strings.Contains(userAgent, "iphone"), strings.Contains(userAgent, "ipad"):
		return "iOS"
	default:
		return "Unknown"
	}
}

func GetBrowser(r *http.Request) string {
	userAgent := strings.ToLower(r.Header.Get("User-Agent"))
	switch {
	case strings.Contains(userAgent, "edg"):
		return "Microsoft Edge"
	case strings.Contains(userAgent, "chrome"):
		return "Google Chrome"
	case strings.Contains(userAgent, "firefox"):
		return "Mozilla Firefox"
	case strings.Contains(userAgent, "safari"):
		return "Apple Safari"
	case strings.Contains(userAgent, "opera"):
		return "Opera"
	case strings.Contains(userAgent, "msie"), strings.Contains(userAgent, "trident"):
		return "Internet Explorer"
	default:
		return "Unknown"
	}
}

func GetDevice(r *http.Request) string {
	userAgent := strings.ToLower(r.Header.Get("User-Agent"))
	if strings.Contains(userAgent, "mobile") {
		return "Mobile"
	}
	return "PC"
}

func GetNetwork(r *http.Request) string {
	actualIp := GetActualIp(r)
	if strings.HasPrefix(actualIp, "192.168.") || strings.HasPrefix(actualIp, "10.") {
		return "WIFI"
	}
	return "Mobile"
}

func ExtractDomain(rawUrl string) string {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	domain := parsedUrl.Host
	if strings.HasPrefix(domain, "www.") {
		domain = domain[4:]
	}
	return domain
}
