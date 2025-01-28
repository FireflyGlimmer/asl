package utils

import (
	"fmt"
	"regexp"
	"strings"
)

type ProgressBar struct {
	current    int
	total      int
	status     string
	lastStrLen int
}

func (p *ProgressBar) SetCurrentValue(current int) {
	p.current = current
}

func (p *ProgressBar) SetTotalValue(total int) {
	p.total = total
}

func (p *ProgressBar) SetStatus(status bool) {
	if status {
		p.status = "Done"
	} else {
		p.status = "Failed"
	}
}

func (p *ProgressBar) Print() {
	percentage := float64(p.current) / float64(p.total) * 100
	bar := fmt.Sprintf("[ %.2f%% ]", percentage)
	if percentage == 100 {
		p.SetStatus(true)
	}
	if p.status == "Done" {
		bar = fmt.Sprintf("[ \033[32mDone\033[0m ]")
	} else if p.status == "Failed" {
		bar = fmt.Sprintf("[ \033[31mFailed\033[0m ]")
	}

	// 使用正则表达式去除ANSI转义码
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	cleanBar := re.ReplaceAllString(bar, "")

	fmt.Print("\0337")                    // 保存当前光标位置
	fmt.Print("\033[A")                   // 光标上移一行
	fmt.Print("\033[999C")                // 移动到行尾
	fmt.Printf("\033[%dD", len(cleanBar)) // 根据字符长度向左移动
	if p.lastStrLen > len(cleanBar) {
		fmt.Printf("\033[%dD", p.lastStrLen-len(cleanBar))         // 根据上一次的字符长度向左移动
		fmt.Print(strings.Repeat(" ", p.lastStrLen-len(cleanBar))) // 用空格覆盖多余的字符
	}
	p.lastStrLen = len(cleanBar)
	fmt.Print(bar)
	fmt.Print("\0338") // 恢复光标位置
}

func NewProgressBar(current, total int) *ProgressBar {
	return &ProgressBar{current: current, total: total}
}
