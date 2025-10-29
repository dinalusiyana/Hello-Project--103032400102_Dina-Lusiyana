//ini update terbaru untuk latihan Git

package main

import (
	"fmt"
)

// tipe alias
type Tanggal string
type Durasi int
type Kalori int

// Tipe Bentukan struct
type Workout struct {
	ID      int
	Tanggal Tanggal
	Jenis   string
	Durasi  Durasi
	Kalori  Kalori
}

// array olahraga
const NMAX int = 1000

var olahraga [NMAX]Workout

func main() {
	var pilihan int
	var nWorkout int = 0

	for pilihan != 6 {
		menu()
		fmt.Scan(&pilihan)
		//pemanggilan func
		if pilihan == 1 {
			tambahWorkout(&nWorkout)
		} else if pilihan == 2 {
			tampilkanData(nWorkout)
		} else if pilihan == 3 {
			cariWorkout(nWorkout)
		} else if pilihan == 4 {
			rekapData(nWorkout)
		} else if pilihan == 5 {
			olahragaFav(nWorkout)
		} else if pilihan == 6 {
			fmt.Println("Terima kasih sudah menggunakan aplikasi")
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menu() {
	//tampilan menu
	fmt.Println("+-------------------------------------+")
	fmt.Println("|     Aplikasi Manajemen Workout      |")
	fmt.Println("+-------------------------------------+")
	fmt.Println("| 1. Tambah Workout                   |")
	fmt.Println("| 2. Tampilkan Semua Data             |")
	fmt.Println("| 3. Cari Workout Berdasarkan Tanggal |")
	fmt.Println("| 4. Rekap Latihan Mingguan           |")
	fmt.Println("| 5. Olahraga Terfavorit              |")
	fmt.Println("| 6. Simpan & Keluar                  |")
	fmt.Println("+-------------------------------------+")
	fmt.Println(" Pilih menu:                         ")

}

func tambahWorkout(nWorkout *int) {
	//penambahan data workout ke arr olahraga dan otomatis memperbarui jumlah data(nWorkout) yang tersimpan
	var tgl, jenis string
	var durasi, kalori int

	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&tgl)
	fmt.Print("Jenis Olahraga: ")
	fmt.Scan(&jenis)
	fmt.Print("Durasi (menit): ")
	fmt.Scan(&durasi)
	fmt.Print("Kalori terbakar: ")
	fmt.Scan(&kalori)

	//Proses penambahan data baru ke arr olhrg pada index nWorkout dengan elemen ID, Tanggal,  Jenis, Durasi, Kalori.
	//Misalnya: jika *nWorkout = 3, maka data baru dimasukkan ke olahraga[3].
	olahraga[*nWorkout] = Workout{
		ID:      *nWorkout + 1, //pemberian identitas utk mempermudah pencarian data
		Tanggal: Tanggal(tgl),
		Jenis:   jenis,
		Durasi:  Durasi(durasi),
		Kalori:  Kalori(kalori),
	}
	*nWorkout++ //data ditambahkan

	fmt.Println("Data berhasil ditambahkan!")

}

// insertion sort
func tampilkanData(nWorkout int) {
	var i, j int
	var temp Workout

	if nWorkout == 0 {
		fmt.Println("Belum ada data workout yang tersimpan")
		return
	}

	// Insertion Sort (Ascending) kalori: menampilkan output berdasarkan kalori yg terbakar
	for i = 1; i < nWorkout; i++ {
		temp = olahraga[i]
		j = i - 1
		for j >= 0 && olahraga[j].Kalori > temp.Kalori {
			olahraga[j+1] = olahraga[j]
			j--
		}
		olahraga[j+1] = temp
	}

	//menampilkan data yang sudah disorting berdasarkan kalori(ascending)
	fmt.Println("\n=== Semua Data Workout ===")
	for i = 0; i < nWorkout; i++ {
		fmt.Printf("ID: %d | Tanggal: %s | Jenis: %s | Durasi: %d menit | Kalori: %d\n",
			olahraga[i].ID, olahraga[i].Tanggal, olahraga[i].Jenis, olahraga[i].Durasi, olahraga[i].Kalori)
	}
}

func cariWorkout(nWorkout int) {
	var i int
	var ditemukan int
	var cariTanggal string

	if nWorkout == 0 {
		fmt.Println("Belum ada data yang tersimpan")
		return
	}

	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scan(&cariTanggal)

	for i = 0; i < nWorkout; i++ {
		//Sequential search
		if olahraga[i].Tanggal == Tanggal(cariTanggal) {
			fmt.Printf("ID: %d | Tanggal: %s | Jenis: %s | Durasi: %d menit | Kalori: %d\n",
				olahraga[i].ID, olahraga[i].Tanggal, olahraga[i].Jenis, olahraga[i].Durasi, olahraga[i].Kalori)
			ditemukan = 1
		}
	}

	if ditemukan == 0 {
		fmt.Println("Tidak terdapat data workout pada tanggal tersebut")
	}
}

func rekapData(nWorkout int) {
	var totalDurasi Durasi
	var totalKalori Kalori
	var i int

	if nWorkout == 0 {
		fmt.Println("Data tidak ditemukan :(")
		return
	}

	// looping untuk hitung total durasi dan kalori
	for i = 0; i < nWorkout; i++ {
		totalDurasi += olahraga[i].Durasi
		totalKalori += olahraga[i].Kalori
	}

	fmt.Println("=== Rekap Workout Mingguan ===")
	fmt.Println("Total Durasi Workout:", totalDurasi, "menit")
	fmt.Println("Total Kalori Terbakar:", totalKalori, "kalori")
}

func olahragaFav(nWorkout int) {
	var jenis string
	var jenisList [NMAX]string
	var jumlahList [NMAX]int
	var nJenis int
	var i, j, idxMax, jumlah, maxJumlah, idx int
	var kiri, kanan, tengah int

	if nWorkout == 0 {
		fmt.Println("Belum ada data workout yang tersimpan.")
		return
	}

	// Selection sort digunakan untuk mengurutkan data
	//berdasarkan jenis olahraga agar bisa dihitung frekuensinya secara berurutan.
	for i = 0; i < nWorkout-1; i++ {
		idx = i
		for j = i + 1; j < nWorkout; j++ {
			if olahraga[j].Jenis < olahraga[idx].Jenis {
				idx = j
			}
		}
		temp := olahraga[i]
		olahraga[i] = olahraga[idx]
		olahraga[idx] = temp
	}
	//loop untuk menghitung frek
	i = 0
	for i < nWorkout {
		jenis = olahraga[i].Jenis
		jumlah = 1
		for i+1 < nWorkout && olahraga[i+1].Jenis == jenis {
			jumlah++
			i++
		}
		jenisList[nJenis] = jenis
		jumlahList[nJenis] = jumlah
		nJenis++
		i++
	}

	// binary search digunakan untuk mencari frekuensi terbesar
	//pada array jenisList karena data tersebut sudah terurut.
	idxMax = -1
	maxJumlah = 0
	for i = 0; i < nJenis; i++ {
		kiri = 0
		kanan = nJenis - 1
		idx = -1
		ditemukan := false

		for kiri <= kanan && !ditemukan {
			tengah = (kiri + kanan) / 2
			if jenisList[tengah] == jenisList[i] {
				idx = tengah
				ditemukan = true
			} else if jenisList[tengah] < jenisList[i] {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}

		if ditemukan && jumlahList[idx] > maxJumlah {
			maxJumlah = jumlahList[idx]
			idxMax = idx
		}
	}

	if idxMax != -1 {
		fmt.Println("\n=== Olahraga Terfavorit ===")
		fmt.Printf("Jenis olahraga: %s (sebanyak %d kali)\n", jenisList[idxMax], maxJumlah)
	} else {
		fmt.Println("Tidak ada olahraga yang ditemukan.")
	}
}
