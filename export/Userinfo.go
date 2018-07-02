package export

//Userinfo user infomation struct
type Userinfo struct {
	UserID     string `json:"user_id"`
	UserName   string `json:"username"`
	Orgroot    string
	Orgcode    string
	CustomerId string
	Name       string
	Theme      string
}
