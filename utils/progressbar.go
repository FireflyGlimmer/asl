package utils

import (
	"fmt"
)

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

func (p *ProgressBar) Print() {
	percentage := float64(p.value) / float64(p.length) * 100
	fmt.Print("\0337")                   // 保存当前光标位置
	fmt.Print("\033[A")                  // 光标上移一行
	fmt.Print("\033[999C")               // 移动到行尾
	fmt.Print("\033[10D")                // 向左移动六位
	fmt.Printf("[ %.2f%% ]", percentage) // 在行尾添加百分比
	fmt.Print("\0338")                   // 恢复光标位置
}

func NewProgressBar(value, length int) *ProgressBar {
	return &ProgressBar{value: value, length: length}
}
