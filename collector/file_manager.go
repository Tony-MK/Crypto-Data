package collector
import (
	"os"
	//"github.com/kr/pretty"
	"encoding/json"
	"strconv"
	"errors"
)
type Filer struct{
	Format string
	Directory string
	ErrorChan chan error
	Collector *collector

}

const UINT_BASE = 10
const STREAM_BYTE_N = 11

func(f *Filer)FileArray(raw []byte){
	switch(raw[STREAM_BYTE_N+1]){
		case 'm':
			var miniTickers []MiniTicker
			err := json.Unmarshal(raw,&miniTickers)
			if(f.Collector.CheckError(err)){return}
			for _,miniTicker := range(miniTickers) {f.writeMiniTicker(&miniTicker)}
			break
		case 't':
			var tickers []Ticker
			if(f.Collector.CheckError(json.Unmarshal(raw,&tickers))){return}
			for _,ticker := range(tickers) {f.writeTicker(&ticker)}
			break
		default:
			f.ErrorChan <- errors.New("Unknown Stream\n"+string(raw[STREAM_BYTE_N:40]));
			break;
	}
	return
}
func(f *Filer)File(raw []byte){

	var s = [2]byte{0,0}
	for i := STREAM_BYTE_N; i < STREAM_BYTE_N+20; i++ {
		if(raw[i] == '@'){
			s[0] = raw[i+1]
			s[1] = raw[i+2]
			break
		}
	}

	if(s[0] == 0){f.ErrorChan <- errors.New("Unknown Stream\n"+string(raw[STREAM_BYTE_N:40]));}
	var err error
	raw,err = FilterStream(raw)
	if(f.Collector.CheckError(err)){return}
	switch(s[0]){
	 	case '!':
	 		f.FileArray(raw);
	 		break;
	 	case 'd':
	 		var depth Depth
			if(f.Collector.CheckError(json.Unmarshal(raw,&depth))){break}
	 		f.writeDepth(&depth);
	 		break;
	 	case 'a':
	 		if(s[1] == 'r'){
	 			f.FileArray(raw);
	 		}else{
		 		var aggTrade AggTrade
				if(f.Collector.CheckError(json.Unmarshal(raw,&aggTrade))){break}
		 		f.writeAggTrade(&aggTrade);
		 	}
	 		break;
	 	case 't':
	 		if(s[1] == 'r'){
	 			var trade Trade
				if(f.Collector.CheckError(json.Unmarshal(raw,&trade))){break}
	 			f.writeTrade(&trade)
	 		}else{
	 			var ticker Ticker
				if(f.Collector.CheckError(json.Unmarshal(raw,&ticker))){break}
	 			f.writeTicker(&ticker)
	 		}
	 		break
	 	case 'm':
	 		var miniTicker MiniTicker
			if(f.Collector.CheckError(json.Unmarshal(raw,&miniTicker))){break}
	 		f.writeMiniTicker(&miniTicker);
	 		break
	 	default:
	 		f.ErrorChan <- errors.New("Unknown Stream...");
	 		break
	} 
	return
}





func(f *Filer)Save(filepath string,time string,data *[]byte,b...bool){
	file, err := os.Create(filepath+"/"+time+".json")
	if(err != nil){
		if(len(b) > 0){
			f.ErrorChan <- err
			return
		}
		if(f.Collector.CheckError(os.MkdirAll(filepath,1))){return}
		f.Save(filepath,time,data)
	}
	defer file.Close();
	file.Write(*data);
	file.Close();
}


func(f *Filer)writeTicker(ticker *Ticker){
	buf, err := json.Marshal(map[string]interface{}{
		"v":ticker.Total_Traded_Quote_Asset_Volume,
		"q":ticker.Total_Traded_Base_Asset_Volume,
		"h":ticker.High,
		"l":ticker.Low,
		"o":ticker.Open_Price,
		"p":ticker.Price_Change,
    	"P":ticker.Price_Change_Percent,
  		"w":ticker.Weighted_Average_Price,
  		"x":ticker.First_Trade_Price,
  		"c":ticker.Last_Trade_Price,
    	"Q":ticker.Last_Quantity,
  		"b":ticker.Best_Bid_Price,
    	"B":ticker.Best_Bid_Quantity,
    	"a":ticker.Best_Ask_Price,
  	 	"A":ticker.Best_Ask_Quantity ,
		"O":ticker.Statistics_Open_Time,
		"C":ticker.Statistics_Close_Time,
		"F":ticker.First_Trade_ID,
		"L":ticker.Last_Trade_ID,
		"n":ticker.Total_Number_Of_Trades,
	})
	if(err != nil){
		f.ErrorChan <- err
		return
	}
	f.Save(f.Directory+"Tickers/"+ticker.Symbol,strconv.FormatUint(ticker.Event_Time,UINT_BASE),&buf)
}



func(f *Filer)writeDepth(depth *Depth){
	data, err := json.Marshal(map[string]interface{}{
		"U":depth.First_Update_ID_In_Event,
		"u":depth.Final_Update_ID_In_Event,
		"b":depth.Bids_To_Be_Updated,
		"a":depth.Asks_To_Be_Updated,
	})
	if(f.Collector.CheckError(err)){return}
	filepath := f.Directory+"/Depth/"+depth.Symbol
	f.Save(filepath,strconv.FormatUint(depth.Event_Time,UINT_BASE),&data)
	return
}

func(f *Filer)writeMiniTicker(ticker *MiniTicker){
	buf, err := json.Marshal(map[string]interface{}{
		"v":ticker.Total_Traded_Quote_Asset_Volume,
		"q":ticker.Total_Traded_Base_Asset_Volume,
		"h":ticker.High,
		"l":ticker.Low,
		"o":ticker.Open_Price,
		"c":ticker.Close_Price,
	})
	if(err != nil){
		panic(err)
	}
	f.Save(f.Directory+"MiniTickers/"+ticker.Symbol,strconv.FormatUint(ticker.Event_Time,UINT_BASE),&buf)
	return
}

func(f *Filer)writeTrade(trade *Trade){
	m := 1
	if(!trade.Market_Maker){m = 0}
	buf, err := json.Marshal(map[string]interface{}{
		"t":trade.Trade_ID,
		"p":trade.Price,
		"q":trade.Quantity,
		"b":trade.Buyer_Order_ID,
		"a":trade.Seller_Order_ID,
		"T":trade.Trade_Time,
		"m":m,
		
	})
	if(f.Collector.CheckError(err)){return}
	f.Save(f.Directory+"Trades/"+trade.Symbol,strconv.FormatUint(trade.Event_Time,UINT_BASE),&buf)
	return
}

func(f *Filer)writeAggTrade(aggTrade *AggTrade){
	m := 1
	if(!aggTrade.Market_Maker){m = 0}
	buf, err := json.Marshal(map[string]interface{}{
		"t":aggTrade.Trade_ID,
		"p":aggTrade.Price,
		"q":aggTrade.Quantity,
		"b":aggTrade.First_Trade_ID,
		"a":aggTrade.Last_Trade_ID,
		"T":aggTrade.Trade_Time,
		"m":m,
		
	})
	
	if(f.Collector.CheckError(err)){return}
	f.Save(f.Directory+"AggTrades/"+aggTrade.Symbol,strconv.FormatUint(aggTrade.Event_Time,UINT_BASE),&buf)
	return
}
var DATAKEY = [7]byte{'"','d','a','t','a','"',':'}
const DATAKEY_LEN = 7
func FilterStream(raw []byte)([]byte,error){
		for i := STREAM_BYTE_N; i < len(raw); i++ {
			for n := 0; n < DATAKEY_LEN; n++ {
				if(raw[i+n] == DATAKEY[n]){
					if(n == DATAKEY_LEN-1){
						raw = raw[i+n+1:len(raw)-1]
						return raw,nil
					}
				}

			}
		}
	return raw,errors.New("Could Not find Array")
}


