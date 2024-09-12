package patterns

type Prototype interface {
	clone() Prototype
}

type Goods struct {
	tag  int
	flag string
}

func (g Goods) clone() Prototype {
	newGoods := Goods{}
	newGoods.tag = g.tag
	newGoods.flag = g.flag
	return newGoods
}
