package main

import "fmt"

type Langganan struct {
	Nama    string
	Biaya   int
	Metode  string
	Tenggat int
}

const LMAX int = 10

type TabLangganan = [LMAX]Langganan

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
		Tenggat: 8,
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

		case 0:
			lanjut = false
		}
	}
}

func menu() {
	fmt.Println("[1] Cetak layanan berlangganan")
	fmt.Println("[2] Tambah layanan berlangganan")
	fmt.Println("[3] Ubah layanan berlangganan")
	fmt.Println("[4] Hapus layanan berlangganan")
	fmt.Println("[5] Cari layanan berlangganan")
	fmt.Println("[6] Urutkan layanan berlangganan")
	fmt.Println("[0] Keluar")
}

func cetakLangganan(d TabLangganan, n int) {
	fmt.Println("Daftar layanan berlangganan:")

	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s %d %s %d\n", i+1, d[i].Nama, d[i].Biaya, d[i].Metode, d[i].Tenggat)
	}
}

func cetakSatuLangganan(d *TabLangganan, n int) {
	var i int

	urutLanggananNama(d, n)

	i = cariLangganan(*d, n)

	fmt.Println("Nama layanan:", d[i].Nama)
	fmt.Println("Biaya layanan:", d[i].Biaya)
	fmt.Println("Metode Pembayaran layanan:", d[i].Metode)
	fmt.Println("Tenggat layanan tiap bulan:", d[i].Tenggat)
}

func tambahLangganan(d *TabLangganan, n *int) {
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

func urutLanggananNama(d *TabLangganan, n int) {
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
