package main

import "fmt"


func main() {

	// 创建实例 
	var i1 Item
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		X: 3,
		// Y: 4,
	}
	fmt.Printf("i3: %#v\n", i3)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))


	// 方法
	i3.Move(100, 200)
	fmt.Printf("i2 (Move): %#v\n", i3)


	p1 := Player{
		Name: "Jie",
		Item: i1,
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	

	// interface 相关
	ms := []mover{
		&i1,
		&i2, 
		&p1,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}


	k := Jade
	fmt.Println("k:", k)
	fmt.Println("key:", Key(17))


	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
}



// Go 版本的枚举 enum 
const (
	Jade Key = iota + 1  // 翡翠
	Copper               // 铜
	Crystal              // 水晶
	invalidKey
)

type Key byte


// 实现了 fmt.Stringer 的接口（魔改）
func (k Key) String() string {
	switch k {
	case Jade:
		return "翡翠"
	case Copper:
		return "铜"
	case Crystal:
		return "水晶"
	}

	return fmt.Sprintf("<Key %d>", k)
}


// 接受 interfaces，返回 types
func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}

}

type mover interface {
	Move(x, y int)
}


func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}

	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil 
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	
	return false
}

type Player struct {
	Name string
	Item  // 内嵌，匿名字段
	Keys []Key
}


// i 被称为 receiver，*i 可以改变实例的值
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// Go 编译器 “逃逸分析”，i 分配到堆上
	return &i, nil
}

// 工厂模式：零值 vs 字面量

const (
	maxX = 1_000
	maxY = 600
)

// Item 游戏中的一个物品
type Item struct {
	X int
	Y int
}

