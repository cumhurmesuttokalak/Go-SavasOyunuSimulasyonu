package main

import (
	"fmt"
	"math/rand"
	"oyunmodulu/oyun"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var takim int
	fmt.Println("Takımlar kaç kişilik olacak?")
	fmt.Scan(&takim)
	dunyaliSlc := []oyun.Dunyali{}
	marsliSlc := []oyun.Marsli{}
	for i := 0; i < takim; i++ {
		dunyali := new(oyun.Dunyali)
		dunyali.Birlik = "Sakarya"
		dunyali.Can = 100
		dunyali.MermiSayisi = 90
		dunyali.Name = "Dunyali" + strconv.Itoa(i+1) + " = "
		dunyaliSlc = append(dunyaliSlc, *dunyali)

		marsli := new(oyun.Marsli)
		marsli.UzayGemisi = "XYZ-5454"
		marsli.Can = 100
		marsli.MermiSayisi = 80
		marsli.Name = "Marsli" + strconv.Itoa(i+1) + " = "
		marsliSlc = append(marsliSlc, *marsli)
	}

	for {
		kontrolD := true
		kontrolM := true

		for i := 0; i < len(dunyaliSlc); i++ {
			if dunyaliSlc[i].Can > 0 {
				kontrolD = false
			}
		}
		if kontrolD {
			fmt.Println("Marslılar kazandı")
			break
		} else {
			for i := 0; i < len(marsliSlc); i++ {
				if marsliSlc[i].Can > 0 {
					kontrolM = false
				}
			}
			if kontrolM {
				fmt.Println("Dünyalılar kazandı")
				break
			}
		}
		random := rand.Intn(2) // enum değeri gibi düşün 0 ya da 1 döndürür
		randomForSlc := rand.Intn(len(dunyaliSlc))
		if random == 0 {
			dunyaliSlc[randomForSlc].AtesEt()
			marsliSlc[randomForSlc].HasarAl()
		} else {
			marsliSlc[randomForSlc].AtesEt()
			dunyaliSlc[randomForSlc].HasarAl()
		}
	}
	fmt.Println(dunyaliSlc)
	fmt.Println(marsliSlc)
}
