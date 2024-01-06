package gosteamweb

type App struct {
	Id                       string `json:"appid"`
	Name                     string `json:"name"`
	PlaytimeTwoWeeks         int    `json:"playtime_2weeks"`
	PlaytimeForever          int    `json:"playtime_forever"`
	PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
	PlaytimeMacForever       int    `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
	ImgIconUrl               string `json:"img_icon_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	RTimeLastPlayed          int    `json:"rtime_last_played"`
	PlaytimeDisconnected     int    `json:"playtime_disconnected"`
}

func (a App) String() string {
	return a.Name
}
