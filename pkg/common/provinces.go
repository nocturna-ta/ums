package common

type Province struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var provinces = []Province{
	{Code: "11", Name: "Aceh"},
	{Code: "51", Name: "Bali"},
	{Code: "36", Name: "Banten"},
	{Code: "17", Name: "Bengkulu"},
	{Code: "34", Name: "Daerah Istimewa Yogyakarta"},
	{Code: "31", Name: "DKI Jakarta"},
	{Code: "75", Name: "Gorontalo"},
	{Code: "15", Name: "Jambi"},
	{Code: "32", Name: "Jawa Barat"},
	{Code: "33", Name: "Jawa Tengah"},
	{Code: "35", Name: "Jawa Timur"},
	{Code: "61", Name: "Kalimantan Barat"},
	{Code: "63", Name: "Kalimantan Selatan"},
	{Code: "62", Name: "Kalimantan Tengah"},
	{Code: "64", Name: "Kalimantan Timur"},
	{Code: "65", Name: "Kalimantan Utara"},
	{Code: "19", Name: "Kepulauan Bangka Belitung"},
	{Code: "21", Name: "Kepulauan Riau"},
	{Code: "18", Name: "Lampung"},
	{Code: "81", Name: "Maluku"},
	{Code: "82", Name: "Maluku Utara"},
	{Code: "52", Name: "Nusa Tenggara Barat"},
	{Code: "53", Name: "Nusa Tenggara Timur"},
	{Code: "91", Name: "Papua"},
	{Code: "92", Name: "Papua Barat"},
	{Code: "95", Name: "Papua Pegunungan"},
	{Code: "93", Name: "Papua Selatan"},
	{Code: "94", Name: "Papua Tengah"},
	{Code: "14", Name: "Riau"},
	{Code: "76", Name: "Sulawesi Barat"},
	{Code: "73", Name: "Sulawesi Selatan"},
	{Code: "72", Name: "Sulawesi Tengah"},
	{Code: "74", Name: "Sulawesi Tenggara"},
	{Code: "71", Name: "Sulawesi Utara"},
	{Code: "13", Name: "Sumatera Barat"},
	{Code: "16", Name: "Sumatera Selatan"},
	{Code: "12", Name: "Sumatera Utara"},
}
