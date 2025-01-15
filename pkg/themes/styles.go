package themes

type BorderStyle string

const (
	Square  BorderStyle = "rounded-none"
	Rounded BorderStyle = "rounded"
	Pill    BorderStyle = "rounded-full"
)

type FormControlSize string

const (
	SmallFormControl  FormControlSize = "px-1"
	MediumFormControl FormControlSize = "px-2 py-1"
	LargeFormControl  FormControlSize = "px-3 py-2"
)
