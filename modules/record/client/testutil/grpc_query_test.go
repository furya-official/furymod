package testutil_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/testutil/rest"
	sdk "github.com/cosmos/cosmos-sdk/types"

	recordcli "github.com/furya-official/furymod/modules/record/client/cli"
	recordtestutil "github.com/furya-official/furymod/modules/record/client/testutil"
	recordtypes "github.com/furya-official/furymod/modules/record/types"
	"github.com/irisnet/irismod/simapp"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := simapp.NewConfig()
	cfg.NumValidators = 1

	s.cfg = cfg

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestQueryRecordGRPC() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	// ---------------------------------------------------------------------------

	from := val.Address
	digest := "digest"
	digestAlgo := "digest-algo"
	uri := "https://example.abc"
	meta := "meta data"

	args := []string{
		fmt.Sprintf("--%s=%s", recordcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", recordcli.FlagMeta, meta),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)

	bz, err := recordtestutil.MsgCreateRecordExec(clientCtx, from.String(), digest, digestAlgo, args...)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	recordID := gjson.Get(txResp.RawLog, "0.events.0.attributes.1.value").String()

	// ---------------------------------------------------------------------------

	baseURL := val.APIAddress
	url := fmt.Sprintf("%s/irismod/record/records/%s", baseURL, recordID)

	respType = proto.Message(&recordtypes.QueryRecordResponse{})
	expectedContents := []recordtypes.Content{{
		Digest:     digest,
		DigestAlgo: digestAlgo,
		URI:        uri,
		Meta:       meta,
	}}

	resp, err := rest.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, respType))
	record := respType.(*recordtypes.QueryRecordResponse).Record
	s.Require().Equal(expectedContents, record.Contents)
}
