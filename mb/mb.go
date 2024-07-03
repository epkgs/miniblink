package mb

import "unsafe"

// Rect 对应 C++ 中的 _Rect 结构体
type Rect struct {
	X, Y, W, H int
}

// mbPoint 对应 C++ 中的 _mbPoint 结构体
type Point struct {
	X, Y int
}

// mbSize 对应 C++ 中的 _mbSize 结构体
type Size struct {
	W, H int
}

// mbMouseFlags 对应 C++ 中的枚举
type MouseFlags int

const (
	LBUTTON MouseFlags = 0x01
	RBUTTON MouseFlags = 0x02
	SHIFT   MouseFlags = 0x04
	CONTROL MouseFlags = 0x08
	MBUTTON MouseFlags = 0x10
)

// mbKeyFlags 对应 C++ 中的枚举
type KeyFlags int

const (
	EXTENDED KeyFlags = 0x0100
	REPEAT   KeyFlags = 0x4000
)

// mbMouseMsg 对应 C++ 中的枚举
type MouseMsg int

const (
	MSG_MOUSEMOVE     MouseMsg = 0x0200
	MSG_LBUTTONDOWN   MouseMsg = 0x0201
	MSG_LBUTTONUP     MouseMsg = 0x0202
	MSG_LBUTTONDBLCLK MouseMsg = 0x0203
	MSG_RBUTTONDOWN   MouseMsg = 0x0204
	MSG_RBUTTONUP     MouseMsg = 0x0205
	MSG_RBUTTONDBLCLK MouseMsg = 0x0206
	MSG_MBUTTONDOWN   MouseMsg = 0x0207
	MSG_MBUTTONUP     MouseMsg = 0x0208
	MSG_MBUTTONDBLCLK MouseMsg = 0x0209
	MSG_MOUSEWHEEL    MouseMsg = 0x020A
)

// wchar_t 的定义，如果平台不支持 wchar_t，这里会定义一个别名
// 但在Go中，我们通常会直接使用 rune 类型来表示 Unicode 字符
type Wchar_t uint16 // 假设 wchar_t 在不支持的平台上是 unsigned short

// Utf8 类型在Go中通常就是 byte 或 []byte，这里我们直接使用 byte
type Utf8 byte

// 保留C枚举的大小写
type ProxyType int

const (
	PROXY_NONE ProxyType = iota
	PROXY_HTTP
	PROXY_SOCKS4
	PROXY_SOCKS4A
	PROXY_SOCKS5
	PROXY_SOCKS5HOSTNAME
)

// 对应的Go结构体
type Proxy struct {
	Type     ProxyType // 枚举类型，用Go中的别名来表示
	Hostname [100]byte // 使用byte数组代替C中的char数组
	Port     uint16    // 无符号短整型
	Username [50]byte
	Password [50]byte
}

// 枚举类型在Go中通过一组常量模拟
type SettingMask uint32 // 使用无符号32位整数来模拟枚举的基础类型

const (
	SETTING_PROXY            SettingMask = 1
	ENABLE_NODEJS            SettingMask = 1 << 3
	ENABLE_DISABLE_H5VIDEO   SettingMask = 1 << 4
	ENABLE_DISABLE_PDFVIEW   SettingMask = 1 << 5
	ENABLE_DISABLE_CC        SettingMask = 1 << 6
	ENABLE_ENABLE_EGLGLES2   SettingMask = 1 << 7 // 测试功能，请勿使用
	ENABLE_ENABLE_SWIFTSHAER SettingMask = 1 << 8 // 测试功能，请勿使用
)

// mbOnBlinkThreadInitCallback 对应的Go函数签名
// 注意：Go没有直接的函数指针类型别名，但我们可以定义函数类型
type OnBlinkThreadInitCallback func(param uintptr) (anyRes uintptr)

// mbSettings 对应 C++ 中的 _mbSettings 结构体
type Settings struct {
	Proxy                        Proxy
	Mask                         uint
	BlinkThreadInitCallback      OnBlinkThreadInitCallback
	BlinkThreadInitCallbackParam uintptr // 使用unsafe.Pointer来存储任意类型的指针
	Version                      uintptr
	MainDllPath                  *uint16 // 假设使用uint16的指针来表示宽字符字符串（通常使用syscall.UTF16PtrFromString转换）
	MainDllHandle                Handle
	Config                       string // 假设config是C风格的字符串（char*），在Go中通常是*byte
}

// mbViewSettings 对应 C++ 中的 _mbViewSettings 结构体
type ViewSettings struct {
	Size    int
	BgColor uint
}

// WebFrameHandle 对应 C++ 中的 void*
type WebFrameHandle uintptr

// NetJob 对应 C++ 中的 void*
type NetJob uintptr

// WebView 对应 C++ 中的 intptr_t
type WebView uintptr

// 假设NULL_WEBVIEW在Go中是一个常量
const NULL_WEBVIEW = 0

// mbCookieVisitorFunc 是对应C中mbCookieVisitor函数指针的Go接口
type CookieVisitorCallback func(
	params interface{}, // params 在C中是void*类型，但在Go中我们通常不使用这种通用指针。
	name, value, domain, path string,
	secure, httpOnly bool,
	expires *int64, // 假设这是指向过期时间秒数的指针
) bool

// mbCookieCommand 枚举
type CookieCommand int

const (
	CookieCommandClearAllCookies CookieCommand = iota
	CookieCommandClearSessionCookies
	CookieCommandFlushCookiesToFile
	CookieCommandReloadCookiesFromFile
)

// mbNavigationType 枚举
type NavigationType int

const (
	NAVIGATION_TYPE_LINKCLICK    NavigationType = iota
	NAVIGATION_TYPE_FORMSUBMITTE                // 注意：原始代码中可能有拼写错误
	NAVIGATION_TYPE_BACKFORWARD
	NAVIGATION_TYPE_RELOAD
	NAVIGATION_TYPE_FORMRESUBMITT // 注意：原始代码中可能有拼写错误
	NAVIGATION_TYPE_OTHER
)

// mbCursorInfoType 对应 C 中的枚举
type CursorInfoType int

const (
	CursorInfoPointer CursorInfoType = iota
	CursorInfoCross
	CursorInfoHand
	CursorInfoIBeam
	CursorInfoWait
	CursorInfoHelp
	CursorInfoEastResize
	CursorInfoNorthResize
	CursorInfoNorthEastResize
	CursorInfoNorthWestResize
	CursorInfoSouthResize
	CursorInfoSouthEastResize
	CursorInfoSouthWestResize
	CursorInfoWestResize
	CursorInfoNorthSouthResize
	CursorInfoEastWestResize
	CursorInfoNorthEastSouthWestResize
	CursorInfoNorthWestSouthEastResize
	CursorInfoColumnResize
	CursorInfoRowResize
	CursorInfoMiddlePanning
	CursorInfoEastPanning
	CursorInfoNorthPanning
	CursorInfoNorthEastPanning
	CursorInfoNorthWestPanning
	CursorInfoSouthPanning
	CursorInfoSouthEastPanning
	CursorInfoSouthWestPanning
	CursorInfoWestPanning
	CursorInfoMove
	CursorInfoVerticalText
	CursorInfoCell
	CursorInfoContextMenu
	CursorInfoAlias
	CursorInfoProgress
	CursorInfoNoDrop
	CursorInfoCopy
	CursorInfoNone
	CursorInfoNotAllowed
	CursorInfoZoomIn
	CursorInfoZoomOut
	CursorInfoGrab
	CursorInfoGrabbing
	CursorInfoCustom
)

// mbWindowFeatures 对应 C 中的结构体
type WindowFeatures struct {
	X, Y, Width, Height int
	MenuBarVisible      bool
	StatusBarVisible    bool
	ToolBarVisible      bool
	LocationBarVisible  bool
	ScrollbarsVisible   bool
	Resizable           bool
	Fullscreen          bool
}

// PrintSettings 对应 C 中的结构体
type PrintSettings struct {
	StructSize               int
	DPI                      int
	Width, Height            int
	MarginTop, MarginBottom  int
	MarginLeft, MarginRight  int
	IsPrintPageHeadAndFooter bool
	IsPrintBackgroud         bool // 注意：这里假设是 Background 的拼写错误，并修正为 Background
	IsLandscape              bool
	IsPrintToMultiPage       bool
}

// String 在Go中通常不会定义为指针类型别名，而是直接定义结构体
// 由于我们没有具体的mbString结构体内容，这里仅作为占位符
type String struct {
	// 假设这里有一些字段，但具体取决于C中的mbString定义
}

// StringPtr 是指向String结构体的指针类型
type StringPtr *String

// MemBuf 对应C中的_MemBuf结构体
type MemBuf struct {
	Size   int
	Data   uintptr // 使用unsafe.Pointer来表示void*
	Length uintptr // uintptr在Go中通常对应uintptr或uint，具体取决于你的平台和编译器
}

// mbStorageType 枚举
type StorageType int

const (
	StorageTypeString         StorageType = iota // String data with an associated MIME type
	StorageTypeFilename                          // Stores the name of one file being dragged into the renderer
	StorageTypeBinaryData                        // An image being dragged out of the renderer
	StorageTypeFileSystemFile                    // Stores the filesystem URL of one file being dragged into the renderer
)

// Item 结构体
type Item struct {
	StorageType        StorageType
	StringType         *MemBuf // Only valid when storageType == StorageTypeString
	StringData         *MemBuf // Only valid when storageType == StorageTypeString
	FilenameData       *MemBuf // Only valid when storageType == StorageTypeFilename
	DisplayNameData    *MemBuf // Only valid when storageType == StorageTypeFilename
	BinaryData         *MemBuf // Only valid when storageType == StorageTypeBinaryData
	Title              *MemBuf // Title associated with a link when stringType == "text/uri-list". Filename when storageType == StorageTypeBinaryData
	FileSystemURL      *MemBuf // Only valid when storageType == StorageTypeFileSystemFile
	FileSystemFileSize int64   // Only valid when storageType == StorageTypeFileSystemFile
	BaseURL            *MemBuf // Only valid when stringType == "text/html"
}

// WebDragData 结构体
type WebDragData struct {
	ItemList         []*Item
	ItemListLength   int
	ModifierKeyState int // State of Shift/Ctrl/Alt/Meta keys
	FileSystemID     *MemBuf
}

type WebDragOperation uint32

// mbWebDragOperation 是对应C枚举的Go常量集
const (
	WebDragOperationNone    WebDragOperation = 0
	WebDragOperationCopy    WebDragOperation = 1
	WebDragOperationLink    WebDragOperation = 2
	WebDragOperationGeneric WebDragOperation = 4
	WebDragOperationPrivate WebDragOperation = 8
	WebDragOperationMove    WebDragOperation = 16
	WebDragOperationDelete  WebDragOperation = 32
	WebDragOperationEvery   WebDragOperation = 0xffffffff
)

// mbWebDragOperationsMask 类型别名，用于表示可能的拖放操作掩码
type WebDragOperationsMask = WebDragOperation

// mbResourceType 是资源类型的枚举
type ResourceType int

const (
	RESOURCE_TYPE_MAIN_FRAME     ResourceType = iota // 顶层页面
	RESOURCE_TYPE_SUB_FRAME                          // 框架或iframe
	RESOURCE_TYPE_STYLESHEET                         // CSS样式表
	RESOURCE_TYPE_SCRIPT                             // 外部脚本
	RESOURCE_TYPE_IMAGE                              // 图片（jpg/gif/png等）
	RESOURCE_TYPE_FONT_RESOURCE                      // 字体资源
	RESOURCE_TYPE_SUB_RESOURCE                       // "其他"子资源
	RESOURCE_TYPE_OBJECT                             // 对象（或embed）标签的插件资源，或插件请求的资源
	RESOURCE_TYPE_MEDIA                              // 媒体资源
	RESOURCE_TYPE_WORKER                             // 专用工作线程的主资源
	RESOURCE_TYPE_SHARED_WORKER                      // 共享工作线程的主资源
	RESOURCE_TYPE_PREFETCH                           // 明确请求的预取资源
	RESOURCE_TYPE_FAVICON                            // 网站图标
	RESOURCE_TYPE_XHR                                // XMLHttpRequest
	RESOURCE_TYPE_PING                               // <a ping>的ping请求
	RESOURCE_TYPE_SERVICE_WORKER                     // 服务工作线程的主资源
	RESOURCE_TYPE_LAST_TYPE
)

// RequestType 枚举对应的Go常量
type RequestType int

const (
	RequestTypeInvalidation RequestType = iota
	RequestTypeGet
	RequestTypePost
	RequestTypePut
)

// Slist 对应C中的结构体
type Slist struct {
	Data *string // 使用指针来模拟C中的char*
	Next *Slist
}

// MenuItemId 枚举对应的Go常量
type MenuItemId uint32

const (
	MenuSelectedAllId      MenuItemId = 1 << 1
	MenuSelectedTextId     MenuItemId = 1 << 2
	MenuUndoId             MenuItemId = 1 << 3
	MenuCopyImageId        MenuItemId = 1 << 4
	MenuInspectElementAtId MenuItemId = 1 << 5
	MenuCutId              MenuItemId = 1 << 6
	MenuPasteId            MenuItemId = 1 << 7
	MenuPrintId            MenuItemId = 1 << 8
	MenuGoForwardId        MenuItemId = 1 << 9
	MenuGoBackId           MenuItemId = 1 << 10
	MenuReloadId           MenuItemId = 1 << 11
)

// typedef void* WebSocketChannel;
type WebSocketChannel uintptr

// 回调函数签名
type OnWillConnectCallback func(webView WebView, param uintptr, channel WebSocketChannel, url string) (StringPtr, bool)
type OnConnectedCallback func(webView WebView, param uintptr, channel WebSocketChannel) bool
type OnReceiveCallback func(webView WebView, param uintptr, channel WebSocketChannel, opCode int, buf *byte, len int) (StringPtr, bool)
type OnSendCallback func(webView WebView, param uintptr, channel WebSocketChannel, opCode int, buf *byte, len int) (StringPtr, bool)
type OnErrorCallback func(webView WebView, param uintptr, channel WebSocketChannel)

// 回调接口定义
type WebsocketHookCallbacks interface {
	OnWillConnect(OnWillConnectCallback)
	OnConnected(OnConnectedCallback)
	OnReceive(OnReceiveCallback)
	OnSend(OnSendCallback)
	OnError(OnErrorCallback)
}

//////////////////////////////////////////////////////////////////////////

// JsType 是枚举类型的别名
type JsType int

// JsType 枚举对应的Go常量
const (
	JsTypeNumber JsType = iota
	JsTypeString
	JsTypeBool
	//kJsTypeObject  // 注释掉的部分在Go中也可以省略
	//kJsTypeFunction // 同上
	JsTypeUndefined
	//kJsTypeArray    // 同上
	JsTypeNull
	JsTypeV8Value
)

// JsValue 是int64的别名
type JsValue uintptr

// JsExecState 是void*的别名，在Go中通常使用unsafe.Pointer来表示
type JsExecState uintptr

// OnGetPdfPageDataCallback 对应C中的回调函数类型
type OnGetPdfPageDataCallback func(webView WebView, param uintptr, data uintptr, size uintptr)

// RunJsCallback 对应C中的回调函数类型
type RunJsCallback func(webView WebView, param uintptr, es JsExecState, v JsValue)

// JsQueryCallback 对应C中的回调函数类型
type JsQueryCallback func(webView WebView, param uintptr, es JsExecState, queryId int64, customMsg int, request string)

// mbTitleChangedCallback
type TitleChangedCallback func(webView WebView, param unsafe.Pointer, title uintptr) uintptr

// mbMouseOverUrlChangedCallback
type MouseOverUrlChangedCallback func(webView WebView, param uintptr, url string)

// mbURLChangedCallback
type URLChangedCallback func(webView WebView, param uintptr, url string, canGoBack bool, canGoForward bool)

// mbURLChangedCallback2
type URLChangedCallback2Callback func(webView WebView, param uintptr, frameId WebFrameHandle, url string)

// PaintUpdatedCallback
// 注意：HDC是Windows特有的类型，这里假设它已经作为uintptr类型定义
type HDC uintptr
type PaintUpdatedCallback func(webView WebView, param uintptr, hdc HDC, x, y, cx, cy int)

// AcceleratedPaintCallback
type AcceleratedPaintCallback func(webView WebView, param uintptr, typ int, dirytRects *Rect, dirytRectsSize uintptr, sharedHandle uintptr)

// PaintBitUpdatedCallback
type PaintBitUpdatedCallback func(webView WebView, param uintptr, buffer uintptr, r *Rect, width, height int)

// mbAlertBoxCallback
type AlertBoxCallback func(webView WebView, param uintptr, msg string)

// mbConfirmBoxCallback
type ConfirmBoxCallback func(webView WebView, param uintptr, msg string) bool

// mbPromptBoxCallback
type PromptBoxCallback func(webView WebView, param uintptr, msg, defaultResult string) StringPtr

// NavigationCallback
type NavigationCallback func(webView WebView, param uintptr, navigationType NavigationType, url string) bool

// CreateViewCallback
type CreateViewCallback func(webView WebView, param uintptr, navigationType NavigationType, url string, windowFeatures *WindowFeatures) WebView

// DocumentReadyCallback
type DocumentReadyCallback func(webView WebView, param uintptr, frameId WebFrameHandle)

// mbCloseCallback
type CloseCallback func(webView WebView, param uintptr, unuse uintptr) bool

// mbDestroyCallback
type DestroyCallback func(webView WebView, param uintptr, unuse uintptr) bool

// mbOnShowDevtoolsCallback
type OnShowDevtoolsCallback func(webView WebView, param uintptr)

// mbDidCreateScriptContextCallback
type DidCreateScriptContextCallback func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, extensionGroup, worldId int)

// mbGetPluginListCallback
type GetPluginListCallback func(refresh bool, pluginListBuilder, param uintptr) bool

// NetResponseCallback
type NetResponseCallback func(webView WebView, param uintptr, url string, job NetJob) bool

// ThreadCallback
type ThreadCallback func(param1, param2 uintptr)

// NodeOnCreateProcessCallback
type WCHAR uint16

// 定义别名以匹配Windows API中的类型
type DWORD uint32
type WORD uint16
type LPWSTR *uint16 // 宽字符字符串指针
type LPBYTE *byte   // 字节指针
type Handle uintptr // 通常使用syscall.Handle作为句柄类型
// STARTUPINFOW 结构体在Go中的表示
type STARTUPINFOW struct {
	Cb              DWORD
	LpReserved      LPWSTR
	LpDesktop       LPWSTR
	LpTitle         LPWSTR
	DwX             DWORD
	DwY             DWORD
	DwXSize         DWORD
	DwYSize         DWORD
	DwXCountChars   DWORD
	DwYCountChars   DWORD
	DwFillAttribute DWORD
	DwFlags         DWORD
	WShowWindow     WORD
	CbReserved2     WORD
	LpReserved2     LPBYTE
	HStdInput       Handle
	HStdOutput      Handle
	HStdError       Handle
}
type NodeOnCreateProcessCallback func(webView WebView, param uintptr, applicationPath, arguments *WCHAR, startup *STARTUPINFOW)

// mbLoadingResult 是一个表示加载结果的类型
type LoadingResult int

// 定义mbLoadingResult的常量
const (
	LOADING_SUCCEEDED LoadingResult = iota // iota是Go语言的枚举计数器
	LOADING_FAILED
	LOADING_CANCELED
)

// mbLoadingFinishCallback
type LoadingFinishCallback func(webView WebView, param uintptr, FrameId WebFrameHandle, URL string, Result LoadingResult, FailedReason string)

// mbDownloadCallback
type DownloadCallback func(webView WebView, param uintptr, FrameId WebFrameHandle, URL string, DownloadJob uintptr) bool

// mbConsoleLevel 枚举
type ConsoleLevel int

const (
	LevelLog ConsoleLevel = iota + 1
	LevelWarning
	LevelError
	LevelDebug
	LevelInfo
	LevelRevokedError
	LevelLast = LevelRevokedError
)

// mbConsoleCallback
type ConsoleCallback func(webView WebView, param uintptr, level ConsoleLevel, message string, sourceName string, sourceLine uint, stackTrace string)

// mbOnCallUiThread
// 这是一个在UI线程上调用的函数类型
type OnCallUiThreadCallback func(webView WebView, param uintptr)

// mbCallUiThread
// 这是一个用于在UI线程上调用OnCallUiThread的函数类型
type CallUiThreadCallback func(webView WebView, funcOnUiThread OnCallUiThreadCallback, param uintptr)

// LoadUrlBeginCallback
type LoadUrlBeginCallback func(webView WebView, param uintptr, url string, job NetJob) bool

// LoadUrlEndCallback
type LoadUrlEndCallback func(webView WebView, param uintptr, url string, job NetJob, buf uintptr, len int)

// 注意：buf 参数是一个 unsafe.Pointer，因为 C 中的 void* 在 Go 中没有直接对应。
// 在实际使用时，你可能需要将其转换为合适的类型（如 []byte）来访问其内容。

// LoadUrlFailCallback
type LoadUrlFailCallback func(webView WebView, param uintptr, url string, job NetJob)

// LoadUrlHeadersReceivedCallback
type LoadUrlHeadersReceivedCallback func(webView WebView, param uintptr, url string, job NetJob)

// LoadUrlFinishCallback
// 注意：这里使用了 Utf8 类型别名，但在 Go 中直接使用 string 即可，因为 Go 的字符串默认就是 UTF-8 编码的。
type LoadUrlFinishCallback func(webView WebView, param uintptr, url string, job NetJob, len int)

// mbDidCreateScriptContextCallback
// type DidCreateScriptContextCallback func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, extensionGroup, worldId int)

// mbWillReleaseScriptContextCallback
type WillReleaseScriptContextCallback func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, worldId int)

// mbNetGetFaviconCallback
type NetGetFaviconCallback func(webView WebView, param uintptr, url string, buf *MemBuf)

// AsynRequestState 枚举类型
type AsynRequestState int

// 枚举常量
const (
	AsynRequestStateOk   AsynRequestState = 0
	AsynRequestStateFail AsynRequestState = 1
)

// CanGoBackForwardCallback
type CanGoBackForwardCallback func(webView WebView, param uintptr, state AsynRequestState, b bool)

// GetCookieCallback
type GetCookieCallback func(webView WebView, param uintptr, state AsynRequestState, cookie string)

// 类型别名定义
type V8ContextPtr uintptr // v8引擎的上下文指针
type V8Isolate uintptr    // v8引擎的隔离区指针

// GetSourceCallback
type GetSourceCallback func(webView WebView, param uintptr, mhtml string)

// GetContentAsMarkupCallback
type GetContentAsMarkupCallback func(webView WebView, param uintptr, content string, size uintptr)

// 结构体指针类型定义
type WebUrlRequest struct{}
type WebUrlResponse struct{}
type WebUrlRequestPtr *WebUrlRequest
type WebUrlResponsePtr *WebUrlResponse

// 回调函数类型定义
// 注意：Go没有__stdcall调用约定，所以这里省略了它
type OnUrlRequestWillRedirectCallback func(webView WebView, param uintptr, oldRequest WebUrlRequestPtr, request WebUrlRequestPtr, redirectResponse WebUrlResponsePtr)
type OnUrlRequestDidReceiveResponseCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, response WebUrlResponsePtr)
type OnUrlRequestDidReceiveDataCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, data *byte, dataLength int)
type OnUrlRequestDidFailCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, error *uint8)
type OnUrlRequestDidFinishLoadingCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, finishTime float64)

// UrlRequestCallbacks 结构体
type UrlRequestCallbacks struct {
	WillRedirectCallback       OnUrlRequestWillRedirectCallback
	DidReceiveResponseCallback OnUrlRequestDidReceiveResponseCallback
	DidReceiveDataCallback     OnUrlRequestDidReceiveDataCallback
	DidFailCallback            OnUrlRequestDidFailCallback
	DidFinishLoadingCallback   OnUrlRequestDidFinishLoadingCallback
}

// DownloadOpt 枚举
type DownloadOpt int

const (
	DownloadOptCancel    DownloadOpt = iota // iota是Go中的计数器，用于枚举值
	DownloadOptCacheData                    // 后续枚举值会自动递增
)

// 函数指针类型
type NetJobDataRecvCallback func(ptr uintptr, job NetJob, data *byte, length int)
type NetJobDataFinishCallback func(ptr uintptr, job NetJob, result LoadingResult)
type PopupDialogSaveNameCallback func(ptr uintptr, filePath *uint16)

// 结构体定义
type NetJobDataBind struct {
	Param          uintptr
	RecvCallback   NetJobDataRecvCallback
	FinishCallback NetJobDataFinishCallback
}

type DownloadBind struct {
	Param            uintptr
	RecvCallback     NetJobDataRecvCallback
	FinishCallback   NetJobDataFinishCallback
	SaveNameCallback PopupDialogSaveNameCallback
}

type FileFilter struct {
	Name       string // 转换为 string 以匹配C中的const Utf8*
	Extensions string // 转换为 string 以匹配C中的const Utf8*
}

// mbDialogProperties 枚举定义了对话框的各种属性
type DialogProperties uint32

const (
	// kDialogPropertiesOpenFile 允许选择文件
	DialogPropertiesOpenFile DialogProperties = 1 << 1
	// kDialogPropertiesOpenDirectory 允许选择文件夹
	DialogPropertiesOpenDirectory DialogProperties = 1 << 2
	// kDialogPropertiesMultiSelections 允许多选
	DialogPropertiesMultiSelections DialogProperties = 1 << 3
	// kDialogPropertiesShowHiddenFiles 显示对话框中的隐藏文件
	DialogPropertiesShowHiddenFiles DialogProperties = 1 << 4
	// kDialogPropertiesCreateDirectory macOS - 允许通过对话框的形式创建新的目录
	DialogPropertiesCreateDirectory DialogProperties = 1 << 5
	// kDialogPropertiesPromptToCreate Windows - 如果输入的文件路径在对话框中不存在，则提示创建
	DialogPropertiesPromptToCreate DialogProperties = 1 << 6
	// kDialogPropertiesNoResolveAliases macOS - 禁用自动的别名路径(符号链接) 解析
	DialogPropertiesNoResolveAliases DialogProperties = 1 << 7
	// kDialogPropertiesTreatPackageAsDirectory macOS - 将包(如.app 文件夹) 视为目录而不是文件
	DialogPropertiesTreatPackageAsDirectory DialogProperties = 1 << 8
	// kDialogPropertiesDontAddToRecent Windows - 不要将正在打开的项目添加到最近的文档列表中
	DialogPropertiesDontAddToRecent DialogProperties = 1 << 9
)

// DialogOptions 结构体定义了对话框的选项
type DialogOptions struct {
	Magic                   int    // 'mbdo'
	Title                   string // 标题
	DefaultPath             string // 默认路径
	ButtonLabel             string // 按钮标签
	Filters                 *FileFilter
	FiltersCount            int
	Prop                    DialogProperties
	Message                 string // 消息
	SecurityScopedBookmarks bool
}

// DownloadOptions 结构体定义了下载选项
type DownloadOptions struct {
	Magic             int
	SaveAsPathAndName bool
}

// mbDownloadInBlinkThreadCallback 函数指针类型
type DownloadInBlinkThreadCallback func(
	webView WebView,
	params uintptr,
	expectedContentLength uintptr,
	url string,
	mime string,
	disposition string,
	job NetJob,
	dataBind *NetJobDataBind,
) DownloadOpt

// mbPdfDatas 结构体
type PdfDatas struct {
	Count int
	Sizes *uintptr  // 指向 uintptr 类型的指针数组
	Datas **uintptr // 指向 void* 类型的指针的指针数组，即 const void**
}

// PrintPdfDataCallback 函数指针类型
type PrintPdfDataCallback func(webview WebView, param uintptr, datas *PdfDatas)

// ScreenshotSettings 结构体
type ScreenshotSettings struct {
	StructSize int
	Width      int
	Height     int
}

// PrintBitmapCallback 函数指针类型
type PrintBitmapCallback func(webview WebView, param uintptr, data *byte, size uintptr)

// OnScreenshot 函数指针类型
type OnScreenshotCallback func(webView WebView, param uintptr, data *byte, size uintptr)

// 枚举 mbHttBodyElementType
type HttBodyElementType int

const (
	HttBodyElementTypeData HttBodyElementType = iota
	HttBodyElementTypeFile
)

// 结构体 PostBodyElement
type PostBodyElement struct {
	Size       int
	Type       HttBodyElementType
	Data       *MemBuf // 假设 MemBuf 已经被定义
	FilePath   StringPtr
	FileStart  int64
	FileLength int64 // -1 表示到文件末尾
}

// PostBodyElements 结构体
type PostBodyElements struct {
	Size        int
	Element     **PostBodyElement
	ElementSize uintptr
	IsDirty     bool
}

// mbWillSendRequestInfo 结构体
type WillSendRequestInfo struct {
	URL              StringPtr
	NewURL           StringPtr
	ResourceType     ResourceType
	HTTPResponseCode int
	Method           StringPtr
	Referrer         StringPtr
	Headers          uintptr // 使用unsafe.Pointer来替代void*
}

// mbViewLoadType 枚举
type ViewLoadType int

const (
	DID_START_LOADING ViewLoadType = iota
	DID_STOP_LOADING
	DID_NAVIGATE
	DID_NAVIGATE_IN_PAGE
	DID_GET_RESPONSE_DETAILS
	DID_GET_REDIRECT_REQUEST
	DID_POST_REQUEST
)

// mbViewLoadCallbackInfo 结构体
type ViewLoadCallbackInfo struct {
	Size                int
	Frame               WebFrameHandle
	WillSendRequestInfo *WillSendRequestInfo
	URL                 string
	PostBody            *PostBodyElements
	Job                 NetJob
}

// mbNetViewLoadInfoCallback 函数指针类型
type NetViewLoadInfoCallback func(webView WebView, param uintptr, type_ ViewLoadType, info *ViewLoadCallbackInfo)

// mbwindow-----------------------------------------------------------------------------------
// mbWindowType 枚举
type WindowType int

const (
	WINDOW_TYPE_POPUP WindowType = iota
	WINDOW_TYPE_TRANSPARENT
	WINDOW_TYPE_CONTROL
)

// mbWindowInfo 枚举
type WindowInfo int

const (
	WINDOW_INFO_SHARTD_TEXTURE_ENABLE WindowInfo = 1 << 16 // 注意：这里可能有个拼写错误，应该是 WINDOW_INFO_SHARED_TEXTURE_ENABLE
)

type RECT struct {
	Left, Top, Right, Bottom int32
}

// mbDraggableRegion 结构体
type DraggableRegion struct {
	Bounds    RECT
	Draggable bool
}

// 回调函数类型
type WindowClosingCallback func(webview WebView, param uintptr) bool
type WindowDestroyCallback func(webview WebView, param uintptr)

// mbDraggableRegionsChangedCallback 回调函数类型
type DraggableRegionsChangedCallback func(webview WebView, param uintptr, rects *DraggableRegion, rectCount int)

// mbPrintintStep 枚举（注意：这里可能有个拼写错误，应该是 mbPrintStep）
type PrintintStep int

const (
	PrintStepStart PrintintStep = iota
	PrintStepPreview
	PrintStepPrinting
)

// mbPrintintSettings 结构体
type PrintintSettings struct {
	Dpi    int
	Width  int
	Height int
	Scale  float32
}

// DefaultPrinterSettings 结构体
type DefaultPrinterSettings struct {
	StructSize         int32 // 默认是4 * 10
	IsLandscape        bool  // 是否为横向打印格式
	IsPrintHeadFooter  bool
	IsPrintBackground  bool
	EdgeDistanceLeft   int32 // 单位：毫米
	EdgeDistanceTop    int32
	EdgeDistanceRight  int32
	EdgeDistanceBottom int32
	Copies             int32
	PaperType          int32 // 假设DMPAPER_A4等是预定义的常量
}

// mbPrintingCallback 函数指针类型
type PrintingCallback func(
	webview WebView,
	param uintptr,
	step PrintintStep,
	hDC Handle, // 假设HDC在Go中用syscall.Handle表示
	settings *PrintintSettings,
	pageCount int32, // 假设页数是32位整数
) bool

// mbImageBufferToDataURLCallback 函数指针类型
type ImageBufferToDataURLCallback func(
	webView WebView,
	param uintptr,
	data *byte, // 假设data是指向字节数组的指针
	size uintptr, // 假设size是uintptr类型，在Go中用uintptr表示
) StringPtr

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type HWND uintptr

type WPARAM uintptr
type LPARAM uintptr
type LRESULT uintptr

type mbInit func(settings *Settings)
type mbUninit func()
type mbCreateInitSettings func() *Settings
type mbSetInitSettings func(settings *Settings, name, value string)
type mbCreateWebView func() WebView
type mbDestroyWebView func(webview WebView)
type mbCreateWebWindow func(typ WindowType, parent Handle, x, y, width, height int) WebView
type mbCreateWebCustomWindow func(parent Handle, style, styleEx uint32, x, y, width, height int) WebView
type mbMoveWindow func(webview WebView, x, y, width, height int)
type mbMoveToCenter func(webview WebView)
type mbSetAutoDrawToHwnd func(webview WebView, b bool)
type mbGetCaretRect func(webview WebView, r *Rect)
type mbSetAudioMuted func(webview WebView, b bool)
type mbIsAudioMuted func(webview WebView) bool
type mbCreateString func(str string, length uintptr) StringPtr
type mbCreateStringWithoutNullTermination func(str string, length uintptr) StringPtr
type mbDeleteString func(str StringPtr)
type mbGetStringLen func(str StringPtr) uintptr
type mbGetString func(str StringPtr) string
type mbSetProxy func(webView WebView, proxy *Proxy)
type mbSetDebugConfig func(webView WebView, debugString, param string)
type mbNetSetData func(jobPtr NetJob, buf *byte, len int)
type mbNetHookRequest func(jobPtr NetJob)
type mbNetChangeRequestUrl func(jobPtr NetJob, url string)
type mbNetContinueJob func(jobPtr NetJob)
type mbNetGetRawHttpHeadInBlinkThread func(jobPtr NetJob) *Slist
type mbNetGetRawResponseHeadInBlinkThread func(jobPtr NetJob) *Slist
type mbNetHoldJobToAsynCommit func(jobPtr NetJob) bool
type mbNetCancelRequest func(jobPtr NetJob)
type mbNetOnResponse func(webviewHandle WebView, callback NetResponseCallback, param uintptr)
type mbNetSetWebsocketCallback func(webview WebView, callbacks *WebsocketHookCallbacks, param uintptr)
type mbNetSendWsText func(channel WebSocketChannel, buf *byte, len uintptr)
type mbNetSendWsBlob func(channel WebSocketChannel, buf *byte, len uintptr)
type mbNetEnableResPacket func(webviewHandle WebView, pathName *uint16)
type mbNetGetPostBody func(jobPtr NetJob) *PostBodyElements
type mbNetCreatePostBodyElements func(webView WebView, length uintptr) *PostBodyElements
type mbNetFreePostBodyElements func(elements *PostBodyElements)
type mbNetCreatePostBodyElement func(webView WebView) *PostBodyElement
type mbNetFreePostBodyElement func(element *PostBodyElement)
type mbNetCreateWebUrlRequest func(url, method, mime string) WebUrlRequestPtr
type mbNetAddHTTPHeaderFieldToUrlRequest func(request WebUrlRequestPtr, name, value string)
type mbNetStartUrlRequest func(webView WebView, request WebUrlRequestPtr, param uintptr, callbacks *UrlRequestCallbacks) int
type mbNetGetHttpStatusCode func(response WebUrlResponsePtr) int
type mbNetGetRequestMethod func(jobPtr NetJob) RequestType
type mbNetGetExpectedContentLength func(response WebUrlResponsePtr) int64
type mbNetGetResponseUrl func(response WebUrlResponsePtr) string
type mbNetCancelWebUrlRequest func(requestId int)
type mbSetViewProxy func(webView WebView, proxy *Proxy)
type mbNetSetMIMEType func(jobPtr NetJob, typeStr string)
type mbNetGetMIMEType func(jobPtr NetJob) string
type mbNetGetHTTPHeaderField func(job NetJob, key string, fromRequestOrResponse bool) string
type mbNetGetReferrer func(jobPtr NetJob) string
type mbNetSetHTTPHeaderField func(jobPtr NetJob, key, value string, response bool)
type mbSetMouseEnabled func(webView WebView, b bool)
type mbSetTouchEnabled func(webView WebView, b bool)
type mbSetSystemTouchEnabled func(webView WebView, b bool)
type mbSetContextMenuEnabled func(webView WebView, b bool)
type mbSetNavigationToNewWindowEnabled func(webView WebView, b bool)
type mbSetHeadlessEnabled func(webView WebView, b bool)
type mbSetDragDropEnable func(webView WebView, b bool)
type mbSetDragEnable func(webView WebView, b bool)
type mbSetContextMenuItemShow func(webView WebView, item MenuItemId, isShow bool)
type mbSetHandle func(webView WebView, wnd HWND)
type mbSetHandleOffset func(webView WebView, x, y int)
type mbGetHostHWND func(webView WebView) HWND
type mbSetTransparent func(webviewHandle WebView, transparent bool)
type mbSetViewSettings func(webviewHandle WebView, settings *ViewSettings)
type mbSetCspCheckEnable func(webView WebView, b bool)
type mbSetNpapiPluginsEnabled func(webView WebView, b bool)
type mbSetMemoryCacheEnable func(webView WebView, b bool)
type mbSetCookie func(webView WebView, url, cookie string)
type mbSetCookieEnabled func(webView WebView, enable bool)
type mbSetCookieJarPath func(webView WebView, path string)
type mbSetCookieJarFullPath func(webView WebView, path string)
type mbSetLocalStorageFullPath func(webView WebView, path *uint16)
type mbGetTitle func(webView WebView) string
type mbSetWindowTitle func(webView WebView, title string)
type mbSetWindowTitleW func(webView WebView, title string)
type mbGetUrl func(webView WebView) string
type mbGetCursorInfoType func(webView WebView) int
type mbAddPluginDirectory func(webView WebView, path string)
type mbSetUserAgent func(webView WebView, userAgent string)
type mbSetZoomFactor func(webView WebView, factor float32)
type mbGetZoomFactor func(webView WebView) float32
type mbSetDiskCacheEnabled func(webView WebView, enable bool)
type mbSetDiskCachePath func(webView WebView, path string)
type mbSetDiskCacheLimit func(webView WebView, limit uint64)
type mbSetDiskCacheLimitDisk func(webView WebView, limit uint64)
type mbSetDiskCacheLevel func(webView WebView, level int)
type mbSetResourceGc func(webView WebView, intervalSec int)
type mbCanGoForward func(webView WebView, callback CanGoBackForwardCallback, param uintptr)
type mbCanGoBack func(webView WebView, callback CanGoBackForwardCallback, param uintptr)
type mbGetCookie func(webView WebView, callback GetCookieCallback, param uintptr)
type mbGetCookieOnBlinkThread func(webView WebView) string
type mbClearCookie func(webView WebView)
type mbResize func(webView WebView, w, h int)
type mbOnNavigation func(webView WebView, callback NavigationCallback, param uintptr)
type mbOnNavigationSync func(webView WebView, callback NavigationCallback, param uintptr)
type mbOnCreateView func(webView WebView, callback CreateViewCallback, param uintptr)
type mbOnDocumentReady func(webView WebView, callback DocumentReadyCallback, param uintptr)
type mbOnDocumentReadyInBlinkThread func(webView WebView, callback DocumentReadyCallback, param uintptr)
type mbOnPaintUpdated func(webView WebView, callback PaintUpdatedCallback, param uintptr)
type mbOnPaintBitUpdated func(webView WebView, callback PaintBitUpdatedCallback, param uintptr)
type mbOnAcceleratedPaint func(webView WebView, callback AcceleratedPaintCallback, param uintptr)
type mbOnLoadUrlBegin func(webView WebView, callback LoadUrlBeginCallback, param uintptr)
type mbOnLoadUrlEnd func(webView WebView, callback LoadUrlEndCallback, param uintptr)
type mbOnLoadUrlFail func(webView WebView, callback LoadUrlFailCallback, param uintptr)
type mbOnLoadUrlHeadersReceived func(webView WebView, callback LoadUrlHeadersReceivedCallback, param uintptr)
type mbOnLoadUrlFinish func(webView WebView, callback LoadUrlFinishCallback, param uintptr)
type mbOnTitleChanged func(webView WebView, callback TitleChangedCallback, callbackParam uintptr)
type mbOnURLChanged func(webView WebView, callback URLChangedCallback, callbackParam uintptr)
type mbOnLoadingFinish func(webView WebView, callback LoadingFinishCallback, param uintptr)
type mbOnDownload func(webView WebView, callback DownloadCallback, param uintptr)
type mbOnDownloadInBlinkThread func(webView WebView, callback DownloadInBlinkThreadCallback, param uintptr)
type mbOnAlertBox func(webView WebView, callback AlertBoxCallback, param uintptr)
type mbOnConfirmBox func(webView WebView, callback ConfirmBoxCallback, param uintptr)
type mbOnPromptBox func(webView WebView, callback PromptBoxCallback, param uintptr)
type mbOnNetGetFavicon func(webView WebView, callback NetGetFaviconCallback, param uintptr)
type mbOnConsole func(webView WebView, callback ConsoleCallback, param uintptr)
type mbOnClose func(webView WebView, callback CloseCallback, param uintptr)
type mbOnDestroy func(webView WebView, callback DestroyCallback, param uintptr)
type mbOnPrinting func(webView WebView, callback PrintingCallback, param uintptr)
type mbOnPluginList func(webView WebView, callback GetPluginListCallback, param uintptr)
type mbOnImageBufferToDataURL func(webView WebView, callback ImageBufferToDataURLCallback, param uintptr)
type mbOnDidCreateScriptContext func(webView WebView, callback DidCreateScriptContextCallback, param uintptr)
type mbOnWillReleaseScriptContext func(webView WebView, callback WillReleaseScriptContextCallback, param uintptr)
type mbGoBack func(webView WebView)
type mbGoForward func(webView WebView)
type mbGoToOffset func(webView WebView, offset int)
type mbGoToIndex func(webView WebView, index int)
type mbNavigateAtIndex func(webView WebView, index int)
type mbGetNavigateIndex func(webView WebView) int
type mbStopLoading func(webView WebView)
type mbReload func(webView WebView)
type mbPerformCookieCommand func(webView WebView, command CookieCommand)
type mbInsertCSSByFrame func(webView WebView, frameId WebFrameHandle, cssText string)
type mbEditorSelectAll func(webView WebView)
type mbEditorUnSelect func(webView WebView)
type mbEditorCopy func(webView WebView)
type mbEditorCut func(webView WebView)
type mbEditorPaste func(webView WebView)
type mbEditorDelete func(webView WebView)
type mbEditorUndo func(webView WebView)
type mbEditorRedo func(webView WebView)
type mbSetEditable func(webView WebView, editable bool)
type mbFireMouseEvent func(webView WebView, message uint, x, y int, flags uint) bool
type mbFireContextMenuEvent func(webView WebView, x, y int, flags uint) bool
type mbFireMouseWheelEvent func(webView WebView, x, y, delta int, flags uint) bool
type mbFireKeyUpEvent func(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool
type mbFireKeyDownEvent func(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool
type mbFireKeyPressEvent func(webView WebView, charCode, flags uint, systemKey bool) bool
type mbFireWindowsMessage func(webView WebView, hWnd HWND, message uint, wParam WPARAM, lParam LPARAM, result *LRESULT) bool
type mbSetFocus func(webView WebView)
type mbKillFocus func(webView WebView)
type mbShowWindow func(webView WebView, show bool)
type mbLoadURL func(webView WebView, url string)
type mbLoadHtmlWithBaseUrl func(webView WebView, html string, baseUrl string)
type mbPostURL func(webView WebView, url string, postData *byte, postLen int)
type mbGetLockedViewDC func(webView WebView) HDC
type mbUnlockViewDC func(webView WebView)
type mbWake func(webView WebView)
type mbJsToV8Value func(es JsExecState, v JsValue) uintptr
type mbGetGlobalExecByFrame func(webView WebView, frameId WebFrameHandle) JsExecState
type mbJsToDouble func(es JsExecState, v JsValue) float64
type mbJsToBoolean func(es JsExecState, v JsValue) bool
type mbJsToString func(es JsExecState, v JsValue) string
type mbGetJsValueType func(es JsExecState, v JsValue) JsType
type mbOnJsQuery func(webView WebView, callback JsQueryCallback, param uintptr)
type mbResponseQuery func(webView WebView, queryId int64, customMsg int, response string)
type mbRunJs func(webView WebView, frameId WebFrameHandle, script string, isInClosure bool, callback RunJsCallback, param, unused uintptr)
type mbRunJsSync func(webView WebView, frameId WebFrameHandle, script string, isInClosure bool) JsValue
type mbWebFrameGetMainFrame func(webView WebView) WebFrameHandle
type mbIsMainFrame func(webView WebView, frameId WebFrameHandle) bool
type mbSetNodeJsEnable func(webView WebView, enable bool)
type mbSetDeviceParameter func(webView WebView, device, paramStr string, paramInt int, paramFloat float32)
type mbGetContentAsMarkup func(webView WebView, callback GetContentAsMarkupCallback, param uintptr, frameId WebFrameHandle)
type mbGetSource func(webView WebView, callback GetSourceCallback, param uintptr)
type mbUtilSerializeToMHTML func(webView WebView, callback GetSourceCallback, param uintptr)
type mbUtilCreateRequestCode func(registerInfo string) string
type mbUtilIsRegistered func(defaultPath string) bool
type mbUtilPrint func(webView WebView, frameId WebFrameHandle, printParams *PrintSettings) bool
type mbUtilBase64Encode func(str string) string
type mbUtilBase64Decode func(str string) string
type mbUtilDecodeURLEscape func(url string) string
type mbUtilEncodeURLEscape func(url string) string
type mbUtilCreateV8Snapshot func(str string) *MemBuf
type mbUtilPrintToPdf func(webView WebView, frameId WebFrameHandle, settings *PrintSettings, callback PrintPdfDataCallback, param uintptr)
type mbUtilPrintToBitmap func(webView WebView, frameId WebFrameHandle, settings *ScreenshotSettings, callback PrintBitmapCallback, param uintptr)
type mbUtilScreenshot func(webView WebView, settings *ScreenshotSettings, callback OnScreenshotCallback, param uintptr)
type mbUtilsSilentPrint func(webView WebView, settings string) bool
type mbUtilSetDefaultPrinterSettings func(webView WebView, setting *DefaultPrinterSettings)
type mbPopupDownloadMgr func(webView WebView, url string, downloadJob uintptr) bool
type mbPopupDialogAndDownload func(webView WebView, dialogOpt *DialogOptions, contentLength uint64, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt
type mbDownloadByPath func(webView WebView, downloadOptions *DownloadOptions, path string, contentLength uint64, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt
type mbGetPdfPageData func(webView WebView, callback OnGetPdfPageDataCallback, param uintptr)
type mbCreateMemBuf func(webView WebView, buf *byte, length uint64) *MemBuf
type mbFreeMemBuf func(buf *MemBuf)
type mbSetUserKeyValue func(webView WebView, key string, value uintptr)
type mbGetUserKeyValue func(webView WebView, key string) uintptr
type mbPluginListBuilderAddPlugin func(builder uintptr, name, description, fileName string)
type mbPluginListBuilderAddMediaTypeToLastPlugin func(builder uintptr, name, description string)
type mbPluginListBuilderAddFileExtensionToLastMediaType func(builder uintptr, fileExtension string)
type mbGetBlinkMainThreadIsolate func() V8Isolate
type mbWebFrameGetMainWorldScriptContext func(webView WebView, frameId WebFrameHandle, contextOut *V8ContextPtr)
type mbEnableHighDPISupport func()
type mbRunMessageLoop func()
type mbGetContentWidth func(webView WebView) int
type mbGetContentHeight func(webView WebView) int
type mbGetWebViewForCurrentContext func() WebView
type mbRegisterEmbedderCustomElement func(webviewHandle WebView, frameId WebFrameHandle, name string, options, outResult uintptr) bool
type mbOnNodeCreateProcess func(webviewHandle WebView, callback NodeOnCreateProcessCallback, param uintptr)
type mbOnThreadIdle func(callback ThreadCallback, param1, param2 uintptr)
type mbGetProcAddr func(name string) uintptr
type mbSetMbDllPath func(dllPath string)
type mbSetMbMainDllPath func(dllPath string)
type mbFillFuncPtr func()
