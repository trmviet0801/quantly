package dto

import "fmt"

type OrderGetResponseDto struct {
	ID             string      `json:"id"`
	ClientOrderID  string      `json:"client_order_id"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
	SubmittedAt    string      `json:"submitted_at"`
	FilledAt       *string     `json:"filled_at"`
	ExpiredAt      *string     `json:"expired_at"`
	CanceledAt     *string     `json:"canceled_at"`
	FailedAt       *string     `json:"failed_at"`
	ReplacedAt     *string     `json:"replaced_at"`
	ReplacedBy     *string     `json:"replaced_by"`
	Replaces       *string     `json:"replaces"`
	AssetID        string      `json:"asset_id"`
	Symbol         string      `json:"symbol"`
	AssetClass     string      `json:"asset_class"`
	Notional       *string     `json:"notional"`
	Qty            string      `json:"qty"`
	FilledQty      string      `json:"filled_qty"`
	FilledAvgPrice *string     `json:"filled_avg_price"`
	OrderClass     string      `json:"order_class"`
	OrderType      string      `json:"order_type"`
	Type           string      `json:"type"`
	Side           string      `json:"side"`
	PositionIntent string      `json:"position_intent"`
	TimeInForce    string      `json:"time_in_force"`
	LimitPrice     string      `json:"limit_price"`
	StopPrice      *string     `json:"stop_price"`
	Status         string      `json:"status"`
	ExtendedHours  bool        `json:"extended_hours"`
	Legs           interface{} `json:"legs"` // Use []OrderResponseDto if this is a slice of sub-orders
	TrailPercent   *string     `json:"trail_percent"`
	TrailPrice     *string     `json:"trail_price"`
	HWM            *string     `json:"hwm"`
	Commission     string      `json:"commission"`
	Subtag         *string     `json:"subtag"`
	Source         string      `json:"source"`
	ExpiresAt      string      `json:"expires_at"`
}

func (o OrderGetResponseDto) String() string {
	return fmt.Sprintf(
		`OrderGetResponseDto:
  ID: %s
  ClientOrderID: %s
  CreatedAt: %s
  UpdatedAt: %s
  SubmittedAt: %s
  FilledAt: %s
  ExpiredAt: %s
  CanceledAt: %s
  FailedAt: %s
  ReplacedAt: %s
  ReplacedBy: %s
  Replaces: %s
  AssetID: %s
  Symbol: %s
  AssetClass: %s
  Notional: %s
  Qty: %s
  FilledQty: %s
  FilledAvgPrice: %s
  OrderClass: %s
  OrderType: %s
  Type: %s
  Side: %s
  PositionIntent: %s
  TimeInForce: %s
  LimitPrice: %s
  StopPrice: %s
  Status: %s
  ExtendedHours: %t
  TrailPercent: %s
  TrailPrice: %s
  HWM: %s
  Commission: %s
  Subtag: %s
  Source: %s
  ExpiresAt: %s`,
		o.ID,
		o.ClientOrderID,
		o.CreatedAt,
		o.UpdatedAt,
		o.SubmittedAt,
		nullableString(o.FilledAt),
		nullableString(o.ExpiredAt),
		nullableString(o.CanceledAt),
		nullableString(o.FailedAt),
		nullableString(o.ReplacedAt),
		nullableString(o.ReplacedBy),
		nullableString(o.Replaces),
		o.AssetID,
		o.Symbol,
		o.AssetClass,
		nullableString(o.Notional),
		o.Qty,
		o.FilledQty,
		nullableString(o.FilledAvgPrice),
		o.OrderClass,
		o.OrderType,
		o.Type,
		o.Side,
		o.PositionIntent,
		o.TimeInForce,
		o.LimitPrice,
		nullableString(o.StopPrice),
		o.Status,
		o.ExtendedHours,
		nullableString(o.TrailPercent),
		nullableString(o.TrailPrice),
		nullableString(o.HWM),
		o.Commission,
		nullableString(o.Subtag),
		o.Source,
		o.ExpiresAt,
	)
}

func nullableString(s *string) string {
	if s == nil {
		return "null"
	}
	return *s
}
