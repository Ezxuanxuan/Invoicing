package api

import (
	"github.com/Invoicing/cookie"
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"time"
)

//Route
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.QueryParam("username")
		password := c.QueryParam("password")

		//MD5加密
		psswd := cookie.MD5(password)

		//用户名长度小于1
		if len(username) < 1 {
			return sendError(errors.INPUT_USER_ERROR, c)
		}

		//不存在该用户
		if models.GetUserCountbyUsername(username) < 1 {
			return sendError(errors.USER_NOT_EXI, c)
		}
		pass := models.GetPasswordbyUsername(username)
		//密码不正确
		if pass != psswd {
			return sendError(errors.PASS_ERROR, c)
		}
		//获取用户id
		userId := models.GetIdbyUsername(username)
		//加密并返回给前端
		IdValue := cookie.EncryptionId(userId)

		return sendSuccess(1, IdValue, "账号密码校验成功", c)
	}
}

//注册用户
func CreateStaff() echo.HandlerFunc {
	return func(c echo.Context) error {
		Name := c.FormValue("name")
		EnglishName := c.FormValue("english_name")
		Password := c.FormValue("password")
		Birthday := c.FormValue("birthday")
		IdCard := c.FormValue("id_card")
		Telephone := c.FormValue("telephone")
		//姓名为空
		if Name == "" {
			return sendError(errors.NAME_ERROR, c)
		}
		//英文名为空
		if EnglishName == "" {
			return sendError(errors.ENGLISH_ERROR, c)
		}
		//身份证不合法
		if len(IdCard) != 18 {
			return sendError(errors.IDCARD_ERROR, c)
		}

		if models.IsExitName(Name) {
			return sendError(errors.NAME_ERROR, c)
		}

		if models.IsExitEnglishName(EnglishName) {
			return sendError(errors.ENGLISH_ERROR, c)
		}

		//	local, _ := time.LoadLocation("local")
		birthday, err := time.ParseInLocation("2006-01-02 15:04:05", Birthday, time.Local)
		//日期转换出错
		if err != nil {
			return sendError(errors.BIRTHDAY_ERROR, c)
		}

		//电话号码非法
		if !isAllNumic(Telephone) {
			return sendError(errors.PHONE_ERROR, c)
		}

		//MD5加密
		password := cookie.MD5(Password)

		staff := models.Staff{
			Name:        Name,
			EnglishName: EnglishName,
			Password:    password,
			Birthday:    birthday,
			IdCard:      IdCard,
			Telephone:   Telephone,
		}
		if !models.CreateUser(staff) {
			return sendError(errors.DO_ERROR, c)
		}
		//successful := &errors.Successful{1, "添加用户成功"}
		return sendSuccess(1, "", "添加用户成功", c)
	}
}

func GetAllStaff() echo.HandlerFunc {
	return func(c echo.Context) error {
		staffs := models.GetAllStaff()
		return sendSuccess(1, staffs, "所有staff信息", c)
	}
}

func ModifyPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserId := c.FormValue("id")
		Password := c.FormValue("password")

		id := cookie.DecryptId(UserId)

		//校验id是否存在
		if !models.GetStaffById(id) {
			return sendError(errors.USER_NOT_EXI, c)
		}

		if Password == "" {
			return sendError(errors.PASS_ERROR, c)
		}
		password := cookie.MD5(Password)
		res := models.UpdatePassword(id, password)
		if !res {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "更新密码成功", c)
	}
}

func UpdateTelephone() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserId := c.FormValue("id")
		Telephone := c.FormValue("telephone")

		//更新前先解密id
		id := cookie.DecryptId(UserId)

		//校验id是否存在
		if !models.GetStaffById(id) {
			return sendError(errors.USER_NOT_EXI, c)
		}

		//将string的电话转成int类型。
		if !isAllNumic(Telephone) || Telephone == "" {
			return sendError(errors.PHONE_ERROR, c)
		}
		res := models.UpdateTelephone(id, Telephone)
		if !res {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "更新电话号码成功", c)
	}
}