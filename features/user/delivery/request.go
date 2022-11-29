package delivery

import "yusnar/clean-arch/features/user"

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Telp_number string `json:"telp_number" form:"telp_number"`
	Address     string `json:"address" form:"address"`
	Role        string `json:"role" form:"role"`
}

func (req *UserRequest) reqToCore() user.Core {
	return user.Core{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		Telp_number: req.Telp_number,
		Address:     req.Address,
		Role:        req.Role,
	}

}
