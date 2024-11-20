package models

type ErrorMessageResponse struct {
	ErrorMessage string `json:"error"`
}

type CurrentTimeRequest struct {
	Timezone string `json:"timezone"`
}

type CurrentTimeResponse struct {
	Timezone string `json:"timezone"`
	Time     string `json:"time"`
}

type TargetTimeRequest struct {
	TargetInfo []TargetTimeInfo `json:"target_infos"`
}

type TargetTimeResponse struct {
	TargetInfo []TargetTimeInfo `json:"target_infos"`
}

type TargetTimeInfo struct {
	Timezone string `json:"timezone"`
	Time     string `json:"time"`
}

type WorldTimeRequest struct {
	Timezone string `json:"timezone"`
}

type WorldTimeResponse struct {
	UserTime     string            `json:"user_time"`     // 사용자의 현재 시간
	UserTimezone string            `json:"user_timezone"` // 사용자의 타임존
	CityTimes    map[string]string `json:"city_times"`
}
