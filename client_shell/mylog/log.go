package mylog

import (
	"fmt"
	"runtime"
	"syscall"
)

var (
	G_Current_Platform string
	kernel32           *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	proc               *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle        *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)

	// 给字体颜色对象赋值
	FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

type Color struct {
	black        int // 黑色
	blue         int // 蓝色
	green        int // 绿色
	cyan         int // 青色
	red          int // 红色
	purple       int // 紫色
	yellow       int // 黄色
	light_gray   int // 淡灰色（系统默认值）
	gray         int // 灰色
	light_blue   int // 亮蓝色
	light_green  int // 亮绿色
	light_cyan   int // 亮青色
	light_red    int // 亮红色
	light_purple int // 亮紫色
	light_yellow int // 亮黄色
	white        int // 白色
}

func init() {
	G_Current_Platform = runtime.GOOS
}

func Black(str string) string {
	return textColor(textBlack, str)
}

func Red(str string) string {
	return textColor(textRed, str)
}
func Yellow(str string) string {
	return textColor(textYellow, str)
}
func Green(str string) string {
	return textColor(textGreen, str)
}
func Cyan(str string) string {
	return textColor(textCyan, str)
}
func Blue(str string) string {
	return textColor(textBlue, str)
}
func Purple(str string) string {
	return textColor(textPurple, str)
}
func White(str string) string {
	return textColor(textWhite, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}

// win 	输出有颜色的字体
func ColorPrint(s string, i int) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	print(s)
	CloseHandle.Call(handle)
}

func Error(msg string) {
	if G_Current_Platform == "windows" {
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.red))
		fmt.Println(msg)
		CloseHandle.Call(handle)
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.white))
		CloseHandle.Call(handle)
	} else {
		fmt.Println(Red(msg))
	}
}

func Warning(msg string) {
	if G_Current_Platform == "windows" {
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.yellow))
		fmt.Println(msg)
		CloseHandle.Call(handle)
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.white))
		CloseHandle.Call(handle)
	} else {
		fmt.Println(Yellow(msg))
	}
}
func Success(msg string) {
	if G_Current_Platform == "windows" {
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.green))
		fmt.Println(msg)
		CloseHandle.Call(handle)

		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.white))
		CloseHandle.Call(handle)
	} else {
		fmt.Println(Green(msg))
	}
}
func Info(msg string) {
	if G_Current_Platform == "windows" {
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.blue))
		fmt.Println(msg)
		CloseHandle.Call(handle)

		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(FontColor.white))
		CloseHandle.Call(handle)
	} else {
		fmt.Println(Blue(msg))
	}
}
