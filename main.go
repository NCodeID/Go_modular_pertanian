package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

type RiwayatPanen struct {
	id             int
	nama, jenis    string
	jumlahKg       float64
	grade, tanggal string
}

var arrTanaman = []Tanaman{}
var arrPenanaman = []Penanaman{}
var arrHasilPanen = []HasilPanen{}
var arrRiwayatPanen = []RiwayatPanen{}
var reader = bufio.NewReader(os.Stdin)

func main() {
	inisiasiDataAwal()
	for {
		bubbleSort(&arrPenanaman, &arrTanaman, &arrHasilPanen)
		insertionSort(&arrRiwayatPanen)
		var Menu int
		menu()
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		Menu, _ = strconv.Atoi(input)

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
			fmt.Println(" ")
			fmt.Println("------ Menu Riwayat Penen ------")
			subMenu()
			pilihMenu(Menu)
		case 5:
			return
		}

	}
}

func menu() {
	fmt.Println("------ Menu ------")
	fmt.Println("1. Penanaman")
	fmt.Println("2. Tanaman")
	fmt.Println("3. Hasil Panen")
	fmt.Println("4. Riwayat Panen")
	fmt.Println("5. Keluar")
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
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	userInput, _ = strconv.Atoi(input)

	switch userInput {
	case 1:
		switch n {
		case 1:
			tambah(&arrPenanaman)
		case 2:
			tambah(&arrTanaman)
		case 3:
			tambah(&arrHasilPanen)
		case 4:
			tambah(&arrHasilPanen)
		default:
			fmt.Println("Input Salah")

		}

	case 2:
		switch n {
		case 1:
			edit(&arrPenanaman)
		case 2:
			edit(&arrTanaman)
		case 3:
			edit(&arrHasilPanen)
		case 4:
			edit(&arrRiwayatPanen)
		default:
			fmt.Println("Input Salah")

		}
	case 3:
		switch n {
		case 1:
			hapus(&arrPenanaman)
		case 2:
			hapus(&arrTanaman)
		case 3:
			hapus(&arrHasilPanen)
		case 4:
			hapus(&arrRiwayatPanen)
		default:
			fmt.Println("Input Salah")

		}
	case 4:
		switch n {
		case 1:
			tampilkan(&arrPenanaman, false)
		case 2:
			tampilkan(&arrTanaman, false)
		case 3:
			tampilkan(&arrHasilPanen, false)
		case 4:
			tampilkan(&arrRiwayatPanen, false)
		default:
			fmt.Println("Input Salah")
		}
	case 5:
		switch n {
		case 1:
			search(&arrPenanaman)
		case 2:
			search(&arrTanaman)
		case 3:
			search(&arrHasilPanen)
		case 4:
			searchBinary(&arrRiwayatPanen)
		default:
			fmt.Println("Input Salah")
		}
	case 6:
		return
	default:
		fmt.Println("Input Salah")
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
	case *[]RiwayatPanen:
		var k RiwayatPanen
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
		fmt.Print("Masukkan tanggal panen tanaman (YYYY-MM-DD): ")
		k.tanggal, _ = reader.ReadString('\n')
		k.tanggal = strings.TrimSpace(k.tanggal)
	}
}

func edit(data interface{}) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama atau jenis tanaman yang ingin diedit: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch v := data.(type) {
	case *[]Penanaman:
		var tempIndices []int
		for i, item := range *v {
			if strings.EqualFold(item.namatanaman, input) || strings.EqualFold(item.jenistanaman, input) {
				tempIndices = append(tempIndices, i)
			}
		}

		if len(tempIndices) == 0 {
			fmt.Println("Data tidak ditemukan.")
			return
		}

		var editIndex int
		var temp = []Penanaman{}
		if len(tempIndices) == 1 {
			editIndex = tempIndices[0]
		} else {
			fmt.Println("Ditemukan beberapa item yang cocok:")
			for _, idx := range tempIndices {
				item := (*v)[idx]
				temp = append(temp, Penanaman{item.id, item.namatanaman, item.jenistanaman, item.lahan, item.jumlahBibit})
			}
			tampilkan(&temp, true)

			fmt.Print("Pilih nomor data yang ingin diedit: ")
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)
			if err != nil || choice < 1 || choice > len(tempIndices) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			editIndex = tempIndices[choice-1]
		}

		// Input data baru
		fmt.Print("Masukkan nama tanaman baru: ")
		namaBaru, _ := reader.ReadString('\n')
		namaBaru = strings.TrimSpace(namaBaru)

		fmt.Print("Masukkan jenis tanaman baru: ")
		jenisBaru, _ := reader.ReadString('\n')
		jenisBaru = strings.TrimSpace(jenisBaru)

		fmt.Print("Masukkan lahan baru: ")
		var lahanBaru int
		fmt.Scanln(&lahanBaru)

		fmt.Print("Masukkan jumlah bibit baru: ")
		var bibitBaru int
		fmt.Scanln(&bibitBaru)

		// Update data
		(*v)[editIndex].namatanaman = namaBaru
		(*v)[editIndex].jenistanaman = jenisBaru
		(*v)[editIndex].lahan = lahanBaru
		(*v)[editIndex].jumlahBibit = bibitBaru

		fmt.Println("Data berhasil diubah.")

	case *[]Tanaman:
		var tempIndices []int
		for i, item := range *v {
			if strings.EqualFold(item.nama, input) || strings.EqualFold(item.jenis, input) {
				tempIndices = append(tempIndices, i)
			}
		}

		if len(tempIndices) == 0 {
			fmt.Println("Data tidak ditemukan.")
			return
		}

		var editIndex int
		var temp = []Tanaman{}
		if len(tempIndices) == 1 {
			editIndex = tempIndices[0]
		} else {
			fmt.Println("Ditemukan beberapa item yang cocok:")
			for _, idx := range tempIndices {
				item := (*v)[idx]
				temp = append(temp, Tanaman{item.id, item.nama, item.jenis, item.grade})
			}
			tampilkan(&temp, true)

			fmt.Print("Pilih nomor data yang ingin diedit: ")
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)
			if err != nil || choice < 1 || choice > len(tempIndices) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			editIndex = tempIndices[choice-1]
		}

		// Input data baru
		fmt.Print("Masukkan nama tanaman baru: ")
		namaBaru, _ := reader.ReadString('\n')
		namaBaru = strings.TrimSpace(namaBaru)

		fmt.Print("Masukkan jenis tanaman baru: ")
		jenisBaru, _ := reader.ReadString('\n')
		jenisBaru = strings.TrimSpace(jenisBaru)
		
		fmt.Print("Masukkan grade tanaman baru: ")
		gradeBaru, _ := reader.ReadString('\n')
		gradeBaru = strings.TrimSpace(gradeBaru)

		// Update data
		(*v)[editIndex].nama = namaBaru
		(*v)[editIndex].jenis = jenisBaru
		(*v)[editIndex].grade = gradeBaru

		fmt.Println("Data berhasil diubah.")
	case *[]HasilPanen:
		var tempIndices []int
		for i, item := range *v {
			if strings.EqualFold(item.nama, input) || strings.EqualFold(item.jenis, input) {
				tempIndices = append(tempIndices, i)
			}
		}

		if len(tempIndices) == 0 {
			fmt.Println("Data tidak ditemukan.")
			return
		}

		var editIndex int
		var temp = []HasilPanen{}
		if len(tempIndices) == 1 {
			editIndex = tempIndices[0]
		} else {
			fmt.Println("Ditemukan beberapa item yang cocok:")
			for _, idx := range tempIndices {
				item := (*v)[idx]
				temp = append(temp, HasilPanen{item.id, item.nama, item.jenis, item.jumlahKg, item.grade})
			}
			tampilkan(&temp, true)

			fmt.Print("Pilih nomor data yang ingin diedit: ")
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)
			if err != nil || choice < 1 || choice > len(tempIndices) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			editIndex = tempIndices[choice-1]
		}

		// Input data baru
		fmt.Print("Masukkan nama tanaman baru: ")
		namaBaru, _ := reader.ReadString('\n')
		namaBaru = strings.TrimSpace(namaBaru)

		fmt.Print("Masukkan jenis tanaman baru: ")
		jenisBaru, _ := reader.ReadString('\n')
		jenisBaru = strings.TrimSpace(jenisBaru)

		fmt.Print("Masukkan jumlah (Kg) baru: ")
		var jumlahKg float64
		fmt.Scanln(&jumlahKg)

		fmt.Print("Masukkan grade baru: ")
		var grade string
		fmt.Scanln(&grade)

		// Update data
		(*v)[editIndex].nama = namaBaru
		(*v)[editIndex].jenis = jenisBaru
		(*v)[editIndex].jumlahKg = jumlahKg
		(*v)[editIndex].grade = grade

		fmt.Println("Data berhasil diubah.")
	
		case *[]RiwayatPanen:
		var tempIndices []int
		for i, item := range *v {
			if strings.EqualFold(item.nama, input) || strings.EqualFold(item.jenis, input) {
				tempIndices = append(tempIndices, i)
			}
		}

		if len(tempIndices) == 0 {
			fmt.Println("Data tidak ditemukan.")
			return
		}

		var editIndex int
		var temp = []RiwayatPanen{}
		if len(tempIndices) == 1 {
			editIndex = tempIndices[0]
		} else {
			fmt.Println("Ditemukan beberapa item yang cocok:")
			for _, idx := range tempIndices {
				item := (*v)[idx]
				temp = append(temp, RiwayatPanen{item.id, item.nama, item.jenis, item.jumlahKg, item.grade, item.tanggal})
			}
			tampilkan(&temp, true)

			fmt.Print("Pilih nomor data yang ingin diedit: ")
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)
			if err != nil || choice < 1 || choice > len(tempIndices) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			editIndex = tempIndices[choice-1]
		}

		// Input data baru
		fmt.Print("Masukkan nama tanaman baru: ")
		namaBaru, _ := reader.ReadString('\n')
		namaBaru = strings.TrimSpace(namaBaru)

		fmt.Print("Masukkan jenis tanaman baru: ")
		jenisBaru, _ := reader.ReadString('\n')
		jenisBaru = strings.TrimSpace(jenisBaru)

		fmt.Print("Masukkan jumlah (Kg) baru: ")
		var jumlahKg float64
		fmt.Scanln(&jumlahKg)
		
		fmt.Print("Masukkan grade tanaman baru: ")
		gradeBaru, _ := reader.ReadString('\n')
		gradeBaru = strings.TrimSpace(gradeBaru)

		fmt.Print("Masukkan tanggal panen baru: ")
		tanggalBaru, _ := reader.ReadString('\n')
		tanggalBaru = strings.TrimSpace(tanggalBaru)

		// Update data
		(*v)[editIndex].nama = namaBaru
		(*v)[editIndex].jenis = jenisBaru
		(*v)[editIndex].jumlahKg = jumlahKg
		(*v)[editIndex].grade = gradeBaru
		(*v)[editIndex].tanggal = tanggalBaru

		fmt.Println("Data berhasil diubah.")
	}
}

func hapus(data interface{}) {
	switch v := data.(type) {
	case *[]Penanaman:
		var sementara []Penanaman
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan ID atau Nama yang ingin dihapus: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Coba konversi ke ID
		if id, err := strconv.Atoi(input); err == nil {
			for i, item := range *v {
				if item.id == id {
					*v = append((*v)[:i], (*v)[i+1:]...)
					fmt.Println("Data dengan ID", id, "berhasil dihapus.")
					return
				}
			}
			fmt.Println("ID tidak ditemukan.")
			return
		}

		// Cari berdasarkan Nama
		var temp []int
		for i, item := range *v {
			if strings.EqualFold(item.jenistanaman, input) || strings.EqualFold(item.namatanaman, input) {
				temp = append(temp, i)
			}
		}

		if len(temp) == 0 {
			fmt.Println("Nama tidak ditemukan.")
			return
		} else if len(temp) == 1 {
			*v = append((*v)[:temp[0]], (*v)[temp[0]+1:]...)
			fmt.Println("Data dengan nama", input, "berhasil dihapus.")
			return
		} else {
			fmt.Println("Ditemukan beberapa item dengan nama tersebut:")
			for _, idx := range temp {
				item := (*v)[idx]
				sementara = append(sementara, Penanaman{item.id, item.namatanaman, item.jenistanaman, item.lahan, item.jumlahBibit})
			}
			tampilkan(&sementara, true)
			fmt.Print("Pilih nomor yang ingin dihapus: ")
			hapus, _ := reader.ReadString('\n')
			hapus = strings.TrimSpace(hapus)
			final, err := strconv.Atoi(hapus)
			if err != nil || final < 1 || final > len(temp) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			indexToDelete := temp[final-1]
			*v = append((*v)[:indexToDelete], (*v)[indexToDelete+1:]...)
			fmt.Println("Data berhasil dihapus.")
		}
	case *[]Tanaman:
		var sementara []Tanaman
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan ID atau Nama yang ingin dihapus: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Coba konversi ke ID
		if id, err := strconv.Atoi(input); err == nil {
			for i, item := range *v {
				if item.id == id {
					*v = append((*v)[:i], (*v)[i+1:]...)
					fmt.Println("Data dengan ID", id, "berhasil dihapus.")
					return
				}
			}
			fmt.Println("ID tidak ditemukan.")
			return
		}

		// Cari berdasarkan Nama
		var temp []int
		for i, item := range *v {
			if strings.EqualFold(item.jenis, input) || strings.EqualFold(item.nama, input) {
				temp = append(temp, i)
			}
		}

		if len(temp) == 0 {
			fmt.Println("Nama tidak ditemukan.")
			return
		} else if len(temp) == 1 {
			*v = append((*v)[:temp[0]], (*v)[temp[0]+1:]...)
			fmt.Println("Data dengan nama", input, "berhasil dihapus.")
			return
		} else {
			fmt.Println("Ditemukan beberapa item dengan nama tersebut:")
			for _, idx := range temp {
				item := (*v)[idx]
				sementara = append(sementara, Tanaman{item.id, item.nama, item.jenis, item.grade})
			}
			tampilkan(&sementara, true)
			fmt.Print("Pilih nomor yang ingin dihapus: ")
			hapus, _ := reader.ReadString('\n')
			hapus = strings.TrimSpace(hapus)
			final, err := strconv.Atoi(hapus)
			if err != nil || final < 1 || final > len(temp) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			indexToDelete := temp[final-1]
			*v = append((*v)[:indexToDelete], (*v)[indexToDelete+1:]...)
			fmt.Println("Data berhasil dihapus.")
		}
	case *[]HasilPanen:
		var sementara []HasilPanen
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan ID atau Nama yang ingin dihapus: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Coba konversi ke ID
		if id, err := strconv.Atoi(input); err == nil {
			for i, item := range *v {
				if item.id == id {
					*v = append((*v)[:i], (*v)[i+1:]...)
					fmt.Println("Data dengan ID", id, "berhasil dihapus.")
					return
				}
			}
			fmt.Println("ID tidak ditemukan.")
			return
		}

		// Cari berdasarkan Nama
		var temp []int
		for i, item := range *v {
			if strings.EqualFold(item.jenis, input) || strings.EqualFold(item.nama, input) {
				temp = append(temp, i)
			}
		}

		if len(temp) == 0 {
			fmt.Println("Nama tidak ditemukan.")
			return
		} else if len(temp) == 1 {
			*v = append((*v)[:temp[0]], (*v)[temp[0]+1:]...)
			fmt.Println("Data dengan nama", input, "berhasil dihapus.")
			return
		} else {
			fmt.Println("Ditemukan beberapa item dengan nama tersebut:")
			for _, idx := range temp {
				item := (*v)[idx]
				sementara = append(sementara, HasilPanen{item.id, item.nama, item.jenis, item.jumlahKg, item.grade})
			}
			tampilkan(&sementara, true)
			fmt.Print("Pilih nomor yang ingin dihapus: ")
			hapus, _ := reader.ReadString('\n')
			hapus = strings.TrimSpace(hapus)
			final, err := strconv.Atoi(hapus)
			if err != nil || final < 1 || final > len(temp) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			indexToDelete := temp[final-1]
			*v = append((*v)[:indexToDelete], (*v)[indexToDelete+1:]...)
			fmt.Println("Data berhasil dihapus.")
		}
	case *[]RiwayatPanen:
		var sementara []RiwayatPanen
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan ID atau Nama yang ingin dihapus: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Coba konversi ke ID
		if id, err := strconv.Atoi(input); err == nil {
			for i, item := range *v {
				if item.id == id {
					*v = append((*v)[:i], (*v)[i+1:]...)
					fmt.Println("Data dengan ID", id, "berhasil dihapus.")
					return
				}
			}
			fmt.Println("ID tidak ditemukan.")
			return
		}

		// Cari berdasarkan Nama
		var temp []int
		for i, item := range *v {
			if strings.EqualFold(item.jenis, input) || strings.EqualFold(item.nama, input) {
				temp = append(temp, i)
			}
		}

		if len(temp) == 0 {
			fmt.Println("Nama tidak ditemukan.")
			return
		} else if len(temp) == 1 {
			*v = append((*v)[:temp[0]], (*v)[temp[0]+1:]...)
			fmt.Println("Data dengan nama", input, "berhasil dihapus.")
			return
		} else {
			fmt.Println("Ditemukan beberapa item dengan nama tersebut:")
			for _, idx := range temp {
				item := (*v)[idx]
				sementara = append(sementara, RiwayatPanen{item.id, item.nama, item.jenis, item.jumlahKg, item.grade, item.tanggal})
			}
			tampilkan(&sementara, true)
			fmt.Print("Pilih nomor yang ingin dihapus: ")
			hapus, _ := reader.ReadString('\n')
			hapus = strings.TrimSpace(hapus)
			final, err := strconv.Atoi(hapus)
			if err != nil || final < 1 || final > len(temp) {
				fmt.Println("Pilihan tidak valid.")
				return
			}

			indexToDelete := temp[final-1]
			*v = append((*v)[:indexToDelete], (*v)[indexToDelete+1:]...)
			fmt.Println("Data berhasil dihapus.")
		}
	}
}
func search(data interface{}) {
	var nama string
	fmt.Print("Masukkan nama tanaman yang ingin dicari: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	found := false

	switch v := data.(type) {
	case *[]Penanaman:
		temp := []Penanaman{}
		for _, p := range *v {
			if strings.HasPrefix(strings.ToLower(p.jenistanaman), strings.ToLower(nama)) {
				temp = append(temp, Penanaman{p.id, p.namatanaman, p.jenistanaman, p.lahan, p.jumlahBibit})
				found = true
			}
		}
		tampilkan(&temp, false)
	case *[]Tanaman:
		temp := []Tanaman{}
		for _, p := range *v {
			if strings.HasPrefix(strings.ToLower(p.jenis), strings.ToLower(nama)) {
				temp = append(temp, Tanaman{p.id, p.nama, p.jenis, p.grade})
				found = true
			}
		}
		tampilkan(&temp, false)
	case *[]HasilPanen:
		temp := []HasilPanen{}
		for _, p := range *v {
			if strings.HasPrefix(strings.ToLower(p.jenis), strings.ToLower(nama)) {
				temp = append(temp, HasilPanen{p.id, p.nama, p.jenis, p.jumlahKg, p.grade})
				found = true
			}
		}
		tampilkan(&temp, false)
	}

	if !found {
		fmt.Println("Tanaman tidak ditemukan")
	}
	fmt.Println()
}

func searchBinary(arrRiwayatPanen *[]RiwayatPanen) {
	var nama string
	fmt.Print("Masukkan nama tanaman yang ingin dicari: ")
	fmt.Scanln(&nama)
	nama = strings.ToLower(nama)

	low, high := 0, len(*arrRiwayatPanen)-1
	found := false

	for low <= high {
		mid := (low + high) / 2
		val := strings.ToLower((*arrRiwayatPanen)[mid].jenis)

		if strings.HasPrefix(val, nama) {
			hasil := (*arrRiwayatPanen)[mid]
			fmt.Printf("id : %d, nama: %s, jenis: %s, jumlah (kg): %.2f, grade: %s, tanggal %s\n",
				hasil.id, hasil.nama, hasil.jenis, hasil.jumlahKg, hasil.grade, hasil.tanggal)
			found = true
			break
		}

		if val < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tanaman tidak ditemukan")
	}
}

func bubbleSort(arrPenanaman *[]Penanaman, arrTanaman *[]Tanaman, arrHasilPanen *[]HasilPanen) {
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

func insertionSort(arrRiwayatPanen *[]RiwayatPanen) {
	for i := 1; i < len(*arrRiwayatPanen); i++ {
		key := (*arrRiwayatPanen)[i]
		j := i - 1

		for j >= 0 && (*arrRiwayatPanen)[j].tanggal > key.tanggal {
			(*arrRiwayatPanen)[j+1] = (*arrRiwayatPanen)[j]
			j--
		}
		(*arrRiwayatPanen)[j+1] = key
	}
}

func tampilkan(data interface{}, a bool) {
	clearTerminal()
	switch arr := data.(type) {
	case *[]Penanaman:
		var label = []string{}
		if a {
			label = []string{"No", "ID", "Nama", "Jenis", "Lahan", "Jumlah"}
		} else {
			label = []string{"ID", "Nama", "Jenis", "Lahan", "Jumlah"}
		}
		// deklarasi array lebar untuk menyimpam lebar dari tabel
		lebarMaksimal := make([]int, len(label))
		// memasukkan data panjang tabel awal berdasarkan length dari label
		for i, v := range label {
			lebarMaksimal[i] = len(v)
		}
		/*
			perulangan untuk melihat apabila ada nilai yang lebih besar
			dan menukar isi dari var lebarMaksimal
		*/
		var kolom = []string{}
		for i, v := range *arr {
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.namatanaman,
					v.jenistanaman,
					strconv.Itoa(v.lahan),
					strconv.Itoa(v.jumlahBibit),
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.namatanaman,
					v.jenistanaman,
					strconv.Itoa(v.lahan),
					strconv.Itoa(v.jumlahBibit),
				}
			}

			for i, v := range kolom {
				if len(v) > lebarMaksimal[i] {
					lebarMaksimal[i] = len(v)
				}
			}
		}
		garisPembatas(lebarMaksimal)
		// spasi
		fmt.Println()
		// memberi garis di awal
		fmt.Print("|")
		// perulangan untuk memasukkan judul tabel
		for i, judul := range label {
			// formatting string
			fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
		}
		// memberikan spasi
		fmt.Println()
		garisPembatas(lebarMaksimal)
		fmt.Println()
		for i, v := range *arr {
			fmt.Print("|")
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.namatanaman,
					v.jenistanaman,
					strconv.Itoa(v.lahan),
					strconv.Itoa(v.jumlahBibit),
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.namatanaman,
					v.jenistanaman,
					strconv.Itoa(v.lahan),
					strconv.Itoa(v.jumlahBibit),
				}
			}

			for i, v := range kolom {
				fmt.Printf(" %-*s |", lebarMaksimal[i], v)
			}
			fmt.Println()
		}
		fmt.Print("+")
		for _, lebar := range lebarMaksimal {
			for i := 0; i < lebar+2; i++ {
				fmt.Print("-")
			}
			fmt.Print("+")
		}
		fmt.Println()
	case *[]Tanaman:
		var label = []string{}
		if a {
			label = []string{"No,", "ID", "Nama", "Jenis", "Grade"}
		} else {
			label = []string{"ID", "Nama", "Jenis", "Grade"}
		}
		// deklarasi array lebar untuk menyimpam lebar dari tabel
		lebarMaksimal := make([]int, len(label))
		// memasukkan data panjang tabel awal berdasarkan length dari label
		for i, v := range label {
			lebarMaksimal[i] = len(v)
		}
		/*
			perulangan untuk melihat apabila ada nilai yang lebih besar
			dan menukar isi dari var lebarMaksimal
		*/
		var kolom = []string{}
		for i, v := range *arr {
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					v.grade,
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					v.grade,
				}
			}

			for i, v := range kolom {
				if len(v) > lebarMaksimal[i] {
					lebarMaksimal[i] = len(v)
				}
			}
		}
		garisPembatas(lebarMaksimal)
		// spasi
		fmt.Println()
		// memberi garis di awal
		fmt.Print("|")
		// perulangan untuk memasukkan judul tabel
		for i, judul := range label {
			// formatting string
			fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
		}
		// memberikan spasi
		fmt.Println()
		garisPembatas(lebarMaksimal)
		fmt.Println()
		for i, v := range *arr {
			fmt.Print("|")
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					v.grade,
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					v.grade,
				}
			}

			for i, v := range kolom {
				fmt.Printf(" %-*s |", lebarMaksimal[i], v)
			}
			fmt.Println()
		}
		fmt.Print("+")
		for _, lebar := range lebarMaksimal {
			for i := 0; i < lebar+2; i++ {
				fmt.Print("-")
			}
			fmt.Print("+")
		}
		fmt.Println()
	case *[]HasilPanen:
		var label = []string{}
		if a {
			label = []string{"No", "ID", "Nama", "Jenis", "Jumlah (Kg)", "Grade"}
		} else {
			label = []string{"ID", "Nama", "Jenis", "Jumlah (Kg)", "Grade"}
		}
		// deklarasi array lebar untuk menyimpam lebar dari tabel
		lebarMaksimal := make([]int, len(label))
		// memasukkan data panjang tabel awal berdasarkan length dari label
		for i, v := range label {
			lebarMaksimal[i] = len(v)
		}
		/*
			perulangan untuk melihat apabila ada nilai yang lebih besar
			dan menukar isi dari var lebarMaksimal
		*/
		var kolom = []string{}
		for i, v := range *arr {
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
				}
			}

			for i, v := range kolom {
				if len(v) > lebarMaksimal[i] {
					lebarMaksimal[i] = len(v)
				}
			}
		}
		garisPembatas(lebarMaksimal)
		// spasi
		fmt.Println()
		// memberi garis di awal
		fmt.Print("|")
		// perulangan untuk memasukkan judul tabel
		for i, judul := range label {
			// formatting string
			fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
		}
		// memberikan spasi
		fmt.Println()
		garisPembatas(lebarMaksimal)
		fmt.Println()
		for i, v := range *arr {
			fmt.Print("|")
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
				}
			}

			for i, v := range kolom {
				fmt.Printf(" %-*s |", lebarMaksimal[i], v)
			}
			fmt.Println()
		}
		fmt.Print("+")
		for _, lebar := range lebarMaksimal {
			for i := 0; i < lebar+2; i++ {
				fmt.Print("-")
			}
			fmt.Print("+")
		}
		fmt.Println()
	case *[]RiwayatPanen:
		var label = []string{}
		if a {
			label = []string{"No", "ID", "Nama", "Jenis", "Jumlah (Kg)", "Grade", "Tanggal"}
		} else {
			label = []string{"ID", "Nama", "Jenis", "Jumlah (Kg)", "Grade", "Tanggal"}
		}
		// deklarasi array lebar untuk menyimpam lebar dari tabel
		lebarMaksimal := make([]int, len(label))
		// memasukkan data panjang tabel awal berdasarkan length dari label
		for i, v := range label {
			lebarMaksimal[i] = len(v)
		}
		/*
			perulangan untuk melihat apabila ada nilai yang lebih besar
			dan menukar isi dari var lebarMaksimal
		*/
		var kolom = []string{}
		for i, v := range *arr {
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
					v.tanggal,
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
					v.tanggal,
				}
			}
			for i, v := range kolom {
				if len(v) > lebarMaksimal[i] {
					lebarMaksimal[i] = len(v)
				}
			}
		}
		garisPembatas(lebarMaksimal)
		// spasi
		fmt.Println()
		// memberi garis di awal
		fmt.Print("|")
		// perulangan untuk memasukkan judul tabel
		for i, judul := range label {
			// formatting string
			fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
		}
		// memberikan spasi
		fmt.Println()
		garisPembatas(lebarMaksimal)
		fmt.Println()
		for i, v := range *arr {
			fmt.Print("|")
			if a {
				kolom = []string{
					strconv.Itoa(i + 1),
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
					v.tanggal,
				}
			} else {
				kolom = []string{
					strconv.Itoa(v.id),
					v.nama,
					v.jenis,
					strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
					v.grade,
					v.tanggal,
				}
			}

			for i, v := range kolom {
				fmt.Printf(" %-*s |", lebarMaksimal[i], v)
			}
			fmt.Println()
		}
		fmt.Print("+")
		for _, lebar := range lebarMaksimal {
			for i := 0; i < lebar+2; i++ {
				fmt.Print("-")
			}
			fmt.Print("+")
		}
		fmt.Println()
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
	arrRiwayatPanen = []RiwayatPanen{
		{1, "Jagung", "Jagung Manis", 320.5, "A", "2024-05-15"},
		{2, "Tomat", "Tomat Cherry", 250.0, "A", "2024-05-02"},
		{3, "Tomat", "Tomat Sayur", 230.0, "B", "2024-05-10"},
		{4, "Cabai", "Cabai Rawit", 180.0, "A", "2024-05-05"},
		{5, "Cabai", "Cabai Merah Besar", 170.0, "B", "2024-05-17"},
		{6, "Padi", "Padi IR64", 750.0, "A", "2024-05-08"},
		{7, "Padi", "Padi Inpari 32", 680.0, "B", "2024-05-01"},
		{8, "Wortel", "Wortel Lokal", 210.0, "B", "2024-05-12"},
		{9, "Kentang", "Kentang Granola", 400.0, "A", "2024-05-20"},
		{10, "Bayam", "Bayam Hijau", 160.0, "B", "2024-05-04"},
		{11, "Bayam", "Bayam Merah", 150.0, "B", "2024-05-19"},
		{12, "Terong", "Terong Ungu", 180.0, "B", "2024-05-07"},
		{13, "Timun", "Timun Suri", 200.0, "A", "2024-05-11"},
		{14, "Melon", "Melon Sky Rocket", 220.0, "A", "2024-05-14"},
		{15, "Semangka", "Semangka Inul", 240.0, "A", "2024-05-13"},
		{16, "Kacang", "Kacang Tanah Takar 1", 190.0, "B", "2024-05-06"},
		{17, "Kedelai", "Kedelai Anjasmoro", 260.0, "A", "2024-05-18"},
		{18, "Bawang", "Bawang Merah Bima", 270.0, "B", "2024-05-09"},
		{19, "Bawang", "Bawang Putih Lumbu Kuning", 280.0, "A", "2024-05-03"},
		{20, "Nanas", "Nanas Queen", 300.0, "A", "2024-05-16"},
	}
}

// lain - lain
func garisPembatas(lebarMaksimal []int) {
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan ---  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok atas kiri
		fmt.Print("+")
	}
}

func clearTerminal() {
	// hanya berkerja di windows
	cmd := exec.Command("cmd", "/c", "cls")
	// output
	cmd.Stdout = os.Stdout
	// error
	cmd.Stderr = os.Stderr
	cmd.Run()
}
