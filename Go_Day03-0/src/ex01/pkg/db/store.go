package db

type Store interface {
	GetPlaces(limit, offset int) ([]Place, int, error)
}

type Place struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
