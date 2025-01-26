package utils

type ProgressBar struct {
	value  int
	length int
}

func (p *ProgressBar) SetValue(value int) {
	p.value = value
}

func (p *ProgressBar) SetLength(length int) {
	p.length = length
}

func (p *ProgressBar) GetProgress() int {
	return p.value * 100 / p.length
}
