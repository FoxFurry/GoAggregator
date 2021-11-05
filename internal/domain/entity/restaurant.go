package entity

type Restaurant struct {
	Name string `json:"name"`
	Address string `json:"address"`
	MenuItems int `json:"menu_items"`
	Menu []Food `json:"menu"`
	Rating float32 `json:"rating"`
}
