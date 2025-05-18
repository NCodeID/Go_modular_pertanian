package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Penanaman struct {
	id                        int
	namatanaman, jenistanaman string
	lahan, jumlahBibit        int
}

type Tanaman struct {
	id                 int
	nama, jenis, grade string
}

type HasilPanen struct {
	id       int
	nama     string
	jenis    string
	jumlahKg float64
	grade    string
}

var arrTanaman = []Tanaman{}
var arrPenanaman = []Penanaman{}
var arrHasilPanen = []HasilPanen{}
var reader = bufio.NewReader(os.Stdin)

func main() {
	inisiasiDataAwal()
	for {
		bubbleSort(&arrPenanaman,&arrTanaman,&arrHasilPanen)
		var Menu int
		menu()

		fmt.Print("> ")
		fmt.Scanln(&Menu)

		switch Menu {
		case 1:
			fmt.Println(" ")
			fmt.Println("------ Menu Penanaman ------")
			subMenu()
			pilihMenu(Menu)
		case 2:
			fmt.Println(" ")
			fmt.Println("------ Menu Tanaman ------")
			subMenu()
			pilihMenu(Menu)
		case 3:
			fmt.Println(" ")
			fmt.Println("------ Menu Hasil Panen ------")
			subMenu()
			pilihMenu(Menu)
		case 4:
			return
		}

	}
}

func menu() {
	fmt.Println("------ Menu ------")
	fmt.Println("1. Penanaman")
	fmt.Println("2. Tanaman")
	fmt.Println("3. Hasil Panen")
	fmt.Println("4. Keluar")
}
func subMenu() {
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
	fmt.Println("4. Tampilkan")
	fmt.Println("5. Pencarian(Search)")
	fmt.Println("6. kembali")
}

func pilihMenu(n int) {
	var userInput int
	fmt.Print("> ")
	fmt.Scanln(&userInput)
	switch userInput {
	case 1:
		switch n {
		case 1:
			tambah(&arrPenanaman)
		case 2:
			tambah(&arrTanaman)
		case 3:
			tambah(&arrHasilPanen)
		}
	case 2:
		edit()
	case 3:
		hapus()
	case 4:
		switch n {
		case 1:
			tampilkan(&arrPenanaman)
		case 2:
			tampilkan(&arrTanaman)
		case 3:
			tampilkan(&arrHasilPanen)
		}
	case 5:
		// search()
	case 6:
		return

	}

}
func tambah(data interface{}) {
	switch v := data.(type) {
	case *[]Penanaman:
		var p Penanaman
		p.id = len(*v) + 1
		fmt.Print("Masukkan nama tanaman: ")
		p.namatanaman, _ = reader.ReadString('\n')
		p.namatanaman = strings.TrimSpace(p.namatanaman)
		fmt.Print("Masukkan jenis tanaman : ")
		p.jenistanaman, _ = reader.ReadString('\n')
		p.jenistanaman = strings.TrimSpace(p.jenistanaman)
		fmt.Print("Masukkan lahan: ")
		fmt.Scanln(&p.lahan)
		fmt.Print("Masukkan jumlah bibit: ")
		fmt.Scanln(&p.jumlahBibit)
		*v = append(*v, p)
	case *[]Tanaman:
		var t Tanaman
		t.id = len(*v) + 1
		fmt.Print("Masukkan nama tanaman: ")
		t.nama, _ = reader.ReadString('\n')
		t.nama = strings.TrimSpace(t.nama)
		fmt.Print("Masukkan jenis tanaman : ")
		t.jenis, _ = reader.ReadString('\n')
		t.jenis = strings.TrimSpace(t.jenis)
		fmt.Print("Masukkan grade: ")
		fmt.Scanln(&t.grade)
		*v = append(*v, t)

	case *[]HasilPanen:
		var k HasilPanen
		k.id = len(*v) + 1
		fmt.Print("Masukkan nama tanaman: ")
		k.nama, _ = reader.ReadString('\n')
		k.nama = strings.TrimSpace(k.nama)
		fmt.Print("Masukkan jenis tanaman : ")
		k.jenis, _ = reader.ReadString('\n')
		k.jenis = strings.TrimSpace(k.jenis)
		fmt.Print("Masukkan Jumlah(kg) : ")
		fmt.Scanln(&k.jumlahKg)
		fmt.Print("Masukkan grade: ")
		fmt.Scanln(&k.grade)
		*v = append(*v, k)
	}
}

func edit() {

}

func hapus() {

}
func search() {

}

func bubbleSort(arrPenanaman *[]Penanaman, arrTanaman *[]Tanaman, arrHasilPanen *[]HasilPanen, ) {
	x := len(*arrPenanaman)
	for i := 0; i < x-1; i++ {
		tukar := false
		for j := 0; j < x-i-1; j++ {
			if (*arrPenanaman)[j].namatanaman > (*arrPenanaman)[j+1].namatanaman {
				(*arrPenanaman)[j], (*arrPenanaman)[j+1] = (*arrPenanaman)[j+1], (*arrPenanaman)[j]
				tukar = true
			}
		}
		if !tukar {
			break
		}
	}
	y := len(*arrTanaman)
	for i := 0; i < y-1; i++ {
		tukar := false
		for j := 0; j < y-i-1; j++ {
			if (*arrTanaman)[j].nama > (*arrTanaman)[j+1].nama {
				(*arrTanaman)[j], (*arrTanaman)[j+1] = (*arrTanaman)[j+1], (*arrTanaman)[j]
				tukar = true
			}
		}
		if !tukar {
			break
		}
	}
	z := len(*arrHasilPanen)
	for i := 0; i < z-1; i++ {
		tukar := false
		for j := 0; j < z-i-1; j++ {
			if (*arrHasilPanen)[j].nama > (*arrHasilPanen)[j+1].nama {
				(*arrHasilPanen)[j], (*arrHasilPanen)[j+1] = (*arrHasilPanen)[j+1], (*arrHasilPanen)[j]
				tukar = true
			}
		}
		if !tukar {
			break
		}
	}
}

func tampilkan(data interface{}) {
	switch v := data.(type) {
	case *[]Penanaman:
		fmt.Println(" ")
		for _, nilai := range *v {
			fmt.Printf("id : %d, nama: %s, jenis: %s, lahan: %d, jumlah bibit: %d\n", nilai.id, nilai.namatanaman, nilai.jenistanaman, nilai.lahan, nilai.jumlahBibit)
		}
		fmt.Println(" ")
	case *[]Tanaman:
		fmt.Println(" ")
		for _, nilai := range *v {
			fmt.Printf("id : %d, nama: %s, jenis: %s, grade: %s\n", nilai.id, nilai.nama, nilai.jenis, nilai.grade)
		}
		fmt.Println(" ")
	case *[]HasilPanen:
		fmt.Println(" ")
		for _, nilai := range *v {
			fmt.Printf("id : %d, nama: %s, jenis: %s, jumlah kg: %.2f, grade: %s\n", nilai.id, nilai.nama, nilai.jenis, nilai.jumlahKg, nilai.grade)
		}
		fmt.Println(" ")
	}
}

func inisiasiDataAwal() {
	arrTanaman = []Tanaman{
		{1, "Jagung", "Jagung Manis", "A"},
		{2, "Tomat", "Tomat Cherry", "A"},
		{3, "Tomat", "Tomat Sayur", "B"},
		{4, "Cabai", "Cabai Rawit", "A"},
		{5, "Cabai", "Cabai Merah Besar", "A"},
		{6, "Padi", "Padi IR64", "A"},
		{7, "Padi", "Padi Inpari 32", "B"},
		{8, "Wortel", "Wortel Lokal", "B"},
		{9, "Kentang", "Kentang Granola", "A"},
		{10, "Bayam", "Bayam Hijau", "B"},
		{11, "Bayam", "Bayam Merah", "B"},
		{12, "Terong", "Terong Ungu", "B"},
		{13, "Timun", "Timun Suri", "A"},
		{14, "Melon", "Melon Sky Rocket", "A"},
		{15, "Semangka", "Semangka Inul", "A"},
		{16, "Kacang", "Kacang Tanah Takar 1", "B"},
		{17, "Kedelai", "Kedelai Anjasmoro", "A"},
		{18, "Bawang", "Bawang Merah Bima", "B"},
		{19, "Bawang", "Bawang Putih Lumbu Kuning", "A"},
		{20, "Nanas", "Nanas Queen", "A"},
	}
	arrPenanaman = []Penanaman{
		{1, "Jagung", "Jagung Manis", 1, 150},
		{2, "Tomat", "Tomat Cherry", 1, 120},
		{3, "Tomat", "Tomat Sayur", 2, 100},
		{4, "Cabai", "Cabai Rawit", 2, 200},
		{5, "Cabai", "Cabai Merah Besar", 3, 180},
		{6, "Padi", "Padi IR64", 3, 300},
		{7, "Padi", "Padi Inpari 32", 4, 250},
		{8, "Wortel", "Wortel Lokal", 4, 170},
		{9, "Kentang", "Kentang Granola", 5, 190},
		{10, "Bayam", "Bayam Hijau", 5, 130},
		{11, "Bayam", "Bayam Merah", 6, 100},
		{12, "Terong", "Terong Ungu", 6, 110},
		{13, "Timun", "Timun Suri", 7, 140},
		{14, "Melon", "Melon Sky Rocket", 7, 160},
		{15, "Semangka", "Semangka Inul", 8, 170},
		{16, "Kacang", "Kacang Tanah Takar 1", 8, 180},
		{17, "Kedelai", "Kedelai Anjasmoro", 9, 150},
		{18, "Bawang", "Bawang Merah Bima", 9, 200},
		{19, "Bawang", "Bawang Putih Lumbu Kuning", 10, 90},
		{20, "Nanas", "Nanas Queen", 10, 100},
	}
	arrHasilPanen = []HasilPanen{
		{1, "Jagung", "Jagung Manis", 320.5, "A"},
		{2, "Tomat", "Tomat Cherry", 250.0, "A"},
		{3, "Tomat", "Tomat Sayur", 230.0, "B"},
		{4, "Cabai", "Cabai Rawit", 180.0, "A"},
		{5, "Cabai", "Cabai Merah Besar", 170.0, "B"},
		{6, "Padi", "Padi IR64", 750.0, "A"},
		{7, "Padi", "Padi Inpari 32", 680.0, "B"},
		{8, "Wortel", "Wortel Lokal", 210.0, "B"},
		{9, "Kentang", "Kentang Granola", 400.0, "A"},
		{10, "Bayam", "Bayam Hijau", 160.0, "B"},
		{11, "Bayam", "Bayam Merah", 150.0, "B"},
		{12, "Terong", "Terong Ungu", 180.0, "B"},
		{13, "Timun", "Timun Suri", 200.0, "A"},
		{14, "Melon", "Melon Sky Rocket", 220.0, "A"},
		{15, "Semangka", "Semangka Inul", 240.0, "A"},
		{16, "Kacang", "Kacang Tanah Takar 1", 190.0, "B"},
		{17, "Kedelai", "Kedelai Anjasmoro", 260.0, "A"},
		{18, "Bawang", "Bawang Merah Bima", 270.0, "B"},
		{19, "Bawang", "Bawang Putih Lumbu Kuning", 280.0, "A"},
		{20, "Nanas", "Nanas Queen", 300.0, "A"},
	}

}
