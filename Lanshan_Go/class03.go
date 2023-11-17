package main

import "fmt"

type Goods struct {
	name  string
	price float64
	num   int
}

type Egoods struct {
	g     Goods
	brand string
	model string
}

type Manager interface {
	Check()
	Update()
	Plot()
	Eplot()
}

func (u *Egoods) Update() {
	var n int
	fmt.Println("请输入该商品现有库存：")
	fmt.Scanf("%d\n", &n)
	u.g.num = n
	fmt.Println("库存已更新！")
}

func (p1 *Egoods) Plot() {
	fmt.Println("名称：", p1.g.name, "\n价格：",
		p1.g.price, "\n库存数量：", p1.g.num)
}

func (c *Egoods) Check() {
	fmt.Println(c.g.num)
}

func (p2 *Egoods) Eplot() {
	fmt.Println("品牌：", p2.brand,
		"\n型号：", p2.model)
}
func main() {

	g1 := Egoods{g: Goods{"手机", 2000,
		500}, brand: "华为", model: "Mt1024"}

	/*g2 := Egoods{g: Goods{"冲锋衣", 1000, 500}}*/

	fmt.Println("c >> 检查产品数量\n" +
		"u >> 更新产品数量\np >> 打印库存信息" +
		"\nep >>>打印品牌型号信息")

	for {
		var do string
		/*var st string
		fmt.Scanf("%s\n", &st)
		m1 := map[string]Egoods{
			"g1": g1,
			"g2": g2,
		}*/
		fmt.Println("请输入您想要的操作：")
		fmt.Scanf("%s\n", &do)
		switch do {
		case "c":
			g1.Check()
		case "u":
			g1.Update()
		case "p":
			g1.Plot()
		case "ep":
			g1.Eplot()
		default:
			fmt.Println("请输入正确的操作！\nc >> 检查产品数量" +
				"\nu >> 更新产品数量\np >> 打印库存信息\n" +
				"ep >>> 打印产品品牌型号信息")
		}
	}
}
