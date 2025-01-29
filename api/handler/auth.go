package handler

import (
	"fmt"
	"net/http"
	"weather/pkg/model"

	"github.com/gin-gonic/gin"
)

// Register ro'yxatdan o'tish uchun foydalanuvchi so'rovini qayta ishlaydi
// @Summary      Ro'yxatdan o'tish
// @Description  Foydalanuvchi ma'lumotlarini qabul qilib, yangi akkaunt yaratadi
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        register body model.RegisterReq true "Ro'yxatdan o'tish ma'lumotlari"
// @Success      201  {object}  model.RegisterResp  "Muvaffaqiyatli ro'yxatdan o'tish"
// @Failure      400  {string}  string      "Noto'g'ri so'rov formati"
// @Failure      500  {string}  string      "Server ichki xatosi"
// @Router       /register [post]
func(H *Handler) Register(c *gin.Context){
	req := model.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		H.Log.Error(fmt.Sprintf("Error is read request body: %v", err))
		c.JSON(http.StatusBadRequest, "So'rov uchun kitirilgan ma'lumotlar noto'g'ri")
		return
	}
	resp, err := H.Service.Register(&req)
	if err != nil {
		H.Log.Error(fmt.Sprintf("Error is registration: %v", err))
		c.JSON(http.StatusInternalServerError, "Server bilan xatolik yuz berdi")
		return 
	}
	c.JSON(201, resp)
}