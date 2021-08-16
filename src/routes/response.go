package routes

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Country  string `json:"country"`
	UserID   int    `json:"userid"`
	Address  string `json:"address"`
}
