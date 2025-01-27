package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(testdata.RandomString(32))
	require.NoError(t, err)

	userID := testdata.RandomInt(0, 1000)

	testCases := []struct {
		name      string
		userID    int64
		duration  time.Duration
		genErr    error
		verifyErr error
	}{
		{
			name:      "valid case",
			userID:    userID,
			duration:  time.Minute,
			genErr:    nil,
			verifyErr: nil,
		},
		{
			name:      "expired token",
			userID:    userID,
			duration:  -time.Minute,
			genErr:    nil,
			verifyErr: jwt.ErrTokenExpired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := maker.Generate(tc.userID, tc.duration)
			if tc.genErr == nil {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			} else {
				require.EqualError(t, err, tc.genErr.Error())
				require.Empty(t, token)
			}
			c, err := maker.Verify(token)
			if tc.verifyErr == nil {
				require.NoError(t, err)
				require.NotEmpty(t, c.ID)
				require.Equal(t, userID, c.UserID)
			} else {
				require.Contains(t, err.Error(), tc.verifyErr.Error())
				require.Nil(t, c)
			}
		})
	}
}
