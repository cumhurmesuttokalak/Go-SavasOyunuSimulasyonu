package oyun

type ISavascı interface {
	AtesEt()
	HasarAl()
}
type Savascı struct {
	Name        string
	Can         int
	MermiSayisi int
}
type Dunyali struct {
	Savascı
	Birlik string
}
type Marsli struct {
	Savascı
	UzayGemisi string
}

func (d *Dunyali) HasarAl() {
	if d.Can > 0 {
		d.Can -= 10
	}
}
func (d *Dunyali) AtesEt() {
	if d.MermiSayisi > 0 {
		d.MermiSayisi--
	}
}
func (m *Marsli) HasarAl() {
	if m.Can > 0 {
		m.Can -= 10
	}
}
func (m *Marsli) AtesEt() {
	if m.MermiSayisi > 0 {
		m.MermiSayisi--
	}
}
