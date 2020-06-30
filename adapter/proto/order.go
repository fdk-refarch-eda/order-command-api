package proto

import (
	"errors"

	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/golang/protobuf/proto"
)

// MarshalOrderEvent func
func MarshalOrderEvent(event domain.Event) ([]byte, error) {
	protoEvent, err := toProtoOrder(event)

	if err != nil {
		return nil, err
	}

	return proto.Marshal(protoEvent)
}

func toProtoOrder(event domain.Event) (*Order, error) {
	switch event.(type) {
	case domain.OrderCreatedEvent:
		return toProtoOrderCreatedEvent(event.(domain.OrderCreatedEvent)), nil
	default:
		return nil, errors.New("Received unknown event type. Don't know how to convert to proto")
	}
}

func toProtoOrderCreatedEvent(event domain.OrderCreatedEvent) *Order {
	return &Order{
		Event: &Order_OrderCreatedEvent{
			OrderCreatedEvent: &OrderCreatedEvent{
				OrderId:              event.OrderID,
				CustomerId:           event.CustomerID,
				ProductId:            event.ProductID,
				Quantity:             uint32(event.Quantity),
				ExpectedDeliveryDate: event.ExpectedDeliveryDate,
				PickupDate:           event.PickupDate,
				PickupAddress:        toProtoAddress(event.PickupAddress),
				DestinationAddress:   toProtoAddress(event.DestinationAddress),
			},
		},
	}
}

// UnmarshalOrderEvent func
func UnmarshalOrderEvent(data []byte) (domain.Event, error) {
	order := &Order{}
	err := proto.Unmarshal(data, order)

	if err != nil {
		return nil, err
	}

	return fromProtoOrder(order)
}

func fromProtoOrder(order *Order) (domain.Event, error) {
	switch order.Event.(type) {
	case *Order_OrderCreatedEvent:
		return fromProtoOrderCreatedEvent(order.GetOrderCreatedEvent()), nil
	default:
		return nil, errors.New("Received unknown order event type. Don't know how to convert from proto")
	}
}

func fromProtoOrderCreatedEvent(event *OrderCreatedEvent) domain.OrderCreatedEvent {
	return domain.OrderCreatedEvent{
		OrderID:              event.OrderId,
		CustomerID:           event.CustomerId,
		ProductID:            event.ProductId,
		Quantity:             int(event.Quantity),
		ExpectedDeliveryDate: event.ExpectedDeliveryDate,
		PickupDate:           event.PickupDate,
		PickupAddress:        fromProtoAddress(event.PickupAddress),
		DestinationAddress:   fromProtoAddress(event.DestinationAddress),
	}
}
