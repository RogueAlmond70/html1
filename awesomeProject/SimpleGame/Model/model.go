package Model

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	posX     int32
	posY     int32
	velocity int32
	radius   float32
	Draw     bool
	Colour   rl.Color
}

type Enemy struct {
	posX        int32
	posY        int32
	draw        bool  // we don't want dead enemies to be drawn to the screen
	isWrecked   bool  // is the ship blown up or not
	wave        int32 // this will denote which wave the enemy is from
	shouldShoot bool  // should the enemy shoot or not

}

type Tree struct {
	posX int32
	posY int32 //this will always be 50, as the trees will be on the ground, obviously
	draw bool  //whether or not we are drawing the tree
}

type PowerUp struct {
	posX  int32
	posY  int32
	draw  bool
	class string
}
