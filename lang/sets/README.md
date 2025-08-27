# Sets Package

Sets 包提供了一套泛型集合操作工具，实现了类似数学中集合的数据结构，支持常见的集合操作如并集、交集、差集等。

## 核心特性

- **泛型支持**：支持所有可比较类型（comparable）的集合
- **完整的集合操作**：提供并集、交集、差集等操作
- **高效的性能**：基于 Go 原生 map 实现，查找和插入操作时间复杂度为 O(1)
- **内存优化**：使用空结构体(struct{})作为值类型，最小化内存占用

## 使用方法

### 基本使用
```go
// 创建字符串集合
s1 := sets.NewString("a", "b", "c")
s2 := sets.NewString("b", "c", "d")

// 并集操作
union := s1.Union(s2) // {"a", "b", "c", "d"}

// 交集操作
intersection := s1.Intersection(s2) // {"b", "c"}

// 差集操作
difference := s1.Difference(s2) // {"a"}
```

### 使用场景
#### 1. 去重操作
```go
// 从切片中去除重复元素
data := []int{1, 2, 1, 3, 2, 4}
uniqueSet := sets.FromSlice(data)
unique := uniqueSet.ToSlice()
// unique = [1, 2, 3, 4] (顺序可能不同)
```

#### 2. 标签或分类管理
```go
// 资源标签管理
resourceTags := sets.FromSlice([]string{"env:prod", "team:backend", "region:us-east"})
requiredTags := sets.FromSlice([]string{"env:prod", "team:backend"})

// 检查资源是否包含所有必需标签
hasAllRequired := requiredTags.Difference(resourceTags).IsEmpty()

```

#### 3. 数据对比
```go
// 比较两组数据
currentData := sets.FromSlice([]string{"a", "b", "c"})
expectedData := sets.FromSlice([]string{"b", "c", "d"})

// 找出新增项
addedItems := currentData.Difference(expectedData)

// 找出缺失项
missingItems := expectedData.Difference(currentData)
```


## 性能优势
* 基于 Go 原生 map 实现，提供 O(1) 时间复杂度的查找、插入和删除操作
* 内存使用优化，避免不必要的内存分配
* 提供批量操作方法，减少循环开销