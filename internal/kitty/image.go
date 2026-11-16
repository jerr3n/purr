package kitty

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
)

// \x1b _G

func Image(img []byte, w io.Writer) {
	encoded := base64.StdEncoding.EncodeToString(img)
	writer := bufio.NewWriter(w)
	defer writer.Flush()
	for i := 0; i < len(encoded); i += 4096 {
		cur := encoded[i:min(i+4096, len(encoded))]
		end := i+4096 >= len(encoded)
		switch {
		case i == 0 && end:
			fmt.Fprintf(writer, "\x1b_Ga=T,f=100;%s\x1b\\", cur)
		case i == 0:
			fmt.Fprintf(writer, "\x1b_Ga=T,f=100,m=1;%s\x1b\\", cur)
		case end:

			fmt.Fprintf(writer, "\x1b_Gm=0;%s\x1b\\", cur)
		default:
			fmt.Fprintf(writer, "\x1b_Gm=1;%s\x1b\\", cur)
		}
	}
}
