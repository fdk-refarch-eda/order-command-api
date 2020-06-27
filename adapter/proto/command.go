package proto

import (
	"errors"

	"github.com/fdk-refarch-eda/order-service/order-command-service/domain"
	"github.com/golang/protobuf/proto"
)

// MarshalOrderCommand func
func MarshalOrderCommand(command domain.Event) ([]byte, error) {
	protoCommand, err := toProtoCommand(command)

	if err != nil {
		return nil, err
	}

	return proto.Marshal(protoCommand)
}

func toProtoCommand(command domain.Event) (*Command, error) {
	switch command.(type) {
	case domain.CreateOrderCommand:
		return toProtoCreateOrderCommand(command.(domain.CreateOrderCommand)), nil
	default:
		return nil, errors.New("Received unknown event type. Don't know how to convert to proto")
	}
}

func toProtoCreateOrderCommand(command domain.CreateOrderCommand) *Command {
	return &Command{
		Command: &Command_CreateOrderCommand{
			CreateOrderCommand: &CreateOrderCommand{
				OrderId:              command.OrderID,
				CustomerId:           command.CustomerID,
				ProductId:            command.ProductID,
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
	command := &Command{}
	err := proto.Unmarshal(data, command)

	if err != nil {
		return nil, err
	}

	return fromProtoCommand(command)
}

func fromProtoCommand(command *Command) (domain.Event, error) {
	switch command.Command.(type) {
	case *Command_CreateOrderCommand:
		return fromProtoCreateOrderCommand(command.GetCreateOrderCommand()), nil
	default:
		return nil, errors.New("Received unknown command type. Don't know how to convert from proto")
	}
}

func fromProtoCreateOrderCommand(command *CreateOrderCommand) domain.CreateOrderCommand {
	return domain.CreateOrderCommand{
		OrderID:              command.OrderId,
		CustomerID:           command.CustomerId,
		ProductID:            command.ProductId,
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
