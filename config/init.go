package config

func (c *Config) Init(screenWidth, screenHeight int) {
	c.ScreenWidth = screenWidth
	c.ScreenHeight = screenHeight
	
	if screenWidth < 0 {
		c.ScreenWidth = DEFAULT_SCREEN_WIDTH
	}
	if screenHeight < 0 {
		c.ScreenHeight = DEFAULT_SCREEN_HEIGHT
	}
}
