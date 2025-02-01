package notifier

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	alertingImages "github.com/grafana/alerting/images"
	alertingModels "github.com/grafana/alerting/models"
	alertingNotify "github.com/grafana/alerting/notify"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/services/ngalert/store"
)

func TestGetImage(t *testing.T) {
	testBytes := []byte("some test bytes")
	testPath := generateTestFile(t, testBytes)

	var (
		imageWithoutPath = models.Image{
			Token:     "test-token-no-path",
			URL:       "https://test.com",
			CreatedAt: time.Now().UTC(),
			ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		}
		testImage = models.Image{
			Token:     "test-token",
			URL:       "https://test.com",
			Path:      testPath,
			CreatedAt: time.Now().UTC(),
			ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		}
		testImageMissingFile = models.Image{
			Token:     "test-token-missing-file",
			URL:       "https://test.com",
			Path:      "/tmp/missing/1234asdf.png",
			CreatedAt: time.Now().UTC(),
			ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		}
	)

	fakeImageStore := store.NewFakeImageStore(t, &imageWithoutPath, &testImage, &testImageMissingFile)
	store := newImageProvider(fakeImageStore, log.NewNopLogger())

	tests := []struct {
		name          string
		uri           string
		expImage      *alertingImages.Image
		expRawData    []byte
		expRawDataErr error
	}{
		{
			name: "Given existing raw token, expect image",
			uri:  testImage.Token,
			expImage: &alertingImages.Image{
				Name: filepath.Base(testImage.Path),
				URL:  testImage.URL,
			},
		}, {
			name: "Given existing combined token, expect image",
			uri:  alertingImages.ImageURI{Token: testImage.Token, URL: testImage.URL}.Annotation(),
			expImage: &alertingImages.Image{
				Name: filepath.Base(testImage.Path),
				URL:  testImage.URL,
			},
		}, {
			name: "Given existing combined token with just token, expect image",
			uri:  alertingImages.ImageURI{Token: testImage.Token}.Annotation(),
			expImage: &alertingImages.Image{
				Name: filepath.Base(testImage.Path),
				URL:  testImage.URL,
			},
		}, {
			name:     "Given existing combined token with just url, expect nil",
			uri:      alertingImages.ImageURI{URL: testImage.URL}.Annotation(),
			expImage: nil,
		}, {
			name:     "Given missing raw token, expect nil",
			uri:      "invalid",
			expImage: nil,
		}, {
			name:     "Given missing combined token, expect nil",
			uri:      alertingImages.ImageURI{Token: "invalid"}.Annotation(),
			expImage: nil,
		}, {
			name: "Given image with Path, expect RawData",
			uri:  alertingImages.ImageURI{Token: testImage.Token}.Annotation(),
			expImage: &alertingImages.Image{
				Name: filepath.Base(testImage.Path),
				URL:  testImage.URL,
			},
			expRawData: testBytes,
		}, {
			name: "Given image with Path but file doesn't exist, expect RawData error",
			uri:  alertingImages.ImageURI{Token: testImageMissingFile.Token}.Annotation(),
			expImage: &alertingImages.Image{
				Name: filepath.Base(testImageMissingFile.Path),
				URL:  testImageMissingFile.URL,
			},
			expRawDataErr: models.ErrImageNotFound,
		}, {
			name: "Given image without Path, expect RawData error",
			uri:  alertingImages.ImageURI{Token: imageWithoutPath.Token}.Annotation(),
			expImage: &alertingImages.Image{
				Name: imageWithoutPath.Token,
				URL:  imageWithoutPath.URL,
			},
			expRawDataErr: models.ErrImageDataUnavailable,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			alert := alertingNotify.Alert{
				Alert: model.Alert{
					Annotations: model.LabelSet{alertingModels.ImageTokenAnnotation: model.LabelValue(test.uri)},
				},
			}
			image, err := store.GetImage(context.Background(), alert)
			require.NoError(tt, err)
			if test.expImage == nil {
				require.Nil(tt, image)
				return
			}
			require.Equal(tt, test.expImage.URL, image.URL)
			require.Equal(tt, test.expImage.Name, image.Name)
			if test.expRawData != nil {
				b, err := image.RawData(context.Background())
				require.NoError(tt, err)
				require.Equal(tt, test.expRawData, b)
			}
			if test.expRawDataErr != nil {
				b, err := image.RawData(context.Background())
				require.ErrorIs(tt, err, test.expRawDataErr)
				require.Nil(tt, b)
			}
		})
	}
}

func generateTestFile(t *testing.T, b []byte) string {
	t.Helper()
	f, err := os.CreateTemp("/tmp", "image")
	require.NoError(t, err)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(f.Name()))
	})

	_, err = f.Write(b)
	require.NoError(t, err)

	return f.Name()
}
