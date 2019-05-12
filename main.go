package main

import (
	"./collector"
	"fmt"
)

func main(){
	var errChannel = make(chan error,1)

	c, err := collector.NewCollector(errChannel,"","")
	if(err != nil){panic(err)}
	c.CheckError(c.CollectTickers());
	c.CheckError(c.CollectMiniTickers());

	c.CheckError(c.CollectMiniTickers("BTCETH","BTCADA","BTCUSDT"))
	c.CheckError(c.CollectTickers("BTCETH","BTCADA","BTCUSDT"))
	c.CheckError(c.CollectDepth("ADABTC","ADAETC","ADAETH"));
	c.CheckError(c.CollectAggTrades("BTCETH","BTCADA","BTCUSDT"));
	c.CheckError(c.CollectTrades("BTCETH","BTCADA","BTCUSDT"));
	go c.StartCollecting();
	print("Started to collect\n")
	for{
		fmt.Println(<-errChannel)
	}
}
