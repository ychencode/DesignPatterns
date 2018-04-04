package main

type IOperation interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

type Operation struct {
	a float64
	b float64
}

func (op *Operation) SetA(a float64) {
	op.a = a
}

func (op *Operation) SetB(b float64) {
	op.b = b
}

type AddOperation struct {
	Operation
}

func (a *AddOperation) GetResult() float64 {
	return a.a + a.b
}

type SubOperation struct {
	Operation
}

func (s *SubOperation) GetResult() float64 {
	return s.a - s.b
}

type MulOperation struct {
	Operation
}

func (m *MulOperation) GetResult() float64 {
	return m.a * m.b
}

type DivOperation struct {
	Operation
}

func (d *DivOperation) GetResult() float64 {
	if d.b == 0 {
		return 0
	}
	return d.a / d.b
}

type IFactory interface {
	CreateOperation() Operation
}

type AddFactory struct {

}

func (a *AddFactory) CreateOperation() IOperation {
	return &AddOperation{}
}

type SubFactory struct {

}

func (s *SubFactory) CreateOperation() IOperation{
	return &SubOperation{}
}

type MulFactory struct {

}

func (m *MulFactory) CreateOperation() IOperation {
	return &MulOperation{}
}

type DivFactory struct {

}

func (d *DivFactory) CreateOperation() IOperation {
	return &DivOperation{}
}

func main() {
	addFactory := &AddFactory{}
	addOperation := addFactory.CreateOperation()
	addOperation.SetA(1)
	addOperation.SetB(2)
	println(addOperation.GetResult())

	mulFactory := &MulFactory{}
	mulOperation := mulFactory.CreateOperation()
	mulOperation.SetA(10)
	mulOperation.SetB(5)
	println(mulOperation.GetResult())
}