package main

import "fmt"

type Negara struct {
	Nama           string
	JumlahEmas     int
	JumlahPerak    int
	JumlahPerunggu int
}

type Peringkat struct {
	NamaNegara     string
	JumlahEmas     int
	JumlahPerak    int
	JumlahPerunggu int
	TotalMedali    int
}

var daftarNegara [11]Negara
var peringkat [11]Peringkat

func main() {
	jumlahNegara := 0
	menu(&jumlahNegara)
}

func menu(jumlahNegara *int) {
	/*
		IS : Terdefinisi jumlahNegara (jumlah negara yang ada di array) sebagai masukan
		FS : Menampilkan menu yang interaktif
	*/

	for {
		fmt.Println("\n=== Aplikasi Pencatatan Medali SEA Games ===")
		fmt.Println("1. Tambah Negara")
		fmt.Println("2. Ubah Negara")
		fmt.Println("3. Hapus Negara")
		fmt.Println("4. Tampilkan Peringkat Negara")
		fmt.Println("5. Cari Negara Dari Peringkat")
		fmt.Println("6. Urutkan Negara Berdasarkan Alphabet")
		fmt.Println("0. Keluar")
		fmt.Println("============================================")
		fmt.Print("Pilih menu: ")

		var menu int
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			var nama string
			var emas, perak, perunggu int
			fmt.Print("Nama Negara: ")
			fmt.Scanln(&nama)
			fmt.Print("Jumlah Emas: ")
			fmt.Scanln(&emas)
			fmt.Print("Jumlah Perak: ")
			fmt.Scanln(&perak)
			fmt.Print("Jumlah Perunggu: ")
			fmt.Scanln(&perunggu)
			*jumlahNegara = tambahNegara(nama, emas, perak, perunggu, *jumlahNegara)
		case 2:
			var nama string
			var emas, perak, perunggu int
			fmt.Print("Nama Negara: ")
			fmt.Scanln(&nama)
			fmt.Print("Jumlah Emas baru: ")
			fmt.Scanln(&emas)
			fmt.Print("Jumlah Perak baru: ")
			fmt.Scanln(&perak)
			fmt.Print("Jumlah Perunggu baru: ")
			fmt.Scanln(&perunggu)
			ubahNegara(nama, emas, perak, perunggu, *jumlahNegara)
		case 3:
			var nama string
			fmt.Print("Nama Negara: ")
			fmt.Scanln(&nama)
			*jumlahNegara = hapusNegara(nama, *jumlahNegara)
		case 4:
			tampilkanPeringkat(*jumlahNegara)
		case 5:
			var rank int
			fmt.Print("Masukan Peringkat: ")
			fmt.Scan(&rank)
			temukanNegaraDariPeringkat(rank, *jumlahNegara)
		case 6:
			urutkanNegaraAlphabet(*jumlahNegara)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func tambahNegara(nama string, emas, perak, perunggu, jumlahNegara int) int {

	/*
		 	Fungsi untuk menambahkan negara ke dalam array jika masih ada slot yang tersedia.
		 	Parameter:
				- nama: nama negara
		 		- emas: jumlah medali emas
		 		- perak: jumlah medali perak
		 		- perunggu: jumlah medali perunggu
		 		- jumlahNegara: jumlah negara saat ini dalam daftar
		 	Return:
		 		- Jumlah negara yang diperbarui dalam array
					Jika masih tersedia slot maka negara akan ditambahkan, menambah isi variabel jumlahNegara, dan mencetak "Negara ditambahkan."
					Jika slot sudah penuh maka varibel jumlahNegara tetap dan mencetak "Kapasitas negara penuh."
	*/

	if jumlahNegara < len(daftarNegara) {
		daftarNegara[jumlahNegara] = Negara{Nama: nama, JumlahEmas: emas, JumlahPerak: perak, JumlahPerunggu: perunggu}
		jumlahNegara++
		fmt.Println("Negara ditambahkan.")
	} else {
		fmt.Println("Kapasitas negara penuh.")
	}
	return jumlahNegara
}

func ubahNegara(nama string, emas, perak, perunggu, jumlahNegara int) {

	/*
		IS: Terdefinisi nama (nama sebuah negara), emas (jumlah medali emas), perak (jumlah medali perak), perunggu (jumlah medali perunggu),
			dan jumlahNegara (jumlah negara yang ada di array) sebagai masukan.
		FS: Mengupdate jumlah medali (emas, perak, perunggu) untuk negara yang ditentukan dalam array global 'daftarNegara'.
			Pencarian menggunakan sequential search.
			Jika negara ditemukan dalam array, fungsi ini mengupdate jumlah medali dan mencetak "Data Negara diubah.".
			Jika negara tidak ditemukan, fungsi ini mencetak "Negara tidak ditemukan.".
	*/

	for i := 0; i < jumlahNegara; i++ {
		if daftarNegara[i].Nama == nama {
			daftarNegara[i].JumlahEmas = emas
			daftarNegara[i].JumlahPerak = perak
			daftarNegara[i].JumlahPerunggu = perunggu
			fmt.Println("Data Negara diubah.")
			return
		}
	}
	fmt.Println("Negara tidak ditemukan.")
}

func hapusNegara(nama string, jumlahNegara int) int {

	/*
			Fungsi untuk menghapus negara di dalam array.
		 	Parameter:
		 		- nama: nama negara
		 		- jumlahNegara: jumlah negara saat ini dalam array
		 	Return:
				- Jumlah negara yang diperbarui dalam array
					Pencarian menggunakan sequential search.
					Jika negara ditemukan dalam array, fungsi ini menghapus negara tersebut, mengurangi isi variabel jumlahNegara, dan mencetak "Negara dihapus."
					Jika negara tidak ditemukan, maka varibel jumlahNegara tetap dan mencetak "Negara tidak ditemukan.".
	*/

	for i := 0; i < jumlahNegara; i++ {
		if daftarNegara[i].Nama == nama {
			for j := i; j < jumlahNegara-1; j++ {
				daftarNegara[j] = daftarNegara[j+1]
			}
			jumlahNegara--
			fmt.Println("Negara dihapus.")
			return jumlahNegara
		}
	}
	fmt.Println("Negara tidak ditemukan.")
	return jumlahNegara
}

func updatePeringkat(jumlahNegara int) {

	/*
				IS: Terdefinisi jumlahNegara (jumlah negara yang ada di array) sebagai masukan
				FS: Mengupdate peringkat negara berdasarkan jumlah total medali (emas, perak, dan perunggu) dalam array global 'daftarNegara'.
					Pertama, prosedur menghitung total medali untuk setiap negara dan menyimpannya dalam struct Peringkat.
		 			Kemudian, prosedur mengurutkan array 'peringkat' menggunakan Insertion Sort berdasarkan total medali dari yang tertinggi ke yang terendah.
	*/

	for i := 0; i < jumlahNegara; i++ {
		totalMedali := daftarNegara[i].JumlahEmas + daftarNegara[i].JumlahPerak + daftarNegara[i].JumlahPerunggu
		peringkat[i] = Peringkat{
			NamaNegara:     daftarNegara[i].Nama,
			JumlahEmas:     daftarNegara[i].JumlahEmas,
			JumlahPerak:    daftarNegara[i].JumlahPerak,
			JumlahPerunggu: daftarNegara[i].JumlahPerunggu,
			TotalMedali:    totalMedali,
		}
	}

	for i := 0; i < jumlahNegara; i++ {
		for j := i + 1; j < jumlahNegara; j++ {
			if peringkat[j].TotalMedali > peringkat[i].TotalMedali {
				peringkat[i], peringkat[j] = peringkat[j], peringkat[i]
			}
		}
	}
}

func tampilkanPeringkat(jumlahNegara int) {

	/*
				IS: Terdefinisi jumlahNegara (jumlah negara yang ada dalam array) sebagai masukan
				FS: Menampilkan daftar peringkat negara peserta SEA Games berdasarkan jumlah total medali (emas, perak, dan perunggu).
					Pertama, prosedur ini memanggil prosedur updatePeringkat untuk memperbarui peringkat negara.
		 			Kemudian, prosedur ini mencetak daftar peringkat negara dengan format yang ditentukan.

	*/

	updatePeringkat(jumlahNegara)
	fmt.Println("Daftar Peringkat Negara Peserta SEA Games:")
	for i := 0; i < jumlahNegara; i++ {
		fmt.Printf("%d. %s - Emas: %d, Perak: %d, Perunggu: %d, Total: %d medali\n",
			i+1, peringkat[i].NamaNegara, peringkat[i].JumlahEmas, peringkat[i].JumlahPerak, peringkat[i].JumlahPerunggu, peringkat[i].TotalMedali)
	}
}

func temukanNegaraDariPeringkat(rank, jumlahNegara int) {

	/* 	IS: Terdefinisi rank (peringkat sebuah negara) dan jumlahNegara (jumlah total negara) sebagai masukan.
		FS: Mencari dan menampilkan nama negara yang memiliki peringkat sesuai yang diminta menggunakan Binary Search.
	      	1. Memeriksa apakah peringkat yang dimasukkan valid (dalam rentang 1 sampai jumlahNegara).
	        Jika tidak valid, mencetak pesan dan keluar dari prosedur.
	      	2. Memperbarui peringkat negara berdasarkan jumlah total medali dengan memanggil prosedur updatePeringkat.
	      	3. Menggunakan pencarian biner untuk menemukan negara dengan peringkat yang diminta.
	      	4. Jika ditemukan, mencetak nama negara tersebut. Jika tidak ditemukan, mencetak pesan bahwa negara tidak ditemukan.
	*/

	if rank < 1 || rank > jumlahNegara {
		fmt.Println("Peringkat tidak valid.")
		return
	}

	updatePeringkat(jumlahNegara)

	low, high := 0, jumlahNegara-1

	for low <= high {
		mid := (low + high) / 2
		if mid+1 == rank {
			fmt.Println("Nama Negara:", peringkat[mid].NamaNegara)
			return
		} else if mid+1 < rank {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Negara dengan peringkat tersebut tidak ditemukan.")
}

func urutkanNegaraAlphabet(jumlahNegara int) {

	/*
		IS: Terdefinisi jumlahNegara (jumlah negara yang ada dalam array) sebagai masukan
		FS: Mengurutkan array global 'daftarNegara' berdasarkan nama negara secara alfabetis menggunakan Selection Sort.
			Kemudian, prosedur mencetak daftar negara yang telah diurutkan beserta jumlah medali mereka.
	*/

	for i := 1; i < jumlahNegara; i++ {
		idx_min := daftarNegara[i]
		j := i - 1
		for j >= 0 && daftarNegara[j].Nama > idx_min.Nama {
			daftarNegara[j+1] = daftarNegara[j]
			j--
		}
		daftarNegara[j+1] = idx_min
	}

	fmt.Println("Daftar Negara diurutkan berdasarkan alfabet:")
	for i := 0; i < jumlahNegara; i++ {
		fmt.Printf("%d. %s - Emas: %d, Perak: %d, Perunggu: %d\n",
			i+1, daftarNegara[i].Nama, daftarNegara[i].JumlahEmas, daftarNegara[i].JumlahPerak, daftarNegara[i].JumlahPerunggu)
	}
}
