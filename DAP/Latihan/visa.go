package main
import "fmt"
func main() {
  var tanggal,bulan,tahun int
  var hari string
  fmt.Print("Tanggal,Bulan,Tahun,Hari: ")
  fmt.Scan(&tanggal,&bulan,&tahun,&hari)
  if hari == "Senin"{
    if tanggal == 31 {
      tanggal = 0
      tanggal += 2
      bulan++
      hari = "Rabu"
    }else{
    tanggal += 2
    hari = "Rabu"}
  }else if hari == "Selasa"{
    if tanggal == 31 {
      tanggal = 0
      tanggal += 2
      bulan++
      hari = "Kamis"
    }else{
    tanggal += 2
    hari = "Kamis"}
  } else if hari == "Rabu"{
      if tanggal == 31 {
      tanggal = 0
      tanggal += 2
      bulan++
      hari = "Jumat"
    }else{
    tanggal += 2
    hari = "Jumat"}
  } else if hari == "Kamis"{
     if tanggal == 31 {
      tanggal = 0
      tanggal += 4
      bulan++
      hari = "Senin"
    }else{
    tanggal += 4
    hari = "Senin"}
  } else if hari == "Jumat"{
      if tanggal == 31 {
        tanggal = 0
        tanggal += 4
        bulan++
        hari = "Selasa"
      }else{
        tanggal += 4
        hari = "Selasa"}
  }
  fmt.Println("Tanggal ",tanggal,"Bulan ",bulan,"Tahun",tahun,"Hari",hari)
}
