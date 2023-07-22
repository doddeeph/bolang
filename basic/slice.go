package main

import "fmt"

/*
https://dasarpemrogramangolang.novalagung.com/A-slice.html

Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya,
jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice.
*/
func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array
	fmt.Println(fruitsA, fruitsB, fruitsC)

	var newFruits = fruits[0:2]
	fmt.Println(newFruits)

	var newFruits1 = fruits[:]
	fmt.Println(newFruits1)

	var newFruits2 = fruits[2:]
	fmt.Println(newFruits2)

	var newFruits3 = fruits[:2]
	fmt.Println(newFruits3)

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var bbFruits = bFruits[0:1]
	fmt.Println(aFruits, bFruits, aaFruits, bbFruits)

	// Buah "grape" diubah menjadi "pinnaple"
	bbFruits[0] = "pinnaple"
	fmt.Println(aFruits, bFruits, aaFruits, bbFruits)

	fmt.Println(len(fruits), cap(fruits))
	fmt.Println(len(aFruits), cap(aFruits))
	fmt.Println(len(bFruits), cap(bFruits))

	/*
		Kode	      Output	             len()	cap()
		fruits[0:4]	  [buah buah buah buah]	  4	     4
		aFruits[0:3]  [buah buah buah ----]	  3	     4
		bFruits[1:4]  ---- [buah buah buah]	  3	     3
	*/

	fmt.Println("========== append() ==========")

	var cFruits = append(fruits, "papaya")
	fmt.Println(cFruits)

	var dFruits = fruits[0:2]
	fmt.Println(len(dFruits), cap(dFruits)) //len(fruits) < cap(fruits)
	fmt.Println(fruits, dFruits)

	var eFruits = append(dFruits, "papaya")
	fmt.Println(fruits, dFruits, eFruits)

	/*
			Ada 3 hal yang perlu diketahui dalam penggunaan fungsi ini.
			- Ketika jumlah elemen dan lebar kapasitas adalah sama (len(fruits) == cap(fruits)),
		      maka elemen baru hasil append() merupakan referensi baru.
			- Ketika jumlah elemen lebih kecil dibanding kapasitas (len(fruits) < cap(fruits)),
		      elemen baru tersebut ditempatkan ke dalam cakupan kapasitas, menjadikan semua elemen slice lain yang referensi-nya sama akan berubah nilainya.
	*/

	fmt.Println("========== copy() ==========")

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst, src, n)

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pinnaple"}
	n = copy(dst, src)
	fmt.Println(dst, src, n)

	fmt.Println("==========  Pengaksesan Elemen Slice Dengan 3 Indeks ==========  ")
	aFruits = fruits[0:2]
	bFruits = fruits[0:2:2]
	fmt.Println(fruits, len(fruits), cap(fruits))
	fmt.Println(aFruits, len(aFruits), cap(aFruits))
	fmt.Println(bFruits, len(bFruits), cap(bFruits))
}
