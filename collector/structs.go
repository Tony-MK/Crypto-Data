package collector


var SYMBOLS  = map[string]string{
	"ETHBTC":"ethbtc",
	"BTCUSDT":"btcusdt",
}


type Ticker struct{
  Event_Type string `json:"e"`
  Event_Time uint64 `json:"E"`
  Symbol string `json:"s"`
  Price_Change string `json:"p"`
  Price_Change_Percent string `json:"P"`
  Weighted_Average_Price string  `json:"w"`
  First_Trade_Price string `json:"x"`
  Last_Trade_Price string `json:"c"`
  Last_Quantity string `json:"Q"`
  Best_Bid_Price string `json:"b"`
  Best_Bid_Quantity string `json:"B"`
  Best_Ask_Price string `json:"a"`
  Best_Ask_Quantity string `json:"A"`
  Open_Price string `json:"o"`
  High string `json:"h"`
  Low string `json:"l"`
  Total_Traded_Base_Asset_Volume string `json:"v"`
  Total_Traded_Quote_Asset_Volume  string `json:"q"`
  Statistics_Open_Time uint64 `json:"O"`
  Statistics_Close_Time uint64 `json:"C"`
  First_Trade_ID uint64 `json:"F"`
  Last_Trade_ID uint64 `json:"L"`
  Total_Number_Of_Trades uint64`json:"n"` 
}

type MiniTicker struct{
   Event_Type string `json:"e"`
   Event_Time uint64 `json:"E"`
   Symbol string `json:"s"`
   Close_Price string `json:"c"`
   Open_Price string `json:"o"`
   High string `json:"h"`
   Low string `json:"l"`
   Total_Traded_Base_Asset_Volume string `json:"v"`
   Total_Traded_Quote_Asset_Volume string `json:"q"`
}

type Depth struct{
  	Event_Type string `json:"e"`
  	Event_Time uint64 `json:"E"`
  	Symbol string `json:"s"`
  	First_Update_ID_In_Event uint64 `json:"U"`
  	Final_Update_ID_In_Event uint64 `json:"u"`
  	Bids_To_Be_Updated [][]string `json:"b"`
  	Asks_To_Be_Updated [][]string `json:"a"`
}

type Trade struct{
	Event_Type string `json:"e"`
  	Event_Time uint64 `json:"E"`
  	Symbol string `json:"s"`
  	Trade_ID uint64 `json:"t"`
   	Price string `json:"p"`
  	Quantity string `json:"q"`
  	Buyer_Order_ID uint64 `json:"b"`
   	Seller_Order_ID uint64 `json:"a"`
   	Trade_Time uint64 `json:"T"`
   	Market_Maker bool `json:"m"`
   	Ignore bool `json:"M"` 
}

type AggTrade struct{
	Event_Type string `json:"e"`
  	Event_Time uint64 `json:"E"`
  	Symbol string `json:"s"`
  	Trade_ID uint64 `json:"t"`
   	Price string `json:"p"`
  	Quantity string `json:"q"`
  	First_Trade_ID uint64 `json:"b"`
   	Last_Trade_ID uint64 `json:"a"`
   	Trade_Time uint64 `json:"T"`
   	Market_Maker bool `json:"m"`
   	Ignore bool `json:"M"` 
}
 