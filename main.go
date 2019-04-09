package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"os/user"
	"encoding/json"
	"io/ioutil"
	"time"
)

type Song struct {
	Title string `json:"title"`
	Artist string `json:"artist"`
	Album string `json:"album"`
	AlbumArtCoverUrl string `json:"albumArt"`
}
type Playback struct {
	Playing bool `json:"playing"`
	Song Song `json:"song"`
}
func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: gpdmp-slack-updater ACCESS_TOKEN")
	}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile := usr.HomeDir + "/.config/Google Play Music Desktop Player/json_store/playback.json"
	api := slack.New(os.Args[1])
	previousTitle := ""
	for {
		raw, err := ioutil.ReadFile(jsonFile)
		if err!= nil {
			log.Fatal(err)
		}
		var c Playback
		json.Unmarshal(raw, &c)
		if previousTitle != c.Song.Title {
			api.SetUserCustomStatus(fmt.Sprintf("%s: %s", c.Song.Artist, c.Song.Title), ":musical_note:")
			previousTitle = c.Song.Title
			log.Println(fmt.Sprintf("Found a new song, updating slack: %s", c.Song.Title))
		}
		time.Sleep(5 * time.Second)
	}
}
