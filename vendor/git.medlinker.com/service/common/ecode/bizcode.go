package ecode

// Business code ( 0 <= x <= 1000).
// 错误码保证先注册后使用原则，以防止定义冲突，统一注册地址：
// https://wiki.medlinker.com/pages/viewpage.action?pageId=17598142
var (
	OK                   = add(0)    // 正确
	CommonErr            = add(3)    // 通用错误(当你不知道该如何归类时使用此错误)
	OperDbErr            = add(2)    // 操作db错误
	ParamInvalid         = add(1)    // 参数错误
	UserNotLogin         = add(101)  // 用户未登录
	UserSessionExpired   = add(102)  // 用户SESSION已过期
	UserNotAuthorized    = add(103)  // 鉴权失败，当前用户没有操作权限
	NeedWechatAuthorized = add(104)  // 需要额外的微信授权
	UserHasNoPhone       = add(105)  // 当前用户未绑定手机号
	SrvParamInvalid      = add(1001) // 服务调用参数出错
	SrvOperDbErr         = add(1002) // 服务操作数据库出错
	SrvCommonErr         = add(1003) // 服务通用错误
)
