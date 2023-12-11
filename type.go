package alistsdk

type Dir struct {
	Name     string `json:"name"`
	Modified string `json:"modified"`
}

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Salt       string `json:"Salt"`
	Password   string `json:"password"`
	BasePath   string `json:"base_path"`
	Role       int    `json:"role"`
	Disabled   bool   `json:"disabled"`
	Permission int    `json:"permission"`
	SSOID      string `json:"sso_id"`
	OTP        bool   `json:"otp"`
}

type File struct {
	Name     string `json:"name"`
	Modified string `json:"modified"`
	Size     int64  `json:"size"`
	IsDir    bool   `json:"is_dir"`
	Sign     string `json:"sign"`
	Thumb    string `json:"thumb"`
	Type     int    `json:"type"`
	RawURL   string `json:"raw_url"`
}

type Settings struct {
	AllowIndexed           string `json:"allow_indexed"`
	AllowMounted           string `json:"allow_mounted"`
	Announcement           string `json:"announcement"`
	AudioAutoplay          string `json:"audio_autoplay"`
	AudioCover             string `json:"audio_cover"`
	AutoUpdateIndex        string `json:"auto_update_index"`
	DefaultPageSize        string `json:"default_page_size"`
	ExternalPreviews       string `json:"external_previews"`
	Favicon                string `json:"favicon"`
	FilenameCharMapping    string `json:"filename_char_mapping"`
	ForwardDrectLinkParams string `json:"forward_drect_link_params"`
	HideFiles              string `json:"hide_files"`
	HomeContainer          string `json:"home_container"`
	HomeIcon               string `json:"home_icon"`
	IframePreviews         string `json:"iframe_previews"`
	Logo                   string `json:"logo"`
	MainColor              string `json:"main_color"`
	OCRAPI                 string `json:"ocr_api"`
	PackageDownload        string `json:"package_download"`
	PaginationType         string `json:"pagination_type"`
	RobotsTXT              string `json:"robots_txt"`
	SearchIndex            string `json:"search_index"`
	SettingsLayout         string `json:"settings_layout"`
	SiteTitle              string `json:"site_title"`
	SSOLoginEnabled        string `json:"sso_login_enabled"`
	SSOLoginPlatform       string `json:"sso_login_platform"`
	Version                string `json:"version"`
	VideoAutoplay          string `json:"video_autoplay"`
}
