// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type OrdersPaymentMethod string

const (
	OrdersPaymentMethodCOD  OrdersPaymentMethod = "COD"
	OrdersPaymentMethodMOMO OrdersPaymentMethod = "MOMO"
	OrdersPaymentMethodBANK OrdersPaymentMethod = "BANK"
)

func (e *OrdersPaymentMethod) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrdersPaymentMethod(s)
	case string:
		*e = OrdersPaymentMethod(s)
	default:
		return fmt.Errorf("unsupported scan type for OrdersPaymentMethod: %T", src)
	}
	return nil
}

type NullOrdersPaymentMethod struct {
	OrdersPaymentMethod OrdersPaymentMethod
	Valid               bool // Valid is true if OrdersPaymentMethod is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrdersPaymentMethod) Scan(value interface{}) error {
	if value == nil {
		ns.OrdersPaymentMethod, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrdersPaymentMethod.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrdersPaymentMethod) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrdersPaymentMethod), nil
}

type OrdersStatus string

const (
	OrdersStatusPending    OrdersStatus = "pending"
	OrdersStatusPaid       OrdersStatus = "paid"
	OrdersStatusProcessing OrdersStatus = "processing"
	OrdersStatusShipped    OrdersStatus = "shipped"
	OrdersStatusDelivered  OrdersStatus = "delivered"
	OrdersStatusCancelled  OrdersStatus = "cancelled"
	OrdersStatusFailed     OrdersStatus = "failed"
)

func (e *OrdersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrdersStatus(s)
	case string:
		*e = OrdersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrdersStatus: %T", src)
	}
	return nil
}

type NullOrdersStatus struct {
	OrdersStatus OrdersStatus
	Valid        bool // Valid is true if OrdersStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrdersStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrdersStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrdersStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrdersStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrdersStatus), nil
}

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Cart struct {
	ID        int32
	UserID    int32
	ProductID int32
	Quantity  int32
}

type Order struct {
	ID              int32
	UserID          int32
	PaymentMethod   OrdersPaymentMethod
	Status          NullOrdersStatus
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	ShippingAddress string
	Total           int64
}

type OrderItem struct {
	ID          int32
	OrderID     int32
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
	CreatedAt   sql.NullTime
}

type Product struct {
	ID          int32
	Name        string
	Description sql.NullString
	Price       int64
	Quantity    int32
	ImageUrl    string
	CreatedAt   sql.NullTime
	DeletedAt   sql.NullTime
}

type User struct {
	ID        int32
	Email     string
	Password  string
	Role      NullUserRole
	CreatedAt sql.NullTime
}

type UserProfile struct {
	ID          int32
	UserID      int32
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
	Address     sql.NullString
}
