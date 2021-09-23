package api

// CheckPermissions will check the permissions of the user against the context which contains the user ID from the JWT token.
//It will return true if the user has permission to access the resource, and false otherwise.
//If the function returns false, it will also return a http error message that can be returned from the handler.
//func CheckPermissions(c echo.Context, userID string) (bool, error) {
//	// extract the payload
//	payload, ok := c.Get("Payload").(UserPayload)
//	// we failed to bind the user payload, so just assume there is an error and do not let the user go further
//	if !ok {
//		return false, unAuthorizedMsg
//	}
//
//	// if the user id's do not match, they are not authorized to access the resource
//	if payload.ID != userID {
//		return false, unAuthorizedMsg
//	}
//
//	// they passed our checks, they can access the resource
//	return true, nil
//}
