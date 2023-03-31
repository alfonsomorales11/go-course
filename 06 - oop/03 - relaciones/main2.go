package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

// un curso puede tener muchos videos
type Curso struct {
	titulo string
	videos []Video
}

// un video solo puede pertenecer a un curso
type Video struct {
	titulo string
	curso  Curso
}

func main() {

	video1 := Video{titulo: "01 introduccion"}
	video2 := Video{titulo: "02 parte dos"}
	curso1 := Curso{"Matematicas", []Video{video1, video2}}

	video1.curso = curso1

	fmt.Println(video1.curso.titulo)

	for _, videos := range curso1.videos {
		fmt.Println(videos.titulo)
	}
}
