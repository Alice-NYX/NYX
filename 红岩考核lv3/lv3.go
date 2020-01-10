package main

import (
	"fmt"
	"math/rand"
	"time"
)

type M struct {
	hp      int
	sp      int
	atk     float64
}
type N struct {
	hp      int
	sp      int
	atk     float64
}
var i int
func xunhuan(A *M,B *N) {
	for i := 0; i <= 100; i++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(1-0+1) + 0//随机数
		if x > 0 {
			B.hp -= 10
			if B.hp <= 0 {
				goto Loop
			}
			xunhuan(A, B)
		} else {
			goto Loop
		}
	}
Loop:
}

func attackA(A *M,B *N) {
	for i := 0; i <= 5; i++ {
		B.hp -= 10
		xunhuan(A, B)
		if B.hp <= 0 {
			break
		}
		fmt.Println("A的血量", "A的蓝量", "A的攻击", "B的血量", "B的蓝量", "B的攻击")
		fmt.Println(A.hp, A.sp, A.atk, B.hp, B.sp, B.atk)
	}
}
func attackB(A*M,B*N) {
	for i := 0; i <= 5; i++ {
		A.hp -= 10
		B.sp += 10
		if A.hp<=0 {
			break
		}
		fmt.Println("A的血量", "A的蓝量", "A的攻击", "B的血量", "B的蓝量", "B的攻击")
		fmt.Println(A.hp, A.sp, A.atk, B.hp, B.sp, B.atk)
	}
	if B.sp>=50 {
		A.atk*=0.9
	}
}
func main(){
	var sum int
	var a float64
	var b float64
	var c float64
	var sum1 float64
	var sum2 float64
	a=0
	b=0
	c=0
	sum1 = 0
	sum2 = 0
	for sum=0;sum<=1000;sum++{
		A :=M{
			hp:100,
			sp:0,
			atk:10,
		}
		B :=N{
			hp:300,
			sp:0,
			atk:20,
		}
		for i=0;i<=10;i++{
			attackA(&A,&B)
			attackB(&A,&B)
			if B.hp<=0{
				a++
				break
			}
			if A.hp<=0{
				b++
				break
			}
		}
		c++
	}
	if c!=0{
		sum1=a/c
		sum2=b/c
	}
	fmt.Printf("A的胜率为%f",sum1)
	fmt.Printf("B的胜率为%f",sum2)
}
