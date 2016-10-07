package main

import (
	"net/http"
	"github.com/dghubble/sling"
	"strings"
)

// Response to a unique hit request.
type Hit struct {
	Data *HitResource `json:"data"`
}

type Hits struct {
	Data []HitResource `json:"data"`
}

type HitResource struct {
	Resource
	Attributes *HitAttributes `json:"attributes"`
}

type HitAttributes struct {
	Limb string `json:"limb"`
	Normal *Vector `json:"normal"`
	Type string `json:"type"`
	Impact *Vector `json:"impact"`
	Direct bool `json:"direct"`
	Cause string `json:"cause"`
	Velocity *Vector `json:"velocity"`
	Radius float64 `json:"radius"`

	TargetTag string `json:"targetTag"`
	TargetSide string `json:"targetSide"`
	TargetPosition *Vector `json:"targetPosition"`
	TargetAim *Vector `json:"targetAim"`
	TargetStance string `json:"targetStance"`

	ShooterTag string `json:"shooterTag"`
	ShooterSide string `json:"shooterSide"`
	ShooterPosition *Vector `json:"shooterPosition"`
	ShooterAim *Vector `json:"shooterAim"`
	ShooterStance string `json:"shooterStance"`
}

func NewHit(attributes *HitAttributes) (*Hit) {
	return &Hit {
		Data: &HitResource {
			Attributes: attributes,
		},
	}
}

type HitService struct {
	sling *sling.Sling
}

func newHitService(sling *sling.Sling) *HitService {
	return &HitService {
		sling: sling.Path("hits"),
	}
}

type HitListParams struct { }

func (h *HitService) List(ids []string, params *HitListParams) (*Hits, *http.Response, error) {
	hits := new(Hits)
	serviceError := new(ServiceError)

	resp, err := h.sling.New().Get("/" + strings.Join(ids, ",")).Receive(hits, serviceError)
	return hits, resp, err
}

func (h *HitService) ListAll(params *HitListParams) (*Hits, *http.Response, error) {
	hits := new(Hits)
	serviceError := new(ServiceError)

	resp, err := h.sling.New().Get("").Receive(hits, serviceError)
	return hits, resp, err
}

func (h *HitService) Create(attributes *HitAttributes) (*Hit, *http.Response, error) {
	hit := new(Hit)
	serviceError := new(ServiceError)

	resp, err := h.sling.New().Post("").BodyJSON(&Hit {
		Data: &HitResource {
			Attributes: attributes,
		},
	}).Receive(hit, serviceError)

	return hit, resp, err
}

func (h *HitService) Read(id string) (*Hit, *http.Response, error) {
	hit := new(Hit)
	serviceError := new(ServiceError)

	resp, err := h.sling.New().Get("/" + id).Receive(hit, serviceError)
	return hit, resp, err
}
