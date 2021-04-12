package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/yourok/go-mpv/mpv"
)

var pauseTextState = 0

type Option uint8

const (
	Easy Option = 0
	Hard Option = 1
)

type UiState struct {
	bgColor nk.Color
	prop    int32
	opt     Option
	board   nk.TextEdit
	thread  nk.TextEdit
	connIp  nk.TextEdit
	vol     int32
	controlTreeState nk.CollapseStates
	clientTreeState nk.CollapseStates
	hostTreeState nk.CollapseStates
}

func controlsMain(win *glfw.Window, ctx *nk.Context, state *UiState, mgl *mpv.MpvGL) {
	nk.NkPlatformNewFrame()

	// Layout
	bounds := nk.NkRect(50, 50, 300, 500)
	update := nk.NkBegin(ctx, "Controls", bounds,
		nk.WindowBorder|nk.WindowMovable|nk.WindowScalable|nk.WindowMinimizable|nk.WindowTitle)

	if update > 0 {
		if nk.NkTreeStatePush(ctx, nk.TreeNode, "Video Controls", &state.controlTreeState) > 0 {
			nk.NkLayoutRowDynamic(ctx, 30, 1)
			{
				pauseText := ""
				switch pauseTextState {
				case 0:
					pauseText = "Pause"
				case 1:
					pauseText = "Play"
				}
				if nk.NkButtonLabel(ctx, pauseText) > 0 {
					pause()
				}
			}
			nk.NkLayoutRowDynamic(ctx, 30, 2)
			{
				if nk.NkButtonLabel(ctx, "Prev") > 0 {
					log.Println("Prev Pressed")
					prev()
				}
				if nk.NkButtonLabel(ctx, "Next") > 0 {
					log.Println("Next Pressed")
					next()
				}
			}
			nk.NkLayoutRowDynamic(ctx, 30, 1)
			{
				nk.NkEditBuffer(ctx, nk.EditField, &state.board, nk.NkFilterDefault)
				nk.NkEditBuffer(ctx, nk.EditField, &state.thread, nk.NkFilterDefault)
				if nk.NkButtonLabel(ctx, "Load thread") > 0 {
					log.Printf("Requested thread %s from baord %s", state.thread.GetGoString(), state.board.GetGoString())
					go loadThread(state.thread.GetGoString(), state.board.GetGoString())
				}
			}
			nk.NkLayoutRowDynamic(ctx, 25, 1)
			{
				nk.NkPropertyInt(ctx, "Volume", 0, &state.vol, 100, 10, 1)
			}

			nk.NkLayoutRowDynamic(ctx, 25, 1)
			{
				urlSplits := strings.Split(media[mediaPos], "/")
				fn := urlSplits[len(urlSplits)-1]
				label := fmt.Sprintf("Playing: %s", fn)
				nk.NkLabel(ctx, label, nk.TextAlignLeft)
			}
			//Do not forget to end tree
			nk.NkTreePop(ctx)
		}
		if host == true {
			if nk.NkTreeStatePush(ctx, nk.TreeNode, "Host Controls", &state.hostTreeState) > 0 {
				nk.NkLayoutRowDynamic(ctx, 30, 1)
				{
					nk.NkLabel(ctx, "hosting", nk.TextAlignCentered)
				}
				nk.NkLayoutRowDynamic(ctx, 30, 1)
				{
					text := fmt.Sprintf("Client count: %d", len(clients))
					nk.NkLabel(ctx, text, nk.TextAlignCentered)
				}

				//Do not forget to end tree
				nk.NkTreePop(ctx)
			}
		} else {
			if nk.NkTreeStatePush(ctx, nk.TreeNode, "Client Controls", &state.clientTreeState) > 0 {
				nk.NkLayoutRowDynamic(ctx, 30, 1)
				{
					nk.NkEditBuffer(ctx, nk.EditField, &state.connIp, nk.NkFilterDefault)
					if nk.NkButtonLabel(ctx, "Connect") > 0 {
						go clientConnect(state.connIp.GetGoString())
					}
				}
				//Do not forget to end tree
				nk.NkTreePop(ctx)
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

func shareMain(win *glfw.Window, ctx *nk.Context) {
	nk.NkPlatformNewFrame()

	// Layout
	bounds := nk.NkRect(350, 50, 230, 250)
	update := nk.NkBegin(ctx, "Share", bounds,
		nk.WindowBorder|nk.WindowMovable|nk.WindowScalable|nk.WindowMinimizable|nk.WindowTitle)

	if update > 0 {
		nk.NkLayoutRowDynamic(ctx, 30, 2)
		{
			if nk.NkButtonLabel(ctx, "Test") > 0 {
				log.Println("Test Pressed")
			}
		}
	}
	nk.NkEnd(ctx)
	nk.NkPlatformRender(nk.AntiAliasingOn, maxVertexBuffer, maxElementBuffer)
}

func onError(code int32, msg string) {
	log.Printf("[glfw ERR]: error %d: %s", code, msg)
}

func bitFlag(v bool) int32 {
	if v {
		return 1
	}
	return 0
}
