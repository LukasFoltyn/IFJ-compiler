package main

func foo0(n float64) (float64) {
	if n < 0.2 {
		return 1.3
	}
	return foo0(n-1.8)
}

func foo1(n float64) (float64) {
	if n < 0.2 {
		return 1.3
	}
	return foo1(n-1.8)
}

func foo2(n float64) (float64) {
	if n < 0.2 {
		return 1.3
	}
	return foo2(n-1.8)
}

func main() {
	o,p,q :=1.2,3.3,7.14

	x0 := o+(+(+2.3))+((((((foo1(4.8)+(+2.3)-(+2.3))))+(-q)-(+2.3))))+o
	print(x0,"\n")

	x1 := (((((-(-1.2))*foo0(7.5)))))+((-(-1.2))*-foo2(6.9))-(+foo2(6.9))-(-3.4)*+6.1*6.1+(+p)+foo0(7.5)/+foo1(4.8)*-((+2.3)-(+(+2.3))/o/foo0(7.5)*+foo2(6.9)-(+(+2.3))/(-3.4)*-o+(-foo2(6.9))+q)
	print(x1,"\n")

	x2 := o+o+(o+(+foo1(4.8))-foo0(7.5)*-(o*+(-3.4))/-(foo0(7.5)/+(6.1-(o/+foo0(7.5)*+foo1(4.8))+(-foo1(4.8))+(-3.4))))/(-3.4)
	print(x2,"\n")

	x3 := ((6.1+((q+(-(+2.3))*(foo1(4.8)+(-foo2(6.9))+(q-(+o)-(-(-1.2)))-(+6.1)+6.1-(-(-1.2))-(+o)*foo2(6.9)))*o))+(+foo0(7.5))-(-3.4)+p)
	print(x3,"\n")

	x4 := (((p+6.1)-(6.1+(+(+2.3))*q)))+q-(foo2(6.9)-(+p)+foo0(7.5)*+foo0(7.5))
	print(x4,"\n")

	x5 := foo2(6.9)*-(foo1(4.8)-o)/(+2.3)+(+(+2.3))-(q/((foo1(4.8)*-foo1(4.8))/(-3.4)*o))+(-3.4)+(+p)-(foo0(7.5)*(6.1/(+2.3))*foo1(4.8)-(-(-1.2))-(+6.1)-foo0(7.5)+foo1(4.8))+foo2(6.9)/+o
	print(x5,"\n")

	x6 := foo0(7.5)-p-(((+2.3)-foo2(6.9))*(foo1(4.8)-(-foo2(6.9))*(-(-1.2))-(-q)/((foo1(4.8)+o+foo2(6.9)-(+p)+(o*foo1(4.8)/foo0(7.5))))))+(-(-(-1.2)))-(-(-1.2))
	print(x6,"\n")

	x7 := (foo0(7.5)*+((-(-1.2))/+(p/((-(-1.2))-foo2(6.9)*-foo0(7.5))-(+foo1(4.8))+o)+(+6.1)/(p/+q+(+o)+(-(-1.2))-(-6.1)-q)+(+2.3))-foo1(4.8)+foo0(7.5))+6.1
	print(x7,"\n")

	x8 := (+2.3)+((o+(+(-3.4))-((-(-1.2))/-(foo0(7.5)/-(((-(-1.2))+(-3.4)+(+(-(-1.2)))+foo1(4.8)-(+(+2.3))-o*(+2.3)*foo0(7.5)-(-o)*foo0(7.5))))*-foo2(6.9))))-(-3.4)*foo0(7.5)+(-foo1(4.8))-(-3.4)-(-(-1.2))
	print(x8,"\n")

	x9 := (foo2(6.9)-(-3.4))-(+(-(-1.2)))-(-3.4)-q/((o/((-(-1.2))+((-(-1.2))-((foo0(7.5)*-foo0(7.5)))-6.1-(-foo0(7.5))-(-(-1.2)))+(-(-3.4))+(-(-1.2)))*-(-(-1.2)))+(-foo0(7.5))-foo2(6.9))-(+foo2(6.9))+p/(+2.3)
	print(x9,"\n")

	x10 := ((q-foo2(6.9)/(-3.4)))-(+q)/((q/-(+2.3))+(+(+2.3))-p)-foo1(4.8)+(-6.1)*(+2.3)+(+foo1(4.8))+(((p+(-foo2(6.9))*(-(-1.2)))*foo2(6.9)+(-(-3.4))*6.1))+foo0(7.5)+foo2(6.9)
	print(x10,"\n")

	x11 := q-(-(-(-1.2)))+(((-(-1.2))+(-(+2.3))+(-(-1.2)))-(+(-(-1.2)))+6.1)+(+6.1)-((foo2(6.9)-(+(-(-1.2)))+((((+2.3)-(+p)-o)+foo0(7.5))/+foo2(6.9)))*(+2.3))-6.1+(-o)-foo0(7.5)
	print(x11,"\n")

	x12 := o+(+p)-o/(((-(-1.2))-(-6.1)*((-(-1.2))*(((+2.3)+q))+(foo1(4.8)*-(-3.4)-(+(-3.4))+o*6.1)/-(+2.3)/foo2(6.9))))
	print(x12,"\n")

	x13 := foo2(6.9)-(+foo0(7.5))*foo2(6.9)*(((-3.4)-q))*(-(-1.2))-(+(+2.3))+(o/-(+2.3)+(-(-(-1.2)))+(-(-1.2)))-(-(-1.2))*foo2(6.9)+(+(-3.4))-foo1(4.8)/-(((+2.3)+(q+(-(-1.2))-(+p)+o)-(+6.1)+foo1(4.8)*(-3.4))+(-foo2(6.9))+p)
	print(x13,"\n")

	x14 := (-(-1.2))+(-3.4)-p+(-q)+(foo2(6.9)*(((q*foo2(6.9))+(-3.4)+(+2.3))*+o-foo0(7.5)+o)/foo0(7.5))*+foo1(4.8)*foo0(7.5)
	print(x14,"\n")

	x15 := p-(+p)-(-(-1.2))*-(((((+2.3)-(+foo0(7.5))*o))+(+foo1(4.8))*foo1(4.8))*(q-foo1(4.8)-(-foo2(6.9))+q-(+foo2(6.9))*6.1)+6.1)+q-(+q)+(-3.4)
	print(x15,"\n")

	x16 := (foo1(4.8)-(-p)*(+2.3)-q)+(-6.1)+foo1(4.8)+(+p)-(foo1(4.8)-(+q)-foo0(7.5))+((6.1-(+p)+((-3.4)+(-o)-(6.1+(-3.4))+(-(+2.3))-foo0(7.5)*-o)*-p)+foo0(7.5))-(+2.3)
	print(x16,"\n")

	x17 := (p+((((q+(+2.3)*(-3.4))*-(((-(-1.2))*foo2(6.9)))*-foo0(7.5))))+(-3.4)*-foo2(6.9))*+foo2(6.9)+6.1/foo2(6.9)
	print(x17,"\n")

	x18 := (6.1+(+6.1)-(((q-(+(-3.4))+(+2.3)+(-p)-o-(+p)/foo0(7.5)))*+((foo0(7.5)+(+(+2.3))-((-(-1.2))-(+q)+(-3.4)+6.1)+(-foo1(4.8))*(-(-1.2))*-foo0(7.5)))-(-3.4)+(+o)*foo1(4.8)+(-o)-o))+(-q)+q+foo0(7.5)
	print(x18,"\n")

	x19 := 6.1/(5.31+(q*foo0(7.5)+(-foo1(4.8))*p*-q)*((((-(-1.2))-o)))*-6.1)*(((-3.4)/-q+(+p)/foo0(7.5)*+6.1-(-foo0(7.5))-(-(-1.2))-p-(+foo2(6.9))-o+(-(-1.2))+(+o)-p+p+(-3.4)+(-(-(-1.2)))-q))*q
	print(x19,"\n")

	x20 := foo1(4.8)/(-3.4)*-((p/-o)+((foo2(6.9)-(-foo0(7.5))-6.1)*+foo1(4.8)*-foo2(6.9)/-foo1(4.8)-(+(-(-1.2)))+(-(-1.2)))-(-o)/foo0(7.5))*+(+2.3)+(-(-(-1.2)))/foo1(4.8)+q
	print(x20,"\n")

	x21 := ((-(-1.2))-((+2.3)-(-3.4)))-((foo1(4.8)*+((-(-1.2))+p)+6.1*-(-(-1.2))+q+(-foo1(4.8))+(+2.3)))-(+o)*foo2(6.9)-(+o)+(-(-1.2))
	print(x21,"\n")

	x22 := ((-3.4)-foo0(7.5)-(+(+2.3))/((-3.4)-(+foo2(6.9))+(((foo1(4.8)-(+o)*(+2.3)+(q-foo2(6.9)-(-q)*q*foo1(4.8))))-(-(+2.3))*(+2.3))-foo1(4.8)/-6.1))+(+(+2.3))+foo1(4.8)
	print(x22,"\n")

	x23 := o-(q-(+6.1)*(((+2.3)+q)-(+2.3)+(-p)+((q*+(foo2(6.9)-(-q)+foo0(7.5))))))
	print(x23,"\n")

	x24 := (((foo1(4.8)-(+q)*(p+6.1+foo0(7.5))-(-foo2(6.9))-(((+2.3)-(-(-1.2))))+p)-(+2.3)*+(+2.3)/foo2(6.9)+(+2.3))+(-(-1.2))*-foo0(7.5)+(+6.1)-foo2(6.9)*-q)/+(+2.3)+(+o)-6.1
	print(x24,"\n")

	x25 := foo0(7.5)*(q-(q+(q*+(-(-1.2))-(+q)*((((foo2(6.9)*(-3.4)))+q)))))
	print(x25,"\n")

	x26 := foo1(4.8)*(-(-1.2))*(foo1(4.8)/+(foo0(7.5)+p)+foo0(7.5))+(+o)/(p*+(-(-1.2))*-(((o*((-3.4)*-foo1(4.8)+(+(-3.4))+foo1(4.8))/-6.1+(+(+2.3))-foo1(4.8)/+(+2.3))*+q)+o-p)-(+(-3.4))-6.1+6.1+foo1(4.8)+o)*foo1(4.8)
	print(x26,"\n")

	x27 := (foo0(7.5)+(((+2.3)-(-foo0(7.5))+(foo1(4.8)+(-(-3.4))-foo1(4.8)-(+2.3)*-(-(-1.2))/-(-3.4)/+6.1+(-foo1(4.8))*o*6.1))+foo1(4.8)))
	print(x27,"\n")

	x28 := (((-(-1.2))-((-3.4)*+(foo1(4.8)+foo0(7.5))-(-(-1.2))))+(-6.1)/(+2.3)+(-(-3.4))-foo2(6.9)*o-foo2(6.9)*foo2(6.9)*+q)*+(+2.3)
	print(x28,"\n")

	x29 := foo0(7.5)-(-foo1(4.8))/(+2.3)*((q+(-6.1)-foo1(4.8)/-((+2.3)+(-p)+(foo1(4.8)/+p)))+o-(+o)+foo1(4.8))+(+foo0(7.5))/6.1+(+foo0(7.5))*(+2.3)
	print(x29,"\n")

	x30 := ((+2.3)+foo0(7.5)*+(-(-1.2))+(-(+2.3))-(((-(-1.2))+(p+p/(6.1+((+2.3)-q)-(+p)+o)+(-(-(-1.2)))-6.1))))+o+foo2(6.9)-(-(-3.4))-foo1(4.8)
	print(x30,"\n")

	x31 := o*foo1(4.8)*((-(-1.2))-(+foo2(6.9))*p)+(+foo0(7.5))+foo1(4.8)-(6.1-(+o)/((foo2(6.9)/-foo2(6.9)+foo1(4.8)+(+2.3))-(-p)-((-3.4)-(+(-3.4))-((6.1*foo0(7.5))+foo2(6.9)-(+q)/foo0(7.5)/(-(-1.2))*-(-(-1.2)))-q+foo1(4.8)+(+(-(-1.2)))+o)+foo2(6.9)-(-3.4)))
	print(x31,"\n")

	x32 := ((((((+2.3)+(+(-(-1.2)))+(-3.4))))))
	print(x32,"\n")

	x33 := foo1(4.8)+(+(-(-1.2)))-((+2.3)*foo2(6.9)+foo0(7.5))-(+p)-((foo0(7.5)*(-3.4)))+(o*-p*foo2(6.9))-(+(-3.4))*foo0(7.5)+(-p)-(foo2(6.9)-(-q)-(foo0(7.5)+(+(-3.4))*foo0(7.5))-(+foo2(6.9))+p)
	print(x33,"\n")

	x34 := q-(+foo0(7.5))+q*-6.1-(-o)*(p-(+foo0(7.5))+(((-(-1.2))-(+(+2.3))+6.1+(-p)*foo0(7.5))-foo2(6.9)-(+o)+((foo1(4.8)/-o/foo2(6.9)-(+o)*foo1(4.8)+q+(-foo2(6.9))-foo0(7.5)+(-o)+foo0(7.5)-(+2.3))))+q/+foo1(4.8)+(+foo0(7.5))*foo1(4.8))+(-q)-6.1+o-(-3.4)
	print(x34,"\n")

	x35 := foo1(4.8)+(-3.4)+(-foo1(4.8))+foo2(6.9)+(p*+(o+(-(-1.2))))*-(-(-1.2))+(-(+2.3))*((-3.4)/+foo1(4.8))/+(-3.4)-(foo1(4.8)*-(foo1(4.8)-(+(-3.4))*(+2.3)-(-(+2.3))-foo0(7.5))+((+2.3)-(-p)-foo2(6.9)))-foo2(6.9)+(+o)-foo0(7.5)-(+(-3.4))+foo2(6.9)+(+p)+p
	print(x35,"\n")

	x36 := (((foo2(6.9)+o)+foo1(4.8)+(6.1/((6.1-(-foo2(6.9))/(foo1(4.8)*6.1))-(-q)*6.1/+(-3.4)-o/(-3.4)))))
	print(x36,"\n")

	x37 := o*+o*(-3.4)*(+2.3)*+o/(-3.4)/+(-(-1.2))+(-foo0(7.5))*(-(-1.2))+(-6.1)*(((6.1+(-(+2.3))*(foo0(7.5)+(foo0(7.5)/((+2.3)*(-3.4))+(+2.3)-(-(-3.4))-q))*+o/-o-(+p)+o+(-(-1.2))*-o))+foo1(4.8))-(-3.4)+foo0(7.5)
	print(x37,"\n")

	x38 := (-3.4)-(+foo2(6.9))-((p*-(((+2.3)*-foo0(7.5)+foo1(4.8)))))*((foo2(6.9)-p))
	print(x38,"\n")

	x39 := (p/-(6.1-((foo1(4.8)/-(foo1(4.8)+(-q)-6.1))+(+p)-(+2.3)+(-6.1)-o)*+((q*q))))+foo1(4.8)/+o
	print(x39,"\n")

	x40 := ((-3.4)*-q-(foo1(4.8)-((-(-1.2))-(-(-3.4))+(+2.3))/+foo1(4.8)))-q-(+foo0(7.5))+foo2(6.9)*p-6.1/+(-(-1.2))+(6.1*+(((-3.4)-6.1+(+2.3)+foo0(7.5)/q/(-3.4)*+foo1(4.8))))
	print(x40,"\n")

	x41 := foo2(6.9)*o+(-(-(-1.2)))+(((+2.3)+(+foo2(6.9))+(((foo2(6.9)-((-3.4)*+(6.1+(-(-3.4))*foo1(4.8))))+q)+foo0(7.5)))/(-(-1.2))+foo2(6.9)*+foo0(7.5))+(-(-1.2))
	print(x41,"\n")

	x42 := (((q+foo0(7.5)))-(-3.4)-(+foo0(7.5))-6.1*(((-3.4)+6.1+(foo1(4.8)/(-3.4)-foo0(7.5)*+(+2.3)/+(-(-1.2)))*(+2.3))+(+p)+p*-(-3.4)))+(+o)+p*foo0(7.5)
	print(x42,"\n")

	x43 := ((((+2.3)-(+q)+(foo0(7.5)+foo1(4.8)-foo0(7.5)+(-foo0(7.5))-foo0(7.5))))-q)
	print(x43,"\n")

	x44 := (6.1-p+(+2.3)*-q-((q+(-foo0(7.5))-(+2.3))+(+6.1)+6.1))-p*-(((((+2.3)/+o*+q)*foo0(7.5))-(-foo0(7.5))*foo1(4.8)-foo2(6.9)+foo0(7.5)*foo2(6.9)*q))
	print(x44,"\n")

	x45 := ((foo0(7.5)+6.1/((-3.4)+(+foo2(6.9))-p-(-p)+(((-(-1.2))+q+(q+(+2.3))+(+p)*foo1(4.8)+o-6.1*foo0(7.5)+p)+(+(-3.4))+6.1)*(-3.4))))+(-foo1(4.8))+foo2(6.9)-(-(-1.2))+(-foo1(4.8))-(foo0(7.5)-(-(-1.2))*foo0(7.5)*-(-(-1.2)))-(+6.1)+(+2.3)+(-(-3.4))-(-(-1.2))
	print(x45,"\n")

	x46 := p-(((((-(-1.2))+(foo1(4.8)*foo1(4.8)*o)+(-o)-p))+(-p)-foo0(7.5)))+(+2.3)
	print(x46,"\n")

	x47 := (q-(+q)/(((p/o)/q)))
	print(x47,"\n")

	x48 := ((((-(-1.2))*(-(-1.2))/((+2.3)-(foo2(6.9)*-foo0(7.5))))))
	print(x48,"\n")

	x49 := (-3.4)-(+p)-((((o-foo2(6.9))-(-(-3.4))+foo0(7.5)+p)*foo2(6.9)))
	print(x49,"\n")

	print(o,p,q,"\n")
}
