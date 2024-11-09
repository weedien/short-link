package main

//import (
//	"fmt"
//	"github.com/bytedance/sonic"
//	"reflect"
//	"shortlink/internal/base/base_event"
//	"shortlink/internal/link/app/event"
//	"shortlink/internal/link/domain/valobj"
//	"time"
//)
//
//func main() {
//	recordInfo := valobj.LinkStatsRecordVo{
//		FullShortUrl: "https://t.cn/8fj3",
//		RemoteAddr:   "175.16.23.33",
//		OS:           "Windows",
//		Browser:      "Chrome 126.5.431",
//		Device:       "PC",
//		Network:      "WIFI",
//		UV:           "aa691b45-0b96-4b6d-9534-8648a8e26b9f",
//		UVFirstFlag:  true,
//		UipFirstFlag: true,
//		Keys:         "Nothing",
//		CurrentDate:  time.Now(),
//	}
//	e := event.NewLinkAccessedEvent(recordInfo)
//	marshal, err := sonic.Marshal(e)
//	if err != nil {
//		fmt.Printf("序列化失败: %v\n", err)
//	}
//	fmt.Printf("序列化后的消息: %v\n", string(marshal))
//
//	// 尝试还原消息
//	eventType := reflect.TypeOf(e)
//	fmt.Printf("消息类型 String: %s, Tag: %s\n", eventType.String(), eventType.Tag())
//	eventPtr := reflect.New(eventType).Interface()
//	if err = sonic.Unmarshal(marshal, eventPtr); err != nil {
//		fmt.Printf("反序列化失败: %v\n", err)
//	}
//	fmt.Printf("反序列化后的消息: %#v\n", eventPtr.(base_event.Event).(*event.RecordLinkVisitEvent))
//}
