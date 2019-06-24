package webstoreservice

import (
        // "encoding/json"
    // "nfg/web/app/logger"
    // "log"
    "io"
    "os"
    // "os/exec"
    "path/filepath"
    "github.com/kataras/iris"

    // "time"
    "fmt"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "strconv"
    "encoding/csv"
    "os/exec"
)

type CatatanJumlahBarang struct {
	ID  int `json:"ID"`
	SKU string `json:"SKU"`
    NamaItem string `json:"NamaItem"`
    Jumlah int `json:"Jumlah"`
}

type CatatanBarangMasuk struct {
	Waktu string `json:"Waktu"`
	SKU string `json:"SKU"`
    NamaBarang string `json:"NamaBarang"`
    JumlahPesanan int `json:"JumlahPesanan"`
    JumlahDiterima int `json:"JumlahDiterima"`
    HargaBeli int `json:"HargaBeli"`
    Total int `json:"Total"`
    NomorKwitansi string `json:"NomorKwitansi"`
    Catatan string `json:"Catatan"`
}

type CatatanBarangKeluar struct {
	Waktu string `json:"Waktu"`
	SKU string `json:"SKU"`
    NamaBarang string `json:"NamaBarang"`
    JumlahKeluar int `json:"JumlahKeluar"`
    HargaJual int `json:"HargaJual"`
    Total int `json:"Total"`
    Catatan string `json:"Catatan"`
}

type LaporanPenjualan struct {
	IDPesanan  string `json:"IDPesanan"`
	Waktu string `json:"Waktu"`
	SKU string `json:"SKU"`
    NamaBarang string `json:"NamaBarang"`
    Jumlah int `json:"Jumlah"`
    HargaJual int `json:"HargaJual"`
    Total int `json:"Total"`
    HargaBeli int `json:"HargaBeli"`
    Laba int `json:"Laba"`
}

type Struk struct {
    // IDStruk string `json:"IDStruk"`
    // Waktu Time.time `json:"Waktu"` 
    SKU string `json:"SKU"`
    NamaBarang string `json:"NamaBarang"`
    Jumlah int `json:"Jumlah"`
    HargaSatuan int `json:"HargaSatuan"`
    // HargaTotal int `json:"HargaTotal"`
}
type Kwitansi struct {
    SKU string `json:"SKU"`
    NamaBarang string `json:"NamaBarang"`
    JumlahPesanan int `json:"JumlahPesanan"`
    JumlahDiterima int `json:"JumlahDiterima"`
    HargaBeli int `json:"HargaBeli"`
}


func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "WebstoreDatabase.db")
    checkErr(err)
	return db
}

func checkErr(err error) {
        if err != nil {
            panic(err)
        }
}

func InsertCatatanJumlahBarang(JumlahBarang CatatanJumlahBarang) error{
	SKU := JumlahBarang.SKU
	NamaItem := JumlahBarang.NamaItem
	Jumlah := JumlahBarang.Jumlah
    db := ConnectDB()

	stmt, err := db.Prepare("INSERT INTO CatatanJumlahBarang( SKU, NamaItem, Jumlah) values(?,?,?)")
    checkErr(err)

    res, err := stmt.Exec(SKU,NamaItem,Jumlah)
    checkErr(err)

    _, err = res.LastInsertId()
    checkErr(err)

    if err != nil {
    	return err
    }
    return nil
}

func UpdateCatatanBarang(Jumlah int, SKU string) error {
    db := ConnectDB()
    stmt, err := db.Prepare("update CatatanJumlahBarang set Jumlah=? where SKU=?")
    checkErr(err)

    res, err := stmt.Exec(Jumlah, SKU)
    checkErr(err)

    _, err = res.RowsAffected()
    checkErr(err)
    if err != nil {
        return err
    }
    return nil
}

func InsertCatatanBarangMasuk(BarangMasuk CatatanBarangMasuk) error{
	Waktu := BarangMasuk.Waktu
	SKU := BarangMasuk.SKU
	NamaBarang := BarangMasuk.NamaBarang
	JumlahPesanan := BarangMasuk.JumlahPesanan
	JumlahDiterima := BarangMasuk.JumlahDiterima
	HargaBeli := BarangMasuk.HargaBeli
	Total := BarangMasuk.Total
	NomorKwitansi := BarangMasuk.NomorKwitansi
	Catatan := BarangMasuk.Catatan
    db := ConnectDB()

	stmt, err := db.Prepare("INSERT INTO CatatanBarangMasuk( Waktu, SKU, NamaBarang, JumlahPesanan, JumlahDiterima, HargaBeli, Total, NomorKwitansi, Catatan) values(?,?,?,?,?,?,?,?,?)")
    checkErr(err)

    res, err := stmt.Exec(Waktu, SKU, NamaBarang, JumlahPesanan, JumlahDiterima, HargaBeli, Total, NomorKwitansi, Catatan)
    checkErr(err)

    _, err = res.LastInsertId()
    checkErr(err)

    if err != nil {
    	return err
    }
    return nil
}



func InsertCatatanBarangKeluar(BarangKeluar CatatanBarangKeluar) error{
	Waktu := BarangKeluar.Waktu
	SKU := BarangKeluar.SKU
	NamaBarang := BarangKeluar.NamaBarang
    JumlahKeluar := BarangKeluar.JumlahKeluar
	HargaJual := BarangKeluar.HargaJual
	Total := BarangKeluar.Total
	Catatan := BarangKeluar.Catatan
    db := ConnectDB()

	stmt, err := db.Prepare("INSERT INTO CatatanBarangKeluar( Waktu, SKU, NamaBarang, JumlahKeluar, HargaJual, Total, Catatan) values(?,?,?,?,?,?,?)")
    checkErr(err)

    res, err := stmt.Exec(Waktu, SKU, NamaBarang, JumlahKeluar, HargaJual, Total, Catatan)
    checkErr(err)

    _, err = res.LastInsertId()
    checkErr(err)

    if err != nil {
    	return err
    }
    return nil
}

func InsertLaporanPenjualan(Laporan LaporanPenjualan) error{
    IDPesanan := Laporan.IDPesanan
    Waktu := Laporan.Waktu
    SKU := Laporan.SKU
    NamaBarang := Laporan.NamaBarang
    Jumlah := Laporan.Jumlah
    HargaJual := Laporan.HargaJual
    Total := Laporan.Total
    HargaBeli := Laporan.HargaBeli
    Laba := Laporan.Laba

    db := ConnectDB()
    stmt, err := db.Prepare("INSERT INTO LaporanPenjualan( IDPesanan, Waktu, SKU, NamaBarang, Jumlah, HargaJual, Total, HargaBeli, Laba) values(?,?,?,?,?,?,?,?,?)")
    checkErr(err)

    res, err := stmt.Exec( IDPesanan, Waktu, SKU, NamaBarang, Jumlah, HargaJual, Total, HargaBeli, Laba)
    checkErr(err)

    _, err = res.LastInsertId()
    checkErr(err)

    if err != nil {
        return err
    }
    return nil
}

func NilaiBarangTotal(SKU string) int {
    var HargaBeliSemua int
    var Total int
    db := ConnectDB()
    SKU = "'"+SKU+"'"
    query := fmt.Sprintf("%s %s" , "SELECT Total FROM CatatanBarangMasuk WHERE SKU = ", SKU)
    rows, err := db.Query(query)
    // fmt.Println("rows:", rows, "err:", err)
    for rows.Next() {
        err = rows.Scan(&Total)
        checkErr(err)
        HargaBeliSemua += Total
        // fmt.Println("HargaBeli:", HargaBeli)
        // fmt.Println("Total:", Total)
    }
    return HargaBeliSemua
}

func CekStokBarang(SKU string) int {
	var JumlahStok int
	db := ConnectDB()
	SKU = "'"+SKU+"'"
	query := fmt.Sprintf("%s %s" , "SELECT Jumlah FROM CatatanJumlahBarang WHERE SKU=", SKU)
	rows, err := db.Query(query)
    checkErr(err)
	for rows.Next() {
            err = rows.Scan(&JumlahStok)
            checkErr(err)
        }
	return JumlahStok
}

func ExportCSV_Catatan_Barang_Masuk(fname string) (string, error) {
    // fname := "Catatan_Barang_Masuk.csv"
    Remove(fname)
    fmt.Println("filename:", fname)
    f, err := os.Create(fname)
    if err != nil {
        return "Fail to Export CSV : Cannot Create File ", err
    }
    defer f.Close()
    w := csv.NewWriter(f)
    defer w.Flush()
    headers := []string{
        "Waktu",
        "SKU",
        "NamaBarang",
        "JumlahPesanan",
        "JumlahDiterima",
        "HargaBeli",
        "Total",
        "NomorKwitansi",
        "Catatan",
    }
    w.Write(headers)

    var Row CatatanBarangMasuk
    db := ConnectDB()
    query := ("SELECT * FROM CatatanBarangMasuk")
    rows, err := db.Query(query)
    
    for rows.Next() {
            err = rows.Scan(&Row.Waktu, &Row.SKU, &Row.NamaBarang, &Row.JumlahPesanan, &Row.JumlahDiterima, &Row.HargaBeli, &Row.Total, &Row.NomorKwitansi, &Row.Catatan)
            if err != nil {
                return "Fail to Export CSV : Fail to Read from Databases ", err
            }
            // fmt.Println("Row:", Row)

            r := make([]string, 0, 9)
            r = append(r, Row.Waktu)
            r = append(r, Row.SKU)
            r = append(r, Row.NamaBarang)
            r = append(r, strconv.Itoa(Row.JumlahPesanan))
            r = append(r, strconv.Itoa(Row.JumlahDiterima))
            r = append(r, strconv.Itoa(Row.HargaBeli))
            r = append(r, strconv.Itoa(Row.Total))
            r = append(r, Row.NomorKwitansi)
            r = append(r, Row.Catatan)
            err := w.Write(r)
            if err != nil {
                return "Fail to Export CSV : Cannot Write ", err
            }
        }   
    // CatatanJson, _ := json.Marshal(Row)
    // fmt.Println(string(CatatanJson))
    res := ("Data Successfully exported to " + fname)
    return res, nil
}

func ExportCSV_Catatan_Barang_Keluar(fname string) (string, error) {
    // fname := "Catatan_Barang_Masuk.csv"
    Remove(fname)
    f, err := os.Create(fname)
    if err != nil {
        return "Fail to Export CSV : Cannot Create File ", err
    }
    defer f.Close()
    w := csv.NewWriter(f)
    defer w.Flush()
    headers := []string{
        "Waktu",
        "SKU",
        "NamaBarang",
        "JumlahKeluar",
        "HargaJual",
        "Total",
        "Catatan",
    }
    w.Write(headers)

    var Row CatatanBarangKeluar
    db := ConnectDB()
    query := ("SELECT * FROM CatatanBarangKeluar")
    rows, err := db.Query(query)
    for rows.Next() {
            err = rows.Scan(&Row.Waktu, &Row.SKU, &Row.NamaBarang, &Row.JumlahKeluar, &Row.HargaJual, &Row.Total, &Row.Catatan)
            // fmt.Println("Row:", Row)

            r := make([]string, 0, 9)
            r = append(r, Row.Waktu)
            r = append(r, Row.SKU)
            r = append(r, Row.NamaBarang)
            r = append(r, strconv.Itoa(Row.JumlahKeluar))
            r = append(r, strconv.Itoa(Row.HargaJual))
            r = append(r, strconv.Itoa(Row.Total))
            r = append(r, Row.Catatan)
            err := w.Write(r)
            if err != nil {
                return "Fail to Export CSV : Cannot Write ", err
            }
        }   
    res := ("Data Successfully exported to " + fname)
    return res, nil
}


func ExportCSV_Laporan_Penjualan(fname string) (string, error) {
    // fname := "Catatan_Barang_Masuk.csv"
    Remove(fname)
    f, err := os.Create(fname)
    if err != nil {
        fmt.Println("go here")
        return "Fail to Export CSV : Cannot Create File ", err
    }
    defer f.Close()
    w := csv.NewWriter(f)
    defer w.Flush()
    headers := []string{
        "IDPesanan",
        "Waktu",
        "SKU",
        "NamaBarang",
        "Jumlah",
        "HargaJual",
        "Total",
        "HargaBeli",
        "Laba",
    }
    w.Write(headers)

    var Row LaporanPenjualan
    db := ConnectDB()
    query := ("SELECT * FROM LaporanPenjualan")
    rows, err := db.Query(query)
    for rows.Next() {
            err = rows.Scan(&Row.IDPesanan, &Row.Waktu, &Row.SKU, &Row.NamaBarang, &Row.Jumlah, &Row.HargaJual, &Row.Total, &Row.HargaBeli, &Row.Laba)
            // fmt.Println("Row:", Row)

            r := make([]string, 0, 9)
            r = append(r, Row.IDPesanan)
            r = append(r, Row.Waktu)
            r = append(r, Row.SKU)
            r = append(r, Row.NamaBarang)
            r = append(r, strconv.Itoa(Row.Jumlah))
            r = append(r, strconv.Itoa(Row.HargaJual))
            r = append(r, strconv.Itoa(Row.Total))
            r = append(r, strconv.Itoa(Row.HargaBeli))
            r = append(r, strconv.Itoa(Row.Laba))
            err := w.Write(r)
            if err != nil {
                return "Fail to Export CSV : Cannot Write ", err
            }
        }   
    // CatatanJson, _ := json.Marshal(Row)
    // fmt.Println(string(CatatanJson))
    res := ("Data Successfully exported to " + fname)
    return res, nil
}

func JumlahDiterima(SKU string) int{
    var JumlahDiterima int
    db := ConnectDB()
    SKU = "'"+SKU+"'"
    query := fmt.Sprintf("%s %s" , "select sum(jumlahDiterima) FROM CatatanBarangMasuk WHERE SKU = ", SKU)
    rows, err := db.Query(query)
    checkErr(err)
    for rows.Next() {
            err = rows.Scan(&JumlahDiterima)
            checkErr(err)
        }
    return JumlahDiterima


}

func ServefileIris(ctx iris.Context, fname string) {

    filetoHandle := "./"+fname
    ctx.SendFile(filetoHandle, "exported.csv")
}


func HandleFile(ctx iris.Context) (bool, string, string){

    // Get the file from the request.
    file, info, err := ctx.FormFile("file")
    fname := info.Filename
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.HTML("Error while uploading [reading the file]: <b>" + err.Error() + "</b>")
        msg := "Error while uploading [reading the file]"
        return false, msg, fname
    }
    defer file.Close()

    RemoveContents("../../"+fname)

    out, err := os.OpenFile("../../"+fname, os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.HTML("Error while uploading2 [creating the file]: <b>" + err.Error() + "</b>")
        msg := "Error while uploading [creating the file]"
        return false ,msg, fname
    }
    defer out.Close()
    io.Copy(out, file)
    msg := "The file has been uploaded"
    return true, msg, fname

}

func RemoveContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}

func Remove(destination string) (){
    echoCmd := exec.Command("echo", "")
    rmCmd := exec.Command("rm", "-rf", destination)
    reader, writer := io.Pipe()

    echoCmd.Stdout = writer
    rmCmd.Stdin = reader

    echoCmd.Start()
    rmCmd.Start()

    echoCmd.Wait()
    writer.Close()

    rmCmd.Wait()
    reader.Close()
}
