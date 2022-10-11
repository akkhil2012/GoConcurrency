package main

type ak47 struct {
	Gun
}

func newAk47() iGun {
	return &ak47{
		Gun: Gun{
			name: "Ak47",
		},
	}
}
