package cli

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	abciServer "github.com/tendermint/tendermint/abci/server"
	tcmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/mock"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
)

var testMbm = module.NewBasicManager(genutil.AppModuleBasic{})

func TestInitCmd(t *testing.T) {
	defer server.SetupViper(t)()
	defer setupClientHome(t)()
	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()

	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)

	cmd := InitCmd(server.NewContext(cfg, log.NewNopLogger()), makeCodec(), testMbm, home)
	_, stdout, _ := tests.ApplyMockIO(cmd)

	require.NoError(t, cmd.RunE(cmd, []string{"appnode-test"}))
	buf, err := ioutil.ReadAll(stdout)
	require.NoError(t, err)

	var jsonMap map[string]interface{}
	require.NoError(t, json.Unmarshal(buf, &jsonMap))
	require.Equal(t, "appnode-test", jsonMap["moniker"])
}

func setupClientHome(t *testing.T) func() {
	clientDir, cleanup := tests.NewTestCaseDir(t)
	viper.Set(flagClientHome, clientDir)
	return cleanup
}

func TestEmptyState(t *testing.T) {
	defer server.SetupViper(t)()
	defer setupClientHome(t)()

	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()

	logger := log.NewNopLogger()
	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)

	ctx := server.NewContext(cfg, logger)
	cdc := makeCodec()


	cmd := InitCmd(ctx, cdc, testMbm, home)
	_, stdout, _ := tests.ApplyMockIO(cmd)
	require.NoError(t, cmd.RunE(cmd, []string{"appnode-test"}))
	var jsonOut map[string]interface{}
	require.NoError(t, json.Unmarshal(stdout.Bytes(), &jsonOut))
	require.Equal(t, "appnode-test", jsonOut["moniker"])

	cmd = server.ExportCmd(ctx, cdc, nil)
	_, s1, _ := tests.ApplyMockIO(cmd)

	require.NoError(t, cmd.RunE(cmd, nil))
	require.NoError(t, json.Unmarshal(s1.Bytes(), &jsonOut), s1.String())
	require.Equal(t, "appnode-test", jsonOut["moniker"])

	// require.Contains(t, out, "genesis_time")
	// require.Contains(t, out, "chain_id")
	// require.Contains(t, out, "consensus_params")
	// require.Contains(t, out, "app_hash")
	// require.Contains(t, out, "app_state")
}

func TestStartStandAlone(t *testing.T) {
	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()
	viper.Set(cli.HomeFlag, home)
	defer setupClientHome(t)()

	logger := log.NewNopLogger()
	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)
	ctx := server.NewContext(cfg, logger)
	cdc := makeCodec()
	initCmd := InitCmd(ctx, cdc, testMbm, home)
	require.NoError(t, initCmd.RunE(nil, []string{"appnode-test"}))

	app, err := mock.NewApp(home, logger)
	require.Nil(t, err)
	svrAddr, _, err := server.FreeTCPAddr()
	require.Nil(t, err)
	svr, err := abciServer.NewServer(svrAddr, "socket", app)
	require.Nil(t, err, "error creating listener")
	svr.SetLogger(logger.With("module", "abci-server"))
	svr.Start()

	timer := time.NewTimer(time.Duration(2) * time.Second)
	select {
	case <-timer.C:
		svr.Stop()
	}
}

func TestInitNodeValidatorFiles(t *testing.T) {
	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()
	viper.Set(cli.HomeFlag, home)
	viper.Set(client.FlagName, "moniker")
	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)
	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(cfg)
	require.Nil(t, err)
	require.NotEqual(t, "", nodeID)
	require.NotEqual(t, 0, len(valPubKey.Bytes()))
}

// custom tx codec
func makeCodec() *codec.Codec {
	var cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
