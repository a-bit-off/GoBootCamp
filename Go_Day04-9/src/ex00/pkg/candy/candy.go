package candy

type CandyRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type Candy struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CandyResponse struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}
