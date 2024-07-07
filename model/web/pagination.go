package web

type PaginationRequest struct {
	Limit      int `form:"limit" binding:"omitempty"`
	Page       int `form:"page" binding:"omitempty"`
	TotalPages int
	TotalData  int64
}

func (p *PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationRequest) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PaginationRequest) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
