package proto

import (
	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/golang/protobuf/proto"
)

// MarshalOrderCommand func
func MarshalOrderCommand(event domain.Event) ([]byte, error) {
	return proto.Marshal(toProto(event))
}

func toProto(event domain.Event) *OrderCommand {
	switch event.(type) {
	case domain.CreateOrderCommand:
		return toProtoCreate(event.(domain.CreateOrderCommand))
	default:
		return nil
	}
}

func toProtoCreate(command domain.CreateOrderCommand) *OrderCommand {
	return &OrderCommand{
		Payload: &OrderCommand_Create{
			Create: &CreateOrderCommand{
				OrderID:              command.OrderID,
				CustomerID:           command.CustomerID,
				ProductID:            command.ProductID,
				Quantity:             uint32(command.Quantity),
				ExpectedDeliveryDate: command.ExpectedDeliveryDate,
				PickupDate:           command.PickupDate,
				PickupAddress:        toProtoAddress(command.PickupAddress),
				DestinationAddress:   toProtoAddress(command.DestinationAddress),
			},
		},
	}
}

func toProtoAddress(address domain.Address) *Address {
	return &Address{
		Street:  address.Street,
		City:    address.City,
		Country: address.Country,
		State:   address.State,
		ZipCode: address.ZipCode,
	}
}

// UnmarshalOrderCommand func
func UnmarshalOrderCommand(data []byte) (domain.Event, error) {
	orderCommand := &OrderCommand{}
	err := proto.Unmarshal(data, orderCommand)

	if err != nil {
		return nil, err
	}

	return fromProto(orderCommand), nil
}

func fromProto(command *OrderCommand) domain.Event {
	switch command.Payload.(type) {
	case *OrderCommand_Create:
		return fromProtoCreate(command.GetCreate())
	default:
		return nil
	}
}

func fromProtoCreate(command *CreateOrderCommand) domain.CreateOrderCommand {
	return domain.CreateOrderCommand{
		OrderID:              command.OrderID,
		CustomerID:           command.CustomerID,
		ProductID:            command.ProductID,
		Quantity:             int(command.Quantity),
		ExpectedDeliveryDate: command.ExpectedDeliveryDate,
		PickupDate:           command.PickupDate,
		PickupAddress:        fromProtoAddress(command.PickupAddress),
		DestinationAddress:   fromProtoAddress(command.DestinationAddress),
	}
}

func fromProtoAddress(address *Address) domain.Address {
	return domain.Address{
		Street:  address.Street,
		City:    address.City,
		Country: address.Country,
		State:   address.State,
		ZipCode: address.ZipCode,
	}
}
