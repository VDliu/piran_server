package req

type RegisterReq struct {
	Nick_name string
	Invite_code string
	Password string
}

type ChargeReq struct {
	Id int64 //充值人id
	Amount float64 //充值金额
	Invite_code string //充值对象
}
