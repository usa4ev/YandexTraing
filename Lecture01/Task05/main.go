package main

import (
	"fmt"
	"os"
)

func main() {
	var k1, p1, n1, m, k2, p2, n2, z int

	f, _ := os.Open("input.txt")
	fmt.Fscanf(f, "%d %d %d %d %d", &k1, &m, &k2, &p2, &n2)

	var valid bool
	if !(p2 == 1 && n2 <= 2 && k1 != k2){
		if p2 == 1 && n2 == 1{
			z = 1
		}else{
			z = k2 / (n2 + ((p2 - 1) * m) - 1)
		}
		

		valid = validate(m, n2, k2, p2 ,z)
	}
	
	

	switch {
	case p2 == 1 && n2 <= 2 && k1 != k2:
		if k1 < k2{
			p1 = 1
		}else{
			p1 = 0
		}
		
		if m == 1{
			n1 = 1
		}else{
			n1 = 0
		}
	case !valid:
		p1, n1 = -1, -1
	case k1 == k2:
		p1, n1 = p2, n2
	case k1 > k2:
		d := k1 - k2

		dp := d / (z * m)

		p1 = p2 + dp

		// квартира может быть в соседнем подъезде
		if dp == 0 && k1 > (m * p2 * z){
			p1++
		}

		n1 = m - (((p1 * m * z) - k1) / z)

	case k1 < k2:
		d := k2 - k1
		
		dp := d / (z * m)

		p1 = p2 - dp

		// квартира может быть в соседнем подъезде
		if dp == 0 && k1 < (m * (p2 - 1) * z){
			p1--
		}

		n1 = m - ((p1 * m * z) - k1) / z
	}

	f,_ = os.Create("output.txt")
	fmt.Fprintf(f, "%d %d", p1, n1)
}


func validate(m, n, k, p, z int) bool{
	// номер квартиры не бьется по этажу и подъезду
	if k <= ((p - 1) * m * z) + ((n - 1) * z) || 
		k > ((p - 1) * m * z) + (n * z){
			return false
		}

	return true
}