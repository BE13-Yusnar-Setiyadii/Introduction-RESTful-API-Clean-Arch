package delivery

import "yusnar/clean-arch/features/book"

type BookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	PublishYear string `json:"publish_year"`
	UserID      uint   `json:"user_id"`
}

func coreToResponse(core book.Core) BookResponse {
	response := BookResponse{
		ID:          core.ID,
		Title:       core.Title,
		Publisher:   core.Publisher,
		Author:      core.Author,
		PublishYear: core.PublishYear,
		UserID:      core.UserID,
	}
	return response

}

func responseList(listRes []book.Core) []BookResponse {
	var resList []BookResponse
	for _, v := range listRes {
		resList = append(resList, coreToResponse(v))

	}
	return resList

}
