package repo

import (
	e "main/internal/core/error"
	"main/internal/core/model"
)

func CheckPermission(userID uint, object model.Object) error {
	if userID != object.GetOwnerID() {
		return e.ErrPermissionDenied
	}
	return nil
}
