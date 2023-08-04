package client

import (
	"context"
	"fmt"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/util"
)

type O struct {
	API     model.APIType
	Auth    AuthType
	Corpora []model.CorpusType
	Config  util.Config
}

type Base struct {
	api     model.APIType
	auth    AuthType
	corpora map[model.CorpusType]bool
	config  util.Config
}

func New(o O) *Base {
	c := &Base{
		api:     o.API,
		auth:    o.Auth,
		corpora: map[model.CorpusType]bool{},
		config:  o.Config,
	}
	for _, corpus := range o.Corpora {
		c.corpora[corpus] = true
	}
	return c
}

func (c Base) API() model.APIType                 { return c.api }
func (c Base) Auth() AuthType                     { return c.auth }
func (c Base) Corpora() map[model.CorpusType]bool { return c.corpora }
func (c Base) Config() util.Config                { return c.config }

func (c Base) Put(ctx context.Context, d *model.APIData) error {
	if !c.Auth().Check(AuthTypePrivateWrite) {
		return nil
	}

	return fmt.Errorf("unimplemented")
}

func (c Base) List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error) {
	if !c.Auth().Check(AuthTypePublic) {
		return nil, nil
	}

	return nil, fmt.Errorf("unimplemented")
}

func (c Base) Remove(ctx context.Context, id string) error {
	if !c.Auth().Check(AuthTypePrivateWrite) {
		return nil
	}

	return fmt.Errorf("unimplemented")
}
