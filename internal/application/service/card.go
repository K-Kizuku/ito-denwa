package service

import (
	"context"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/cards/rpc"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/cards/resources"
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/pkg/uuid"
)

type CardService struct {
	mu           sync.RWMutex
	cards        map[string]*entity.TelephoneCard
	templates    map[string]*entity.TemplateTelephoneCard
	userCards    map[string][]string
}

func NewCardService() *CardService {
	svc := &CardService{
		cards:     make(map[string]*entity.TelephoneCard),
		templates: make(map[string]*entity.TemplateTelephoneCard),
		userCards: make(map[string][]string),
	}
	
	svc.initializeTemplates()
	return svc
}

func (s *CardService) initializeTemplates() {
	templates := []*entity.TemplateTelephoneCard{
		{
			ID:             "template-1",
			Name:           "Basic Card",
			DefaultBalance: 1000,
			Type:           entity.CardTypeA,
		},
		{
			ID:             "template-2",
			Name:           "Premium Card",
			DefaultBalance: 5000,
			Type:           entity.CardTypeB,
		},
		{
			ID:             "template-3",
			Name:           "Gold Card",
			DefaultBalance: 10000,
			Type:           entity.CardTypeC,
		},
	}
	
	for _, template := range templates {
		s.templates[template.ID] = template
	}
}

func (s *CardService) GetTemplateCards(ctx context.Context, req *connect.Request[rpc.GetTemplateCardsRequest]) (*connect.Response[rpc.GetTemplateCardsResponse], error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var templates []*resources.TemplateCard
	for _, template := range s.templates {
		templates = append(templates, &resources.TemplateCard{
			Id:     template.ID,
			Name:   template.Name,
			Credit: int32(template.DefaultBalance),
			Type:   resources.CardType(template.Type),
		})
	}
	
	return connect.NewResponse(&rpc.GetTemplateCardsResponse{
		TemplateStrings: templates,
	}), nil
}

func (s *CardService) GetCards(ctx context.Context, req *connect.Request[rpc.GetCardsRequest]) (*connect.Response[rpc.GetCardsResponse], error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var cards []*resources.Card
	for _, card := range s.cards {
		cards = append(cards, &resources.Card{
			Id:     card.ID,
			Name:   card.TemplateTelephoneCard.Name,
			Credit: int32(card.Balance),
			Type:   resources.CardType(card.TemplateTelephoneCard.Type),
		})
	}
	
	return connect.NewResponse(&rpc.GetCardsResponse{
		Cards: cards,
	}), nil
}

func (s *CardService) BuyCard(ctx context.Context, req *connect.Request[rpc.BuyCardRequest]) (*connect.Response[rpc.BuyCardResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	name := req.Msg.Name
	credit := req.Msg.Credit
	
	// Find template by name (simplified)
	var template *entity.TemplateTelephoneCard
	for _, t := range s.templates {
		if t.Name == name {
			template = t
			break
		}
	}
	
	if template == nil {
		return connect.NewResponse(&rpc.BuyCardResponse{
			Success: false,
			Message: "Template not found",
		}), nil
	}
	
	cardID := uuid.New(ctx)
	now := time.Now()
	
	card := &entity.TelephoneCard{
		ID:                    cardID,
		Balance:               int(credit),
		History:               []int{},
		CreatedAt:             now,
		UpdatedAt:             now,
		TemplateTelephoneCard: *template,
	}
	
	s.cards[cardID] = card
	
	return connect.NewResponse(&rpc.BuyCardResponse{
		Success: true,
		Message: "Card purchased successfully",
	}), nil
}

