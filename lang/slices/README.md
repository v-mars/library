# Slices 包

## 简介

Slices 包是一个基于 Go 泛型实现的切片工具包，提供了常见的切片操作功能，包括转换、分组、去重、填充、分块、映射转换和反转等操作。

## 核心特性

- 基于泛型实现，支持任意类型
- 提供高效的切片操作函数
- 简化常见的切片处理逻辑
- 类型安全，在编译时进行类型检查

## 功能列表

1. **Transform** - 将切片中的每个元素通过转换函数映射为另一种类型的元素
2. **TransformWithErrorCheck** - 带错误检查的转换函数
3. **GroupBy** - 根据指定函数将切片元素分组
4. **Unique** - 去除切片中的重复元素
5. **Fill** - 创建指定长度并用指定值填充的切片
6. **Chunks** - 将切片分割成指定大小的块
7. **ToMap** - 将切片转换为映射
8. **Reverse** - 反转切片中元素的顺序

## 使用方法

### 安装

```bash
go get your-project/pkg/lang/slices
```

### 导入

```go
import "your-project/pkg/lang/slices"
```

### Transform 函数

将切片中的每个元素通过转换函数映射为另一种类型的元素：

```go
// 将整数切片转换为字符串切片
numbers := []int{1, 2, 3, 4, 5}
strings := slices.Transform(numbers, func(n int) string {
    return fmt.Sprintf("number-%d", n)
})
// 结果: ["number-1", "number-2", "number-3", "number-4", "number-5"]
```

### TransformWithErrorCheck 函数

带错误检查的转换函数，当转换过程中可能出错时使用：

```go
// 将字符串切片转换为整数切片，可能会出错
strings := []string{"1", "2", "invalid", "4"}
numbers, err := slices.TransformWithErrorCheck(strings, func(s string) (int, error) {
    return strconv.Atoi(s)
})
if err != nil {
    // 处理转换错误
    log.Printf("转换错误: %v", err)
}
```

### GroupBy 函数

根据指定函数将切片元素分组：

```go
// 按照年龄将人员分组
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 25},
    {"David", 30},
}

grouped := slices.GroupBy(people, func(p Person) (int, Person) {
    return p.Age, p
})
// 结果: map[25:[{Alice 25} {Charlie 25}] 30:[{Bob 30} {David 30}]]
```

### Unique 函数

去除切片中的重复元素：

```go
// 去除整数切片中的重复元素
numbers := []int{1, 2, 2, 3, 3, 4, 5, 5}
uniqueNumbers := slices.Unique(numbers)
// 结果: [1, 2, 3, 4, 5]
```

### Fill 函数

创建指定长度并用指定值填充的切片：

```go
// 创建长度为5，用"hello"填充的字符串切片
filled := slices.Fill("hello", 5)
// 结果: ["hello", "hello", "hello", "hello", "hello"]
```

### Chunks 函数

将切片分割成指定大小的块：

```go
// 将整数切片分割成大小为3的块
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
chunks := slices.Chunks(numbers, 3)
// 结果: [[1, 2, 3], [4, 5, 6], [7, 8, 9], [10]]
```

### ToMap 函数

将切片转换为映射：

```go
// 将人员切片转换为以姓名为键的映射
people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 35},
}

peopleMap := slices.ToMap(people, func(p Person) (string, Person) {
    return p.Name, p
})
// 结果: map[Alice:{Alice 25} Bob:{Bob 30} Charlie:{Charlie 35}]
```

### Reverse 函数

反转切片中元素的顺序：

```go
// 反转整数切片
numbers := []int{1, 2, 3, 4, 5}
reversed := slices.Reverse(numbers)
// 结果: [5, 4, 3, 2, 1]
// 注意：这是原地反转，numbers切片本身也会被修改
```

## 性能特点

- 基于 Go 原生切片操作实现，性能高效
- 预分配内存空间，减少内存重新分配
- 避免不必要的数据复制，提高执行效率

## 注意事项

1. Reverse 函数会原地修改传入的切片
2. 当传入 nil 切片时，大多数函数会返回 nil
3. 使用泛型确保类型安全，避免运行时类型错误