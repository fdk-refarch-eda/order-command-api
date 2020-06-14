package event

import (
	"fmt"
	"log"
	"reflect"
)

// CommandEventEmitterMock type
type CommandEventEmitterMock struct{}

// Emit Func
func (emitter CommandEventEmitterMock) Emit(event interface{}) {
	log.Println(fmt.Sprintf("Emitting %s: %+v", reflect.TypeOf(event), event))
}
