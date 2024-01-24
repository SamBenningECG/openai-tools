package looker

import "github.com/Suburban-Street-Trading/openai-tools/openai"

type Looker struct {
	client openai.Client
}

func NewLooker(client openai.Client) Looker {
	return Looker{
		client: client,
	}
}

func (l Looker) LookAtImage() {

}
