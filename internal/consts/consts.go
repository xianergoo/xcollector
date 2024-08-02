package consts

// TReaderCmdCode 是命令代码的枚举类型
type TReaderCmdCode int

const (
	UnknownCmd TReaderCmdCode = iota
	EchoCmd
	GetVersionCmd
	GetGroupNumberCmd
	CameraModeCmd
	LogOnCmd
	LogOffCmd
	SetTimeCmd
	SetUserCmd
	SetLanguageCmd
	SetPasswordCmd
	DownloadFontCmd
	ClearFontCmd
	LedCmd
	BeepCmd
	SetBeepCmd
	ClearScreenCmd
	TextOutCmd
	TextOutUTF8Cmd
	CollectDataCmd
	DeleteDataCmd
	FormatDataCmd
	GetDataCmd
	// 空白项
	GetSerialNoCmd
	GetDeviceIDCmd
	SetDeviceIDCmd
	SendCommandGroupCmd
	SetGroupNumberCmd
	// 空白项
)

// READER_CMD_VALUES 是命令值的数组
var READER_CMD_VALUES = []int{
	-1,
	1, 2, 3, 5, 0x51, 0x52,
	0x10, 0x53, 0x11, 0x12, 0x18, 0x19,
	0x21, 0x22, 0x23, 0x30, 0x31, 0x32,
	0x40, 0x41, 0x42, 0x48, 0x49, 0x44,
	0x45, 0x46, 0x47, 0x60, 0xC0,
}

// READER_CMD_NAMES 是命令名称的数组
var READER_CMD_NAMES = []string{
	"Unknown",

	"Echo", "Get Version", "Get Group Number", "Camera Mode", "Log On", "Log Off",

	"Set Time", "Set User", "Set Language", "Set Password", "Download Font", "Clear Font",

	"Led", "Beep", "Set Beep", "Clear Screen", "Text Out", "Text Out(UTF8)",

	"Collect Data", "Delete Data", "Format Data", "Get Data", "", "Get Serial No",

	"Get Device ID", "Set Device ID", "Send Command Group", "Set Group Number", "",
}

// lookupCmdName 通过给定的 TReaderCmdCode 查找对应的命令名称
func LookupCmdName(cmdCode TReaderCmdCode) string {
	if cmdCode >= UnknownCmd && cmdCode <= SetGroupNumberCmd {
		return READER_CMD_NAMES[cmdCode]
	}
	return "Unknown Command"
}
