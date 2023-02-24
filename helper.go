package main

import (
	"encoding/json"
	"log"
	"os"
)

func getCookie() (string, error) {
	file, err := os.ReadFile("bask_data/cookie.json")
	if err != nil {
		return "", err
	}

	var cookieData CookieData
	err = json.Unmarshal(file, &cookieData)
	if err != nil {
		return "", err
	}

	return cookieData.Cookie, nil
}

func cookieExists() bool {
	if _, err := os.Stat("bask_data/cookie.json"); err == nil {
		return true
	} else {
		return false
	}
}

func setCookie(cookie string) {
	var cookieData CookieData
	cookieData.Cookie = cookie
	json, err := json.Marshal(cookieData)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("bask_data/cookie.json", json, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
