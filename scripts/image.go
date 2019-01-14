// Download photos from a Google Photos album.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	log "github.com/golang/glog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/sync/errgroup"
	googphotos "google.golang.org/api/photoslibrary/v1"
)

var (
	albumName        = flag.String("album_name", "Blog Photos", "The name of the Google Photos album of images to download.")
	maxWidth         = flag.Int("max_width", 582, "The maximum width of an image.")
	clientSecretPath = flag.String("client_secret_path", "/home/eric/src/ericgar.com/scripts/client_secret.json", "Path to the Google API client secret")
	oauthTokenPath   = flag.String("oauth_token", "/home/eric/src/ericgar.com/scripts/token", "Path to the Google API client secret")
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Exit("Requires directory argument")
	}
	photoDir := flag.Arg(0)

	httpClient, err := oauthHTTPClient(context.Background(), *clientSecretPath)
	if err != nil {
		log.Exit(err)
	}

	client, err := googphotos.New(httpClient)
	if err != nil {
		log.Exitf("Error creating new Google Photos client: %s", err)
	}

	log.Info("Finding album ", *albumName)
	album, err := findAlbum(client, *albumName)
	if err != nil {
		log.Exit(err)
	}

	log.Info("Getting media items in album ", album.Title)
	searchCall := client.MediaItems.Search(&googphotos.SearchMediaItemsRequest{
		AlbumId: album.Id,
	})

	var count int
	var wg errgroup.Group
	searchCall.Pages(context.Background(), func(resp *googphotos.SearchMediaItemsResponse) error {
		for i, p := range resp.MediaItems {
			photoDestPath := filepath.Join(photoDir, strconv.Itoa(i)+".jpg")
			p := p
			wg.Go(func() error {
				return downloadPhoto(httpClient, photoDestPath, p.BaseUrl)
			})
			count++
		}
		return nil
	})
	if count == 0 {
		log.Exitf("No photos found in album %q", *albumName)
	}
	if err := wg.Wait(); err != nil {
		log.Exit(err)
	}
}

func oauthHTTPClient(ctx context.Context, secretPath string) (*http.Client, error) {
	oauthConfigBuf, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return nil, fmt.Errorf("error reading OAuth credential: %s", err)
	}

	oauthConfig, err := google.ConfigFromJSON(oauthConfigBuf)
	if err != nil {
		return nil, fmt.Errorf("error parsing Auth credential: %s", err)
	}
	oauthConfig.Scopes = []string{"https://www.googleapis.com/auth/photoslibrary.readonly"}

	var tok *oauth2.Token
	if tokenBuf, err := ioutil.ReadFile(*oauthTokenPath); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error opening OAuth token file: %s", err)
		}
		authCodeURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog: %v", authCodeURL)

		var code string
		if _, err := fmt.Scan(&code); err != nil {
			return nil, fmt.Errorf("error reading auth code: %s", err)
		}
		tok, err = oauthConfig.Exchange(ctx, code)
		if err != nil {
			return nil, fmt.Errorf("error exchanging oauth token: %s", err)
		}

		tokJSON, err := json.Marshal(tok)
		if err != nil {
			return nil, fmt.Errorf("error marshalling OAuth token: %s", err)
		}

		if err := ioutil.WriteFile(*oauthTokenPath, []byte(tokJSON), 0640); err != nil {
			log.Warningf("Error storing OAuth token: %s", err)
		}
	} else {
		tok = new(oauth2.Token)
		if err := json.Unmarshal(tokenBuf, tok); err != nil {
			return nil, fmt.Errorf("error unmarshalling OAuth token: %s", err)
		}
	}

	return oauthConfig.Client(ctx, tok), nil
}

func findAlbum(client *googphotos.Service, name string) (*googphotos.Album, error) {
	listCall := client.Albums.List()
	cancellableCtx, cancel := context.WithCancel(context.Background())
	var album *googphotos.Album
	err := listCall.Pages(cancellableCtx, func(resp *googphotos.ListAlbumsResponse) error {
		for _, al := range resp.Albums {
			if al.Title == name {
				album = al
				cancel()
				return nil
			}
		}
		return nil
	})
	if err != nil && err != context.Canceled {
		return nil, fmt.Errorf("error fetching Google Photos albums: %s", err)
	}
	if album == nil {
		return nil, fmt.Errorf("Google Photos album %q not found.", *albumName)
	}
	return album, nil
}

func downloadPhoto(client *http.Client, destPath, photoURL string) error {
	log.Infof("Downloading photo to %s", destPath)
	photoURL = fmt.Sprintf("%s=w%d-h1024", photoURL, *maxWidth)
	resp, err := client.Get(photoURL)
	if err != nil {
		return fmt.Errorf("error fetching photo data: %s", err)
	}
	f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		return fmt.Errorf("error creating photo file: %s", err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		return fmt.Errorf("error writing photo file: %s", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("error writing photo file: %s", err)
	}
	return nil
}
