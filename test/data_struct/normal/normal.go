package datastruct

import "sync"

type NormalStruct struct {
	Age int `年龄`

	Name string `姓名`

	// - 1.当你需要在其他包中使用这个结构体或其字段时，必须首字母大写
	// - 2.小写字段在包外完全不可见，即使你能访问结构体本身
	ptr *int

	Address uint64

	// 等同于Slice []any{}，目前主流写any
	// 动态大小的数组切片
	// 类似于其他语言中的"通用数组"或"对象数组"
	Slice []interface{}

	Call func() byte

	Map map[string]string

	Once sync.Once

	// func (n *NormalStruct) init() {
	// 	n.Once.Do(func() {
	// 		// 初始化操作
	// 		fmt.Println("初始化操作")
	// 	})
	// }
}
