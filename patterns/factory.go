package patterns

type Product interface {
	info() string
}

type ProductA struct {
}

func (s ProductA) info() string {
	return "ProductA"
}

type ProductB struct {
}

func (s ProductB) info() string {
	return "ProductB"
}

type ProductType1 int

const (
	ProductType1A ProductType1 = iota
	ProductType1B
)

func CreateProduct(tp ProductType1) Product {
	switch tp {
	case ProductType1A:
		return ProductA{}
	case ProductType1B:
		return ProductB{}
	default:
		panic("invalid product type")
	}
}

type Factory interface {
	CreateProduct() Product
}

type FactoryA struct {
	Factory
}

func (f FactoryA) CreateProduct() Product {
	return ProductA{}
}

func c() {
	var f = FactoryA{}
	f.CreateProduct()
}

type FactoryB struct{}

func (f FactoryB) CreateProduct() Product {
	return ProductB{}
}
