package main

func foo0() (int){
	return 1
}

func foo1() (int){
	return 2
}

func foo2() (int){
	return 3
}

func main() {
	o,p,q :=1,2,3

	x0 := !!!((foo0()-foo1())-(-(-(-1)))*foo0() > 6)
	print(x0,"\n")

	x1 := (foo0()*q*(-3)) != (-(-1)) || !!((foo0()+(+(+2))-(-(-1)))+(+(-(-1)))-foo0() == foo0())
	print(x1,"\n")

	x2 := (((foo0()*-p-foo2()))+(+foo2())-o)-(+(-(-1)))+foo1()
	print(x2,"\n")

	x3 := !((foo1()*+q) != q+((+2)-(-(-1))))
	print(x3,"\n")

	x4 := !!!!((foo1()-(-q)/foo1()+(+6)*(6+(-q)+(-3)))+foo2()*p != q)
	print(x4,"\n")

	x5 := !!!(((+2)+p*+foo2()) < p)
	print(x5,"\n")

	x6 := (q*-foo0()+(+(+2))-6*o-(+(-(-1)))-foo1()) < (-(-1))
	print(x6,"\n")

	x7 := !((-(-1))+foo2()+(-foo1())/((-3)-(-p)/(-(-1))-o)-foo0() >= p*-foo0())
	print(x7,"\n")

	x8 := ((-(-1))*q-foo1()*-(foo0()-(-3)))
	print(x8,"\n")

	x9 := (((-(-1))+foo2())*+foo0()) <= foo0()
	print(x9,"\n")

	x10 := (+2)+foo0() >= (6-foo0()*foo0()+((-(-1))+(+2)))/foo0()
	print(x10,"\n")

	x11 := !!!(foo2() == p-(((o-(-q)+p+(-3)+foo2()+(-p)*foo1()*(-3)-(+2)+(+6)+o-(-(-1)))+(-(+2))*(+2)))) || !(foo2() < (+2))
	print(x11,"\n")

	x12 := o > (((foo1()-(+(-3))*(+2)+(-(-(-1)))-foo1()))+(-(+2))-(-3))
	print(x12,"\n")

	x13 := o*+foo2() == (o*-(-(-1))) && (-3)+(+(-(-1)))+(6-p)/-((-(-1))*-foo0()-(-3))-(+(-3))-foo1()+(+2)+(-q)/(-3) >= foo1()
	print(x13,"\n")

	x14 := !((foo0()+(-p)+(+2))/-foo2() >= (-(-1)))
	print(x14,"\n")

	print(o,p,q,"\n")
}
