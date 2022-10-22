package main

import (
	"Api_With_AbdulHamid/config"
	"Api_With_AbdulHamid/domain/Phone"
	"Api_With_AbdulHamid/pkg/id"
	"Api_With_AbdulHamid/repository"
	"Api_With_AbdulHamid/service"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println("config load ", err)
		return
	}

	repo, err := repository.NewPostgres(cfg.PostgresConfig)
	if err != nil {
		log.Println("new postgres ", err)
		return
	}

	PhoneFactory := Phone.NewFactory(id.Generator{})

	svc := service.New(repo, PhoneFactory)
	svr :=

		fmt.Println(repo)

	aaa()

}

func aaa() {
	m := make(map[string][]int)
	slc := make([]int, 0)
	for i := 0; i < 5; i++ {
		slc = append(slc, i)
	}
	m["A"] = slc
	fmt.Println(m["A"])
}
