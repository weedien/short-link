package database

import (
	"errors"
	"fmt"
	"shortlink/internal/user/adapter/po"
	"testing"
)

func TestShardingSuffixes(t *testing.T) {
	numberOfShards := 16
	var expected []string
	for i := 0; i < numberOfShards; i++ {
		expected = append(expected, fmt.Sprintf("_%d", i%numberOfShards))
	}

	result := shardingSuffix(numberOfShards)()
	for i, suffix := range result {
		if suffix != expected[i] {
			t.Errorf("Expected %s, but got %s", expected[i], suffix)
		}
	}
}

func TestHashModeShardingAlgorithm(t *testing.T) {
	shardingAlg := hashModeShardingAlgorithm()

	tests := []struct {
		value    any
		expected string
		err      error
	}{
		{"example_username", "_14", nil},
		{"another_username", "_1", nil},
		{"", "", errors.New("invalid username")},
	}

	for _, test := range tests {
		result, err := shardingAlg(test.value)
		if result != test.expected || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf("For value %v, expected %s and error %v, but got %s and error %v", test.value, test.expected, test.err, result, err)
		}
	}
}

func TestInitSharding(t *testing.T) {
	// This test will just check if InitSharding runs without errors
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InitSharding panicked with error: %v", r)
		}
	}()

	ConnectToDatabase()
}

// 分片功能测试
//
//	func TestSharding(t *testing.T) {
//		InitSharding()
//		now := time.Now()
//		err := DB.Create(&entity.Group{ID: 1752265619253805057, Gid: "tSUBMP", Tag: "默认分组", Username: "admin", SortOrder: 0, CreateTime: now, UpdateTime: now, DelFlag: 0}).Error
//		if err != nil {
//			t.Errorf("Create group failed: %v", err)
//		}
//		var group entity.Group
//		err = DB.Model(&entity.Group{}).Where("username = ?", "admin").First(&group).Error
//		if err != nil {
//			t.Errorf("Query group failed: %v", err)
//		}
//		t.Logf("group: %v", group)
//	}
func TestSharding(t *testing.T) {
	DB := ConnectToDatabase()
	err := DB.Create(&po.User{
		Username: "admin", Password: "admin123456",
		RealName: "admin", Phone: "yKZz0xLyjNb9LSCOCfJD4w==", Mail: "02/9oF/nWTBK0cM8UPtCOw==",
	}).Error
	if err != nil {
		t.Errorf("Create user failed: %v", err)
	}
	var user po.User
	err = DB.Model(&po.User{}).Where("username = ?", "admin").First(&user).Error
	if err != nil {
		t.Errorf("Query user failed: %v", err)
	}
	t.Logf("user: %v", user)

	//q := query.Use(DB)
	//q.User.WithContext(nil).Where(q.User.Username.Eq("admin")).First()
}
