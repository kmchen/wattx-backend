package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	pb "github.com/wattx-backend/proto"
)

type TestSuite struct {
	suite.Suite
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupSuite() {
}

func (ts *TestSuite) TestToProtoConversion() {
	var conversions = Conversion{
		Data: map[string]Currency{
			"EOS10": Currency{
				Symbol: "EOS10",
				Quote: map[string]Price{
					"USD": {
						Price: 10,
					},
				},
			},
		},
	}
	pbConversion := ToProtoConversion(conversions)
	assetValue := &pb.AssetValue{
		Key:   "EOS10",
		Value: 10,
	}

	pbConversionExpect := pb.Data{Data: []*pb.AssetValue{assetValue}}
	assert.Equal(ts.T(), pbConversion, pbConversionExpect)
}
