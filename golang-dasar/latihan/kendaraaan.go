package latihan

type Kendaraan interface {
	CaraJalan() string
	SetWarna(string)
	GetWarna() string
	GetRoda() int
}

type becak struct {
	roda  int
	warna string
}

func (o *becak) SetWarna(s string) {
	o.warna = s
}

func (o *becak) GetWarna() string {
	return o.warna
}

func (o *becak) GetRoda() int {
	return 3
}

func (o *becak) CaraJalan() string {
	return "dikayuh"
}

// NewBecak function untuk membuat objek becak
func NewBecak() Kendaraan {
	return &becak{}
}
