package collector
import(
	"github.com/gorilla/websocket"
	"github.com/kr/pretty"
)


const (
	ENDPOINT = "wss://stream.binance.com:9443/stream?streams="
	DATA_DIRECTORY = "./Data/"
)
var STREAMS = map[string]string{"!miniTicker@arr":"miniTickers"};
type Options struct{
	Format string
	Directory string
}
type collector struct{
	Filer *Filer
	Conn *websocket.Conn
	Url string
	ErrorChan chan error
}

func NewCollector(err chan error,dir string,format string)(*collector,error){
	if(dir == ""){
		dir = DATA_DIRECTORY
	}
	if(format == ""){
		format = "JSON"
	}
	c := &collector{
		Filer:&Filer{
			Format:format,
			Directory:dir,
			Collector:nil,
		},
		Conn: nil,
		Url: "",
		ErrorChan: err,
	}
	c.Filer.Collector = c;
	return c,nil
}


// 24hr rolling window mini-ticker statistics
func(c *collector)CollectTickers(symbols...string)error{
	if(len(symbols) == 0){
		c.Url += "/!ticker@arr";
		return nil
	}
	return c.collect("@ticker",symbols)
}


//24hr rolling window ticker statistics
func(c *collector)CollectMiniTickers(symbols...string)error{
	if(len(symbols) == 0){
		c.Url += "/!miniTicker@arr";
		return nil;
	}
	return c.collect("@miniTicker",symbols);
}

//Order book price and quantity depth updates used to locally manage an order book 
func(c *collector)CollectDepth(symbols...string)error{
	return c.collect("@depth",symbols)
}

//The Trade Streams push raw trade information; each trade has a unique buyer and seller.
func(c *collector)CollectTrades(symbols...string)error{
	return c.collect("@trade",symbols)
}


//The Aggregate Trade Streams push trade information that is aggregated for a single taker order.
func(c *collector)CollectAggTrades(symbols...string)error{
	return c.collect("@aggTrade",symbols)
}

func(c *collector)collect(stream string,symbols []string)error{

	if(len(symbols) == 0){
		panic("Number of Symbol has to be GREATER than ZERO")
	}
	for n := len(symbols)-1; 0 < n ;n--{
		sym,err := toLower(symbols[n])
		if err != nil{return err}
		c.Url +=  "/"+sym+stream;
	}
	return nil
}

func(c *collector)CheckError(err error)bool{
	if(err != nil){
		c.ErrorChan <- err;
		return true
	}
	return false
}

func(c *collector)StartCollecting(){
	var err error
	pretty.Printf("Conecting to Binance via WebSockets streams ..\nUrl: %s", c.Url[1:])
	c.Conn, _, err = websocket.DefaultDialer.Dial(ENDPOINT+c.Url[1:],nil);
	defer c.Conn.Close()
	if(c.CheckError(err)){return}
	for {
		_,data,err := c.Conn.ReadMessage();
		if(c.CheckError(err)){return}
		go c.Filer.File(data)
	}	
}


