package main
import"fmt"

const NMAX = 999

type resep struct{
     nama string
     bahan [NMAX] string
	 jumlah_bahan int
     langkah string
     estimasi_waktu int
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
		fmt.Print(" ❖ Pilih menu: ")
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
				fmt.Print(" ❖ Pilih aksi: ")
				fmt.Scan(&pilihan_kelola)
				
				if pilihan_kelola == 5 {
					break 
				}

				switch pilihan_kelola {
				case 1:
					tambah(&r, &n)
				case 2:
					fmt.Println("\n[Fitur Edit Resep]")
				case 3:
					fmt.Println("\n[Fitur Hapus Resep]")
				case 4:
					show(r, n)
				default:
					fmt.Println("\n⚠ Pilihan tidak valid.")
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
				fmt.Print(" ❖ Pilih metode: ")
				fmt.Scan(&pilihan_cari)
				
				if pilihan_cari == 3 {
					break 
				}

				switch pilihan_cari {
				case 1:
					binary(r, n)
				case 2:
					fmt.Println("\n[Fitur Sequential Search]")
				default:
					fmt.Println("\n⚠ Pilihan tidak valid.")
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
				fmt.Print(" ❖ Pilih algoritma sort: ")
				fmt.Scan(&pilihan_sort)
				
				if pilihan_sort == 3 {
					break
				}

				switch pilihan_sort {
				case 1:
					fmt.Println("\n[Fitur Selection Sort]")
				case 2:
					fmt.Println("\n[Fitur Insertion Sort]")
				default:
					fmt.Println("\n⚠ Pilihan tidak valid.")
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
				fmt.Print(" ❖ Pilih statistik: ")
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
			fmt.Println("\n⚠ Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func tambah(r *tabresep, n *int) {
	var i int
	var lagi string
	
	for {
		fmt.Println("\n=========================================")
		fmt.Printf("          INPUT DATA RESEP KE-%d          \n", *n+1)
		fmt.Println("=========================================")

		fmt.Print("Masukan Nama Resep     : ")
		fmt.Scan(&(*r)[*n].nama)
		
		fmt.Print("Masukan Jumlah Bahan   : ")
		fmt.Scan(&(*r)[*n].jumlah_bahan)
		
		for i = 0; i < r[*n].jumlah_bahan; i++ {
			fmt.Printf("  -> Masukan Bahan ke-%d : ", i+1)
			fmt.Scan(&r[*n].bahan[i])
		}
		
		fmt.Print("Masukan Langkah        : ")
		fmt.Scan(&(*r)[*n].langkah)
		
		fmt.Print("Masukan Estimasi Waktu (Menit) : ")
		fmt.Scan(&(*r)[*n].estimasi_waktu)
		*n = *n + 1

		fmt.Println("-----------------------------------------")
		fmt.Println("Apakah Ingin Masukan Data Lagi? (y/n): ")
		fmt.Scan(&lagi)
		if lagi == "n" || lagi == "N" {
			return
		}
	}
}

func show(r tabresep, n int) {
	var i, j int   
	var lagi string

	for {
		fmt.Println("\n=========================================")
		fmt.Println("             DAFTAR SEMUA RESEP          ")
		fmt.Println("=========================================")

		for i = 0; i < n; i++ {
			fmt.Printf("--- [Resep Ke-%d] ---\n", i+1)
			fmt.Printf("%-16s: %s\n", "Nama Resep", r[i].nama)
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

		fmt.Print("\nApakah Ingin Tampilkan Data Lagi? (y/n): ")
		fmt.Scan(&lagi)
		if lagi == "n" || lagi == "N" {
			return
		}
	}
}


func binary(r tabresep, n int) {
	var right, mid, left, idx, j int
	var cari, lagi string
	var found bool

	for {
		left = 0
		right = n - 1
		found = false

		fmt.Print("Masukkan Nama Resep yang Ingin Dicari: ")
		fmt.Scan(&cari)

		for left <= right && !found {
			mid = (left + right) / 2
			if r[mid].nama == cari {
				found = true
				idx = mid
			} else if r[mid].nama > cari {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		if found {
			fmt.Printf("\n✓ Data '%s' DITEMUKAN pada Indeks ke-%d\n", cari, idx)
			fmt.Println("=========================================")
			fmt.Println("Nama Resep     :", r[idx].nama)
			fmt.Println("Jumlah Bahan   :", r[idx].jumlah_bahan)
			fmt.Print("Bahan          : ")
			for j = 0; j < r[idx].jumlah_bahan; j++ {
				fmt.Print(r[idx].bahan[j], " ")
			}
			fmt.Println("\nLangkah        :", r[idx].langkah)
			fmt.Println("Estimasi Waktu :", r[idx].estimasi_waktu, "menit")
			fmt.Println("=========================================")
		} else {
			fmt.Printf("\n✗ Data '%s' Tidak Ditemukan\n", cari)
		}

		fmt.Print("\nApakah Ingin Cari Data Lagi? (y/n): ")
		fmt.Scan(&lagi)
		if lagi == "n" || lagi == "N" {
			return
		}
		fmt.Println()
	}
}

