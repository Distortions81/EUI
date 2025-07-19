package main

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var softShadowShader *ebiten.Shader

//go:embed shaders/drop_shadow.kage
var softShadowShaderSrc []byte

func init() {
	var err error
	softShadowShader, err = ebiten.NewShader(softShadowShaderSrc)
	if err != nil {
		panic(err)
	}
}
