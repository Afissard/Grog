package entity

type Position struct {
	X int32
	Y int32
}

type Entity struct {
	Pos          Position
	Sprite_sheet string
}

var Entity_list = make(map[string]Entity) // global dictionary of all Entity created and alive

func Create_entity(name string, x int32, y int32, sprite_sheet string) {
	/*
		Create an Entity and add it to the global dict : Entity_list
	*/
	position := Position{X: x, Y: y}
	Entity_list[name] = Entity{Pos: position, Sprite_sheet: sprite_sheet}
}
