package dm01

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
)

// HunterFish ...
func HunterFish(c *match.Card) {

	c.Name = "Hunter Fish"
	c.Power = 3000
	c.Civ = civ.Water
	c.Family = family.Fish
	c.ManaCost = 2
	c.ManaRequirement = []string{civ.Water}

	c.Use(fx.Creature, fx.Blocker, fx.CantAttackPlayers, fx.CantAttackCreatures)

}

// Seamine ...
func Seamine(c *match.Card) {

	c.Name = "Seamine"
	c.Power = 4000
	c.Civ = civ.Water
	c.Family = family.Fish
	c.ManaCost = 6
	c.ManaRequirement = []string{civ.Water}

	c.Use(fx.Creature, fx.Blocker)

}