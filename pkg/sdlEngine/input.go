package sdlEngine

import (
	"github.com/veandco/go-sdl2/sdl"
)

func Get_input() (string, error) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.QuitEvent:
			//app.Running = false
			return "Quit", nil
		case *sdl.KeyboardEvent:
			keyCode := ev.Keysym.Sym
			keys := ""

			/*
				// Modifier keys
				switch ev.Keysym.Mod {
				case sdl.KMOD_LALT:
					keys += "Left Alt"
				case sdl.KMOD_LCTRL:
					keys += "Left Control"
				case sdl.KMOD_LSHIFT:
					keys += "Left Shift"
				case sdl.KMOD_LGUI:
					keys += "Left Meta or Windows key"
				case sdl.KMOD_RALT:
					keys += "Right Alt"
				case sdl.KMOD_RCTRL:
					keys += "Right Control"
				case sdl.KMOD_RSHIFT:
					keys += "Right Shift"
				case sdl.KMOD_RGUI:
					keys += "Right Meta or Windows key"
				case sdl.KMOD_NUM:
					keys += "Num Lock"
				case sdl.KMOD_CAPS:
					keys += "Caps Lock"
				case sdl.KMOD_MODE:
					keys += "AltGr Key"
				}
			*/

			if keyCode < 10000 {
				if keys != "" {
					keys += " + "
				}

				// If the key is held down, this will fire
				if ev.Repeat > 0 {
					keys += string(keyCode) + " repeating"
				} else {
					if ev.State == sdl.RELEASED {
						keys += string(keyCode) + " released"
					} else if ev.State == sdl.PRESSED {
						keys += string(keyCode) + " pressed"
					}
				}

			}

			if keys != "" {
				//log.Print(keys)
				return keys, nil
			}
		}
	}
	return "", nil // if there is no input
}
