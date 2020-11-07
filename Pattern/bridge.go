package pattern

import "fmt"

type car interface {
	name() string
	run()
}

type MercedesBenz struct {
}

func (m MercedesBenz) name() string {
	return "Mercedes Benz"
}
func (m MercedesBenz) run() {
	fmt.Println("benz running")
}

type Audi struct {
}

func (a Audi) name() string {
	return "Audi"
}
func (a Audi) run() {
	fmt.Println("audi running")
}

type shiftMode interface {
	up()
	down()
}
type Manual struct {
	car car
}

func (m Manual) up() {
	fmt.Printf("Manual up the %s car\n", m.car.name())
	m.car.run()
}
func (m Manual) down() {
	fmt.Printf("Manual down the %s car\n", m.car.name())
	m.car.run()
}
func getNewManualShift(c car) *Manual {
	return &Manual{car: c}
}

type Automatic struct {
	car car
}

func (a Automatic) up() {
	fmt.Printf("Automatic up the %s car\n", a.car.name())
	a.car.run()
}
func (a Automatic) down() {
	fmt.Printf("Automatic down the %s car\n", a.car.name())
	a.car.run()
}
func getNewAutoShift(c car) *Automatic {
	return &Automatic{car: c}
}

func TestBridge() {
	ma := getNewAutoShift(&MercedesBenz{})
	mm := getNewManualShift(&MercedesBenz{})
	aa := getNewAutoShift(&Audi{})
	am := getNewManualShift(&Audi{})

	ma.up()
	ma.down()
	println("+++++++++++++++++++++++++")
	mm.up()
	mm.down()
	println("+++++++++++++++++++++++++")
	aa.up()
	aa.down()
	println("+++++++++++++++++++++++++")
	am.up()
	am.down()
	println("+++++++++++++++++++++++++")

}
