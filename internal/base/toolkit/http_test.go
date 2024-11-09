package toolkit

import "testing"

func TestGetTitleByUrl(t *testing.T) {
	title, err := GetTitleByUrl("https://cubox.pro/my/card?id=7172619058599168541&query=true")
	if err != nil {
		t.Errorf("GetTitleByUrl() error = %v", err)
	}
	t.Logf("Title: %s", title)
}

//func TestGetTitleByUrl(t *testing.T) {
//	type args struct {
//		rawUrl string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{
//			name: "Test GetTitleByUrl",
//			args: args{
//				rawUrl: "https://www.google.com",
//			},
//			want:    "Google",
//			wantErr: false,
//		},
//		{
//			name: "Test GetTitleByUrl",
//			args: args{
//				rawUrl: "https://www.google.com",
//			},
//			want:    "Google",
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetTitleByUrl(tt.args.rawUrl)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetTitleByUrl() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GetTitleByUrl() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
