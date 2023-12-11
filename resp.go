package alistsdk

type CommonResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginResp struct {
	CommonResp
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type Login2FAResp struct {
	CommonResp
	Data struct {
		QR     string `json:"qr"`
		Secret string `json:"secret"`
	} `json:"data"`
}

type UserInfoResp struct {
	CommonResp
	Data *User `json:"data"`
}

type DirsResp struct {
	CommonResp
	Data []Dir `json:"data"`
}

type ListResp struct {
	CommonResp
	Data struct {
		Content  []File `json:"content"`
		Total    int    `json:"total"`
		Readme   string `json:"readme"`
		Write    bool   `json:"write"`
		Provider string `json:"provider"`
	} `json:"data"`
}

type GetResp struct {
	CommonResp
	Data *File `json:"data"`
}

type SettingsResp struct {
	CommonResp
	Data *Settings `json:"data"`
}
