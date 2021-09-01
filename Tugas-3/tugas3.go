package main

import (
	"fmt"
	"strconv"
)

func main() {
  // soal 1
  var panjangPersegiPanjang string = "8"
  var lebarPersegiPanjang string = "5"
  var alasSegitiga string = "9"
  var tinggiSegitiga string = "5"
  
  var kelilingPersegiPanjang int
  var luasSegitiga int

  pPersegiPanjang, _ := strconv.Atoi(panjangPersegiPanjang)
  lPersegiPanjang, _ := strconv.Atoi(lebarPersegiPanjang)
  aSegitiga, _ := strconv.Atoi(alasSegitiga)
  tSegitiga, _ := strconv.Atoi(tinggiSegitiga)

  kelilingPersegiPanjang = 2*(pPersegiPanjang+lPersegiPanjang)
  luasSegitiga = aSegitiga*tSegitiga/2
  fmt.Printf("Keliling Persegi Panjang: %d \n", kelilingPersegiPanjang)
  fmt.Printf("Luas Segitiga: %d \n", luasSegitiga)
  
  // soal 2
  var nilaiJohn int = 80 
  var nilaiDoe int = 50
  
  if nilaiJohn >= 80 {
      fmt.Println("Nilai John = A")
  } else if nilaiJohn >= 70 && nilaiJohn < 80 {
      fmt.Println("Nilai John = B")
  } else if nilaiJohn >= 60 && nilaiJohn < 70 {
      fmt.Println("Nilai John = C")
  } else if nilaiJohn >= 50 && nilaiJohn < 60 {
      fmt.Println("Nilai John = D")
  } else {
      fmt.Println("Nilai John = E")
  }
  
  if nilaiDoe >= 80 {
      fmt.Println("Nilai Doe = A")
  } else if nilaiDoe >= 70 && nilaiDoe < 80 {
      fmt.Println("Nilai Doe = B")
  } else if nilaiDoe >= 60 && nilaiDoe < 70 {
      fmt.Println("Nilai Doe = C")
  } else if nilaiDoe >= 50 && nilaiDoe < 60 {
      fmt.Println("Nilai Doe = D")
  } else {
      fmt.Println("Nilai Doe = E")
  }
  
  // soal 3
  var tanggal int = 9
  var bulan int = 12
  var tahun int = 1998
  
  switch bulan {
    case 1:
      fmt.Printf("%d Januari %d \n", tanggal, tahun)
    case 2:
      fmt.Printf("%d Februari %d \n", tanggal, tahun)
    case 3:
      fmt.Printf("%d Maret %d \n", tanggal, tahun)
    case 4:
      fmt.Printf("%d April %d \n", tanggal, tahun)
    case 5:
      fmt.Printf("%d Mei %d \n", tanggal, tahun)
    case 6:
      fmt.Printf("%d Juni %d \n", tanggal, tahun)
    case 7:
      fmt.Printf("%d Juli %d \n", tanggal, tahun)
    case 8:
      fmt.Printf("%d Agustus %d \n", tanggal, tahun)
    case 9:
      fmt.Printf("%d September %d \n", tanggal, tahun)
    case 10:
      fmt.Printf("%d Oktober %d \n", tanggal, tahun)
    case 11:
      fmt.Printf("%d November %d \n", tanggal, tahun)
    case 12:
      fmt.Printf("%d Desember %d \n", tanggal, tahun)
    default:
     fmt.Println("Tidak ada bulan")
  }
  
  // soal 4
  if tahun >= 1994 && tahun < 1964 {
      fmt.Println("Baby boomer")
  } else if tahun >= 1965 && tahun < 1979 {
      fmt.Println("Generasi X")
  } else if tahun >= 1980 && tahun < 1994 {
      fmt.Println("Generasi Y (Millenials)")
  } else if tahun >= 1995 && tahun < 2015 {
      fmt.Println("Generasi Z")
  }
}
