package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	unsplashAccessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
	width := os.Getenv("SWAY_WIDTH")
	height := os.Getenv("SWAY_HEIGHT")
	query := os.Getenv("UNSPLASH_SEARCH")
	unsplashURL := "https://api.unsplash.com/photos/random?w=" + width + "&h=" + height + "&auto=compress&fm=png&client_id=" + unsplashAccessKey + "&query=" + query
	wallpaperPath := "/tmp/swaywallpaper.png"
	// Fetch and save the wallpaper
	if err := fetchAndSaveWallpaper(unsplashURL, wallpaperPath); err != nil {
		log.Fatalf("Failed to fetch and save wallpaper: %v", err)
	}

	// Set the wallpaper on SwayWM
	if err := setSwayWallpaper(wallpaperPath); err != nil {
		log.Fatalf("Failed to set wallpaper on SwayWM: %v", err)
	}
}

func fetchAndSaveWallpaper(url, filepath string) error {
	resp, err := http.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return io.ErrUnexpectedEOF
	}

	var result struct {
		URLs struct {
			Full string `json:"full"`
		} `json:"urls"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	checkErr(err)

	imageResp, err := http.Get(result.URLs.Full)
	checkErr(err)
	defer imageResp.Body.Close()

	file, err := os.Create(filepath)
	checkErr(err)
	defer file.Close()

	_, err = io.Copy(file, imageResp.Body)
	checkErr(err)
	return err
}

func setSwayWallpaper(filepath string) error {
	cmd := exec.Command("/usr/bin/swaymsg", "output", "*", "bg", filepath, "fill")
	return cmd.Run()
}
