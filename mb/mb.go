package mb

import (
	"syscall"
	"unsafe"
)

// Rect 对应 C++ 中的 _Rect 结构体
type Rect struct {
	X, Y, W, H int32
}

// mbPoint 对应 C++ 中的 _mbPoint 结构体
type Point struct {
	X, Y int32
}

// mbSize 对应 C++ 中的 _mbSize 结构体
type Size struct {
	W, H int32
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
type OnBlinkThreadInitCallback func()

// mbSettings 对应 C++ 中的 _mbSettings 结构体
type Settings struct {
	Proxy                        Proxy
	Mask                         uint32
	BlinkThreadInitCallback      OnBlinkThreadInitCallback
	BlinkThreadInitCallbackParam uintptr // 使用unsafe.Pointer来存储任意类型的指针
	Version                      uintptr
	MainDllPath                  *uint16 // 假设使用uint16的指针来表示宽字符字符串（通常使用syscall.UTF16PtrFromString转换）
	MainDllHandle                Handle
	Config                       string // 假设config是C风格的字符串（char*），在Go中通常是*byte
}

// mbViewSettings 对应 C++ 中的 _mbViewSettings 结构体
type ViewSettings struct {
	Size    int32
	BgColor uint32
}

// WebFrameHandle 对应 C++ 中的 void*
type WebFrameHandle uintptr

// NetJob 对应 C++ 中的 void*
type NetJob uintptr

// WebView 对应 C++ 中的 intptr_t
type WebView uintptr

// type bool uint32

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
	X, Y, Width, Height int32
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
	StructSize               int32
	DPI                      int32
	Width, Height            int32
	MarginTop, MarginBottom  int32
	MarginLeft, MarginRight  int32
	IsPrintPageHeadAndFooter bool
	IsPrintBackgroud         bool // 注意：这里假设是 Background 的拼写错误，并修正为 Background
	IsLandscape              bool
	IsPrintToMultiPage       bool
}

type MbString uintptr

// MemBuf 对应C中的_MemBuf结构体
type MemBuf struct {
	Unuse  int32 // 这字段暂时没啥用
	Data   uintptr
	Length uintptr
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
	ItemListLength   int32
	ModifierKeyState int32 // State of Shift/Ctrl/Alt/Meta keys
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
	Data *byte // 使用指针来模拟C中的char*
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

type BOOL uintptr

const (
	FALSE BOOL = 0
	TRUE  BOOL = 1
)

// typedef void* WebSocketChannel;
type WebSocketChannel uintptr

// 回调函数签名
type OnWillConnectCallback func(webView WebView, channel WebSocketChannel, url string, needHook bool) MbString
type OnConnectedCallback func(webView WebView, channel WebSocketChannel) bool
type OnReceiveCallback func(webView WebView, channel WebSocketChannel, opCode int, buf []byte, isContinue bool) MbString
type OnSendCallback func(webView WebView, channel WebSocketChannel, opCode int, buf []byte, isContinue bool) MbString
type OnErrorCallback func(webView WebView, channel WebSocketChannel)

// 回调接口定义
type WebsocketHookCallbacks struct {
	OnWillConnect OnWillConnectCallback
	OnConnected   OnConnectedCallback
	OnReceive     OnReceiveCallback
	OnSend        OnSendCallback
	OnError       OnErrorCallback
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

type JsValue uintptr
type JsExecState uintptr
type HDC uintptr

// typedef void(MB_CALL_TYPE *mbOnGetPdfPageDataCallback)(mbWebView webView, void* param, void* data, size_t size);
type mbOnGetPdfPageDataCallback func(webView WebView, param uintptr, data uintptr, size uintptr) (void uintptr)
type OnGetPdfPageDataCallback func(webView WebView, data []byte)

// typedef void(MB_CALL_TYPE *mbRunJsCallback)(mbWebView webView, void* param, mbJsExecState es, mbJsValue v);
type mbRunJsCallback func(webView WebView, param uintptr, es JsExecState, v JsValue) (void uintptr)
type RunJsCallback func(webView WebView, es JsExecState, v JsValue)

// typedef void(MB_CALL_TYPE* mbJsQueryCallback)(mbWebView webView, void* param, mbJsExecState es, int64_t queryId, int customMsg, const utf8* request);
type mbJsQueryCallback func(webView WebView, param uintptr, es JsExecState, queryId int64, customMsg int, request uintptr) (void uintptr)
type JsQueryCallback func(webView WebView, es JsExecState, queryId int64, customMsg int, request string)

// typedef void(MB_CALL_TYPE *mbTitleChangedCallback)(mbWebView webView, void* param, const utf8* title);
type mbTitleChangedCallback func(webView WebView, param uintptr, title uintptr) (void uintptr)
type TitleChangedCallback func(webView WebView, title string)

// typedef void(MB_CALL_TYPE *mbMouseOverUrlChangedCallback)(mbWebView webView, void* param, const utf8* url);
type mbMouseOverUrlChangedCallback func(webView WebView, param uintptr, url uintptr) (void uintptr)
type MouseOverUrlChangedCallback func(webView WebView, url string)

// typedef void(MB_CALL_TYPE *mbURLChangedCallback)(mbWebView webView, void* param, const utf8* url, BOOL canGoBack, BOOL canGoForward);
type mbURLChangedCallback func(webView WebView, param uintptr, url uintptr, canGoBack, canGoForward bool) (void uintptr)
type URLChangedCallback func(webView WebView, url string, canGoBack bool, canGoForward bool)

// typedef void(MB_CALL_TYPE *mbURLChangedCallback2)(mbWebView webView, void* param, mbWebFrameHandle frameId, const utf8* url);
type mbURLChangedCallback2 func(webView WebView, param uintptr, frameId WebFrameHandle, url uintptr) (void uintptr)
type URLChangedCallback2 func(webView WebView, frameId WebFrameHandle, url string)

// typedef void(MB_CALL_TYPE *mbPaintUpdatedCallback)(mbWebView webView, void* param, const HDC hdc, int x, int y, int cx, int cy);
type mbPaintUpdatedCallback func(webView WebView, param uintptr, hdc HDC, x, y, cx, cy int32) (void uintptr)
type PaintUpdatedCallback func(webView WebView, hdc HDC, x, y, cx, cy int32)

// typedef void(MB_CALL_TYPE* mbAcceleratedPaintCallback)(mbWebView webView, void* param, int type, const mbRect* dirytRects, const size_t dirytRectsSize,void* sharedHandle);
type mbAcceleratedPaintCallback func(webView WebView, param uintptr, typ int32, dirytRects *Rect, dirytRectsSize, sharedHandle uintptr) (void uintptr)
type AcceleratedPaintCallback func(webView WebView, typ int32, dirytRects *Rect, dirytRectsSize uintptr, sharedHandle uintptr)

// typedef void(MB_CALL_TYPE *mbPaintBitUpdatedCallback)(mbWebView webView, void* param, const void* buffer, const mbRect* r, int width, int height);
type mbPaintBitUpdatedCallback func(webView WebView, param, buffer uintptr, r *Rect, width, height int32) (void uintptr)
type PaintBitUpdatedCallback func(webView WebView, buffer uintptr, r *Rect, width, height int32)

// typedef void(MB_CALL_TYPE *mbAlertBoxCallback)(mbWebView webView, void* param, const utf8* msg);
type mbAlertBoxCallback func(webView WebView, param uintptr, msg uintptr) (void uintptr)
type AlertBoxCallback func(webView WebView, msg string)

// typedef BOOL(MB_CALL_TYPE *mbConfirmBoxCallback)(mbWebView webView, void* param, const utf8* msg);
type mbConfirmBoxCallback func(webView WebView, param uintptr, msg uintptr) BOOL
type ConfirmBoxCallback func(webView WebView, msg string) bool

// typedef mbStringPtr(MB_CALL_TYPE *mbPromptBoxCallback)(mbWebView webView, void* param, const utf8* msg, const utf8* defaultResult, BOOL* result);
type mbPromptBoxCallback func(webView WebView, param uintptr, msg, defaultResult uintptr, result *BOOL) uintptr
type PromptBoxCallback func(webView WebView, msg, defaultResult string) MbString

// typedef BOOL(MB_CALL_TYPE *mbNavigationCallback)(mbWebView webView, void* param, mbNavigationType navigationType, const utf8* url);
type mbNavigationCallback func(webView WebView, param uintptr, navigationType NavigationType, url uintptr) BOOL
type NavigationCallback func(webView WebView, navigationType NavigationType, url string) bool

// typedef mbWebView(MB_CALL_TYPE *mbCreateViewCallback)(mbWebView webView, void* param, mbNavigationType navigationType, const utf8* url, const mbWindowFeatures* windowFeatures);
type mbCreateViewCallback func(webView WebView, param uintptr, navigationType NavigationType, url uintptr, windowFeatures *WindowFeatures) WebView
type CreateViewCallback func(webView WebView, navigationType NavigationType, url string, windowFeatures *WindowFeatures) WebView

// typedef void(MB_CALL_TYPE *mbDocumentReadyCallback)(mbWebView webView, void* param, mbWebFrameHandle frameId);
type mbDocumentReadyCallback func(webView WebView, param uintptr, frameId WebFrameHandle) (void uintptr)
type DocumentReadyCallback func(webView WebView, frameId WebFrameHandle)

// typedef void(MB_CALL_TYPE *mbLoadUrlFinishCallback)(mbWebView webView, void* param, const utf8* url, mbNetJob job, int len);
type mbLoadUrlFinishCallback func(webView WebView, param uintptr, url uintptr, job NetJob, length int) (void uintptr)
type LoadUrlFinishCallback func(webView WebView, url string, job NetJob, length int)

// typedef void(MB_CALL_TYPE *mbLoadUrlHeadersReceivedCallback)(mbWebView webView, void* param, const char* url, mbNetJob job);
type mbLoadUrlHeadersReceivedCallback func(webView WebView, param uintptr, url uintptr, job NetJob) (void uintptr)
type LoadUrlHeadersReceivedCallback func(webView WebView, url string, job NetJob)

// typedef BOOL(MB_CALL_TYPE *mbCloseCallback)(mbWebView webView, void* param, void* unuse);
type mbCloseCallback func(webView WebView, param uintptr, unuse uintptr) BOOL
type CloseCallback func(webView WebView, unuse uintptr) bool

// typedef BOOL(MB_CALL_TYPE *mbDestroyCallback)(mbWebView webView, void* param, void* unuse);
type mbDestroyCallback func(webView WebView, param uintptr, unuse uintptr) BOOL
type DestroyCallback func(webView WebView, unuse uintptr) bool

// typedef void(MB_CALL_TYPE *mbOnShowDevtoolsCallback)(mbWebView webView, void* param);
type mbOnShowDevtoolsCallback func(webView WebView, param uintptr) (void uintptr)
type OnShowDevtoolsCallback func(webView WebView)

// typedef void(MB_CALL_TYPE *mbDidCreateScriptContextCallback)(mbWebView webView, void* param, mbWebFrameHandle frameId, void* context, int extensionGroup, int worldId);
type mbDidCreateScriptContextCallback func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, extensionGroup, worldId int) (void uintptr)
type DidCreateScriptContextCallback func(webView WebView, frameId WebFrameHandle, context uintptr, extensionGroup, worldId int)

// typedef BOOL(MB_CALL_TYPE *mbGetPluginListCallback)(BOOL refresh, void* pluginListBuilder, void* param);
type mbGetPluginListCallback func(refresh bool, pluginListBuilder uintptr, param uintptr) BOOL
type GetPluginListCallback func(refresh bool, pluginListBuilder uintptr) bool

// typedef BOOL(MB_CALL_TYPE *mbNetResponseCallback)(mbWebView webView, void* param, const utf8* url, mbNetJob job);
type mbNetResponseCallback func(webView WebView, param uintptr, url uintptr, job NetJob) BOOL
type NetResponseCallback func(webView WebView, url string, job NetJob) bool

// typedef void(MB_CALL_TYPE* mbThreadCallback)(void* param1, void* param2);
type mbThreadCallback func(param1, param2 uintptr) (void uintptr)
type ThreadCallback func()

// typedef void(MB_CALL_TYPE* mbNodeOnCreateProcessCallback)(mbWebView webView, void* param, const WCHAR* applicationPath, const WCHAR* arguments, STARTUPINFOW* startup);
type mbNodeOnCreateProcessCallback func(webView WebView, param, applicationPath, arguments, startup uintptr) (void uintptr)
type NodeOnCreateProcessCallback func(webView WebView, applicationPath, arguments string, startup *STARTUPINFOW)

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

// mbLoadingResult 是一个表示加载结果的类型
type LoadingResult int

// 定义mbLoadingResult的常量
const (
	LOADING_SUCCEEDED LoadingResult = iota // iota是Go语言的枚举计数器
	LOADING_FAILED
	LOADING_CANCELED
)

// mbLoadingFinishCallback
type LoadingFinishCallback func(webView WebView, frameId WebFrameHandle, url string, result LoadingResult, failedReason string)

// mbDownloadCallback
type DownloadCallback func(webView WebView, frameId WebFrameHandle, url string, downloadJob uintptr) bool

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

// typedef void(MB_CALL_TYPE *mbConsoleCallback)(mbWebView webView, void* param, mbConsoleLevel level, const utf8* message, const utf8* sourceName, unsigned sourceLine, const utf8* stackTrace);
type mbConsoleCallback func(webView WebView, param uintptr, level ConsoleLevel, message uintptr, sourceName uintptr, sourceLine uint, stackTrace uintptr) (void uintptr)
type ConsoleCallback func(webView WebView, level ConsoleLevel, message string, sourceName string, sourceLine uint, stackTrace string)

// typedef void(MB_CALL_TYPE *mbOnCallUiThread)(mbWebView webView, void* paramOnInThread);
type mbOnCallUiThread func(webView WebView, param uintptr) (void uintptr)
type OnCallUiThreadCallback func(webView WebView)

type mbCallUiThreadPtr uintptr

func (addr mbCallUiThreadPtr) Call(webView WebView, param uintptr) (void uintptr) {
	syscall.SyscallN(uintptr(addr), uintptr(webView), param)
	return
}

// typedef void(MB_CALL_TYPE *mbCallUiThread)(mbWebView webView, mbOnCallUiThread func, void* param);
type mbCallUiThread func(webView WebView, fn mbCallUiThreadPtr, param uintptr) (void uintptr)
type CallUiThreadCallback func(webView WebView, funcOnUiThread OnCallUiThreadCallback)

// typedef BOOL(MB_CALL_TYPE *mbLoadUrlBeginCallback)(mbWebView webView, void* param, const char* url, void* job);
type mbLoadUrlBeginCallback func(webView WebView, param, url uintptr, job NetJob) BOOL
type LoadUrlBeginCallback func(webView WebView, url string, job NetJob) bool

// typedef void(MB_CALL_TYPE *mbLoadUrlEndCallback)(mbWebView webView, void* param, const char* url, void* job, void* buf, int len);
type mbLoadUrlEndCallback func(webView WebView, param, url uintptr, job NetJob, buf uintptr, len int) (void uintptr)
type LoadUrlEndCallback func(webView WebView, url string, job NetJob, buf []byte)

// typedef void(MB_CALL_TYPE *mbLoadUrlFailCallback)(mbWebView webView, void* param, const char* url, void* job);
type mbLoadUrlFailCallback func(webView WebView, param, url uintptr, job NetJob) (void uintptr)
type LoadUrlFailCallback func(webView WebView, url string, job NetJob)

// typedef void(MB_CALL_TYPE *mbWillReleaseScriptContextCallback)(mbWebView webView, void* param, mbWebFrameHandle frameId, void* context, int worldId);
type mbWillReleaseScriptContextCallback func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, worldId int) (void uintptr)
type WillReleaseScriptContextCallback func(webView WebView, frameId WebFrameHandle, context uintptr, worldId int)

// typedef void(MB_CALL_TYPE *mbNetGetFaviconCallback)(mbWebView webView, void* param, const utf8* url, mbMemBuf* buf);
type mbNetGetFaviconCallback func(webView WebView, param uintptr, url uintptr, buf *MemBuf) (void uintptr)
type NetGetFaviconCallback func(webView WebView, url string, buf *MemBuf)

// AsynRequestState 枚举类型
type AsynRequestState int

// 枚举常量
const (
	AsynRequestStateOk   AsynRequestState = 0
	AsynRequestStateFail AsynRequestState = 1
)

// typedef void(MB_CALL_TYPE* mbCanGoBackForwardCallback)(mbWebView webView, void* param, MbAsynRequestState state, BOOL b);
type mbCanGoBackForwardCallback func(webView WebView, param uintptr, state AsynRequestState, b BOOL) (void uintptr)
type CanGoBackForwardCallback func(webView WebView, state AsynRequestState, b bool)

// typedef void(MB_CALL_TYPE* mbGetCookieCallback)(mbWebView webView, void* param, MbAsynRequestState state, const utf8* cookie);
type mbGetCookieCallback func(webView WebView, param uintptr, state AsynRequestState, cookie uintptr) (void uintptr)
type GetCookieCallback func(webView WebView, state AsynRequestState, cookie string)

// 类型别名定义
type V8ContextPtr uintptr // v8引擎的上下文指针
type V8Isolate uintptr    // v8引擎的隔离区指针

// typedef void(MB_CALL_TYPE* mbGetSourceCallback)(mbWebView webView, void* param, const utf8* mhtml);
type mbGetSourceCallback func(webView WebView, param, mhtml uintptr) (void uintptr)
type GetSourceCallback func(webView WebView, mhtml string)

// typedef void(MB_CALL_TYPE* mbGetContentAsMarkupCallback)(mbWebView webView, void* param, const utf8* content, size_t size);
type mbGetContentAsMarkupCallback func(webView WebView, param, content, size uintptr) (void uintptr)
type GetContentAsMarkupCallback func(webView WebView, content []byte)

// 结构体指针类型定义
type WebUrlRequest struct{}
type WebUrlResponse struct{}
type WebUrlRequestPtr *WebUrlRequest
type WebUrlResponsePtr *WebUrlResponse

// typedef void(MB_CALL_TYPE* mbOnUrlRequestWillRedirectCallback)(mbWebView webView, void* param, mbWebUrlRequestPtr oldRequest, mbWebUrlRequestPtr request, mbWebUrlResponsePtr redirectResponse);
type mbOnUrlRequestWillRedirectCallback func(webView WebView, param uintptr, oldRequest WebUrlRequestPtr, request WebUrlRequestPtr, redirectResponse WebUrlResponsePtr) (void uintptr)
type OnUrlRequestWillRedirectCallback func(webView WebView, oldRequest WebUrlRequestPtr, request WebUrlRequestPtr, redirectResponse WebUrlResponsePtr)

// typedef void(MB_CALL_TYPE* mbOnUrlRequestDidReceiveResponseCallback)(mbWebView webView, void* param, mbWebUrlRequestPtr request, mbWebUrlResponsePtr response);
type mbOnUrlRequestDidReceiveResponseCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, response WebUrlResponsePtr) (void uintptr)
type OnUrlRequestDidReceiveResponseCallback func(webView WebView, request WebUrlRequestPtr, response WebUrlResponsePtr)

// typedef void(MB_CALL_TYPE* mbOnUrlRequestDidReceiveDataCallback)(mbWebView webView, void* param, mbWebUrlRequestPtr request, const char* data, int dataLength);
type mbOnUrlRequestDidReceiveDataCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, data *byte, dataLength int) (void uintptr)
type OnUrlRequestDidReceiveDataCallback func(webView WebView, request WebUrlRequestPtr, data []byte)

// typedef void(MB_CALL_TYPE* mbOnUrlRequestDidFailCallback)(mbWebView webView, void* param, mbWebUrlRequestPtr request, const utf8* error);
type mbOnUrlRequestDidFailCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, error uintptr) (void uintptr)
type OnUrlRequestDidFailCallback func(webView WebView, request WebUrlRequestPtr, err string)

// typedef void(MB_CALL_TYPE* mbOnUrlRequestDidFinishLoadingCallback)(mbWebView webView, void* param, mbWebUrlRequestPtr request, double finishTime);
type mbOnUrlRequestDidFinishLoadingCallback func(webView WebView, param uintptr, request WebUrlRequestPtr, finishTime uintptr) (void uintptr)
type OnUrlRequestDidFinishLoadingCallback func(webView WebView, request WebUrlRequestPtr, finishTime float64)

// UrlRequestCallbacks 结构体
type UrlRequestCallbacks struct {
	WillRedirectCallback       OnUrlRequestWillRedirectCallback
	DidReceiveResponseCallback OnUrlRequestDidReceiveResponseCallback
	DidReceiveDataCallback     OnUrlRequestDidReceiveDataCallback
	DidFailCallback            OnUrlRequestDidFailCallback
	DidFinishLoadingCallback   OnUrlRequestDidFinishLoadingCallback
}

// DownloadOpt 枚举
type DownloadOpt uintptr

const (
	DownloadOptCancel    DownloadOpt = iota // iota是Go中的计数器，用于枚举值
	DownloadOptCacheData                    // 后续枚举值会自动递增
)

// typedef void(MB_CALL_TYPE*mbNetJobDataRecvCallback)(void* ptr, mbNetJob job, const char* data, int length);
type mbNetJobDataRecvCallback func(ptr uintptr, job NetJob, data uintptr, length int) (void uintptr)
type NetJobDataRecvCallback func(ptr uintptr, job NetJob, data []byte)

// func (cb NetJobDataRecvCallback) ToC() mbNetJobDataRecvCallback {
// 	return func(ptr uintptr, job NetJob, data uintptr, length int) (void uintptr) {
// 		cb(ptr, job, unsafe.Slice((*byte)(unsafe.Pointer(data)), length))
// 		return
// 	}
// }
// func (cb mbNetJobDataRecvCallback) ToGo() NetJobDataRecvCallback {
// 	return func(ptr uintptr, job NetJob, data []byte) {
// 		cb(ptr, job, uintptr(unsafe.Pointer(&data[0])), len(data))
// 	}
// }

type NetJobDataRecvCallbackPtr uintptr

func (cb NetJobDataRecvCallbackPtr) Call(ptr uintptr, job NetJob, data []byte) (void uintptr) {
	_, _, _ = syscall.SyscallN(uintptr(cb), ptr, uintptr(job), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return
}

// typedef void(MB_CALL_TYPE*mbNetJobDataFinishCallback)(void* ptr, mbNetJob job, mbLoadingResult result);
type mbNetJobDataFinishCallback func(ptr uintptr, job NetJob, result LoadingResult) (void uintptr)
type NetJobDataFinishCallback func(ptr uintptr, job NetJob, result LoadingResult)

type NetJobDataFinishCallbackPtr uintptr

func (cb NetJobDataFinishCallbackPtr) Call(ptr uintptr, job NetJob, result LoadingResult) (void uintptr) {
	_, _, _ = syscall.SyscallN(uintptr(cb), ptr, uintptr(job), uintptr(result))
	return
}

// typedef void(MB_CALL_TYPE*mbPopupDialogSaveNameCallback)(void* ptr, const WCHAR* filePath);
type mbPopupDialogSaveNameCallback func(ptr uintptr, filePath *uint16) (void uintptr)
type PopupDialogSaveNameCallback func(ptr uintptr, filePath string)

type PopupDialogSaveNameCallbackPtr uintptr

func (cb PopupDialogSaveNameCallbackPtr) Call(ptr uintptr, filePath string) (void uintptr) {
	_, _, _ = syscall.SyscallN(uintptr(cb), ptr, StrToWPtr(filePath))
	return
}

// 结构体定义
type NetJobDataBind struct {
	Param          uintptr
	RecvCallback   NetJobDataRecvCallbackPtr
	FinishCallback NetJobDataFinishCallbackPtr
}

type DownloadBind struct {
	Param            uintptr
	RecvCallback     NetJobDataRecvCallbackPtr
	FinishCallback   NetJobDataFinishCallbackPtr
	SaveNameCallback PopupDialogSaveNameCallbackPtr
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
	Magic                   int32  // 'mbdo'
	Title                   string // 标题
	DefaultPath             string // 默认路径
	ButtonLabel             string // 按钮标签
	Filters                 *FileFilter
	FiltersCount            int32
	Prop                    DialogProperties
	Message                 string // 消息
	SecurityScopedBookmarks bool
}

// DownloadOptions 结构体定义了下载选项
type DownloadOptions struct {
	Magic             int32
	SaveAsPathAndName bool
}

// typedef mbDownloadOpt(MB_CALL_TYPE*mbDownloadInBlinkThreadCallback)(mbWebView webView, void* param,size_t expectedContentLength,const char* url, const char* mime, const char* disposition, mbNetJob job, mbNetJobDataBind* dataBind);
type mbDownloadInBlinkThreadCallback func(webView WebView, param uintptr, expectedContentLength uintptr, url, mime, disposition uintptr, job NetJob, dataBind *NetJobDataBind) uintptr
type DownloadInBlinkThreadCallback func(webView WebView, expectedContentLength uintptr, url string, mime string, disposition string, job NetJob, dataBind *NetJobDataBind) DownloadOpt

// mbPdfDatas 结构体
type PdfDatas struct {
	Count int32
	Sizes *uintptr  // 指向 uintptr 类型的指针数组
	Datas **uintptr // 指向 void* 类型的指针的指针数组，即 const void**
}

// PrintPdfDataCallback 函数指针类型
type PrintPdfDataCallback func(webview WebView, datas *PdfDatas)

// ScreenshotSettings 结构体
type ScreenshotSettings struct {
	StructSize int32
	Width      int32
	Height     int32
}

// PrintBitmapCallback 函数指针类型
type PrintBitmapCallback func(webview WebView, data []byte)

// OnScreenshot 函数指针类型
type OnScreenshotCallback func(webview WebView, data []byte)

// 枚举 mbHttBodyElementType
type HttBodyElementType int

const (
	HttBodyElementTypeData HttBodyElementType = iota
	HttBodyElementTypeFile
)

// 结构体 PostBodyElement
type PostBodyElement struct {
	Size       int32
	Type       HttBodyElementType
	Data       *MemBuf // 假设 MemBuf 已经被定义
	FilePath   MbString
	FileStart  int64
	FileLength int64 // -1 表示到文件末尾
}

// PostBodyElements 结构体
type PostBodyElements struct {
	Size        int32
	Element     **PostBodyElement
	ElementSize uintptr
	IsDirty     bool
}

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
// typedef BOOL(MB_CALL_TYPE *mbWindowClosingCallback)(mbWebView webview, void* param);
type WindowClosingCallback func(webview WebView) bool

// typedef void(MB_CALL_TYPE *mbWindowDestroyCallback)(mbWebView webview, void* param);
type WindowDestroyCallback func(webview WebView)

// typedef void(MB_CALL_TYPE *mbDraggableRegionsChangedCallback)(mbWebView webview, void* param, const mbDraggableRegion* rects, int rectCount);
type DraggableRegionsChangedCallback func(webview WebView, rects *DraggableRegion, rectCount int)

// mbPrintintStep 枚举（注意：这里可能有个拼写错误，应该是 mbPrintStep）
type PrintintStep int

const (
	PrintStepStart PrintintStep = iota
	PrintStepPreview
	PrintStepPrinting
)

// mbPrintintSettings 结构体
type PrintintSettings struct {
	Dpi    int32
	Width  int32
	Height int32
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
type PrintingCallback func(webview WebView, step PrintintStep, hDC Handle, settings *PrintintSettings, pageCount int32) bool

// mbImageBufferToDataURLCallback 函数指针类型
type ImageBufferToDataURLCallback func(webView WebView, data []byte) MbString

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type HWND uintptr

type WPARAM uintptr
type LPARAM uintptr
type LRESULT uintptr
