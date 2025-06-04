package main

import (
	"bufio"
	"os/exec"
	"runtime"

	// "encoding/csv"
	"fmt"
	"os"
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

/*
===============================
|| Menu Handler (Pakai Switch) ||
===============================
*/
func menu() bool {
	var counterPenanaman = 20
	var counterTanaman = 20
	var counterHasilPanen = 20
	var counterRiwayatPanen = 20

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
		menuPenanaman(1, &counterPenanaman)

	case 2:
		fmt.Println("=====  (Tanaman)  =====")
		menuTanaman(2, &counterTanaman)

	case 3:
		fmt.Println("===== (Hasil Panen) =====")
		menuHasilPanen(3, &counterHasilPanen)

	case 4:
		fmt.Println("===== (Riwayat Panen) =====")
		menuRiwayatPanen(4, &counterRiwayatPanen)

	case 5:
		return false
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
	return true
}

/* ===============================
|| CSV HANDLER (import Encoding)||
============================== */

/* ===============================
|| Sort (Selection & Insertion) ||
=============================== */

/*
	===============================

|| Search (Sequential & Binary) ||
===============================
*/
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

/*
===============================
|| SubMenu                   ||
===============================
*/
func menuPenanaman(arrTujuan int, x *int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
	fmt.Println("4. Tampilkan")
	fmt.Println("5. Pencarian(Search)")
	fmt.Println("6. kembali")
	fmt.Print("Pilih Menu: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	input, _ := strconv.Atoi(userInput)
	switch input {
	case 1:
		clearTerminal()
		TambahData(1, x)
	case 2:
		clearTerminal()

	case 3:
		clearTerminal()

	case 4:
		clearTerminal()
		tampilkanPenanaman(arrPenanaman)
	case 5:
		clearTerminal()
		SequentialSearch(arrTujuan)
	case 6:
		clearTerminal()
		return
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
}
func menuTanaman(arrTujuan int, x *int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
	fmt.Println("4. Tampilkan")
	fmt.Println("5. Pencarian(Search)")
	fmt.Println("6. kembali")
	fmt.Print("Pilih Menu: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	input, _ := strconv.Atoi(userInput)
	switch input {
	case 1:
		clearTerminal()
		TambahData(arrTujuan, x)
	case 2:
		clearTerminal()

	case 3:
		clearTerminal()

	case 4:
		clearTerminal()
		tampilkanTanaman(arrTanaman)
	case 5:
		clearTerminal()
		SequentialSearch(arrTujuan)
	case 6:
		clearTerminal()
		return
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
}
func menuHasilPanen(arrTujuan int, x *int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
	fmt.Println("4. Tampilkan")
	fmt.Println("5. Pencarian(Search)")
	fmt.Println("6. kembali")
	fmt.Print("Pilih Menu: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	input, _ := strconv.Atoi(userInput)
	switch input {
	case 1:
		clearTerminal()
		TambahData(arrTujuan, x)
	case 2:
		clearTerminal()

	case 3:
		clearTerminal()

	case 4:
		clearTerminal()
		tampilkanHasilPanen(arrHasilPanen)
	case 5:
		clearTerminal()
		SequentialSearch(arrTujuan)
	case 6:
		clearTerminal()
		return
	default:
		clearTerminal()
		fmt.Println("===== Error =====")
		fmt.Printf("Tidak Valid: %d\n", input)
	}
}

func menuRiwayatPanen(arrTujuan int, x *int) {
	clearTerminal()
	tampilkanRiwayatPanen(arrRiwayatPanen)
}

/*
===============================
|| Tabel (Tampilan)          ||
===============================
*/
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

/*
===============================
|| Tambah Data                 ||
===============================
*/
func TambahData(arrTujuan int, x *int) {
	var reader = bufio.NewReader(os.Stdin)
	switch arrTujuan {
	case 1:
		var lahan, jumlahBibit int
		arrPenanaman[*x].id = *x + 1
		fmt.Print("Masukkan nama tanaman: ")
		namatanaman, _ := reader.ReadString('\n')
		namatanaman = strings.TrimSpace(namatanaman)
		arrPenanaman[*x].namatanaman = namatanaman
		fmt.Print("Masukkan jenis tanaman : ")
		jenistanaman, _ := reader.ReadString('\n')
		jenistanaman = strings.TrimSpace(jenistanaman)
		arrPenanaman[*x].jenistanaman = jenistanaman
		fmt.Print("Masukkan lahan: ")
		fmt.Scanln(&lahan)
		arrPenanaman[*x].lahan = lahan
		fmt.Print("Masukkan jumlah bibit: ")
		fmt.Scanln(&jumlahBibit)
		arrPenanaman[*x].jumlahBibit = jumlahBibit
		*x++
	case 2:
		arrTanaman[*x].id = *x + 1
		fmt.Print("Masukkan nama tanaman: ")
		namatanaman, _ := reader.ReadString('\n')
		namatanaman = strings.TrimSpace(namatanaman)
		arrTanaman[*x].nama = namatanaman
		fmt.Print("Masukkan jenis tanaman : ")
		jenistanaman, _ := reader.ReadString('\n')
		jenistanaman = strings.TrimSpace(jenistanaman)
		arrTanaman[*x].jenis = jenistanaman
		fmt.Print("Masukkan grade: ")
		grade, _ := reader.ReadString('\n')
		grade = strings.TrimSpace(grade)
		arrTanaman[*x].grade = grade
		*x++
	case 3:
		var jumlahKg float64
		arrHasilPanen[*x].id = *x + 1
		fmt.Print("Masukkan nama tanaman: ")
		namatanaman, _ := reader.ReadString('\n')
		namatanaman = strings.TrimSpace(namatanaman)
		arrHasilPanen[*x].nama = namatanaman
		fmt.Print("Masukkan jenis tanaman : ")
		jenistanaman, _ := reader.ReadString('\n')
		jenistanaman = strings.TrimSpace(jenistanaman)
		arrHasilPanen[*x].jenis = jenistanaman
		fmt.Print("Masukkan grade: ")
		grade, _ := reader.ReadString('\n')
		grade = strings.TrimSpace(grade)
		arrTanaman[*x].grade = grade
		fmt.Print("Masukkan jumlah bibit: ")
		fmt.Scanln(&jumlahKg)
		arrHasilPanen[*x].jumlahKg = jumlahKg

		*x++
	case 4:

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

/*
===============================
|| Inisiasi Data Array        ||
===============================
*/
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

/*
Program harus dibuat secara modular dengan menggunakan subprogram. [x]
Program memiliki tema. Tema Bebas (contoh : peternakan, pertanian, pendidikan, kesehatan dll) [X]
Struktur subprogram ditentukan oleh masing-masing kelompok. Setiap subprogram harus dilengkapi dengan parameter dan spesifikasinya. []
Program harus mengimplementasikan Array(min 3) dan Tipe Bentukan(min 3). Array statis harus digunakan, bukan array dinamis atau slice. []
Program harus mengimplementasikan algoritma Sequential dan Binary Search untuk pencarian data, pengubahan (edit), atau penghapusan data tertentu. []
Program harus mengimplementasikan algoritma Selection Sort dan Insertion Sort untuk mengurutkan data dalam kategori yang berbeda saat menampilkan hasil. Setiap kategori harus bisa diurutkan dalam urutan naik (ascending) maupun turun (descending). []
Penggunaan statement break (kecuali dalam repeat-until) dan continue tidak diperbolehkan. []
Variabel global hanya diperbolehkan untuk menyimpan array utama yang akan diproses. []
*/
