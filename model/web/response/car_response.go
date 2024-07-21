package response

type CarResponse struct {
	ID                  uint    `json:"id" example:"1" extensions:"x-order=0"`
	BrandID             uint    `json:"brand_id" example:"2" extensions:"x-order=1"`
	Name                string  `json:"name" example:"Toyota Yaris" extensions:"x-order=2"`
	Model               string  `json:"model" example:"SUV" extensions:"x-order=2"`
	Year                int16   `json:"year" example:"2020" extensions:"x-order=3"`
	ImageUrl            string  `json:"image_url" example:"image url" extensions:"x-order=4"`
	Width               int16   `json:"width" example:"462" extensions:"x-order=5"`
	Height              int16   `json:"height" example:"184" extensions:"x-order=6"`
	Length              int16   `json:"length" example:"137" extensions:"x-order=7"`
	Engine              string  `json:"engine" example:"2.0L EA113 CDLA TFSI In-Line 4 + Mild Hybrid 48V" extensions:"x-order=8"`
	Torque              int16   `json:"torque" example:"370" extensions:"x-order=9"`
	Transmission        string  `json:"transmission" example:"Manual" extensions:"x-order=10"`
	Acceleration        float32 `json:"acceleration" example:"5.6" extensions:"x-order=11"`
	HorsePower          int16   `json:"horse_power" example:"265" extensions:"x-order=12"`
	BreakingSystemFront string  `json:"breaking_system_front" example:"Ventilated Disc" extensions:"x-order=13"`
	BreakingSystemBack  string  `json:"breaking_system_back" example:"Disc" extensions:"x-order=14"`
	Fuel                string  `json:"fuel" example:"Electric" extensions:"x-order=15"`
}
