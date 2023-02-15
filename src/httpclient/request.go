package httpclient

type CreateOrderRequest struct {
	Consignee   ConsigneRequest    `json:"consignee"`
	Consigner   ConsigneRequest    `json:"consigner"`
	Courier     CourierRequest     `json:"courier"`
	Coverage    string             `json:"coverage"`
	Destination DestinationRequest `json:"destination"`
	ExternalID  string             `json:"external_id"`
	Origin      DestinationRequest `json:"origin"`
	Package     PackageRequest     `json:"package"`
	PaymentType string             `json:"payment_type"`
}

type ConsigneRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type CourierRequest struct {
	Cod          bool  `json:"cod"`
	RateID       int64 `json:"rate_id"`
	UseInsurance bool  `json:"use_insurance"`
}

type DestinationRequest struct {
	Address string `json:"address"`
	AreaID  int64  `json:"area_id"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

type PackageRequest struct {
	Height      int64                `json:"height"`
	Items       []ItemPackageRequest `json:"items"`
	Length      int64                `json:"length"`
	PackageType int64                `json:"package_type"`
	Price       int64                `json:"price"`
	Weight      float64              `json:"weight"`
	Width       int64                `json:"width"`
}

type ItemPackageRequest struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Qty   int64  `json:"qty"`
}
