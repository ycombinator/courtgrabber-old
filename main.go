package main

import (
	"errors"
	"log"
	"os"

	"github.com/ycombinator/courtgrabber/scheduler"
)

// TODO use cobra or similar to pass in:
//  - username
//  - password
//  - reservation date (default: 3 days from now)
//  - reservation start time
//  - reservation duration (default: 90 minutes)
//  - reservation type, [s]ingles or [d]oubles (default: singles)

func main() {
	cfg, err := config()
	if err != nil {
		log.Fatal(err)
	}

	defer scheduler.Close()

	err = scheduler.Login(cfg.Username, cfg.Password)
	if err != nil {
		log.Fatal(err)
	}

	err = scheduler.SelectDay(cfg.StartTime)
	if err != nil {
		log.Fatal(err)
	}

	court, err := scheduler.FindFreeCourt(cfg.StartTime, cfg.Duration)
	if err != nil {
		log.Fatal(err)
	}

	if court == nil {
		log.Print("free court could not be found")
		return
	}

	err = court.Book()
	if err != nil {
		log.Fatal(err)
	}
}

func config() (*scheduler.Config, error) {
	cfg := scheduler.DefaultConfig()

	username, exists := os.LookupEnv("USERNAME")
	if !exists {
		return nil, errors.New("USERNAME environment variable not specified")
	}
	cfg.Username = username

	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		return nil, errors.New("PASSWORD environment variable not specified")
	}
	cfg.Password = password

	return &cfg, nil
}
