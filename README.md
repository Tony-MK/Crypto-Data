# Crypto-Data
A Simple Framework for collect data from the Binance Exchange websockets API.

Data feed avaiable  for 
1) Trades
2) Tickers
3) MiniTickers
4) Depth

# Installation

1) `git clone https://github.com/Tony-MK/Crypto-Data`

# Usage
`
  // Create an err Channel 
  var errChannel = make(chan error,1)
  //Create a Collector object 
  
	c, err := collector.NewCollector(errChannel,"","")
	if(err != nil){panic(err)}
  
  // ALL Currenices
  //Tickers
	c.CheckError(c.CollectTickers());
  
  //Mini Tickers 
	c.CheckError(c.CollectMiniTickers());
  
  // Specific Currenices
  // Tickers
	c.CheckError(c.CollectTickers("BTCETH","BTCADA","BTCUSDT"))
  // Mini Tickers
	c.CheckError(c.CollectMiniTickers("BTCETH","BTCADA","BTCUSDT"))
  // Depth
	c.CheckError(c.CollectDepth("ADABTC","ADAETC","ADAETH"));
  //Aggerate Trades
	c.CheckError(c.CollectAggTrades("BTCETH","BTCADA","BTCUSDT"));
  //Trades
	c.CheckError(c.CollectTrades("BTCETH","BTCADA","BTCUSDT"));
  
  
	go c.StartCollecting();
  
	for{
    // logging Errors 
		fmt.Println(<-errChannel)
	}
  

`
