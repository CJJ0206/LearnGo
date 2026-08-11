package main

import (
	"encoding/json"
	"fmt"
)

// 8.11 fixme 结构体里的内容在地址里是连续分布的

type Point struct {
	x, y int
}

type Rectangle struct {
	leftUp, rightDown Point // 坐标
}

type Rectangle2 struct {
	leftUp, rightDown *Point // 坐标
}

func detail1() {
	rect := Rectangle{Point{1, 2}, Point{3, 4}}
	fmt.Printf("右上的地址为%p,右下 的地址为%p", &rect.leftUp, &rect.rightDown)
	// fixme 右上的地址为0x4a603de6280,右下 的地址为0x4a603de6290  十六进制进了一位，所以是16字节，刚好是两个int
	// fixme 所以结构体在内存中是连续的，所以他底层取内容就是用地址找的，快！ 没有指针掺和，这里就完全是一条扁平连续的内存

	rp := Rectangle2{
		leftUp:    &Point{1, 2},
		rightDown: &Point{3, 4}}
	fmt.Println(rp)
	// fixme 如果是指针的话，指针本身是连续的，但是指向的地址不连续

}

// --------------------------------------------------------------

type A struct {
	Num int
}

type B struct {
	Num int
}

func detail2() {
	var a A
	var b B
	a = A(b) // fixme 这里之所以能强转，是因为这两个结构体里面的属性是一样的，否则就毫无办法了
	// fixme 要完全一样（属性个数、类型、名字都要一样！！）
	fmt.Println(a, b)
}

// fixme 还有一点是：如果用type对结构体取别名，之后go会认为两个是不同的数据类型（这个不管是结构体还是基础数据类型都是这样的）需要强制转换下

// -----------------------------------------------------------
/*
struct 的每个字段（属性）上可以写一个tag，该tag可以通过反射机制获取，常见的场景就是序列化和反序列化

将struct变量进行json处理
	问题：json处理后的字段名也是首字母大写，这样如果我们是将json后的字符串返回给其他程序用，比如php，那么可能他们不适用这个格式怎么办？
	解决方案：使用tag来解决
fixme 这个就和SCADA项目里需要JS映射是一个道理（复习）到时候学反射要认真
*/

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
	// fixme 有了这个之后，最后输出的字段就变成了小写
}

func detail3() {
	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	// 怎么把这个实例序列化成字符串呢,fixme 使用 encoding/json 库处理
	jsonMonster, err := json.Marshal(monster) // fixme 这里面用到了反射才实现这个效果的
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("序列化结果为：", string(jsonMonster)) // 需要转换下，否则是字节输出
}

// fixme 这个就和项目里使用的方法对上点了

// ScriptTrigger 脚本触发配置（t_scripts.trigger JSON 列）。
type ScriptTrigger struct {
	Type       string   `json:"type"`                 // schedule | event | manual
	IntervalMs int      `json:"interval_ms"`          // schedule 周期，≥500
	WatchTags  []string `json:"watch_tags,omitempty"` // event 监听测点
	MissPolicy string   `json:"miss_policy"`          // 默认 skip，防堆积
}

// ScriptDef 组态脚本定义（对应配置库 t_scripts 表）。
type ScriptDef struct {
	ID         string                 `json:"id"`                   // PK 脚本唯一标识
	Name       string                 `json:"name"`                 // 展示名
	Version    int                    `json:"version"`              // 发布自增版本
	Runtime    string                 `json:"runtime"`              // 固定 javascript
	Entry      string                 `json:"entry"`                // 默认 onTick；手动可 main
	Source     string                 `json:"source"`               // JS 源码
	Params     map[string]interface{} `json:"params,omitempty"`     // 工程参数
	Trigger    ScriptTrigger          `json:"trigger"`              // 触发配置
	TimeoutSec float64                `json:"timeout_sec"`          // 单次超时（秒），< interval_ms/1000
	Enabled    bool                   `json:"enabled"`              // 是否调度
	CreatedAt  int64                  `json:"created_at,omitempty"` // Unix 时间戳（节点本地元数据）
	UpdatedAt  int64                  `json:"updated_at,omitempty"` // Unix 时间戳（节点本地元数据）
}
