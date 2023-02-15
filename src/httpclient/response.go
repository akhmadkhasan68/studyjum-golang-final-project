package httpclient

type Metadata struct {
	Path       string `json:"path"`
	StatusCode uint8  `json:"http_status_code"`
	Status     string `json:"http_status"`
	Timestamp  int64  `json:"timestamp"`
}

type Pagination struct {
	CurrentPage     uint64 `json:"current_page"`
	CurrentElements uint64 `json:"current_elements"`
	TotalPage       uint64 `json:"total_pages"`
	TotalElements   uint64 `json:"total_elements"`
}

type GetProvinces struct {
	Metadata   Metadata   `json:"metadata"`
	Data       []Province `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Province struct {
	ID      uint64  `json:"id"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Country Country `json:"country"`
}

type Country struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateOrderResponse struct {
	Metadata CreateOrderResponseMetadata `json:"metadata"`
	Data     CreateOrderResponseData     `json:"data"`
}

type CreateOrderResponseData struct {
	Coverage    string                         `json:"coverage"`
	OrderID     string                         `json:"order_id"`
	PaymentType string                         `json:"payment_type"`
	Courier     CourierCreateOrderResponse     `json:"courier"`
	Consignee   ConsigneCreateOrderResponse    `json:"consignee"`
	Consigner   ConsigneCreateOrderResponse    `json:"consigner"`
	Destination DestinationCreateOrderResponse `json:"destination"`
	Origin      DestinationCreateOrderResponse `json:"origin"`
	Package     PackageOrderResponse           `json:"package"`
}

type ConsigneCreateOrderResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type CourierCreateOrderResponse struct {
	RateID          int64 `json:"rate_id"`
	Amount          int64 `json:"amount"`
	UseInsurance    bool  `json:"use_insurance"`
	InsuranceAmount int64 `json:"insurance_amount"`
	Cod             bool  `json:"cod"`
}

type DestinationCreateOrderResponse struct {
	Address      string `json:"address"`
	AreaID       int64  `json:"area_id"`
	AreaName     string `json:"area_name"`
	CityID       int64  `json:"city_id"`
	CityName     string `json:"city_name"`
	CountryID    int64  `json:"country_id"`
	CountryName  string `json:"country_name"`
	Lat          string `json:"lat"`
	Lng          string `json:"lng"`
	Postcode     string `json:"postcode"`
	ProvinceID   int64  `json:"province_id"`
	ProvinceName string `json:"province_name"`
	SuburbID     int64  `json:"suburb_id"`
	SuburbName   string `json:"suburb_name"`
}

type PackageOrderResponse struct {
	PackageType int64                      `json:"package_type"`
	Weight      float64                    `json:"weight"`
	Length      int64                      `json:"length"`
	Width       int64                      `json:"width"`
	Height      int64                      `json:"height"`
	Price       int64                      `json:"price"`
	Items       []ItemPackageOrderResponse `json:"items"`
}

type ItemPackageOrderResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Qty   int64  `json:"qty"`
	Price int64  `json:"price"`
}

type CreateOrderResponseMetadata struct {
	Path           string `json:"path"`
	HTTPStatusCode int64  `json:"http_status_code"`
	HTTPStatus     string `json:"http_status"`
	Timestamp      int64  `json:"timestamp"`
}
