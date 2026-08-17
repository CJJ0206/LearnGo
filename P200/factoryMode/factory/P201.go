package factory

/*
8.13
golang的结构体没有构造函数，通常可以使用工厂模式来解决这个问题

比如：
	一个结构体的声明是:
	package model
	type Student struct{
		Name string
	}
	因为里面的Student首字母S是大写的，如果我们想在其他包创建Student的实例，引入model包后，就可以直接创建Student结构体的变量
	但是问题来了，如果首字母小写呢，比如type student struct()就不行了，咋办  --> 工厂模式
	fixme 所以工厂模式解决的就是既要首字母小写，又要能偶在别的包用（小写呢是为了保证封装性）
*/

// fixme 因为student的首字母是小写的，所以相当于是私有数据
type student struct {
	name  string
	score float64
}

func NewStudent(name string, score float64) *student {
	return &student{
		name:  name,
		score: score,
	}
}

// GetScore fixme 这里如果是返回指针是会导致底层值被篡改的
func (s *student) GetScore() *float64 {
	return &s.score
	// fixme 但是如果单纯返回 float 时，做的是一个值拷贝，外部拿到的是副本，不会影响真实值
}

// fixme 总结：不论是整个类型还是里面的属性，能够被外部看见并使用，都是通过一个共有的函数来向外界暴露
// fixme 其实到这里已经开始有接口的感觉了
