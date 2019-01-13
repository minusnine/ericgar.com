// Download photos from a Google Photos album.
package main

import (
	"context"
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
	googphotos "google.golang.org/api/photoslibrary/v1"
)

var (
	albumName        = flag.String("album_name", "Blog Photos", "The name of the Google Photos album of images to download.")
	maxWidth         = flag.Int("max_width", 582, "The maximum width of an image.")
	clientSecretPath = flag.String("client_secret_path", "/home/eric/src/ericgar.com/scripts/client_secret.json", "Path to the Google API client secret")
)

func main() {
	flag.Parse()
	if len(os.Args) != 2 {
		log.Exit("Requires directory argument")
	}
	photoDir := os.Args[1]

	httpClient, err := oauthHTTPClient(context.Background(), *clientSecretPath)
	if err != nil {
		log.Exit(err)
	}

	client, err := googphotos.New(httpClient)
	if err != nil {
		log.Exitf("Error creating new Google Photos client: %s", err)
	}

	album, err := findAlbum(client, *albumName)
	if err != nil {
		log.Exit(err)
	}

	searchCall := client.MediaItems.Search(&googphotos.SearchMediaItemsRequest{
		AlbumId: album.Id,
	})

	var count int
	searchCall.Pages(context.Background(), func(resp *googphotos.SearchMediaItemsResponse) error {
		for i, p := range resp.MediaItems {
			photoDestPath := filepath.Join(photoDir, strconv.Itoa(i)+".jpg")
			go downloadPhoto(httpClient, photoDestPath, p.BaseUrl)
			count++
		}
		return nil
	})
	if count == 0 {
		log.Exitf("No photos found in album %q", *albumName)
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

	authCodeURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", authCodeURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("error reading auth code: %s", err)
	}
	tok, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("error exchanging oauth token: %s", err)
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

func downloadPhoto(client *http.Client, destPath, photoURL string) {
	photoURL = fmt.Sprintf("%s=w%d-h1024", photoURL, *maxWidth)
	resp, err := client.Get(photoURL)
	if err != nil {
		log.Errorf("Error fetching photo data: %s", err)
		return
	}
	f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0x640)
	if err != nil {
		log.Errorf("Error creating photo file: %s", err)
		return
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Errorf("Error writing photo file: %s", err)
		f.Close()
		return
	}
	if err := f.Close(); err != nil {
		log.Errorf("Error writing photo file: %s", err)
	}
}
