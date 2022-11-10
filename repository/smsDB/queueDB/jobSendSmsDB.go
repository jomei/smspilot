package queueDB

type JobSendSmsDB struct {
	Id           int    `json:"id,omitempty"`
	Userid       int    `json:"userid,omitempty"`
	Message      string `json:"message,omitempty"`
	Phone        int    `json:"phone,omitempty"`
	DatetimeSend string `json:"datetimeSend,omitempty"`
}
