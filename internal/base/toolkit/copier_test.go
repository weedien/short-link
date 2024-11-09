package toolkit

//
//import (
//	"testing"
//)
//
//type Address struct {
//	City    string
//	Country string
//}
//
//type Person struct {
//	Tag    string
//	Age     int
//	Address *Address
//	Friends []string
//}
//
//type Student struct {
//	Tag    string
//	Age     int
//	Address *Address
//	friends []string
//}
//
//func TestDeepCopy(t *testing.T) {
//	original := &Person{
//		Tag: "Alice",
//		Age:  30,
//		Address: &Address{
//			City:    "New York",
//			Country: "USA",
//		},
//		Friends: []string{"Bob", "Charlie"},
//	}
//
//	var clone Person
//	if err := DeepCopy(&clone, original); err != nil {
//		t.Errorf("DeepCopy failed: %v", err)
//		return
//	}
//
//	t.Logf("Original: %+v\n", original)
//	t.Logf("Copy: %+v\n", clone)
//
//	// Modify the original to verify deep copy
//	original.Tag = "Alice Changed"
//	original.Address.City = "Los Angeles"
//	original.Friends[0] = "Bob Changed"
//
//	t.Logf("After modification, Original: %+v\n", original)
//	t.Logf("After modification, Copy: %+v\n", clone)
//}
//
//func TestWithOutDeepCopy(t *testing.T) {
//	original := &Person{
//		Tag: "Alice",
//		Age:  30,
//		Address: &Address{
//			City:    "New York",
//			Country: "USA",
//		},
//		Friends: []string{"Bob", "Charlie"},
//	}
//
//	clone := original
//
//	t.Logf("Original: %+v\n", original)
//	t.Logf("Copy: %+v\n", clone)
//
//	// Modify the original to verify deep copy
//	original.Tag = "Alice Changed"
//	original.Address.City = "Los Angeles"
//	original.Friends[0] = "Bob Changed"
//
//	t.Logf("After modification, Original: %+v\n", original)
//	t.Logf("After modification, Copy: %+v\n", clone)
//}
//
//// goos: windows
//// goarch: amd64
//// pkg: shortlink/pkg/toolkit
//// cpu: Intel(R) Core(TM) i5-10200H CPU @ 2.40GHz
//// BenchmarkDeepCopy
//// BenchmarkDeepCopy-8   	 3372711	       348.4 ns/op
//func BenchmarkDeepCopy(b *testing.B) {
//	original := &Person{
//		Tag: "Alice",
//		Age:  30,
//		Address: &Address{
//			City:    "New York",
//			Country: "USA",
//		},
//		Friends: []string{"Bob", "Charlie"},
//	}
//
//	var clone Person
//	for i := 0; i < b.N; i++ {
//		err := DeepCopy(&clone, original)
//		if err != nil {
//			b.Errorf("DeepCopy failed: %v", err)
//			return
//		}
//	}
//}
//
//func TestCopyStruct(t *testing.T) {
//	student := Student{
//		Tag: "Alice",
//		Age:  30,
//		Address: &Address{
//			City:    "New York",
//			Country: "USA",
//		},
//		friends: []string{"Bob", "Charlie"},
//	}
//
//	t.Logf("Student: %+v\n", student)
//
//	var person Person
//	if err := CopyStruct(&person, &student); err != nil {
//		t.Errorf("CopyStruct failed: %v", err)
//	}
//
//	t.Logf("Person: %+v\n", person)
//}
