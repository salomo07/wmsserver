package models

//Buat validator/verifikator object
// Type = "Buy/Sell"
type InvoiceModel struct {
	Id          string               `json:"_id"`
	Rev         string               `json:"_rev"`
	InvoiceNo   string               `json:"invoiceNo"`
	Type        string               `json:"type"`
	Billto      string               `json:"billto"`
	Shipto      string               `json:"shipto"`
	InvoiceDate string               `json:"invoicedate"`
	PoNo        string               `json:"pono"`
	DueDate     string               `json:"duedate"`
	Details     []InvoiceDetailModel `json:"details"`
	Loadby      []string             `json:"loadby"`
	Tax         float32              `json:"tax"`
	Shipping    float32              `json:"shipping"`
	Other       float32              `json:"other"`
	Total       float32              `json:"total"`
	Isactive    bool                 `json:"isactive" binding:"required"`
	Table       string               `json:"table" binding:"required"`
}
type InvoiceDetailModel struct {
	IdProduct string  `json:"idproduct"`
	Qty       int32   `json:"qty"`
	Unitprice float32 `json:"unitprice"`
	Checkby   string  `json:"checkby"`
	Approveby string  `json:"approveby"`
}
type ThirdPartyModel struct {
	Id          string  `json:"_id"`
	Rev         string  `json:"_rev"`
	Name        string  `json:"idproduct"`
	Category    int32   `json:"qty"`
	Phone       float32 `json:"phone"`
	BankAcc     string  `json:"bankAcc"`
	Address     string  `json:"address"`
	Shipaddress string  `json:"shipaddress"`
	IsActive    bool    `json:"isActive"`
	Table       string  `json:"table"`
}
