package usersDb

type UserDB struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Status      int    `json:"status"`
	PostbackUrl string `json:"postbackUrl"`
}
