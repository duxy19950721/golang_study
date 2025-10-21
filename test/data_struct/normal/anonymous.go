package datastruct

type AnonymousStruct struct {
	Age int `年龄`

	Name string `姓名`

	// 匿名字段
	string
}

// 获取匿名字段的值
func (a *AnonymousStruct) GetAnonymous() string {
	return a.string
}

// 设置匿名字段的值
func (a *AnonymousStruct) SetAnonymous(s string) {
	a.string = s
}
