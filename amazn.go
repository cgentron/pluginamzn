package pluginamzn

// yaegi:tags safe
import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/cgentron/api/iface"
)

var _ iface.ResolverHandler = (*Amazn)(nil)

// Config ...
type Config struct{}

// CreateConfig ...
func CreateConfig() *Config {
	return &Config{}
}

// Amazn ...
type Amazn struct {
	name string
}

// New ...
func New(config *Config, name string) (iface.ResolverHandler, error) {
	return &Amazn{
		name: name,
	}, nil
}

func (a *Amazn) Resolve(context context.Context, req interface{}, resp interface{}) error {
	in := req.(proto.Message)

	// opts := in.ProtoReflect().Descriptor().Options()
	// ext := proto.GetExtension(opts, pb.E_Messages).(*pb.Messages)

	// if ext.GetLambda() != nil {
	// 	return nil
	// }

	fmt.Println(in)

	return nil
}
