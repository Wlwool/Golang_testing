package main

import (
	// "fmt"
	"math/rand/v2"
)


// func main() {
// 	// fmt.Println(Attack())
// 	// attackWithDamageBoost := DamageBoostDecorator(Attack)
// 	// fmt.Println(attackWithDamageBoost())

// 	// attackWithCriticalHit := CriticalHitDecorator(Attack)
// 	// fmt.Println(attackWithCriticalHit)

// 	// attackWithSlowEffect := SlowEffectDecorator(Attack)
// 	// fmt.Println(attackWithSlowEffect)
// 	attackWithDamageBoost := DamageBoostDecorator(Attack)
// 	attackWithCriticalHit := CriticalHitDecorator(attackWithDamageBoost)
// 	attackWithSlowEffect := SlowEffectDecorator(attackWithCriticalHit)
// 	fmt.Println(attackWithSlowEffect())
// }

func Attack() string {
	return "Атака выполнена"
}

func DamageBoostDecorator(attackFunc func() string) func() string {
	return func() string {
		return "Вам улыбнулась удача, нанесение урона увеличино на 10%! " + attackFunc()
	}
}

func CriticalHitDecorator(attackFunc func() string) func() string {
	return func() string {
		if rand.Float64() < 0.25 {
			return "Критический удар! Урон удвоен! " + attackFunc()
		}
		return attackFunc()
	}
}

func SlowEffectDecorator(attackFunc func() string) func () string {
	return func() string {
		return attackFunc() + " Цель замедлена на 2 хода!"
	}
}
