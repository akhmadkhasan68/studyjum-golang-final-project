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
	Metadata MetadataResponse        `json:"metadata"`
	Data     CreateOrderResponseData `json:"data"`
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

type CancelOrderResponse struct {
	Metadata MetadataResponse        `json:"metadata"`
	Data     DataCancelOrderResponse `json:"data"`
}

type DataCancelOrderResponse struct {
	CancelOrder CancelOrderDataResponse `json:"cancel_order"`
}

type CancelOrderDataResponse struct {
	OrderID string `json:"order_id"`
	Cancel  bool   `json:"cancel"`
}

type MetadataResponse struct {
	Path           string `json:"path"`
	HTTPStatusCode int64  `json:"http_status_code"`
	HTTPStatus     string `json:"http_status"`
	Timestamp      int64  `json:"timestamp"`
}

type GetOrderDetailResponse struct {
	Metadata MetadataResponse        `json:"metadata"`
	Data     DataOrderDetailResponse `json:"data"`
}

type DataOrderDetailResponse struct {
	Consignee        Consigne         `json:"consignee"`
	Consigner        Consigne         `json:"consigner"`
	Origin           Destination      `json:"origin"`
	Destination      Destination      `json:"destination"`
	ExternalID       string           `json:"external_id"`
	OrderID          string           `json:"order_id"`
	Courier          Courier          `json:"courier"`
	Package          Package          `json:"package"`
	ProofOfDelivery  ProofOfDelivery  `json:"proof_of_delivery"`
	TimeSlotSelected TimeSlotSelected `json:"time_slot_selected"`
	PaymentType      string           `json:"payment_type"`
	Driver           Driver           `json:"driver"`
	LabelCheckSum    string           `json:"label_check_sum"`
	CreationDate     string           `json:"creation_date"`
	LastUpdatedDate  string           `json:"last_updated_date"`
	AwbNumber        string           `json:"awb_number"`
	Trackings        []Tracking       `json:"trackings"`
}

type Consigne struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Courier struct {
	RateID          int64 `json:"rate_id"`
	Amount          int64 `json:"amount"`
	UseInsurance    bool  `json:"use_insurance"`
	InsuranceAmount int64 `json:"insurance_amount"`
	Cod             bool  `json:"cod"`
	MinDay          int64 `json:"min_day"`
	MaxDay          int64 `json:"max_day"`
}

type Destination struct {
	ID           int64  `json:"id"`
	StopID       int64  `json:"stop_id"`
	Address      string `json:"address"`
	Direction    string `json:"direction"`
	Postcode     string `json:"postcode"`
	AreaID       int64  `json:"area_id"`
	AreaName     string `json:"area_name"`
	SuburbID     int64  `json:"suburb_id"`
	SuburbName   string `json:"suburb_name"`
	CityID       int64  `json:"city_id"`
	CityName     string `json:"city_name"`
	ProvinceID   int64  `json:"province_id"`
	ProvinceName string `json:"province_name"`
	CountryID    int64  `json:"country_id"`
	CountryName  string `json:"country_name"`
	Lat          string `json:"lat"`
	Lng          string `json:"lng"`
}

type Driver struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	VehicleType   string `json:"vehicle_type"`
	VehicleNumber string `json:"vehicle_number"`
}

type Package struct {
	Weight       float64 `json:"weight"`
	Length       int64   `json:"length"`
	Width        int64   `json:"width"`
	Height       int64   `json:"height"`
	VolumeWeight int64   `json:"volume_weight"`
	PackageType  int64   `json:"package_type"`
	Items        []Item  `json:"items"`
}

type Item struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Qty   int64  `json:"qty"`
}

type ProofOfDelivery struct {
	Photo     string `json:"photo"`
	Signature string `json:"signature"`
}

type TimeSlotSelected struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Tracking struct {
	ShipperStatus  Status `json:"shipper_status"`
	LogisticStatus Status `json:"logistic_status"`
	CreatedDate    string `json:"created_date"`
}

type Status struct {
	Code        int64  `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
