package event

import (
	"fmt"
	"log"
	"reflect"

	"github.com/asaskevich/EventBus"
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
)

var eventBus EventBus.Bus = EventBus.New()

// SimpleEventBusEmitter type
type SimpleEventBusEmitter struct{}

// Emit Func
func (emitter SimpleEventBusEmitter) Emit(event interface{}) {
	log.Println(fmt.Sprintf("Emitting %s: %+v", reflect.TypeOf(event), event))
	eventBus.Publish("order-commands", event)
}

// SimpleEventBusListener type
type SimpleEventBusListener struct {
	Processor domain.EventProcessor
}

// Listen func
func (listener SimpleEventBusListener) Listen() {
	eventBus.Subscribe("order-commands", listener.Processor.Process)
}
