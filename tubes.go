package main
import "fmt"

const max = 10
type barang struct{
	nama string
	harga int
	stok int
}
type produsen struct{
	nama string
	alamat string
	telp string
}
type tempat struct{
	nama string
	alamat string
	telp string
}

type tblBarang [max]barang
type tblProdusen [max]produsen
type tblTempat [max]tempat

func main() {
    var (
		menu int
		databarang tblBarang = tblBarang{
			barang{"Indomie",5000,4},
			barang{"Kecap",6000,10},
			barang{"Saos",3000,2},
			barang{"Sirup Marjan",15000,15},
			barang{"Buah",9000,13},
			barang{"Apel",3000,12},
			barang{"Jeruk",4000,9},
			barang{"Mangga",8000,19},
			barang{"Pepaya",10000,6},
			barang{"Nangka",12000,15},
		}

		dataprodusen tblProdusen = tblProdusen{
			produsen{"Indomie","5000","4"},
			produsen{"Kecap","6000","10"},
			produsen{"Saos","3000","2"},
			produsen{"Sirup Marjan","15000","15"},
			produsen{"Buah","9000","13"},
			produsen{"Apel","3000","12"},
			produsen{"Jeruk","4000","9"},
			produsen{"Mangga","8000","19"},
			produsen{"Pepaya","10000","6"},
			produsen{"Nangka","12000","15"},
		}

		datatempat tblTempat = tblTempat{
			tempat{"Indomie","5000","4"},
			tempat{"Kecap","6000","10"},
			tempat{"Saos","3000","2"},
			tempat{"Sirup Marjan","15000","15"},
			tempat{"Buah","9000","13"},
			tempat{"Apel","3000","12"},
			tempat{"Jeruk","4000","9"},
			tempat{"Mangga","8000","19"},
			tempat{"Pepaya","10000","6"},
			tempat{"Nangka","12000","15"},
		}
	)
   for {
   fmt.Println("=============================================")
   fmt.Println("|         SISTEM INFORMASI INVENTORY        |")
   fmt.Println("=============================================")
   fmt.Println("|                                           |")
   fmt.Println("|             1. Lihat Data                 |")
   fmt.Println("|              2. Ubah Data                 |")
   fmt.Println("|               3. Hapus Data               |")
   fmt.Println("|                4. Tambah Data             |")
   fmt.Println("|                                           |")
   fmt.Println("=============================================")
   for (menu!=1) && (menu!=2) && (menu!=3) && (menu!=4) {
	 fmt.Print("  Masukan menu pilihan anda : ")
	 fmt.Scan(&menu)
	 if (menu!=1) && (menu!=2) && (menu!=3) && (menu!=4) {
		fmt.Println("   *Pilihan tidak tersedia, Silahkan masukkan nomor yang valid!")
		fmt.Println(" ")
	 }
	}
	 if (menu==1) {
		tampilBarang(databarang)
		tampilProdusen(dataprodusen)
		tampilTempat(datatempat)
		fmt.Println(" ")
		menu=0
	 }else if (menu==2) {
		edit(databarang,dataprodusen,datatempat)
		fmt.Println(" ")
		menu=0
	 }else if (menu==3) {
		hapus(databarang,dataprodusen,datatempat)
		fmt.Println(" ")
		menu=0
	 }else if (menu==4) {
		fmt.Println(" ")
		menu=0
	 }
   }
}

// func tambah(data tblBarang)  {

// }

func tampilBarang(barang tblBarang)  {
	var no int
	no = 1
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|                Data Barang                |")
	fmt.Println("=============================================")

	for _, val := range barang {
		fmt.Println("No.",no, "=> Nama Barang :",val.nama,", Harga :",val.harga,", Stok :",val.stok)
		no=no+1
	}
}

func tampilProdusen(produsen tblProdusen)  {	
	var no int
	no = 1
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|               Data Produsen               |")
	fmt.Println("=============================================")

	for _, val := range produsen {
		fmt.Println("No.",no, "=> Nama Produsen :",val.nama,", Alamat :",val.alamat,", Telp :",val.telp)
		no=no+1
	}
}

func tampilTempat(tempat tblTempat)  {
	var no int
	no = 1
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|           Data Tempat Distribusi          |")
	fmt.Println("=============================================")

	for _, val := range tempat {
		fmt.Println("No.",no, "=> Nama Agen :",val.nama,", Alamat :",val.alamat,", Telp :",val.telp)
		no=no+1
	}
}

func edit(barang tblBarang, produsen tblProdusen, tempat tblTempat)  {
	var (
		pilih,index,harga,stok int
		nama,alamat,telp string
	)
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|   SISTEM INFORMASI INVENTORY / UBAH DATA  |")
	fmt.Println("=============================================")
	fmt.Println("|                                           |")
	fmt.Println("|                1. Data Barang             |")
	fmt.Println("|             2. Data Produsen              |")
	fmt.Println("|          3. Data Tempat Distribusi        |")
	fmt.Println("|                                           |")
	fmt.Println("=============================================")
	for (pilih!=1) && (pilih!=2) && (pilih!=3) {
		fmt.Print("  Pilih data mana yang akan diubah : ")
		fmt.Scan(&pilih)
		if (pilih!=1) && (pilih!=2) && (pilih!=3) {
			fmt.Println("   *Pilihan tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	if (pilih==1) {
		tampilBarang(barang)
	}else if (pilih==2) {
		tampilProdusen(produsen)
	}else if (pilih==3) {
		tampilTempat(tempat)
	}
	fmt.Println("=============================================")
	for (index < 1) || (index > max) {
		fmt.Print("  Masukkan nomor data yang akan diubah : ")
		fmt.Scan(&index)
		if (index < 1) || (index > max) {
			fmt.Println("   *Nomor",index,"tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	index=index-1
	fmt.Println("=============================================")
	if (pilih==1) {
		fmt.Print("Nama Barang : ")
		fmt.Scan(&nama)
		fmt.Print("Harga       : ")
		fmt.Scan(&harga)
		fmt.Print("Stok        : ")
		fmt.Scan(&stok)
		barang[index].nama = nama
		barang[index].harga = harga
		barang[index].stok = stok
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menyimpan data!")
		fmt.Println("")
		tampilBarang(barang)
	}else if (pilih==2) {
		fmt.Print("Nama Produsen : ")
		fmt.Scan(&nama)
		fmt.Print("Alamat        : ")
		fmt.Scan(&alamat)
		fmt.Print("Telepon       : ")
		fmt.Scan(&telp)
		produsen[index].nama = nama
		produsen[index].alamat = alamat
		produsen[index].telp = telp
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menyimpan data!")
		fmt.Println("")
		tampilProdusen(produsen)
	}else if (pilih==3) {
		fmt.Print("Nama Agen : ")
		fmt.Scan(&nama)
		fmt.Print("Alamat    : ")
		fmt.Scan(&alamat)
		fmt.Print("Telepon   : ")
		fmt.Scan(&telp)
		tempat[index].nama = nama
		tempat[index].alamat = alamat
		tempat[index].telp = telp
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menyimpan data!")
		fmt.Println("")
		tampilTempat(tempat)
	}
}

func hapus(barang tblBarang, produsen tblProdusen, tempat tblTempat)  {
	var (
		pilih,index int
	)
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|  SISTEM INFORMASI INVENTORY / HAPUS DATA  |")
	fmt.Println("=============================================")
	fmt.Println("|                                           |")
	fmt.Println("|                1. Data Barang             |")
	fmt.Println("|             2. Data Produsen              |")
	fmt.Println("|          3. Data Tempat Distribusi        |")
	fmt.Println("|                                           |")
	fmt.Println("=============================================")
	for (pilih!=1) && (pilih!=2) && (pilih!=3) {
		fmt.Print("  Pilih jenis data yang akan dihapus : ")
		fmt.Scan(&pilih)
		if (pilih!=1) && (pilih!=2) && (pilih!=3) {
			fmt.Println("   *Pilihan tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	if (pilih==1) {
		tampilBarang(barang)
	}else if (pilih==2) {
		tampilProdusen(produsen)
	}else if (pilih==3) {
		tampilTempat(tempat)
	}
	fmt.Println("=============================================")
	for (index < 1) || (index > max) {
		fmt.Print("  Masukkan nomor data yang akan dihapus : ")
		fmt.Scan(&index)
		if (index < 1) || (index > max) {
			fmt.Println("   *Nomor",index,"tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	index=index-1
	fmt.Println("=============================================")
	if (pilih==1) {
		barang[index].nama = " "
		barang[index].harga = 0
		barang[index].stok = 0
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menghapus data!")
		fmt.Println("")
		tampilBarang(barang)
	}else if (pilih==2) {
		produsen[index].nama = " "
		produsen[index].alamat = " "
		produsen[index].telp = " "
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menghapus data!")
		fmt.Println("")
		tampilProdusen(produsen)
	}else if (pilih==3) {
		tempat[index].nama = " "
		tempat[index].alamat = " "
		tempat[index].telp = " "
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menghapus data!")
		fmt.Println("")
		tampilTempat(tempat)
	}
}

// func cari(data tblBarang)  {

// }