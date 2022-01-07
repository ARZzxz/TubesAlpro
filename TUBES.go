package main

import "fmt"

const Nmax = 101

type Mahasiswa struct {
	nama  string
	nim   int
	nilai int
}

type arrM struct {
	tabmhs [Nmax]Mahasiswa
	n      int
}

func isiData_Mahasiswa(data *arrM, N int) { // ariz

	for i := data.n; i < data.n+N; i++ {
		fmt.Print("Masukan Nama : ")
		fmt.Scan(&data.tabmhs[i].nama)
		fmt.Print("Masukan NIM : ")
		fmt.Scan(&data.tabmhs[i].nim)
		fmt.Print("Masukan Nilai : ")
		fmt.Scan(&data.tabmhs[i].nilai)
	}
	data.n = data.n + N
}

/*func DelMahasiswa() { //ariz
	var idx_del, i int

	idx_del = Seq_CariMahasiswa()
}

/*func Seq_CariMahasiswa(data *arrM, nama string) {
	var a Mahasiswa
	var hasil int

	i:=0
	for i < data.n && data.[i].nama != nama {
		i = i + 1
	} if tab.T[i].nama == nama {
		return i
	} else {
		return -1
	}
}
/*
func Bin_() { apis

}

func insort_cariNilaiTinggi() { apis

}

func selsort_cariNilaiTerendah() { apis

}

*/

func Show_data(data arrM) { // ariz
	for i := 0; i < data.n; i++ {
		fmt.Println("Nama Mahasiswa : ", data.tabmhs[i].nama, "NIM Mahasiswa : ", data.tabmhs[i].nim, "Nilai Mahasiswa : ", data.tabmhs[i].nilai, "\n")
	}
}

func main() {
	var input, N int
	var ar arrM

	fmt.Println("==========MENU==========")
	fmt.Println("1. Tambahkan mahasiswa")
	fmt.Println("2. Tampilkan data seluruh mahasiswa")
	fmt.Println("0. Selesai")
	fmt.Println("====================")
	fmt.Print("PILIH MENU : ")
	fmt.Scan(&input)
	fmt.Println("")

	for input != 0 {
		if input == 1 {
			fmt.Scan(&N)
			isiData_Mahasiswa(&ar, N)
			fmt.Println("Data Telah ditambahkan!!")
		} else if input == 2 {
			Show_data(ar)
			fmt.Println("Ini data Mahasiswa!!")
		} /*else if input == 3 {
			Seq_CariMahasiswa(ar)
		}*/
		fmt.Println("==========MENU==========")
		fmt.Println("1. Tambahkan mahasiswa")
		fmt.Println("2. Tampilkan data seluruh mahasiswa")
		fmt.Println("0. Selesai")
		fmt.Println("====================")
		fmt.Print("PILIH MENU : ")
		fmt.Scan(&input)
	}
}
