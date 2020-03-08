package transport

// page参数
type Page struct {
	PageSize   uint `form:"page" json:"username" binding:"required"`
	PageNumber uint `form:"page" json:"page" binding:"required"`
}

func (p *Page) GetPageSize() uint {
	if p.PageSize == 0 {
		return 10
	}
	return p.PageSize
}

func (p *Page) GetPageNumber() uint {
	if p.PageNumber == 0 {
		return 1
	}
	return p.PageNumber
}
