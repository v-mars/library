package slices

import (
	"cmdb/pkg/lang/conv"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试Transform函数
func TestTransform(t *testing.T) {
	// 测试正常情况
	numbers := []int{1, 2, 3, 4, 5}
	strings := Transform(numbers, func(n int) string {
		return string(rune('a' + n - 1))
	})
	expected := []string{"a", "b", "c", "d", "e"}
	assert.Equal(t, expected, strings)

	// 测试空切片
	var empty []int
	emptyResult := Transform(empty, func(n int) string {
		return ""
	})
	assert.Nil(t, emptyResult)

	// 测试nil切片
	var nilSlice []int
	nilResult := Transform(nilSlice, func(n int) string {
		return ""
	})
	assert.Nil(t, nilResult)
}

// 测试TransformWithErrorCheck函数
func TestTransformWithErrorCheck(t *testing.T) {
	// 测试正常情况
	strings := []string{"1", "2", "3", "4", "5"}
	numbers, err := TransformWithErrorCheck(strings, func(s string) (int, error) {
		switch s {
		case "1":
			return 1, nil
		case "2":
			return 2, nil
		case "3":
			return 3, nil
		case "4":
			return 4, nil
		case "5":
			return 5, nil
		default:
			return 0, errors.New("invalid number")
		}
	})
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, numbers)

	// 测试错误情况
	invalidStrings := []string{"1", "2", "invalid"}
	_, err = TransformWithErrorCheck(invalidStrings, func(s string) (int, error) {
		switch s {
		case "1":
			return 1, nil
		case "2":
			return 2, nil
		default:
			return 0, errors.New("invalid number")
		}
	})
	assert.Error(t, err)
	assert.Equal(t, "invalid number", err.Error())

	// 测试空切片
	var empty []string
	emptyResult, err := TransformWithErrorCheck(empty, func(s string) (int, error) {
		return 0, nil
	})
	assert.NoError(t, err)
	assert.Nil(t, emptyResult)

	// 测试nil切片
	var nilSlice []string
	nilResult, err := TransformWithErrorCheck(nilSlice, func(s string) (int, error) {
		return 0, nil
	})
	assert.NoError(t, err)
	assert.Nil(t, nilResult)
}

// 测试GroupBy函数
func TestGroupBy(t *testing.T) {
	// 定义测试结构体
	type Person struct {
		Name string
		Age  int
		//List []string
	}

	// 测试正常情况
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"David", 30},
	}
	// func GroupBy[A, K comparable, V any](src []A, fn func(A) (K, V)) map[K][]V {})
	grouped := GroupBy(people, func(p Person) (int, Person) {
		return p.Age, p
	})

	expected := map[int][]Person{
		25: {{"Alice", 25}, {"Charlie", 25}},
		30: {{"Bob", 30}, {"David", 30}},
	}
	assert.Equal(t, expected, grouped)

	// 测试空切片
	var empty []Person
	emptyResult := GroupBy(empty, func(p Person) (int, Person) {
		return p.Age, p
	})
	assert.Nil(t, emptyResult)

	// 测试nil切片
	var nilSlice []Person
	nilResult := GroupBy(nilSlice, func(p Person) (int, Person) {
		return p.Age, p
	})
	assert.Nil(t, nilResult)

	type Person2 struct {
		Name string
		Age  int
		Data map[string]string
	}

	// 测试正常情况
	people2 := []Person2{
		{"Alice", 25, map[string]string{"key1": "value1", "key2": "value2"}},
		{"Bob", 30, map[string]string{"key1": "value2", "key2": "value2"}},
		{"Charlie", 35, map[string]string{"key1": "value1", "key2": "value2"}},
		{"Charlie2", 36, map[string]string{"key1": "value1", "key2": "value2"}},
		{"Charlie3", 36, map[string]string{"key1": "value3", "key2": "value3"}},
	}
	nilResult2 := GroupBy(people2, func(p Person2) (string, Person2) {
		return p.Data["key1"], p
	})
	fmt.Println("nilResult2:", conv.DebugJsonToStr(nilResult2))
}

// 测试Unique函数
func TestUnique(t *testing.T) {
	// 测试正常情况
	numbers := []int{1, 2, 2, 3, 3, 4, 5, 5}
	uniqueNumbers := Unique(numbers)
	// 注意：Unique函数不保证顺序，只要元素唯一即可
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, uniqueNumbers)

	// 测试字符串
	strings := []string{"a", "b", "b", "c", "a"}
	uniqueStrings := Unique(strings)
	assert.ElementsMatch(t, []string{"a", "b", "c"}, uniqueStrings)

	// 测试空切片
	var empty []int
	emptyResult := Unique(empty)
	assert.Nil(t, emptyResult)

	// 测试nil切片
	var nilSlice []int
	nilResult := Unique(nilSlice)
	assert.Nil(t, nilResult)
}

// 测试Fill函数
func TestFill(t *testing.T) {
	// 测试正常情况
	filled := Fill("hello", 5)
	expected := []string{"hello", "hello", "hello", "hello", "hello"}
	assert.Equal(t, expected, filled)

	// 测试填充0个元素
	empty := Fill("hello", 0)
	assert.Equal(t, []string{}, empty)

	// 测试填充数字
	numbers := Fill(42, 3)
	expectedNumbers := []int{42, 42, 42}
	assert.Equal(t, expectedNumbers, numbers)
}

// 测试Chunks函数
func TestChunks(t *testing.T) {
	// 测试正常情况
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunks := Chunks(numbers, 3)
	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
	assert.Equal(t, expected, chunks)

	// 测试块大小大于切片长度
	smallNumbers := []int{1, 2, 3}
	smallChunks := Chunks(smallNumbers, 5)
	expectedSmall := [][]int{{1, 2, 3}}
	assert.Equal(t, expectedSmall, smallChunks)

	// 测试空切片
	var empty []int
	emptyChunks := Chunks(empty, 3)
	assert.Equal(t, [][]int{}, emptyChunks)

	// 测试块大小为0
	zeroChunks := Chunks(numbers, 0)
	// 当chunkSize为0时，应该返回包含整个切片的二维切片
	expectedZero := [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	assert.Equal(t, expectedZero, zeroChunks)

	// 测试块大小为负数
	negativeChunks := Chunks(numbers, -1)
	// 当chunkSize为负数时，应该返回包含整个切片的二维切片
	expectedNegative := [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	assert.Equal(t, expectedNegative, negativeChunks)
}

// 测试ToMap函数
func TestToMap(t *testing.T) {
	// 定义测试结构体
	type Person struct {
		Name string
		Age  int
	}

	// 测试正常情况
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	peopleMap := ToMap(people, func(p Person) (string, Person) {
		return p.Name, p
	})

	expected := map[string]Person{
		"Alice":   {"Alice", 25},
		"Bob":     {"Bob", 30},
		"Charlie": {"Charlie", 35},
	}
	assert.Equal(t, expected, peopleMap)

	// 测试空切片
	var empty []Person
	emptyResult := ToMap(empty, func(p Person) (string, Person) {
		return p.Name, p
	})
	assert.Nil(t, emptyResult)

	// 测试nil切片
	var nilSlice []Person
	nilResult := ToMap(nilSlice, func(p Person) (string, Person) {
		return p.Name, p
	})
	assert.Nil(t, nilResult)
}

// 测试Reverse函数
func TestReverse(t *testing.T) {
	// 测试正常情况
	numbers := []int{1, 2, 3, 4, 5}
	reversed := Reverse(numbers)
	expected := []int{5, 4, 3, 2, 1}
	assert.Equal(t, expected, reversed)
	// 注意：Reverse是原地修改，numbers本身也被修改了
	assert.Equal(t, expected, numbers)

	// 测试偶数个元素
	evenNumbers := []int{1, 2, 3, 4}
	evenReversed := Reverse(evenNumbers)
	expectedEven := []int{4, 3, 2, 1}
	assert.Equal(t, expectedEven, evenReversed)

	// 测试单个元素
	single := []int{42}
	singleReversed := Reverse(single)
	assert.Equal(t, []int{42}, singleReversed)

	// 测试空切片
	var empty []int
	emptyReversed := Reverse(empty)
	assert.Equal(t, []int{}, emptyReversed)
}
