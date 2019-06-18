// 12.3. Display, a Recursive Value Printer
package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen, Jack D. Ripper":  "Streling Hayden",
			`Maj. T.J. "King" Kone`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", strangelove)
	// Display strangelove (main.Movie):
	//strangelove.Title = "Dr. Strangelove"
	//strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
	//strangelove.Year = 1964
	//strangelove.Color = false
	//strangelove.Actor["Brig. Gen, Jack D. Ripper"] = "Streling Hayden"
	//strangelove.Actor["Maj. T.J. \"King\" Kone"] = "Slim Pickens"
	//strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
	//strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
	//strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
	//strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
	//strangelove.Oscars[0] = "Best Actor (Nomin.)"
	//strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
	//strangelove.Oscars[2] = "Best Director (Nomin.)"
	//strangelove.Oscars[3] = "Best Picture (Nomin.)"
	//strangelove.Sequel = nil

	Display("os.Stderr", os.Stderr)
	// Display os.Stderr (*os.File):
	//(*(*os.Stderr).file).pfd.fdmu.state = 0
	//(*(*os.Stderr).file).pfd.fdmu.rsema = 0
	//(*(*os.Stderr).file).pfd.fdmu.wsema = 0
	//(*(*os.Stderr).file).pfd.Sysfd = 2
	//(*(*os.Stderr).file).pfd.pd.runtimeCtx = 0
	//(*(*os.Stderr).file).pfd.iovecs = nil
	//(*(*os.Stderr).file).pfd.csema = 0
	//(*(*os.Stderr).file).pfd.isBlocking = 1
	//(*(*os.Stderr).file).pfd.IsStream = true
	//(*(*os.Stderr).file).pfd.ZeroReadIsEOF = true
	//(*(*os.Stderr).file).pfd.isFile = true
	//(*(*os.Stderr).file).name = "/dev/stderr"
	//(*(*os.Stderr).file).dirinfo = nil
	//(*(*os.Stderr).file).nonblock = false
	//(*(*os.Stderr).file).stdoutOrErr = true

	Display("rV", reflect.ValueOf(os.Stderr))
	// Display rV (reflect.Value):
	//(*rV.typ).size = 8
	//(*rV.typ).ptrdata = 8
	//(*rV.typ).hash = 871609668
	//(*rV.typ).tflag = 1
	//(*rV.typ).align = 8
	//(*rV.typ).fieldAlign = 8
	//(*rV.typ).kind = 54
	//(*(*rV.typ).alg).hash = func(unsafe.Pointer, uintptr) uintptr 0x1050fd0
	//(*(*rV.typ).alg).equal = func(unsafe.Pointer, unsafe.Pointer) bool 0x1002e20
	//(*(*rV.typ).gcdata) = 1
	//(*rV.typ).str = 8697
	//(*rV.typ).ptrToThis = 0
	//rV.ptr = unsafe.Pointer value
	//rV.flag = 22

	var i interface{} = 3
	Display("i", i)
	// Display i (int):
	// i = 3

	Display("&i", &i)
	// Display &i (*interface {}):
	//(*&i).type = int
	//(*&i).value = 3

	//// a struct that points to itself
	//type Cycle struct {
	//	Value int
	//	Tail *Cycle
	//}
	//var c Cycle
	//c = Cycle{42, &c}
	//Display("c", c)
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
