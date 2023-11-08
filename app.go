package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"cloud.google.com/go/storage"
)

const defaultPage = "index.html"

// cleanPath returns the canonical path for p, eliminating . and .. elements.
//
// Stolen in part from net/http/server.go.
func cleanPath(p string) string {
	if p == "/" || p == "" {
		return defaultPage
	}

	// Add / so .. cannot go beyond this root.
	// Such that /../../bar = /bar
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)

	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' {
		if np == "/" {
			np = "/" + defaultPage
		} else {
			np += "/" + defaultPage
		}
	}

	// Strip the prefix of /.
	return np[1:]
}

type gcsProxy struct {
	bucket *storage.BucketHandle
}

func (g *gcsProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := cleanPath(req.URL.Path)
	obj := g.bucket.Object(path)
	r, err := obj.NewReader(req.Context())
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			http.NotFound(w, req)
		} else {
			http.Error(w, fmt.Sprintf("500 internal server error: encountered bucket error: %v", err), http.StatusInternalServerError)
		}
		return
	}
	defer r.Close()

	if _, err := io.Copy(w, r); err != nil {
		http.Error(w, fmt.Sprintf("500 internal server error: encountered bucket error: %v", err), http.StatusInternalServerError)
		return
	}
}

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bkt := client.Bucket(os.Getenv("BUCKET"))

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: &gcsProxy{bucket: bkt},
	}
	log.Fatal(s.ListenAndServe())
}
