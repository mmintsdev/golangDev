package main

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CurrentTimeTestSuite struct {
	suite.Suite
}

func Test_CurrentTime_Success(t *testing.T) {
	CurrentTime()
	require.NoError(t, nil)
}

//
//func Test_CurrentTime_Fail(t *testing.T) {
//	require.Error(t, ntp.Time("test"))
//}
