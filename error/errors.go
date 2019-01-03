package errors

type Restful struct {
	Code int
	Msg  string
}
type Successful struct {
	Code int
	Data interface{}
	Msg  string
}

var (
	DO_ERROR                 = &Restful{-1000, "操作失败"}
	USER_NOT_EXI             = &Restful{-1001, "用户名不存在"}
	INPUT_USER_ERROR         = &Restful{-1002, "输入用户不合法"}
	PASS_ERROR               = &Restful{-1003, "密码错误"}
	NAME_ERROR               = &Restful{-1004, "姓名为空或者已存在"}
	ENGLISH_ERROR            = &Restful{-1005, "英文名为空或者已存在"}
	IDCARD_ERROR             = &Restful{-1005, "身份证非法"}
	BIRTHDAY_ERROR           = &Restful{-1006, "生日转换过程出错"}
	PHONE_ERROR              = &Restful{-1007, "电话号码非法"}
	STAFF_ID_ERROR           = &Restful{-1008, "staff_id 非法"}
	POWER_INPUT_ERROR        = &Restful{-1009, "权限类型输入错误"}
	INPUT_ERROR              = &Restful{-1010, "输入非法"}
	COMPONENT_NO_ERROR       = &Restful{-1011, "零件编码不能为空"}
	COMPONENT_NAME_ERROR     = &Restful{-1012, "零件名不能为空"}
	COMPONENT_MATERIAL_ERROR = &Restful{-1013, "零件材质输入不合法"}
	COMPONENT_QUALITY_ERROR  = &Restful{-1014, "零件质量输入不合法"}
	COMPONENT_NO_EXSIT       = &Restful{-1015, "零件编号已经存在"}
)