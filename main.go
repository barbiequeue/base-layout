package main

func main() {
	var err error

	if err = conf.loadFromEnv(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	initLogger(conf.logLevel.String())

	log.Println("Hello base layout")
}
