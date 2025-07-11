package main

type Negara struct {
	IDNegara   int    `json:"id_negara"`
	KodeNegara string `json:"kode_negara"`
	NamaNegara string `json:"nama_negara"`
}

type Pelabuhan struct {
	IDPelabuhan   string `json:"id_pelabuhan"`
	NamaPelabuhan string `json:"nama_pelabuhan"`
	IDNegara      string `json:"id_negara"`
}

type Barang struct {
	IDBarang    int     `json:"id_barang"`
	NamaBarang  string  `json:"nama_barang"`
	IDPelabuhan int     `json:"id_pelabuhan"`
	Description string  `json:"description"`
	Diskon      float64 `json:"diskon"`
	Harga       float64 `json:"harga"`
}

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
