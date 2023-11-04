package main

import (
	"context"
	"syscall"

	"github.com/lxn/win"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	var val *uint16
	var err error
	val, err = syscall.UTF16PtrFromString("RainbowOverlay")
	if err != nil {
		print(err)
	}
	hwnd := win.FindWindow(nil, val)
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
}
