package util

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TimeTestStruct struct {
	Now UnixTimestamp `json:"now"`
}

func TestUnixTimestampMarshalJSON(t *testing.T) {
	now := time.Now()

	expected := fmt.Sprintf("{\"now\":\"%d\"}", now.Unix())

	subject := &TimeTestStruct{
		Now: UnixTimestamp(now),
	}

	marshalled, err := json.Marshal(subject)

	assert.Nil(t, err)
	assert.Equal(t, expected, string(marshalled))
}

func TestUnixTimestampUnmarshalJSON(t *testing.T) {
	var subject TimeTestStruct

	err := json.Unmarshal([]byte("{\"now\":\"1445754159\"}"), &subject)

	assert.Nil(t, err)
	assert.Equal(t, int64(1445754159), time.Time(subject.Now).Unix())
}

func TestUnixTimestampUnmarshalJSON_withInvalid(t *testing.T) {
	var subject TimeTestStruct

	err := json.Unmarshal([]byte("{\"now\":\"\"}"), &subject)

	assert.NotNil(t, err)

	err = json.Unmarshal([]byte(""), &subject)

	assert.NotNil(t, err)
}
