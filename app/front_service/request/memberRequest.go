package request

type CreateMemberRequest struct {
	Name     string `json:"name" label:"名稱" validate:"required"`     //會員名稱
	Account  string `json:"account" label:"帳號" validate:"required"`  //會員帳號
	Password string `json:"password" label:"密碼" validate:"required"` //會員密碼
}
