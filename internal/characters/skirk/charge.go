package skirk

import (
	"fmt"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/geometry"
)

var (
	chargeFrames        []int
	chargeSkillFrames   []int
	chargeHitmarks      = []int{27, 27 + 8, 27 + 8 + 8}
	chargeSkillHitmarks = []int{28, 28 + 9, 28 + 9 + 9}
	chargeOffsets       = []float64{1, 1.3, 1.3}
)

func init() {
	chargeFrames = frames.InitAbilSlice(43) // CA -> N1
	chargeFrames[action.ActionDash] = 28    // CA -> D
	chargeFrames[action.ActionJump] = 28    // CA -> J
	chargeFrames[action.ActionSwap] = 42    // CA -> Swap
	chargeFrames[action.ActionWalk] = 53    // CA -> Walk

	chargeSkillFrames = frames.InitAbilSlice(54) // CA -> N1
	chargeSkillFrames[action.ActionDash] = 28    // CA -> D
	chargeSkillFrames[action.ActionJump] = 28    // CA -> J
	chargeSkillFrames[action.ActionWalk] = 52    // CA -> Walk
	chargeSkillFrames[action.ActionSwap] = 42    // CA -> Swap
	chargeSkillFrames[action.ActionBurst] = 43   // CA -> Burst
}

func (c *char) ChargeAttack(p map[string]int) (action.Info, error) {
	if c.StatusIsActive(skillKey) {
		return c.ChargeAttackSkill(p)
	}
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		AttackTag:  attacks.AttackTagExtra,
		ICDTag:     attacks.ICDTagNormalAttack,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeSlash,
		Element:    attributes.Physical,
		Durability: 25,
	}

	for i, mult := range charge {
		ai.Mult = mult[c.TalentLvlAttack()]
		ai.Abil = fmt.Sprintf("Charge %v", i)
		c.Core.QueueAttack(
			ai,
			combat.NewCircleHitOnTarget(
				c.Core.Combat.Player(),
				geometry.Point{Y: chargeOffsets[i]},
				2.2,
			),
			chargeHitmarks[i],
			chargeHitmarks[i],
		)
	}

	return action.Info{
		Frames:          frames.NewAbilFunc(chargeFrames),
		AnimationLength: chargeFrames[action.InvalidAction],
		CanQueueAfter:   chargeFrames[action.ActionJump], // earliest cancel
		State:           action.ChargeAttackState,
	}, nil
}

func (c *char) ChargeAttackSkill(p map[string]int) (action.Info, error) {
	ai := combat.AttackInfo{
		ActorIndex:     c.Index,
		AttackTag:      attacks.AttackTagExtra,
		ICDTag:         attacks.ICDTagNormalAttack,
		ICDGroup:       attacks.ICDGroupDefault,
		StrikeType:     attacks.StrikeTypeSlash,
		Element:        attributes.Cryo,
		Durability:     25,
		IgnoreInfusion: true,
	}

	for i, mult := range skillCharge {
		var cb func(a combat.AttackCB)
		if i == 2 {
			cb = c.absorbVoidRiftCB
		} else {
			cb = nil
		}
		ai.Mult = mult[c.TalentLvlSkill()]
		ai.Abil = fmt.Sprintf("Charge (Skill) %v", i)
		c.Core.QueueAttack(
			ai,
			combat.NewCircleHitOnTarget(
				c.Core.Combat.Player(),
				geometry.Point{Y: chargeOffsets[i]},
				2.2,
			),
			chargeSkillHitmarks[i],
			chargeSkillHitmarks[i],
			cb,
		)
	}

	return action.Info{
		Frames:          frames.NewAbilFunc(chargeFrames),
		AnimationLength: chargeFrames[action.InvalidAction],
		CanQueueAfter:   chargeFrames[action.ActionJump], // earliest cancel
		State:           action.ChargeAttackState,
	}, nil
}
