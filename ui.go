package main

import (
	"log"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/yourok/go-mpv/mpv"
)

var pauseTextState = 0

func gfxMain(win *glfw.Window, ctx *nk.Context, state *State, mgl *mpv.MpvGL) {
	nk.NkPlatformNewFrame()

	// Layout
	bounds := nk.NkRect(50, 50, 230, 250)
	update := nk.NkBegin(ctx, "Demo", bounds,
		nk.WindowBorder|nk.WindowMovable|nk.WindowScalable|nk.WindowMinimizable|nk.WindowTitle)

	if update > 0 {
		nk.NkLayoutRowStatic(ctx, 30, 80, 1)
		{
			pauseText := ""
			switch pauseTextState {
			case 0: pauseText = "Pause"
			case 1: pauseText = "Play"
			}
			if nk.NkButtonLabel(ctx, pauseText) > 0 {
				pause()
			}
		}
		nk.NkLayoutRowDynamic(ctx, 30, 2)
		{
			if nk.NkButtonLabel(ctx, "Prev") > 0 {
				log.Println("Next Pressed")
			}
			if nk.NkButtonLabel(ctx, "Next") > 0 {
				log.Println("Prev Pressed")
			}
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkEditBuffer(ctx, nk.EditField, &state.board, nk.NkFilterDefault)
			nk.NkEditBuffer(ctx, nk.EditField, &state.thread, nk.NkFilterDefault)
			if nk.NkButtonLabel(ctx, "Print Entered Text") > 0 {
				log.Printf("Requested thread %s from baord %s", state.thread.GetGoString(), state.board.GetGoString())
			}
		}
	}
	nk.NkEnd(ctx)

	// Render
	bg := make([]float32, 4)
	nk.NkColorFv(bg, state.bgColor)
	width, height := win.GetSize()
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(bg[0], bg[1], bg[2], bg[3])
	mgl.Draw(0, width, height)
	nk.NkPlatformRender(nk.AntiAliasingOn, maxVertexBuffer, maxElementBuffer)
}

type Option uint8

const (
	Easy Option = 0
	Hard Option = 1
)

type State struct {
	bgColor nk.Color
	prop    int32
	opt     Option
	board    nk.TextEdit
	thread   nk.TextEdit
}

func onError(code int32, msg string) {
	log.Printf("[glfw ERR]: error %d: %s", code, msg)
}

func flag(v bool) int32 {
	if v {
		return 1
	}
	return 0
}