package models

type ChangePassword struct {
	NewPassword string `json:"newPassword" query:"newPassword" form:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword" query:"confirmNewPassword" form:"confirmNewPassword"`
}
