package delivery

import "yusnar/clean-arch/features/user"

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Telp_number string `json:"telp_number"`
	Address     string `json:"address"`
	Role        string `json:"role"`
}

func coreToResponse(core user.Core) UserResponse {
	response := UserResponse{
		ID:          core.ID,
		Name:        core.Name,
		Email:       core.Email,
		Telp_number: core.Telp_number,
		Address:     core.Address,
		Role:        core.Role,
	}
	return response

}

func responseList(listRes []user.Core) []UserResponse {
	var resList []UserResponse
	for _, v := range listRes {
		resList = append(resList, coreToResponse(v))

	}
	return resList

}
