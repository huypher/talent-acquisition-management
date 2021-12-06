package level

//func (d *talentDelivery) getUser() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		uid := auth.UIDFromContext(c)
//		if uid == (auth.UID{}) {
//			http_response.Response(c, http.StatusBadRequest, "invalid request", nil)
//			return
//		}
//
//		http_response.Success(c, "success", container.Map{
//			"id":       uid.ID,
//			"username": uid.Username,
//			"name":     uid.Username,
//		})
//	}
//}
