# Webstore RESTful API



# Services

The Services resource represents all web services currently available via Webstore RESTful API.

| Method  | Service 	   											| Header | Request Payload 							      | Response 					 | Description                        |
| ------- | --------------------------------- | ------ | ---------------------------------- | ------------------ | ------------------------------     |
| POST	  | /webstore/barang-masuk						| *none* |            Kwitansi Body					 	| Kwitansi Details   | Insert Items to inventory          |
| POST	  | /webstore/jual-barang/<*StrukID*> | *none* |             Struk Body			    		| Struk Details      | Sell Items                         |
| POST	  | /webstore/export-catatan-masuk 		| *none* |              File Name             | csv								 | export Catatan Masuk to csv files  |
| POST	  | /webstore/export-catatan-keluar		| *none* |              File Name             | csv								 | export Catatan Keluar to csv files  |
| POST	  | /webstore/export-laporan-penjualan| *none* |              File Name             | csv								 | export Laporan Penjaualan to csv files  |


# Add Items to inventory

Add items to inventory, the input is "kwitansi" which each line is product information (SKU, NamaBarang, JumlahPesanan JumlahDiterima, HargaBeli). This service will update CatatanMasuk Table, and update "Jumlah" column in "CatatanJumlahBarang" table if the units is already on our inventory, or will add new line if the we have not the product yet.   

Definition :
```bash
http://<webstore host>:<rest api port>/barang-masuk
```
The body must contain :
```bash
"kwitansi" :
	"SKU": number of Store Keeping Unit
	"NamaBarang": Name Items
	"JumlahPesanan": Number of ordered
	"JumlahDiterima": Number of accepted
	"HargaBeli": Price per unit
```

Example :
```bash
curl -X POST -d
{
	"kwitansi":[{
        		"SKU": "SKUTEST01",
    			"NamaBarang": "NamaTEST01",
    			"JumlahPesanan": 10,
    			"JumlahDiterima": 10,
    			"HargaBeli": 55000
   			},
            {
        		"SKU": "SKUTEST02",
    			"NamaBarang": "NamaTEST02",
    			"JumlahPesanan": 10,
    			"JumlahDiterima": 10,
    			"HargaBeli": 60000
   			}      
    	]
} http://localhost:8888/webstore/barang-masuk
```

# Sell items

Sell Items Service. the input is "strtuk" which each line caontain product information (SKU, NamaBarang, Jumlah, HargaSatuan). This service will check wheter the product is on our inventory or not, if not, the service will be fail, if it is exist, update "CatatanKeluar" table, calculate our profit, update "LaporanPenjualan" table, and update "Jumlah" column on "CatatanJumlahBarang" table.  

Definition :
```bash
http://<webstore host>:<rest api port>/jual-barang/<*StrukID*>  
```
The body must contain :
```bash
"struk":
	"SKU": number of Store Keeping Unit
	"NamaBarang": Name Items
	"Jumlah": Number of sold item
	"HargaSatuan": Price per unit
```

Example :
```bash
curl -X POST -d
{
	"struk":[{
      			"SKU": "SKUTEST05",
      			"NamaBarang": "NamaTEST05",
      			"Jumlah": 1 ,
      			"HargaSatuan": 110000
   				},
                {
      			"SKU": "SKUTEST06",
      			"NamaBarang": "NamaTEST06",
      			"Jumlah": 1 ,
      			"HargaSatuan": 60000
   				}  
    	]
} http://localhost:8888/webstore/jual-barang/struk-011-20190623
```
# export catatan masuk

Definition :
```bash
http://<webstore host>:<rest api port>/webstore/export-catatan-masuk
```
The body must contain :
```bash
"FileName": Expected output filename
```

Example :
```bash
curl -X POST -d
{
	"FileName": "Catatan-masuk.csv"
} http://localhost:8888/webstore/export-catatan-masuk/
```

The exported CSV will be on same folder as the application location



# export catatan keluar

Definition :
```bash
http://<webstore host>:<rest api port>/webstore/export-catatan-keluar
```
The body must contain :
```bash
"FileName": Expected output filename
```

Example :
```bash
curl -X POST -d
{
	"FileName": "Catatan-keluar.csv"

} http://localhost:8888/webstore/export-catatan-keluar/
```



# export Laporan Penjualan

Definition :
```bash
http://<webstore host>:<rest api port>/webstore/export-laporan-penjualan
```
The body must contain :
```bash
"FileName": Expected output filename
```

Example :
```bash
curl -X POST -d
{
	"FileName": "Laporan.csv"

} http://localhost:8888/webstore/export-laporan-penjualan/
```
