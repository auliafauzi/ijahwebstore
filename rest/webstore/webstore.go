package webstore

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/context"
    "ijahwebstore/service/webstore"
    "ijahwebstore/tools"
    // "strconv"
    "ijahwebstore/logger"

    // "fmt"
    "reflect"
    "time"
    "encoding/json"

    "bufio"
    "encoding/csv"
    // "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    // "reflect"
)

func ListCatatanJumlahBarang(ctx iris.Context) {

}

func ListCatatanMasuk(ctx iris.Context) {

}

func ListCatatanKeluar(ctx iris.Context) {

}


func JualBarang(ctx iris.Context) {
    var struk []webstoreservice.Struk
    var objJson map[string]interface{}
    var stok int
    err := ctx.ReadJSON(&objJson)
    IDPesanan := ctx.Params().Get("StrukID")
    Waktu := fmt.Sprintf("%s", time.Now())

    if err != nil{
        logger.Error.Println("Error Struk JSON", err)
        ctx.JSON(context.Map{"result": false, "message": err})
        return
    }
    
    marshaled,_ := json.Marshal(objJson["struk"])
    errmar :=  json.Unmarshal(marshaled, &struk)
    if errmar != nil {
        logger.Error.Println("Marshalling failed", errmar)
        ctx.JSON(context.Map{"result": false, "message": errmar})
        return
    }

    // cek stok barang 
    for i,_ := range struk{
        stok = webstoreservice.CekStokBarang(struk[i].SKU)
        stok = stok - struk[i].Jumlah
        if stok < 1 {
            logger.Error.Println("Stok Kosong, SKU=", struk[i].SKU)
            ctx.JSON(context.Map{"result": false, "message": "SKU= "+struk[i].SKU+", Stok Kosong"})
            return
        }
    }
    ////

    for i,_ := range struk{
        // catat barang CatatanKeluar
        var CatatanKeluar webstoreservice.CatatanBarangKeluar
        CatatanKeluar.Waktu = Waktu
        CatatanKeluar.SKU = struk[i].SKU
        CatatanKeluar.NamaBarang = struk[i].NamaBarang
        CatatanKeluar.JumlahKeluar = struk[i].Jumlah
        CatatanKeluar.HargaJual = struk[i].HargaSatuan
        CatatanKeluar.Total = struk[i].Jumlah*struk[i].HargaSatuan
        CatatanKeluar.Catatan = IDPesanan
        err1 := webstoreservice.InsertCatatanBarangKeluar(CatatanKeluar)
        if err1 != nil {
            logger.Error.Println("Error Insert to Catatan Keluar", err1)
            ctx.JSON(context.Map{"result": false, "message": err1})
            return
        }
        // update jumlah di inventory
        JumlahKeluar := struk[i].Jumlah
        Stok := webstoreservice.CekStokBarang(struk[i].SKU) 
        StokAkhir := Stok - JumlahKeluar
        err2 := webstoreservice.UpdateCatatanBarang(StokAkhir,struk[i].SKU)
        if err2 != nil {
            logger.Error.Println("Error Update Catatan Jumlah Barang", err2)
            ctx.JSON(context.Map{"result": false, "message": err2})
            return
        }
        // hitung rata2 harga beli
        HargaBeliTotal := webstoreservice.NilaiBarangTotal(struk[i].SKU)
        JumlahDiterima := webstoreservice.JumlahDiterima(struk[i].SKU)
        HargaBeliRata2 := int(HargaBeliTotal/JumlahDiterima)

        // catat laporan penjualan
        var Laporan webstoreservice.LaporanPenjualan
        Laporan.IDPesanan = IDPesanan
        Laporan.Waktu = Waktu
        Laporan.SKU = struk[i].SKU
        Laporan.NamaBarang = struk[i].NamaBarang
        Laporan.Jumlah = struk[i].Jumlah
        Laporan.HargaJual = struk[i].HargaSatuan
        Laporan.Total = struk[i].Jumlah*struk[i].HargaSatuan
        Laporan.HargaBeli = HargaBeliRata2
        Laporan.Laba =  (struk[i].HargaSatuan-HargaBeliRata2)*struk[i].Jumlah
        err3 := webstoreservice.InsertLaporanPenjualan(Laporan)
        if err3 != nil {
            logger.Error.Println("Error Laporan Penjualan", err3)
            ctx.JSON(context.Map{"result": false, "message": err3})
            return
        }
    }
}

func BarangMasuk(ctx iris.Context) {
    var kwitansi []webstoreservice.Kwitansi
    var objJson map[string]interface{}
    err := ctx.ReadJSON(&objJson)
    IDPesanan := ctx.Params().Get("StrukID")
    Waktu := fmt.Sprintf("%s", time.Now())

    if err != nil{
        logger.Error.Println("Error Struct JSON", err)
        ctx.JSON(context.Map{"result": false, "message": err})
        return
    }
    
    marshaled,_ := json.Marshal(objJson["kwitansi"])
    errmar :=  json.Unmarshal(marshaled, &kwitansi)
    if errmar != nil {
        logger.Error.Println("Marshalling failed", errmar)
        ctx.JSON(context.Map{"result": false, "message": errmar})
        return
    }
    fmt.Println(kwitansi)

    for i,_ := range kwitansi{
        // catat barang CatatanMasuk
        var CatatanMasuk webstoreservice.CatatanBarangMasuk
        CatatanMasuk.Waktu = Waktu
        CatatanMasuk.SKU = kwitansi[i].SKU
        CatatanMasuk.NamaBarang = kwitansi[i].NamaBarang
        CatatanMasuk.JumlahPesanan = kwitansi[i].JumlahPesanan
        CatatanMasuk.JumlahDiterima = kwitansi[i].JumlahDiterima
        CatatanMasuk.HargaBeli = kwitansi[i].HargaBeli
        CatatanMasuk.Total = kwitansi[i].JumlahDiterima*kwitansi[i].HargaBeli
        CatatanMasuk.Catatan = IDPesanan
        err1 := webstoreservice.InsertCatatanBarangMasuk(CatatanMasuk)
        if err1 != nil {
            logger.Error.Println("Error Insert to Catatan Keluar", err1)
            ctx.JSON(context.Map{"result": false, "message": err1})
            return
        }
        // update jumlah di inventory
        Stok := webstoreservice.CekStokBarang(kwitansi[i].SKU) 
        if Stok < 1 {
            var insertinventory webstoreservice.CatatanJumlahBarang
            insertinventory.SKU = kwitansi[i].SKU
            insertinventory.NamaItem = kwitansi[i].NamaBarang
            insertinventory.Jumlah = kwitansi[i].JumlahDiterima
            err := webstoreservice.InsertCatatanJumlahBarang(insertinventory)
            if err != nil {
                logger.Error.Println("Error Update Catatan Jumlah Barang", err)
                ctx.JSON(context.Map{"result": false, "message": err})
                return
            }
        } 
        JumlahMasuk := kwitansi[i].JumlahDiterima
        Stok = webstoreservice.CekStokBarang(kwitansi[i].SKU) 
        StokAkhir := Stok + JumlahMasuk
        err2 := webstoreservice.UpdateCatatanBarang(StokAkhir,kwitansi[i].SKU)
        if err2 != nil {
            logger.Error.Println("Error Update Catatan Jumlah Barang", err2)
            ctx.JSON(context.Map{"result": false, "message": err2})
            return
        }   
    }
    ctx.JSON(context.Map{"result":true, "kwitansi": kwitansi})
}

func ImportCSV_Catatan_Jumlah_Barang(ctx iris.Context) {
    result, msg, fname := webstoreservice.HandleFile(ctx)
    if result == false {
        tools.ResponseJSON(ctx, 500, context.Map{"result": result, "errors": nil , "message": msg })
        return
    }
    csvFile, _ := os.Open(fname)
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var Catatan []webstoreservice.CatatanJumlahBarang
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        angka, _ := strconv.Atoi(line[2])
        Catatan = append(Catatan, webstoreservice.CatatanJumlahBarang{
            SKU: line[0],
            NamaItem:  line[1],
            Jumlah: angka,
        })
    }

    CatatanJson, _ := json.Marshal(Catatan)
    fmt.Println(string(CatatanJson))
    fmt.Println("\n", Catatan, reflect.TypeOf(Catatan))

    // insert to sqlite
    for i,_ := range Catatan {
        err := webstoreservice.InsertCatatanJumlahBarang(Catatan[i])
        if err != nil {
            logger.Error.Println("Error Insert to Catatan Jumlah Barang", err)
            ctx.JSON(context.Map{"result": false, "message": err})
            return
        }
    }
}

func ImportCSV_Catatan_Barang_Masuk(ctx iris.Context) {
    result, msg, fname := webstoreservice.HandleFile(ctx)
    if result == false {
        tools.ResponseJSON(ctx, 500, context.Map{"result": result, "errors": nil , "message": msg })
        return
    }
    csvFile, _ := os.Open(fname)
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var CatatanMasuk []webstoreservice.CatatanBarangMasuk
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        JumlahPesanan, _ := strconv.Atoi(line[3])
        JumlahDiterima, _ := strconv.Atoi(line[4])
        HargaBeli, _ := strconv.Atoi(line[5])
        Total, _ := strconv.Atoi(line[6])
        CatatanMasuk = append(CatatanMasuk, webstoreservice.CatatanBarangMasuk{
            Waktu: line[0],
            SKU: line[1],
            NamaBarang:  line[2],
            JumlahPesanan: JumlahPesanan,
            JumlahDiterima: JumlahDiterima,
            HargaBeli: HargaBeli,
            Total: Total,
            NomorKwitansi: line[7],
            Catatan: line[8],
        })
    }

    CatatanJson, _ := json.Marshal(CatatanMasuk)
    fmt.Println(string(CatatanJson))
    fmt.Println("\n", CatatanMasuk, reflect.TypeOf(CatatanMasuk))

    // insert to sqlite
    for i,_ := range CatatanMasuk {
        err := webstoreservice.InsertCatatanBarangMasuk(CatatanMasuk[i])
        if err != nil {
            logger.Error.Println("Error Insert to Catatan Masuk ", err)
            ctx.JSON(context.Map{"result": false, "message": err})
            return
        }
    }
}

func ImportCSV_Catatan_Barang_Keluar(ctx iris.Context) {
    result, msg, fname := webstoreservice.HandleFile(ctx)
    if result == false {
        tools.ResponseJSON(ctx, 500, context.Map{"result": result, "errors": nil , "message": msg })
        return
    }
    csvFile, _ := os.Open(fname)
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var CatatanKeluar []webstoreservice.CatatanBarangKeluar
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        JumlahKeluar, _ := strconv.Atoi(line[3])
        HargaJual, _ := strconv.Atoi(line[4])
        Total, _ := strconv.Atoi(line[5])
        CatatanKeluar = append(CatatanKeluar, webstoreservice.CatatanBarangKeluar{
            Waktu: line [0],
            SKU: line[1],
            NamaBarang:  line[2],
            JumlahKeluar: JumlahKeluar,
            HargaJual: HargaJual,
            Total: Total,
            Catatan: line[6],
        })
    }

    CatatanJson, _ := json.Marshal(CatatanKeluar)
    fmt.Println(string(CatatanJson))
    fmt.Println("\n", CatatanKeluar, reflect.TypeOf(CatatanKeluar))

    // insert to sqlite
    for i,_ := range CatatanKeluar {
        err := webstoreservice.InsertCatatanBarangKeluar(CatatanKeluar[i])
        if err != nil {
            logger.Error.Println("Error Insert to Catatan Keluar ", err)
            ctx.JSON(context.Map{"result": false, "message": err})
            return
        }
    }
}

func ImportCSV_Laporan_Penjualan(ctx iris.Context) {
    result, msg, fname := webstoreservice.HandleFile(ctx)
    if result == false {
        tools.ResponseJSON(ctx, 500, context.Map{"result": result, "errors": nil , "message": msg })
        return
    }
    csvFile, _ := os.Open(fname)
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var Laporan []webstoreservice.LaporanPenjualan
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        Jumlah, _ := strconv.Atoi(line[4])
        HargaJual, _ := strconv.Atoi(line[5])
        Total, _ := strconv.Atoi(line[6])
        HargaBeli, _ := strconv.Atoi(line[7])
        Laba, _ := strconv.Atoi(line[8])
        Laporan = append(Laporan, webstoreservice.LaporanPenjualan{
            IDPesanan: line[0],
            Waktu: line [1],
            SKU: line[2],
            NamaBarang:  line[3],
            Jumlah: Jumlah,
            HargaJual: HargaJual,
            Total: Total,
            HargaBeli: HargaBeli,
            Laba: Laba,
        })
    }

    CatatanJson, _ := json.Marshal(Laporan)
    fmt.Println(string(CatatanJson))
    fmt.Println("\n", Laporan, reflect.TypeOf(Laporan))

    // insert to sqlite
    for i,_ := range Laporan {
        err := webstoreservice.InsertLaporanPenjualan(Laporan[i])
        if err != nil {
            logger.Error.Println("Error Insert to Catatan Keluar ", err)
            ctx.JSON(context.Map{"result": false, "message": err})
            return
        }
    }
}

func Export_Catatan_Masuk(ctx iris.Context){
    var fname string
    var objJson map[string]interface{}
    err := ctx.ReadJSON(&objJson)

    if err != nil{
        logger.Error.Println("Error Struct JSON", err)
        ctx.JSON(context.Map{"result": false, "message": err})
        return
    }
    
    marshaled,_ := json.Marshal(objJson["FileName"])
    errmar :=  json.Unmarshal(marshaled, &fname)
    if errmar != nil {
        logger.Error.Println("Marshalling failed", errmar)
        ctx.JSON(context.Map{"result": false, "message": errmar})
        return
    }


    fmt.Println("FileName:", fname)
    res, err := webstoreservice.ExportCSV_Catatan_Barang_Masuk(fname)
    if err != nil {
            logger.Error.Println(res, err)
            ctx.JSON(context.Map{"result": false, "message": err})
            return
        }
    webstoreservice.ServefileIris(ctx, fname)
}

func Export_Catatan_Keluar(ctx iris.Context){
    var fname string
    var objJson map[string]interface{}
    err := ctx.ReadJSON(&objJson)

    if err != nil{
        logger.Error.Println("Error Struct JSON", err)
        ctx.JSON(context.Map{"result": false, "message": err})
        return
    }
    
    marshaled,_ := json.Marshal(objJson["FileName"])
    errmar :=  json.Unmarshal(marshaled, &fname)
    if errmar != nil {
        logger.Error.Println("Marshalling failed", errmar)
        ctx.JSON(context.Map{"result": false, "message": errmar})
        return
    }


    fmt.Println("FileName:", fname)
    res, err := webstoreservice.ExportCSV_Catatan_Barang_Keluar(fname)
    if err != nil {
            logger.Error.Println(res, err)
            ctx.JSON(context.Map{"result": false, "message": err})
            return
        }
    webstoreservice.ServefileIris(ctx, fname)
}

func Export_Laporan_Penjualan(ctx iris.Context){
    var fname string
    var objJson map[string]interface{}
    err := ctx.ReadJSON(&objJson)

    if err != nil{
        logger.Error.Println("Error Struct JSON", err)
        ctx.JSON(context.Map{"result": false, "message": err})
        return
    }
    
    marshaled,_ := json.Marshal(objJson["FileName"])
    errmar :=  json.Unmarshal(marshaled, &fname)
    if errmar != nil {
        logger.Error.Println("Marshalling failed", errmar)
        ctx.JSON(context.Map{"result": false, "message": errmar})
        return
    }


    fmt.Println("FileName:", fname)
    res, err := webstoreservice.ExportCSV_Laporan_Penjualan(fname)
    if err != nil {
            logger.Error.Println(res, err)
            ctx.JSON(context.Map{"result": false, "message": err})
            return
        }
    webstoreservice.ServefileIris(ctx, fname)
}


func TestConnection(ctx iris.Context){
    // var obj nfgservice.Rule

    ctx.JSON(context.Map{"result":true, "message": "Connection Established"})
}




func Register(app *iris.Application, endpoint string, crs ...context.Handler){
    r := app.Party(endpoint, crs...).AllowMethods(iris.MethodOptions)

    r.Get("/CatatanJumlahBarang", ListCatatanJumlahBarang)
    r.Get("/CatatanMasuk", ListCatatanMasuk)
    r.Get("/CatatanKeluar", ListCatatanKeluar)
    r.Get("/test-connection", TestConnection)

    r.Post("/export-catatan-masuk", Export_Catatan_Masuk)
    r.Post("/export-catatan-keluar", Export_Catatan_Keluar)
    r.Post("/export-laporan-penjualan", Export_Laporan_Penjualan)
    r.Post("/barang-masuk", BarangMasuk)
    r.Post("/jual-barang/{StrukID:string}", JualBarang)


    // r.Get("/", RuleList)
    // r.Delete("/{id:string}", DeleteRule)
    // r.Patch("/{id:string}", UpdateRule)
    // r.Get("/getOptions", GetOptions)
    // r.Get("/statistics", StatisticLists)
}