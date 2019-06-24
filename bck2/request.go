type OrderDetails struct {
	CurrencyId		int64
	Amount 			float64
	Rate			float64
	RateCurrencyId 	int64
	OrderType 		string
	Token			string
	RequestOrders 	RequestOrders
}

type RequestOrders struct {
	AllUser			string
	ExcludeMine 	string
	OrderType 	 	string
	IsProcessed		string
}


type OrderId struct {
	OrderId 	 				int64
	Token							string
	RequestOrderList 	RequestOrderList
}
