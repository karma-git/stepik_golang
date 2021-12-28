/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.

Закрывающая фигурная скобка в конце main() вам не видна, но она есть.

Пакет main объявлять не нужно!

Удачи!

Sample Input:

Sample Output:

*/
package main

type Gta struct {
	On    bool
	Ammo  int
	Power int
}

func (g *Gta) Shoot() bool {
	if !g.On || g.Ammo <= 0 {
		return false
	} else {
		*&g.Ammo--
		return true
	}
}

func (g *Gta) RideBike() bool {
	// gta not on
	if !g.On || g.Power <= 0 {
		return false
	} else {
		*&g.Power--
	}
	return true
}

func main() {
	/* test case
	g := Gta{true, 5, 5}
	fmt.Println(g.Shoot(), g.Ammo)
	fmt.Println(g.RideBike(), g.Power)
	*/
	// we can pass pointer to our struct via new() or &
	testStruct := new(Gta)
}
