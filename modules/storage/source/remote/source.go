package remote

import (
	storagetypes "github.com/MonikaCat/canine-chain/v2/x/storage/types"
	"github.com/forbole/juno/v4/node/remote"

	storagesource "github.com/forbole/bdjuno/v4/modules/storage/source"
)

var (
	_ storagesource.Source = &Source{}
)

// Source implements storagesource.Source using a remote node
type Source struct {
	*remote.Source
	querier storagetypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, querier storagetypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// Params implements storagesource.Source
func (s Source) Params(height int64) (storagetypes.Params, error) {
	res, err := s.querier.Params(remote.GetHeightRequestContext(s.Ctx, height), &storagetypes.QueryParamsRequest{})
	if err != nil {
		return storagetypes.Params{}, nil
	}

	return res.Params, nil
}

// Providers implements storagesource.Source
func (s Source) Providers(height int64) ([]storagetypes.Providers, error) {
	res, err := s.querier.ProvidersAll(remote.GetHeightRequestContext(s.Ctx, height), &storagetypes.QueryAllProvidersRequest{})
	if err != nil {
		return []storagetypes.Providers{}, nil
	}

	return res.Providers, nil
}

// Strays implements storagesource.Source
func (s Source) Strays(height int64) ([]storagetypes.Strays, error) {
	res, err := s.querier.StraysAll(remote.GetHeightRequestContext(s.Ctx, height), &storagetypes.QueryAllStraysRequest{})
	if err != nil {
		return []storagetypes.Strays{}, nil
	}

	return res.Strays, nil
}
