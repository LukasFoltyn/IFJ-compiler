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

	x0 := !(((-(-1.2))-(+6.1)/(o*((foo2(6.9)+q-(-q)*(+2.3)-(+p)-q-(-(+2.3))-foo2(6.9)*6.1+foo1(4.8)))/foo0(7.5)-(+p)+o*(+2.3)-(-o)*(+2.3)+p-(+o)/6.1-(-3.4)-(+2.3)+(-p)/q))+(+(+2.3))-o-p >= (-3.4))
	print(x0,"\n")

	x1 := ((o*(foo2(6.9)-(-foo2(6.9))*(+2.3)*(-(-1.2))-(+foo1(4.8))-foo0(7.5)*+o)-(+q)*((foo0(7.5)/(foo2(6.9)-(+q)*(foo1(4.8)-o+(-p)+(+2.3)-q)-(+o)+foo1(4.8)*-foo1(4.8)+q-(-3.4)))+(+2.3)-(-o)-(-3.4)/o)+foo2(6.9)))*+o <= q
	print(x1,"\n")

	x2 := ((-(-1.2))+foo1(4.8)) > (foo1(4.8)*q-6.1+((q+(foo2(6.9)-(-o)-foo2(6.9)))))+6.1
	print(x2,"\n")

	x3 := q*((((foo0(7.5)+(-q)+foo0(7.5)-p))-o*q*6.1+foo0(7.5))) >= (-(-1.2)) || !(p <= o) || !(q/(-3.4) != foo2(6.9)) && !((-3.4) != p*-(-3.4))
	print(x3,"\n")

	x4 := !(q*((+2.3)*(+2.3))+(foo0(7.5)*-(6.1-((+2.3)*+(foo0(7.5)-6.1)*-p/+(p-((-(-1.2))+(-foo1(4.8))+foo2(6.9))*-(-(-1.2)))+(+6.1)+6.1-(-foo0(7.5))+foo2(6.9)/6.1))*(-(-1.2))+(+(-(-1.2)))+(-(-1.2))+(+foo1(4.8))-foo0(7.5)-(-p)+foo2(6.9)/(+2.3)) >= o) || !(6.1 >= foo0(7.5))
	print(x4,"\n")

	x5 := (((+2.3)+o)) != (+2.3)-(o-(+(+2.3))-(-(-1.2))) || (-(-1.2)) > foo2(6.9)*-(6.1*(foo1(4.8)-(foo0(7.5)+(-foo0(7.5))-foo2(6.9)))*+6.1-foo0(7.5)*+p)+(-foo2(6.9))/(-3.4)-(+6.1)-foo0(7.5)
	print(x5,"\n")

	x6 := (((o*+6.1)*+(p*q+(+foo1(4.8))-foo1(4.8)+(+o)+((foo2(6.9)-(+foo1(4.8))-(6.1+p+(-(+2.3))-(-(-1.2)))+(+2.3))*6.1)+foo0(7.5)-(+q)*(+2.3))*foo1(4.8)*-p)*foo0(7.5)+(-q)-o) > 6.1
	print(x6,"\n")

	x7 := !(q-(((((+2.3)*6.1+(+o)+(-3.4)-(-foo1(4.8))-(-3.4)))+(-3.4))) < (+2.3))
	print(x7,"\n")

	x8 := !!!(((+2.3)-((o+(-o)*foo1(4.8)+(+p)-(6.1+(+(-(-1.2)))+q))))*-(foo2(6.9)-foo0(7.5)-(+p)*foo2(6.9)+(-o)/(-(-1.2))+o*o+(-(-1.2))-foo1(4.8)+(+p)+q*6.1)/(-3.4)+(+(-3.4))*p > (-3.4))
	print(x8,"\n")

	x9 := !((foo0(7.5)*-((((+2.3)-(+6.1)-o-foo1(4.8))-(+6.1)/o)))+q+(+(-3.4))+q < 6.1) && !(o+(-(-3.4))-6.1 == (-3.4))
	print(x9,"\n")

	x10 := !((((+2.3)-(+q)*((o+(+foo0(7.5))/q+foo2(6.9)-(+6.1)+foo0(7.5))+foo0(7.5))-(-3.4))-(-q)-(+2.3)) == foo1(4.8)/+6.1)
	print(x10,"\n")

	x11 := !(p-((o*+foo2(6.9))*-o) > p+(-(-3.4))-foo1(4.8)*foo0(7.5)) || foo0(7.5)/+(((-(-1.2))*q-(+p)*((6.1-((-3.4)+(+p)-(+2.3)))-foo2(6.9)-(-q)+(-(-1.2))/-foo0(7.5)*-foo2(6.9)*-q/q))*p)-foo1(4.8)*q-(-foo2(6.9))+foo0(7.5)/+foo0(7.5) == 6.1
	print(x11,"\n")

	x12 := 6.1 <= (-3.4) || p > foo1(4.8)-(-foo1(4.8))*q || ((p-(+q)+foo1(4.8)*-(+2.3)-(+p)+(6.1-6.1*+6.1+(+foo1(4.8))+(+2.3)-((foo1(4.8)-((-3.4)+(-(-(-1.2)))*foo1(4.8)+(+foo1(4.8))/6.1-foo1(4.8)))/+foo0(7.5)-(-o)+6.1))*+foo2(6.9))+(-(-3.4))*p+(-o)-foo0(7.5)*-foo1(4.8)) >= p-(+(+2.3))+(-3.4) || foo1(4.8) == (-(-1.2))/(+2.3)
	print(x12,"\n")

	x13 := !!((foo1(4.8)*+o) == q*o) || ((p+(((foo0(7.5)*(-(-1.2)))-(-foo0(7.5))-(-(-1.2))+(-foo2(6.9))+q))+(q+6.1))) >= o+(-o)+(-3.4)-o/-6.1 || foo1(4.8)-(+(+2.3))+q*-o > q+foo1(4.8)
	print(x13,"\n")

	x14 := !!!((-(-1.2)) < 6.1) && !(((-(-1.2))*q) == (((+2.3)+(o*p+(6.1*foo2(6.9)+(+6.1)*(foo0(7.5)+(+o)-foo0(7.5)*+(-(-1.2))-(-(+2.3))-q-6.1-(-6.1)*o*+foo1(4.8)+foo0(7.5)))))))
	print(x14,"\n")

	x15 := !((foo1(4.8)+(+6.1)*(q/foo2(6.9))*((o-q*+o/+foo1(4.8))+(-foo2(6.9))/foo2(6.9))) == foo1(4.8))
	print(x15,"\n")

	x16 := !((-(-1.2)) > (+2.3)-((foo2(6.9)*-(6.1-(-(-(-1.2)))-foo2(6.9)/(q-(+(-3.4))*q))-p))) && !((+2.3)-(-(-(-1.2)))+(-(-1.2)) >= q)
	print(x16,"\n")

	x17 := o <= ((-(-1.2))*6.1) && p <= ((-(-1.2))*foo0(7.5)+(-foo2(6.9))-p)*(-(-1.2))*+(-(-1.2))+(+(+2.3))/p/o+(((+2.3)/+foo1(4.8))-(+foo2(6.9))+foo1(4.8))
	print(x17,"\n")

	x18 := !(foo1(4.8) == 6.1) && (-(-1.2)) > foo1(4.8)+(-6.1)/(+2.3) && !((+2.3) <= ((-3.4)-((q*+foo2(6.9)))*+((-(-1.2))+(-(-1.2))))-(+p)-(foo2(6.9)+(o-(+o)-p)-foo2(6.9)/q)) && (-(-1.2))+(+foo2(6.9))-foo1(4.8) <= foo0(7.5)-o+foo2(6.9)-(+p)*o
	print(x18,"\n")

	x19 := ((+2.3)+o) < o+(+6.1)-6.1 || !(q != ((((+2.3)+(-(+2.3))*(-(-1.2))+(-(+2.3))-foo1(4.8)-foo2(6.9))))-(-3.4)) || foo1(4.8) != (-(-1.2))
	print(x19,"\n")

	x20 := !!((-(-1.2))*(q+(+6.1)-((-3.4)-(+q)*(((-(-1.2))-foo1(4.8))-(-(-3.4))-(((+2.3)-(-3.4))+(-3.4))+foo0(7.5)+(-6.1)-foo0(7.5))/6.1)) < (-3.4))
	print(x20,"\n")

	x21 := (+2.3)+foo0(7.5) >= foo1(4.8) && (((6.1/((6.1+(+(-3.4))+p))+q))-(+2.3)-6.1)+(-(+2.3))/q > q
	print(x21,"\n")

	x22 := ((((((6.1*6.1))-(+foo2(6.9))-(-(-1.2)))+(-6.1)-foo2(6.9))+(-(-1.2))/foo1(4.8)/q+(+2.3))) <= (-(-1.2)) && !(foo2(6.9) < (-(-1.2)))
	print(x22,"\n")

	x23 := (foo1(4.8)-(foo2(6.9)+(+(+2.3))*foo0(7.5)/((-(-1.2))*-(6.1+(-o)-foo2(6.9)))-foo2(6.9)+(-q)-6.1+o))/foo2(6.9)+p+(+(+2.3))-foo2(6.9) > foo1(4.8)
	print(x23,"\n")

	x24 := p*foo2(6.9) <= (p+(((((-3.4)-(+(+2.3))-(+2.3)))-p+(-o)*(+2.3)))+(-3.4)-(+(-3.4))+o)-(+(-(-1.2)))-foo2(6.9)
	print(x24,"\n")

	x25 := ((foo1(4.8)*+((((foo0(7.5)-(+(+2.3))+(+2.3)-(-3.4))-(+(-3.4))+foo0(7.5))-foo2(6.9))-o)-(+2.3))+(+foo0(7.5))/p/+foo2(6.9)/-q) != p || !(foo1(4.8) > (-(-1.2)))
	print(x25,"\n")

	x26 := !(foo2(6.9)-foo0(7.5) < ((foo0(7.5)-(-(-(-1.2)))+(-3.4)))) || !!!(foo0(7.5)*o < (-3.4)-p*((+2.3)-(((+2.3)-(6.1+((-3.4)*-foo0(7.5))*foo1(4.8))))-(+foo2(6.9))*(-3.4))) && !(foo0(7.5)*-(-3.4) > p)
	print(x26,"\n")

	x27 := !!((q+p-(+foo0(7.5))-((+2.3)*q)) < (((((+2.3)-(-(-1.2))+(-(+2.3))-(-(-1.2))))*-o)))
	print(x27,"\n")

	x28 := (((q*((foo2(6.9)*-q)*(+2.3)/6.1-(-foo0(7.5))-o-6.1+(+o)/foo1(4.8)))+(-3.4))+(-(-3.4))*(-3.4)) <= (+2.3) && !!!!((-(-1.2)) <= (-3.4))
	print(x28,"\n")

	x29 := !((foo2(6.9)+(+foo2(6.9))*(p+(-foo2(6.9))-6.1-(((6.1-q+(+2.3)))+(+q)/(-(-1.2))+(+(-(-1.2)))-p))*-((+2.3)+(6.1-(-3.4)+(+p)*6.1)-(+foo1(4.8))+o)+(-foo2(6.9))*q*foo0(7.5))+p <= p)
	print(x29,"\n")

	x30 := !(((p-(+(-(-1.2)))+(+2.3))) <= (6.1*o/+6.1)) || !((-(-1.2))+((-(-1.2))+(+(-(-1.2)))+p+(-(-1.2))) > q+(+foo0(7.5))+(+2.3)*-foo0(7.5)-(-6.1)/(-(-1.2))) && !!!!((+2.3)/((-(-1.2))-(+(-3.4))*(-(-1.2))-(+foo1(4.8))+foo0(7.5)-6.1)-(+(-3.4))+(-3.4)-(+(-(-1.2)))*(+2.3) <= foo2(6.9)/-foo2(6.9))
	print(x30,"\n")

	x31 := (p+(+(-3.4))+(-(-1.2))+(-foo1(4.8))+6.1-p*(o+(-(-1.2))+(-o)+(((foo1(4.8)-(+q)-(foo0(7.5)+((-3.4)-q-foo2(6.9)+(-q)-(-(-1.2))-q+foo1(4.8)))*foo2(6.9))))-(+o)-(+2.3)))-(+(+2.3))+o < o
	print(x31,"\n")

	x32 := (((((-3.4)*+p+(-foo1(4.8))*p)+(-foo0(7.5))-(-(-1.2))+(-3.4))*-q)) >= o
	print(x32,"\n")

	x33 := o < p*(o+(-foo1(4.8))+(((foo2(6.9)/(foo2(6.9)-6.1/+(6.1*(+2.3)))+(-foo1(4.8))*(p-foo2(6.9)+foo1(4.8))))))+(+6.1)+(+2.3)
	print(x33,"\n")

	x34 := (-3.4)+(-o)*(((-(-1.2))-p)+((-(-1.2))*(q+(+foo2(6.9))+foo0(7.5)/p+(+6.1)*(-(-1.2))+(-3.4))-(-foo2(6.9))*(-3.4)))-(-q)+(-(-1.2))/q == p
	print(x34,"\n")

	x35 := ((+2.3)-(+foo0(7.5))/(((((+2.3)/6.1)-(-foo1(4.8))-((o/6.1-6.1)+(-6.1)+q))+(+(-(-1.2)))-(+2.3))*-(-(-1.2))-q)) <= o
	print(x35,"\n")

	x36 := (p-(-foo0(7.5))-6.1)*(((foo2(6.9)+(-(-1.2)))*6.1*(-(-1.2))+(o+p)*6.1+((-3.4)*(-(-1.2)))/-(-(-1.2))-(+(-3.4))/6.1))+(-(+2.3))*foo0(7.5) == ((+2.3)*(-(-1.2))/+q*-foo0(7.5)-foo1(4.8)/-(-3.4))
	print(x36,"\n")

	x37 := !((p+((((-3.4)*foo2(6.9)-(-foo1(4.8))-((q*-((-3.4)*-o))))/-(+2.3)+(-6.1)*foo1(4.8)))-6.1-q) != (+2.3)) && foo1(4.8)-foo1(4.8)*(-(-1.2))/-(-(-1.2)) != foo2(6.9)
	print(x37,"\n")

	x38 := !!!!(o < ((p+foo2(6.9))-(+foo1(4.8))-foo1(4.8))+(+q)*6.1) || (q-(+6.1)/foo0(7.5)) < (q-(-foo1(4.8))+6.1) || foo1(4.8) >= foo2(6.9)
	print(x38,"\n")

	x39 := (6.1/foo2(6.9)-(foo2(6.9)-(+foo0(7.5))*foo0(7.5)))+p*(o+(+foo2(6.9))-(o*+q+(+(-(-1.2)))-q))/+(foo0(7.5)+(-p)-(-3.4))+(-(-(-1.2)))/q >= p
	print(x39,"\n")

	x40 := ((foo2(6.9)+q+(+p)-((foo2(6.9)+(+(+2.3))*foo2(6.9)))*-(-(-1.2))))/(-3.4)+(+p)*(-(-1.2)) == 6.1/q
	print(x40,"\n")

	x41 := (p+p) <= o && (-3.4)*-((+2.3)*p*foo2(6.9)+(-foo1(4.8))-foo0(7.5))-(-foo2(6.9))*((foo2(6.9)/q+(p+(foo0(7.5)+(+(-(-1.2)))/foo1(4.8)))))-foo0(7.5) < o
	print(x41,"\n")

	x42 := ((p-foo1(4.8)-(-foo0(7.5))*(-3.4)-(-(-1.2))))*-6.1 < p || !!((foo1(4.8)+foo0(7.5))+(+o)-foo1(4.8)/+(((foo1(4.8)-(-p)+6.1*+(-3.4))-foo2(6.9))) > 6.1)
	print(x42,"\n")

	x43 := !!((p+(+(-3.4))+6.1)-(+6.1)+o*(((p-o-(+foo2(6.9))*6.1/(+2.3)*-p))-foo1(4.8)-(+(+2.3))*o) > o)
	print(x43,"\n")

	x44 := (foo2(6.9)/((+2.3)-(-(+2.3))*(p+(-6.1)+(+2.3)+(+6.1)-o))+(-foo2(6.9))*(o-(+foo1(4.8))*(p-(-foo2(6.9))-o/q-(+foo2(6.9))-o-(+foo1(4.8))+p))+(+(-(-1.2)))+6.1-(-o)*foo1(4.8)-(+foo1(4.8))-6.1) != foo1(4.8)
	print(x44,"\n")

	x45 := !!(o+foo2(6.9) != p) && !!((((p+(+o)*(6.1*(foo2(6.9)+foo2(6.9)+foo2(6.9)-(+(+2.3))+foo2(6.9)))-(+(+2.3))*foo2(6.9)))+(-6.1)*p+(+2.3)) == 6.1)
	print(x45,"\n")

	x46 := !(q-((+2.3)*((6.1+(-6.1)+(-(-1.2))/+6.1))+(+foo0(7.5))-o)+(q*-(q*p-(+q)-foo1(4.8)+foo0(7.5))) != o)
	print(x46,"\n")

	x47 := !(p-foo1(4.8)+(-(+2.3))+(6.1/-((foo0(7.5)+foo2(6.9)+(+(+2.3))*(((p-foo2(6.9))))-(foo0(7.5)/-(-(-1.2))-foo1(4.8)))))-(-o)+6.1*-foo0(7.5)*(+2.3)*+6.1 > 6.1)
	print(x47,"\n")

	x48 := (-3.4)-(-foo1(4.8))*p*(-3.4) < (+2.3)-(-o)/(+2.3)-(p+(-foo1(4.8))-(-(-1.2)))+6.1 || (foo1(4.8)-((-(-1.2))*(foo1(4.8)-6.1-(-3.4)-(+foo0(7.5))/(+2.3)-(-(-3.4))-p+(-foo0(7.5))*o)-(+foo1(4.8))+(+2.3))+(+foo1(4.8))+(+2.3)-(+o)-o-(+o)/o*-6.1-(+2.3)-(-q)-(+2.3)*(+2.3)) <= p
	print(x48,"\n")

	x49 := foo0(7.5) >= 6.1-((q-(-(-1.2)))-(-6.1)+foo2(6.9)-p-((-3.4)/+((o-(-p)+(foo1(4.8)-foo0(7.5)*(-3.4)-(-3.4)+(+(-3.4))-foo0(7.5)/+(-3.4)*o/(-3.4))))*+o-p*+foo0(7.5)*(+2.3))+foo2(6.9))
	print(x49,"\n")

	print(o,p,q,"\n")
}