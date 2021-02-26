package pokemon

import (
	"context"
)

type Server struct{}

func (s *Server) GetPokemon(ctx context.Context, in *PokeQuery) (*Pokemon, error) {
	return &Pokemon{
		IndexNo:        120,
		FormType:       1,
		FormName:       "GOのすがた",
		Name:           "けつばん",
		Type1:          "ノーマル",
		Type2:          "",
		Hp:             80,
		Attack:         120,
		Defence:        70,
		SpecialAttack:  130,
		SpecialDefence: 40,
		Speed:          131,
		Ability1:       "ふみん",
		Ability2:       "",
		HiddenAbility:  "なまけ",
	}, nil
}
