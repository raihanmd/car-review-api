package request

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
	BrandID             uint    `json:"brand_id" binding:"omitempty" extensions:"x-order=0"`
	Model               string  `json:"model" binding:"omitempty" extensions:"x-order=1"`
	Year                int16   `json:"year" binding:"omitempty,min=1878" extensions:"x-order=2"`
	ImageUrl            string  `json:"image_url" binding:"omitempty,url" extensions:"x-order=3"`
	Width               int16   `json:"width" binding:"omitempty" extensions:"x-order=4"`
	Height              int16   `json:"height" binding:"omitempty" extensions:"x-order=5"`
	Length              int16   `json:"length" binding:"omitempty" extensions:"x-order=6"`
	Engine              string  `json:"engine" binding:"omitempty" extensions:"x-order=7"`
	Torque              int16   `json:"torque" binding:"omitempty" extensions:"x-order=8"`
	Transmission        string  `json:"transmission" binding:"omitempty" extensions:"x-order=9"`
	Acceleration        float32 `json:"acceleration" binding:"omitempty" extensions:"x-order=10"`
	HorsePower          int16   `json:"horse_power" binding:"omitempty" extensions:"x-order=11"`
	BreakingSystemFront string  `json:"breaking_system_front" binding:"omitempty" extensions:"x-order=12"`
	BreakingSystemBack  string  `json:"breaking_system_back" binding:"omitempty" extensions:"x-order=13"`
	Fuel                string  `json:"fuel" binding:"omitempty" extensions:"x-order=14"`
}
