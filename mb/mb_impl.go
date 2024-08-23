package mb

import (
	"syscall"
	"unsafe"

	"github.com/epkgs/miniblink/internal/lib"
)

var mbLib *lib.DLL

func LoadLibrary(name string) error {
	var err error
	mbLib, err = lib.LoadLibrary(name)
	if err != nil {
		return err
	}
	return nil
}

func Release() error {
	return lib.Release()
}

// void mbUninit();
func Init(settings *Settings) {
	mbLib.MustFindProc("mbInit").Call(uintptr(unsafe.Pointer(settings)))
}

func Uninit() {
	mbLib.MustFindProc("mbUninit").Call()
}

// 方便c#等其他语言创建setting结构体
// mbSettings* mbCreateInitSettings();
func CreateInitSettings() *Settings {
	r1, _, _ := mbLib.MustFindProc("mbCreateInitSettings").Call()
	return (*Settings)(unsafe.Pointer(r1))
}

// void mbSetInitSettings(mbSettings* settings, const char* name, const char* value);
func SetInitSettings(settings *Settings, name, value string) {
	mbLib.MustFindProc("mbSetInitSettings").Call(uintptr(unsafe.Pointer(settings)), StrToPtr(name), StrToPtr(value))
}

// mbWebView mbCreateWebView();
func CreateWebView() WebView {
	r1, _, _ := mbLib.MustFindProc("mbCreateWebView").Call()
	return WebView(r1)
}

// 用于GTK绑定窗口
// mbWebView mbCreateWebViewBindGtkWindow(void* rootWindow, void* drawingArea, const char* type, DWORD style, DWORD styleEx, int width, int height)
func CreateWebViewBindGtkWindow(rootWindow, drawingArea uintptr, typ string, style, styleEx uint32, width, height int32) WebView {
	r1, _, _ := mbLib.MustFindProc("mbCreateWebViewBindGtkWindow").Call(rootWindow, drawingArea, StrToPtr(typ), uintptr(style), uintptr(styleEx), uintptr(width), uintptr(height))
	return WebView(r1)
}

// void mbDestroyWebView(mbWebView webview);
func DestroyWebView(webview WebView) {
	mbLib.MustFindProc("mbDestroyWebView").Call(uintptr(webview))
}

// mbWebView mbCreateWebWindow(mbWindowType type, HWND parent, int x, int y, int width, int height);
func CreateWebWindow(typ WindowType, parent HWND, x, y, width, height int32) WebView {
	r1, _, _ := mbLib.MustFindProc("mbCreateWebWindow").Call(uintptr(typ), uintptr(parent), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return WebView(r1)
}

// mbWebView mbCreateWebCustomWindow(HWND parent, DWORD style, DWORD styleEx, int x, int y, int width, int height);
func CreateWebCustomWindow(parent HWND, style, styleEx uint32, x, y, width, height int32) WebView {
	r1, _, _ := mbLib.MustFindProc("mbCreateWebCustomWindow").Call(uintptr(parent), uintptr(style), uintptr(styleEx), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return WebView(r1)
}

// void mbMoveWindow(mbWebView webview, int x, int y, int w, int h);
func MoveWindow(webview WebView, x, y, width, height int32) {
	mbLib.MustFindProc("mbMoveWindow").Call(uintptr(webview), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

// void mbMoveToCenter(mbWebView webview);
func MoveToCenter(webview WebView) {
	mbLib.MustFindProc("mbMoveToCenter").Call(uintptr(webview))
}

// 离屏模式下控制是否自动上屏
// void mbSetAutoDrawToHwnd(mbWebView webview, BOOL b);
func SetAutoDrawToHwnd(webview WebView, b bool) {
	mbLib.MustFindProc("mbSetAutoDrawToHwnd").Call(uintptr(webview), BoolToPtr(b))
}

// void mbGetCaretRect(mbWebView webviewHandle, mbRect* r);
func GetCaretRect(webview WebView, r *Rect) {
	mbLib.MustFindProc("mbGetCaretRect").Call(uintptr(webview), uintptr(unsafe.Pointer(r)))
}

// void mbSetAudioMuted(mbWebView webview, BOOL b);
func SetAudioMuted(webview WebView, b bool) {
	mbLib.MustFindProc("mbSetAudioMuted").Call(uintptr(webview), BoolToPtr(b))
}

// BOOL mbIsAudioMuted(mbWebView webview);
func IsAudioMuted(webview WebView) bool {
	r1, _, _ := mbLib.MustFindProc("mbIsAudioMuted").Call(uintptr(webview))
	return 0 != r1
}

// mbStringPtr mbCreateString(const utf8* str, size_t length);
func CreateString(str string, length uintptr) MbString {
	r1, _, _ := mbLib.MustFindProc("mbCreateString").Call(StrToPtr(str), length)
	return MbString(r1)
}

// mbStringPtr mbCreateStringWithoutNullTermination(const utf8* str, size_t length);
func CreateStringWithoutNullTermination(str string, length uintptr) MbString {
	r1, _, _ := mbLib.MustFindProc("mbCreateStringWithoutNullTermination").Call(StrToPtr(str), length)
	return MbString(r1)
}

// void mbDeleteString(mbStringPtr str);
func DeleteString(mbStr MbString) {
	mbLib.MustFindProc("mbDeleteString").Call(uintptr(mbStr))
}

// size_t mbGetStringLen(mbStringPtr str);
func GetStringLen(mbStr MbString) uintptr {
	r1, _, _ := mbLib.MustFindProc("mbGetStringLen").Call(uintptr(mbStr))
	return r1
}

// const utf8* mbGetString(mbStringPtr str);
func GetString(mbStr MbString) string {
	r1, _, _ := mbLib.MustFindProc("mbGetString").Call(uintptr(mbStr))
	return StrFromPtr(r1)
}

// void mbSetProxy(mbWebView webView, const mbProxy* proxy);
func SetProxy(webView WebView, proxy *Proxy) {
	mbLib.MustFindProc("mbSetProxy").Call(uintptr(webView), uintptr(unsafe.Pointer(proxy)))
}

// void mbSetDebugConfig(mbWebView webView, const char* debugString, const char* param);
func SetDebugConfig(webView WebView, debugString, param string) {
	mbLib.MustFindProc("mbSetDebugConfig").Call(uintptr(webView), StrToPtr(debugString), StrToPtr(param))
}

// 调用此函数后,网络层收到数据会存储在一buf内,接收数据完成后响应OnLoadUrlEnd事件.#此调用严重影响性能,慎用
// 此函数和mbNetSetData的区别是，mbNetHookRequest会在接受到真正网络数据后再调用回调，并允许回调修改网络数据。
// 而mbNetSetData是在网络数据还没发送的时候修改
// void mbNetSetData(mbNetJob jobPtr, void* buf, int len);
func NetSetData(jobPtr NetJob, buf []byte) {
	mbLib.MustFindProc("mbNetSetData").Call(uintptr(jobPtr), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

// void mbNetHookRequest(mbNetJob jobPtr);
func NetHookRequest(jobPtr NetJob) {
	mbLib.MustFindProc("mbNetHookRequest").Call(uintptr(jobPtr))
}

// void mbNetChangeRequestUrl(mbNetJob jobPtr, const char* url);
func NetChangeRequestUrl(jobPtr NetJob, url string) {
	mbLib.MustFindProc("mbNetChangeRequestUrl").Call(uintptr(jobPtr), StrToPtr(url))
}

// void mbNetContinueJob(mbNetJob jobPtr);
func NetContinueJob(jobPtr NetJob) {
	mbLib.MustFindProc("mbNetContinueJob").Call(uintptr(jobPtr))
}

// const mbSlist* mbNetGetRawHttpHeadInBlinkThread(mbNetJob jobPtr);
func NetGetRawHttpHeadInBlinkThread(jobPtr NetJob) *Slist {
	r1, _, _ := mbLib.MustFindProc("mbNetGetRawHttpHeadInBlinkThread").Call(uintptr(jobPtr))
	return (*Slist)(unsafe.Pointer(r1))
}

// const mbSlist* mbNetGetRawResponseHeadInBlinkThread(mbNetJob jobPtr);
func NetGetRawResponseHeadInBlinkThread(jobPtr NetJob) *Slist {
	r1, _, _ := mbLib.MustFindProc("mbNetGetRawResponseHeadInBlinkThread").Call(uintptr(jobPtr))
	return (*Slist)(unsafe.Pointer(r1))
}

// void mbNetHoldJobToAsynCommit(mbNetJob jobPtr);
func NetHoldJobToAsynCommit(jobPtr NetJob) {
	_, _, _ = mbLib.MustFindProc("mbNetHoldJobToAsynCommit").Call(uintptr(jobPtr))
}

// void mbNetCancelRequest(mbNetJob jobPtr);
func NetCancelRequest(jobPtr NetJob) {
	mbLib.MustFindProc("mbNetCancelRequest").Call(uintptr(jobPtr))
}

// 注意此接口的回调是在另外个线程
// void mbNetOnResponse(mbWebView webviewHandle, mbNetResponseCallback callback, void* param);
func NetOnResponse(webviewHandle WebView, callback NetResponseCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbNetResponseCallback = func(webView WebView, param, url uintptr, job NetJob) BOOL {
			if callback(webView, StrFromPtr(url), job) {
				return TRUE
			} else {
				return FALSE
			}
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbNetOnResponse").Call(uintptr(webviewHandle), cbPtr, 0)
}

// void mbNetSetWebsocketCallback(mbWebView webview, const mbWebsocketHookCallbacks* callbacks, void* param);
func NetSetWebsocketCallback(webview WebView, callbacks *WebsocketHookCallbacks) {

	type mbWebsocketHookCallbacks struct {
		onWillConnect uintptr
		onConnected   uintptr
		onReceive     uintptr
		onSend        uintptr
		onError       uintptr
	}

	var onWillConnect = func(webView WebView, param uintptr, channel WebSocketChannel, url uintptr, needHook uintptr) uintptr {
		rst := callbacks.OnWillConnect(webView, channel, StrFromPtr(url), BoolFromPtr(needHook))
		return uintptr(rst)
	}

	var onConnected = func(webView WebView, param uintptr, channel WebSocketChannel) BOOL {
		if callbacks.OnConnected(webView, channel) {
			return TRUE
		} else {
			return FALSE
		}
	}

	var onReceive = func(webView WebView, param uintptr, channel WebSocketChannel, opCode int, buf *byte, len uintptr, isContinue uintptr) uintptr {
		data := unsafe.Slice(buf, len)
		rst := callbacks.OnReceive(webView, channel, opCode, data, BoolFromPtr(isContinue))
		return uintptr(rst)
	}

	var onSend = func(webView WebView, param uintptr, channel WebSocketChannel, opCode int, buf *byte, len uintptr, isContinue uintptr) uintptr {
		data := unsafe.Slice(buf, len)
		rst := callbacks.OnSend(webView, channel, opCode, data, BoolFromPtr(isContinue))
		return uintptr(rst)
	}

	var onError = func(webView WebView, param uintptr, channel WebSocketChannel) (void uintptr) {
		callbacks.OnError(webView, channel)
		return
	}

	cbs := mbWebsocketHookCallbacks{
		onWillConnect: CallbackToPtr(onWillConnect),
		onConnected:   CallbackToPtr(onConnected),
		onReceive:     CallbackToPtr(onReceive),
		onSend:        CallbackToPtr(onSend),
		onError:       CallbackToPtr(onError),
	}

	mbLib.MustFindProc("mbNetSetWebsocketCallback").Call(uintptr(webview), uintptr(unsafe.Pointer(&cbs)), 0)
}

// void mbNetSendWsText(mbWebSocketChannel channel, const char* buf, size_t len);
func NetSendWsText(channel WebSocketChannel, buf []byte) {
	mbLib.MustFindProc("mbNetSendWsText").Call(uintptr(channel), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

// void mbNetSendWsBlob(mbWebSocketChannel channel, const char* buf, size_t len);
func NetSendWsBlob(channel WebSocketChannel, buf []byte) {
	mbLib.MustFindProc("mbNetSendWsBlob").Call(uintptr(channel), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

// void mbNetEnableResPacket(mbWebView webviewHandle, const WCHAR* pathName);
func NetEnableResPacket(webviewHandle WebView, pathName string) {
	mbLib.MustFindProc("mbNetEnableResPacket").Call(uintptr(webviewHandle), StrToWPtr(pathName))
}

// mbPostBodyElements* mbNetGetPostBody(mbNetJob jobPtr);
func NetGetPostBody(jobPtr NetJob) *PostBodyElements {
	r1, _, _ := mbLib.MustFindProc("mbNetGetPostBody").Call(uintptr(jobPtr))
	return (*PostBodyElements)(unsafe.Pointer(r1))
}

// mbPostBodyElements* mbNetCreatePostBodyElements(mbWebView webView, size_t length);
func NetCreatePostBodyElements(webView WebView, length uintptr) *PostBodyElements {
	r1, _, _ := mbLib.MustFindProc("mbNetCreatePostBodyElements").Call(uintptr(webView), length)
	return (*PostBodyElements)(unsafe.Pointer(r1))
}

// void mbNetFreePostBodyElements(mbPostBodyElements* elements);
func NetFreePostBodyElements(elements *PostBodyElements) {
	mbLib.MustFindProc("mbNetFreePostBodyElements").Call(uintptr(unsafe.Pointer(elements)))
}

// mbPostBodyElement* mbNetCreatePostBodyElement(mbWebView webView);
func NetCreatePostBodyElement(webView WebView) *PostBodyElement {
	r1, _, _ := mbLib.MustFindProc("mbNetCreatePostBodyElement").Call(uintptr(webView))
	return (*PostBodyElement)(unsafe.Pointer(r1))
}

// void mbNetFreePostBodyElement(mbPostBodyElement* element);
func NetFreePostBodyElement(element *PostBodyElement) {
	mbLib.MustFindProc("mbNetFreePostBodyElement").Call(uintptr(unsafe.Pointer(element)))
}

// mbWebUrlRequestPtr mbNetCreateWebUrlRequest(const utf8* url, const utf8* method, const utf8* mime);
func NetCreateWebUrlRequest(url, method, mime string) *WebUrlRequest {
	r1, _, _ := mbLib.MustFindProc("mbNetCreateWebUrlRequest").Call(StrToPtr(url), StrToPtr(method), StrToPtr(mime))
	return (*WebUrlRequest)(unsafe.Pointer(r1))
}

// void mbNetAddHTTPHeaderFieldToUrlRequest(mbWebUrlRequestPtr request, const utf8* name, const utf8* value);
func NetAddHTTPHeaderFieldToUrlRequest(request *WebUrlRequest, name, value string) {
	mbLib.MustFindProc("mbNetAddHTTPHeaderFieldToUrlRequest").Call(uintptr(unsafe.Pointer(request)), StrToPtr(name), StrToPtr(value))
}

// int mbNetStartUrlRequest(mbWebView webView, mbWebUrlRequestPtr request, void* param, const mbUrlRequestCallbacks* callbacks);
func NetStartUrlRequest(webView WebView, request *WebUrlRequest, callbacks *UrlRequestCallbacks) int {
	type Callbacks struct {
		willRedirectCallback       uintptr
		didReceiveResponseCallback uintptr
		didReceiveDataCallback     uintptr
		didFailCallback            uintptr
		didFinishLoadingCallback   uintptr
	}

	var willRedirectCallback mbOnUrlRequestWillRedirectCallback = func(webView WebView, param uintptr, oldRequest WebUrlRequestPtr, request WebUrlRequestPtr, redirectResponse WebUrlResponsePtr) (void uintptr) {
		callbacks.WillRedirectCallback(webView, oldRequest, request, redirectResponse)
		return
	}

	var didReceiveResponseCallback mbOnUrlRequestDidReceiveResponseCallback = func(webView WebView, param uintptr, request WebUrlRequestPtr, response WebUrlResponsePtr) (void uintptr) {
		callbacks.DidReceiveResponseCallback(webView, request, response)
		return
	}

	var didReceiveDataCallback mbOnUrlRequestDidReceiveDataCallback = func(webView WebView, param uintptr, request WebUrlRequestPtr, data *byte, dataLength int) (void uintptr) {
		byts := unsafe.Slice(data, dataLength)
		callbacks.DidReceiveDataCallback(webView, request, byts)
		return
	}

	var didFailCallback mbOnUrlRequestDidFailCallback = func(webView WebView, param uintptr, request WebUrlRequestPtr, err uintptr) (void uintptr) {
		callbacks.DidFailCallback(webView, request, StrFromPtr(err))
		return
	}

	var didFinishLoadingCallback mbOnUrlRequestDidFinishLoadingCallback = func(webView WebView, param uintptr, request WebUrlRequestPtr, finishTime uintptr) (void uintptr) {
		callbacks.DidFinishLoadingCallback(webView, request, F64FromPtr(finishTime))
		return
	}

	cbs := Callbacks{
		willRedirectCallback:       CallbackToPtr(willRedirectCallback),
		didReceiveResponseCallback: CallbackToPtr(didReceiveResponseCallback),
		didReceiveDataCallback:     CallbackToPtr(didReceiveDataCallback),
		didFailCallback:            CallbackToPtr(didFailCallback),
		didFinishLoadingCallback:   CallbackToPtr(didFinishLoadingCallback),
	}

	r1, _, _ := mbLib.MustFindProc("mbNetStartUrlRequest").Call(uintptr(webView), uintptr(unsafe.Pointer(request)), 0, uintptr(unsafe.Pointer(&cbs)))
	return int(r1)
}

// int mbNetGetHttpStatusCode(mbWebUrlResponsePtr response);
func NetGetHttpStatusCode(response *WebUrlResponse) int {
	r1, _, _ := mbLib.MustFindProc("mbNetGetHttpStatusCode").Call(uintptr(unsafe.Pointer(response)))
	return int(r1)
}

// mbRequestType mbNetGetRequestMethod(mbNetJob jobPtr);
func NetGetRequestMethod(jobPtr NetJob) RequestType {
	r1, _, _ := mbLib.MustFindProc("mbNetGetRequestMethod").Call(uintptr(jobPtr))
	return RequestType(r1)
}

// __int64 mbNetGetExpectedContentLength(mbWebUrlResponsePtr response);
func NetGetExpectedContentLength(response *WebUrlResponse) int64 {
	r1, _, _ := mbLib.MustFindProc("mbNetGetExpectedContentLength").Call(uintptr(unsafe.Pointer(response)))
	return int64(r1)
}

// const utf8* mbNetGetResponseUrl(mbWebUrlResponsePtr response);
func NetGetResponseUrl(response *WebUrlResponse) string {
	r1, _, _ := mbLib.MustFindProc("mbNetGetResponseUrl").Call(uintptr(unsafe.Pointer(response)))
	return StrFromPtr(r1)
}

// void mbNetCancelWebUrlRequest(int requestId);
func NetCancelWebUrlRequest(requestId int) {
	mbLib.MustFindProc("mbNetCancelWebUrlRequest").Call(uintptr(requestId))
}

// void mbSetViewProxy(mbWebView webView, const mbProxy* proxy);
func SetViewProxy(webView WebView, proxy *Proxy) {
	mbLib.MustFindProc("mbSetViewProxy").Call(uintptr(webView), uintptr(unsafe.Pointer(proxy)))
}

// void mbNetSetMIMEType(mbNetJob jobPtr, const char* type);
func NetSetMIMEType(jobPtr NetJob, typeStr string) {
	mbLib.MustFindProc("mbNetSetMIMEType").Call(uintptr(jobPtr), StrToPtr(typeStr))
}

// 只能在blink线程调用（非主线程）
// const char* mbNetGetMIMEType(mbNetJob jobPtr);
func NetGetMIMEType(jobPtr NetJob) string {
	r1, _, _ := mbLib.MustFindProc("mbNetGetMIMEType").Call(uintptr(jobPtr))
	return StrFromPtr(r1)
}

// const utf8* mbNetGetHTTPHeaderField(mbNetJob job, const char* key, BOOL fromRequestOrResponse);
func NetGetHTTPHeaderField(job NetJob, key string, fromRequestOrResponse bool) string {
	r1, _, _ := mbLib.MustFindProc("mbNetGetHTTPHeaderField").Call(uintptr(job), StrToPtr(key), BoolToPtr(fromRequestOrResponse))
	return StrFromPtr(r1)
}

// void mbNetSetHTTPHeaderField(mbNetJob jobPtr, const WCHAR* key, const WCHAR* value, BOOL response);
func NetSetHTTPHeaderField(jobPtr NetJob, key, value string, response bool) {
	mbLib.MustFindProc("mbNetSetHTTPHeaderField").Call(uintptr(jobPtr), StrToWPtr(key), StrToWPtr(value), BoolToPtr(response))
}

// void mbNetSetHTTPHeaderFieldUtf8(mbNetJob jobPtr, const utf8* key, const utf8* value, BOOL response)
func NetSetHTTPHeaderFieldUtf8(jobPtr NetJob, key, value string, response bool) {
	mbLib.MustFindProc("mbNetSetHTTPHeaderFieldUtf8").Call(uintptr(jobPtr), StrToPtr(key), StrToPtr(value), BoolToPtr(response))
}

// void mbSetMouseEnabled(mbWebView webView, BOOL b);
func SetMouseEnabled(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetMouseEnabled").Call(uintptr(webView), BoolToPtr(b))
}

// void mbSetTouchEnabled(mbWebView webView, BOOL b);
func SetTouchEnabled(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetTouchEnabled").Call(uintptr(webView), BoolToPtr(b))
}

// void mbSetSystemTouchEnabled(mbWebView webView, BOOL b);
func SetSystemTouchEnabled(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetSystemTouchEnabled").Call(uintptr(webView), BoolToPtr(b))
}

// void mbSetContextMenuEnabled(mbWebView webView, BOOL b);
func SetContextMenuEnabled(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetContextMenuEnabled").Call(uintptr(webView), BoolToPtr(b))
}

// void mbSetNavigationToNewWindowEnable(mbWebView webView, BOOL b);
func SetNavigationToNewWindowEnabled(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetNavigationToNewWindowEnabled").Call(uintptr(webView), BoolToPtr(b))
}

// 可以关闭渲染
// void mbSetHeadlessEnabled(mbWebView webView, BOOL b);
func SetHeadlessEnabled(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetHeadlessEnabled").Call(uintptr(webView), BoolToPtr(b))
}

// 可以关闭拖拽文件、文字
// void mbSetDragDropEnable(mbWebView webView, BOOL b);
func SetDragDropEnable(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetDragDropEnable").Call(uintptr(webView), BoolToPtr(b))
}

// 可关闭自动响应WM_DROPFILES消息让网页加载本地文件
// void mbSetDragEnable(mbWebView webView, BOOL b);
func SetDragEnable(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetDragEnable").Call(uintptr(webView), BoolToPtr(b))
}

// 设置某项menu是否显示
// void mbSetContextMenuItemShow(mbWebView webView, mbMenuItemId item, BOOL isShow);
func SetContextMenuItemShow(webView WebView, item MenuItemId, isShow bool) {
	mbLib.MustFindProc("mbSetContextMenuItemShow").Call(uintptr(webView), uintptr(item), BoolToPtr(isShow))
}

// void mbSetHandle(mbWebView webView, HWND wnd);
func SetHandle(webView WebView, wnd HWND) {
	mbLib.MustFindProc("mbSetHandle").Call(uintptr(webView), uintptr(wnd))
}

// void mbSetHandleOffset(mbWebView webView, int x, int y);
func SetHandleOffset(webView WebView, x, y int) {
	mbLib.MustFindProc("mbSetHandleOffset").Call(uintptr(webView), uintptr(x), uintptr(y))
}

// linux下是gtk句柄，windows下是hwnd
// void* mbGetPlatformWindowHandle(mbWebView webView)
func GetPlatformWindowHandle(webView WebView) uintptr {
	r1, _, _ := mbLib.MustFindProc("mbGetPlatformWindowHandle").Call(uintptr(webView))
	return r1
}

// linux下永远返回nullptr，windows下是hwnd
// HWND mbGetHostHWND(mbWebView webView);
func GetHostHWND(webView WebView) HWND {
	r1, _, _ := mbLib.MustFindProc("mbGetHostHWND").Call(uintptr(webView))
	return HWND(r1)
}

// void mbSetTransparent(mbWebView webviewHandle, BOOL transparent);
func SetTransparent(webviewHandle WebView, transparent bool) {
	mbLib.MustFindProc("mbSetTransparent").Call(uintptr(webviewHandle), BoolToPtr(transparent))
}

// void mbSetViewSettings(mbWebView webviewHandle, const mbViewSettings* settings);
func SetViewSettings(webviewHandle WebView, settings *ViewSettings) {
	mbLib.MustFindProc("mbSetViewSettings").Call(uintptr(webviewHandle), uintptr(unsafe.Pointer(settings)))
}

// void mbSetCspCheckEnable(mbWebView webView, BOOL b);
func SetCspCheckEnable(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetCspCheckEnable").Call(uintptr(webView), BoolToPtr(b))
}

// void mbSetNpapiPluginsEnabled(mbWebView webView, BOOL b);
func SetNpapiPluginsEnabled(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetNpapiPluginsEnabled").Call(uintptr(webView), BoolToPtr(b))
}

// void mbSetMemoryCacheEnable(mbWebView webView, BOOL b);
func SetMemoryCacheEnable(webView WebView, b bool) {
	mbLib.MustFindProc("mbSetMemoryCacheEnable").Call(uintptr(webView), BoolToPtr(b))
}

// cookie格式必须是:PRODUCTINFO=webxpress; domain=.fidelity.com; path=/; secure
// void mbSetCookie(mbWebView webView, const utf8* url, const utf8* cookie);
func SetCookie(webView WebView, url, cookie string) {
	mbLib.MustFindProc("mbSetCookie").Call(uintptr(webView), StrToPtr(url), StrToPtr(cookie))
}

// void mbSetCookieEnabled(mbWebView webView, BOOL enable);
func SetCookieEnabled(webView WebView, enable bool) {
	mbLib.MustFindProc("mbSetCookieEnabled").Call(uintptr(webView), BoolToPtr(enable))
}

// void mbSetCookieJarPath(mbWebView webView, const WCHAR* path);
func SetCookieJarPath(webView WebView, path string) {
	mbLib.MustFindProc("mbSetCookieJarPath").Call(uintptr(webView), StrToWPtr(path))
}

// void mbSetCookieJarFullPath(mbWebView webView, const WCHAR* path);
func SetCookieJarFullPath(webView WebView, path string) {
	mbLib.MustFindProc("mbSetCookieJarFullPath").Call(uintptr(webView), StrToWPtr(path))
}

// void mbSetLocalStorageFullPath(mbWebView webView, const WCHAR* path);
func SetLocalStorageFullPath(webView WebView, path string) {
	mbLib.MustFindProc("mbSetLocalStorageFullPath").Call(uintptr(webView), StrToWPtr(path))
}

// const utf8* mbGetTitle(mbWebView webView);
func GetTitle(webView WebView) string {
	r1, _, _ := mbLib.MustFindProc("mbGetTitle").Call(uintptr(webView))
	return StrFromPtr(r1)
}

// void mbSetWindowTitle(mbWebView webView, const utf8* title);
func SetWindowTitle(webView WebView, title string) {
	mbLib.MustFindProc("mbSetWindowTitle").Call(uintptr(webView), StrToPtr(title))
}

// void mbSetWindowTitleW(mbWebView webView, const wchar_t* title);
func SetWindowTitleW(webView WebView, title string) {
	mbLib.MustFindProc("mbSetWindowTitleW").Call(uintptr(webView), StrToWPtr(title))
}

// const utf8* mbGetUrl(mbWebView webView);
func GetUrl(webView WebView) string {
	r1, _, _ := mbLib.MustFindProc("mbGetUrl").Call(uintptr(webView))
	return StrFromPtr(r1)
}

// int mbGetCursorInfoType(mbWebView webView);
func GetCursorInfoType(webView WebView) int {
	r1, _, _ := mbLib.MustFindProc("mbGetCursorInfoType").Call(uintptr(webView))
	return int(r1)
}

// void mbAddPluginDirectory(mbWebView webView, const WCHAR* path);
func AddPluginDirectory(webView WebView, path string) {
	mbLib.MustFindProc("mbAddPluginDirectory").Call(uintptr(webView), StrToPtr(path))
}

// void mbSetUserAgent(mbWebView webView, const utf8* userAgent);
func SetUserAgent(webView WebView, userAgent string) {
	mbLib.MustFindProc("mbSetUserAgent").Call(uintptr(webView), StrToPtr(userAgent))
}

// void mbSetZoomFactor(mbWebView webView, float factor);
func SetZoomFactor(webView WebView, factor float32) {
	mbLib.MustFindProc("mbSetZoomFactor").Call(uintptr(webView), F32ToPtr(factor))
}

// float mbGetZoomFactor(mbWebView webView);
func GetZoomFactor(webView WebView) float32 {
	_, r2, _ := mbLib.MustFindProc("mbGetZoomFactor").Call(uintptr(webView))
	return F32FromPtr(r2)
}

// void mbSetDiskCacheEnabled(mbWebView webView, BOOL enable);
func SetDiskCacheEnabled(webView WebView, enable bool) {
	mbLib.MustFindProc("mbSetDiskCacheEnabled").Call(uintptr(webView), BoolToPtr(enable))
}

// void mbSetDiskCachePath(mbWebView webView, const WCHAR* path);
func SetDiskCachePath(webView WebView, path string) {
	mbLib.MustFindProc("mbSetDiskCachePath").Call(uintptr(webView), StrToPtr(path))
}

// void mbSetDiskCacheLimit(mbWebView webView, size_t limit);
func SetDiskCacheLimit(webView WebView, limit uint64) {
	mbLib.MustFindProc("mbSetDiskCacheLimit").Call(uintptr(webView), uintptr(limit))
}

// void mbSetDiskCacheLimitDisk(mbWebView webView, size_t limit);
func SetDiskCacheLimitDisk(webView WebView, limit uint64) {
	mbLib.MustFindProc("mbSetDiskCacheLimitDisk").Call(uintptr(webView), uintptr(limit))
}

// void mbSetDiskCacheLevel(mbWebView webView, int Level);
func SetDiskCacheLevel(webView WebView, level int) {
	mbLib.MustFindProc("mbSetDiskCacheLevel").Call(uintptr(webView), uintptr(level))
}

// void mbSetResourceGc(mbWebView webView, int intervalSec);
func SetResourceGc(webView WebView, intervalSec int) {
	mbLib.MustFindProc("mbSetResourceGc").Call(uintptr(webView), uintptr(intervalSec))
}

// void mbCanGoForward(mbWebView webView, mbCanGoBackForwardCallback callback, void* param);
func CanGoForward(webView WebView, callback CanGoBackForwardCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbCanGoBackForwardCallback = func(webView WebView, param uintptr, state AsynRequestState, b BOOL) (void uintptr) {
			callback(webView, state, BoolFromPtr(uintptr(b)))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbCanGoForward").Call(uintptr(webView), cbPtr, 0)
}

// void mbCanGoBack(mbWebView webView, mbCanGoBackForwardCallback callback, void* param);
func CanGoBack(webView WebView, callback CanGoBackForwardCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbCanGoBackForwardCallback = func(webView WebView, param uintptr, state AsynRequestState, b BOOL) (void uintptr) {
			callback(webView, state, BoolFromPtr(uintptr(b)))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbCanGoBack").Call(uintptr(webView), cbPtr, 0)
}

// void mbGetCookie(mbWebView webView, mbGetCookieCallback callback, void* param);
func GetCookie(webView WebView, callback GetCookieCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbGetCookieCallback = func(webView WebView, param uintptr, state AsynRequestState, cookie uintptr) (void uintptr) {
			callback(webView, state, StrFromPtr(cookie))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbGetCookie").Call(uintptr(webView), cbPtr, 0)
}

// const utf8* mbGetCookieOnBlinkThread(mbWebView webView);
func GetCookieOnBlinkThread(webView WebView) string {
	r1, _, _ := mbLib.MustFindProc("mbGetCookieOnBlinkThread").Call(uintptr(webView))
	return StrFromPtr(r1)
}

// void mbClearCookie(mbWebView webView);
func ClearCookie(webView WebView) {
	mbLib.MustFindProc("mbClearCookie").Call(uintptr(webView))
}

// void mbResize(mbWebView webView, int w, int h);
func Resize(webView WebView, w, h int) {
	mbLib.MustFindProc("mbResize").Call(uintptr(webView), uintptr(w), uintptr(h))
}

// void mbGetSize(mbWebView webView, mbRect* rc)
func GetSize(webView WebView, rc *Rect) {
	mbLib.MustFindProc("mbGetSize").Call(uintptr(webView), uintptr(unsafe.Pointer(rc)))
}

// void mbOnNavigation(mbWebView webView, mbNavigationCallback callback, void* param);
func OnNavigation(webView WebView, callback NavigationCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbNavigationCallback = func(webView WebView, param uintptr, navigationType NavigationType, url uintptr) BOOL {
			if callback(webView, navigationType, StrFromPtr(url)) {
				return TRUE
			} else {
				return FALSE
			}
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnNavigation").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnNavigationSync(mbWebView webView, mbNavigationCallback callback, void* param);
func OnNavigationSync(webView WebView, callback NavigationCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbNavigationCallback = func(webView WebView, param uintptr, navigationType NavigationType, url uintptr) BOOL {
			if callback(webView, navigationType, StrFromPtr(url)) {
				return TRUE
			} else {
				return FALSE
			}
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnNavigationSync").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnCreateView(mbWebView webView, mbCreateViewCallback callback, void* param);
func OnCreateView(webView WebView, callback CreateViewCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbCreateViewCallback = func(webView WebView, param uintptr, navigationType NavigationType, url uintptr, windowFeatures *WindowFeatures) WebView {
			return callback(webView, navigationType, StrFromPtr(url), windowFeatures)
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnCreateView").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnDocumentReady(mbWebView webView, mbDocumentReadyCallback callback, void* param);
func OnDocumentReady(webView WebView, callback DocumentReadyCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbDocumentReadyCallback = func(webView WebView, param uintptr, frameId WebFrameHandle) (void uintptr) {
			callback(webView, frameId)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnDocumentReady").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnPaintUpdated(mbWebView webView, mbPaintUpdatedCallback callback, void* callbackParam);
func OnPaintUpdated(webView WebView, callback PaintUpdatedCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbPaintUpdatedCallback = func(webView WebView, param uintptr, hdc HDC, x, y, cx, cy int32) (void uintptr) {
			callback(webView, hdc, x, y, cx, cy)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnPaintUpdated").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnPaintBitUpdated(mbWebView webView, mbPaintBitUpdatedCallback callback, void* callbackParam);
func OnPaintBitUpdated(webView WebView, callback PaintBitUpdatedCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbPaintBitUpdatedCallback = func(webView WebView, param, buffer uintptr, r *Rect, width, height int32) (void uintptr) {
			callback(webView, buffer, r, width, height)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnPaintBitUpdated").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnAcceleratedPaint(mbWebView webView, mbAcceleratedPaintCallback callback, void* callbackParam);
func OnAcceleratedPaint(webView WebView, callback AcceleratedPaintCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbAcceleratedPaintCallback = func(webView WebView, param uintptr, typ int32, dirytRects *Rect, dirytRectsSize, sharedHandle uintptr) (void uintptr) {
			callback(webView, typ, dirytRects, dirytRectsSize, sharedHandle)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnAcceleratedPaint").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnLoadUrlBegin(mbWebView webView, mbLoadUrlBeginCallback callback, void* callbackParam);
func OnLoadUrlBegin(webView WebView, callback LoadUrlBeginCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbLoadUrlBeginCallback = func(webView WebView, param uintptr, url uintptr, job NetJob) BOOL {
			if callback(webView, StrFromPtr(url), job) {
				return TRUE
			} else {
				return FALSE
			}
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnLoadUrlBegin").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnLoadUrlEnd(mbWebView webView, mbLoadUrlEndCallback callback, void* callbackParam);
func OnLoadUrlEnd(webView WebView, callback LoadUrlEndCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbLoadUrlEndCallback = func(webView WebView, param uintptr, url uintptr, job NetJob, buf uintptr, len int) (void uintptr) {
			data := unsafe.Slice((*byte)(unsafe.Pointer(buf)), len)
			callback(webView, StrFromPtr(url), job, data)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnLoadUrlEnd").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnLoadUrlFail(mbWebView webView, mbLoadUrlFailCallback callback, void* callbackParam);
func OnLoadUrlFail(webView WebView, callback LoadUrlFailCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbLoadUrlFailCallback = func(webView WebView, param uintptr, url uintptr, job NetJob) (void uintptr) {
			callback(webView, StrFromPtr(url), job)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnLoadUrlFail").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnTitleChanged(mbWebView webView, mbTitleChangedCallback callback, void* callbackParam);
func OnTitleChanged(webView WebView, callback TitleChangedCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbTitleChangedCallback = func(webView WebView, param uintptr, title uintptr) (void uintptr) {
			callback(webView, StrFromPtr(title))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnTitleChanged").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnURLChanged(mbWebView webView, mbURLChangedCallback callback, void* callbackParam);
func OnURLChanged(webView WebView, callback URLChangedCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbURLChangedCallback = func(webView WebView, param uintptr, url uintptr, canGoBack, canGoForward bool) (void uintptr) {
			callback(webView, StrFromPtr(url), canGoBack, canGoForward)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnURLChanged").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnLoadingFinish(mbWebView webView, mbLoadingFinishCallback callback, void* param);
func OnLoadingFinish(webView WebView, callback LoadingFinishCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb = func(webView WebView, param uintptr, FrameId WebFrameHandle, url uintptr, Result LoadingResult, FailedReason uintptr) (void uintptr) {
			callback(webView, FrameId, StrFromPtr(url), Result, StrFromPtr(FailedReason))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnLoadingFinish").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnDownload(mbWebView webView, mbDownloadCallback callback, void* param);
func OnDownload(webView WebView, callback DownloadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb = func(webView WebView, param uintptr, frameId WebFrameHandle, url uintptr, downloadJob uintptr) BOOL {
			if callback(webView, frameId, StrFromPtr(url), downloadJob) {
				return TRUE
			}
			return FALSE
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnDownload").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnDownloadInBlinkThread(mbWebView webView, mbDownloadInBlinkThreadCallback callback, void* param);
func OnDownloadInBlinkThread(webView WebView, callback DownloadInBlinkThreadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbDownloadInBlinkThreadCallback = func(webView WebView, params uintptr, expectedContentLength uintptr, url uintptr, mime uintptr, disposition uintptr, job NetJob, dataBind *NetJobDataBind) uintptr {
			opt := callback(webView, expectedContentLength, StrFromPtr(url), StrFromPtr(mime), StrFromPtr(disposition), job, dataBind)
			return uintptr(opt)
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnDownloadInBlinkThread").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnAlertBox(mbWebView webView, mbAlertBoxCallback callback, void* param);
func OnAlertBox(webView WebView, callback AlertBoxCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbAlertBoxCallback = func(webView WebView, param uintptr, msg uintptr) (void uintptr) {
			callback(webView, StrFromPtr(msg))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnAlertBox").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnConfirmBox(mbWebView webView, mbConfirmBoxCallback callback, void* param);
func OnConfirmBox(webView WebView, callback ConfirmBoxCallback) {

	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbConfirmBoxCallback = func(webView WebView, param uintptr, msg uintptr) BOOL {
			if callback(webView, StrFromPtr(msg)) {
				return TRUE
			}
			return FALSE
		}

		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnConfirmBox").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnPromptBox(mbWebView webView, mbPromptBoxCallback callback, void* param);
func OnPromptBox(webView WebView, callback PromptBoxCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbPromptBoxCallback = func(webView WebView, param, msg, defaultResult uintptr, result *BOOL) uintptr {
			rst := callback(webView, StrFromPtr(msg), StrFromPtr(defaultResult))
			return uintptr(rst)
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnPromptBox").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnNetGetFavicon(mbWebView webView, mbNetGetFaviconCallback callback, void* param);
func OnNetGetFavicon(webView WebView, callback NetGetFaviconCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbNetGetFaviconCallback = func(webView WebView, param uintptr, url uintptr, buf *MemBuf) (void uintptr) {
			callback(webView, StrFromPtr(url), buf)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnNetGetFavicon").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnConsole(mbWebView webView, mbConsoleCallback callback, void* param);
func OnConsole(webView WebView, callback ConsoleCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbConsoleCallback = func(webView WebView, param uintptr, level ConsoleLevel, message uintptr, sourceName uintptr, sourceLine uint, stackTrace uintptr) (void uintptr) {
			callback(webView, level, StrFromPtr(message), StrFromPtr(sourceName), sourceLine, StrFromPtr(stackTrace))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnConsole").Call(uintptr(webView), cbPtr, 0)
}

// BOOL mbOnClose(mbWebView webView, mbCloseCallback callback, void* param);
func OnClose(webView WebView, callback CloseCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbCloseCallback = func(webView WebView, param uintptr, unuse uintptr) BOOL {
			if callback(webView, unuse) {
				return TRUE
			}
			return FALSE
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnClose").Call(uintptr(webView), cbPtr, 0)
}

// BOOL mbOnDestroy(mbWebView webView, mbDestroyCallback callback, void* param);
func OnDestroy(webView WebView, callback DestroyCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbDestroyCallback = func(webView WebView, param uintptr, unuse uintptr) BOOL {
			if callback(webView, unuse) {
				return TRUE
			}
			return FALSE
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnDestroy").Call(uintptr(webView), cbPtr, 0)
}

// BOOL mbOnPrinting(mbWebView webView, mbPrintingCallback callback, void* param);
func OnPrinting(webView WebView, callback PrintingCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb = func(webview WebView, param uintptr, step PrintintStep, hDC Handle, settings *PrintintSettings, pageCount int32) BOOL {
			if callback(webview, step, hDC, settings, pageCount) {
				return TRUE
			}
			return FALSE
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnPrinting").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnDidCreateScriptContext(mbWebView webView, mbDidCreateScriptContextCallback callback, void* callbackParam);
func OnDidCreateScriptContext(webView WebView, callback DidCreateScriptContextCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbDidCreateScriptContextCallback = func(webView WebView, param uintptr, frameId WebFrameHandle, context uintptr, extensionGroup, worldId int) (void uintptr) {
			callback(webView, frameId, context, extensionGroup, worldId)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnDidCreateScriptContext").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnPluginList(mbWebView webView, mbGetPluginListCallback callback, void* callbackParam);
func OnPluginList(webView WebView, callback GetPluginListCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbGetPluginListCallback = func(refresh bool, pluginListBuilder, param uintptr) BOOL {
			if callback(refresh, pluginListBuilder) {
				return TRUE
			}
			return FALSE
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnPluginList").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnImageBufferToDataURL(mbWebView webView, mbImageBufferToDataURLCallback callback, void* callbackParam);
func OnImageBufferToDataURL(webView WebView, callback ImageBufferToDataURLCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb = func(webView WebView, param uintptr, data *byte, size uintptr) uintptr {
			byts := unsafe.Slice(data, size)
			return uintptr(callback(webView, byts))
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnImageBufferToDataURL").Call(uintptr(webView), cbPtr, 0)
}

// void mbGoBack(mbWebView webView);
func GoBack(webView WebView) {
	mbLib.MustFindProc("mbGoBack").Call(uintptr(webView))
}

// void mbGoForward(mbWebView webView);
func GoForward(webView WebView) {
	mbLib.MustFindProc("mbGoForward").Call(uintptr(webView))
}

// void mbNavigateAtIndex(mbWebView webView, int index);
func NavigateAtIndex(webView WebView, index int) {
	mbLib.MustFindProc("mbNavigateAtIndex").Call(uintptr(webView), uintptr(index))
}

// int mbGetNavigateIndex(mbWebView webView);
func GetNavigateIndex(webView WebView) int {
	r1, _, _ := mbLib.MustFindProc("mbGetNavigateIndex").Call(uintptr(webView))
	return int(r1)
}

// void mbStopLoading(mbWebView webView);
func StopLoading(webView WebView) {
	mbLib.MustFindProc("mbStopLoading").Call(uintptr(webView))
}

// void mbReload(mbWebView webView);
func Reload(webView WebView) {
	mbLib.MustFindProc("mbReload").Call(uintptr(webView))
}

// void mbPerformCookieCommand(mbWebView webView, mbCookieCommand command);
func PerformCookieCommand(webView WebView, command CookieCommand) {
	mbLib.MustFindProc("mbPerformCookieCommand").Call(uintptr(webView), uintptr(command))
}

// void mbEditorSelectAll(mbWebView webView);
func EditorSelectAll(webView WebView) {
	mbLib.MustFindProc("mbEditorSelectAll").Call(uintptr(webView))
}

// void mbEditorCopy(mbWebView webView);
func EditorCopy(webView WebView) {
	mbLib.MustFindProc("mbEditorCopy").Call(uintptr(webView))
}

// void mbEditorCut(mbWebView webView);
func EditorCut(webView WebView) {
	mbLib.MustFindProc("mbEditorCut").Call(uintptr(webView))
}

// void mbEditorPaste(mbWebView webView);
func EditorPaste(webView WebView) {
	mbLib.MustFindProc("mbEditorPaste").Call(uintptr(webView))
}

// void mbEditorDelete(mbWebView webView);
func EditorDelete(webView WebView) {
	mbLib.MustFindProc("mbEditorDelete").Call(uintptr(webView))
}

// void mbEditorUndo(mbWebView webView);
func EditorUndo(webView WebView) {
	mbLib.MustFindProc("mbEditorUndo").Call(uintptr(webView))
}

// BOOL mbFireMouseEvent(mbWebView webView, unsigned int message, int x, int y, unsigned int flags);
func FireMouseEvent(webView WebView, message uint, x, y int, flags uint) bool {
	r1, _, _ := mbLib.MustFindProc("mbFireMouseEvent").Call(uintptr(webView), uintptr(message), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r1)
}

// BOOL mbFireContextMenuEvent(mbWebView webView, int x, int y, unsigned int flags);
func FireContextMenuEvent(webView WebView, x, y int, flags uint) bool {
	r1, _, _ := mbLib.MustFindProc("mbFireContextMenuEvent").Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r1)
}

// BOOL mbFireMouseWheelEvent(mbWebView webView, int x, int y, int delta, unsigned int flags);
func FireMouseWheelEvent(webView WebView, x, y, delta int, flags uint) bool {
	r1, _, _ := mbLib.MustFindProc("mbFireMouseWheelEvent").Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
	return BoolFromPtr(r1)
}

// BOOL mbFireKeyUpEvent(mbWebView webView, unsigned int virtualKeyCode, unsigned int flags, BOOL systemKey);
func FireKeyUpEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r1, _, _ := mbLib.MustFindProc("mbFireKeyUpEvent").Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r1)
}

// BOOL mbFireKeyDownEvent(mbWebView webView, unsigned int virtualKeyCode, unsigned int flags, BOOL systemKey);
func FireKeyDownEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r1, _, _ := mbLib.MustFindProc("mbFireKeyDownEvent").Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r1)
}

// BOOL mbFireKeyPressEvent(mbWebView webView, unsigned int charCode, unsigned int flags, BOOL systemKey);
func FireKeyPressEvent(webView WebView, charCode, flags uint, systemKey bool) bool {
	r1, _, _ := mbLib.MustFindProc("mbFireKeyPressEvent").Call(uintptr(webView), uintptr(charCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r1)
}

// BOOL mbFireWindowsMessage(mbWebView webView, HWND hWnd, UINT message, WPARAM wParam, LPARAM lParam, LRESULT* result);
func FireWindowsMessage(webView WebView, hWnd HWND, message uint, wParam WPARAM, lParam LPARAM, result *LRESULT) bool {
	r1, _, _ := mbLib.MustFindProc("mbFireWindowsMessage").Call(uintptr(webView), uintptr(hWnd), uintptr(message), uintptr(wParam), uintptr(lParam), uintptr(unsafe.Pointer(result)))
	return BoolFromPtr(r1)
}

// void mbSetFocus(mbWebView webView);
func SetFocus(webView WebView) {
	mbLib.MustFindProc("mbSetFocus").Call(uintptr(webView))
}

// void mbKillFocus(mbWebView webView);
func KillFocus(webView WebView) {
	mbLib.MustFindProc("mbKillFocus").Call(uintptr(webView))
}

// void mbShowWindow(mbWebView webview, BOOL show);
func ShowWindow(webView WebView, show bool) {
	mbLib.MustFindProc("mbShowWindow").Call(uintptr(webView), BoolToPtr(show))
}

// void mbLoadURL(mbWebView webView, const utf8* url);
func LoadURL(webView WebView, url string) {
	mbLib.MustFindProc("mbLoadURL").Call(uintptr(webView), StrToPtr(url))
}

// void mbLoadHtmlWithBaseUrl(mbWebView webView, const utf8* html, const utf8* baseUrl);
func LoadHtmlWithBaseUrl(webView WebView, html string, baseUrl string) {
	mbLib.MustFindProc("mbLoadHtmlWithBaseUrl").Call(uintptr(webView), StrToPtr(html), StrToPtr(baseUrl))
}

// void mbPostURL(mbWebView webView, const utf8* url, const char* postData, int postLen);
func PostURL(webView WebView, url string, postData []byte) {
	mbLib.MustFindProc("mbPostURL").Call(uintptr(webView), StrToPtr(url), uintptr(unsafe.Pointer(&postData[0])), uintptr(len(postData)))
}

// HDC mbGetLockedViewDC(mbWebView webView);
func GetLockedViewDC(webView WebView) HDC {
	r1, _, _ := mbLib.MustFindProc("mbGetLockedViewDC").Call(uintptr(webView))
	return HDC(r1)
}

// void mbUnlockViewDC(mbWebView webView);
func UnlockViewDC(webView WebView) {
	mbLib.MustFindProc("mbUnlockViewDC").Call(uintptr(webView))
}

// void mbWake(mbWebView webView);
func Wake(webView WebView) {
	mbLib.MustFindProc("mbWake").Call(uintptr(webView))
}

// double mbJsToDouble(mbJsExecState es, mbJsValue v);
func JsToDouble(es JsExecState, v JsValue) float64 {
	_, r2, _ := mbLib.MustFindProc("mbJsToDouble").Call(uintptr(es), uintptr(v))
	return F64FromPtr(r2)
}

// BOOL mbJsToBoolean(mbJsExecState es, mbJsValue v);
func JsToBoolean(es JsExecState, v JsValue) bool {
	r1, _, _ := mbLib.MustFindProc("mbJsToBoolean").Call(uintptr(es), uintptr(v))
	return BoolFromPtr(r1)
}

// const utf8* mbJsToString(mbJsExecState es, mbJsValue v);
func JsToString(es JsExecState, v JsValue) string {
	r1, _, _ := mbLib.MustFindProc("mbJsToString").Call(uintptr(es), uintptr(v))
	return StrFromPtr(r1)
}

// mbJsType mbGetJsValueType(mbJsExecState es, mbJsValue v);
func GetJsValueType(es JsExecState, v JsValue) JsType {
	r1, _, _ := mbLib.MustFindProc("mbGetJsValueType").Call(uintptr(es), uintptr(v))
	return JsType(r1)
}

// void mbOnJsQuery(mbWebView webView, mbJsQueryCallback callback, void* param);
func OnJsQuery(webView WebView, callback JsQueryCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbJsQueryCallback = func(webView WebView, param uintptr, es JsExecState, queryId int64, customMsg int, request uintptr) (void uintptr) {
			callback(webView, es, queryId, customMsg, StrFromPtr(request))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnJsQuery").Call(uintptr(webView), cbPtr, 0)
}

// void mbResponseQuery(mbWebView webView, int64_t queryId, int customMsg, const utf8* response);
func ResponseQuery(webView WebView, queryId int64, customMsg int, response string) {
	mbLib.MustFindProc("mbResponseQuery").Call(uintptr(webView), uintptr(queryId), uintptr(customMsg), StrToPtr(response))
}

// void mbRunJs(mbWebView webView, mbWebFrameHandle frameId, const utf8* script, BOOL isInClosure, mbRunJsCallback callback, void* param, void* unuse);
func RunJs(webView WebView, frameId WebFrameHandle, script string, isInClosure bool, callback RunJsCallback, unuse uintptr) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbRunJsCallback = func(webView WebView, param uintptr, es JsExecState, v JsValue) (void uintptr) {
			callback(webView, es, v)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbRunJs").Call(uintptr(webView), uintptr(frameId), StrToPtr(script), BoolToPtr(isInClosure), cbPtr, 0, unuse)
}

// mbJsValue mbRunJsSync(mbWebView webView, mbWebFrameHandle frameId, const utf8* script, BOOL isInClosure);
func RunJsSync(webView WebView, frameId WebFrameHandle, script string, isInClosure bool) JsValue {
	r1, _, _ := mbLib.MustFindProc("mbRunJsSync").Call(uintptr(webView), uintptr(frameId), StrToPtr(script), BoolToPtr(isInClosure))
	return JsValue(r1)
}

// mbWebFrameHandle mbWebFrameGetMainFrame(mbWebView webView);
func WebFrameGetMainFrame(webView WebView) WebFrameHandle {
	r1, _, _ := mbLib.MustFindProc("mbWebFrameGetMainFrame").Call(uintptr(webView))
	return WebFrameHandle(r1)
}

// BOOL mbIsMainFrame(mbWebView webView, mbWebFrameHandle frameId);
func IsMainFrame(webView WebView, frameId WebFrameHandle) bool {
	r1, _, _ := mbLib.MustFindProc("mbIsMainFrame").Call(uintptr(webView), uintptr(frameId))
	return BoolFromPtr(r1)
}

// void mbSetNodeJsEnable(mbWebView webView, BOOL b);
func SetNodeJsEnable(webView WebView, enable bool) {
	mbLib.MustFindProc("mbSetNodeJsEnable").Call(uintptr(webView), BoolToPtr(enable))
}

// void mbSetDeviceParameter(mbWebView webView, const char* device, const char* paramStr, int paramInt, float paramFloat);
func SetDeviceParameter(webView WebView, device, paramStr string, paramInt int, paramFloat float32) {
	mbLib.MustFindProc("mbSetDeviceParameter").Call(uintptr(webView), StrToPtr(device), StrToPtr(paramStr), uintptr(paramInt), F32ToPtr(paramFloat))
}

// void mbGetContentAsMarkup(mbWebView webView, mbGetContentAsMarkupCallback calback, void* param, mbWebFrameHandle frameId);
func GetContentAsMarkup(webView WebView, callback GetContentAsMarkupCallback, frameId WebFrameHandle) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbGetContentAsMarkupCallback = func(webView WebView, param uintptr, content uintptr, size uintptr) (void uintptr) {
			data := unsafe.Slice((*byte)(unsafe.Pointer(content)), size)
			callback(webView, data)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbGetContentAsMarkup").Call(uintptr(webView), cbPtr, 0, uintptr(frameId))
}

// void mbGetSource(mbWebView webView, mbGetSourceCallback calback, void* param);
func GetSource(webView WebView, callback GetSourceCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbGetSourceCallback = func(webView WebView, param uintptr, mhtml uintptr) (void uintptr) {
			callback(webView, StrFromPtr(mhtml))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbGetSource").Call(uintptr(webView), cbPtr)
}

// void mbUtilSerializeToMHTML(mbWebView webView, mbGetSourceCallback calback, void* param);
func UtilSerializeToMHTML(webView WebView, callback GetSourceCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbGetSourceCallback = func(webView WebView, param uintptr, mhtml uintptr) (void uintptr) {
			callback(webView, StrFromPtr(mhtml))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbUtilSerializeToMHTML").Call(uintptr(webView), cbPtr)
}

// const char* mbUtilCreateRequestCode(const char* registerInfo);
func UtilCreateRequestCode(registerInfo string) string {
	r1, _, _ := mbLib.MustFindProc("mbUtilCreateRequestCode").Call(StrToPtr(registerInfo))
	return StrFromPtr(r1)
}

// BOOL mbUtilIsRegistered(const wchar_t* defaultPath);
func UtilIsRegistered(defaultPath string) bool {
	r1, _, _ := mbLib.MustFindProc("mbUtilIsRegistered").Call(StrToWPtr(defaultPath))
	return BoolFromPtr(r1)
}

// BOOL mbUtilPrint(mbWebView webView, mbWebFrameHandle frameId, const mbPrintSettings* printParams);
func UtilPrint(webView WebView, frameId WebFrameHandle, printParams *PrintSettings) bool {
	r1, _, _ := mbLib.MustFindProc("mbUtilPrint").Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(printParams)))
	return BoolFromPtr(r1)
}

// const utf8* mbUtilBase64Encode(const utf8* str);
func UtilBase64Encode(str string) string {
	r1, _, _ := mbLib.MustFindProc("mbUtilBase64Encode").Call(StrToPtr(str))
	return StrFromPtr(r1)
}

// const utf8* mbUtilBase64Decode(const utf8* str);
func UtilBase64Decode(str string) string {
	r1, _, _ := mbLib.MustFindProc("mbUtilBase64Decode").Call(StrToPtr(str))
	return StrFromPtr(r1)
}

// const utf8* mbUtilDecodeURLEscape(const utf8* url);
func UtilDecodeURLEscape(url string) string {
	r1, _, _ := mbLib.MustFindProc("mbUtilDecodeURLEscape").Call(StrToPtr(url))
	return StrFromPtr(r1)
}

// const utf8* mbUtilEncodeURLEscape(const utf8* url);
func UtilEncodeURLEscape(url string) string {
	r1, _, _ := mbLib.MustFindProc("mbUtilEncodeURLEscape").Call(StrToPtr(url))
	return StrFromPtr(r1)
}

// const mbMemBuf* mbUtilCreateV8Snapshot(const utf8* str);
func UtilCreateV8Snapshot(str string) *MemBuf {
	r1, _, _ := mbLib.MustFindProc("mbUtilCreateV8Snapshot").Call(StrToPtr(str))
	return (*MemBuf)(unsafe.Pointer(r1))
}

// void mbUtilPrintToPdf(mbWebView webView, mbWebFrameHandle frameId, const mbPrintSettings* settings, mbPrintPdfDataCallback callback, void* param);
func UtilPrintToPdf(webView WebView, frameId WebFrameHandle, settings *PrintSettings, callback PrintPdfDataCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb = func(webview WebView, param uintptr, datas uintptr) (void uintptr) {
			callback(webview, (*PdfDatas)(unsafe.Pointer(datas)))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbUtilPrintToPdf").Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)), cbPtr, 0)
}

// void mbUtilPrintToBitmap(mbWebView webView, mbWebFrameHandle frameId, const mbScreenshotSettings* settings, mbPrintBitmapCallback callback, void* param);
func UtilPrintToBitmap(webView WebView, frameId WebFrameHandle, settings *ScreenshotSettings, callback PrintBitmapCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb = func(webview WebView, param uintptr, data uintptr, size uintptr) (void uintptr) {
			d := unsafe.Slice((*byte)(unsafe.Pointer(data)), size)
			callback(webview, d)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbUtilPrintToBitmap").Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)), cbPtr, 0)
}

// void mbUtilScreenshot(mbWebView webView, const mbScreenshotSettings* settings, mbOnScreenshot callback, void* param);
func UtilScreenshot(webView WebView, settings *ScreenshotSettings, callback OnScreenshotCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb = func(webview WebView, param uintptr, data uintptr, size uintptr) (void uintptr) {
			d := unsafe.Slice((*byte)(unsafe.Pointer(data)), size)
			callback(webview, d)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbUtilScreenshot").Call(uintptr(webView), uintptr(unsafe.Pointer(settings)), cbPtr, 0)
}

// BOOL mbUtilsSilentPrint(mbWebView webView, const char* settings);
func UtilsSilentPrint(webView WebView, settings string) bool {
	r1, _, _ := mbLib.MustFindProc("mbUtilsSilentPrint").Call(uintptr(webView), StrToPtr(settings))
	return BoolFromPtr(r1)
}

// BOOL mbPopupDownloadMgr(mbWebView webView, const char* url, void* downloadJob);
func PopupDownloadMgr(webView WebView, url string, downloadJob uintptr) bool {
	r1, _, _ := mbLib.MustFindProc("mbPopupDownloadMgr").Call(uintptr(webView), StrToPtr(url), downloadJob)
	return BoolFromPtr(r1)
}

// mbDownloadOpt mbPopupDialogAndDownload(mbWebView webView, const mbDialogOptions* dialogOpt, size_t contentLength, const char* url, const char* mime, const char* disposition, mbNetJob job, mbNetJobDataBind* dataBind, mbDownloadBind* callbackBind);
func PopupDialogAndDownload(webView WebView, dialogOpt *DialogOptions, contentLength uint64, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	r1, _, _ := mbLib.MustFindProc("mbPopupDialogAndDownload").Call(uintptr(webView), uintptr(unsafe.Pointer(dialogOpt)), uintptr(contentLength), StrToPtr(url), StrToPtr(mime), StrToPtr(disposition), uintptr(job), uintptr(unsafe.Pointer(dataBind)), uintptr(unsafe.Pointer(callbackBind)))
	return DownloadOpt(r1)
}

// mbDownloadOpt mbDownloadByPath(mbWebView webView, const mbDownloadOptions* downloadOptions, const WCHAR* path, size_t expectedContentLength, const char* url, const char* mime, const char* disposition, mbNetJob job, mbNetJobDataBind* dataBind, mbDownloadBind* callbackBind);
func DownloadByPath(webView WebView, downloadOptions *DownloadOptions, path string, expectedContentLength uintptr, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	r1, _, _ := mbLib.MustFindProc("mbDownloadByPath").Call(uintptr(webView), uintptr(unsafe.Pointer(downloadOptions)), StrToPtr(path), expectedContentLength, StrToPtr(url), StrToPtr(mime), StrToPtr(disposition), uintptr(job), uintptr(unsafe.Pointer(dataBind)), uintptr(unsafe.Pointer(callbackBind)))
	return DownloadOpt(r1)
}

// void mbGetPdfPageData(mbWebView webView, mbOnGetPdfPageDataCallback callback, void* param);
func GetPdfPageData(webView WebView, callback OnGetPdfPageDataCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbOnGetPdfPageDataCallback = func(webView WebView, param uintptr, data uintptr, size uintptr) (void uintptr) {
			byts := unsafe.Slice((*byte)(unsafe.Pointer(data)), size)
			callback(webView, byts)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbGetPdfPageData").Call(uintptr(webView), cbPtr, 0)
}

// mbMemBuf* mbCreateMemBuf(mbWebView webView, void* buf, size_t length);
func CreateMemBuf(webView WebView, buf []byte) *MemBuf {
	r1, _, _ := mbLib.MustFindProc("mbCreateMemBuf").Call(uintptr(webView), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return (*MemBuf)(unsafe.Pointer(r1))
}

// void mbFreeMemBuf(mbMemBuf* buf);
func FreeMemBuf(buf *MemBuf) {
	mbLib.MustFindProc("mbFreeMemBuf").Call(uintptr(unsafe.Pointer(buf)))
}

// void mbPluginListBuilderAddPlugin(void* builder, const utf8* name, const utf8* description, const utf8* fileName);
func PluginListBuilderAddPlugin(builder uintptr, name, description, fileName string) {
	mbLib.MustFindProc("mbPluginListBuilderAddPlugin").Call(builder, StrToPtr(name), StrToPtr(description), StrToPtr(fileName))
}

// void mbPluginListBuilderAddMediaTypeToLastPlugin(void* builder, const utf8* name, const utf8* description);
func PluginListBuilderAddMediaTypeToLastPlugin(builder uintptr, name, description string) {
	mbLib.MustFindProc("mbPluginListBuilderAddMediaTypeToLastPlugin").Call(builder, StrToPtr(name), StrToPtr(description))
}

// void mbPluginListBuilderAddFileExtensionToLastMediaType(void* builder, const utf8* fileExtension);
func PluginListBuilderAddFileExtensionToLastMediaType(builder uintptr, fileExtension string) {
	mbLib.MustFindProc("mbPluginListBuilderAddFileExtensionToLastMediaType").Call(builder, StrToPtr(fileExtension))
}

// void mbEnableHighDPISupport();
func EnableHighDPISupport() {
	mbLib.MustFindProc("mbEnableHighDPISupport").Call()
}

// void mbRunMessageLoop();
func RunMessageLoop() {
	mbLib.MustFindProc("mbRunMessageLoop").Call()
}

// void mbExitMessageLoop();
func ExitMessageLoop() {
	mbLib.MustFindProc("mbExitMessageLoop").Call()
}

// void mbOnLoadUrlFinish(mbWebView webView, mbLoadUrlFinishCallback callback, void* callbackParam);
func OnLoadUrlFinish(webView WebView, callback LoadUrlFinishCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbLoadUrlFinishCallback = func(webView WebView, param, url uintptr, job NetJob, length int) (void uintptr) {
			callback(webView, StrFromPtr(url), job, length)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnLoadUrlFinish").Call(uintptr(webView), cbPtr, 0)
}

// void mbOnLoadUrlHeadersReceived(mbWebView webView, mbLoadUrlHeadersReceivedCallback callback, void* callbackParam);
func OnLoadUrlHeadersReceived(webView WebView, callback LoadUrlHeadersReceivedCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbLoadUrlHeadersReceivedCallback = func(webView WebView, param, url uintptr, job NetJob) (void uintptr) {
			callback(webView, StrFromPtr(url), job)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnLoadUrlHeadersReceived").Call(uintptr(webView), cbPtr, 0)
}

// ITERATOR3(void, mbOnDocumentReadyInBlinkThread, mbWebView webView, mbDocumentReadyCallback callback, void* param, "") \
func OnDocumentReadyInBlinkThread(webView WebView, callback DocumentReadyCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbDocumentReadyCallback = func(webView WebView, param uintptr, frameId WebFrameHandle) (void uintptr) {
			callback(webView, frameId)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnDocumentReadyInBlinkThread").Call(uintptr(webView), cbPtr, 0)
}

// ITERATOR2(void, mbUtilSetDefaultPrinterSettings, mbWebView webView, const mbDefaultPrinterSettings* setting, "") \
func UtilSetDefaultPrinterSettings(webView WebView, setting *DefaultPrinterSettings) {
	mbLib.MustFindProc("mbUtilSetDefaultPrinterSettings").Call(uintptr(webView), uintptr(unsafe.Pointer(setting)))
}

// int mbGetContentWidth(mbWebView webView);
func GetContentWidth(webView WebView) int {
	r1, _, _ := mbLib.MustFindProc("mbGetContentWidth").Call(uintptr(webView))
	return int(r1)
}

// int mbGetContentHeight(mbWebView webView);
func GetContentHeight(webView WebView) int {
	r1, _, _ := mbLib.MustFindProc("mbGetContentHeight").Call(uintptr(webView))
	return int(r1)
}

// mbWebView mbGetWebViewForCurrentContext();
func GetWebViewForCurrentContext() WebView {
	r1, _, _ := mbLib.MustFindProc("mbGetWebViewForCurrentContext").Call()
	return WebView(r1)
}

// BOOL mbRegisterEmbedderCustomElement(mbWebView webviewHandle, mbWebFrameHandle frameId, const char* name, void* options, void* outResult);
func RegisterEmbedderCustomElement(webviewHandle WebView, frameId WebFrameHandle, name string, options, outResult uintptr) bool {
	r1, _, _ := mbLib.MustFindProc("mbRegisterEmbedderCustomElement").Call(uintptr(webviewHandle), uintptr(frameId), StrToPtr(name), options, outResult)
	return BoolFromPtr(r1)
}

// void mbOnNodeCreateProcess(mbWebView webviewHandle, mbNodeOnCreateProcessCallback callback, void* param);
func OnNodeCreateProcess(webviewHandle WebView, callback NodeOnCreateProcessCallback) {
	var cbPtr uintptr = 0
	if callback != nil {

		var cb mbNodeOnCreateProcessCallback = func(webView WebView, param, applicationPath, arguments, startup uintptr) (void uintptr) {
			callback(webView, StrFromWPtr(applicationPath), StrFromWPtr(arguments), (*STARTUPINFOW)(unsafe.Pointer(startup)))
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnNodeCreateProcess").Call(uintptr(webviewHandle), cbPtr, 0)
}

// ITERATOR2(mbJsExecState, mbGetGlobalExecByFrame, mbWebView webView, mbWebFrameHandle frameId, "") \
func GetGlobalExecByFrame(webView WebView, frameId WebFrameHandle) JsExecState {
	r1, _, _ := mbLib.MustFindProc("mbGetGlobalExecByFrame").Call(uintptr(webView), uintptr(frameId))
	return JsExecState(r1)
}

// ITERATOR2(void*, mbJsToV8Value, mbJsExecState es, mbJsValue v, "") \
func JsToV8Value(es JsExecState, v JsValue) any {
	r1, _, _ := mbLib.MustFindProc("mbJsToV8Value").Call(uintptr(es), uintptr(v))
	return *(*any)(unsafe.Pointer(r1))
}

// void mbOnThreadIdle(mbThreadCallback callback, void* param1, void* param2);
func OnThreadIdle(callback ThreadCallback, param1, param2 uintptr) {
	var cbPtr uintptr = 0
	if callback != nil {
		var cb mbThreadCallback = func(param1, param2 uintptr) (void uintptr) {
			callback()
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnThreadIdle").Call(cbPtr, 0, 0)
}

// void mbOnBlinkThreadInit(mbThreadCallback callback, void* param1, void* param2);
func OnBlinkThreadInit(callback ThreadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {
		var cb mbThreadCallback = func(param1, param2 uintptr) (void uintptr) {
			callback()
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbOnBlinkThreadInit").Call(cbPtr, 0, 0)
}

// ITERATOR3(void, mbCallBlinkThreadAsync, mbThreadCallback callback, void* param1, void* param2, "") \
func CallBlinkThreadAsync(callback ThreadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {
		var cb mbThreadCallback = func(param1, param2 uintptr) (void uintptr) {
			callback()
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbCallBlinkThreadAsync").Call(cbPtr, 0, 0)
}

// ITERATOR3(void, mbCallBlinkThreadSync, mbThreadCallback callback, void* param1, void* param2, "") \
func CallBlinkThreadSync(callback ThreadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {
		var cb mbThreadCallback = func(param1, param2 uintptr) (void uintptr) {
			callback()
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbCallBlinkThreadSync").Call(cbPtr, 0, 0)
}

// ITERATOR3(void, mbCallUiThreadSync, mbThreadCallback callback, void* param1, void* param2, "") \
func CallUiThreadSync(callback ThreadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {
		var cb mbThreadCallback = func(param1, param2 uintptr) (void uintptr) {
			callback()
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbCallUiThreadSync").Call(cbPtr, 0, 0)
}

// ITERATOR3(void, mbCallUiThreadAsync, mbThreadCallback callback, void* param1, void* param2, "") \
func CallUiThreadAsync(callback ThreadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {
		var cb mbThreadCallback = func(param1, param2 uintptr) (void uintptr) {
			callback()
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbCallUiThreadAsync").Call(cbPtr, 0, 0)
}

// ITERATOR3(void, mbSetUserKeyValue, mbWebView webView, const char* key, void* value, "") \
func SetUserKeyValue(webView WebView, key string, value any) {
	mbLib.MustFindProc("mbSetUserKeyValue").Call(uintptr(webView), StrToPtr(key), uintptr(unsafe.Pointer(&value)), 0)
}

// ITERATOR2(void*, mbGetUserKeyValue, mbWebView webView, const char* key, "") \
func GetUserKeyValue(webView WebView, key string) (r1 uintptr, r2 uintptr, err syscall.Errno) {
	return mbLib.MustFindProc("mbGetUserKeyValue").Call(uintptr(webView), StrToPtr(key))
}

// ITERATOR2(void, mbGoToOffset, mbWebView webView, int offset, "") \
func GoToOffset(webView WebView, offset int) {
	mbLib.MustFindProc("mbGoToOffset").Call(uintptr(webView), uintptr(offset))
}

// ITERATOR2(void, mbGoToIndex, mbWebView webView, int index, "") \
func GoToIndex(webView WebView, index int) {
	mbLib.MustFindProc("mbGoToIndex").Call(uintptr(webView), uintptr(index))
}

// ITERATOR1(void, mbEditorRedo, mbWebView webView, "") \
func EditorRedo(webView WebView) {
	mbLib.MustFindProc("mbEditorRedo").Call(uintptr(webView))
}

// ITERATOR1(void, mbEditorUnSelect, mbWebView webView, "") \
func EditorUnSelect(webView WebView) {
	mbLib.MustFindProc("mbEditorUnSelect").Call(uintptr(webView))
}

// v8Isolate mbGetBlinkMainThreadIsolate();
func GetBlinkMainThreadIsolate() V8Isolate {
	r1, _, _ := mbLib.MustFindProc("mbGetBlinkMainThreadIsolate").Call()
	return V8Isolate(r1)
}

// ITERATOR3(void, mbInsertCSSByFrame, mbWebView webView, mbWebFrameHandle frameId, const utf8* cssText, "") \
func InsertCSSByFrame(webView WebView, frameId WebFrameHandle, cssText string) {
	mbLib.MustFindProc("mbInsertCSSByFrame").Call(uintptr(webView), uintptr(frameId), StrToPtr(cssText))
}

// void mbWebFrameGetMainWorldScriptContext(mbWebView webView, mbWebFrameHandle frameId, v8ContextPtr contextOut);
func WebFrameGetMainWorldScriptContext(webView WebView, frameId WebFrameHandle, contextOut *V8ContextPtr) {
	mbLib.MustFindProc("mbWebFrameGetMainWorldScriptContext").Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(contextOut)))
}

// ITERATOR1(const char*, mbNetGetReferrer, mbNetJob jobPtr, "获取request的referrer") \
func NetGetReferrer(jobPtr NetJob) string {
	r1, _, _ := mbLib.MustFindProc("mbNetGetReferrer").Call(uintptr(jobPtr))
	return StrFromPtr(r1)
}

// ITERATOR2(void, mbPostToUiThread, mbOnCallUiThread callback, void* param, "") \
func PostToUiThread(callback OnCallUiThreadCallback) {
	var cbPtr uintptr = 0
	if callback != nil {
		var cb mbOnCallUiThread = func(webView WebView, param uintptr) (void uintptr) {
			callback(webView)
			return
		}
		cbPtr = CallbackToPtr(cb)
	}
	mbLib.MustFindProc("mbPostToUiThread").Call(cbPtr, 0)
}

// ITERATOR2(void, mbSetEditable, mbWebView webView, bool editable, "") \
func SetEditable(webView WebView, editable bool) {
	mbLib.MustFindProc("mbSetEditable").Call(uintptr(webView), BoolToPtr(editable))
}

// ITERATOR2(int, mbQueryState, mbWebView webviewHandle, const char* type, "") \
func QueryState(webviewHandle WebView, queryType string) int {
	r1, _, _ := mbLib.MustFindProc("mbQueryState").Call(uintptr(webviewHandle), StrToPtr(queryType))
	return int(r1)
}
