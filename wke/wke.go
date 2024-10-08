package wke

type Rect struct {
	X, Y, W, H int32
}

// wkePoint 结构体
type Point struct {
	X int32
	Y int32
}

// wkeSize 结构体
type Size struct {
	W int32
	H int32
}

// MouseFlags 位标志类型
type MouseFlags uint32

const (
	LBUTTON MouseFlags = 0x01
	RBUTTON MouseFlags = 0x02
	SHIFT   MouseFlags = 0x04
	CONTROL MouseFlags = 0x08
	MBUTTON MouseFlags = 0x10
)

// KeyFlags 位标志类型
type KeyFlags uint32

const (
	EXTENDED KeyFlags = 0x0100
	REPEAT   KeyFlags = 0x4000
)

// MouseMsg 枚举类型，在Go中通常使用int32或uint32代替枚举
type MouseMsg uint32

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

// JsExecState 定义为void*的别名，在Go中通常使用uintptr
type JsExecState uintptr

type JsValue uintptr

// _tagWkeWebView 结构体定义
// type _tagWkeWebView struct{}
type WebView uintptr

// _tagWkeString 结构体定义
// type _tagWkeString struct{}
type WkeString uintptr

// type _tagWkeMediaPlayer struct{}
type MediaPlayer uintptr

// wkeMediaPlayerClient
// type _tagWkeMediaPlayerClient struct{}
type MediaPlayerClient uintptr

// blinkWebURLRequestPtr
// type _tabblinkWebURLRequestPtr struct{}
type blinkWebURLRequestPtr uintptr

// 枚举类型通常使用int来表示
type ProxyType int32

const (
	PROXY_NONE ProxyType = iota
	PROXY_HTTP
	PROXY_SOCKS4
	PROXY_SOCKS4A
	PROXY_SOCKS5
	PROXY_SOCKS5HOSTNAME
)

// 结构体类型使用struct关键字定义
type Proxy struct {
	Type     ProxyType
	Hostname [100]byte // 在Go中，字符串通常不是定长的，但这里为了匹配C的类型我们使用[100]byte
	Port     uint16
	Username [50]byte
	Password [50]byte
}

// 枚举类型用于位掩码
type SettingMask uint32

const (
	SETTING_PROXY     SettingMask = 1
	SETTING_EXTENSION             = 1 << 2
)

// 结构体
type Settings struct {
	Proxy     Proxy
	Mask      SettingMask
	Extension *byte // 在cgo中，通常使用*C.char来表示C的const char*
}

// 另一个结构体
type ViewSettings struct {
	Size    int32
	BgColor uint32
}

// 结构体，表示地理位置信息
type GeolocationPosition struct {
	Timestamp                float64
	Latitude, Longitude      float64
	Accuracy                 float64
	ProvidesAltitude         bool
	Altitude                 float64
	ProvidesAltitudeAccuracy bool
	AltitudeAccuracy         float64
	ProvidesHeading          bool
	Heading                  float64
	ProvidesSpeed            bool
	Speed                    float64
}

// 对于void*类型的WkeWebFrameHandle，在Go中通常使用uintptr来表示
type WebFrameHandle uintptr

// 枚举类型通常使用int来表示
type MenuItemId int32

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
	MenuSaveImageId        MenuItemId = 1 << 12
)

// 文件操作函数类型
type FILE_OPEN func(path string) (void uintptr)
type FILE_CLOSE func(handle uintptr) (void uintptr)
type FILE_SIZE func(handle uintptr) (size uintptr)
type FILE_READ func(handle uintptr, buffer uintptr, size uintptr) int32
type FILE_SEEK func(handle uintptr, offset int, origin int) int32

// 文件存在检查函数类型
type EXISTS_FILE func(path string) bool

// 回调函数类型
type ON_TITLE_CHANGED func(clientHandler *ClientHandler, title WkeString) (void uintptr)
type ON_URL_CHANGED func(clientHandler *ClientHandler, url WkeString) (void uintptr)

// 客户端处理结构
type ClientHandler struct {
	OnTitleChanged ON_TITLE_CHANGED
	OnURLChanged   ON_URL_CHANGED
}

// 访问cookie的函数类型
type CookieVisitor func(
	params uintptr,
	name, value, domain, path *byte, // 注意：在Go中通常使用*byte来表示C中的char*
	secure, httpOnly int32,
	expires *int32, // 注意：在Go中我们使用*int来接收一个可能被修改的指针参数
) bool

// 枚举类型在Go中通常用int来表示
type CookieCommand int32

const (
	CookieCommandClearAllCookies CookieCommand = iota
	CookieCommandClearSessionCookies
	CookieCommandFlushCookiesToFile
	CookieCommandReloadCookiesFromFile
)

type NavigationType int32

const (
	NAVIGATION_TYPE_LINKCLICK    NavigationType = iota
	NAVIGATION_TYPE_FORMSUBMITTE                // 注意：这里可能是拼写错误，原始应为NAVIGATION_TYPE_FORMSUBMITTED
	NAVIGATION_TYPE_BACKFORWARD
	NAVIGATION_TYPE_RELOAD
	NAVIGATION_TYPE_FORMRESUBMITT // 注意：这里也可能是拼写错误，原始应为NAVIGATION_TYPE_FORMRESUBMITTED
	NAVIGATION_TYPE_OTHER
)

type CursorInfoType int32

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

// 结构体类型在Go中直接使用struct关键字定义
type WindowFeatures struct {
	X                  int32
	Y                  int32
	Width              int32
	Height             int32
	MenuBarVisible     bool
	StatusBarVisible   bool
	ToolBarVisible     bool
	LocationBarVisible bool
	ScrollbarsVisible  bool
	Resizable          bool
	Fullscreen         bool
}

type StorageType int32

const (
	StorageTypeString         StorageType = iota // WkeString data with an associated MIME type. Depending on the MIME type, there may be optional metadata attributes as well.
	StorageTypeFilename                          // Stores the name of one file being dragged into the renderer.
	StorageTypeBinaryData                        // An image being dragged out of the renderer. Contains a buffer holding the image data as well as the suggested name for saving the image to.
	StorageTypeFileSystemFile                    // Stores the filesystem URL of one file being dragged into the renderer.
)

type MemBuf struct {
	unuse  int32
	data   uintptr
	length uintptr // Note: using uintptr for uintptr
}

type WebDragDataItem struct {
	StorageType        StorageType
	StringType         *MemBuf // Only valid when storageType == StorageTypeString.
	StringData         *MemBuf
	FilenameData       *MemBuf // Only valid when storageType == StorageTypeFilename.
	DisplayNameData    *MemBuf
	BinaryData         *MemBuf // Only valid when storageType == StorageTypeBinaryData.
	Title              *MemBuf // Title associated with a link when stringType == "text/uri-list". Filename when storageType == StorageTypeBinaryData.
	FileSystemURL      *MemBuf // Only valid when storageType == StorageTypeFileSystemFile.
	FileSystemFileSize int64
	BaseURL            *MemBuf // Only valid when stringType == "text/html".
}

type WebDragData struct {
	ItemList         *WebDragDataItem
	ItemListLength   int32
	ModifierKeyState int32 // State of Shift/Ctrl/Alt/Meta keys.
	FilesystemId     *MemBuf
}

type WebDragOperationsMask uintptr // 假设wkeWebDragOperation是uintptr或其他整数类型
const (
	WebDragOperationNone    WebDragOperationsMask = 0
	WebDragOperationCopy    WebDragOperationsMask = 1
	WebDragOperationLink    WebDragOperationsMask = 2
	WebDragOperationGeneric WebDragOperationsMask = 4
	WebDragOperationPrivate WebDragOperationsMask = 8
	WebDragOperationMove    WebDragOperationsMask = 16
	WebDragOperationDelete  WebDragOperationsMask = 32
	WebDragOperationEvery   WebDragOperationsMask = 0xffffffff
)

type ResourceType int32

const (
	RESOURCE_TYPE_MAIN_FRAME     ResourceType = iota // top level page
	RESOURCE_TYPE_SUB_FRAME                          // frame or iframe
	RESOURCE_TYPE_STYLESHEET                         // a CSS stylesheet
	RESOURCE_TYPE_SCRIPT                             // an external script
	RESOURCE_TYPE_IMAGE                              // an image (jpg/gif/png/etc)
	RESOURCE_TYPE_FONT_RESOURCE                      // a font
	RESOURCE_TYPE_SUB_RESOURCE                       // an "other" subresource.
	RESOURCE_TYPE_OBJECT                             // an object (or embed) tag for a plugin, or a resource that a plugin requested.
	RESOURCE_TYPE_MEDIA                              // a media resource.
	RESOURCE_TYPE_WORKER                             // the main resource of a dedicated worker.
	RESOURCE_TYPE_SHARED_WORKER                      // the main resource of a shared worker.
	RESOURCE_TYPE_PREFETCH                           // an explicitly requested prefetch
	RESOURCE_TYPE_FAVICON                            // a favicon
	RESOURCE_TYPE_XHR                                // a XMLHttpRequest
	RESOURCE_TYPE_PING                               // a ping request for <a ping>
	RESOURCE_TYPE_SERVICE_WORKER                     // the main resource of a service worker.
	RESOURCE_TYPE_LAST_TYPE      = RESOURCE_TYPE_SERVICE_WORKER
)

type WillSendRequestInfo struct {
	Url              WkeString
	NewUrl           WkeString
	ResourceType     ResourceType
	HttpResponseCode int32
	Method           WkeString
	Referrer         WkeString
	Headers          uintptr // void*在Go中通常转换为uintptr
}

type HttBodyElementType int32

const (
	HttBodyElementTypeData HttBodyElementType = iota
	HttBodyElementTypeFile
)

type PostBodyElement struct {
	Size       int32
	Type       HttBodyElementType
	Data       *MemBuf // 假设WkeMemBuf*是指向WkeMemBuf的指针
	FilePath   WkeString
	FileStart  int64 // __int64在Go中通常为int64
	FileLength int64 // -1 means to the end of the file.
}

type PostBodyElements struct {
	Size        int32
	Element     **PostBodyElement
	ElementSize uintptr // size_t在Go中通常为uintptr
	IsDirty     bool
}
type NetJob uintptr

type Slist struct {
	Data *byte
	Next *Slist
}

type TempCallbackInfo struct {
	Size                int32
	Frame               WebFrameHandle
	WillSendRequestInfo *WillSendRequestInfo
	Url                 string
	PostBody            *PostBodyElements
	Job                 NetJob
}

type RequestType int32

const (
	RequestTypeInvalidation RequestType = iota
	RequestTypeGet
	RequestTypePost
	RequestTypePut
)

type PdfDatas struct {
	Count int32
	Sizes *uintptr
	Datas **uintptr
}

type BOOL int32

type PrintSettings struct {
	StructSize               int32
	Dpi                      int32
	Width                    int32 // in px
	Height                   int32
	MarginTop                int32
	MarginBottom             int32
	MarginLeft               int32
	MarginRight              int32
	IsPrintPageHeadAndFooter BOOL
	IsPrintBackgroud         BOOL
	IsLandscape              BOOL
	IsPrintToMultiPage       BOOL
}

type ScreenshotSettings struct {
	StructSize int32
	Width      int32
	Height     int32
}

type HDC uintptr
type HWND uintptr

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

type COLORREF DWORD

// wkeCaretChangedCallback
type CaretChangedCallback func(webView WebView, param uintptr, r *Rect) (void uintptr)

// wkeTitleChangedCallback
type TitleChangedCallback func(webView WebView, param uintptr, title WkeString) (void uintptr)

// wkeURLChangedCallback
type URLChangedCallback func(webView WebView, param uintptr, url WkeString) (void uintptr)

// wkeURLChangedCallback2
type URLChangedCallback2 func(webView WebView, param uintptr, frameId WebFrameHandle, url WkeString) (void uintptr)

// wkePaintUpdatedCallback
type PaintUpdatedCallback func(webView WebView, param uintptr, hdc HDC, x, y, cx, cy int32) (void uintptr)

// wkePaintBitUpdatedCallback
type PaintBitUpdatedCallback func(webView WebView, param uintptr, buffer uintptr, r *Rect, width, height int32) (void uintptr)

// wkeAlertBoxCallback
type AlertBoxCallback func(webView WebView, param uintptr, msg WkeString) (void uintptr)

// wkeConfirmBoxCallback
type ConfirmBoxCallback func(webView WebView, param uintptr, msg WkeString) bool

// wkePromptBoxCallback
type PromptBoxCallback func(webView WebView, param uintptr, msg, defaultResult WkeString) (result WkeString, ok bool)

// wkeNavigationCallback
type NavigationCallback func(webView WebView, param uintptr, navigationType NavigationType, url WkeString) bool

// wkeCreateViewCallback
type CreateViewCallback func(webView WebView, param uintptr, navigationType NavigationType, url WkeString, windowFeatures *WindowFeatures) WebView

// wkeDocumentReadyCallback
type DocumentReadyCallback func(webView WebView, param uintptr) (void uintptr)

// wkeDocumentReady2Callback
type DocumentReady2Callback func(webView WebView, param uintptr, frameId WebFrameHandle) (void uintptr)

type OnShowDevtoolsCallback func(webView WebView, param uintptr) (void uintptr)

type NodeOnCreateProcessCallback func(webView WebView, param uintptr, applicationPath *uint16, arguments *uint16, startup *STARTUPINFOW) (void uintptr)
type OnPluginFindCallback func(webView WebView, param uintptr, mime *byte, initializeFunc uintptr, getEntryPointsFunc uintptr, shutdownFunc uintptr) (void uintptr)

type OnPrintCallback func(webView WebView, param uintptr, frameId WebFrameHandle, printParams uintptr) (void uintptr)
type OnScreenshotCallback func(webView WebView, param uintptr, data *byte, size uintptr) (void uintptr)

type ImageBufferToDataURL func(webView WebView, param uintptr, data *byte, size uintptr) WkeString

type MediaLoadInfo struct {
	Size     int32
	Width    int32
	Height   int32
	Duration float64
}
type WillMediaLoadCallback func(webView WebView, param uintptr, url *byte, info *MediaLoadInfo) (void uintptr)

type StartDraggingCallback func(
	webView WebView,
	param uintptr,
	frame WebFrameHandle,
	data *WebDragData,
	mask WebDragOperationsMask,
	image uintptr,
	dragImageOffset *Point,
)

type UiThreadRunCallback func(hWnd HWND, param uintptr) (void uintptr)
type UiThreadPostTaskCallback func(hWnd HWND, callback UiThreadRunCallback, param uintptr) int

type OtherLoadType int

const (
	DID_START_LOADING OtherLoadType = iota
	DID_STOP_LOADING
	DID_NAVIGATE
	DID_NAVIGATE_IN_PAGE
	DID_GET_RESPONSE_DETAILS
	DID_GET_REDIRECT_REQUEST
	DID_POST_REQUEST
)

type OnOtherLoadCallback func(webView WebView, param uintptr, type_ OtherLoadType, info *TempCallbackInfo) (void uintptr)

type OnContextMenuItemClickType int

const kWkeContextMenuItemClickTypePrint OnContextMenuItemClickType = 0x01

type OnContextMenuItemClickStep int

const (
	kWkeContextMenuItemClickStepShow OnContextMenuItemClickStep = 0x01
	kWkeContextMenuItemClickStepClick
)

type OnContextMenuItemClickCallback func(
	webView WebView,
	param uintptr,
	type_ OnContextMenuItemClickType,
	step OnContextMenuItemClickStep,
	frameId WebFrameHandle,
	info uintptr,
) bool

type LoadingResult int

const (
	LOADING_SUCCEEDED LoadingResult = iota
	LOADING_FAILED
	LOADING_CANCELED
)

type DownloadOpt int

const (
	kWkeDownloadOptCancel DownloadOpt = iota
	kWkeDownloadOptCacheData
)

type NetJobDataRecvCallback func(ptr uintptr, job NetJob, data *byte, length int) (void uintptr)
type NetJobDataFinishCallback func(ptr uintptr, job NetJob, result LoadingResult) (void uintptr)

type NetJobDataBind struct {
	Param          uintptr
	RecvCallback   NetJobDataRecvCallback
	FinishCallback NetJobDataFinishCallback
}

type PopupDialogSaveNameCallback func(ptr uintptr, filePath *uint16) (void uintptr)

type DownloadBind struct {
	Param            uintptr
	RecvCallback     NetJobDataRecvCallback
	FinishCallback   NetJobDataFinishCallback
	SaveNameCallback PopupDialogSaveNameCallback
}

type DialogProperties int

const (
	kWkeDialogPropertiesOpenFile                DialogProperties = 1 << 1 // 允许选择文件
	kWkeDialogPropertiesOpenDirectory           DialogProperties = 1 << 2 // 允许选择文件夹
	kWkeDialogPropertiesMultiSelections         DialogProperties = 1 << 3 // 允许多选。
	kWkeDialogPropertiesShowHiddenFiles         DialogProperties = 1 << 4 // 显示对话框中的隐藏文件。
	kWkeDialogPropertiesCreateDirectory         DialogProperties = 1 << 5 // macOS - 允许你通过对话框的形式创建新的目录。
	kWkeDialogPropertiesPromptToCreate          DialogProperties = 1 << 6 // Windows - 如果输入的文件路径在对话框中不存在, 则提示创建。
	kWkeDialogPropertiesNoResolveAliases        DialogProperties = 1 << 7 // macOS - 禁用自动的别名路径(符号链接) 解析。
	kWkeDialogPropertiesTreatPackageAsDirectory DialogProperties = 1 << 8 // macOS - 将包(如.app 文件夹) 视为目录而不是文件。
	kWkeDialogPropertiesDontAddToRecent         DialogProperties = 1 << 9 // Windows - 不要将正在打开的项目添加到最近的文档列表中。
)

type FileFilter struct {
	Name       string // 例如"image"、"Movies"
	Extensions string // 例如"jpg|png|gif"
}

type DialogOptions struct {
	Magic                   int32
	Title                   string
	DefaultPath             string
	ButtonLabel             string
	Filters                 *FileFilter
	FiltersCount            int32
	Prop                    DialogProperties
	Message                 string
	SecurityScopedBookmarks bool
}

type DownloadInBlinkThreadCallback func(
	webView WebView,
	param uintptr,
	expectedContentLength uintptr,
	url string,
	mime string,
	disposition string,
	job NetJob,
	dataBind *NetJobDataBind,
) DownloadOpt

type LoadingFinishCallback func(
	webView WebView,
	param uintptr,
	url WkeString,
	result LoadingResult,
	failedReason WkeString,
)

type DownloadCallback func(
	webView WebView,
	param uintptr,
	url string,
) bool

type Download2Callback func(
	webView WebView,
	param uintptr,
	expectedContentLength uintptr,
	url string,
	mime string,
	disposition string,
	job NetJob,
	dataBind *NetJobDataBind,
) DownloadOpt

type ConsoleLevel int

const (
	LevelDebug        ConsoleLevel = 4
	LevelLog          ConsoleLevel = 1
	LevelInfo         ConsoleLevel = 5
	LevelWarning      ConsoleLevel = 2
	LevelError        ConsoleLevel = 3
	LevelRevokedError ConsoleLevel = 6
	LevelLast         ConsoleLevel = LevelInfo
)

type ConsoleCallback func(webView WebView, param uintptr, level ConsoleLevel, message WkeString, sourceName WkeString, sourceLine uint, stackTrace WkeString) (void uintptr)

type OnCallUiThread func(webView WebView, paramOnInThread uintptr) (void uintptr)
type CallUiThread func(webView WebView, func_ OnCallUiThread, param uintptr) (void uintptr)

type MediaPlayerFactory func(webView WebView, client MediaPlayerClient, npBrowserFuncs uintptr, npPluginFuncs uintptr) MediaPlayer
type OnIsMediaPlayerSupportsMIMEType func(mime string) bool

// wkeNet
type WebUrlRequest struct{}
type WebUrlRequestPtr *WebUrlRequest
type WebUrlResponse struct{}
type WebUrlResponsePtr *WebUrlResponse

type OnUrlRequestWillRedirectCallback func(webView WebView, param uintptr, oldRequest WebUrlRequestPtr, request WebUrlRequestPtr, redirectResponse WebUrlResponsePtr) (void uintptr)
type OnUrlRequestDidReceiveResponseCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, response WebUrlResponsePtr) (void uintptr)
type OnUrlRequestDidReceiveDataCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, data *byte, dataLength int) (void uintptr)
type OnUrlRequestDidFailCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, error string) (void uintptr)
type OnUrlRequestDidFinishLoadingCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, finishTime float64) (void uintptr)

type UrlRequestCallbacks struct {
	WillRedirectCallback       OnUrlRequestWillRedirectCallback
	DidReceiveResponseCallback OnUrlRequestDidReceiveResponseCallback
	DidReceiveDataCallback     OnUrlRequestDidReceiveDataCallback
	DidFailCallback            OnUrlRequestDidFailCallback
	DidFinishLoadingCallback   OnUrlRequestDidFinishLoadingCallback
}

type LoadUrlBeginCallback func(webView WebView, param uintptr, url string, job NetJob) bool
type LoadUrlEndCallback func(webView WebView, param uintptr, url string, job NetJob, buf uintptr, len int) (void uintptr)
type LoadUrlFailCallback func(webView WebView, param uintptr, url string, job NetJob) (void uintptr)
type LoadUrlHeadersReceivedCallback func(webView WebView, param uintptr, url string, job NetJob) (void uintptr)
type LoadUrlFinishCallback func(webView WebView, param uintptr, url string, job NetJob, len int) (void uintptr)

type DidCreateScriptContextCallback func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, extensionGroup int, worldId int) (void uintptr)
type WillReleaseScriptContextCallback func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, worldId int) (void uintptr)
type NetResponseCallback func(webView WebView, param uintptr, url string, job NetJob) bool
type OnNetGetFaviconCallback func(webView WebView, param uintptr, url string, buf *MemBuf) (void uintptr)

type V8ContextPtr uintptr
type V8Isolate uintptr

// wke window
type WindowType int

const (
	WINDOW_TYPE_POPUP WindowType = iota
	WINDOW_TYPE_TRANSPARENT
	WINDOW_TYPE_CONTROL
)

type WindowCreateInfo struct {
	Size    int32
	Parent  HWND
	Style   DWORD
	StyleEx DWORD
	X       int32
	Y       int32
	Width   int32
	Height  int32
	Color   COLORREF
}

type DefaultPrinterSettings struct {
	StructSize         int32
	IsLandscape        bool
	IsPrintHeadFooter  bool
	IsPrintBackgroud   bool
	EdgeDistanceLeft   int32
	EdgeDistanceTop    int32
	EdgeDistanceRight  int32
	EdgeDistanceBottom int32
	Copies             int32
	PaperType          int32
}

type WindowClosingCallback func(webWindow WebView, param uintptr) bool
type WindowDestroyCallback func(webWindow WebView, param uintptr) (void uintptr)

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type DraggableRegion struct {
	Bounds    RECT
	Draggable bool
}
type DraggableRegionsChangedCallback func(webView WebView, param uintptr, rects *DraggableRegion, rectCount int)

type JsType int

const (
	JSTYPE_NUMBER JsType = iota
	JSTYPE_STRING
	JSTYPE_BOOLEAN
	JSTYPE_OBJECT
	JSTYPE_FUNCTION
	JSTYPE_UNDEFINED
	JSTYPE_ARRAY
	JSTYPE_NULL
)

// cexer JS对象、函数绑定支持
type JsGetPropertyCallback func(es JsExecState, object JsValue, propertyName *byte) JsValue
type JsSetPropertyCallback func(es JsExecState, object JsValue, propertyName *byte, value JsValue) bool
type JsCallAsFunctionCallback func(es JsExecState, object JsValue, args *[]JsValue, argCount int) JsValue
type JsFinalizeCallback func(data *JsData) (void uintptr)

type JsData struct {
	TypeName       [100]byte
	PropertyGet    JsGetPropertyCallback
	PropertySet    JsSetPropertyCallback
	Finalize       JsFinalizeCallback
	CallAsFunction JsCallAsFunctionCallback
}

type JsExceptionInfo struct {
	Message            string
	SourceLine         string
	ScriptResourceName string
	LineNumber         int32
	StartPosition      int32
	EndPosition        int32
	StartColumn        int32
	EndColumn          int32
	CallstackString    string
}

type JsKeys struct {
	Length uint
	Keys   **byte
}

type WPARAM uintptr
type LPARAM uintptr
type LRESULT uintptr

type wkeJsNativeFunction func(es JsExecState, param uintptr) JsValue
type wkeShutdown func()
type wkeShutdownForDebug func() // 测试使用，不了解千万别用！
type wkeVersion func() uint
type wkeVersionString func() string
type wkeGC func(webView WebView, intervalSec int64)
type wkeSetResourceGc func(webView WebView, intervalSec int64)
type wkeSetFileSystem func(pfnOpen FILE_OPEN, pfnClose FILE_CLOSE, pfnSize FILE_SIZE, pfnRead FILE_READ, pfnSeek FILE_SEEK)
type wkeWebViewName func(webView WebView) string
type wkeSetWebViewName func(webView WebView, name string)
type wkeIsLoaded func(webView WebView) bool
type wkeIsLoadFailed func(webView WebView) bool
type wkeIsLoadComplete func(webView WebView) bool
type wkeGetSource func(webView WebView) string
type wkeTitle func(webView WebView) string
type wkeTitleW func(webView WebView) []uint16
type wkeWidth func(webView WebView) int
type wkeHeight func(webView WebView) int
type wkeContentsWidth func(webView WebView) int
type wkeContentsHeight func(webView WebView) int
type wkeSelectAll func(webView WebView)
type wkeCopy func(webView WebView)
type wkeCut func(webView WebView)
type wkePaste func(webView WebView)
type wkeDelete func(webView WebView)
type wkeCookieEnabled func(webView WebView) bool
type wkeMediaVolume func(webView WebView) float32 // 假设float为float32，根据实际情况可能是float64
type wkeMouseEvent func(webView WebView, message uint, x int, y int, flags uint) bool
type wkeContextMenuEvent func(webView WebView, x int, y int, flags uint) bool
type wkeMouseWheel func(webView WebView, x int, y int, delta int, flags uint) bool
type wkeKeyUp func(webView WebView, virtualKeyCode uint, flags uint, systemKey bool) bool
type wkeKeyDown func(webView WebView, virtualKeyCode uint, flags uint, systemKey bool) bool
type wkeKeyPress func(webView WebView, virtualKeyCode uint, flags uint, systemKey bool) bool
type wkeFocus func(webView WebView)
type wkeUnfocus func(webView WebView)
type wkeGetCaret func(webView WebView) Rect
type wkeAwaken func(webView WebView)
type wkeZoomFactor func(webView WebView) float32 // 假设float为float32
type wkeSetClientHandler func(webView WebView, handler *ClientHandler)
type wkeGetClientHandler func(webView WebView) *ClientHandler
type jsToString func(es JsExecState, v JsValue) string
type jsToStringW func(es JsExecState, v JsValue) []uint16
type wkeSetViewSettings func(webView WebView, settings *ViewSettings)
type wkeSetDebugConfig func(webView WebView, debugString string, param string)
type wkeToString func(wkeStr WkeString) string
type wkeToStringW func(wkeStr WkeString) []uint16
type wkeGetDebugConfig func(webView WebView, debugString string) uintptr // 假设返回void*对应Go的uintptr
type wkeIsInitialize func() bool
type wkeFinalize func()
type wkeUpdate func()
type wkeGetVersion func() uint
type wkeGetVersionString func() string
type wkeCreateWebView func() WebView
type wkeDestroyWebView func(webView WebView)
type wkeSetMemoryCacheEnable func(webView WebView, b bool)
type wkeSetMouseEnabled func(webView WebView, b bool)
type wkeSetTouchEnabled func(webView WebView, b bool)
type wkeSetSystemTouchEnabled func(webView WebView, b bool)
type wkeSetContextMenuEnabled func(webView WebView, b bool)
type wkeSetNavigationToNewWindowEnable func(webView WebView, b bool)
type wkeSetCspCheckEnable func(webView WebView, b bool)
type wkeSetNpapiPluginsEnabled func(webView WebView, b bool)
type wkeSetHeadlessEnabled func(webView WebView, b bool) // 可以关闭渲染
type wkeSetDragEnable func(webView WebView, b bool)      // 可关闭拖拽文件加载网页
type wkeSetDragDropEnable func(webView WebView, b bool)  // 可关闭拖拽到其他进程
type wkeSetLanguage func(webView WebView, language string)
type wkeSetViewNetInterface func(webView WebView, netInterface string)
type wkeSetContextMenuItemShow func(webView WebView, item MenuItemId, isShow bool) // 设置某项menu是否显示
type wkeSetProxy func(proxy *Proxy)
type wkeGetName func(webView WebView) string // 假设char为*C.char或其他等效类型
type wkeIsTransparent func(webView WebView) bool
type wkeGetUserAgent func(webView WebView) string // 假设char为*C.char或其他等效类型
type wkeSetViewProxy func(webView WebView, proxy *Proxy)
type wkeSetName func(webView WebView, name string) // 假设char为*C.char或其他等效类型
type wkeSetHandle func(webView WebView, wnd HWND)
type wkeSetTransparent func(webView WebView, transparent bool)
type wkeSetUserAgent func(webView WebView, userAgent string)
type wkeSetUserAgentW func(webView WebView, userAgent []uint16)
type wkeSetHandleOffset func(webView WebView, x int, y int)
type wkeShowDevtools func(webView WebView, path []uint16, callback OnShowDevtoolsCallback, param uintptr) // 假设使用uintptr表示void*
type wkeLoadW func(webView WebView, url []uint16)
type wkeLoadURL func(webView WebView, url string)
type wkeLoadURLW func(webView WebView, url []uint16)
type wkePostURL func(wkeView WebView, url string, postData []byte, postLen int)
type wkePostURLW func(wkeView WebView, url []uint16, postData []byte, postLen int)
type wkeLoadHTML func(webView WebView, html string)
type wkeLoadHtmlWithBaseUrl func(webView WebView, html string, baseUrl string)
type wkeLoadHTMLW func(webView WebView, html []uint16)
type wkeLoadFile func(webView WebView, filename string)
type wkeLoadFileW func(webView WebView, filename []uint16)
type wkeGetURL func(webView WebView) string
type wkeGetFrameUrl func(webView WebView, frameId WebFrameHandle) string
type wkeIsLoading func(webView WebView) bool
type wkeIsLoadingSucceeded func(webView WebView) bool
type wkeIsLoadingFailed func(webView WebView) bool
type wkeIsLoadingCompleted func(webView WebView) bool
type wkeIsDocumentReady func(webView WebView) bool
type wkeStopLoading func(webView WebView)
type wkeReload func(webView WebView)
type wkeGoToOffset func(webView WebView, offset int)
type wkeGoToIndex func(webView WebView, index int)
type wkeGetWebviewId func(webView WebView) int
type wkeIsWebviewAlive func(id int) bool
type wkeIsWebviewValid func(webView WebView) bool
type wkeGetDocumentCompleteURL func(webView WebView, frameId WebFrameHandle, partialURL string) string
type wkeCreateMemBuf func(webView WebView, buf uintptr, length uintptr) *MemBuf
type wkeFreeMemBuf func(buf *MemBuf)
type wkeGetTitle func(webView WebView) string
type wkeGetTitleW func(webView WebView) []uint16
type wkeResize func(webView WebView, w int, h int)
type wkeGetWidth func(webView WebView) int
type wkeGetHeight func(webView WebView) int
type wkeGetContentWidth func(webView WebView) int
type wkeGetContentHeight func(webView WebView) int
type wkeSetDirty func(webView WebView, dirty bool)
type wkeIsDirty func(webView WebView) bool
type wkeAddDirtyArea func(webView WebView, x int, y int, w int, h int)
type wkeLayoutIfNeeded func(webView WebView)
type wkePaint2 func(webView WebView, bits uintptr, bufWid int, bufHei int, xDst int, yDst int, w int, h int, xSrc int, ySrc int, bCopyAlpha bool)
type wkePaint func(webView WebView, bits uintptr, pitch int)
type wkeRepaintIfNeeded func(webView WebView)
type wkeGetViewDC func(webView WebView) HDC
type wkeUnlockViewDC func(webView WebView)
type wkeGetHostHWND func(webView WebView) HWND
type wkeCanGoBack func(webView WebView) bool
type wkeGoBack func(webView WebView) bool
type wkeCanGoForward func(webView WebView) bool
type wkeGoForward func(webView WebView) bool
type wkeNavigateAtIndex func(webView WebView, index int) bool
type wkeGetNavigateIndex func(webView WebView) int
type wkeEditorSelectAll func(webView WebView)
type wkeEditorUnSelect func(webView WebView)
type wkeEditorCopy func(webView WebView)
type wkeEditorCut func(webView WebView)
type wkeEditorPaste func(webView WebView)
type wkeEditorDelete func(webView WebView)
type wkeEditorUndo func(webView WebView)
type wkeEditorRedo func(webView WebView)
type wkeGetCookieW func(webView WebView) []uint16
type wkeGetCookie func(webView WebView) string
type wkeSetCookie func(webView WebView, url string, cookie string)
type wkeVisitAllCookie func(webView WebView, params uintptr, visitor CookieVisitor)
type wkePerformCookieCommand func(webView WebView, command CookieCommand)
type wkeSetCookieEnabled func(webView WebView, enable bool)
type wkeIsCookieEnabled func(webView WebView) bool
type wkeSetCookieJarPath func(webView WebView, path []uint16)
type wkeSetCookieJarFullPath func(webView WebView, path []uint16)
type wkeClearCookie func(webView WebView)
type wkeSetLocalStorageFullPath func(webView WebView, path []uint16)
type wkeAddPluginDirectory func(webView WebView, path []uint16)
type wkeSetMediaVolume func(webView WebView, volume float32)
type wkeGetMediaVolume func(webView WebView) float32
type wkeFireMouseEvent func(webView WebView, message uint, x int, y int, flags uint) bool
type wkeFireContextMenuEvent func(webView WebView, x int, y int, flags uint) bool
type wkeFireMouseWheelEvent func(webView WebView, x int, y int, delta int, flags uint) bool
type wkeFireKeyUpEvent func(webView WebView, virtualKeyCode uint, flags uint, systemKey bool) bool
type wkeFireKeyDownEvent func(webView WebView, virtualKeyCode uint, flags uint, systemKey bool) bool
type wkeFireKeyPressEvent func(webView WebView, charCode uint, flags uint, systemKey bool) bool
type wkeFireWindowsMessage func(webView WebView, hWnd HWND, message uint, wParam WPARAM, lParam LPARAM, result *LRESULT) bool
type wkeSetFocus func(webView WebView)
type wkeKillFocus func(webView WebView)
type wkeGetCaretRect func(webView WebView) Rect
type wkeGetCaretRect2 func(webView WebView) *Rect // 给一些不方便获取返回结构体的语言调用
type wkeRunJS func(webView WebView, script string) JsValue
type wkeRunJSW func(webView WebView, script []uint16) JsValue
type wkeGlobalExec func(webView WebView) JsExecState
type wkeGetGlobalExecByFrame func(webView WebView, frameId WebFrameHandle) JsExecState
type wkeSleep func(webView WebView)
type wkeWake func(webView WebView)
type wkeIsAwake func(webView WebView) bool
type wkeSetZoomFactor func(webView WebView, factor float32)
type wkeGetZoomFactor func(webView WebView) float32
type wkeEnableHighDPISupport func()
type wkeSetEditable func(webView WebView, editable bool)
type wkeGetString func(wkeStr WkeString) string
type wkeGetStringW func(wkeStr WkeString) []uint16
type wkeSetString func(wkeStr WkeString, str *byte, len uintptr)
type wkeSetStringWithoutNullTermination func(wkeStr WkeString, str *byte, len uintptr)
type wkeSetStringW func(wkeStr WkeString, str *uint16, len uintptr)
type wkeCreateString func(str *byte, len uintptr) WkeString
type wkeCreateStringW func(str *uint16, len uintptr) WkeString
type wkeCreateStringWithoutNullTermination func(str *byte, len uintptr) WkeString
type wkeGetStringLen func(wkeStr WkeString) uintptr
type wkeDeleteString func(wkeStr WkeString)
type wkeGetWebViewForCurrentContext func() WebView
type wkeSetUserKeyValue func(webView WebView, key string, value uintptr)
type wkeGetUserKeyValue func(webView WebView, key string) uintptr
type wkeGetCursorInfoType func(webView WebView) int
type wkeSetCursorInfoType func(webView WebView, type_ int)
type wkeSetDragFiles func(webView WebView, clintPos *Point, screenPos *Point, files *WkeString, filesCount int)
type wkeSetDeviceParameter func(webView WebView, device string, paramStr string, paramInt int, paramFloat float32)
type wkeGetTempCallbackInfo func(webView WebView) *TempCallbackInfo
type wkeOnCaretChanged func(webView WebView, callback CaretChangedCallback, callbackParam uintptr)
type wkeOnMouseOverUrlChanged func(webView WebView, callback TitleChangedCallback, callbackParam uintptr)
type wkeOnTitleChanged func(webView WebView, callback TitleChangedCallback, callbackParam uintptr)
type wkeOnURLChanged func(webView WebView, callback URLChangedCallback, callbackParam uintptr)
type wkeOnURLChanged2 func(webView WebView, callback URLChangedCallback2, callbackParam uintptr)
type wkeOnPaintUpdated func(webView WebView, callback PaintUpdatedCallback, callbackParam uintptr)
type wkeOnPaintBitUpdated func(webView WebView, callback PaintBitUpdatedCallback, callbackParam uintptr)
type wkeOnAlertBox func(webView WebView, callback AlertBoxCallback, callbackParam uintptr)
type wkeOnConfirmBox func(webView WebView, callback ConfirmBoxCallback, callbackParam uintptr)
type wkeOnPromptBox func(webView WebView, callback PromptBoxCallback, callbackParam uintptr)
type wkeOnNavigation func(webView WebView, callback NavigationCallback, param uintptr)
type wkeOnCreateView func(webView WebView, callback CreateViewCallback, param uintptr)
type wkeOnDocumentReady func(webView WebView, callback DocumentReadyCallback, param uintptr)
type wkeOnDocumentReady2 func(webView WebView, callback DocumentReady2Callback, param uintptr)
type wkeOnLoadingFinish func(webView WebView, callback LoadingFinishCallback, param uintptr)
type wkeOnDownload func(webView WebView, callback DownloadCallback, param uintptr)
type wkeOnDownload2 func(webView WebView, callback Download2Callback, param uintptr)
type wkeOnConsole func(webView WebView, callback ConsoleCallback, param uintptr)
type wkeSetUIThreadCallback func(webView WebView, callback CallUiThread, param uintptr)
type wkeOnLoadUrlBegin func(webView WebView, callback LoadUrlBeginCallback, callbackParam uintptr)
type wkeOnLoadUrlEnd func(webView WebView, callback LoadUrlEndCallback, callbackParam uintptr)
type wkeOnLoadUrlHeadersReceived func(webView WebView, callback LoadUrlHeadersReceivedCallback, callbackParam uintptr)
type wkeOnLoadUrlFinish func(webView WebView, callback LoadUrlFinishCallback, callbackParam uintptr)
type wkeOnLoadUrlFail func(webView WebView, callback LoadUrlFailCallback, callbackParam uintptr)
type wkeOnDidCreateScriptContext func(webView WebView, callback DidCreateScriptContextCallback, callbackParam uintptr)
type wkeOnWillReleaseScriptContext func(webView WebView, callback WillReleaseScriptContextCallback, callbackParam uintptr)
type wkeOnWindowClosing func(webWindow WebView, callback WindowClosingCallback, param uintptr)
type wkeOnWindowDestroy func(webWindow WebView, callback WindowDestroyCallback, param uintptr)
type wkeOnDraggableRegionsChanged func(webView WebView, callback DraggableRegionsChangedCallback, param uintptr)
type wkeOnWillMediaLoad func(webView WebView, callback WillMediaLoadCallback, param uintptr)
type wkeOnStartDragging func(webView WebView, callback StartDraggingCallback, param uintptr)
type wkeOnPrint func(webView WebView, callback OnPrintCallback, param uintptr)
type wkeScreenshot func(webView WebView, settings *ScreenshotSettings, callback OnScreenshotCallback, param uintptr)
type wkeOnOtherLoad func(webView WebView, callback OnOtherLoadCallback, param uintptr)
type wkeOnContextMenuItemClick func(webView WebView, callback OnContextMenuItemClickCallback, param uintptr)
type wkeIsProcessingUserGesture func(webView WebView) bool
type wkeNetSetMIMEType func(jobPtr NetJob, type_ string)
type wkeNetGetMIMEType func(jobPtr NetJob, mime *WkeString) string
type wkeNetGetReferrer func(jobPtr NetJob) string
type wkeNetSetHTTPHeaderField func(jobPtr NetJob, key, value uint16, response bool)
type wkeNetGetHTTPHeaderField func(jobPtr NetJob, key string) string
type wkeNetGetHTTPHeaderFieldFromResponse func(jobPtr NetJob, key string) string

// 调用此函数后,网络层收到数据会存储在一buf内,接收数据完成后响应OnLoadUrlEnd事件.#此调用严重影响性能,慎用
// 此函数和wkeNetSetData的区别是，wkeNetHookRequest会在接受到真正网络数据后再调用回调，并允许回调修改网络数据。
// 而wkeNetSetData是在网络数据还没发送的时候修改
type wkeNetSetData func(jobPtr NetJob, buf *byte, len int)
type wkeNetHookRequest func(jobPtr NetJob)
type wkeNetGetRequestMethod func(jobPtr NetJob) RequestType
type wkeNetContinueJob func(jobPtr NetJob)
type wkeNetGetUrlByJob func(jobPtr NetJob) string
type wkeNetGetRawHttpHead func(jobPtr NetJob) Slist
type wkeNetGetRawResponseHead func(jobPtr NetJob) Slist
type wkeNetCancelRequest func(jobPtr NetJob)
type wkeNetHoldJobToAsynCommit func(jobPtr NetJob) bool
type wkeNetOnResponse func(webView WebView, callback NetResponseCallback, param uintptr)
type wkeNetGetFavicon func(webView WebView, callback OnNetGetFaviconCallback, param uintptr) int
type wkeNetCreateWebUrlRequest func(url, method, mime string) WebUrlRequestPtr
type wkeNetCreateWebUrlRequest2 func(request blinkWebURLRequestPtr) WebUrlRequestPtr
type wkeNetCopyWebUrlRequest func(jobPtr NetJob, needExtraData bool) blinkWebURLRequestPtr
type wkeNetDeleteBlinkWebURLRequestPtr func(request blinkWebURLRequestPtr)
type wkeNetAddHTTPHeaderFieldToUrlRequest func(request WebUrlRequestPtr, name, value string)
type wkeNetStartUrlRequest func(webView WebView, request WebUrlRequestPtr, param uintptr, callbacks *UrlRequestCallbacks) int
type wkeNetGetHttpStatusCode func(response WebUrlResponsePtr) int
type wkeNetGetExpectedContentLength func(response WebUrlResponsePtr) int64
type wkeNetGetResponseUrl func(response WebUrlResponsePtr) string
type wkeNetCancelWebUrlRequest func(requestId int)
type wkeNetGetPostBody func(jobPtr NetJob) *PostBodyElements
type wkeNetFreePostBodyElements func(elements *PostBodyElements)
type wkeNetCreatePostBodyElements func(webView WebView, length uintptr) *PostBodyElements
type wkeNetCreatePostBodyElement func(webView WebView) *PostBodyElement
type wkeNetFreePostBodyElement func(element *PostBodyElement)
type wkePopupDialogAndDownload func(
	webviewHandle WebView,
	dialogOpt *DialogOptions,
	expectedContentLength uintptr,
	url string,
	mime string,
	disposition string,
	job NetJob,
	dataBind *NetJobDataBind,
	callbackBind *DownloadBind,
) DownloadOpt
type wkeDownloadByPath func(
	webviewHandle WebView,
	dialogOpt *DialogOptions,
	path []uint16,
	expectedContentLength uintptr,
	url string,
	mime string,
	disposition string,
	job NetJob,
	dataBind *NetJobDataBind,
	callbackBind *DownloadBind,
) DownloadOpt
type wkeIsMainFrame func(webView WebView, frameId WebFrameHandle) bool
type wkeIsWebRemoteFrame func(webView WebView, frameId WebFrameHandle) bool
type wkeWebFrameGetMainFrame func(webView WebView) WebFrameHandle
type wkeRunJsByFrame func(webView WebView, frameId WebFrameHandle, script string, isInClosure bool) JsValue
type wkeInsertCSSByFrame func(webView WebView, frameId WebFrameHandle, cssText string)
type wkeWebFrameGetMainWorldScriptContext func(webView WebView, webFrameId WebFrameHandle, contextOut *V8ContextPtr)
type wkeGetBlinkMainThreadIsolate func() V8Isolate
type wkeCreateWebWindow func(
	typ WindowType,
	parent HWND,
	x, y, width, height int,
) WebView
type wkeCreateWebCustomWindow func(info *WindowCreateInfo) WebView
type wkeDestroyWebWindow func(webWindow WebView)
type wkeGetWindowHandle func(webWindow WebView) HWND
type wkeShowWindow func(webWindow WebView, show bool)
type wkeEnableWindow func(webWindow WebView, enable bool)
type wkeMoveWindow func(webWindow WebView, x, y, width, height int)
type wkeMoveToCenter func(webWindow WebView)
type wkeResizeWindow func(webWindow WebView, width, height int)
type wkeDragTargetDragEnter func(
	webView WebView,
	webDragData *WebDragData,
	clientPoint *Point,
	screenPoint *Point,
	operationsAllowed WebDragOperationsMask,
	modifiers int,
) WebDragOperationsMask
type wkeDragTargetDragOver func(
	webView WebView,
	clientPoint *Point,
	screenPoint *Point,
	operationsAllowed WebDragOperationsMask,
	modifiers int,
) WebDragOperationsMask
type wkeDragTargetDragLeave func(webView WebView)
type wkeDragTargetDrop func(
	webView WebView,
	clientPoint *Point,
	screenPoint *Point,
	modifiers int,
)
type wkeDragTargetEnd func(
	webView WebView,
	clientPoint *Point,
	screenPoint *Point,
	operation WebDragOperationsMask,
)
type wkeUtilSetUiCallback func(callback UiThreadPostTaskCallback)
type wkeUtilSerializeToMHTML func(webView WebView) string
type wkeUtilPrintToPdf func(
	webView WebView,
	frameId WebFrameHandle,
	settings *PrintSettings,
) *PdfDatas
type wkePrintToBitmap func(
	webView WebView,
	frameId WebFrameHandle,
	settings *ScreenshotSettings,
) *MemBuf
type wkeUtilRelasePrintPdfDatas func(datas *PdfDatas)
type wkeSetWindowTitle func(webWindow WebView, title string)
type wkeSetWindowTitleW func(webWindow WebView, title uint16)
type wkeNodeOnCreateProcess func(webView WebView, callback NodeOnCreateProcessCallback, param uintptr)
type wkeOnPluginFind func(webView WebView, mime string, callback OnPluginFindCallback, param uintptr)
type wkeAddNpapiPlugin func(webView WebView, initializeFunc, getEntryPointsFunc, shutdownFunc uintptr)
type wkePluginListBuilderAddPlugin func(builder uintptr, name, description, fileName string)
type wkePluginListBuilderAddMediaTypeToLastPlugin func(builder uintptr, name, description string)
type wkePluginListBuilderAddFileExtensionToLastMediaType func(builder uintptr, fileExtension string)
type wkeGetWebViewByNData func(ndata uintptr) WebView
type wkeRegisterEmbedderCustomElement func(
	webView WebView,
	frameId WebFrameHandle,
	name string,
	options uintptr, // 使用uintptr来表示任意类型的指针
	outResult uintptr, // 同样使用uintptr
) bool
type wkeSetMediaPlayerFactory func(
	webView WebView,
	factory MediaPlayerFactory,
	callback OnIsMediaPlayerSupportsMIMEType,
)
type wkeGetContentAsMarkup func(
	webView WebView,
	frame WebFrameHandle,
	size *uintptr,
) string
type wkeUtilDecodeURLEscape func(url string) string
type wkeUtilEncodeURLEscape func(url string) string
type wkeUtilBase64Encode func(str string) string
type wkeUtilBase64Decode func(str string) string
type wkeUtilCreateV8Snapshot func(str string) *MemBuf
type wkeRunMessageLoop func()
type wkeSaveMemoryCache func(webView WebView)

// type jsBindFunction func(name string, fn JsNativeFunction, argCount uint)
// type jsBindGetter func(name string, fn JsNativeFunction)
// type jsBindSetter func(name string, fn JsNativeFunction)
type wkeJsBindFunction func(name string, fn wkeJsNativeFunction, param uintptr, argCount uint)
type wkeJsBindGetter func(name string, fn wkeJsNativeFunction, param uintptr)
type wkeJsBindSetter func(name string, fn wkeJsNativeFunction, param uintptr)
type jsArgCount func(es JsExecState) int
type jsArgType func(es JsExecState, argIdx int) JsType
type jsArg func(es JsExecState, argIdx int) JsValue
type jsTypeOf func(v JsValue) JsType
type jsIsNumber func(v JsValue) bool
type jsIsString func(v JsValue) bool
type jsIsBoolean func(v JsValue) bool
type jsIsObject func(v JsValue) bool
type jsIsFunction func(v JsValue) bool
type jsIsUndefined func(v JsValue) bool
type jsIsNull func(v JsValue) bool
type jsIsArray func(v JsValue) bool
type jsIsTrue func(v JsValue) bool
type jsIsFalse func(v JsValue) bool
type jsToInt func(es JsExecState, v JsValue) int
type jsToFloat func(es JsExecState, v JsValue) float32
type jsToDouble func(es JsExecState, v JsValue) float64
type jsToDoubleString func(es JsExecState, v JsValue) string
type jsToBoolean func(es JsExecState, v JsValue) bool
type jsArrayBuffer func(es JsExecState, buffer *byte, size uintptr) JsValue
type jsGetArrayBuffer func(es JsExecState, value JsValue) *MemBuf
type jsToTempString func(es JsExecState, v JsValue) string
type jsToTempStringW func(es JsExecState, v JsValue) uint16
type v8ValuePtr uintptr // 假设v8ValuePtr是v8::Persistent<v8::Value>*的某种Go接口或类型
type jsToV8Value func(es JsExecState, v JsValue) v8ValuePtr
type jsInt func(n int) JsValue
type jsFloat func(f float32) JsValue
type jsDouble func(d float64) JsValue
type jsDoubleString func(str string) JsValue
type jsBoolean func(b bool) JsValue
type jsUndefined func() JsValue
type jsNull func() JsValue
type jsTrue func() JsValue
type jsFalse func() JsValue
type jsString func(es JsExecState, str string) JsValue
type jsStringW func(es JsExecState, str uint16) JsValue
type jsEmptyObject func(es JsExecState) JsValue
type jsEmptyArray func(es JsExecState) JsValue
type jsObject func(es JsExecState, obj JsData) JsValue
type jsFunction func(es JsExecState, obj JsData) JsValue
type jsGetData func(es JsExecState, object JsValue) JsData
type jsGet func(es JsExecState, object JsValue, prop string) JsValue
type jsSet func(es JsExecState, object JsValue, prop string, v JsValue)
type jsGetAt func(es JsExecState, object JsValue, index int) JsValue
type jsSetAt func(es JsExecState, object JsValue, index int, v JsValue)
type jsGetKeys func(es JsExecState, object JsValue) JsKeys
type jsIsJsValueValid func(es JsExecState, object JsValue) bool
type jsIsValidExecState func(es JsExecState) bool
type jsDeleteObjectProp func(es JsExecState, object JsValue, prop string)
type jsGetLength func(es JsExecState, object JsValue) int
type jsSetLength func(es JsExecState, object JsValue, length int)
type jsGlobalObject func(es JsExecState) JsValue
type jsGetWebView func(es JsExecState) WebView
type jsEval func(es JsExecState, str string) JsValue
type jsEvalW func(es JsExecState, str uint16) JsValue
type jsEvalExW func(es JsExecState, str uint16, isInClosure bool) JsValue
type jsCall func(es JsExecState, funcObj JsValue, thisObject JsValue, args []JsValue, argCount int) JsValue
type jsCallGlobal func(es JsExecState, funcObj JsValue, args []JsValue, argCount int) JsValue
type jsGetGlobal func(es JsExecState, prop string) JsValue
type jsSetGlobal func(es JsExecState, prop string, v JsValue)
type jsGC func()
type jsAddRef func(es JsExecState, val JsValue) bool
type jsReleaseRef func(es JsExecState, val JsValue) bool
type jsGetLastErrorIfException func(es JsExecState) *JsExceptionInfo
type jsThrowException func(es JsExecState, exception string) JsValue
type jsGetCallstack func(es JsExecState) string
type wkeInit func()
type wkeInitialize func()
type wkeInitializeEx func(settings *Settings)
