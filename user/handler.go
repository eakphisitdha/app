package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser Profile
	// เรียก BindJSON เพื่อผูก JSON ที่รับมากับ newUser
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	if IsExits(newUser.Phone) {
		//ถ้ามีแล้ว return หมายเลขนี้ได้ลงทะเบียนไปแล้ว (TEXT) Status 200
		c.String(http.StatusOK, "หมายเลขนี้ได้ลงทะเบียนไปแล้ว")

	} else {
		if err := AddUser(newUser); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		//ถ้าไม่เจอ return User Profile (JSON) Status 201
		c.JSON(http.StatusCreated, newUser)
	}
}
