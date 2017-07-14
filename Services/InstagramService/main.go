package main

import (
	"flag"

	"./models"
)

var cfg models.Configuration

func init() {
	pathPtr := flag.String("config", defaultConfigPath, "Path for configuration file")
	flag.Parse()

	if *pathPtr == "" {
		panic("No config path")
	}

	bytes, err := ioutil.ReadFile(*pathPtr)

	if err != nil {
		panic("Read config file error")
	}

	if *pathPtr == defaultConfigPath {
		currentDir, _ := os.Getwd()
		cfg = models.Config{FilePath: currentDir + string(os.PathSeparator) + "cfg.json"}
	} else {
		cfg = models.Config{FilePath: *pathPtr}
	}

	// set default values
	cfg.UpdateMinutes = 30

	json.Unmarshal(bytes, &conf)

	if len(cfg.Driver) == 0 || len(cfg.ConnectionString) == 0 {
		panic("DB connection configuration failed")
	}
}

func main() {
	// get users with instagram

}
