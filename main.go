package main

import (
	"os"

	"github.com/ahojukka5/gontti/utils"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: ./gontti <username/repository>")
		return
	}
	repo := os.Args[1]
	XRA := os.Getenv("DOCKER_XRA")
	if XRA == "" {
		println("Environment variable DOCKER_XRA not defined, unable to push image to Docker Hub.")
		return
	}
	println("Building image", repo)
	id := utils.BuildImage(repo)
	println("Build id: ", id)
	println("Tagging image ...")
	utils.TagImage(id, repo)
	println("Pushing image to Docker container registry")
	utils.PushImage(repo, XRA)
}
