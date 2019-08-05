package main

import (
	"learn/channels"
)

func main(){
	channels.Test_channels()
	channels.Test_chan_buff()
	channels.Test_worker()
	channels.Test_chan_direct()
	channels.Test_select()
	channels.Test_timeout()
}