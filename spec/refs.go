package spec

import (
	"errors"
	"fmt"
)

var (
	ErrUnresolvedRef = errors.New("found unresolved ref")
)

func foundUnresolvedRef(ref string) error {
	return fmt.Errorf("%q is not resolved: %w", ref, ErrUnresolvedRef)
}

type ChannelRef = RefG[*Channel]
type MessageRef = RefG[*Message]
type ParameterRef = RefG[*Parameter]
type MessageTraitRef = RefG[*MessageTrait]
type OperationTraitRef = RefG[*OperationTrait]
type OperationRef = RefG[*Operation]
type CorrelationIDRef = RefG[*CorrelationID]
type ServerBindingsRef = RefG[*ServerBindings]
type ChannelBindingsRef = RefG[*ChannelBindings]
type OperationBindingsRef = RefG[*OperationBindings]
type MessageBindingsRef = RefG[*MessageBindings]
