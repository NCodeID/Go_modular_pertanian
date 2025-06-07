package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// deklarasi struct
type Penanaman struct {
	// {Struct untuk melihat lokasi lahan dan jumlah bibit yang tertanam}
	id                        int
	namatanaman, jenistanaman string
	lahan, jumlahBibit        int
}

type Tanaman struct {
	// {Struct untuk melihat nama sayuran / tanaman yang terdiri dari Nama, Jenis, Grade (kualitas)}
	id                 int
	nama, jenis, grade string
}

type HasilPanen struct {
	// {Struct untuk melihat stock dalam jumlahKg dan grade yang tersedia}
	id       int
	nama     string
	jenis    string
	jumlahKg float64
	grade    string
}

type RiwayatPanen struct {
	// {Menampilkan riwayatPanen}
	id             int
	nama, jenis    string
	jumlahKg       float64
	grade, tanggal string
}

// Membuka list data Tanaman

// deklarasi 3 Slice
var arrTanaman = [100]Tanaman{}
var arrPenanaman = [100]Penanaman{}
var arrHasilPanen = [100]HasilPanen{}
var arrRiwayatPanen = [100]RiwayatPanen{}

func main() {
	inisiasiDataAwal()

	for {
		kondisi := menu()
		if !kondisi {
			return
		} else {
			continue
		}
	}
}

func idMaxPenanaman(arrPenanaman [100]Penanaman) int {
	penanamanTertinggi := arrPenanaman[0]

	for i := 1; i < len(arrPenanaman); i++ {
		if arrPenanaman[i].id > penanamanTertinggi.id {
			penanamanTertinggi = arrPenanaman[i]
		}
	}
	return penanamanTertinggi.id
}

func idMaxTanaman(arrTanaman [100]Tanaman) int {
	tanamanTertinggi := arrTanaman[0]

	for i := 1; i < len(arrTanaman); i++ {
		if arrTanaman[i].id > tanamanTertinggi.id {
			tanamanTertinggi = arrTanaman[i]
		}
	}
	return tanamanTertinggi.id
}

func idMaxHasilpanen(arrHasilPanen [100]HasilPanen) int {
	HasilPanenTertinggi := arrHasilPanen[0]

	for i := 1; i < len(arrHasilPanen); i++ {
		if arrHasilPanen[i].id > HasilPanenTertinggi.id {
			HasilPanenTertinggi = arrHasilPanen[i]
		}
	}
	return HasilPanenTertinggi.id
}

func idMaxRiwayatpanen(arrRiwayatPanen [100]RiwayatPanen) int {
	HasilRiwayatpanen := arrRiwayatPanen[0]

	for i := 1; i < len(arrRiwayatPanen); i++ {
		if arrRiwayatPanen[i].id > HasilRiwayatpanen.id {
			HasilRiwayatpanen = arrRiwayatPanen[i]
		}
	}
	return HasilRiwayatpanen.id
}

/*===============================
|| Menu Handler (Pakai Switch) ||
===============================*/

func menu() bool {

	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("------ Menu ------")
	fmt.Println("1. Penanaman")
	fmt.Println("2. Tanaman")
	fmt.Println("3. Hasil Panen")
	fmt.Println("4. Riwayat Panen")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih Menu : ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	input, _ := strconv.Atoi(userInput)
	switch input {
	case 1:
		clearTerminal()
		fmt.Println("===== (Penanaman) =====")
		menuPenanaman(1)

	case 2:
		fmt.Println("=====  (Tanaman)  =====")
		menuTanaman(2)

	case 3:
		fmt.Println("===== (Hasil Panen) =====")
		menuHasilPanen(3)

	case 4:
		fmt.Println("===== (Riwayat Panen) =====")
		menuRiwayatPanen(4)

	case 5:
		return false
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
	return true
}

/* =====================
|| Insertion Sort     ||
======================*/

func insertionSorTanaman(arrTanaman *[100]Tanaman) {
	n := len(*arrTanaman)
	for i := 1; i < n; i++ {
		key := (*arrTanaman)[i]
		j := i - 1

		for j >= 0 && (*arrTanaman)[j].id < key.id {
			(*arrTanaman)[j+1] = (*arrTanaman)[j]
			j--
		}
		(*arrTanaman)[j+1] = key
	}
}

func insertionSortPenanaman(arrPenanaman *[100]Penanaman) {
	n := len(*arrPenanaman)
	for i := 1; i < n; i++ {
		key := (*arrPenanaman)[i]
		j := i - 1

		for j >= 0 && (*arrPenanaman)[j].id < key.id {
			(*arrPenanaman)[j+1] = (*arrPenanaman)[j]
			j--
		}
		(*arrPenanaman)[j+1] = key
	}
}

func insertionSortHasilPanen(arrHasilPanen *[100]HasilPanen) {
	n := len(*arrHasilPanen)
	for i := 1; i < n; i++ {
		key := (*arrHasilPanen)[i]
		j := i - 1

		for j >= 0 && (*arrHasilPanen)[j].id < key.id {
			(*arrHasilPanen)[j+1] = (*arrHasilPanen)[j]
			j--
		}
		(*arrHasilPanen)[j+1] = key
	}
}

/*================================
|| Sort (Selection & Insertion) ||
================================*/

func selectionSortPenanaman(arrPenanaman *[100]Penanaman, jumlahData int) {
	for i := 0; i < jumlahData-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if (*arrPenanaman)[j].id < (*arrPenanaman)[minIdx].id {
				minIdx = j
			}
		}
		if minIdx != i {
			(*arrPenanaman)[i], (*arrPenanaman)[minIdx] = (*arrPenanaman)[minIdx], (*arrPenanaman)[i]
		}
	}
}

func selectionSortTanaman(arrTanaman *[100]Tanaman, jumlahData int) {
	for i := 0; i < jumlahData-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if (*arrTanaman)[j].id < (*arrTanaman)[minIdx].id {
				minIdx = j
			}
		}
		if minIdx != i {
			(*arrTanaman)[i], (*arrTanaman)[minIdx] = (*arrTanaman)[minIdx], (*arrTanaman)[i]
		}
	}
}

func selectionSortHasilpanen(arrHasilPanen *[100]HasilPanen, jumlahData int) {
	for i := 0; i < jumlahData-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if (*arrHasilPanen)[j].id < (*arrHasilPanen)[minIdx].id {
				minIdx = j
			}
		}
		if minIdx != i {
			(*arrHasilPanen)[i], (*arrHasilPanen)[minIdx] = (*arrHasilPanen)[minIdx], (*arrHasilPanen)[i]
		}
	}
}

/*================================
|| Search (Sequential & Binary) ||
================================*/

func SequentialSearch(arrTujuan int) {
	var input string
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama tanaman yang ingin dicari: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	switch arrTujuan {
	case 1:
		var temp = [100]Penanaman{}
		count := 0
		for i, v := range arrPenanaman {
			if strings.Contains(strings.ToLower(input), strings.ToLower(v.namatanaman)) || strings.HasPrefix(strings.ToLower(v.namatanaman), strings.ToLower(input)) {
				temp[count] = arrPenanaman[i]
				count++
			}
		}
		tampilkanPenanaman(temp)
	case 2:
		var temp = [100]Tanaman{}
		count := 0
		for i, v := range arrTanaman {
			if strings.Contains(strings.ToLower(input), strings.ToLower(v.nama)) || strings.HasPrefix(strings.ToLower(v.nama), strings.ToLower(input)) {
				temp[count] = arrTanaman[i]
				count++
			}
		}
		// tampilkanPenanaman(temp)
	case 3:
		var temp = [100]HasilPanen{}
		count := 0
		for i, v := range arrHasilPanen {
			if strings.Contains(strings.ToLower(input), strings.ToLower(v.nama)) || strings.HasPrefix(strings.ToLower(v.nama), strings.ToLower(input)) {
				temp[count] = arrHasilPanen[i]
				count++
			}
		}
		// tampilkanPenanaman(temp)
	case 4:
		var temp = [100]RiwayatPanen{}
		count := 0
		for i, v := range arrRiwayatPanen {
			if strings.Contains(strings.ToLower(input), strings.ToLower(v.nama)) || strings.HasPrefix(strings.ToLower(v.nama), strings.ToLower(input)) {
				temp[count] = arrRiwayatPanen[i]
				count++
			}
		}
		// tampilkanPenanaman(temp)
	}
}
func binarySearchPenanaman(arr [100]Penanaman) {
	// Urutkan terlebih dahulu
	selectionSortPenanaman(&arr, len(arr))

	var targetID int
	fmt.Print("Masukkan ID yang ingin dicari: ")
	fmt.Scanln(&targetID)

	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid].id == targetID {
			var temp = [100]Penanaman{}
			temp[0] = arr[mid]
			tampilkanPenanaman(temp)
			return
		}
		if arr[mid].id < targetID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Data dengan ID tersebut tidak ditemukan.")
}
func binarySearchTanaman(arr [100]Tanaman) {
	// Urutkan terlebih dahulu
	selectionSortTanaman(&arr, len(arr))

	var targetID int
	fmt.Print("Masukkan ID yang ingin dicari: ")
	fmt.Scanln(&targetID)

	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid].id == targetID {
			var temp = [100]Tanaman{}
			temp[0] = arr[mid]
			tampilkanTanaman(temp)
			return
		}
		if arr[mid].id < targetID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Data dengan ID tersebut tidak ditemukan.")
}
func binarySearchHasilPanen(arr [100]HasilPanen) {
	selectionSortHasilpanen(&arr, len(arr))

	var targetID int
	fmt.Print("Masukkan ID yang ingin dicari: ")
	fmt.Scanln(&targetID)

	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid].id == targetID {
			var temp = [100]HasilPanen{}
			temp[0] = arr[mid]
			tampilkanHasilPanen(temp)
			return
		}
		if arr[mid].id < targetID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Data dengan ID tersebut tidak ditemukan.")
}

/*=============================
|| SubMenu                   ||
=============================*/

func menuPenanaman(arrTujuan int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
	fmt.Println("4. Tampilkan")
	fmt.Println("5. Pencarian(Search by ID)")
	fmt.Println("6. Pencarian(Search by Nama)")
	fmt.Println("7. Selection Sort(Ascending)")
	fmt.Println("8. Insertion Sort(Descending)")
	fmt.Println("9. kembali")
	fmt.Print("Pilih Menu: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	input, _ := strconv.Atoi(userInput)
	switch input {
	case 1:
		clearTerminal()
		TambahData(arrTujuan)
	case 2:
		clearTerminal()
		editDataPenanaman(&arrPenanaman)
	case 3:
		clearTerminal()
		n := idMaxPenanaman(arrPenanaman)
		hapusPenanaman(&arrPenanaman, &arrRiwayatPanen, &n)
		tampilkanPenanaman(arrPenanaman)
	case 4:
		clearTerminal()
		tampilkanPenanaman(arrPenanaman)
	case 5:
		clearTerminal()
		binarySearchPenanaman(arrPenanaman)
	case 6:
		clearTerminal()
		SequentialSearch(arrTujuan)
	case 7:
		clearTerminal()
		selectionSortPenanaman(&arrPenanaman, idMaxPenanaman(arrPenanaman))
		tampilkanPenanaman(arrPenanaman)
	case 8:
		clearTerminal()
		insertionSortPenanaman(&arrPenanaman)
		tampilkanPenanaman(arrPenanaman)
	case 9:
		clearTerminal()
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
}
func menuTanaman(arrTujuan int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
	fmt.Println("4. Tampilkan")
	fmt.Println("5. Pencarian(Search by ID)")
	fmt.Println("6. Pencarian(Search by Nama)")
	fmt.Println("7. Selection Sort(Ascending)")
	fmt.Println("8. Insertion Sort(Descending)")
	fmt.Println("9. kembali")
	fmt.Print("Pilih Menu: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	input, _ := strconv.Atoi(userInput)
	switch input {
	case 1:
		clearTerminal()
		TambahData(arrTujuan)
	case 2:
		clearTerminal()
		editDataTanaman(&arrTanaman)
	case 3:
		clearTerminal()
		hapusTanaman(&arrTanaman)
		tampilkanTanaman(arrTanaman)
	case 4:
		clearTerminal()
		tampilkanTanaman(arrTanaman)
	case 5:
		clearTerminal()
		binarySearchTanaman(arrTanaman)
	case 6:
		clearTerminal()
		SequentialSearch(arrTujuan)
	case 7:
		clearTerminal()
		selectionSortTanaman(&arrTanaman, idMaxTanaman(arrTanaman))
		tampilkanTanaman(arrTanaman)
	case 8:
		clearTerminal()
		insertionSorTanaman(&arrTanaman)
		tampilkanTanaman(arrTanaman)
	case 9:
		clearTerminal()
		return
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
}

func menuHasilPanen(arrTujuan int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
	fmt.Println("5. Pencarian(Search by ID)")
	fmt.Println("6. Pencarian(Search by Nama)")
	fmt.Println("7. Selection Sort(Ascending)")
	fmt.Println("8. Insertion Sort(Descending)")
	fmt.Println("9. kembali")
	fmt.Print("Pilih Menu: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	input, _ := strconv.Atoi(userInput)
	switch input {
	case 1:
		clearTerminal()
		TambahData(arrTujuan)
	case 2:
		clearTerminal()
		editDataHasilPanen(&arrHasilPanen)

	case 3:
		clearTerminal()
		hapusHasilPanen(&arrHasilPanen)
		tampilkanHasilPanen(arrHasilPanen)

	case 4:
		clearTerminal()
		tampilkanHasilPanen(arrHasilPanen)
	case 5:
		clearTerminal()
		binarySearchHasilPanen(arrHasilPanen)
	case 6:
		clearTerminal()
		SequentialSearch(arrTujuan)
	case 7:
		clearTerminal()
		selectionSortHasilpanen(&arrHasilPanen, idMaxHasilpanen(arrHasilPanen))
		tampilkanHasilPanen(arrHasilPanen)
	case 8:
		clearTerminal()
		insertionSortHasilPanen(&arrHasilPanen)
		tampilkanHasilPanen(arrHasilPanen)
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
}

func menuRiwayatPanen(arrTujuan int) {
	clearTerminal()
	tampilkanRiwayatPanen(arrRiwayatPanen)
}

/*===============================
|| Tambah Data                 ||
===============================*/

func TambahData(arrTujuan int) {
	var reader = bufio.NewReader(os.Stdin)
	switch arrTujuan {
	case 1:
		var lahan, jumlahBibit int
		var last = idMaxPenanaman(arrPenanaman)
		arrPenanaman[last].id += last + 1
		fmt.Print("Masukkan nama tanaman: ")
		namatanaman, _ := reader.ReadString('\n')
		namatanaman = strings.TrimSpace(namatanaman)
		arrPenanaman[last].namatanaman = namatanaman
		fmt.Print("Masukkan jenis tanaman : ")
		jenistanaman, _ := reader.ReadString('\n')
		jenistanaman = strings.TrimSpace(jenistanaman)
		arrPenanaman[last].jenistanaman = jenistanaman
		fmt.Print("Masukkan lahan: ")
		fmt.Scanln(&lahan)
		arrPenanaman[last].lahan = lahan
		fmt.Print("Masukkan jumlah bibit: ")
		fmt.Scanln(&jumlahBibit)
		arrPenanaman[last].jumlahBibit = jumlahBibit
	case 2:
		var last = idMaxTanaman(arrTanaman)
		arrTanaman[last].id = last + 1
		fmt.Print("Masukkan nama tanaman: ")
		namatanaman, _ := reader.ReadString('\n')
		namatanaman = strings.TrimSpace(namatanaman)
		arrTanaman[last].nama = namatanaman
		fmt.Print("Masukkan jenis tanaman : ")
		jenistanaman, _ := reader.ReadString('\n')
		jenistanaman = strings.TrimSpace(jenistanaman)
		arrTanaman[last].jenis = jenistanaman
		fmt.Print("Masukkan grade: ")
		grade, _ := reader.ReadString('\n')
		grade = strings.TrimSpace(grade)
		arrTanaman[last].grade = grade
	case 3:
		var jumlahKg float64
		var last = idMaxHasilpanen(arrHasilPanen)
		arrHasilPanen[last].id = last + 1
		fmt.Print("Masukkan nama tanaman: ")
		namatanaman, _ := reader.ReadString('\n')
		namatanaman = strings.TrimSpace(namatanaman)
		arrHasilPanen[last].nama = namatanaman
		fmt.Print("Masukkan jenis tanaman : ")
		jenistanaman, _ := reader.ReadString('\n')
		jenistanaman = strings.TrimSpace(jenistanaman)
		arrHasilPanen[last].jenis = jenistanaman
		fmt.Print("Masukkan grade: ")
		grade, _ := reader.ReadString('\n')
		grade = strings.TrimSpace(grade)
		arrTanaman[last].grade = grade
		fmt.Print("Masukkan jumlah bibit: ")
		fmt.Scanln(&jumlahKg)
		arrHasilPanen[last].jumlahKg = jumlahKg
	}
}

/*===============================
|| EditData (cls)              ||
===============================*/

func editDataPenanaman(arr *[100]Penanaman) {
	var lahan, jumlahBibit int
	var userInput int
	reader := bufio.NewReader(os.Stdin)

	// Menampilkan data yang bisa diedit
	tampilkanPenanaman(*arr)

	fmt.Print("Masukkan nomor data yang ingin diedit: ")
	fmt.Scanln(&userInput)

	// Flush newline sebelum membaca string
	reader.ReadString('\n') // Untuk membersihkan newline sisa dari Scanln

	fmt.Print("Nama Tanaman: ")
	namatanaman, _ := reader.ReadString('\n')
	namatanaman = strings.TrimSpace(namatanaman)

	fmt.Print("Jenis Tanaman: ")
	jenisTanaman, _ := reader.ReadString('\n')
	jenisTanaman = strings.TrimSpace(jenisTanaman)

	fmt.Print("Lahan: ")
	fmt.Scanln(&lahan)

	fmt.Print("Jumlah Bibit: ")
	fmt.Scanln(&jumlahBibit)

	// Menyimpan perubahan
	(*arr)[userInput].namatanaman = namatanaman
	(*arr)[userInput].jenistanaman = jenisTanaman
	(*arr)[userInput].lahan = lahan
	(*arr)[userInput].jumlahBibit = jumlahBibit

	fmt.Println("Data berhasil diperbarui.")
}

func editDataTanaman(arr *[100]Tanaman) {
	var grade string
	var userInput int
	reader := bufio.NewReader(os.Stdin)

	// Menampilkan data yang bisa diedit
	tampilkanTanaman(*arr)

	fmt.Print("Masukkan nomor data yang ingin diedit: ")
	fmt.Scanln(&userInput)

	// Flush newline sebelum membaca string
	reader.ReadString('\n') // Untuk membersihkan newline sisa dari Scanln

	fmt.Print("Nama Tanaman: ")
	namatanaman, _ := reader.ReadString('\n')
	namatanaman = strings.TrimSpace(namatanaman)

	fmt.Print("Jenis Tanaman: ")
	jenisTanaman, _ := reader.ReadString('\n')
	jenisTanaman = strings.TrimSpace(jenisTanaman)

	fmt.Print("Grade: ")
	fmt.Scanln(&grade)

	// Menyimpan perubahan
	(*arr)[userInput].nama = namatanaman
	(*arr)[userInput].jenis = jenisTanaman
	(*arr)[userInput].grade = grade

	fmt.Println("Data berhasil diperbarui.")
}
func editDataHasilPanen(arr *[100]HasilPanen) {
	var jumlahKg float64
	var grade string
	var userInput int
	reader := bufio.NewReader(os.Stdin)

	// Menampilkan data yang bisa diedit
	tampilkanHasilPanen(*arr)

	fmt.Print("Masukkan nomor data yang ingin diedit: ")
	fmt.Scanln(&userInput)

	// Flush newline sebelum membaca string
	reader.ReadString('\n') // Untuk membersihkan newline sisa dari Scanln

	fmt.Print("Nama Tanaman: ")
	namatanaman, _ := reader.ReadString('\n')
	namatanaman = strings.TrimSpace(namatanaman)

	fmt.Print("Jenis Tanaman: ")
	jenisTanaman, _ := reader.ReadString('\n')
	jenisTanaman = strings.TrimSpace(jenisTanaman)

	fmt.Print("Jumlah (Kg): ")
	fmt.Scanln(&jumlahKg)

	fmt.Print("grade: ")
	fmt.Scanln(&grade)

	// Menyimpan perubahan
	(*arr)[userInput].nama = namatanaman
	(*arr)[userInput].jenis = jenisTanaman
	(*arr)[userInput].jumlahKg = jumlahKg
	(*arr)[userInput].grade = grade

	fmt.Println("Data berhasil diperbarui.")
}

/*
===============================
|| Hapus                       ||
===============================
*/
func hapusPenanaman(arrPenanaman *[100]Penanaman, arrRiwayatPanen *[100]RiwayatPanen, n *int) {
	var reader = bufio.NewReader(os.Stdin)
	nRiwayat := idMaxRiwayatpanen(*arrRiwayatPanen)

	var target int
	var jumlahKg float64

	fmt.Print("Masukkan ID tanaman yang ingin dihapus: ")
	fmt.Scanln(&target)

	(*arrRiwayatPanen)[*&nRiwayat].id = nRiwayat + 1
	(*arrRiwayatPanen)[*&nRiwayat].nama = arrPenanaman[target-1].namatanaman
	(*arrRiwayatPanen)[*&nRiwayat].jenis = arrPenanaman[target-1].jenistanaman

	fmt.Print("Masukkan jumlah (kg) setelah dipanen: ")
	fmt.Scanln(&jumlahKg)
	(*arrRiwayatPanen)[*&nRiwayat].jumlahKg = jumlahKg
	fmt.Print("Masukkan grade setelah dipanen: ")
	grade, _ := reader.ReadString('\n')
	grade = strings.TrimSpace(grade)
	(*arrRiwayatPanen)[*&nRiwayat].grade = grade
	fmt.Print("Masukkan tanggal : ")
	tanggal, _ := reader.ReadString('\n')
	tanggal = strings.TrimSpace(tanggal)
	(*arrRiwayatPanen)[*&nRiwayat].tanggal = tanggal

	fmt.Println()
	for i := 0; i < len(*arrPenanaman); i++ {
		if (*arrPenanaman)[i].id == target {
			for j := i; j < len(*arrPenanaman)-1; j++ {
				(*arrPenanaman)[j] = (*arrPenanaman)[j+1]
			}
		}
	}
}

func hapusTanaman(arrTanaman *[100]Tanaman) {
	var target int

	fmt.Print("Masukkan ID tanaman yang ingin dihapus: ")
	fmt.Scanln(&target)

	for i := 0; i < len(*arrTanaman); i++ {
		if (*arrTanaman)[i].id == target {
			for j := i; j < len(*arrTanaman)-1; j++ {
				(*arrTanaman)[j] = (*arrTanaman)[j+1]
			}
		}
	}
}

func hapusHasilPanen(arrHasilPanen *[100]HasilPanen) {
	var target int

	fmt.Print("Masukkan ID tanaman yang ingin dihapus: ")
	fmt.Scanln(&target)

	for i := 0; i < len(*&arrHasilPanen); i++ {
		if (*&arrHasilPanen)[i].id == target {
			for j := i; j < len(*&arrHasilPanen)-1; j++ {
				(*&arrHasilPanen)[j] = (*&arrHasilPanen)[j+1]
			}
		}
	}
}

/*===============================
|| ClearTerminal (cls)         ||
===============================*/

func clearTerminal() {
	var cmd *exec.Cmd
	sistemOperasi := runtime.GOOS
	if sistemOperasi == "windows" {
		// hanya berkerja di windows
		cmd = exec.Command("cmd", "/c", "cls")
		// output
	} else if sistemOperasi == "linux" {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	// error
	cmd.Stderr = os.Stderr
	cmd.Run()

}

/*===============================
|| Inisiasi Data Array         ||
===============================*/

func inisiasiDataAwal() {
	arrTanaman = [100]Tanaman{
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
	arrPenanaman = [100]Penanaman{
		{20, "Jagung", "Jagung Manis", 1, 150},
		{19, "Tomat", "Tomat Cherry", 1, 120},
		{18, "Tomat", "Tomat Sayur", 2, 100},
		{17, "Cabai", "Cabai Rawit", 2, 200},
		{16, "Cabai", "Cabai Merah Besar", 3, 180},
		{15, "Padi", "Padi IR64", 3, 300},
		{14, "Padi", "Padi Inpari 32", 4, 250},
		{13, "Wortel", "Wortel Lokal", 4, 170},
		{12, "Kentang", "Kentang Granola", 5, 190},
		{11, "Bayam", "Bayam Hijau", 5, 130},
		{10, "Bayam", "Bayam Merah", 6, 100},
		{9, "Terong", "Terong Ungu", 6, 110},
		{8, "Timun", "Timun Suri", 7, 140},
		{7, "Melon", "Melon Sky Rocket", 7, 160},
		{6, "Semangka", "Semangka Inul", 8, 170},
		{5, "Kacang", "Kacang Tanah Takar 1", 8, 180},
		{4, "Kedelai", "Kedelai Anjasmoro", 9, 150},
		{3, "Bawang", "Bawang Merah Bima", 9, 200},
		{2, "Bawang", "Bawang Putih Lumbu Kuning", 10, 90},
		{1, "Nanas", "Nanas Queen", 10, 100},
	}
	arrHasilPanen = [100]HasilPanen{
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
	arrRiwayatPanen = [100]RiwayatPanen{
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

/*=============================
|| Tabel (Tampilan)          ||
=============================*/

func tampilkanPenanaman(x [100]Penanaman) {
	label := []string{"ID", "Nama", "Jenis", "Lahan", "Jumlah"}
	lebarMaksimal := [5]int{}
	for i, v := range label {
		lebarMaksimal[i] = len(v)
	}

	for _, v := range x {
		kolom := [5]string{
			strconv.Itoa(v.id),
			v.namatanaman,
			v.jenistanaman,
			strconv.Itoa(v.lahan),
			strconv.Itoa(v.jumlahBibit),
		}

		for i, v := range kolom {
			if len(v) > lebarMaksimal[i] {
				lebarMaksimal[i] = len(v)
			}
		}
	}
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// memberi garis di awal
	fmt.Print("|")
	for i, judul := range label {
		// formatting string
		fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
	}
	// newline
	fmt.Println()
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// menampilkan isi array
	for _, v := range x {
		if len(v.namatanaman) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		} else {
			fmt.Print("|")
		}
		kolom := [5]string{
			strconv.Itoa(v.id),
			v.namatanaman,
			v.jenistanaman,
			strconv.Itoa(v.lahan),
			strconv.Itoa(v.jumlahBibit),
		}
		if len(v.namatanaman) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		}
		for i, v := range kolom {

			fmt.Printf(" %-*s |", lebarMaksimal[i], v)
		}
		fmt.Println()
	}
	fmt.Println()

	// akhir dari tabel (penutup)
	// pojok kanan

}
func tampilkanRiwayatPanen(x [100]RiwayatPanen) {
	label := []string{"ID", "Nama", "Jenis", "Grade", "Jumlah (Kg)", "Tanggal"}
	lebarMaksimal := [6]int{}
	for i, v := range label {
		lebarMaksimal[i] = len(v)
	}

	for _, v := range x {
		kolom := [6]string{
			strconv.Itoa(v.id),
			v.nama,
			v.jenis,
			v.grade,
			strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
			v.tanggal,
		}

		for i, v := range kolom {
			if len(v) > lebarMaksimal[i] {
				lebarMaksimal[i] = len(v)
			}
		}
	}
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// memberi garis di awal
	fmt.Print("|")
	for i, judul := range label {
		// formatting string
		fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
	}
	// newline
	fmt.Println()
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// menampilkan isi array
	for _, v := range x {
		if len(v.nama) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		} else {
			fmt.Print("|")
		}
		kolom := [6]string{
			strconv.Itoa(v.id),
			v.nama,
			v.jenis,
			v.grade,
			strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
			v.tanggal,
		}
		if len(v.nama) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		}
		for i, v := range kolom {

			fmt.Printf(" %-*s |", lebarMaksimal[i], v)
		}
		fmt.Println()
	}
	fmt.Println()

	// akhir dari tabel (penutup)
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}

}
func tampilkanTanaman(x [100]Tanaman) {
	label := []string{"ID", "Nama", "Jenis", "Grade"}
	lebarMaksimal := [4]int{}
	for i, v := range label {
		lebarMaksimal[i] = len(v)
	}

	for _, v := range x {
		kolom := [4]string{
			strconv.Itoa(v.id),
			v.nama,
			v.jenis,
			v.grade,
		}

		for i, v := range kolom {
			if len(v) > lebarMaksimal[i] {
				lebarMaksimal[i] = len(v)
			}
		}
	}
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// memberi garis di awal
	fmt.Print("|")
	for i, judul := range label {
		// formatting string
		fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
	}
	// newline
	fmt.Println()
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// menampilkan isi array
	for _, v := range x {
		if len(v.nama) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		} else {
			fmt.Print("|")
		}
		kolom := [4]string{
			strconv.Itoa(v.id),
			v.nama,
			v.jenis,
			v.grade,
		}
		if len(v.nama) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		}
		for i, v := range kolom {

			fmt.Printf(" %-*s |", lebarMaksimal[i], v)
		}
		fmt.Println()
	}
	fmt.Println()

	// akhir dari tabel (penutup)
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}

}
func tampilkanHasilPanen(x [100]HasilPanen) {
	label := []string{"ID", "Nama", "Jenis", "Jumlah (Kg)", "Grade"}
	lebarMaksimal := [5]int{}
	for i, v := range label {
		lebarMaksimal[i] = len(v)
	}

	for _, v := range x {
		kolom := [5]string{
			strconv.Itoa(v.id),
			v.nama,
			v.jenis,
			strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
			v.grade,
		}

		for i, v := range kolom {
			if len(v) > lebarMaksimal[i] {
				lebarMaksimal[i] = len(v)
			}
		}
	}
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// memberi garis di awal
	fmt.Print("|")
	for i, judul := range label {
		// formatting string
		fmt.Printf(" %-*s |", lebarMaksimal[i], judul)
	}
	// newline
	fmt.Println()
	// pojok kanan
	fmt.Print("+")
	for _, lebar := range lebarMaksimal {
		// perulangan untuk menambahkan `---`  sesuai dengan
		for i := 0; i < lebar+2; i++ { // +2 spasi untuk readability
			fmt.Print("-")
		}
		// pojok kiri
		fmt.Print("+")
	}
	// newline
	fmt.Println()
	// menampilkan isi array
	for _, v := range x {
		if len(v.nama) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		} else {
			fmt.Print("|")
		}
		kolom := [5]string{
			strconv.Itoa(v.id),
			v.nama,
			v.jenis,
			strconv.FormatFloat(v.jumlahKg, 'f', 1, 64),
			v.grade,
		}
		if len(v.nama) == 0 {
			fmt.Print("+")
			for _, lebar := range lebarMaksimal {
				for i := 0; i < lebar+2; i++ {
					fmt.Print("-")
				}
				// pojok kiri
				fmt.Print("+")
			}
			fmt.Println()
			return
		}
		for i, v := range kolom {

			fmt.Printf(" %-*s |", lebarMaksimal[i], v)
		}
		fmt.Println()
	}
	fmt.Println()

	// akhir dari tabel (penutup)
	// pojok kanan

}
