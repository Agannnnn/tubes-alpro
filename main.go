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
	Nama     string
	Biaya    int
	Metode   string
	Tanggal  time.Time
	Kategori string // Pemasukan, Langganan, dll. (Sesuai user)
}

const LMAX int = 10
const TMAX int = 100

type TabLangganan = [LMAX]Langganan
type TabTransaksi = [TMAX]Transaksi

func main() {
	var input int
	var lanjut bool
	var daftarLangganan TabLangganan
	var daftarTransaksi TabTransaksi
	var nDaftarLangganan, nDaftarTransaksi int

	nDaftarLangganan = 0
	nDaftarTransaksi = 0

	lanjut = true
	for lanjut {
		tenggatTerdekatLangganan(daftarLangganan, nDaftarLangganan)

		menu()

		fmt.Print(">> ")
		fmt.Scan(&input)

		switch input {
		case 1:
			menuLangganan(&daftarLangganan, &nDaftarLangganan)
		case 2:
			menuTransaksi(&daftarTransaksi, &nDaftarTransaksi)
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

	fmt.Println("[1] Menu langganan")
	fmt.Println("[2] Menu transaksi")
	fmt.Println("[0] Keluar")
}

func menuLangganan(d *TabLangganan, n *int) {
	/*
		Initial state:
		input: tidak terdefinisi.
		lanjut: bernilai true.
		Final state:
		lanjut: digunakan sebagai flag agar perulangan menu terus berjalan.
		input: berisi menu yang dipilih pengguna.
		Akan menampilkan daftar menu dan memanggil fungsi - fungsi dan prosedur sesuai menu yang dipilih.
	*/

	var input int
	var lanjut bool
	lanjut = true

	for lanjut {
		fmt.Println("  _                                 _             _                                         ")
		fmt.Println(" | |__ _ _  _ __ _ _ _  __ _ _ _   | |__  ___ _ _| |__ _ _ _  __ _ __ _ __ _ _ _  __ _ _ _  ")
		fmt.Println(" | / _` | || / _` | ' \\/ _` | ' \\  | '_ \\/ -_) '_| / _` | ' \\/ _` / _` / _` | ' \\/ _` | ' \\ ")
		fmt.Println(" |_\\__,_|\\_, \\__,_|_||_\\__,_|_||_| |_.__/\\___|_| |_\\__,_|_||_\\__, \\__, \\__,_|_||_\\__,_|_||_|")
		fmt.Println("         |__/                                                |___/|___/                     ")

		fmt.Println("[1] Cetak layanan berlangganan")
		fmt.Println("[2] Tambah layanan berlangganan")
		fmt.Println("[3] Ubah layanan berlangganan")
		fmt.Println("[4] Hapus layanan berlangganan")
		fmt.Println("[5] Cari layanan berlangganan")
		fmt.Println("[6] Urutkan layanan berlangganan")
		fmt.Println("[7] Total pengeluaran layanan berlangganan")
		fmt.Println("[8] Rekomendasi")
		fmt.Println("[0] Keluar")

		fmt.Print(">> ")
		fmt.Scan(&input)

		switch input {
		case 1:
			listLangganan(*d, *n)
		case 2:
			simpanLangganan(d, n)
		case 3:
			ubahLangganan(d, *n)
		case 4:
			hapusLangganan(d, n)
		case 5:
			cetakLangganan(d, *n)
		case 6:
			urutLangganan(d, *n)
		case 7:
			cetakTotalBiayaLangganan(*d, *n)
		case 8:
			cetakPengeluaranTerbesar(*d, *n)
		case 0:
			lanjut = false
		}

		fmt.Println("---")
	}
}

func listLangganan(d TabLangganan, n int) {
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

func cetakLangganan(d *TabLangganan, n int) {
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
	fmt.Println("Metode pembayaran layanan:", d[i].Metode)
	fmt.Println("Tenggat layanan tiap bulan:", d[i].Tenggat)
}

func simpanLangganan(d *TabLangganan, n *int) {
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
		Akan memanggil prosedur simpanLangganan sebagai prosedur untuk merubah nilai dari
		layanan yang telah dipilih.
	*/

	var i int

	listLangganan(*d, n)

	fmt.Print("Masukkan index layanan >> ")
	fmt.Scan(&i)
	i--

	if i < 0 || i > n {
		fmt.Printf("Index harus antara 1 - %d\n", n)
	}

	simpanLangganan(d, &i)
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
		dihapus.
	*/

	var i int

	listLangganan(*d, *n)

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
		listLangganan(*d, n)
	case "dsc":
		urutLanggananDesc(d, n)
		listLangganan(*d, n)
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

func urutLanggananDesc(d *TabLangganan, n int) {
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

func tenggatTerdekatLangganan(d TabLangganan, n int) {
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

func menuTransaksi(d *TabTransaksi, n *int) {
	/*
		Initial state:
		input: tidak terdefinisi.
		lanjut: bernilai true.
		Final state:
		lanjut: digunakan sebagai flag agar perulangan menu terus berjalan.
		input: berisi menu yang dipilih pengguna.
		Akan menampilkan daftar menu dan memanggil fungsi - fungsi dan prosedur sesuai menu yang dipilih.
	*/

	var input, total, batas int
	var lanjut bool

	lanjut = true

	for lanjut {
		total = totalTransaksi(d, *n)
		fmt.Println("Nominal hingga batas transaksi:", batas-total)

		fmt.Println("  _____                       _       _ ")
		fmt.Println(" |_   _| _ __ _ _ _  ___ __ _| |__ __(_)")
		fmt.Println("   | || '_/ _` | ' \\(_-</ _` | / /(_-< |")
		fmt.Println("   |_||_| \\__,_|_||_/__/\\__,_|_\\_\\/__/_|")
		fmt.Println("                                        ")

		fmt.Println("[1] Cetak transaksi")
		fmt.Println("[2] Tambah transaksi")
		fmt.Println("[3] Ubah transaksi")
		fmt.Println("[4] Hapus transaksi")
		fmt.Println("[5] Cari transaksi")
		fmt.Println("[6] Urutkan transaksi")
		fmt.Println("[7] Simpan batas transaksi")
		fmt.Println("[0] Keluar")

		fmt.Print(">> ")
		fmt.Scan(&input)

		switch input {
		case 1:
			listTransaksi(*d, *n)
		case 2:
			simpanTransaksi(d, n)
		case 3:
			ubahTransaksi(d, *n)
		case 4:
			hapusTransaksi(d, n)
		case 5:
			cetakTransaksi(d, *n)
		case 6:
			urutTransaksi(d, *n)
		case 7:
			simpanBatasTransaksi(&batas)
		case 0:
			lanjut = false
		}

		fmt.Println("---")
	}
}

func listTransaksi(d TabTransaksi, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		Final state:
		i: nilai diperbarui untuk perulangan.
		Mengeluarkan output berupa isi dari array TabTransaksi d.
	*/

	fmt.Println("Daftar transaksi:")

	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s %s %d %s %s\n", i+1, d[i].Tanggal.Format(time.DateOnly), d[i].Nama, d[i].Biaya, d[i].Metode, d[i].Kategori)
	}
}

func cetakTransaksi(d *TabTransaksi, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		Final state:
		i: berisikan indeks dari layanan yang dicari berdasarkan namanya.
		Mengeluarkan barisan string yang berisikan nilai dari layanan yang dipilih.
	*/

	var i int

	urutTransaksiNama(d, n)

	i = cariTransaksi(*d, n)

	fmt.Println("Nama transaksi:", d[i].Nama)
	fmt.Println("Biaya transaksi:", d[i].Biaya)
	fmt.Println("Metode pembayaran transaksi:", d[i].Metode)
	fmt.Println("Waktu transaksi:", d[i].Tanggal)
	fmt.Println("Kategori transaksi:", d[i].Kategori)
}

func simpanTransaksi(d *TabTransaksi, n *int) {
	/*
		Initial state:
		nm, m: tidak terdefinisi.
		b, t: tidak terdefinisi.
		Final state:
		nm: berisikan nama transaksi baru.
		m: berisikan metode pembayaran transaksi baru.
		b: berisikan biaya transaksi baru.
		t: berisikan tanggal tenggat transaksi baru.
		Akan merubah isi dari array TabTransaksi d dengna menambahkan item baru, lalu
		menambahkan nilai n sebanyak 1.
	*/

	fmt.Println("Masukkan data transaksi")

	var nm, m, k string
	var b int

	fmt.Print("Nama >> ")
	fmt.Scan(&nm)
	fmt.Print("Biaya >> ")
	fmt.Scan(&b)
	fmt.Print("Metode >> ")
	fmt.Scan(&m)
	fmt.Print("Kategori >> ")
	fmt.Scan(&k)

	d[*n].Nama = nm
	d[*n].Biaya = b
	d[*n].Metode = m
	d[*n].Kategori = k
	d[*n].Tanggal = time.Now()
	*n++

	fmt.Println("Data transaksi telah disimpan")
}

func ubahTransaksi(d *TabTransaksi, n int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		FInal state:
		Mencetak daftar transaksi yang sdah disimpan.
		i: bernilai index dari transaksi yang ingin dirubah.
		Akan memanggil prosedur simpanTransaksi sebagai prosedur untuk merubah nilai dari
		transaksi yang telah dipilih.
	*/

	var i int

	listTransaksi(*d, n)

	fmt.Print("Masukkan index transaksi >> ")
	fmt.Scan(&i)
	i--

	if i < 0 || i > n {
		fmt.Printf("Index harus antara 1 - %d\n", n)
	}

	simpanTransaksi(d, &i)
}

func hapusTransaksi(d *TabTransaksi, n *int) {
	/*
		Initial state:
		i: tidak terdefinisi.
		Final state:
		Mencetak daftar transaksi yang sudah terdaftar.
		Mengurtkan daftar transaksi berdasarkan nama secara menaik (ascending).
		i: berisikan index dari transaksi yang akan dihapus. Dicari berdasarkan nama dari transaksi tersebut.
		Akan menggeser posisi dari nilai - nilai yang terdapat dalam array dimulai dari indeks transaksi yang
		dihapus.
	*/

	var i int

	listTransaksi(*d, *n)

	urutTransaksiNama(d, *n)

	i = cariTransaksi(*d, *n)

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

func cariTransaksi(d TabTransaksi, n int) int {
	/*
		Initial state:
		x, i: tidak terdefinisi.
		idx: bernilai -1.
		Final state:
		x: berisikan nama dari transaksi yang dicari.
		Akan mengembalikan indeks dari transaksi yang dicari apabila ada dalam array. Jika tidak
		ada, maka fungsi akan mengembalikan nilai -1 sebagai penanda bahwa transaksi
		tidak dapat ditemukan.
	*/

	var x string
	var i, idx int

	fmt.Print("Masukkan nama transaksi >> ")
	fmt.Scan(&x)

	idx = -1

	if x > d[n-1].Nama {
		return idx
	}

	i = 0
	for idx == -1 && idx < n {
		if d[i].Nama == x {
			idx = i
		}
		i++
	}

	return idx
}

func urutTransaksi(d *TabTransaksi, n int) {
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

	fmt.Print("Urut berdasarkan [nama/harga/kategori/tanggal] >> ")
	fmt.Scan(&x)
	fmt.Print("Pilih jenis urutan [asc/dsc] >> ")
	fmt.Scan(&y)

	switch x {
	case "nama":
		urutTransaksiNama(d, n)
	case "harga":
		urutTransaksiHarga(d, n)
	case "kategori":
		urutTransaksiKategori(d, n)
	case "tanggal":
		urutTransaksiTanggal(d, n)
	default:
		fmt.Println("Pilih opsi yang tersedia")
		return
	}

	switch y {
	case "asc":
		listTransaksi(*d, n)
	case "dsc":
		urutTransaksiDesc(d, n)
		listTransaksi(*d, n)
	default:
		fmt.Println("Pilih opsi yang tersedia")
		return
	}
}

func urutTransaksiNama(d *TabTransaksi, n int) {
	/*
		Initial state:
		i, j, min: tidak terdefinisi.
		temp: tidak terdefinisi.
		Final state:
		i, j: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari transaksi dengan nama terkecil.
		temp: menyimpan nilai dari array ke min yang akan ditukar nilainya
		dengan array ke i.
	*/

	var i, j, min int
	var temp Transaksi

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

func urutTransaksiHarga(d *TabTransaksi, n int) {
	/*
		Initial state:
		i, j, min: tidak terdefinisi.
		temp: tidak terdefinisi.
		Final state:
		i, j: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari transaksi dengan biaya terkecil.
		temp: menyimpan nilai dari array ke min yang akan ditukar nilainya
		dengan array ke i.
	*/

	var i, j, min int
	var temp Transaksi

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

func urutTransaksiKategori(d *TabTransaksi, n int) {
	/*
		Initial state:
		i, j, min: tidak terdefinisi.
		temp: tidak terdefinisi.
		Final state:
		i, j: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari transaksi dengan kategori terkecil.
		temp: menyimpan nilai dari array ke min yang akan ditukar nilainya
		dengan array ke i.
	*/

	var i, j, min int
	var temp Transaksi

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if d[j].Kategori < d[min].Kategori {
				min = j
			}
		}

		temp = d[min]
		d[min] = d[i]
		d[i] = temp
	}
}

func urutTransaksiTanggal(d *TabTransaksi, n int) {
	/*
		Initial state:
		i, j, min: tidak terdefinisi.
		temp: tidak terdefinisi.
		Final state:
		i, j: digunakan sebagai counter perulangan.
		min: menyimpan indeks dari transaksi dengan tanggal terkecil.
		temp: menyimpan nilai dari array ke min yang akan ditukar nilainya
		dengan array ke i.
	*/

	var i, j, min int
	var temp Transaksi

	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if d[j].Tanggal.Unix() < d[min].Tanggal.Unix() {
				min = j
			}
		}

		temp = d[min]
		d[min] = d[i]
		d[i] = temp
	}
}

func urutTransaksiDesc(d *TabTransaksi, n int) {
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
	var temp Transaksi

	i = 0
	for i < n/2 && d[i] != d[n-1-i] {
		temp = d[i]
		d[i] = d[n-1-i]
		d[n-1-i] = temp

		i++
	}
}

func simpanBatasTransaksi(b *int) {
	fmt.Print("Masukkan batas transaksi (0 jika tidak ingin mengatur batas) >> ")
	fmt.Scan(b)
}

func totalTransaksi(d *TabTransaksi, n int) int {
	var i, total int

	total = 0

	for i = 0; i < n; i++ {
		if d[i].Tanggal.Month() == time.Now().Month() {
			total += d[i].Biaya
		}
	}

	return total
}
