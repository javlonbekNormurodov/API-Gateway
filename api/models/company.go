package models

type CreateCompany struct {
	Name string `json:"name" binding:"required"`
}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetAllCompanyResponse struct {
	Companies []Company `json:"professions"`
	Count       uint32       `json:"count"`
}

type MsgResponse struct {
  Msg string `json:"message"`
}