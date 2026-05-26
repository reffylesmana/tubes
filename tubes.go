package main

import "fmt"

const NMAX = 999

type resep struct {
	nama           string
	kategori 			 string
	bahan_utama    string
	bahan          [NMAX]string
	jumlah_bahan   int
	langkah        string
	estimasi_waktu int
	jumlah_dicari  int
}

type tabresep [NMAX]resep

func main() {
	var r tabresep
	var n int
	var pilihan, pilihan_kelola, pilihan_cari, pilihan_sort, pilihan_statistik int

	for {
		fmt.Println("\n┌───────────────────────────────────────┐")
		fmt.Println("│              MENU UTAMA               │")
		fmt.Println("├───────────────────────────────────────┤")
		fmt.Println("│ 1. Kelola Resep                       │")
		fmt.Println("│ 2. Cari Resep                         │")
		fmt.Println("│ 3. Tampilkan Semua Resep              │")
		fmt.Println("│ 4. Statistik Resep                    │")
		fmt.Println("│ 5. Keluar                             │")
		fmt.Println("└───────────────────────────────────────┘")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			for {
				fmt.Println("\n┌───────────────────────────────────────┐")
				fmt.Println("│             KELOLA RESEP              │")
				fmt.Println("├───────────────────────────────────────┤")
				fmt.Println("│ 1. Tambah Resep                       │")
				fmt.Println("│ 2. Edit Resep                         │")
				fmt.Println("│ 3. Hapus Resep                        │")
				fmt.Println("│ 4. Tampilkan Resep                    │")
				fmt.Println("│ 5. Kembali ke Menu Utama              │")
				fmt.Println("└───────────────────────────────────────┘")
				fmt.Print("Pilih aksi: ")
				fmt.Scan(&pilihan_kelola)

				if pilihan_kelola == 5 {
					break
				}

				switch pilihan_kelola {
				case 1:
					tambah(&r, &n)
				case 2:
					edit(&r, n)
				case 3:
					hapus(&r, &n)
				case 4:
					show(r, n)
				default:
					fmt.Println("\nPilihan tidak valid.")
				}
			}

		case 2:
			for {
				fmt.Println("\n┌───────────────────────────────────────┐")
				fmt.Println("│              CARI RESEP               │")
				fmt.Println("├───────────────────────────────────────┤")
				fmt.Println("│ 1. Binary Search                      │")
				fmt.Println("│ 2. Sequential Search                  │")
				fmt.Println("│ 3. Kembali ke Menu Utama              │")
				fmt.Println("└───────────────────────────────────────┘")
				fmt.Print("Pilih metode: ")
				fmt.Scan(&pilihan_cari)

				if pilihan_cari == 3 {
					break
				}

				switch pilihan_cari {
				case 1:
					binary(&r, n)
				case 2:
					sequential(&r, n)
				default:
					fmt.Println("\n Pilihan tidak valid.")
				}
			}

		case 3:
			for {
				fmt.Println("\n┌───────────────────────────────────────┐")
				fmt.Println("│         TAMPILKAN SEMUA RESEP         │")
				fmt.Println("├───────────────────────────────────────┤")
				fmt.Println("│ 1. Selection Sort                     │")
				fmt.Println("│ 2. Insertion Sort                     │")
				fmt.Println("│ 3. Kembali ke Menu Utama              │")
				fmt.Println("└───────────────────────────────────────┘")
				fmt.Print("Pilih algoritma sort: ")
				fmt.Scan(&pilihan_sort)

				if pilihan_sort == 3 {
					break
				}

				switch pilihan_sort {
				case 1:
					selection_sort(&r, n)
				case 2:
					fmt.Println("\n[Fitur Insertion Sort]")
				default:
					fmt.Println("\n Pilihan tidak valid.")
				}
			}

		case 4:
			for {
				fmt.Println("\n┌───────────────────────────────────────┐")
				fmt.Println("│            STATISTIK RESEP            │")
				fmt.Println("├───────────────────────────────────────┤")
				fmt.Println("│ 1. Jumlah Resep per Bahan             │")
				fmt.Println("│ 2. Resep Paling Sering Dicari         │")
				fmt.Println("│ 3. Kembali ke Menu Utama              │")
				fmt.Println("└───────────────────────────────────────┘")
				fmt.Print("Pilih statistik: ")
				fmt.Scan(&pilihan_statistik)

				if pilihan_statistik == 3 {
					break
				}

				switch pilihan_statistik {
				case 1:
					fmt.Println("\n[Fitur Jumlah Resep per Kategori Bahan]")
				case 2:
					fmt.Println("\n[Fitur Daftar Menu yang Paling Sering Dicari]")
				default:
					fmt.Println("\n⚠ Pilihan tidak valid.")
				}
			}

		case 5:
			fmt.Println("\n=========================================")
			fmt.Println(" Keluar dari program. Terima kasih!      ")
			fmt.Println("=========================================")
			return

		default:
			fmt.Println("\n Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func tambah(r *tabresep, n *int) {
	var i int
	var lagi string

	for {
		fmt.Println("\n=========================================")
		fmt.Printf("      INPUT DATA RESEP KE-%d\n", *n+1)
		fmt.Println("=========================================")

		fmt.Print("Masukan Nama Resep           : ")
		fmt.Scan(&(*r)[*n].nama)

		fmt.Print("Masukan Kategori Resep       : ")
		fmt.Scan(&(*r)[*n].kategori)

		fmt.Print("Masukan Bahan Utama          : ")
		fmt.Scan(&(*r)[*n].bahan_utama)

		fmt.Print("Masukan Jumlah Bahan         : ")
		fmt.Scan(&(*r)[*n].jumlah_bahan)

		for i = 0; i < (*r)[*n].jumlah_bahan; i++ {
			fmt.Printf("  -> Masukan Bahan ke-%d      : ", i+1)
			fmt.Scan(&(*r)[*n].bahan[i])
		}

		fmt.Print("Masukan Langkah              : ")
		fmt.Scan(&(*r)[*n].langkah)

		fmt.Print("Masukan Estimasi Waktu (Menit)       : ")
		fmt.Scan(&(*r)[*n].estimasi_waktu)

		*n++

		fmt.Println("-----------------------------------------")
		fmt.Print("Tambah data lagi? (y/n): ")
		fmt.Scan(&lagi)

		if lagi == "n" || lagi == "N" {
			return
		}
	}
}

func edit(r *tabresep, n int) {
	var cari string
	var i, j int
	var found bool

	if n == 0 {
		fmt.Println("\nData resep masih kosong.")
		return
	}

	fmt.Print("Masukkan Nama Resep yang Ingin Diedit: ")
	fmt.Scan(&cari)

	found = false
	i = 0

	for i < n && !found {
		if (*r)[i].nama == cari {
			found = true

			fmt.Println("\n=========================================")
			fmt.Println("           EDIT DATA RESEP")
			fmt.Println("=========================================")

			fmt.Print("Masukan Nama Resep Baru           : ")
			fmt.Scan(&(*r)[i].nama)

			fmt.Print("Masukan Kategori Resep Baru       : ")
			fmt.Scan(&(*r)[i].kategori)

			fmt.Print("Masukan Bahan Utama Baru          : ")
			fmt.Scan(&(*r)[i].bahan_utama)

			fmt.Print("Masukan Jumlah Bahan Baru   : ")
			fmt.Scan(&(*r)[i].jumlah_bahan)

			for j = 0; j < (*r)[i].jumlah_bahan; j++ {
				fmt.Printf("  -> Masukan Bahan ke-%d        : ", j+1)
				fmt.Scan(&(*r)[i].bahan[j])
			}

			fmt.Print("Masukan Langkah Baru               : ")
			fmt.Scan(&(*r)[i].langkah)

			fmt.Print("Masukan Estimasi Waktu Baru (Menit)       : ")
			fmt.Scan(&(*r)[i].estimasi_waktu)

			fmt.Println("\nData resep berhasil diubah.")
		}
		i++
	}

	if !found {
		fmt.Println("\nResep tidak ditemukan.")
	}
}

func hapus(r *tabresep, n *int) {
	var cari string
	var i, j int
	var found bool

	if *n == 0 {
		fmt.Println("\nData resep masih kosong.")
		return
	}

	fmt.Print("Masukkan Nama Resep yang Ingin Dihapus: ")
	fmt.Scan(&cari)

	found = false
	i = 0

	for i < *n && !found {
		if (*r)[i].nama == cari {
			found = true

			for j = i; j < *n-1; j++ {
				(*r)[j] = (*r)[j+1]
			}

			*n = *n - 1

			fmt.Println("\nData resep berhasil dihapus.")
		}
		i++
	}

	if !found {
		fmt.Println("\nResep tidak ditemukan.")
	}
}

func show(r tabresep, n int) {
	var i, j int
	var lagi string

	for {
		fmt.Println("\n=========================================")
		fmt.Println("             DAFTAR SEMUA RESEP          ")
		fmt.Println("=========================================")

		if n == 0 {
			fmt.Println("Resep belum ada.")
		} else {
			for i = 0; i < n; i++ {
				fmt.Printf("--- [Resep Ke-%d] ---\n", i+1)
				fmt.Printf("%-16s: %s\n", "Nama Resep", r[i].nama)
				fmt.Printf("%-16s: %s\n", "Kategori Resep", r[i].kategori)
				fmt.Printf("%-16s: %s\n", "Bahan Utama", r[i].bahan_utama)
				fmt.Printf("%-16s: %d\n", "Jumlah Bahan", r[i].jumlah_bahan)

				fmt.Printf("%-16s: ", "Bahan")
				for j = 0; j < r[i].jumlah_bahan; j++ {
					fmt.Print(r[i].bahan[j], " ")
				}
				fmt.Println()

				fmt.Printf("%-16s: %s\n", "Langkah", r[i].langkah)
				fmt.Printf("%-16s: %d menit\n", "Estimasi Waktu", r[i].estimasi_waktu)
				fmt.Println("-----------------------------------------")
			}
		}

		fmt.Print("\nApakah Ingin Tampilkan Data Lagi? (y/n): ")
		fmt.Scan(&lagi)
		if lagi == "n" || lagi == "N" {
			return
		}
	}
}

func sequential(r *tabresep, n int) { 
	var i, idx, j int
	var cari, lagi string
	var found bool

	for {
		found = false
		i = 0

		fmt.Print("Masukkan Bahan Utama yang Ingin Dicari: ")
		fmt.Scan(&cari)

		for i < n && !found {
			if (*r)[i].bahan_utama == cari {
				found = true
				idx = i
			}
			i++
		}

		if found {
			(*r)[idx].jumlah_dicari = (*r)[idx].jumlah_dicari + 1
			fmt.Printf("\n Data '%s' DITEMUKAN pada Indeks ke-%d\n", cari, idx)
			fmt.Println("=========================================")
			fmt.Println("Nama Resep     :", (*r)[idx].nama)
			fmt.Println("Kategori Resep :", (*r)[idx].kategori)
			fmt.Println("Bahan Utama    :", (*r)[idx].bahan_utama)
			fmt.Println("Jumlah Bahan   :", (*r)[idx].jumlah_bahan)
			fmt.Print("Bahan          : ")

			for j = 0; j < (*r)[idx].jumlah_bahan; j++ {
				fmt.Print((*r)[idx].bahan[j], " ")
			}

			fmt.Println("\nLangkah        :", (*r)[idx].langkah)
			fmt.Println("Estimasi Waktu :", (*r)[idx].estimasi_waktu, "menit")
			fmt.Println("=========================================")
		} else {
			fmt.Printf("Data '%s' Tidak Ditemukan\n", cari)
		}

		fmt.Print("\nApakah Ingin Cari Data Lagi? (y/n): ")
		fmt.Scan(&lagi)

		if lagi == "n" || lagi == "N" {
			return
		}

		fmt.Println()
	}
}


func binary(r *tabresep, n int) { 
	var right, mid, left, idx, j int
	var cari, lagi string
	var found bool

	urut_bahan_utama(r, n)

	for {
		left = 0
		right = n - 1
		found = false

		fmt.Print("Masukkan Bahan Utama yang Ingin Dicari: ")
		fmt.Scan(&cari)

		for left <= right && !found {
			mid = (left + right) / 2
			if (*r)[mid].bahan_utama == cari {
				found = true
				idx = mid
			} else if (*r)[mid].bahan_utama > cari {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		if found {
			(*r)[idx].jumlah_dicari = (*r)[idx].jumlah_dicari + 1
			fmt.Printf("\n Data '%s' DITEMUKAN pada Indeks ke-%d\n", cari, idx)
			fmt.Println("=========================================")
			fmt.Println("Nama Resep     :", (*r)[idx].nama)
			fmt.Println("Kategori Resep :", (*r)[idx].kategori)
			fmt.Println("Bahan Utama    :", (*r)[idx].bahan_utama)
			fmt.Println("Jumlah Bahan   :", (*r)[idx].jumlah_bahan)
			fmt.Print("Bahan          : ")
			for j = 0; j < (*r)[idx].jumlah_bahan; j++ {
				fmt.Print((*r)[idx].bahan[j], " ")
			}
			fmt.Println("\nLangkah        :", (*r)[idx].langkah)
			fmt.Println("Estimasi Waktu :", (*r)[idx].estimasi_waktu, "menit")
			fmt.Println("=========================================")
		} else {
			fmt.Printf("Data '%s' Tidak Ditemukan\n", cari)
		}

		fmt.Print("\nApakah Ingin Cari Data Lagi? (y/n): ")
		fmt.Scan(&lagi)
		if lagi == "n" || lagi == "N" {
			return
		}
		fmt.Println()
	}
}

func urut_bahan_utama(r * tabresep, n int){
	var i, j, idx int
	var temp resep

	for i = 0; i < n-1; i++{
		idx = i

		for j = i + 1; j < n; j++{
			if (*r)[j].bahan_utama < (*r)[idx].bahan_utama {
				idx = j
			}
		}

		if idx != i {
			temp = (*r)[i]
			(*r)[i] = (*r)[idx]
			(*r)[idx] = temp
		}
		
	}
}

func selection_sort(r * tabresep, n int) {
	var i, j, idx, k, m int 
	var temp resep

	for i = 0; i < n-1; i++ {
		idx = i

		for j = i + 1; j < n; j++ {
			if (*r)[j].estimasi_waktu < (*r)[idx].estimasi_waktu {
				idx = j
			}
		}

		if idx != i {
			temp = (*r)[i]
			(*r)[i] = (*r)[idx]
			(*r)[idx] = temp
		}
	}

	fmt.Println("\n=== DAFTAR RESEP SETELAH DIURUTKAN (ESTIMASI WAKTU) ===")
	for k = 0; k < n; k++ {
		fmt.Printf("%d. Nama Resep     : %s\n", k+1, (*r)[k].nama)
		fmt.Printf("   Kategori Resep : %s\n", (*r)[k].kategori)
		fmt.Printf("   Bahan Utama    : %s\n", (*r)[k].bahan_utama)
		fmt.Printf("   Jumlah Bahan   : %d\n", (*r)[k].jumlah_bahan)
		
		fmt.Print("   Bahan          : ")
		for m = 0; m < (*r)[k].jumlah_bahan; m++ {
			fmt.Print((*r)[k].bahan[m], " ")
		}
		fmt.Println()
		
		fmt.Printf("   Langkah        : %s\n", (*r)[k].langkah)
		fmt.Printf("   Estimasi Waktu : %d menit\n", (*r)[k].estimasi_waktu)
		fmt.Println("=======================================================")
	}
}


