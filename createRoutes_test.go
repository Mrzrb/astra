package astra

import (
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestService_CreateRoutes(t *testing.T) {
	t.Run("creates routes", func(t *testing.T) {
		var service1, service2 bool

		service := &Service{
			Inputs: []Input{
				{
					CreateRoutes: func(service *Service) error {
						service1 = true
						return nil
					},
				},
				{
					CreateRoutes: func(service *Service) error {
						service2 = true
						return nil
					},
				},
			},
		}

		err := service.CreateRoutes()
		require.NoError(t, err)

		require.True(t, service1)
		require.True(t, service2)
	})

	t.Run("caches service if enabled", func(t *testing.T) {
		t.Run("enabled", func(t *testing.T) {
			service := &Service{
				CacheEnabled: true,
			}

			err := service.Setup()
			require.NoError(t, err)
			defer func() {
				err = service.ClearCache()
				require.NoError(t, err)
			}()

			err = service.CreateRoutes()
			require.NoError(t, err)

			stat, err := os.Stat(path.Join(service.getAstraDirPath(), cacheFileName))
			require.NoError(t, err)
			require.NotNil(t, stat)
		})

		t.Run("disabled", func(t *testing.T) {
			service := &Service{
				CacheEnabled: false,
			}

			err := service.Setup()
			require.NoError(t, err)

			err = service.CreateRoutes()
			require.NoError(t, err)

			_, err = os.Stat(path.Join(service.getAstraDirPath(), cacheFileName))
			require.Error(t, err)
		})
	})
}
