package aurora

import (
	"github.com/hcnet/go/protocols/aurora"
	"github.com/hcnet/go/services/aurora/internal/actions"
	"github.com/hcnet/go/services/aurora/internal/paths"
	"github.com/hcnet/go/services/aurora/internal/resourceadapter"
	"github.com/hcnet/go/support/render/hal"
)

// Interface verification
var _ actions.JSONer = (*PathIndexAction)(nil)

// PathIndexAction provides path finding
type PathIndexAction struct {
	Action
	Query   paths.Query
	Records []paths.Path
	Page    hal.BasePage
}

// JSON implements actions.JSON
func (action *PathIndexAction) JSON() error {
	action.Do(
		action.loadQuery,
		action.loadSourceAssets,
		action.loadRecords,
		action.loadPage,
		func() { hal.Render(action.W, action.Page) },
	)
	return action.Err
}

func (action *PathIndexAction) loadQuery() {
	action.Query.DestinationAmount = action.GetPositiveAmount("destination_amount")
	action.Query.DestinationAddress = action.GetAddress("destination_account", actions.RequiredParam)
	action.Query.DestinationAsset = action.GetAsset("destination_")
}

func (action *PathIndexAction) loadSourceAssets() {
	action.Err = action.CoreQ().AssetsForAddress(
		&action.Query.SourceAssets,
		action.GetAddress("source_account"),
	)
}

func (action *PathIndexAction) loadRecords() {
	action.Records, action.Err = action.App.paths.Find(action.Query, action.App.config.MaxPathLength)
}

func (action *PathIndexAction) loadPage() {
	action.Page.Init()
	for _, p := range action.Records {
		var res aurora.Path
		action.Err = resourceadapter.PopulatePath(action.R.Context(), &res, action.Query, p)

		if action.Err != nil {
			return
		}
		action.Page.Add(res)
	}
}
