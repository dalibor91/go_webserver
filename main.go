package main

import (
	"fmt"
	"flag"
	host "./host"
	"./helpers"
	"log"
	"net/http"
	"strconv"
)

func listConfigs(path string) {
	fmt.Printf("Domains found in %s\n", path)
	for _, file := range helpers.GetConfigs(path) {
		fmt.Printf("  - %s\n", file)
	}
}

func main() {

	path 	:= "."
	config 	:= fmt.Sprintf("%s/domains", path)
	static 	:= fmt.Sprintf("%s/static", path)


	sc  := flag.Bool("sc", false, "Show availible config")
	h 	:= flag.Bool("h", false, "Show help")
	r 	:= flag.Bool("run", false, "Start webserver")
	scf := flag.String("config", "", "Config to load")


	flag.Parse()

	if *sc {
		listConfigs(config)
	} else if *scf != "" && !*r {
		fmt.Printf("Print properties for config %s \n", *scf)
		if u := host.LoadFromFile(fmt.Sprintf("%s/%s", config, *scf)); u != nil {
			u.PrintProps()
		}
	} else if *scf != "" && *r {
		fmt.Printf("Run config \n%s\n" , *scf)
		u := host.LoadFromFile(fmt.Sprintf("%s/%s", config, *scf))
		port, err := strconv.ParseInt(u.GetProp("port"), 10, 0)

		if err != nil {
			log.Println(u.GetProp("port"))
			log.Fatal(err)
		}

		if !u.HasProp("port") && port <= 0 {
			log.Fatal("Configuration does not have port or port invalid")
		}

		if !u.HasProp("static") {
			log.Fatal("Static directory not found")
		}

		fmt.Printf("Port    : %s\nStatic  : static/%s\n", u.GetProp("port"), u.GetProp("static"))


		fs := http.FileServer(http.Dir(fmt.Sprintf("%s/%s", static, u.GetProp("static"))))
		http.Handle("/", fs)

		log.Println("Listening...")
		logFatal(http.ListenAndServe(fmt.Sprintf(":%s", u.GetProp("port")), nil))
	} else if *h || true {
		fmt.Println("Runs http server")
		flag.PrintDefaults()
	}
}