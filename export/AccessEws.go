package export

//AccessEws 访问EWS
type AccessEws struct {
	CURL
}

//NewAccessEws
func NewAccessEws() *AccessEws {
	return &AccessEws{}
}
