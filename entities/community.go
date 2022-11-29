package entities

type Community struct {
	UID      uint    `json:"id"`
	Name     string  `json:"name"`
	Photo    string  `json:"photo"`
	Win      float64 `json:"win"`
	Roe      float64 `json:"roe"`
	Day      float64 `json:"day"`
	Premium  string  `json:"premium"`
	Monetize string  `json:"monetize"`
	Verified string  `json:"verified"`
	Packages string  `json:"packages"`
	Follower float64 `json:"follower"`
	Signal   float64 `json:"signal"`
	Desc     string  `json:"desc"`
}
