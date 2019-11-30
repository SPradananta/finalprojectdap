/**    DATA KELOMPOK */
/*     Judul 		: Aplikasi Inventori 2 
       Anggota Kelompok : 1301190243 Muhammad Satria Pradananta
	                      1301190437 Rifqi Ramadhan
*/
package main
import (
	"fmt"
	"time"
	"os"
	"os/exec"
)
const max = 100
var (
	markbarang int = 10
	markprodusen int = 10
	marktempat int = 10
	pilih int
)
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

func clear() {
	cmd := exec.Command("cmd","/c","cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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
			produsen{"Pt_Indofood","Jakarta","(022)5231799"},
			produsen{"Pt_Prima_Sinar_Mulia","Bandung","(022)5231819"},
			produsen{"Pt_Sky_Food_Industry","Bekasi","(021)8250311"},
			produsen{"Pt_URC_Indonesia","Cikarang","(021)89982585"},
			produsen{"Pt_Simba_Indosnack_Makmur","Bogor","(021)8674818"},
			produsen{"Pt_Alison_Agung","Jakarta_Barat","(021)6294630"},
			produsen{"Pt_Djaya_Giri_Citra_Lestari","Jakarta_Barat","(021)6253448"},
			produsen{"Pt_Rasa_Mutu_Utama","Bogor","(021)8670360"},
			produsen{"Pada_Suka","Jakbar","(021)63861335"},
			produsen{"CV_Sari_Rezeki","Kampung_Krendang","(021)6384227"},
		}

		datatempat tblTempat = tblTempat{
			tempat{"OshaSnack","Bogor","(022)5231799"},
			tempat{"PontusBakery","Cengkareng","(022)5231819"},
			tempat{"Sokita","Bandung","(021)6384227"},
			tempat{"StarindoGemilang","Solo","(0271)728046"},
			tempat{"CentralSnack","Pagarsih","(022)6012739"},
			tempat{"DailySnack","Citarum","(022)92839009"},
			tempat{"SnackShop","PasirKaliki","(022)6004950"},
			tempat{"DistributorKeripikMaicih","SumurBandung","(087)880316160"},
			tempat{"PDSenang","Bandung","(022) 6032126"},
			tempat{"WaroengJajananKitaKita","Dipatiukur","(085)624276856"},
		}
	)
	clear()
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
		clear()
		tampil(databarang,dataprodusen,datatempat)
		fmt.Println(" ")
		menu=0
	 }else if (menu==2) {
		clear()
		edit(&databarang,&dataprodusen,&datatempat)
		fmt.Println(" ")
		menu=0
	 }else if (menu==3) {
		clear()
		hapus(&databarang,&dataprodusen,&datatempat)
		fmt.Println(" ")
		menu=0
	 }else if (menu==4) {
		clear()
		tambah(&databarang,&dataprodusen,&datatempat)
		fmt.Println(" ")
		menu=0
	 }
   }
}

func secondmenu(judul,ket string)  {
	pilih=0
	fmt.Println("=============================================")
	fmt.Println(judul)
	fmt.Println("=============================================")
	fmt.Println("|                                           |")
	fmt.Println("|                1. Data Barang             |")
	fmt.Println("|             2. Data Produsen              |")
	fmt.Println("|          3. Data Tempat Distribusi        |")
	fmt.Println("|                 4. Kembali                |")
	fmt.Println("|                                           |")
	fmt.Println("=============================================")
	for (pilih!=1) && (pilih!=2) && (pilih!=3) && (pilih!=4) {
		fmt.Print("  Pilih data mana yang akan ",ket," : ")
		fmt.Scan(&pilih)
		if (pilih!=1) && (pilih!=2) && (pilih!=3) && (pilih!=4) {
			fmt.Println("   *Pilihan tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	if (pilih==4) {
		clear()
		main()
	}
}

func tambah(barang *tblBarang, produsen *tblProdusen, tempat *tblTempat)  {
	var (
		harga,stok int
		nama,alamat,telp string
	)
	secondmenu("|  SISTEM INFORMASI INVENTORY / TAMBAH DATA |","ditambah")
	clear()
	fmt.Println("=============================================")
	fmt.Println("    Silahkan Isi Form Untuk Menambah Data    ")
	fmt.Println("=============================================")
	fmt.Println(" ")
	if (pilih==1) && (markbarang < max) {
		fmt.Print("Nama Barang : ")
		fmt.Scan(&nama)
		fmt.Print("Harga       : ")
		fmt.Scan(&harga)
		fmt.Print("Stok        : ")
		fmt.Scan(&stok)
		barang[markbarang].nama = nama
		barang[markbarang].harga = harga
		barang[markbarang].stok = stok
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menyimpan data!")
		fmt.Println("")
		markbarang=markbarang+1
		time.Sleep(1 * time.Second)
		clear()
		tampilBarang(*barang)
	}else if (pilih==2) && (markprodusen < max) {
		fmt.Print("Nama Produsen : ")
		fmt.Scan(&nama)
		fmt.Print("Alamat        : ")
		fmt.Scan(&alamat)
		fmt.Print("Telepon       : ")
		fmt.Scan(&telp)
		produsen[markprodusen].nama = nama
		produsen[markprodusen].alamat = alamat
		produsen[markprodusen].telp = telp
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menyimpan data!")
		fmt.Println("")
		markprodusen=markprodusen+1
		time.Sleep(1 * time.Second)
		clear()
		tampilProdusen(*produsen)
	}else if (pilih==3) && (marktempat < max) {
		fmt.Print("Nama Agen : ")
		fmt.Scan(&nama)
		fmt.Print("Alamat    : ")
		fmt.Scan(&alamat)
		fmt.Print("Telepon   : ")
		fmt.Scan(&telp)
		tempat[marktempat].nama = nama
		tempat[marktempat].alamat = alamat
		tempat[marktempat].telp = telp
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menyimpan data!")
		fmt.Println("")
		marktempat=marktempat+1
		time.Sleep(1 * time.Second)
		clear()
		tampilTempat(*tempat)
	}else if (markbarang == max) || (markprodusen == max) || (marktempat == max){
		fmt.Println(" Maaf, Kapasitas Data Sudah Tidak Tersedia.")
		fmt.Println("  *Silahkan hapus sebagian data!")
	}
}

func tampil(barang tblBarang, produsen tblProdusen, tempat tblTempat)  {
	var stats int
	secondmenu("|  SISTEM INFORMASI INVENTORY / TAMPIL DATA |","ditampilkan")
	clear()
	if (pilih==1) {
		stats=1
		tampilBarang(barang)
	}else if (pilih==2) {
		stats=2
		tampilProdusen(produsen)
	}else if (pilih==3) {
		stats=3
		tampilTempat(tempat)
	}
	fmt.Println(" ")
	fmt.Println("---------------------------------------------")
	fmt.Println("| Cari Data (1) | Urutkan (2) | Kembali (3) |")
	fmt.Println("---------------------------------------------")
	pilih=0
	for (pilih!=1) && (pilih!=2) && (pilih!=3) {
		fmt.Print("  Masukkan menu pilihan anda : ")
		fmt.Scan(&pilih)
		if (pilih!=1) && (pilih!=2) && (pilih!=3) {
			fmt.Println("   *Pilihan tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	if (pilih==1) {
		cari(barang,produsen,tempat,stats)
	}else if (pilih==2) {
		urut(barang,produsen,tempat,stats)
	}else if (pilih==3) {
		clear()
		tampil(barang,produsen,tempat)
	}
}

func tampilBarang(data tblBarang)  {
	var no int
	no = 1
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|                Data Barang                |")
	fmt.Println("=============================================")

	for i:=0;i<markbarang;i++ {
		fmt.Println("No.",no, "=> Nama Barang :",data[i].nama,", Harga :",data[i].harga,", Stok :",data[i].stok)
		no=no+1
	}
}

func tampilProdusen(data tblProdusen)  {	
	var no int
	no = 1
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|               Data Produsen               |")
	fmt.Println("=============================================")

	for i:=0;i<markprodusen;i++ {
		fmt.Println("No.",no, "=> Nama Produsen :",data[i].nama,", Alamat :",data[i].alamat,", Telp :",data[i].telp)
		no=no+1
	}
}

func tampilTempat(data tblTempat)  {
	var no int
	no = 1
	fmt.Println(" ")
	fmt.Println("=============================================")
	fmt.Println("|           Data Tempat Distribusi          |")
	fmt.Println("=============================================")

	for i:=0;i<marktempat;i++ {
		fmt.Println("No.",no, "=> Nama Agen :",data[i].nama,", Alamat :",data[i].alamat,", Telp :",data[i].telp)
		no=no+1
	}
}

func edit(barang *tblBarang, produsen *tblProdusen, tempat *tblTempat)  {
	var (
		marker,index,harga,stok int
		nama,alamat,telp string
	)
	secondmenu("|   SISTEM INFORMASI INVENTORY / UBAH DATA  |","diubah")
	clear()
	if (pilih==1) {
		marker=markbarang
		tampilBarang(*barang)
	}else if (pilih==2) {
		marker=markprodusen
		tampilProdusen(*produsen)
	}else if (pilih==3) {
		marker=marktempat
		tampilTempat(*tempat)
	}
	fmt.Println("=============================================")
	for (index < 1) || (index > marker) {
		fmt.Print("  Masukkan nomor data yang akan diubah : ")
		fmt.Scan(&index)
		if (index < 1) || (index > marker) {
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
		time.Sleep(1 * time.Second)
		clear()
		tampilBarang(*barang)
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
		time.Sleep(1 * time.Second)
		clear()
		tampilProdusen(*produsen)
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
		time.Sleep(1 * time.Second)
		clear()
		tampilTempat(*tempat)
	}
}

func hapus(barang *tblBarang, produsen *tblProdusen, tempat *tblTempat)  {
	var index,marker int
	secondmenu("|  SISTEM INFORMASI INVENTORY / HAPUS DATA  |","dihapus")
	clear()
	if (pilih==1) {
		marker=markbarang
		tampilBarang(*barang)
	}else if (pilih==2) {
		marker=markprodusen
		tampilProdusen(*produsen)
	}else if (pilih==3) {
		marker=marktempat
		tampilTempat(*tempat)
	}
	fmt.Println("=============================================")
	for (index < 1) || (index > marker) {
		fmt.Print("  Masukkan nomor data yang akan dihapus : ")
		fmt.Scan(&index)
		if (index < 1) || (index > marker) {
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
		time.Sleep(1 * time.Second)
		clear()
		tampilBarang(*barang)
	}else if (pilih==2) {
		produsen[index].nama = " "
		produsen[index].alamat = " "
		produsen[index].telp = " "
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menghapus data!")
		time.Sleep(1 * time.Second)
		clear()
		tampilProdusen(*produsen)
	}else if (pilih==3) {
		tempat[index].nama = " "
		tempat[index].alamat = " "
		tempat[index].telp = " "
		fmt.Println(" ")
		fmt.Println("  *Berhasil Menghapus data!")
		time.Sleep(1 * time.Second)
		clear()
		tampilTempat(*tempat)
	}
}

func menucarisort(ket1,ket2,ket3,ket4,ket5 string){
	fmt.Println("=============================================")
	fmt.Println("|           ",ket4,"Data Berdasarkan          |")
	fmt.Println("=============================================")
	fmt.Println("|                                           |")
	fmt.Println(ket1)
	fmt.Println(ket2)
	fmt.Println(ket3)
	fmt.Println("|          4. Kembali Ke Menu Utama         |")
	fmt.Println("|                                           |")
	fmt.Println("=============================================")
	for (pilih!=1) && (pilih!=2) && (pilih!=3) && (pilih!=4) {
		fmt.Print("  Pilih berdasarkan apa data akan ",ket5," : ")
		fmt.Scan(&pilih)
		if (pilih!=1) && (pilih!=2) && (pilih!=3) && (pilih!=4) {
			fmt.Println("   *Pilihan tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	if (pilih==4) {
		clear()
		main()
	}
}

func cari(barang tblBarang, produsen tblProdusen, tempat tblTempat,datastat int)  {
	var (
		no,i,key int
		keystr string
		found bool
	)
	pilih=0
	clear()
	found=false
	if (datastat==1) {
		menucarisort("|               1. Nama Barang              |","|               2. Harga Barang             |","|               3. Stok Barang              |","Cari","dicari")
		clear()
		for found==false{
			fmt.Println("=============================================")
			fmt.Print(" Masukkan Kata Kunci : ")
				if (pilih==1) {
					i=0
					no=1
					fmt.Scan(&keystr)
					fmt.Println("=============================================")
					for (i<markbarang){
						if barang[i].nama==keystr {
							fmt.Println("No.",no, "=> Nama Barang :",barang[i].nama,", Harga :",barang[i].harga,", Stok :",barang[i].stok)
							no=no+1
							found=true
						}
						i=i+1
					}
				}else if (pilih==2) {
					i=0
					no=1
					fmt.Scan(&key)
					fmt.Println("=============================================")
					for (i<markbarang){
						if barang[i].harga==key {
							fmt.Println("No.",no, "=> Nama Barang :",barang[i].nama,", Harga :",barang[i].harga,", Stok :",barang[i].stok)
							no=no+1
							found=true
						}
						i=i+1
					}
				}else if (pilih==3) {
					i=0
					no=1
					fmt.Scan(&key)
					fmt.Println("=============================================")
					for (i<markbarang){
						if barang[i].stok==key {
							fmt.Println("No.",no, "=> Nama Barang :",barang[i].nama,", Harga :",barang[i].harga,", Stok :",barang[i].stok)
							no=no+1
							found=true
						}
						i=i+1
					}
				}
			if (found==true) {
				fmt.Println("_____________________________________________")
				fmt.Println(no-1,"hasil pencarian data telah ditemukan.")
			}else {
				fmt.Println(" Maaf, Kami Tidak Menemukan Data Apapun Disini.")
				fmt.Println("   *Silahkan Periksa Kata kunci yang anda masukkan.")
				time.Sleep(3 * time.Second)
				clear()
			}
		}	
	}else if (datastat==2) {
		menucarisort("|               1. Nama Produsen            |","|               2. Alamat Produsen          |","|               3. Telepon Produsen         |","Cari","dicari")
		clear()
		for found==false{
			fmt.Println("=============================================")
			fmt.Print(" Masukkan Kata Kunci : ")
				if (pilih==1) {
					i=0
					no=1
					fmt.Scan(&keystr)
					fmt.Println("=============================================")
					for (i<markprodusen){
						if produsen[i].nama==keystr {
							fmt.Println("No.",no, "=> Nama Produsen :",produsen[i].nama,", Alamat :",produsen[i].alamat,", Telp :",produsen[i].telp)
							no=no+1
							found=true
						}
						i=i+1
					}
				}else if (pilih==2) {
					i=0
					no=1
					fmt.Scan(&keystr)
					fmt.Println("=============================================")
					for (i<markprodusen){
						if produsen[i].alamat==keystr {
							fmt.Println("No.",no, "=> Nama Produsen :",produsen[i].nama,", Alamat :",produsen[i].alamat,", Telp :",produsen[i].telp)
							no=no+1
							found=true
						}
						i=i+1
					}
				}else if (pilih==3) {
					i=0
					no=1
					fmt.Scan(&keystr)
					fmt.Println("=============================================")
					for (i<markprodusen){
						if produsen[i].telp==keystr {
							fmt.Println("No.",no, "=> Nama produsen :",produsen[i].nama,", Alamat :",produsen[i].alamat,", Telp :",produsen[i].telp)
							no=no+1
							found=true
						}
						i=i+1
					}
				}
			if (found==true) {
				fmt.Println("_____________________________________________")
				fmt.Println(no-1,"hasil pencarian data telah ditemukan.")
			}else {
				fmt.Println(" Maaf, Kami Tidak Menemukan Data Apapun Disini.")
				fmt.Println("   *Silahkan Periksa Kata kunci yang anda masukkan.")
				time.Sleep(3 * time.Second)
				clear()
			}
		}
	}else if (datastat==3) {
		menucarisort("|               1. Nama Agen                |","|               2. Alamat Agen              |","|               3. Telepon Agen             |","Cari","dicari")
		clear()
		for found==false{
			fmt.Println("=============================================")
			fmt.Print(" Masukkan Kata Kunci : ")
				if (pilih==1) {
					i=0
					no=1
					fmt.Scan(&keystr)
					fmt.Println("=============================================")
					for (i<marktempat){
						if tempat[i].nama==keystr {
							fmt.Println("No.",no, "=> Nama Agen :",tempat[i].nama,", Alamat :",tempat[i].alamat,", Telp :",tempat[i].telp)
							no=no+1
							found=true
						}
						i=i+1
					}
				}else if (pilih==2) {
					i=0
					no=1
					fmt.Scan(&keystr)
					fmt.Println("=============================================")
					for (i<marktempat){
						if tempat[i].alamat==keystr {
							fmt.Println("No.",no, "=> Nama Agen :",tempat[i].nama,", Alamat :",tempat[i].alamat,", Telp :",tempat[i].telp)
							no=no+1
							found=true
						}
						i=i+1
					}
				}else if (pilih==3) {
					i=0
					no=1
					fmt.Scan(&keystr)
					fmt.Println("=============================================")
					for (i<marktempat){
						if tempat[i].telp==keystr {
							fmt.Println("No.",no, "=> Nama Agen :",tempat[i].nama,", Alamat :",tempat[i].alamat,", Telp :",tempat[i].telp)
							no=no+1
							found=true
						}
						i=i+1
					}
				}
			if (found==true) {
				fmt.Println("_____________________________________________")
				fmt.Println(no-1,"hasil pencarian data telah ditemukan.")
			}else {
				fmt.Println(" Maaf, Kami Tidak Menemukan Data Apapun Disini.")
				fmt.Println("   *Silahkan Periksa Kata kunci yang anda masukkan.")
				time.Sleep(3 * time.Second)
				clear()
			}
		}
	}
}

func menuurut(urutstat *int){
	fmt.Println("=============================================")
	fmt.Println("|           Mengrutkan Data Menurut         |")
	fmt.Println("=============================================")
	fmt.Println("|                                           |")
	fmt.Println("|               1. Ascending                |")
	fmt.Println("|               2. Descending               |")
	fmt.Println("|                                           |")
	fmt.Println("=============================================")
	for (*urutstat!=1) && (*urutstat!=2) {
		fmt.Print("  Pilih menurut apa data akan diurutkan : ")
		fmt.Scan(&*urutstat)
		if (*urutstat!=1) && (*urutstat!=2) {
			fmt.Println("   *Pilihan tidak tersedia, Silahkan masukkan nomor yang valid!")
			fmt.Println(" ")
		}
	}
	clear()
}

func urut(barang tblBarang, produsen tblProdusen, tempat tblTempat,datastat int)  {
	var (
		idx,temp,urutstat int
		tempstr string
	)

	pilih=0
	clear()
	if (datastat==1) {
		menucarisort("|               1. Nama Barang              |","|               2. Harga Barang             |","|               3. Stok Barang              |","Urut","diurutkan")
		clear()
		menuurut(&urutstat)
		//Ascending (Selection Sort)
		if urutstat==1 {
			if (pilih==1) {
				for i:=0;i < markbarang;i++ {
					idx = i
					for k:=i;k < markbarang;k++ {
						if barang[k].nama < barang[idx].nama {
							idx = k
						}
					}
					//Geser Nama
					tempstr=barang[i].nama
					barang[i].nama=barang[idx].nama
					barang[idx].nama=tempstr
					//Geser Harga
					temp=barang[i].harga
					barang[i].harga=barang[idx].harga
					barang[idx].harga=temp
					//Geser Stok
					temp=barang[i].stok
					barang[i].stok=barang[idx].stok
					barang[idx].stok=temp
				}
				tampilBarang(barang)

			}else if (pilih==2) {
				for i:=0;i < markbarang;i++ {
					idx = i
					for k:=i;k < markbarang;k++ {
						if barang[k].harga < barang[idx].harga {
							idx = k
						}
					}
					//Geser Nama
					tempstr=barang[i].nama
					barang[i].nama=barang[idx].nama
					barang[idx].nama=tempstr
					//Geser Harga
					temp=barang[i].harga
					barang[i].harga=barang[idx].harga
					barang[idx].harga=temp
					//Geser Stok
					temp=barang[i].stok
					barang[i].stok=barang[idx].stok
					barang[idx].stok=temp
				}
				tampilBarang(barang)

			}else if (pilih==3) {
				for i:=0;i < markbarang;i++ {
					idx = i
					for k:=i;k < markbarang;k++ {
						if barang[k].stok < barang[idx].stok {
							idx = k
						}
					}
					//Geser Nama
					tempstr=barang[i].nama
					barang[i].nama=barang[idx].nama
					barang[idx].nama=tempstr
					//Geser Harga
					temp=barang[i].harga
					barang[i].harga=barang[idx].harga
					barang[idx].harga=temp
					//Geser Stok
					temp=barang[i].stok
					barang[i].stok=barang[idx].stok
					barang[idx].stok=temp
				}
				tampilBarang(barang)
			}	
		//Descending (Insertion Sort)
		}else if urutstat==2 {
			if (pilih==1) {
				for i:=1;i < markbarang;i++ {
					idx = i
					for idx > 0 {
						if barang[idx-1].nama < barang[idx].nama {
							//Geser Nama
							tempstr=barang[idx-1].nama
							barang[idx-1].nama=barang[idx].nama
							barang[idx].nama=tempstr
							//Geser Harga
							temp=barang[idx-1].harga
							barang[idx-1].harga=barang[idx].harga
							barang[idx].harga=temp
							//Geser Stok
							temp=barang[idx-1].stok
							barang[idx-1].stok=barang[idx].stok
							barang[idx].stok=temp
						}
						idx=idx-1
					}
				}
				tampilBarang(barang)

			}else if (pilih==2) {
				for i:=1;i < markbarang;i++ {
					idx = i
					for idx > 0 {
						if barang[idx-1].harga < barang[idx].harga {
							//Geser Nama
							tempstr=barang[idx-1].nama
							barang[idx-1].nama=barang[idx].nama
							barang[idx].nama=tempstr
							//Geser Harga
							temp=barang[idx-1].harga
							barang[idx-1].harga=barang[idx].harga
							barang[idx].harga=temp
							//Geser Stok
							temp=barang[idx-1].stok
							barang[idx-1].stok=barang[idx].stok
							barang[idx].stok=temp
						}
						idx=idx-1
					}
				}
				tampilBarang(barang)

			}else if (pilih==3) {
				for i:=1;i < markbarang;i++ {
					idx = i
					for idx > 0 {
						if barang[idx-1].stok < barang[idx].stok {
							//Geser Nama
							tempstr=barang[idx-1].nama
							barang[idx-1].nama=barang[idx].nama
							barang[idx].nama=tempstr
							//Geser Harga
							temp=barang[idx-1].harga
							barang[idx-1].harga=barang[idx].harga
							barang[idx].harga=temp
							//Geser Stok
							temp=barang[idx-1].stok
							barang[idx-1].stok=barang[idx].stok
							barang[idx].stok=temp
						}
						idx=idx-1
					}
				}
				tampilBarang(barang)
			}	
		}
	}else if (datastat==2) {
		menucarisort("|               1. Nama Produsen            |","|               2. Alamat Produsen          |","|               3. Telepon Produsen         |","Urut","diurutkan")
		clear()
		menuurut(&urutstat)
		//Ascending (Selection Sort)
		if urutstat==1 {
			if (pilih==1) {
				for i:=0;i < markprodusen;i++ {
					idx = i
					for k:=i;k < markprodusen;k++ {
						if produsen[k].nama < produsen[idx].nama {
							idx = k
						}
					}
					//Geser Nama
					tempstr=produsen[i].nama
					produsen[i].nama=produsen[idx].nama
					produsen[idx].nama=tempstr
					//Geser Alamat
					tempstr=produsen[i].alamat
					produsen[i].alamat=produsen[idx].alamat
					produsen[idx].alamat=tempstr
					//Geser Telp
					tempstr=produsen[i].telp
					produsen[i].telp=produsen[idx].telp
					produsen[idx].telp=tempstr
				}
				tampilProdusen(produsen)

			}else if (pilih==2) {
				for i:=0;i < markprodusen;i++ {
					idx = i
					for k:=i;k < markprodusen;k++ {
						if produsen[k].alamat < produsen[idx].alamat {
							idx = k
						}
					}
					//Geser Nama
					tempstr=produsen[i].nama
					produsen[i].nama=produsen[idx].nama
					produsen[idx].nama=tempstr
					//Geser Alamat
					tempstr=produsen[i].alamat
					produsen[i].alamat=produsen[idx].alamat
					produsen[idx].alamat=tempstr
					//Geser Telp
					tempstr=produsen[i].telp
					produsen[i].telp=produsen[idx].telp
					produsen[idx].telp=tempstr
				}
				tampilProdusen(produsen)

			}else if (pilih==3) {
				for i:=0;i < markprodusen;i++ {
					idx = i
					for k:=i;k < markprodusen;k++ {
						if produsen[k].telp < produsen[idx].telp {
							idx = k
						}
					}
					//Geser Nama
					tempstr=produsen[i].nama
					produsen[i].nama=produsen[idx].nama
					produsen[idx].nama=tempstr
					//Geser Alamat
					tempstr=produsen[i].alamat
					produsen[i].alamat=produsen[idx].alamat
					produsen[idx].alamat=tempstr
					//Geser Telp
					tempstr=produsen[i].telp
					produsen[i].telp=produsen[idx].telp
					produsen[idx].telp=tempstr
				}
				tampilProdusen(produsen)
			}	
		//Descending (Insertion Sort)
		}else if urutstat==2 {
			if (pilih==1) {
				for i:=1;i < markprodusen;i++ {
					idx = i
					for idx > 0 {
						if produsen[idx-1].nama < produsen[idx].nama {
							//Geser Nama
							tempstr=produsen[idx-1].nama
							produsen[idx-1].nama=produsen[idx].nama
							produsen[idx].nama=tempstr
							//Geser Alamat
							tempstr=produsen[idx-1].alamat
							produsen[idx-1].alamat=produsen[idx].alamat
							produsen[idx].alamat=tempstr
							//Geser Telp
							tempstr=produsen[idx-1].telp
							produsen[idx-1].telp=produsen[idx].telp
							produsen[idx].telp=tempstr
						}
						idx=idx-1
					}
				}
				tampilProdusen(produsen)

			}else if (pilih==2) {
				for i:=1;i < markprodusen;i++ {
					idx = i
					for idx > 0 {
						if produsen[idx-1].alamat < produsen[idx].alamat {
							//Geser Nama
							tempstr=produsen[idx-1].nama
							produsen[idx-1].nama=produsen[idx].nama
							produsen[idx].nama=tempstr
							//Geser Alamat
							tempstr=produsen[idx-1].alamat
							produsen[idx-1].alamat=produsen[idx].alamat
							produsen[idx].alamat=tempstr
							//Geser Telp
							tempstr=produsen[idx-1].telp
							produsen[idx-1].telp=produsen[idx].telp
							produsen[idx].telp=tempstr
						}
						idx=idx-1
					}
				}
				tampilProdusen(produsen)

			}else if (pilih==3) {
				for i:=1;i < markprodusen;i++ {
					idx = i
					for idx > 0 {
						if produsen[idx-1].telp < produsen[idx].telp {
							//Geser Nama
							tempstr=produsen[idx-1].nama
							produsen[idx-1].nama=produsen[idx].nama
							produsen[idx].nama=tempstr
							//Geser Alamat
							tempstr=produsen[idx-1].alamat
							produsen[idx-1].alamat=produsen[idx].alamat
							produsen[idx].alamat=tempstr
							//Geser Telp
							tempstr=produsen[idx-1].telp
							produsen[idx-1].telp=produsen[idx].telp
							produsen[idx].telp=tempstr
						}
						idx=idx-1
					}
				}
				tampilProdusen(produsen)
			}	
		}
	}else if (datastat==3) {
		menucarisort("|               1. Nama Agen                |","|               2. Alamat Agen              |","|               3. Telepon Agen             |","Urut","diurutkan")
		clear()
		menuurut(&urutstat)
		//Ascending (Selection Sort)
		if urutstat==1 {
			if (pilih==1) {
				for i:=0;i < marktempat;i++ {
					idx = i
					for k:=i;k < marktempat;k++ {
						if tempat[k].nama < tempat[idx].nama {
							idx = k
						}
					}
					//Geser Nama
					tempstr=tempat[i].nama
					tempat[i].nama=tempat[idx].nama
					tempat[idx].nama=tempstr
					//Geser Alamat
					tempstr=tempat[i].alamat
					tempat[i].alamat=tempat[idx].alamat
					tempat[idx].alamat=tempstr
					//Geser Telp
					tempstr=tempat[i].telp
					tempat[i].telp=tempat[idx].telp
					tempat[idx].telp=tempstr
				}
				tampilTempat(tempat)

			}else if (pilih==2) {
				for i:=0;i < marktempat;i++ {
					idx = i
					for k:=i;k < marktempat;k++ {
						if tempat[k].alamat < tempat[idx].alamat {
							idx = k
						}
					}
					//Geser Nama
					tempstr=tempat[i].nama
					tempat[i].nama=tempat[idx].nama
					tempat[idx].nama=tempstr
					//Geser Alamat
					tempstr=tempat[i].alamat
					tempat[i].alamat=tempat[idx].alamat
					tempat[idx].alamat=tempstr
					//Geser Telp
					tempstr=tempat[i].telp
					tempat[i].telp=tempat[idx].telp
					tempat[idx].telp=tempstr
				}
				tampilTempat(tempat)

			}else if (pilih==3) {
				for i:=0;i < marktempat;i++ {
					idx = i
					for k:=i;k < marktempat;k++ {
						if tempat[k].telp < tempat[idx].telp {
							idx = k
						}
					}
					//Geser Nama
					tempstr=tempat[i].nama
					tempat[i].nama=tempat[idx].nama
					tempat[idx].nama=tempstr
					//Geser Alamat
					tempstr=tempat[i].alamat
					tempat[i].alamat=tempat[idx].alamat
					tempat[idx].alamat=tempstr
					//Geser Telp
					tempstr=tempat[i].telp
					tempat[i].telp=tempat[idx].telp
					tempat[idx].telp=tempstr
				}
				tampilTempat(tempat)
			}	
		//Descending (Insertion Sort)
		}else if urutstat==2 {
			if (pilih==1) {
				for i:=1;i < marktempat;i++ {
					idx = i
					for idx > 0 {
						if tempat[idx-1].nama < tempat[idx].nama {
							//Geser Nama
							tempstr=tempat[idx-1].nama
							tempat[idx-1].nama=tempat[idx].nama
							tempat[idx].nama=tempstr
							//Geser Alamat
							tempstr=tempat[idx-1].alamat
							tempat[idx-1].alamat=tempat[idx].alamat
							tempat[idx].alamat=tempstr
							//Geser Telp
							tempstr=tempat[idx-1].telp
							tempat[idx-1].telp=tempat[idx].telp
							tempat[idx].telp=tempstr
						}
						idx=idx-1
					}
				}
				tampilTempat(tempat)

			}else if (pilih==2) {
				for i:=1;i < marktempat;i++ {
					idx = i
					for idx > 0 {
						if tempat[idx-1].alamat < tempat[idx].alamat {
							//Geser Nama
							tempstr=tempat[idx-1].nama
							tempat[idx-1].nama=tempat[idx].nama
							tempat[idx].nama=tempstr
							//Geser Alamat
							tempstr=tempat[idx-1].alamat
							tempat[idx-1].alamat=tempat[idx].alamat
							tempat[idx].alamat=tempstr
							//Geser Telp
							tempstr=tempat[idx-1].telp
							tempat[idx-1].telp=tempat[idx].telp
							tempat[idx].telp=tempstr
						}
						idx=idx-1
					}
				}
				tampilTempat(tempat)

			}else if (pilih==3) {
				for i:=1;i < marktempat;i++ {
					idx = i
					for idx > 0 {
						if tempat[idx-1].telp < tempat[idx].telp {
							//Geser Nama
							tempstr=tempat[idx-1].nama
							tempat[idx-1].nama=tempat[idx].nama
							tempat[idx].nama=tempstr
							//Geser Alamat
							tempstr=tempat[idx-1].alamat
							tempat[idx-1].alamat=tempat[idx].alamat
							tempat[idx].alamat=tempstr
							//Geser Telp
							tempstr=tempat[idx-1].telp
							tempat[idx-1].telp=tempat[idx].telp
							tempat[idx].telp=tempstr
						}
						idx=idx-1
					}
				}
				tampilTempat(tempat)
			}	
		}
	}
}
