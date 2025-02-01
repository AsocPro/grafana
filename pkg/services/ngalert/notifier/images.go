package notifier

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	alertingImages "github.com/grafana/alerting/images"
	alertingNotify "github.com/grafana/alerting/notify"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/services/ngalert/store"
)

// storedTokenImageProvider is an implementation of the ImageProvider interface that retrieves images from a database
// using tokens.
// Safety from Path Traversal is provided by the fact that the paths are stored in the database and not user-controlled.
type storedTokenImageProvider struct {
	store  store.ImageStore
	logger log.Logger
}

func newImageProvider(store store.ImageStore, logger log.Logger) alertingImages.Provider {
	return &storedTokenImageProvider{
		store:  store,
		logger: logger,
	}
}

func (i storedTokenImageProvider) GetImage(ctx context.Context, alert alertingNotify.Alert) (*alertingImages.Image, error) {
	uri := alertingImages.GetImageURI(alert)
	if uri == nil {
		return nil, nil
	}

	// TODO: If the URI has a URL, we could consider returning it directly instead of hitting the database right away we could
	//   	postpone the database hit until the raw image data is actually requested in Image.RawData().

	// Assume the uri is a token because we used to store tokens as plain strings.
	logger := i.logger.New("token", uri.Token)
	logger.Debug("Received an image token in annotations")
	image, err := i.store.GetImage(ctx, uri.Token)
	if err != nil {
		if errors.Is(err, models.ErrImageNotFound) {
			logger.Info("Image not found in database")
			return nil, nil
		}
		return nil, err
	}

	name := image.Token
	if image.Path != "" {
		// Most (all?) images from Grafana have a Path, so this keeps the filetype suffix in the name for uploads.
		name = filepath.Base(image.Path)
	}

	return &alertingImages.Image{
		Name: name,
		URL:  image.URL,
		RawData: func(_ context.Context) ([]byte, error) {
			if image.Path == "" {
				return nil, models.ErrImageDataUnavailable
			}
			return readImage(image.Path, logger)
		},
	}, nil
}

// readImage returns an image from the given path.
func readImage(path string, logger log.Logger) ([]byte, error) {
	fp := filepath.Clean(path)
	_, err := os.Stat(fp)
	if os.IsNotExist(err) || os.IsPermission(err) {
		return nil, models.ErrImageNotFound
	}

	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			logger.Error("Failed to close image file", "error", err)
		}
	}()

	return io.ReadAll(f)
}
