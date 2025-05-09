// Code generated by "pipeline"; DO NOT EDIT.
package citlali

import (
	_ "embed"

	"fmt"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/gcs/validation"
	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"slices"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData
var paramKeysValidation = map[action.Action][]string{
	3: {"travel"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Citlali, ValidateParamKeys)
}

func ValidateParamKeys(a action.Action, keys []string) error {
	valid, ok := paramKeysValidation[a]
	if !ok {
		return nil
	}
	for _, v := range keys {
		if !slices.Contains(valid, v) {
			return fmt.Errorf("key %v is invalid for action %v", v, a.String())
		}
	}
	return nil
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.434072,
		0.466627,
		0.499183,
		0.54259,
		0.575145,
		0.607701,
		0.651108,
		0.694515,
		0.737922,
		0.78133,
		0.824737,
		0.868144,
		0.922403,
		0.976662,
		1.030921,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.388136,
		0.417246,
		0.446356,
		0.48517,
		0.51428,
		0.54339,
		0.582204,
		0.621018,
		0.659831,
		0.698645,
		0.737458,
		0.776272,
		0.824789,
		0.873306,
		0.921823,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.537712,
		0.57804,
		0.618369,
		0.67214,
		0.712468,
		0.752797,
		0.806568,
		0.860339,
		0.91411,
		0.967882,
		1.021653,
		1.075424,
		1.142638,
		1.209852,
		1.277066,
	}
	// attack: charge = [3]
	charge = []float64{
		0.992,
		1.0664,
		1.1408,
		1.24,
		1.3144,
		1.3888,
		1.488,
		1.5872,
		1.6864,
		1.7856,
		1.8848,
		1.984,
		2.108,
		2.232,
		2.356,
	}
	// attack: collision = [5]
	collision = []float64{
		0.568288,
		0.614544,
		0.6608,
		0.72688,
		0.773136,
		0.826,
		0.898688,
		0.971376,
		1.044064,
		1.12336,
		1.202656,
		1.281952,
		1.361248,
		1.440544,
		1.51984,
	}
	// attack: highPlunge = [7]
	highPlunge = []float64{
		1.419344,
		1.534872,
		1.6504,
		1.81544,
		1.930968,
		2.063,
		2.244544,
		2.426088,
		2.607632,
		2.80568,
		3.003728,
		3.201776,
		3.399824,
		3.597872,
		3.79592,
	}
	// attack: lowPlunge = [6]
	lowPlunge = []float64{
		1.136335,
		1.228828,
		1.32132,
		1.453452,
		1.545944,
		1.65165,
		1.796995,
		1.94234,
		2.087686,
		2.246244,
		2.404802,
		2.563361,
		2.721919,
		2.880478,
		3.039036,
	}
	// skill: frostfall = [4]
	frostfall = []float64{
		0.17024,
		0.183008,
		0.195776,
		0.2128,
		0.225568,
		0.238336,
		0.25536,
		0.272384,
		0.289408,
		0.306432,
		0.323456,
		0.34048,
		0.36176,
		0.38304,
		0.40432,
	}
	// skill: shieldEM = [1]
	shieldEM = []float64{
		5.76,
		6.192,
		6.624,
		7.2,
		7.632,
		8.064,
		8.64,
		9.216,
		9.792,
		10.368,
		10.944,
		11.52,
		12.24,
		12.96,
		13.68,
	}
	// skill: shieldFlat = [2]
	shieldFlat = []float64{
		1386.6759,
		1525.3628,
		1675.6068,
		1837.4082,
		2010.7667,
		2195.6826,
		2392.1558,
		2600.186,
		2819.7734,
		3050.9182,
		3293.6204,
		3547.8796,
		3813.696,
		4091.0698,
		4380.0005,
	}
	// skill: skill = [0]
	skill = []float64{
		0.7296,
		0.78432,
		0.83904,
		0.912,
		0.96672,
		1.02144,
		1.0944,
		1.16736,
		1.24032,
		1.31328,
		1.38624,
		1.4592,
		1.5504,
		1.6416,
		1.7328,
	}
	// burst: iceStorm = [0]
	iceStorm = []float64{
		5.376,
		5.7792,
		6.1824,
		6.72,
		7.1232,
		7.5264,
		8.064,
		8.6016,
		9.1392,
		9.6768,
		10.2144,
		10.752,
		11.424,
		12.096,
		12.768,
	}
	// burst: spiritVessel = [1]
	spiritVessel = []float64{
		1.344,
		1.4448,
		1.5456,
		1.68,
		1.7808,
		1.8816,
		2.016,
		2.1504,
		2.2848,
		2.4192,
		2.5536,
		2.688,
		2.856,
		3.024,
		3.192,
	}
)
