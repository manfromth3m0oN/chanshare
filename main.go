package main

import (
	"C"
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/xlab/closer"
)
import "github.com/yourok/go-mpv/mpv"

const (
	winWidth  = 400
	winHeight = 500

	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
)

var m *mpv.Mpv
var media []string
var mediaPos int
var requesting bool
var host bool

func init() {
	runtime.LockOSThread()
}


func main() {
	flag.BoolVar(&host, "h", false, "Set your chanshare instance as host")
	flag.Parse()
	log.Printf("Hosting mode is set to: %v", host)
	if host == true {
		go hostFunc()
	}

	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "chanshare", nil, nil)
	if err != nil {
		closer.Fatalln(err)
	}
	win.MakeContextCurrent()

	width, height := win.GetSize()
	log.Printf("glfw: created window %dx%d", width, height)

	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	m = mpv.Create()
	defer m.TerminateDestroy()
	err = m.SetOption("loop-playlist", mpv.FORMAT_FLAG, true)
	if err != nil {
		log.Fatalf("Failed to set loop: %v", err)
	}
	err = m.RequestLogMessages("trace")
	if err != nil {
		log.Fatalf("Unable to get mpv logs: %s", err)
	}

	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	sansFont := nk.NkFontAtlasAddDefault(atlas, 13, nil)
	nk.NkFontStashEnd()
	if sansFont != nil {
		nk.NkStyleSetFont(ctx, sansFont.Handle())
	}

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	state := &State{
		bgColor: nk.NkRgba(28, 48, 62, 255),
		vol: 35,
	}
	nk.NkTexteditInitDefault(&state.board)
	nk.NkTexteditInitDefault(&state.thread)

	//err = m.SetOptionString("vo", "opengl-cb")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	err = m.Initialize()
	if err != nil {
		log.Fatalf("Mpv Init: %v", err)
	}

	mgl := m.GetSubApiGL()
	if mgl == nil {
		return
	}

	err = mgl.InitGL()
	if err != nil {
		log.Printf("InitGL Error: %v", err)
	}
	defer mgl.UninitGL()


	fpsTicker := time.NewTicker(time.Second / 15)
	for {
		if requesting == false {
		select {
		case <-exitC:
			nk.NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			err := m.SetOption("volume", mpv.FORMAT_INT64, int(state.vol))
			if err != nil {
				log.Fatalf("Failed to change volume: %v", err)
			}
			glfw.PollEvents()
			controlsMain(win, ctx, state, mgl)
			win.SwapBuffers()
		}
	}
	}
}
