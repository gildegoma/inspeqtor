package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{

	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 35: // ['#','#']
			return 4
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 44: // [',',',']
			return 5
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case r == 97: // ['a','a']
			return 6
		case 98 <= r && r <= 114: // ['b','r']
			return 2
		case r == 115: // ['s','s']
			return 7
		case r == 116: // ['t','t']
			return 8
		case r == 117: // ['u','u']
			return 2
		case r == 118: // ['v','v']
			return 9
		case r == 119: // ['w','w']
			return 10
		case 120 <= r && r <= 122: // ['x','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S1
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S2
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S3
	func(r rune) int {
		switch {
		case r == 92: // ['\','\']
			return 11

		default:
			return 12
		}

	},

	// S4
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 13
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		default:
			return 14
		}

	},

	// S5
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S6
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 107: // ['a','k']
			return 2
		case r == 108: // ['l','l']
			return 15
		case 109 <= r && r <= 122: // ['m','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S7
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 100: // ['a','d']
			return 2
		case r == 101: // ['e','e']
			return 16
		case 102 <= r && r <= 122: // ['f','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S8
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 110: // ['a','n']
			return 2
		case r == 111: // ['o','o']
			return 17
		case 112 <= r && r <= 122: // ['p','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S9
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 104: // ['a','h']
			return 2
		case r == 105: // ['i','i']
			return 18
		case 106 <= r && r <= 122: // ['j','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S10
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 104: // ['a','h']
			return 2
		case r == 105: // ['i','i']
			return 19
		case 106 <= r && r <= 122: // ['j','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S11
	func(r rune) int {
		switch {

		default:
			return 20
		}

	},

	// S12
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 21
		case r == 92: // ['\','\']
			return 22

		default:
			return 12
		}

	},

	// S13
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S14
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 13

		default:
			return 14
		}

	},

	// S15
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 100: // ['a','d']
			return 2
		case r == 101: // ['e','e']
			return 23
		case 102 <= r && r <= 122: // ['f','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S16
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 109: // ['a','m']
			return 2
		case r == 110: // ['n','n']
			return 24
		case 111 <= r && r <= 115: // ['o','s']
			return 2
		case r == 116: // ['t','t']
			return 25
		case 117 <= r && r <= 122: // ['u','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S17
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S18
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case r == 97: // ['a','a']
			return 26
		case 98 <= r && r <= 122: // ['b','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S19
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 115: // ['a','s']
			return 2
		case r == 116: // ['t','t']
			return 27
		case 117 <= r && r <= 122: // ['u','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S20
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 21
		case r == 92: // ['\','\']
			return 22

		default:
			return 12
		}

	},

	// S21
	func(r rune) int {
		switch {

		}
		return NoState

	},

	// S22
	func(r rune) int {
		switch {

		default:
			return 20
		}

	},

	// S23
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 113: // ['a','q']
			return 2
		case r == 114: // ['r','r']
			return 28
		case 115 <= r && r <= 122: // ['s','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S24
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 99: // ['a','c']
			return 2
		case r == 100: // ['d','d']
			return 29
		case 101 <= r && r <= 122: // ['e','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S25
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S26
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S27
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 103: // ['a','g']
			return 2
		case r == 104: // ['h','h']
			return 30
		case 105 <= r && r <= 122: // ['i','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S28
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 115: // ['a','s']
			return 2
		case r == 116: // ['t','t']
			return 31
		case 117 <= r && r <= 122: // ['u','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S29
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S30
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S31
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 114: // ['a','r']
			return 2
		case r == 115: // ['s','s']
			return 32
		case 116 <= r && r <= 122: // ['t','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S32
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 2
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 62: // ['>','>']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},
}
