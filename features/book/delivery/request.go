package delivery

import "yusnar/clean-arch/features/book"

type BookRequest struct {
	Title       string `json:"title" form:"title"`
	Publisher   string `json:"publisher" form:"publisher"`
	Author      string `json:"author" form:"author"`
	PublishYear string `json:"publish_year" form:"publish_year"`
	UserID      uint   `json:"user_id" form:"user_id"`
}

func (req *BookRequest) reqToCore() book.Core {
	return book.Core{
		Title:       req.Title,
		Publisher:   req.Publisher,
		Author:      req.Author,
		PublishYear: req.PublishYear,
		UserID:      req.UserID,
	}

}
