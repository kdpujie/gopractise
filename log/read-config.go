package main

import (
	"flag"
	"log"

	"github.com/robfig/config"
)

var (
	configFile = flag.String("configFile", "config.ini", "General configuration file")
)

func main() {
	readConfigSection()
}

func readConfigSection() {
	flag.Parse()
	cfg, err := config.ReadDefault(*configFile)

	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	for i, section := range cfg.Sections() {
		if section == config.DEFAULT_SECTION {
			continue
		}
		log.Printf("第%d个section，value=%s \n", i, section)
		readSecionOptions(cfg, section)
	}
}

func readSecionOptions(cfg *config.Config, section string) {
	if options, err := cfg.SectionOptions(section); err == nil {
		for i, option := range options {
			value, _ := cfg.String(section, option)
			log.Printf("\t 第%d个option，key=%s \t value=%s \n", i, option, value)
		}
	}
}
