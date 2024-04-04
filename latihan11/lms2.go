package main

import "fmt"

type ride struct {
	name, team    string
	fp1, fp2, fp3 time
	q             string
}
type time struct {
	minute, second int
}

func isiData(pembalap *ride) {
	fmt.Scan(&pembalap.name, &pembalap.team, pembalap.fp1, &pembalap.fp2, &pembalap.fp3)
}

func kualifikasi(pembalap *ride) {
	var temp1, temp2, temp3, temp4, temp5, temp6 int
	isiData(pembalap)
	temp1 = pembalap.fp1.minute * 60
	temp2 = pembalap.fp1.second
	temp3 = pembalap.fp2.minute * 60
	temp4 = pembalap.fp2.second
	temp5 = pembalap.fp3.minute * 60
	temp6 = pembalap.fp3.second
	sum1 := temp1 + temp2
	sum2 := temp3 + temp4
	sum3 := temp5 + temp6
	if sum1 <= 75 && sum2 <= 75 && sum3 <= 75 {
		pembalap.q = "Kualifikasi 2"
	} else if (temp1 == 0 && temp2 == 0) || (temp3 == 0 && temp4 == 0) || (temp5 == 0 && temp6 == 0) {
		pembalap.q = "Diskualifikasi"
	} else if sum1 > 75 && sum2 > 75 && sum3 > 75 {
		pembalap.q = "Kualifikasi 1"
	}

}
func main() {
	var A ride
	isiData(&A)
	kualifikasi(&A)
	fmt.Print(A.q)
}
