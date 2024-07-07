package request

type CarQueryRequest struct {
	BrandID *uint   `form:"brand_id"`
	Model   *string `form:"model"`
	MinYear *int    `form:"min_year"`
	MaxYear *int    `form:"max_year"`
}

type CarCreateRequest struct {
	BrandID             uint    `json:"brand_id" binding:"required" extensions:"x-order=0"`
	Model               string  `json:"model" binding:"required" extensions:"x-order=1"`
	Year                int16   `json:"year" binding:"required,min=1878" extensions:"x-order=2"`
	ImageUrl            string  `json:"image_url" binding:"required,url" extensions:"x-order=3"`
	Width               int16   `json:"width" binding:"required" extensions:"x-order=4"`
	Height              int16   `json:"height" binding:"required" extensions:"x-order=5"`
	Length              int16   `json:"length" binding:"required" extensions:"x-order=6"`
	Engine              string  `json:"engine" binding:"required" extensions:"x-order=7"`
	Torque              int16   `json:"torque" binding:"required" extensions:"x-order=8"`
	Transmission        string  `json:"transmission" binding:"required" extensions:"x-order=9"`
	Acceleration        float32 `json:"acceleration" binding:"required" extensions:"x-order=10"`
	HorsePower          int16   `json:"horse_power" binding:"required" extensions:"x-order=11"`
	BreakingSystemFront string  `json:"breaking_system_front" binding:"required" extensions:"x-order=12"`
	BreakingSystemBack  string  `json:"breaking_system_back" binding:"required" extensions:"x-order=13"`
	Fuel                string  `json:"fuel" binding:"required" extensions:"x-order=14"`
}

type CarUpdateRequest struct {
	BrandID             uint    `json:"brand_id" extensions:"x-order=0"`
	Model               string  `json:"model" extensions:"x-order=1"`
	Year                int16   `json:"year" binding:"omitempty,min=1878" extensions:"x-order=2"`
	ImageUrl            string  `json:"image_url" binding:"omitempty,url" extensions:"x-order=3"`
	Width               int16   `json:"width" extensions:"x-order=4"`
	Height              int16   `json:"height" extensions:"x-order=5"`
	Length              int16   `json:"length" extensions:"x-order=6"`
	Engine              string  `json:"engine" extensions:"x-order=7"`
	Torque              int16   `json:"torque" extensions:"x-order=8"`
	Transmission        string  `json:"transmission" extensions:"x-order=9"`
	Acceleration        float32 `json:"acceleration" extensions:"x-order=10"`
	HorsePower          int16   `json:"horse_power" extensions:"x-order=11"`
	BreakingSystemFront string  `json:"breaking_system_front" extensions:"x-order=12"`
	BreakingSystemBack  string  `json:"breaking_system_back" extensions:"x-order=13"`
	Fuel                string  `json:"fuel" extensions:"x-order=14"`
}
