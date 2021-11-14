package fonts

type Font string

type font struct {
	height  int
	letters func(r rune) []string
}

const (
	Small  Font = "small"
	Native Font = "native"
)

func FontHeight(f Font) int {
	return fontMap[f].height
}

var fontMap = make(map[Font]font)

func init() {
	small := font{}
	small.height = 4
	// Straight: https://patorjk.com/software/taag/#p=display&f=Straight&t=SNAKE%20SCORE%200123456789
	letters := map[rune][]string{
		'A': {
			`    `,
			` /\ `,
			`/--\`,
			`    `,
		},
		'C': {
			` __`,
			`/  `,
			`\__`,
			`   `,
		},
		'E': {
			` __`,
			`|_ `,
			`|__`,
			`   `,
		},
		'G': {
			` __ `,
			`/ _ `,
			`\__)`,
			`    `,
		},
		'I': {
			` `,
			`|`,
			`|`,
			` `,
		},
		'L': {
			`   `,
			`|  `,
			`|__`,
			`   `,
		},
		'M': {
			`    `,
			`|\/|`,
			`|  |`,
			`    `,
		},
		'N': {
			`    `,
			`|\ |`,
			`| \|`,
			`    `,
		},
		'O': {
			` __ `,
			`/  \`,
			`\__/`,
			`    `,
		},
		'P': {
			` __ `,
			`|__)`,
			`|   `,
			`    `,
		},
		'R': {
			` __ `,
			`|__)`,
			`| \ `,
			`    `,
		},
		'S': {
			` __`,
			`(_ `,
			`__)`,
			`   `,
		},
		'T': {
			`___`,
			` | `,
			` | `,
			`   `,
		},
		'U': {
			`    `,
			`/  \`,
			`\__/`,
			`    `,
		},
		'V': {
			`    `,
			`\  /`,
			` \/ `,
			`    `,
		},
		':': {
			`  `,
			`. `,
			`. `,
			`  `,
		},
		' ': {
			`   `,
			`   `,
			`   `,
			`   `,
		},
		'0': {
			` __  `,
			`/  \ `,
			`\__/ `,
			`     `,
		},
		'1': {
			`   `,
			`/| `,
			` | `,
			`   `,
		},
		'2': {
			`__  `,
			` _) `,
			`/__ `,
			`    `,
		},
		'3': {
			`__  `,
			` _) `,
			`__) `,
			`    `,
		},
		'4': {
			`     `,
			`|__| `,
			`   | `,
			`     `,
		},
		'5': {
			` __ `,
			`|_  `,
			`__) `,
			`    `,
		},
		'6': {
			` __  `,
			`/__  `,
			`\__) `,
			`     `,
		},
		'7': {
			`___ `,
			`  / `,
			` /  `,
			`    `,
		},
		'8': {
			` __  `,
			`(__) `,
			`(__) `,
			`     `,
		},
		'9': {
			` __  `,
			`(__\ `,
			` __/ `,
			`     `,
		},
	}

	fontMap[Small] = font{
		height: 4,
		letters: func(r rune) []string {
			return letters[r]
		},
	}
}

func init() {
	fontMap[Native] = font{
		height: 1,
		letters: func(r rune) []string {
			return []string{string(r)}
		},
	}
}
