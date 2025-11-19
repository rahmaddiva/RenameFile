package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// --- 1. SETUP FLAG (Argumen) ---
	dirPtr := flag.String("dir", ".", "Folder target (default: folder saat ini)")
	findPtr := flag.String("find", "", "Teks yang ingin dicari di nama file")
	replacePtr := flag.String("replace", "", "Teks pengganti")
	executePtr := flag.Bool("execute", false, "Jalankan rename (jika false, hanya preview/dry-run)")

	flag.Parse()

	// Validasi input
	if *findPtr == "" {
		fmt.Println("Error: Harap tentukan teks yang dicari dengan -find")
		fmt.Println("Contoh: go run main.go -find=\"IMG_\" -replace=\"Liburan_\"")
		return
	}

	// --- 2. BACA DIRECTORY ---
	files, err := os.ReadDir(*dirPtr)
	if err != nil {
		fmt.Printf("Gagal membaca folder: %v\n", err)
		return
	}

	fmt.Println("--- Memulai Proses ---")
	if !*executePtr {
		fmt.Println("MODE: PREVIEW (Gunakan -execute untuk menerapkan perubahan)")
	} else {
		fmt.Println("MODE: EKSEKUSI (Perubahan akan diterapkan)")
	}
	fmt.Println("----------------------")

	count := 0

	// --- 3. LOOPING FILE ---
	for _, file := range files {
		// Kita hanya ingin rename file, bukan folder
		if file.IsDir() {
			continue
		}

		originalName := file.Name()

		// Cek apakah nama file mengandung teks yang dicari
		if strings.Contains(originalName, *findPtr) {
			// Buat nama baru
			newName := strings.Replace(originalName, *findPtr, *replacePtr, 1)

			// Gabungkan dengan path folder untuk mendapatkan full path
			oldPath := filepath.Join(*dirPtr, originalName)
			newPath := filepath.Join(*dirPtr, newName)

			// --- 4. EKSEKUSI ATAU PREVIEW ---
			if *executePtr {
				// Lakukan Rename Asli
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("[GAGAL] %s -> %v\n", originalName, err)
				} else {
					fmt.Printf("[OK] %s -> %s\n", originalName, newName)
				}
			} else {
				// Hanya Preview (Dry Run)
				fmt.Printf("[PREVIEW] %s -> %s\n", originalName, newName)
			}
			count++
		}
	}

	if count == 0 {
		fmt.Println("Tidak ada file yang cocok ditemukan.")
	} else {
		fmt.Printf("\nSelesai. %d file diproses.\n", count)
	}
}