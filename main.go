package main

import (
	"fmt"
	"time"
)

type Langganan struct {
	Nama    string
	Biaya   int
	Metode  string
	Tenggat int
}

type Transaksi struct {
	Nama    string
	Biaya   int
	Metode  string
	Tanggal time.Time
}

const LMAX int = 10
const TMAX int = 100

type TabLangganan = [LMAX]Langganan
type TabTransaksi = [TMAX]Transaksi

func main() {
	var input int

	var lanjut bool
	lanjut = true

	var daftarLangganan TabLangganan
	var n int
	daftarLangganan[0] = Langganan{
		Nama:    "Youtube Music",
		Biaya:   60000,
		Metode:  "Gopay",
		Tenggat: 2,
	}
	daftarLangganan[2] = Langganan{
		Nama:    "Netflix",
		Biaya:   90000,
		Metode:  "Mandiri",
		Tenggat: 23,
	}
	daftarLangganan[1] = Langganan{
		Nama:    "Spotify",
		Biaya:   62000,
		Metode:  "Gopay",
		Tenggat: 1,
	}
	daftarLangganan[3] = Langganan{
		Nama:    "Adobe",
		Biaya:   380000,
		Metode:  "Mandiri",
		Tenggat: 12,
	}
	daftarLangganan[4] = Langganan{
		Nama:    "The New York Times",
		Biaya:   80000,
		Metode:  "Paypal",
		Tenggat: 4,
	}
	n = 5

	for lanjut {
		cariTenggatTerdekat(daftarLangganan, n)
		menu()
		fmt.Print(">> ")
		fmt.Scan(&input)

		switch input {
		case 1:
			cetakLangganan(daftarLangganan, n)
		case 2:
			tambahLangganan(&daftarLangganan, &n)
		case 3:
			ubahLangganan(&daftarLangganan, n)
		case 4:
			hapusLangganan(&daftarLangganan, &n)
		case 5:
			cetakSatuLangganan(&daftarLangganan, n)
		case 6:
			urutLangganan(&daftarLangganan, n)
		case 7:
			cetakTotalBiayaLangganan(daftarLangganan, n)
		case 8:
			cetakPengeluaranTerbesar(daftarLangganan, n)
		case 0:
			lanjut = false
		}
		fmt.Println("---")
	}
}

func menu() {
	/*
		Initial state:
		Final state: Mengeluarkan barisan string berisikan nama apalikasi dan menu - menu
		yang dapat dipilih pengguna.
	*/
	fmt.Println("  ___ ___ __  __ ___ _    ___   ___ _   _ ___ ___  ___ ___ ___ ___ _____ ___ ___  _  _   __  __   _   _  _   _   ___ ___ ___ ")
	fmt.Println(" / __|_ _|  \\/  | _ \\ |  | __| / __| | | | _ ) __|/ __| _ \\_ _| _ \\_   _|_ _/ _ \\| \\| | |  \\/  | /_\\ | \\| | /_\\ / __| __| _ \\")
	fmt.Println(" \\__ \\| || |\\/| |  _/ |__| _|  \\__ \\ |_| | _ \\__ \\ (__|   /| ||  _/ | |  | | (_) | .` | | |\\/| |/ _ \\| .` |/ _ \\ (_ | _||   /")
	fmt.Println(" |___/___|_|  |_|_| |____|___| |___/\\___/|___/___/\\___|_|_\\___|_|   |_| |___\\___/|_|\\_| |_|  |_/_/ \\_\\_|\\_/_/ \\_\\___|___|_|_\\")
	fmt.Println("                                                                                                                             ")

	fmt.Println("[1] Cetak layanan berlangganan")
	fmt.Println("[2] Tambah layanan berlangganan")
	fmt.Println("[3] Ubah layanan berlangganan")
	fmt.Println("[4] Hapus layanan berlangganan")
	fmt.Println("[5] Cari layanan berlangganan")
	fmt.Println("[6] Urutkan layanan berlangganan")
	fmt.Println("[7] Total pengeluaran layanan berlangganan")
	fmt.Println("[8] Rekomendasi")
	fmt.Println("[0] Keluar")
}

func cetakLangganan(d TabLangganan, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		Final state:
		i: nilai diperbarui untuk perulangan.
		Mengeluarkan output berupa isi dari array TabLangganan d.
	*/

	fmt.Println("Daftar layanan berlangganan:")

	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s %d %s %d\n", i+1, d[i].Nama, d[i].Biaya, d[i].Metode, d[i].Tenggat)
	}
}

func cetakSatuLangganan(d *TabLangganan, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		Final state:
		i: berisikan indeks dari layanan yang dicari berdasarkan namanya.
		Mengeluarkan barisan string yang berisikan nilai dari layanan yang dipilih.
	*/

	var i int

	urutLanggananNama(d, n)

	i = cariLangganan(*d, n)

	fmt.Println("Nama layanan:", d[i].Nama)
	fmt.Println("Biaya layanan:", d[i].Biaya)
	fmt.Println("Metode Pembayaran layanan:", d[i].Metode)
	fmt.Println("Tenggat layanan tiap bulan:", d[i].Tenggat)
}

func tambahLangganan(d *TabLangganan, n *int) {
	/*
		Initial state:
		nm, m: tidak terdefinisi.
		b, t: tidak terdefinisi.
		Final state:
		nm: berisikan nama layanan baru.
		m: berisikan metode pembayaran layanan baru.
		b: berisikan biaya layanan baru.
		t: berisikan tanggal tenggat layanan baru.
		Akan merubah isi dari array TabLangganan d dengna menambahkan item baru, lalu
		menambahkan nilai n sebanyak 1.
	*/

	fmt.Println("Masukkan data layanan berlangganan")

	var nm, m string
	var b, t int

	fmt.Print("Nama >> ")
	fmt.Scan(&nm)
	fmt.Print("Biaya >> ")
	fmt.Scan(&b)
	fmt.Print("Metode >> ")
	fmt.Scan(&m)
	fmt.Print("Tenggat >> ")
	fmt.Scan(&t)

	d[*n].Nama = nm
	d[*n].Biaya = b
	d[*n].Metode = m
	d[*n].Tenggat = t
	*n++

	fmt.Println("Data layanan berlangganan telah disimpan")
}

func ubahLangganan(d *TabLangganan, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		FInal state:
		Mencetak daftar layanan berlangganan yang sdah disimpan.
		i: bernilai index dari layanan yang ingin dirubah.
		Akan memanggil prosedur tambahLangganan sebagai prosedur untuk merubah nilai dari
		layanan yang telah dipilih.
	*/

	var i int

	cetakLangganan(*d, n)

	fmt.Print("Masukkan index layanan >> ")
	fmt.Scan(&i)
	i--

	if i < 0 || i > n {
		fmt.Printf("Index harus antara 1 - %d\n", n)
	}

	tambahLangganan(d, &i)
}

func hapusLangganan(d *TabLangganan, n *int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		Final state:
		Mencetak daftar layanan berlangganan yang sudah terdaftar.
		Mengurtkan daftar layanan berdasarkan nama secara menaik (ascending).
		i: berisikan index dari layanan yang akan dihapus. Dicari berdasarkan nama dari layanan tersebut.
		Akan menggeser posisi dari nilai - nilai yang terdapat dalam array dimulai dari indeks layanan yang
		dihaps.
	*/

	var i int

	cetakLangganan(*d, *n)

	urutLanggananNama(d, *n)

	i = cariLangganan(*d, *n)

	if i < 0 || i > *n {
		fmt.Println("Data tidak dapat ditemukan")
		return
	}

	for i < *n-1 {
		d[i] = d[i+1]
		i++
	}
	*n--

	fmt.Println("Data telah dihapus")
}

func cariLangganan(d TabLangganan, n int) int {
	/*
		Initial state:
		x: tidak terdefinisi.
		bawah: berniilai 0.
		atas: bernilai n - 1.
		tengah: tidak terdefinisi
		i: bernilai -1.
		Final state:
		x: berisikan nama dari layanan yang dicari.
		tengah: berisikan hasil dari (bawah + atas) / 2 yang akan berubah - ubah setiap iterasi.
		Akan mengembalikan indeks dari layanan yang dicari apabila ada dalam array. Jika tidak
		ada, maka fungsi akan mengembalikan nilai -1 sebagai penanda bahwa layanan
		tidak dapat ditemukan.
	*/

	var x string
	var bawah, tengah, atas, i int

	fmt.Print("Masukkan nama layanan >> ")
	fmt.Scan(&x)

	i = -1
	bawah = 0
	atas = n - 1
	for bawah <= atas && i == -1 {
		tengah = (bawah + atas) / 2

		if x == d[tengah].Nama {
			i = tengah
		} else if x > d[tengah].Nama {
			bawah = tengah + 1
		} else if x < d[tengah].Nama {
			atas = tengah - 1
		}
	}

	return i
}

func urutLangganan(d *TabLangganan, n int) {
	/*
		Initial state:
		x, y: tidak terdefinisi.
		Final state:
		x: berisikan atribut yang akan digunakan untuk mengurutkan array.
		y: berisikan metode pengurutan, menaik (ascending) atau menurun (descending).
		Akan merubah urutan array berdasarkan nilai x dan y dengan memanggil prosedur yang
		sesuai dengan parameter yang diberikan.
	*/

	var x, y string

	fmt.Print("Urut berdasarkan [nama/harga/tenggat] >> ")
	fmt.Scan(&x)
	fmt.Print("Pilih jenis urutan [asc/dsc] >> ")
	fmt.Scan(&y)

	switch x {
	case "nama":
		urutLanggananNama(d, n)
	case "harga":
		urutLanggananHarga(d, n)
	case "tenggat":
		urutLanggananTenggat(d, n)
	default:
		fmt.Println("Pilih opsi yang tersedia")
		return
	}

	switch y {
	case "asc":
		cetakLangganan(*d, n)
	case "dsc":
		urutLanggananBalik(d, n)
		cetakLangganan(*d, n)
	default:
		fmt.Println("Pilih opsi yang tersedia")
		return
	}
}

func urutLanggananNama(d *TabLangganan, n int) {
	/*
		Initial state:
		i, j, min: tidak terdefinisi.
		temp: tidak terdefinisi.
		Final state:
		i, j: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari layanan dengan nama terkecil.
		temp: menyimpan nilai dari array ke min yang akan ditukar nilainya
		dengan array ke i.
	*/

	var i, j, min int
	var temp Langganan

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if d[j].Nama < d[min].Nama {
				min = j
			}
		}

		temp = d[min]
		d[min] = d[i]
		d[i] = temp
	}
}

func urutLanggananHarga(d *TabLangganan, n int) {
	/*
		Initial state:
		i, j, min: tidak terdefinisi.
		temp: tidak terdefinisi.
		Final state:
		i, j: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari layanan dengan biaya terkecil.
		temp: menyimpan nilai dari array ke min yang akan ditukar nilainya
		dengan array ke i.
	*/

	var i, j, min int
	var temp Langganan

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if d[j].Biaya < d[min].Biaya {
				min = j
			}
		}

		temp = d[min]
		d[min] = d[i]
		d[i] = temp
	}
}

func urutLanggananTenggat(d *TabLangganan, n int) {
	/*
		Initial state:
		i, j, min: tidak terdefinisi.
		temp: tidak terdefinisi.
		Final state:
		i, j: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari layanan dengan tenggat terdekat.
		temp: menyimpan nilai dari array ke min yang akan ditukar nilainya
		dengan array ke i.
	*/

	var i, j, min int
	var temp Langganan

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if d[j].Tenggat < d[min].Tenggat {
				min = j
			}
		}

		temp = d[min]
		d[min] = d[i]
		d[i] = temp
	}
}

func urutLanggananBalik(d *TabLangganan, n int) {
	/*
		Initial state:
		i: bernilai 0.
		temp: tidak terdefinisi.
		Final state:
		i: digunakan sebagai counter perulangan.
		temp: digunakan untuk menyimpan nilai sementara dari array ke i yang akan ditukarkan nilainya.
		Akan melakukan perulangan dari 0 hingga n/2 untuk membalikkan urutan array.
	*/

	var i int
	var temp Langganan

	i = 0
	for i < n/2 && d[i] != d[n-1-i] {
		temp = d[i]
		d[i] = d[n-1-i]
		d[n-1-i] = temp

		i++
	}
}

func cariTenggatTerdekat(d TabLangganan, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		min: bernilai 0.
		hariIni: bernilai tanggal prosedur ini dijalankan.
		Final state:
		i: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari layanan dengan tenggat terdekat.
		Akan mengeluarkan layanan dengan tenggat terdekat dari sekarang.
	*/

	var i, min, hariIni int

	hariIni = time.Now().Day()

	min = 0
	for i = 1; i < n; i++ {
		// Logika sekali pakai, jangan diganti 🙏😭
		if d[min].Tenggat <= hariIni && d[i].Tenggat <= hariIni {
			if d[i].Tenggat < d[min].Tenggat {
				min = i
			}
		} else if d[min].Tenggat < hariIni && d[i].Tenggat >= hariIni {
			min = i
		} else if d[min].Tenggat >= hariIni && d[i].Tenggat >= hariIni {
			if d[i].Tenggat < d[min].Tenggat {
				min = i
			}
		}
	}

	fmt.Println("Layanan", d[min].Nama, "akan diperbarui pada tanggal", d[min].Tenggat)
}

func cetakTotalBiayaLangganan(d TabLangganan, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		total: bernilai 0.
		Final state:
		i: digunakan sebagai counter perulangan.
		total: menyimpan akumulasi biaya layanan yang dimiliki.
		Akan mengeluarkan nilai total sebagai total pengeluaran layanan berlangganan bulanan.
	*/

	var i, total int

	total = 0

	for i = 0; i < n; i++ {
		total = d[i].Biaya
	}

	fmt.Println("Total pengeluaran (langganan):", total)
}

func cetakPengeluaranTerbesar(d TabLangganan, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		max: bernilai 0.
		Final state:
		i: digunakan sebagai counter perulangan.
		max: berisikan indeks dari layanan yang memiliki biaya terbesar.
		Akan menampilkan nama dan biaya dari layanan dengan biaya terbesar sebagai rekomendasi
		pembatalan langganan untuk memotong pengeluaran pengguna.
	*/

	var i, max int

	max = 0
	for i = 0; i < n; i++ {
		if d[i].Biaya > d[max].Biaya {
			max = i
		}
	}

	fmt.Println("Pengeluaran terbesar jatuh kepada", d[max].Nama, "dengan biaya", d[max].Biaya)
}
