package main

import (
	"github.com/dghubble/sling"
	"strings"
	"net/http"
)

type Player struct {
	Data *PlayerResource `json:"data"`
}

type Players struct {
	Data []PlayerResource `json:"data"`
}

type PlayerResource struct {
	Resource
	Attributes *PlayerAttributes `json:"attributes"`
}

type PlayerAttributes struct {
	Name string `json:"name"`
	Handle string `json:"handle"`
	Human bool `json:"human"`
}

type PlayerService struct {
	sling *sling.Sling
}

func newPlayerService(sling *sling.Sling) *PlayerService {
	return &PlayerService {
		sling: sling.Path("players"),
	}
}

type PlayerListParams struct { }

func (p *PlayerService) List(ids []string, params *PlayerListParams) (*Players, *http.Response, error) {
	players := new(Players)
	serviceError := new(ServiceError)

	resp, err := p.sling.New().Get("/" + strings.Join(ids, ",")).Receive(players, serviceError)

	return players, resp, err
}

func (p *PlayerService) ListAll(params *PlayerListParams) (*Players, *http.Response, error) {
	players := new(Players)
	serviceError := new(ServiceError)

	resp, err := p.sling.New().Get("").Receive(players, serviceError)

	return players, resp, err
}

func (p *PlayerService) Create(attributes *PlayerAttributes) (*Player, *http.Response, error) {
	player := new(Player)
	serviceError := new(ServiceError)

	resp, err := p.sling.New().Post("").BodyJSON(Player {
		Data: &PlayerResource {
			Resource: Resource {
				Type: "players",
			},
			Attributes: attributes,
		},
	}).Receive(player, serviceError)

	return player, resp, err
}

func (p *PlayerService) Read(id string) (*Player, *http.Response, error) {
	player := new(Player)
	serviceError := new(ServiceError)

	resp, err := p.sling.New().Get("/" + id).Receive(player, serviceError)
	return player, resp, err
}
