package main

import (
	"fmt"
	_"go_code/project05/model"
	
)

type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}  

func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}


type Camera struct {
	name string
}
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

type Pupil struct { 
	Student //匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

type Graduate struct {
	Student //匿名结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods//多重继承，尽量不要使用
	Brand	
}

type TV2 struct {
	*Goods
	*Brand	
}

type Monster struct  {
	Name string
	Age int
}

type E struct {
	Monster
	int
	n int
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type AInterface interface {
	Say()
	//不能有任何变量
}
type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type Usb3 interface {
	Say()
}
type Stu3 struct {
}
func (this *Stu3) Say() {
	fmt.Println("Say()")
}

type AInterface2 interface {//接口也可继承多个接口
	BInterface
	CInterface//继承的两个接口不能有相同的方法名
	test03()
}

type Stu2 struct {
}
func (stu Stu2) test01() {
}
func (stu Stu2) test02() {
}
func (stu Stu2) test03() {	
}

type  Hero struct{
	Name string
	Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//等价上面
	hs[i], hs[j] = hs[j], hs[i]
}


func main(){

//工厂模式
	// var stu = model.NewStudent("mm",19.2)
	// fmt.Println(stu)
	// fmt.Println(stu.GetScore())

//封装
	// p := model.NewPerson("smith")
	// p.SetAge(18)
	// p.SetSal(5000)
	// fmt.Println(p)
	// fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())

//继承
	// pupil := &Pupil{}
	// pupil.Student.Name = "aa"//其实Student可以省略
	// pupil.Age = 8
	// pupil.testing() 
	// pupil.SetScore(70)
	// pupil.Student.ShowInfo()
	// fmt.Println("res=", pupil.Student.GetSum(1, 2))
	// //若结构体和匿名结构体有相同的字段或方法，采用就近访问原则
	// //若嵌入多个匿名结构体，有相同的字段或方法，访问时就必须写上匿名结构体的名字
	// //若嵌套了一个有名结构体（组合关系），访问时就必须带上结构体名字
	
	// tv := TV{ Goods{"001", 5000.99},  Brand{"海尔", "山东"}, }

	// tv2 := TV{ 
	// 	Goods{
	// 		Price : 5000.99,
	// 		Name : "002", 
	// 	},  
	// 	Brand{
	// 		Name : "夏普", 
	// 		Address :"北京",
	// 	}, 
	// }

	// fmt.Println(tv.Goods.Name)
	// fmt.Println(tv.Price) 
	// fmt.Println("tv", tv)
	// fmt.Println("tv2", tv2)

	// tv3 := TV2{ &Goods{"003", 7000.99},  &Brand{"创维", "河南"}, }
	// tv4 := TV2{ 
	// 		&Goods{
	// 			Name : "004", 
	// 			Price : 9000.99,
	// 		},  
	// 		&Brand{
	// 			Name : "长虹", 
	// 			Address : "四川",
	// 		}, 
	// 	}
	// fmt.Println("tv3", *tv3.Goods, tv3.Brand)
	// fmt.Println("tv4", *tv4.Goods, *tv4.Brand)

	// var e E
	// e.Name = "aaa"
	// e.Age = 300
	// e.int = 20
	// e.n = 40
	// fmt.Println("e=", e)

//接口
	//接口类型默认是一个指针
	// var stu Stu 
	// var a AInterface = stu
	// a.Say()

	// var stu2 Stu2
	// var a2 AInterface2 = stu2
	// a2.test01()

	// var stu3 Stu3 = Stu3{}
	// var u3 Usb3= &stu3
	// //error//var u Usb3 = stu 3 
	// u3.Say()
	// fmt.Println("here", u3)
	
//实现接口进行排序
	// var heroes HeroSlice
	// for i := 0; i < 10 ; i++ {
	// 	hero := Hero{
	// 		Name : fmt.Sprintf("%d", rand.Intn(100)),
	// 		Age : rand.Intn(100),
	// 	}
	// 	heroes = append(heroes, hero)
	// }

	// for _ , v := range heroes {
	// 	fmt.Println(v)
	// }

	// sort.Sort(heroes)
	// fmt.Println()
	// for _ , v := range heroes {
	// 	fmt.Println(v)
	// }

//多态
	// var usbArr [3]Usb
	// usbArr[0] = Phone{"vivo"}
	// usbArr[1] = Phone{"小米"}
	// usbArr[2] = Camera{"尼康"}

	// fmt.Println(usbArr)

//
}
//工厂模式、抽象、封装、继承、接口、多态、