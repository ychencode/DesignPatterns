package main

type Data struct {
	a float64
	b float64
}

func (d *Data) SetA(a float64) {
	d.a = a
}

func (d *Data) SetB(b float64) {
	d.b = b
}

func (d *Data) GetA() float64 {
	return d.a
}

func (d *Data) GetB() float64 {
	return d.b
}

type IOperation interface {
	GetResult(*Data) float64
}

type AddOperation struct {
}

type SubOperation struct {
}

type MulOperation struct {
}

type DivOperation struct {
}

func (a *AddOperation) GetResult(data *Data) float64 {
	return data.GetA() + data.GetB()
}

func (a *SubOperation) GetResult(data *Data) float64 {
	return data.GetA() - data.GetB()
}

func (a *MulOperation) GetResult(data *Data) float64 {
	return data.GetA() * data.GetB()
}

func (a *DivOperation) GetResult(data *Data) float64 {
	if data.GetB() == 0 {
		return 0
	}
	return data.GetA() / data.GetB()
}

type OperationFactory struct {
}

func (o *OperationFactory) GetOperation(opt string) IOperation {
	switch opt {
	case "+":
		return &AddOperation{}
	case "-":
		return &SubOperation{}
	case "*":
		return &MulOperation{}
	case "/":
		return &DivOperation{}
	default:
		return nil
	}
}

func main() {
	factory := &OperationFactory{}
	addOpt := factory.GetOperation("+")
	subOpt := factory.GetOperation("-")
	mulOpt := factory.GetOperation("*")
	divOpt := factory.GetOperation("/")
	data := &Data{}
	data.SetA(1)
	data.SetB(2)
	println(addOpt.GetResult(data))
	println(subOpt.GetResult(data))
	println(mulOpt.GetResult(data))
	println(divOpt.GetResult(data))
}
