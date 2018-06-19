package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Location : ", t.Location(), " Time : ", t) // local time

	location, err := time.LoadLocation("Europe/Zurich")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Location : ", location, " Time : ", t.In(location)) // Europe/Zurich

	locc, _ := time.LoadLocation("Africa/Tunis")
	nowc := time.Now().In(locc)
	fmt.Println("Location : ", locc, " Time : ", nowc) // Africa/Tunis

	locb, _ := time.LoadLocation("Europe/Paris")
	nowb := time.Now().In(locb)
	fmt.Println("Location : ", locb, " Time : ", nowb) // Europe/Paris

	loce3, _ := time.LoadLocation("Europe/Madrid")
	nowe3 := time.Now().In(loce3)
	fmt.Println("Location : ", loce3, " Time : ", nowe3) // Europe/Madrid

	loce1, _ := time.LoadLocation("Europe/Dublin")
	nowe1 := time.Now().In(loce1)
	fmt.Println("Location : ", loce1, " Time : ", nowe1) // Europe/Dublin

	loce2, _ := time.LoadLocation("Europe/London")
	nowe2 := time.Now().In(loce2)
	fmt.Println("Location : ", loce2, " Time : ", nowe2) // Europe/London

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	fmt.Println("Location : ", loc, " Time : ", now) // Asia/Shanghai

	locam, _ := time.LoadLocation("America/Montreal")
	nowam := time.Now().In(locam)
	fmt.Println("Location : ", locam, " Time : ", nowam) // America/Montreal

	locamm, _ := time.LoadLocation("America/New_York")
	nowamm := time.Now().In(locamm)
	fmt.Println("Location : ", locamm, " Time : ", nowamm) // America/New_York

}
