package go_watch

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"net/http"
	"sync"
	"log"
	"net/smtp"
	"strconv"
	"time"
)

/**
*  Author: Younis Shah
**/

// WaitGroup for waiting for a group of goroutines
// else the program will exit a soon as it's run.
var wg sync.WaitGroup

func Watch() {

	// Read the config.json file
	config := getConfig()

	// Time ticker - Thanks Sherlock!
	var ticker *time.Ticker

	switch config.Every.Unit {
	case `H`:
		// Make an hourly ticker
		ticker = time.NewTicker(time.Duration(config.Every.Duration) * time.Hour)
		break;
	case `M` :
		// Make a minute-wise ticker
		ticker = time.NewTicker(time.Duration(config.Every.Duration) * time.Minute)
	default:
		// Make a second-wise ticker
		ticker = time.NewTicker(time.Duration(config.Every.Duration) * time.Second)
	}

	for {
		_watch(config)
		<-ticker.C
	}
}

func getConfig() Config {
	pwd, _ := os.Getwd()
	// config.json as byte array
	_json, err := ioutil.ReadFile(pwd + "/config.json")
	if (err != nil) {
		panic(err)
	}
	// Config struct
	var config Config
	// map JSON data to Config struct
	err = json.Unmarshal(_json, &config)
	if (err != nil) {
		panic(err)
	}
	return config
}

func _watch(config Config) {

	// Add delta to the WaitGroup counter equal to the number of sites
	// in the config.json file + 1 for sending email goroutine
	wg.Add(len(config.Sites) + 1)
	for _, v := range config.Sites {
		go checkSite(v.URL, v.Port, config.EmailSettings)
	}
	// Wait for all the goroutines to finish!
	wg.Wait()
}

// Check whether the site is up using a simple HTTP GET
// If the site is not responding, send an email!
// Note: This function runs in a goroutine!
// TODO - find some more efficient method to check the status of the website!
func checkSite(url string, port int, emailSettings NotificationEmailSettings) {
	defer wg.Done()
	if (port != 80) {
		url += ":" + string(port)
	}
	_, err := http.Get(url)
	if (err != nil) {
		go sendEmail(err.Error(), url, emailSettings)
	} else {
		log.Println("ALL WELL :)")
	}
}

// Send email
// Gmail supported so far!
// Haven't tested with other email providers! TODO
func sendEmail(downErr, site string, emailSetting NotificationEmailSettings) {

	defer wg.Done()

	log.Print("Sending email....")

	body := "From: " + emailSetting.Username + "\n" +
		"To: " + emailSetting.Username + "\n" +
		"Subject: " + site + " down \n\n" + downErr

	auth := smtp.PlainAuth("", emailSetting.Username, emailSetting.Password, emailSetting.Smtp)
	err := smtp.SendMail(
		emailSetting.Smtp + ":" + strconv.Itoa(emailSetting.Port),
		auth, emailSetting.Username,
		[]string{emailSetting.Username},
		[]byte(body))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sent!")
}