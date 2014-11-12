
/*
*/
package parser

const numNTSymbols = 5
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Config
		3, // RouteStatement
		-1, // ChannelParameters
		2, // SetStatement
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S2
		
		-1, // S'
		6, // Config
		3, // RouteStatement
		-1, // ChannelParameters
		2, // SetStatement
		

	},
	gotoRow{ // S3
		
		-1, // S'
		7, // Config
		3, // RouteStatement
		-1, // ChannelParameters
		2, // SetStatement
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		18, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		23, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		24, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Config
		-1, // RouteStatement
		-1, // ChannelParameters
		-1, // SetStatement
		

	},
	
}
