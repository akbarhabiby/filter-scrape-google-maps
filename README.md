# Filter Scrape Google Maps

## Input example
**Filetype: json**

place the .json inside **input** folder
```json
{
	"scrapes": [
		{
			"id" : 9273441,
			"full_address" : "Jl. Jambi No.10-37, Simpang III Sipin, Kec. Kota Baru, Kota Jambi, Jambi 36124, Indonesia",
			"number" : "10-37",
			"district" : "Kecamatan Kota Baru",
			"city" : "Kota Jambi",
			"province" : "Jambi",
			"postal_code" : "36124",
			"country" : "Indonesia",
			"latitude" : "-1.6213922",
			"longitude" : "103.5991754",
			"plus_code" : null,
			"created_at" : "2022-07-28T03:32:11.521Z"
		},
		{
			"id" : 9163815,
			"full_address" : "Jl. Adi Sucipto No.87, Ampenan Utara, Kec. Ampenan, Kota Mataram, Nusa Tenggara Bar. 83511, Indonesia",
			"number" : "87",
			"district" : "Kecamatan Ampenan",
			"city" : "Kota Mataram",
			"province" : "Nusa Tenggara Barat",
			"postal_code" : "83511",
			"country" : "Indonesia",
			"latitude" : "-8.5654823",
			"longitude" : "116.0825766",
			"plus_code" : "6P3RC3MM+R2",
			"created_at" : "2022-07-27T22:33:50.180Z"
		}
	]
}
```

## Log example
if there is an address that cannot be parsed
```log
2022/08/27 17:56:43 failed to parse [C-33]
2022/08/27 17:56:43 failed to parse [f-15]
2022/08/27 17:56:43 failed to parse [F-14]
2022/08/27 17:56:43 failed to parse [e-09]
2022/08/27 17:56:43 failed to parse [2 C-D]
2022/08/27 17:56:43 failed to parse [400-B]
```