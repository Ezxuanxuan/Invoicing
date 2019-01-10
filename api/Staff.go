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
		username := c.FormValue("username")
		password := c.FormValue("password")

		//MD5加密
		psswd := cookie.MD5(password)

		//用户名长度小于1
		if len(username) < 1 {
			return sendError(errors.INPUT_USER_ERROR, c)
		}

		total, err := models.GetUserCountbyUsername(username)
		//不存在该用户
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if total < 1 {
			return sendError(errors.USER_NOT_EXI, c)
		}

		pass, err := models.GetPasswordbyUsername(username)
		//密码不正确
		if err != nil {
			return err
		}
		if pass != psswd {
			return sendError(errors.PASS_ERROR, c)
		}
		//获取用户id
		userId, err := models.GetIdbyUsername(username)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		//加密并返回给前端
		IdValue := cookie.EncryptionId(userId)
		permission, err := models.GetPermissionByStaff(userId)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		revalue := reValue{
			Id:         IdValue,
			Permission: permission,
		}

		return sendSuccess(1, revalue, "账号密码校验成功", c)
	}
}

type reValue struct {
	Id         string
	Permission string
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

		//判断姓名是否存在
		has, err := models.IsExitName(Name)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if has {
			return sendError(errors.NAME_ERROR, c)
		}

		//p判断英文名是否存在
		has2, err := models.IsExitEnglishName(Name)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if has2 {
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

		staff := models.Staffs{
			Name:        Name,
			EnglishName: EnglishName,
			Password:    password,
			Birthday:    birthday,
			IdCard:      IdCard,
			Telephone:   Telephone,
		}

		affected, err := models.CreateUser(staff)
		if err != nil || affected != 1 {
			return sendError(errors.DO_ERROR, c)
		}
		//successful := &errors.Successful{1, "添加用户成功"}

		id, _ := models.GetIdbyUsername(Name)
		models.InitPermission(id)
		return sendSuccess(1, "", "添加用户成功", c)
	}
}

func GetAllStaff() echo.HandlerFunc {
	return func(c echo.Context) error {
		staffs, err := models.GetAllStaff()
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, staffs, "所有staff信息", c)
	}
}

func ModifyPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserId := c.FormValue("id")
		Password := c.FormValue("password")

		id := cookie.DecryptId(UserId)

		//校验id是否存在
		has, err := models.IsExitStaffById(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.USER_NOT_EXI, c)
		}

		if Password == "" {
			return sendError(errors.PASS_ERROR, c)
		}

		password := cookie.MD5(Password)
		affected, err := models.UpdatePassword(id, password)
		if err != nil || affected < 1 {
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
		has, err := models.IsExitStaffById(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.USER_NOT_EXI, c)
		}

		//将string的电话转成int类型。
		if !isAllNumic(Telephone) || Telephone == "" {
			return sendError(errors.PHONE_ERROR, c)
		}
		affected, err := models.UpdateTelephone(id, Telephone)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "更新电话号码成功", c)
	}
}
