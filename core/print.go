package core

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

var banner string = `
██╗   ██╗██████╗ ██╗     ██████╗ ██████╗ ██╗   ██╗████████╗███████╗
██║   ██║██╔══██╗██║     ██╔══██╗██╔══██╗██║   ██║╚══██╔══╝██╔════╝
██║   ██║██████╔╝██║     ██████╔╝██████╔╝██║   ██║   ██║   █████╗  
██║   ██║██╔══██╗██║     ██╔══██╗██╔══██╗██║   ██║   ██║   ██╔══╝  
╚██████╔╝██║  ██║███████╗██████╔╝██║  ██║╚██████╔╝   ██║   ███████╗
 ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚══════╝
                                                                   `

func Banner() {
	red := color.New(color.FgRed)
	red.Println(banner)
}

//
//

func Success(data string) {
	green := color.New(color.FgGreen)
	green.Printf("[+]")
	fmt.Println(data)
}

func SuccessYellow(data string) {
	yellow := color.New(color.FgYellow)
	yellow.Printf("[+]")
	fmt.Println(data)
}

func Error(data string) {
	red := color.New(color.FgRed)
	red.Printf("[-]")
	fmt.Println(data)
}

func Info(data string) {
	blue := color.New(color.FgBlue)
	blue.Printf("[*]")
	fmt.Println(data)
}

//
//

func Find(slice []int, val int) (int, bool) {
	// Check if element existst in array

	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//
//

func DirFound(url string, statuscode int, codes []int) {
	if _, i := Find(codes, statuscode); !i {
		return
	}

	// Colors
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	yellowHi := color.New(color.FgHiYellow)
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)

	// Status codes
	azul := []int{100, 101, 102}                                                                                                                                                                                                        // Informational codes
	verde := []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}                                                                                                                                                                    // Success codes
	amarelo := []int{300, 301, 302, 303, 304, 305, 307, 308}                                                                                                                                                                            // Redirection codes
	vermelho := []int{400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 426, 428, 429, 431, 444, 451, 499, 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511, 599} // Error codes

	fmt.Printf("%s ", url) // Print Dir
	yellowHi.Print("-> ")

	// Print status code
	if _, i := Find(azul, statuscode); i {
		blue.Printf("%d %s\n", statuscode, http.StatusText(statuscode))

	} else if _, i := Find(verde, statuscode); i {
		green.Printf("%d %s\n", statuscode, http.StatusText(statuscode))

	} else if _, i := Find(amarelo, statuscode); i {
		yellow.Printf("%d %s\n", statuscode, http.StatusText(statuscode))

	} else if _, i := Find(vermelho, statuscode); i {
		red.Printf("%d %s\n", statuscode, http.StatusText(statuscode))

	} else {
		yellow.Printf("%d %s\n", statuscode, http.StatusText(statuscode))
	}
}
