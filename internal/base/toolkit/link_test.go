package toolkit

import "testing"

func TestGetFavicon(t *testing.T) {
	targetUrl := "https://blog.csdn.net/csdnnews/article/details/143586469"
	favicon, err := GetFavicon(targetUrl)
	if err != nil {
		t.Errorf("GetFavicon error: %v", err)
	} else {
		if favicon == "" {
			t.Errorf("GetFavicon failed: favicon is empty")
		} else {
			t.Logf("GetFavicon success: favicon is %s", favicon)
		}
	}
}

func TestGetTitleAndFavicon(t *testing.T) {
	targetUrl := "https://blog.csdn.net/csdnnews/article/details/143586469"
	title, favicon, err := GetTitleAndFavicon(targetUrl)
	if err != nil {
		t.Errorf("GetTitleAndFavicon error: %v", err)
	} else {
		if title == "" {
			t.Errorf("GetTitleAndFavicon failed: title is empty")
		} else {
			t.Logf("GetTitleAndFavicon success: title is %s", title)
		}
		if favicon == "" {
			t.Errorf("GetTitleAndFavicon failed: favicon is empty")
		} else {
			t.Logf("GetTitleAndFavicon success: favicon is %s", favicon)
		}
	}
}

func BenchmarkGetTitleAndFavicon(b *testing.B) {
	targetUrl := "https://blog.csdn.net/csdnnews/article/details/143586469"
	for i := 0; i < b.N; i++ {
		_, _, err := GetTitleAndFavicon(targetUrl)
		if err != nil {
			b.Errorf("GetTitleAndFavicon error: %v", err)
		}
	}
}
