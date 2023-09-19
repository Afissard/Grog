package main

import (
	"log"

	"Grog/pkg/entity"
	"Grog/pkg/sdlEngine"

	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	Title   string
	Running bool
	Width   int32
	Height  int32
	Window  *sdl.Window
	Surface *sdl.Surface
	Event   *sdl.KeyboardEvent
	err     error
}

func (app App) Check_fatal() {
	if app.err != nil {
		log.Fatal("Fatal error : ", app.err)
		//os.Exit(1)
	}
}

func main() {
	app := App{
		Title:   "Grog",
		Running: true,
		Width:   800, //TODO: make it resizable
		Height:  600,
	}

	app.err = sdl.Init(sdl.INIT_VIDEO)
	app.Window, app.err = sdl.CreateWindow(
		app.Title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		app.Width, app.Height,
		sdl.WINDOW_SHOWN,
	)
	app.Surface, app.err = app.Window.GetSurface()

	app.Check_fatal()
	defer app.Window.Destroy()

	// SERVER Init :
	// GAME Init :
	entity.Create_player() //TODO: remove -> do a proper init for the whole game
	player := entity.Entity_list["Player"]

	// GAME LOOP :
	for app.Running {
		var inputs string = ""
		inputs, app.err = sdlEngine.Get_input()
		app.Check_fatal()
		if inputs != "" {
			if inputs == "Quit" {
				app.Running = false
				break
			} else if string(inputs[0]) == "z" { //TODO: replace const "z" with a var (from a struct) to be able to modify the bindings
				log.Print("up")
				player.Pos.Y -= 5
			} else if string(inputs[0]) == "s" {
				log.Print("down")
				player.Pos.Y += 5
			} else if string(inputs[0]) == "q" {
				log.Print("left")
				player.Pos.X -= 5
			} else if string(inputs[0]) == "d" {
				log.Print("right")
				player.Pos.X += 5
			}
		}

		//HERE: update func for game event
		app.err = app.Surface.FillRect(&sdl.Rect{X: 0, Y: 0, W: app.Width, H: app.Height}, 0)  // "clear" the screen
		app.err = sdlEngine.Draw(app.Surface, player.Sprite_sheet, player.Pos.X, player.Pos.Y) // add parameters ? draw all or what ?
		app.Check_fatal()

		// Update the window surface with what we have drawn
		app.err = app.Window.UpdateSurface()
		sdl.Delay(32)
		app.Check_fatal()
	}

	//app.err = errors.New("Forced quit")

}
