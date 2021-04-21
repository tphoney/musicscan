// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package cli

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/tphoney/musicscan/cli/util"
	"github.com/tphoney/musicscan/types"

	"gopkg.in/alecthomas/kingpin.v2"
)

type scanCommand struct {
	proj int64
}

func (c *scanCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	basePath, err := ioutil.ReadDir("/media/tp/stuff/Music")
	if err != nil {
		return err
	}
	for _, f := range basePath {
		if !f.IsDir() {
			continue
		}

		var foundArtist *types.Artist
		fmt.Println(f.Name())
		artistPath := fmt.Sprintf("/media/tp/stuff/Music/%s", f.Name())
		// try to find the artist
		foundArtist, err = client.ArtistName(c.proj, f.Name())
		if err != nil {
			// artist not found create it
			inArtist := &types.Artist{
				Name:   f.Name(),
				Desc:   artistPath,
				Wanted: true,
			}
			foundArtist, err = client.ArtistCreate(c.proj, inArtist)
			if err != nil {
				fmt.Printf("got here %s", err.Error())
				return err
			}
		}
		albumPaths, _ := ioutil.ReadDir(artistPath)
		for _, albumPath := range albumPaths {
			if albumPath.IsDir() {
				fmt.Println("  " + albumPath.Name())
				_, err = client.AlbumName(c.proj, foundArtist.ID, albumPath.Name())
				if err != nil {
					abs := artistPath + "/" + albumPath.Name()
					mp3Matches, _ := filepath.Glob(abs + "/*.mp3")
					flacMatches, _ := filepath.Glob(abs + "/*.flac")
					format := ""
					if len(mp3Matches) != 0 && len(flacMatches) != 0 {
						format = "mp3+flac"
					} else if len(mp3Matches) != 0 {
						format = "mp3"
					} else if len(flacMatches) != 0 {
						format = "flac"
					}
					inputAlbum := &types.Album{
						Name:   albumPath.Name(),
						Desc:   abs,
						Format: format,
					}
					_, err := client.AlbumCreate(c.proj, foundArtist.ID, inputAlbum)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// helper function to register the scan command.
func registerScan(app *kingpin.Application) {
	c := new(scanCommand)

	cmd := app.Command("scan", "scan directory").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)
}
