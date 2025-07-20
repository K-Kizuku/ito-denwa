package service

import (
	"context"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/strings/rpc"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/strings/resources"
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/pkg/uuid"
)

type StringItemService struct {
	mu           sync.RWMutex
	strings      map[string]*entity.String
	templates    map[string]*entity.TemplateString
	userStrings  map[string][]string
}

func NewStringItemService() *StringItemService {
	svc := &StringItemService{
		strings:     make(map[string]*entity.String),
		templates:   make(map[string]*entity.TemplateString),
		userStrings: make(map[string][]string),
	}
	
	svc.initializeTemplates()
	return svc
}

func (s *StringItemService) initializeTemplates() {
	templates := []*entity.TemplateString{
		{
			ID:                "template-1",
			Name:              "Short String",
			DefaultDurability: 100,
			Type:              entity.StringTypeA,
		},
		{
			ID:                "template-2",
			Name:              "Medium String",
			DefaultDurability: 500,
			Type:              entity.StringTypeB,
		},
		{
			ID:                "template-3",
			Name:              "Long String",
			DefaultDurability: 1000,
			Type:              entity.StringTypeC,
		},
		{
			ID:                "template-4",
			Name:              "Extra Long String",
			DefaultDurability: 2000,
			Type:              entity.StringTypeD,
		},
	}
	
	for _, template := range templates {
		s.templates[template.ID] = template
	}
}

func (s *StringItemService) GetTemplateStrings(ctx context.Context, req *connect.Request[rpc.GetTemplateStringsRequest]) (*connect.Response[rpc.GetTemplateStringsResponse], error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var templates []*resources.TemplateString
	for _, template := range s.templates {
		templates = append(templates, &resources.TemplateString{
			Id:                template.ID,
			Name:              template.Name,
			DefaultDurability: int32(template.DefaultDurability),
			Type:              resources.StringType(template.Type),
		})
	}
	
	return connect.NewResponse(&rpc.GetTemplateStringsResponse{
		TemplateStrings: templates,
	}), nil
}

func (s *StringItemService) GetStrings(ctx context.Context, req *connect.Request[rpc.GetStringsRequest]) (*connect.Response[rpc.GetStringsResponse], error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var strings []*resources.String
	for _, str := range s.strings {
		strings = append(strings, &resources.String{
			Id:         str.ID,
			Name:       str.TemplateString.Name,
			Length:     int32(str.Length),
			Durability: int32(str.Durability),
			Type:       resources.StringType(str.TemplateString.Type),
		})
	}
	
	return connect.NewResponse(&rpc.GetStringsResponse{
		Strings: strings,
	}), nil
}

func (s *StringItemService) BuyString(ctx context.Context, req *connect.Request[rpc.BuyStringRequest]) (*connect.Response[rpc.BuyStringResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	templateID := req.Msg.Id
	length := int(req.Msg.Length)
	
	template, exists := s.templates[templateID]
	if !exists {
		return connect.NewResponse(&rpc.BuyStringResponse{
			Success: false,
			Message: "Template not found",
		}), nil
	}
	
	stringID := uuid.New(ctx)
	now := time.Now()
	
	str := &entity.String{
		ID:             stringID,
		Length:         length,
		Durability:     template.DefaultDurability,
		CreatedAt:      now,
		UpdatedAt:      now,
		TemplateString: *template,
	}
	
	s.strings[stringID] = str
	
	return connect.NewResponse(&rpc.BuyStringResponse{
		Success: true,
		Message: "String purchased successfully",
		String_: &resources.String{
			Id:         str.ID,
			Name:       str.TemplateString.Name,
			Length:     int32(str.Length),
			Durability: int32(str.Durability),
			Type:       resources.StringType(str.TemplateString.Type),
		},
	}), nil
}