package Ets

type ThreeDS struct {
	Pares string `json:"pares"`
}

func NewThreeDS(pares string) *ThreeDS {
	return &ThreeDS{
		Pares: pares,
	}
}
