package config

type Config struct{
	ScreenWidth, ScreenHeight int
}

var Global Config

const (
	DEFAULT_SCREEN_WIDTH int = 900
	DEFAULT_SCREEN_HEIGHT int = 600
)