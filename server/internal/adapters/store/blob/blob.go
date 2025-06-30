package blob

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/blob/memblob"
	"gocloud.dev/gcp"
	"golang.org/x/oauth2/google"
)

type Config struct {
	StoragePath string `yaml:"storage_path" default:""`
	StorageSecret string `yaml:"storage_secret" default:""`
}

type IBucketStore interface {
	WriteAll(ctx context.Context, key string, p []byte, opts *blob.WriterOptions) error
	ReadAll(ctx context.Context, key string) ([]byte, error)
	List(opts *blob.ListOptions) *blob.ListIterator
	Delete(ctx context.Context, key string) error
	Close() error
}

func NewStore(ctx context.Context, cfg Config) (IBucketStore, error) {
	
	if strings.TrimSpace(cfg.StoragePath) == "" {
		return memblob.OpenBucket(nil), nil
	}

	var errBadSecretURL = errors.Errorf(`unsupported storage config %s, possible schemes supported: "env:// file:// val://" for example: "val://username:password"`, cfg.StorageSecret)
	var errBadStorageURL = errors.Errorf("unsupported storage config %s", cfg.StoragePath)

	var storageSecretValue []byte
	if cfg.StorageSecret != "" {
		parsedSecretURL, err := url.Parse(cfg.StorageSecret)

		if err != nil {
			return nil, errBadSecretURL
		}

		switch parsedSecretURL.Scheme {
		case "env":
			{
				storageSecretValue = []byte(os.Getenv(parsedSecretURL.Hostname()))
			}
		case "file":
			{
				fileContent, err := os.ReadFile(parsedSecretURL.Path)

				if err != nil {
					return nil, errors.Wrap(err, "failed to read secret content at "+parsedSecretURL.Path)
				}

				storageSecretValue = fileContent
			}
		case "val":
			{
				storageSecretValue = []byte(parsedSecretURL.Hostname())
			}
		default:
			return nil, errBadSecretURL
		}
	}

	parsedStorageURL, err := url.Parse(cfg.StoragePath)

	if err != nil {
		return nil, errors.Wrap(err, errBadStorageURL.Error())
	}

	switch parsedStorageURL.Scheme {
		case "gs":
			if strings.TrimSpace(cfg.StorageSecret) == "" {
				return nil, errors.Errorf("%s secret not configured for fs", cfg.StoragePath)
			}

			creds, err := google.CredentialsFromJSON(ctx, storageSecretValue, "https://www.googleapis.com/auth/cloud-platform")
			
			if err != nil {
				return nil, err
			}
			
			client, err := gcp.NewHTTPClient(
				gcp.DefaultTransport(),
				gcp.CredentialsTokenSource(creds),
			)
			
				if err != nil {
				return nil, err
			}

			gcsBucket, err := gcsblob.OpenBucket(ctx, client, parsedStorageURL.Host, nil)
			
			if err != nil {
				return nil, err
			}

			// create a *blob.Bucket
			prefix := fmt.Sprintf("%s/", strings.Trim(parsedStorageURL.Path, "/\\"))
			
			if strings.TrimSpace(prefix) == "" {
				return gcsBucket, nil
			}

			return blob.PrefixedBucket(gcsBucket, prefix), nil
		
		case "file":
			return fileblob.OpenBucket(parsedStorageURL.Path, &fileblob.Options{
				CreateDir: true,
				Metadata:  fileblob.MetadataDontWrite,
			})

		case "mem":
			return memblob.OpenBucket(nil), nil
	}

	return nil, errBadStorageURL
}
