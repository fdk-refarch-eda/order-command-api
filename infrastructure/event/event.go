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
type SimpleEventBusEmitter struct {
	Topic string
}

// Emit Func
func (emitter SimpleEventBusEmitter) Emit(event interface{}) {
	log.Println(fmt.Sprintf("Emitting %s: %+v to topic: %s", reflect.TypeOf(event), event, emitter.Topic))
	eventBus.Publish(emitter.Topic, event)
}

// SimpleEventBusListener type
type SimpleEventBusListener struct {
	Topic     string
	Processor domain.EventProcessor
}

// Listen func
func (listener SimpleEventBusListener) Listen() {
	eventBus.Subscribe(listener.Topic, listener.Processor.Process)
}
