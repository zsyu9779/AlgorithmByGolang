package pattern

import "fmt"

type PType int
const (
	ProductA PType= 1
	ProductB PType= 2
)

type Factory struct {
}

func (f *Factory) Generate(t PType) Product {
	switch t {
	case ProductA:
		return &Product1{}
	case ProductB:
		return &Product2{}
	default:
		return nil
	}

}
type Product interface {
	create()
}
type Product1 struct {
}

func (p *Product1) create() {
	fmt.Println("create product1")
}

type Product2 struct {
}

func (p *Product2) create() {
	fmt.Println("create product2")
}

func FactoryTest()  {
	factory := new(Factory)
	p1 :=factory.Generate(ProductA)
	p1.create()
	p2 :=factory.Generate(ProductB)
	p2.create()

}