package mb

import (
	"github.com/epkgs/miniblink/internal/lib"
)

func LoadLibrary(libfile string) error {
	return lib.LoadLibrary(libfile)
}

func LoadEmbedLibrary() error {
	return lib.LoadEmbedLibrary()
}

func Release() error {
	return lib.Release()
}

var _mbInit = lib.LazyFunc[mbInit]{Name: "mbInit"}
var _mbUninit = lib.LazyFunc[mbUninit]{Name: "mbUninit"}
var _mbCreateInitSettings = lib.LazyFunc[mbCreateInitSettings]{Name: "mbCreateInitSettings"}
var _mbSetInitSettings = lib.LazyFunc[mbSetInitSettings]{Name: "mbSetInitSettings"}
var _mbCreateWebView = lib.LazyFunc[mbCreateWebView]{Name: "mbCreateWebView"}
var _mbDestroyWebView = lib.LazyFunc[mbDestroyWebView]{Name: "mbDestroyWebView"}
var _mbCreateWebWindow = lib.LazyFunc[mbCreateWebWindow]{Name: "mbCreateWebWindow"}
var _mbCreateWebCustomWindow = lib.LazyFunc[mbCreateWebCustomWindow]{Name: "mbCreateWebCustomWindow"}
var _mbMoveWindow = lib.LazyFunc[mbMoveWindow]{Name: "mbMoveWindow"}
var _mbMoveToCenter = lib.LazyFunc[mbMoveToCenter]{Name: "mbMoveToCenter"}
var _mbSetAutoDrawToHwnd = lib.LazyFunc[mbSetAutoDrawToHwnd]{Name: "mbSetAutoDrawToHwnd"}
var _mbGetCaretRect = lib.LazyFunc[mbGetCaretRect]{Name: "mbGetCaretRect"}
var _mbSetAudioMuted = lib.LazyFunc[mbSetAudioMuted]{Name: "mbSetAudioMuted"}
var _mbIsAudioMuted = lib.LazyFunc[mbIsAudioMuted]{Name: "mbIsAudioMuted"}
var _mbCreateString = lib.LazyFunc[mbCreateString]{Name: "mbCreateString"}
var _mbCreateStringWithoutNullTermination = lib.LazyFunc[mbCreateStringWithoutNullTermination]{Name: "mbCreateStringWithoutNullTermination"}
var _mbDeleteString = lib.LazyFunc[mbDeleteString]{Name: "mbDeleteString"}
var _mbGetStringLen = lib.LazyFunc[mbGetStringLen]{Name: "mbGetStringLen"}
var _mbGetString = lib.LazyFunc[mbGetString]{Name: "mbGetString"}
var _mbSetProxy = lib.LazyFunc[mbSetProxy]{Name: "mbSetProxy"}
var _mbSetDebugConfig = lib.LazyFunc[mbSetDebugConfig]{Name: "mbSetDebugConfig"}
var _mbNetSetData = lib.LazyFunc[mbNetSetData]{Name: "mbNetSetData"}
var _mbNetHookRequest = lib.LazyFunc[mbNetHookRequest]{Name: "mbNetHookRequest"}
var _mbNetChangeRequestUrl = lib.LazyFunc[mbNetChangeRequestUrl]{Name: "mbNetChangeRequestUrl"}
var _mbNetContinueJob = lib.LazyFunc[mbNetContinueJob]{Name: "mbNetContinueJob"}
var _mbNetGetRawHttpHeadInBlinkThread = lib.LazyFunc[mbNetGetRawHttpHeadInBlinkThread]{Name: "mbNetGetRawHttpHeadInBlinkThread"}
var _mbNetGetRawResponseHeadInBlinkThread = lib.LazyFunc[mbNetGetRawResponseHeadInBlinkThread]{Name: "mbNetGetRawResponseHeadInBlinkThread"}
var _mbNetHoldJobToAsynCommit = lib.LazyFunc[mbNetHoldJobToAsynCommit]{Name: "mbNetHoldJobToAsynCommit"}
var _mbNetCancelRequest = lib.LazyFunc[mbNetCancelRequest]{Name: "mbNetCancelRequest"}
var _mbNetOnResponse = lib.LazyFunc[mbNetOnResponse]{Name: "mbNetOnResponse"}
var _mbNetSetWebsocketCallback = lib.LazyFunc[mbNetSetWebsocketCallback]{Name: "mbNetSetWebsocketCallback"}
var _mbNetSendWsText = lib.LazyFunc[mbNetSendWsText]{Name: "mbNetSendWsText"}
var _mbNetSendWsBlob = lib.LazyFunc[mbNetSendWsBlob]{Name: "mbNetSendWsBlob"}
var _mbNetEnableResPacket = lib.LazyFunc[mbNetEnableResPacket]{Name: "mbNetEnableResPacket"}
var _mbNetGetPostBody = lib.LazyFunc[mbNetGetPostBody]{Name: "mbNetGetPostBody"}
var _mbNetCreatePostBodyElements = lib.LazyFunc[mbNetCreatePostBodyElements]{Name: "mbNetCreatePostBodyElements"}
var _mbNetFreePostBodyElements = lib.LazyFunc[mbNetFreePostBodyElements]{Name: "mbNetFreePostBodyElements"}
var _mbNetCreatePostBodyElement = lib.LazyFunc[mbNetCreatePostBodyElement]{Name: "mbNetCreatePostBodyElement"}
var _mbNetFreePostBodyElement = lib.LazyFunc[mbNetFreePostBodyElement]{Name: "mbNetFreePostBodyElement"}
var _mbNetCreateWebUrlRequest = lib.LazyFunc[mbNetCreateWebUrlRequest]{Name: "mbNetCreateWebUrlRequest"}
var _mbNetAddHTTPHeaderFieldToUrlRequest = lib.LazyFunc[mbNetAddHTTPHeaderFieldToUrlRequest]{Name: "mbNetAddHTTPHeaderFieldToUrlRequest"}
var _mbNetStartUrlRequest = lib.LazyFunc[mbNetStartUrlRequest]{Name: "mbNetStartUrlRequest"}
var _mbNetGetHttpStatusCode = lib.LazyFunc[mbNetGetHttpStatusCode]{Name: "mbNetGetHttpStatusCode"}
var _mbNetGetRequestMethod = lib.LazyFunc[mbNetGetRequestMethod]{Name: "mbNetGetRequestMethod"}
var _mbNetGetExpectedContentLength = lib.LazyFunc[mbNetGetExpectedContentLength]{Name: "mbNetGetExpectedContentLength"}
var _mbNetGetResponseUrl = lib.LazyFunc[mbNetGetResponseUrl]{Name: "mbNetGetResponseUrl"}
var _mbNetCancelWebUrlRequest = lib.LazyFunc[mbNetCancelWebUrlRequest]{Name: "mbNetCancelWebUrlRequest"}
var _mbSetViewProxy = lib.LazyFunc[mbSetViewProxy]{Name: "mbSetViewProxy"}
var _mbNetSetMIMEType = lib.LazyFunc[mbNetSetMIMEType]{Name: "mbNetSetMIMEType"}
var _mbNetGetMIMEType = lib.LazyFunc[mbNetGetMIMEType]{Name: "mbNetGetMIMEType"}
var _mbNetGetHTTPHeaderField = lib.LazyFunc[mbNetGetHTTPHeaderField]{Name: "mbNetGetHTTPHeaderField"}
var _mbNetGetReferrer = lib.LazyFunc[mbNetGetReferrer]{Name: "mbNetGetReferrer"}
var _mbNetSetHTTPHeaderField = lib.LazyFunc[mbNetSetHTTPHeaderField]{Name: "mbNetSetHTTPHeaderField"}
var _mbSetMouseEnabled = lib.LazyFunc[mbSetMouseEnabled]{Name: "mbSetMouseEnabled"}
var _mbSetTouchEnabled = lib.LazyFunc[mbSetTouchEnabled]{Name: "mbSetTouchEnabled"}
var _mbSetSystemTouchEnabled = lib.LazyFunc[mbSetSystemTouchEnabled]{Name: "mbSetSystemTouchEnabled"}
var _mbSetContextMenuEnabled = lib.LazyFunc[mbSetContextMenuEnabled]{Name: "mbSetContextMenuEnabled"}
var _mbSetNavigationToNewWindowEnabled = lib.LazyFunc[mbSetNavigationToNewWindowEnabled]{Name: "mbSetNavigationToNewWindowEnabled"}
var _mbSetHeadlessEnabled = lib.LazyFunc[mbSetHeadlessEnabled]{Name: "mbSetHeadlessEnabled"}
var _mbSetDragDropEnable = lib.LazyFunc[mbSetDragDropEnable]{Name: "mbSetDragDropEnable"}
var _mbSetDragEnable = lib.LazyFunc[mbSetDragEnable]{Name: "mbSetDragEnable"}
var _mbSetContextMenuItemShow = lib.LazyFunc[mbSetContextMenuItemShow]{Name: "mbSetContextMenuItemShow"}
var _mbSetHandle = lib.LazyFunc[mbSetHandle]{Name: "mbSetHandle"}
var _mbSetHandleOffset = lib.LazyFunc[mbSetHandleOffset]{Name: "mbSetHandleOffset"}
var _mbGetHostHWND = lib.LazyFunc[mbGetHostHWND]{Name: "mbGetHostHWND"}
var _mbSetTransparent = lib.LazyFunc[mbSetTransparent]{Name: "mbSetTransparent"}
var _mbSetViewSettings = lib.LazyFunc[mbSetViewSettings]{Name: "mbSetViewSettings"}
var _mbSetCspCheckEnable = lib.LazyFunc[mbSetCspCheckEnable]{Name: "mbSetCspCheckEnable"}
var _mbSetNpapiPluginsEnabled = lib.LazyFunc[mbSetNpapiPluginsEnabled]{Name: "mbSetNpapiPluginsEnabled"}
var _mbSetMemoryCacheEnable = lib.LazyFunc[mbSetMemoryCacheEnable]{Name: "mbSetMemoryCacheEnable"}
var _mbSetCookie = lib.LazyFunc[mbSetCookie]{Name: "mbSetCookie"}
var _mbSetCookieEnabled = lib.LazyFunc[mbSetCookieEnabled]{Name: "mbSetCookieEnabled"}
var _mbSetCookieJarPath = lib.LazyFunc[mbSetCookieJarPath]{Name: "mbSetCookieJarPath"}
var _mbSetCookieJarFullPath = lib.LazyFunc[mbSetCookieJarFullPath]{Name: "mbSetCookieJarFullPath"}
var _mbSetLocalStorageFullPath = lib.LazyFunc[mbSetLocalStorageFullPath]{Name: "mbSetLocalStorageFullPath"}
var _mbGetTitle = lib.LazyFunc[mbGetTitle]{Name: "mbGetTitle"}
var _mbSetWindowTitle = lib.LazyFunc[mbSetWindowTitle]{Name: "mbSetWindowTitle"}
var _mbSetWindowTitleW = lib.LazyFunc[mbSetWindowTitleW]{Name: "mbSetWindowTitleW"}
var _mbGetUrl = lib.LazyFunc[mbGetUrl]{Name: "mbGetUrl"}
var _mbGetCursorInfoType = lib.LazyFunc[mbGetCursorInfoType]{Name: "mbGetCursorInfoType"}
var _mbAddPluginDirectory = lib.LazyFunc[mbAddPluginDirectory]{Name: "mbAddPluginDirectory"}
var _mbSetUserAgent = lib.LazyFunc[mbSetUserAgent]{Name: "mbSetUserAgent"}
var _mbSetZoomFactor = lib.LazyFunc[mbSetZoomFactor]{Name: "mbSetZoomFactor"}
var _mbGetZoomFactor = lib.LazyFunc[mbGetZoomFactor]{Name: "mbGetZoomFactor"}
var _mbSetDiskCacheEnabled = lib.LazyFunc[mbSetDiskCacheEnabled]{Name: "mbSetDiskCacheEnabled"}
var _mbSetDiskCachePath = lib.LazyFunc[mbSetDiskCachePath]{Name: "mbSetDiskCachePath"}
var _mbSetDiskCacheLimit = lib.LazyFunc[mbSetDiskCacheLimit]{Name: "mbSetDiskCacheLimit"}
var _mbSetDiskCacheLimitDisk = lib.LazyFunc[mbSetDiskCacheLimitDisk]{Name: "mbSetDiskCacheLimitDisk"}
var _mbSetDiskCacheLevel = lib.LazyFunc[mbSetDiskCacheLevel]{Name: "mbSetDiskCacheLevel"}
var _mbSetResourceGc = lib.LazyFunc[mbSetResourceGc]{Name: "mbSetResourceGc"}
var _mbCanGoForward = lib.LazyFunc[mbCanGoForward]{Name: "mbCanGoForward"}
var _mbCanGoBack = lib.LazyFunc[mbCanGoBack]{Name: "mbCanGoBack"}
var _mbGetCookie = lib.LazyFunc[mbGetCookie]{Name: "mbGetCookie"}
var _mbGetCookieOnBlinkThread = lib.LazyFunc[mbGetCookieOnBlinkThread]{Name: "mbGetCookieOnBlinkThread"}
var _mbClearCookie = lib.LazyFunc[mbClearCookie]{Name: "mbClearCookie"}
var _mbResize = lib.LazyFunc[mbResize]{Name: "mbResize"}
var _mbOnNavigation = lib.LazyFunc[mbOnNavigation]{Name: "mbOnNavigation"}
var _mbOnNavigationSync = lib.LazyFunc[mbOnNavigationSync]{Name: "mbOnNavigationSync"}
var _mbOnCreateView = lib.LazyFunc[mbOnCreateView]{Name: "mbOnCreateView"}
var _mbOnDocumentReady = lib.LazyFunc[mbOnDocumentReady]{Name: "mbOnDocumentReady"}
var _mbOnDocumentReadyInBlinkThread = lib.LazyFunc[mbOnDocumentReadyInBlinkThread]{Name: "mbOnDocumentReadyInBlinkThread"}
var _mbOnPaintUpdated = lib.LazyFunc[mbOnPaintUpdated]{Name: "mbOnPaintUpdated"}
var _mbOnPaintBitUpdated = lib.LazyFunc[mbOnPaintBitUpdated]{Name: "mbOnPaintBitUpdated"}
var _mbOnAcceleratedPaint = lib.LazyFunc[mbOnAcceleratedPaint]{Name: "mbOnAcceleratedPaint"}
var _mbOnLoadUrlBegin = lib.LazyFunc[mbOnLoadUrlBegin]{Name: "mbOnLoadUrlBegin"}
var _mbOnLoadUrlEnd = lib.LazyFunc[mbOnLoadUrlEnd]{Name: "mbOnLoadUrlEnd"}
var _mbOnLoadUrlFail = lib.LazyFunc[mbOnLoadUrlFail]{Name: "mbOnLoadUrlFail"}
var _mbOnLoadUrlHeadersReceived = lib.LazyFunc[mbOnLoadUrlHeadersReceived]{Name: "mbOnLoadUrlHeadersReceived"}
var _mbOnLoadUrlFinish = lib.LazyFunc[mbOnLoadUrlFinish]{Name: "mbOnLoadUrlFinish"}
var _mbOnTitleChanged = lib.LazyFunc[mbOnTitleChanged]{Name: "mbOnTitleChanged"}
var _mbOnURLChanged = lib.LazyFunc[mbOnURLChanged]{Name: "mbOnURLChanged"}
var _mbOnLoadingFinish = lib.LazyFunc[mbOnLoadingFinish]{Name: "mbOnLoadingFinish"}
var _mbOnDownload = lib.LazyFunc[mbOnDownload]{Name: "mbOnDownload"}
var _mbOnDownloadInBlinkThread = lib.LazyFunc[mbOnDownloadInBlinkThread]{Name: "mbOnDownloadInBlinkThread"}
var _mbOnAlertBox = lib.LazyFunc[mbOnAlertBox]{Name: "mbOnAlertBox"}
var _mbOnConfirmBox = lib.LazyFunc[mbOnConfirmBox]{Name: "mbOnConfirmBox"}
var _mbOnPromptBox = lib.LazyFunc[mbOnPromptBox]{Name: "mbOnPromptBox"}
var _mbOnNetGetFavicon = lib.LazyFunc[mbOnNetGetFavicon]{Name: "mbOnNetGetFavicon"}
var _mbOnConsole = lib.LazyFunc[mbOnConsole]{Name: "mbOnConsole"}
var _mbOnClose = lib.LazyFunc[mbOnClose]{Name: "mbOnClose"}
var _mbOnDestroy = lib.LazyFunc[mbOnDestroy]{Name: "mbOnDestroy"}
var _mbOnPrinting = lib.LazyFunc[mbOnPrinting]{Name: "mbOnPrinting"}
var _mbOnPluginList = lib.LazyFunc[mbOnPluginList]{Name: "mbOnPluginList"}
var _mbOnImageBufferToDataURL = lib.LazyFunc[mbOnImageBufferToDataURL]{Name: "mbOnImageBufferToDataURL"}
var _mbOnDidCreateScriptContext = lib.LazyFunc[mbOnDidCreateScriptContext]{Name: "mbOnDidCreateScriptContext"}
var _mbOnWillReleaseScriptContext = lib.LazyFunc[mbOnWillReleaseScriptContext]{Name: "mbOnWillReleaseScriptContext"}
var _mbGoBack = lib.LazyFunc[mbGoBack]{Name: "mbGoBack"}
var _mbGoForward = lib.LazyFunc[mbGoForward]{Name: "mbGoForward"}
var _mbGoToOffset = lib.LazyFunc[mbGoToOffset]{Name: "mbGoToOffset"}
var _mbGoToIndex = lib.LazyFunc[mbGoToIndex]{Name: "mbGoToIndex"}
var _mbNavigateAtIndex = lib.LazyFunc[mbNavigateAtIndex]{Name: "mbNavigateAtIndex"}
var _mbGetNavigateIndex = lib.LazyFunc[mbGetNavigateIndex]{Name: "mbGetNavigateIndex"}
var _mbStopLoading = lib.LazyFunc[mbStopLoading]{Name: "mbStopLoading"}
var _mbReload = lib.LazyFunc[mbReload]{Name: "mbReload"}
var _mbPerformCookieCommand = lib.LazyFunc[mbPerformCookieCommand]{Name: "mbPerformCookieCommand"}
var _mbInsertCSSByFrame = lib.LazyFunc[mbInsertCSSByFrame]{Name: "mbInsertCSSByFrame"}
var _mbEditorSelectAll = lib.LazyFunc[mbEditorSelectAll]{Name: "mbEditorSelectAll"}
var _mbEditorUnSelect = lib.LazyFunc[mbEditorUnSelect]{Name: "mbEditorUnSelect"}
var _mbEditorCopy = lib.LazyFunc[mbEditorCopy]{Name: "mbEditorCopy"}
var _mbEditorCut = lib.LazyFunc[mbEditorCut]{Name: "mbEditorCut"}
var _mbEditorPaste = lib.LazyFunc[mbEditorPaste]{Name: "mbEditorPaste"}
var _mbEditorDelete = lib.LazyFunc[mbEditorDelete]{Name: "mbEditorDelete"}
var _mbEditorUndo = lib.LazyFunc[mbEditorUndo]{Name: "mbEditorUndo"}
var _mbEditorRedo = lib.LazyFunc[mbEditorRedo]{Name: "mbEditorRedo"}
var _mbSetEditable = lib.LazyFunc[mbSetEditable]{Name: "mbSetEditable"}
var _mbFireMouseEvent = lib.LazyFunc[mbFireMouseEvent]{Name: "mbFireMouseEvent"}
var _mbFireContextMenuEvent = lib.LazyFunc[mbFireContextMenuEvent]{Name: "mbFireContextMenuEvent"}
var _mbFireMouseWheelEvent = lib.LazyFunc[mbFireMouseWheelEvent]{Name: "mbFireMouseWheelEvent"}
var _mbFireKeyUpEvent = lib.LazyFunc[mbFireKeyUpEvent]{Name: "mbFireKeyUpEvent"}
var _mbFireKeyDownEvent = lib.LazyFunc[mbFireKeyDownEvent]{Name: "mbFireKeyDownEvent"}
var _mbFireKeyPressEvent = lib.LazyFunc[mbFireKeyPressEvent]{Name: "mbFireKeyPressEvent"}
var _mbFireWindowsMessage = lib.LazyFunc[mbFireWindowsMessage]{Name: "mbFireWindowsMessage"}
var _mbSetFocus = lib.LazyFunc[mbSetFocus]{Name: "mbSetFocus"}
var _mbKillFocus = lib.LazyFunc[mbKillFocus]{Name: "mbKillFocus"}
var _mbShowWindow = lib.LazyFunc[mbShowWindow]{Name: "mbShowWindow"}
var _mbLoadURL = lib.LazyFunc[mbLoadURL]{Name: "mbLoadURL"}
var _mbLoadHtmlWithBaseUrl = lib.LazyFunc[mbLoadHtmlWithBaseUrl]{Name: "mbLoadHtmlWithBaseUrl"}
var _mbPostURL = lib.LazyFunc[mbPostURL]{Name: "mbPostURL"}
var _mbGetLockedViewDC = lib.LazyFunc[mbGetLockedViewDC]{Name: "mbGetLockedViewDC"}
var _mbUnlockViewDC = lib.LazyFunc[mbUnlockViewDC]{Name: "mbUnlockViewDC"}
var _mbWake = lib.LazyFunc[mbWake]{Name: "mbWake"}
var _mbJsToV8Value = lib.LazyFunc[mbJsToV8Value]{Name: "mbJsToV8Value"}
var _mbGetGlobalExecByFrame = lib.LazyFunc[mbGetGlobalExecByFrame]{Name: "mbGetGlobalExecByFrame"}
var _mbJsToDouble = lib.LazyFunc[mbJsToDouble]{Name: "mbJsToDouble"}
var _mbJsToBoolean = lib.LazyFunc[mbJsToBoolean]{Name: "mbJsToBoolean"}
var _mbJsToString = lib.LazyFunc[mbJsToString]{Name: "mbJsToString"}
var _mbGetJsValueType = lib.LazyFunc[mbGetJsValueType]{Name: "mbGetJsValueType"}
var _mbOnJsQuery = lib.LazyFunc[mbOnJsQuery]{Name: "mbOnJsQuery"}
var _mbResponseQuery = lib.LazyFunc[mbResponseQuery]{Name: "mbResponseQuery"}
var _mbRunJs = lib.LazyFunc[mbRunJs]{Name: "mbRunJs"}
var _mbRunJsSync = lib.LazyFunc[mbRunJsSync]{Name: "mbRunJsSync"}
var _mbWebFrameGetMainFrame = lib.LazyFunc[mbWebFrameGetMainFrame]{Name: "mbWebFrameGetMainFrame"}
var _mbIsMainFrame = lib.LazyFunc[mbIsMainFrame]{Name: "mbIsMainFrame"}
var _mbSetNodeJsEnable = lib.LazyFunc[mbSetNodeJsEnable]{Name: "mbSetNodeJsEnable"}
var _mbSetDeviceParameter = lib.LazyFunc[mbSetDeviceParameter]{Name: "mbSetDeviceParameter"}
var _mbGetContentAsMarkup = lib.LazyFunc[mbGetContentAsMarkup]{Name: "mbGetContentAsMarkup"}
var _mbGetSource = lib.LazyFunc[mbGetSource]{Name: "mbGetSource"}
var _mbUtilSerializeToMHTML = lib.LazyFunc[mbUtilSerializeToMHTML]{Name: "mbUtilSerializeToMHTML"}
var _mbUtilCreateRequestCode = lib.LazyFunc[mbUtilCreateRequestCode]{Name: "mbUtilCreateRequestCode"}
var _mbUtilIsRegistered = lib.LazyFunc[mbUtilIsRegistered]{Name: "mbUtilIsRegistered"}
var _mbUtilPrint = lib.LazyFunc[mbUtilPrint]{Name: "mbUtilPrint"}
var _mbUtilBase64Encode = lib.LazyFunc[mbUtilBase64Encode]{Name: "mbUtilBase64Encode"}
var _mbUtilBase64Decode = lib.LazyFunc[mbUtilBase64Decode]{Name: "mbUtilBase64Decode"}
var _mbUtilDecodeURLEscape = lib.LazyFunc[mbUtilDecodeURLEscape]{Name: "mbUtilDecodeURLEscape"}
var _mbUtilEncodeURLEscape = lib.LazyFunc[mbUtilEncodeURLEscape]{Name: "mbUtilEncodeURLEscape"}
var _mbUtilCreateV8Snapshot = lib.LazyFunc[mbUtilCreateV8Snapshot]{Name: "mbUtilCreateV8Snapshot"}
var _mbUtilPrintToPdf = lib.LazyFunc[mbUtilPrintToPdf]{Name: "mbUtilPrintToPdf"}
var _mbUtilPrintToBitmap = lib.LazyFunc[mbUtilPrintToBitmap]{Name: "mbUtilPrintToBitmap"}
var _mbUtilScreenshot = lib.LazyFunc[mbUtilScreenshot]{Name: "mbUtilScreenshot"}
var _mbUtilsSilentPrint = lib.LazyFunc[mbUtilsSilentPrint]{Name: "mbUtilsSilentPrint"}
var _mbUtilSetDefaultPrinterSettings = lib.LazyFunc[mbUtilSetDefaultPrinterSettings]{Name: "mbUtilSetDefaultPrinterSettings"}
var _mbPopupDownloadMgr = lib.LazyFunc[mbPopupDownloadMgr]{Name: "mbPopupDownloadMgr"}
var _mbPopupDialogAndDownload = lib.LazyFunc[mbPopupDialogAndDownload]{Name: "mbPopupDialogAndDownload"}
var _mbDownloadByPath = lib.LazyFunc[mbDownloadByPath]{Name: "mbDownloadByPath"}
var _mbGetPdfPageData = lib.LazyFunc[mbGetPdfPageData]{Name: "mbGetPdfPageData"}
var _mbCreateMemBuf = lib.LazyFunc[mbCreateMemBuf]{Name: "mbCreateMemBuf"}
var _mbFreeMemBuf = lib.LazyFunc[mbFreeMemBuf]{Name: "mbFreeMemBuf"}
var _mbSetUserKeyValue = lib.LazyFunc[mbSetUserKeyValue]{Name: "mbSetUserKeyValue"}
var _mbGetUserKeyValue = lib.LazyFunc[mbGetUserKeyValue]{Name: "mbGetUserKeyValue"}
var _mbPluginListBuilderAddPlugin = lib.LazyFunc[mbPluginListBuilderAddPlugin]{Name: "mbPluginListBuilderAddPlugin"}
var _mbPluginListBuilderAddMediaTypeToLastPlugin = lib.LazyFunc[mbPluginListBuilderAddMediaTypeToLastPlugin]{Name: "mbPluginListBuilderAddMediaTypeToLastPlugin"}
var _mbPluginListBuilderAddFileExtensionToLastMediaType = lib.LazyFunc[mbPluginListBuilderAddFileExtensionToLastMediaType]{Name: "mbPluginListBuilderAddFileExtensionToLastMediaType"}
var _mbGetBlinkMainThreadIsolate = lib.LazyFunc[mbGetBlinkMainThreadIsolate]{Name: "mbGetBlinkMainThreadIsolate"}
var _mbWebFrameGetMainWorldScriptContext = lib.LazyFunc[mbWebFrameGetMainWorldScriptContext]{Name: "mbWebFrameGetMainWorldScriptContext"}
var _mbEnableHighDPISupport = lib.LazyFunc[mbEnableHighDPISupport]{Name: "mbEnableHighDPISupport"}
var _mbRunMessageLoop = lib.LazyFunc[mbRunMessageLoop]{Name: "mbRunMessageLoop"}
var _mbGetContentWidth = lib.LazyFunc[mbGetContentWidth]{Name: "mbGetContentWidth"}
var _mbGetContentHeight = lib.LazyFunc[mbGetContentHeight]{Name: "mbGetContentHeight"}
var _mbGetWebViewForCurrentContext = lib.LazyFunc[mbGetWebViewForCurrentContext]{Name: "mbGetWebViewForCurrentContext"}
var _mbRegisterEmbedderCustomElement = lib.LazyFunc[mbRegisterEmbedderCustomElement]{Name: "mbRegisterEmbedderCustomElement"}
var _mbOnNodeCreateProcess = lib.LazyFunc[mbOnNodeCreateProcess]{Name: "mbOnNodeCreateProcess"}
var _mbOnThreadIdle = lib.LazyFunc[mbOnThreadIdle]{Name: "mbOnThreadIdle"}
var _mbGetProcAddr = lib.LazyFunc[mbGetProcAddr]{Name: "mbGetProcAddr"}
var _mbSetMbDllPath = lib.LazyFunc[mbSetMbDllPath]{Name: "mbSetMbDllPath"}
var _mbSetMbMainDllPath = lib.LazyFunc[mbSetMbMainDllPath]{Name: "mbSetMbMainDllPath"}
var _mbFillFuncPtr = lib.LazyFunc[mbFillFuncPtr]{Name: "mbFillFuncPtr"}

func Init(settings *Settings) {
	_mbInit.LoadOnce()
	_mbInit.Call(settings)
}

func Uinit() {
	_mbUninit.LoadOnce()
	_mbUninit.Call()
}

func CreateInitSettings() *Settings {
	_mbCreateInitSettings.LoadOnce()
	return _mbCreateInitSettings.Call()
}

func SetInitSettings(settings *Settings, name, value string) {
	_mbSetInitSettings.LoadOnce()
	_mbSetInitSettings.Call(settings, name, value)
}

func CreateWebView() WebView {
	_mbCreateWebView.LoadOnce()
	return _mbCreateWebView.Call()
}
func DestroyWebView(webview WebView) {
	_mbDestroyWebView.LoadOnce()
	_mbDestroyWebView.Call(webview)
}

func CreateWebWindow(typ WindowType, parent Handle, x, y, width, height int) WebView {
	_mbCreateWebWindow.LoadOnce()
	return _mbCreateWebWindow.Call(typ, parent, x, y, width, height)
}

func CreateWebCustomWindow(parent Handle, style, styleEx uint32, x, y, width, height int) WebView {
	_mbCreateWebCustomWindow.LoadOnce()
	return _mbCreateWebCustomWindow.Call(parent, style, styleEx, x, y, width, height)
}

func MoveWindow(webview WebView, x, y, width, height int) {
	_mbMoveWindow.LoadOnce()
	_mbMoveWindow.Call(webview, x, y, width, height)
}

func MoveToCenter(webview WebView) {
	_mbMoveToCenter.LoadOnce()
	_mbMoveToCenter.Call(webview)
}

func SetAutoDrawToHwnd(webview WebView, b bool) {
	_mbSetAutoDrawToHwnd.LoadOnce()
	_mbSetAutoDrawToHwnd.Call(webview, b)
}

func GetCaretRect(webview WebView, r *Rect) {
	_mbGetCaretRect.LoadOnce()
	_mbGetCaretRect.Call(webview, r)
}

func SetAudioMuted(webview WebView, b bool) {
	_mbSetAudioMuted.LoadOnce()
	_mbSetAudioMuted.Call(webview, b)
}

func IsAudioMuted(webview WebView) bool {
	_mbIsAudioMuted.LoadOnce()
	return _mbIsAudioMuted.Call(webview)
}

func CreateString(str string, length uintptr) StringPtr {
	_mbCreateString.LoadOnce()
	return _mbCreateString.Call(str, length)
}

func CreateStringWithoutNullTermination(str string, length uintptr) StringPtr {
	_mbCreateStringWithoutNullTermination.LoadOnce()
	return _mbCreateStringWithoutNullTermination.Call(str, length)
}

func DeleteString(str StringPtr) {
	_mbDeleteString.LoadOnce()
	_mbDeleteString.Call(str)
}

func GetStringLen(str StringPtr) uintptr {
	_mbGetStringLen.LoadOnce()
	return _mbGetStringLen.Call(str)
}

func GetString(str StringPtr) string {
	_mbGetString.LoadOnce()
	return _mbGetString.Call(str)
}
func SetProxy(webView WebView, proxy *Proxy) {
	_mbSetProxy.LoadOnce()
	_mbSetProxy.Call(webView, proxy)
}

func SetDebugConfig(webView WebView, debugString, param string) {
	_mbSetDebugConfig.LoadOnce()
	_mbSetDebugConfig.Call(webView, debugString, param)
}

func NetSetData(jobPtr NetJob, buf []byte) {
	_mbNetSetData.LoadOnce()
	_mbNetSetData.Call(jobPtr, &buf[0], len(buf))
}

func NetHookRequest(jobPtr NetJob) {
	_mbNetHookRequest.LoadOnce()
	_mbNetHookRequest.Call(jobPtr)
}

func NetChangeRequestUrl(jobPtr NetJob, url string) {
	_mbNetChangeRequestUrl.LoadOnce()
	_mbNetChangeRequestUrl.Call(jobPtr, url)
}

func NetContinueJob(jobPtr NetJob) {
	_mbNetContinueJob.LoadOnce()
	_mbNetContinueJob.Call(jobPtr)
}

func NetGetRawHttpHeadInBlinkThread(jobPtr NetJob) *Slist {
	_mbNetGetRawHttpHeadInBlinkThread.LoadOnce()
	return _mbNetGetRawHttpHeadInBlinkThread.Call(jobPtr)
}

func NetGetRawResponseHeadInBlinkThread(jobPtr NetJob) *Slist {
	_mbNetGetRawResponseHeadInBlinkThread.LoadOnce()
	return _mbNetGetRawResponseHeadInBlinkThread.Call(jobPtr)
}

func NetHoldJobToAsynCommit(jobPtr NetJob) bool {
	_mbNetHoldJobToAsynCommit.LoadOnce()
	return _mbNetHoldJobToAsynCommit.Call(jobPtr)
}

func NetCancelRequest(jobPtr NetJob) {
	_mbNetCancelRequest.LoadOnce()
	_mbNetCancelRequest.Call(jobPtr)
}

func NetOnResponse(webviewHandle WebView, callback NetResponseCallback, param uintptr) {
	_mbNetOnResponse.LoadOnce()
	_mbNetOnResponse.Call(webviewHandle, callback, param)
}

func NetSetWebsocketCallback(webview WebView, callbacks *WebsocketHookCallbacks, param uintptr) {
	_mbNetSetWebsocketCallback.LoadOnce()
	_mbNetSetWebsocketCallback.Call(webview, callbacks, param)
}

func NetSendWsText(channel WebSocketChannel, buf *byte, len uintptr) {
	_mbNetSendWsText.LoadOnce()
	_mbNetSendWsText.Call(channel, buf, len)
}

func NetSendWsBlob(channel WebSocketChannel, buf *byte, len uintptr) {
	_mbNetSendWsBlob.LoadOnce()
	_mbNetSendWsBlob.Call(channel, buf, len)
}

func NetEnableResPacket(webviewHandle WebView, pathName *uint16) {
	_mbNetEnableResPacket.LoadOnce()
	_mbNetEnableResPacket.Call(webviewHandle, pathName)
}

func NetGetPostBody(jobPtr NetJob) *PostBodyElements {
	_mbNetGetPostBody.LoadOnce()
	return _mbNetGetPostBody.Call(jobPtr)
}

func NetCreatePostBodyElements(webView WebView, length uintptr) *PostBodyElements {
	_mbNetCreatePostBodyElements.LoadOnce()
	return _mbNetCreatePostBodyElements.Call(webView, length)
}

func NetFreePostBodyElements(elements *PostBodyElements) {
	_mbNetFreePostBodyElements.LoadOnce()
	_mbNetFreePostBodyElements.Call(elements)
}

func NetCreatePostBodyElement(webView WebView) *PostBodyElement {
	_mbNetCreatePostBodyElement.LoadOnce()
	return _mbNetCreatePostBodyElement.Call(webView)
}

func NetFreePostBodyElement(element *PostBodyElement) {
	_mbNetFreePostBodyElement.LoadOnce()
	_mbNetFreePostBodyElement.Call(element)
}

func NetCreateWebUrlRequest(url, method, mime string) WebUrlRequestPtr {
	_mbNetCreateWebUrlRequest.LoadOnce()
	return _mbNetCreateWebUrlRequest.Call(url, method, mime)
}

func NetAddHTTPHeaderFieldToUrlRequest(request WebUrlRequestPtr, name, value string) {
	_mbNetAddHTTPHeaderFieldToUrlRequest.LoadOnce()
	_mbNetAddHTTPHeaderFieldToUrlRequest.Call(request, name, value)
}

func NetStartUrlRequest(webView WebView, request WebUrlRequestPtr, param uintptr, callbacks *UrlRequestCallbacks) int {
	_mbNetStartUrlRequest.LoadOnce()
	return _mbNetStartUrlRequest.Call(webView, request, param, callbacks)
}

func NetGetHttpStatusCode(response WebUrlResponsePtr) int {
	_mbNetGetHttpStatusCode.LoadOnce()
	return _mbNetGetHttpStatusCode.Call(response)
}

func NetGetRequestMethod(jobPtr NetJob) RequestType {
	_mbNetGetRequestMethod.LoadOnce()
	return _mbNetGetRequestMethod.Call(jobPtr)
}

func NetGetExpectedContentLength(response WebUrlResponsePtr) int64 {
	_mbNetGetExpectedContentLength.LoadOnce()
	return _mbNetGetExpectedContentLength.Call(response)
}

func NetGetResponseUrl(response WebUrlResponsePtr) string {
	_mbNetGetResponseUrl.LoadOnce()
	return _mbNetGetResponseUrl.Call(response)
}

func NetCancelWebUrlRequest(requestId int) {
	_mbNetCancelWebUrlRequest.LoadOnce()
	_mbNetCancelWebUrlRequest.Call(requestId)
}

func SetViewProxy(webView WebView, proxy *Proxy) {
	_mbSetViewProxy.LoadOnce()
	_mbSetViewProxy.Call(webView, proxy)
}

func NetSetMIMEType(jobPtr NetJob, typeStr string) {
	_mbNetSetMIMEType.LoadOnce()
	_mbNetSetMIMEType.Call(jobPtr, typeStr)
}

func NetGetMIMEType(jobPtr NetJob) string {
	_mbNetGetMIMEType.LoadOnce()
	return _mbNetGetMIMEType.Call(jobPtr)
}

func NetGetHTTPHeaderField(job NetJob, key string, fromRequestOrResponse bool) string {
	_mbNetGetHTTPHeaderField.LoadOnce()
	return _mbNetGetHTTPHeaderField.Call(job, key, fromRequestOrResponse)
}

func NetGetReferrer(jobPtr NetJob) string {
	_mbNetGetReferrer.LoadOnce()
	return _mbNetGetReferrer.Call(jobPtr)
}
func NetSetHTTPHeaderField(jobPtr NetJob, key, value string, response bool) {
	_mbNetSetHTTPHeaderField.LoadOnce()
	_mbNetSetHTTPHeaderField.Call(jobPtr, key, value, response)
}

func SetMouseEnabled(webView WebView, b bool) {
	_mbSetMouseEnabled.LoadOnce()
	_mbSetMouseEnabled.Call(webView, b)
}

func SetTouchEnabled(webView WebView, b bool) {
	_mbSetTouchEnabled.LoadOnce()
	_mbSetTouchEnabled.Call(webView, b)
}
func SetSystemTouchEnabled(webView WebView, b bool) {
	_mbSetSystemTouchEnabled.LoadOnce()
	_mbSetSystemTouchEnabled.Call(webView, b)
}

func SetContextMenuEnabled(webView WebView, b bool) {
	_mbSetContextMenuEnabled.LoadOnce()
	_mbSetContextMenuEnabled.Call(webView, b)
}

func SetNavigationToNewWindowEnabled(webView WebView, b bool) {
	_mbSetNavigationToNewWindowEnabled.LoadOnce()
	_mbSetNavigationToNewWindowEnabled.Call(webView, b)
}

func SetHeadlessEnabled(webView WebView, b bool) {
	_mbSetHeadlessEnabled.LoadOnce()
	_mbSetHeadlessEnabled.Call(webView, b)
}

func SetDragDropEnable(webView WebView, b bool) {
	_mbSetDragDropEnable.LoadOnce()
	_mbSetDragDropEnable.Call(webView, b)
}

func SetDragEnable(webView WebView, b bool) {
	_mbSetDragEnable.LoadOnce()
	_mbSetDragEnable.Call(webView, b)
}

func SetContextMenuItemShow(webView WebView, item MenuItemId, isShow bool) {
	_mbSetContextMenuItemShow.LoadOnce()
	_mbSetContextMenuItemShow.Call(webView, item, isShow)
}

func SetHandle(webView WebView, wnd HWND) {
	_mbSetHandle.LoadOnce()
	_mbSetHandle.Call(webView, wnd)
}

func SetHandleOffset(webView WebView, x, y int) {
	_mbSetHandleOffset.LoadOnce()
	_mbSetHandleOffset.Call(webView, x, y)
}

func GetHostHWND(webView WebView) HWND {
	_mbGetHostHWND.LoadOnce()
	return _mbGetHostHWND.Call(webView)
}

func SetTransparent(webviewHandle WebView, transparent bool) {
	_mbSetTransparent.LoadOnce()
	_mbSetTransparent.Call(webviewHandle, transparent)
}

func SetViewSettings(webviewHandle WebView, settings *ViewSettings) {
	_mbSetViewSettings.LoadOnce()
	_mbSetViewSettings.Call(webviewHandle, settings)
}

func SetCspCheckEnable(webView WebView, b bool) {
	_mbSetCspCheckEnable.LoadOnce()
	_mbSetCspCheckEnable.Call(webView, b)
}

func SetNpapiPluginsEnabled(webView WebView, b bool) {
	_mbSetNpapiPluginsEnabled.LoadOnce()
	_mbSetNpapiPluginsEnabled.Call(webView, b)
}

func SetMemoryCacheEnable(webView WebView, b bool) {
	_mbSetMemoryCacheEnable.LoadOnce()
	_mbSetMemoryCacheEnable.Call(webView, b)
}

func SetCookie(webView WebView, url, cookie string) {
	_mbSetCookie.LoadOnce()
	_mbSetCookie.Call(webView, url, cookie)
}

func SetCookieEnabled(webView WebView, enable bool) {
	_mbSetCookieEnabled.LoadOnce()
	_mbSetCookieEnabled.Call(webView, enable)
}

func SetCookieJarPath(webView WebView, path string) {
	_mbSetCookieJarPath.LoadOnce()
	_mbSetCookieJarPath.Call(webView, path)
}

func SetCookieJarFullPath(webView WebView, path string) {
	_mbSetCookieJarFullPath.LoadOnce()
	_mbSetCookieJarFullPath.Call(webView, path)
}

func SetLocalStorageFullPath(webView WebView, path string) {
	_mbSetLocalStorageFullPath.LoadOnce()
	_mbSetLocalStorageFullPath.Call(webView, path)
}

func GetTitle(webView WebView) string {
	_mbGetTitle.LoadOnce()
	return _mbGetTitle.Call(webView)
}

func SetWindowTitle(webView WebView, title string) {
	_mbSetWindowTitle.LoadOnce()
	_mbSetWindowTitle.Call(webView, title)
}

func SetWindowTitleW(webView WebView, title string) {
	_mbSetWindowTitleW.LoadOnce()
	_mbSetWindowTitleW.Call(webView, title)
}

func GetUrl(webView WebView) string {
	_mbGetUrl.LoadOnce()
	return _mbGetUrl.Call(webView)
}

func GetCursorInfoType(webView WebView) int {
	_mbGetCursorInfoType.LoadOnce()
	return _mbGetCursorInfoType.Call(webView)
}

func AddPluginDirectory(webView WebView, path string) {
	_mbAddPluginDirectory.LoadOnce()
	_mbAddPluginDirectory.Call(webView, path)
}

func SetUserAgent(webView WebView, userAgent string) {
	_mbSetUserAgent.LoadOnce()
	_mbSetUserAgent.Call(webView, userAgent)
}

func SetZoomFactor(webView WebView, factor float32) {
	_mbSetZoomFactor.LoadOnce()
	_mbSetZoomFactor.Call(webView, factor)
}

func GetZoomFactor(webView WebView) float32 {
	_mbGetZoomFactor.LoadOnce()
	return _mbGetZoomFactor.Call(webView)
}
func SetDiskCacheEnabled(webView WebView, enable bool) {
	_mbSetDiskCacheEnabled.LoadOnce()
	_mbSetDiskCacheEnabled.Call(webView, enable)
}

func SetDiskCachePath(webView WebView, path string) {
	_mbSetDiskCachePath.LoadOnce()
	_mbSetDiskCachePath.Call(webView, path)
}

func SetDiskCacheLimit(webView WebView, limit uint64) {
	_mbSetDiskCacheLimit.LoadOnce()
	_mbSetDiskCacheLimit.Call(webView, limit)
}

func SetDiskCacheLimitDisk(webView WebView, limit uint64) {
	_mbSetDiskCacheLimitDisk.LoadOnce()
	_mbSetDiskCacheLimitDisk.Call(webView, limit)
}

func SetDiskCacheLevel(webView WebView, level int) {
	_mbSetDiskCacheLevel.LoadOnce()
	_mbSetDiskCacheLevel.Call(webView, level)
}

func SetResourceGc(webView WebView, intervalSec int) {
	_mbSetResourceGc.LoadOnce()
	_mbSetResourceGc.Call(webView, intervalSec)
}

func CanGoForward(webView WebView, callback CanGoBackForwardCallback, param uintptr) {
	_mbCanGoForward.LoadOnce()
	_mbCanGoForward.Call(webView, callback, param)
}

func CanGoBack(webView WebView, callback CanGoBackForwardCallback, param uintptr) {
	_mbCanGoBack.LoadOnce()
	_mbCanGoBack.Call(webView, callback, param)
}

func GetCookie(webView WebView, callback GetCookieCallback, param uintptr) {
	_mbGetCookie.LoadOnce()
	_mbGetCookie.Call(webView, callback, param)
}

func GetCookieOnBlinkThread(webView WebView) string {
	_mbGetCookieOnBlinkThread.LoadOnce()
	return _mbGetCookieOnBlinkThread.Call(webView)
}

func ClearCookie(webView WebView) {
	_mbClearCookie.LoadOnce()
	_mbClearCookie.Call(webView)
}

func Resize(webView WebView, w, h int) {
	_mbResize.LoadOnce()
	_mbResize.Call(webView, w, h)
}

func OnNavigation(webView WebView, callback NavigationCallback, param uintptr) {
	_mbOnNavigation.LoadOnce()
	_mbOnNavigation.Call(webView, callback, param)
}

func OnNavigationSync(webView WebView, callback NavigationCallback, param uintptr) {
	_mbOnNavigationSync.LoadOnce()
	_mbOnNavigationSync.Call(webView, callback, param)
}

func OnCreateView(webView WebView, callback CreateViewCallback, param uintptr) {
	_mbOnCreateView.LoadOnce()
	_mbOnCreateView.Call(webView, callback, param)
}

func OnDocumentReady(webView WebView, callback DocumentReadyCallback, param uintptr) {
	_mbOnDocumentReady.LoadOnce()
	_mbOnDocumentReady.Call(webView, callback, param)
}

func OnDocumentReadyInBlinkThread(webView WebView, callback DocumentReadyCallback, param uintptr) {
	_mbOnDocumentReadyInBlinkThread.LoadOnce()
	_mbOnDocumentReadyInBlinkThread.Call(webView, callback, param)
}

func OnPaintUpdated(webView WebView, callback PaintUpdatedCallback, param uintptr) {
	_mbOnPaintUpdated.LoadOnce()
	_mbOnPaintUpdated.Call(webView, callback, param)
}

func OnPaintBitUpdated(webView WebView, callback PaintBitUpdatedCallback, param uintptr) {
	_mbOnPaintBitUpdated.LoadOnce()
	_mbOnPaintBitUpdated.Call(webView, callback, param)
}

func OnAcceleratedPaint(webView WebView, callback AcceleratedPaintCallback, param uintptr) {
	_mbOnAcceleratedPaint.LoadOnce()
	_mbOnAcceleratedPaint.Call(webView, callback, param)
}

func OnLoadUrlBegin(webView WebView, callback LoadUrlBeginCallback, param uintptr) {
	_mbOnLoadUrlBegin.LoadOnce()
	_mbOnLoadUrlBegin.Call(webView, callback, param)
}

func OnLoadUrlEnd(webView WebView, callback LoadUrlEndCallback, param uintptr) {
	_mbOnLoadUrlEnd.LoadOnce()
	_mbOnLoadUrlEnd.Call(webView, callback, param)
}

func OnLoadUrlFail(webView WebView, callback LoadUrlFailCallback, param uintptr) {
	_mbOnLoadUrlFail.LoadOnce()
	_mbOnLoadUrlFail.Call(webView, callback, param)
}

func OnLoadUrlHeadersReceived(webView WebView, callback LoadUrlHeadersReceivedCallback, param uintptr) {
	_mbOnLoadUrlHeadersReceived.LoadOnce()
	_mbOnLoadUrlHeadersReceived.Call(webView, callback, param)
}

func OnLoadUrlFinish(webView WebView, callback LoadUrlFinishCallback, param uintptr) {
	_mbOnLoadUrlFinish.LoadOnce()
	_mbOnLoadUrlFinish.Call(webView, callback, param)
}
func OnTitleChanged(webView WebView, callback TitleChangedCallback, callbackParam uintptr) {
	_mbOnTitleChanged.LoadOnce()
	_mbOnTitleChanged.Call(webView, callback, callbackParam)
}
func OnURLChanged(webView WebView, callback URLChangedCallback, callbackParam uintptr) {
	_mbOnURLChanged.LoadOnce()
	_mbOnURLChanged.Call(webView, callback, callbackParam)
}
func OnLoadingFinish(webView WebView, callback LoadingFinishCallback, param uintptr) {
	_mbOnLoadingFinish.LoadOnce()
	_mbOnLoadingFinish.Call(webView, callback, param)
}
func OnDownload(webView WebView, callback DownloadCallback, param uintptr) {
	_mbOnDownload.LoadOnce()
	_mbOnDownload.Call(webView, callback, param)
}
func OnDownloadInBlinkThread(webView WebView, callback DownloadInBlinkThreadCallback, param uintptr) {
	_mbOnDownloadInBlinkThread.LoadOnce()
	_mbOnDownloadInBlinkThread.Call(webView, callback, param)
}
func OnAlertBox(webView WebView, callback AlertBoxCallback, param uintptr) {
	_mbOnAlertBox.LoadOnce()
	_mbOnAlertBox.Call(webView, callback, param)
}
func OnConfirmBox(webView WebView, callback ConfirmBoxCallback, param uintptr) {
	_mbOnConfirmBox.LoadOnce()
	_mbOnConfirmBox.Call(webView, callback, param)
}
func OnPromptBox(webView WebView, callback PromptBoxCallback, param uintptr) {
	_mbOnPromptBox.LoadOnce()
	_mbOnPromptBox.Call(webView, callback, param)
}
func OnNetGetFavicon(webView WebView, callback NetGetFaviconCallback, param uintptr) {
	_mbOnNetGetFavicon.LoadOnce()
	_mbOnNetGetFavicon.Call(webView, callback, param)
}
func OnConsole(webView WebView, callback ConsoleCallback, param uintptr) {
	_mbOnConsole.LoadOnce()
	_mbOnConsole.Call(webView, callback, param)
}
func OnClose(webView WebView, callback CloseCallback, param uintptr) {
	_mbOnClose.LoadOnce()
	_mbOnClose.Call(webView, callback, param)
}
func OnDestroy(webView WebView, callback DestroyCallback, param uintptr) {
	_mbOnDestroy.LoadOnce()
	_mbOnDestroy.Call(webView, callback, param)
}
func OnPrinting(webView WebView, callback PrintingCallback, param uintptr) {
	_mbOnPrinting.LoadOnce()
	_mbOnPrinting.Call(webView, callback, param)
}
func OnPluginList(webView WebView, callback GetPluginListCallback, param uintptr) {
	_mbOnPluginList.LoadOnce()
	_mbOnPluginList.Call(webView, callback, param)
}
func OnImageBufferToDataURL(webView WebView, callback ImageBufferToDataURLCallback, param uintptr) {
	_mbOnImageBufferToDataURL.LoadOnce()
	_mbOnImageBufferToDataURL.Call(webView, callback, param)
}
func OnDidCreateScriptContext(webView WebView, callback DidCreateScriptContextCallback, param uintptr) {
	_mbOnDidCreateScriptContext.LoadOnce()
	_mbOnDidCreateScriptContext.Call(webView, callback, param)
}
func OnWillReleaseScriptContext(webView WebView, callback WillReleaseScriptContextCallback, param uintptr) {
	_mbOnWillReleaseScriptContext.LoadOnce()
	_mbOnWillReleaseScriptContext.Call(webView, callback, param)
}
func GoBack(webView WebView) {
	_mbGoBack.LoadOnce()
	_mbGoBack.Call(webView)
}
func GoForward(webView WebView) {
	_mbGoForward.LoadOnce()
	_mbGoForward.Call(webView)
}
func GoToOffset(webView WebView, offset int) {
	_mbGoToOffset.LoadOnce()
	_mbGoToOffset.Call(webView, offset)
}
func GoToIndex(webView WebView, index int) {
	_mbGoToIndex.LoadOnce()
	_mbGoToIndex.Call(webView, index)
}
func NavigateAtIndex(webView WebView, index int) {
	_mbNavigateAtIndex.LoadOnce()
	_mbNavigateAtIndex.Call(webView, index)
}
func GetNavigateIndex(webView WebView) int {
	_mbGetNavigateIndex.LoadOnce()
	return _mbGetNavigateIndex.Call(webView)
}
func StopLoading(webView WebView) {
	_mbStopLoading.LoadOnce()
	_mbStopLoading.Call(webView)
}
func Reload(webView WebView) {
	_mbReload.LoadOnce()
	_mbReload.Call(webView)
}
func PerformCookieCommand(webView WebView, command CookieCommand) {
	_mbPerformCookieCommand.LoadOnce()
	_mbPerformCookieCommand.Call(webView, command)
}
func InsertCSSByFrame(webView WebView, frameId WebFrameHandle, cssText string) {
	_mbInsertCSSByFrame.LoadOnce()
	_mbInsertCSSByFrame.Call(webView, frameId, cssText)
}
func EditorSelectAll(webView WebView) {
	_mbEditorSelectAll.LoadOnce()
	_mbEditorSelectAll.Call(webView)
}
func EditorUnSelect(webView WebView) {
	_mbEditorUnSelect.LoadOnce()
	_mbEditorUnSelect.Call(webView)
}
func EditorCopy(webView WebView) {
	_mbEditorCopy.LoadOnce()
	_mbEditorCopy.Call(webView)
}
func EditorCut(webView WebView) {
	_mbEditorCut.LoadOnce()
	_mbEditorCut.Call(webView)
}
func EditorPaste(webView WebView) {
	_mbEditorPaste.LoadOnce()
	_mbEditorPaste.Call(webView)
}
func EditorDelete(webView WebView) {
	_mbEditorDelete.LoadOnce()
	_mbEditorDelete.Call(webView)
}
func EditorUndo(webView WebView) {
	_mbEditorUndo.LoadOnce()
	_mbEditorUndo.Call(webView)
}
func EditorRedo(webView WebView) {
	_mbEditorRedo.LoadOnce()
	_mbEditorRedo.Call(webView)
}
func SetEditable(webView WebView, editable bool) {
	_mbSetEditable.LoadOnce()
	_mbSetEditable.Call(webView, editable)
}
func FireMouseEvent(webView WebView, message uint, x, y int, flags uint) bool {
	_mbFireMouseEvent.LoadOnce()
	return _mbFireMouseEvent.Call(webView, message, x, y, flags)
}
func FireContextMenuEvent(webView WebView, x, y int, flags uint) bool {
	_mbFireContextMenuEvent.LoadOnce()
	return _mbFireContextMenuEvent.Call(webView, x, y, flags)
}
func FireMouseWheelEvent(webView WebView, x, y, delta int, flags uint) bool {
	_mbFireMouseWheelEvent.LoadOnce()
	return _mbFireMouseWheelEvent.Call(webView, x, y, delta, flags)
}
func FireKeyUpEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	_mbFireKeyUpEvent.LoadOnce()
	return _mbFireKeyUpEvent.Call(webView, virtualKeyCode, flags, systemKey)
}
func FireKeyDownEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	_mbFireKeyDownEvent.LoadOnce()
	return _mbFireKeyDownEvent.Call(webView, virtualKeyCode, flags, systemKey)
}
func FireKeyPressEvent(webView WebView, charCode, flags uint, systemKey bool) bool {
	_mbFireKeyPressEvent.LoadOnce()
	return _mbFireKeyPressEvent.Call(webView, charCode, flags, systemKey)
}
func FireWindowsMessage(webView WebView, hWnd HWND, message uint, wParam WPARAM, lParam LPARAM, result *LRESULT) bool {
	_mbFireWindowsMessage.LoadOnce()
	return _mbFireWindowsMessage.Call(webView, hWnd, message, wParam, lParam, result)
}
func SetFocus(webView WebView) {
	_mbSetFocus.LoadOnce()
	_mbSetFocus.Call(webView)
}
func KillFocus(webView WebView) {
	_mbKillFocus.LoadOnce()
	_mbKillFocus.Call(webView)
}
func ShowWindow(webView WebView, show bool) {
	_mbShowWindow.LoadOnce()
	_mbShowWindow.Call(webView, show)
}

func LoadURL(webView WebView, url string) {
	_mbLoadURL.LoadOnce()
	_mbLoadURL.Call(webView, url)
}

func LoadHtmlWithBaseUrl(webView WebView, html string, baseUrl string) {
	_mbLoadHtmlWithBaseUrl.LoadOnce()
	_mbLoadHtmlWithBaseUrl.Call(webView, html, baseUrl)
}
func PostURL(webView WebView, url string, postData *byte, postLen int) {
	_mbPostURL.LoadOnce()
	_mbPostURL.Call(webView, url, postData, postLen)
}
func GetLockedViewDC(webView WebView) HDC {
	_mbGetLockedViewDC.LoadOnce()
	return _mbGetLockedViewDC.Call(webView)
}
func UnlockViewDC(webView WebView) {
	_mbUnlockViewDC.LoadOnce()
	_mbUnlockViewDC.Call(webView)
}
func Wake(webView WebView) {
	_mbWake.LoadOnce()
	_mbWake.Call(webView)
}
func JsToV8Value(es JsExecState, v JsValue) uintptr {
	_mbJsToV8Value.LoadOnce()
	return _mbJsToV8Value.Call(es, v)
}
func GetGlobalExecByFrame(webView WebView, frameId WebFrameHandle) JsExecState {
	_mbGetGlobalExecByFrame.LoadOnce()
	return _mbGetGlobalExecByFrame.Call(webView, frameId)
}
func JsToDouble(es JsExecState, v JsValue) float64 {
	_mbJsToDouble.LoadOnce()
	return _mbJsToDouble.Call(es, v)
}
func JsToBoolean(es JsExecState, v JsValue) bool {
	_mbJsToBoolean.LoadOnce()
	return _mbJsToBoolean.Call(es, v)
}
func JsToString(es JsExecState, v JsValue) string {
	_mbJsToString.LoadOnce()
	return _mbJsToString.Call(es, v)
}
func GetJsValueType(es JsExecState, v JsValue) JsType {
	_mbGetJsValueType.LoadOnce()
	return _mbGetJsValueType.Call(es, v)
}
func OnJsQuery(webView WebView, callback JsQueryCallback, param uintptr) {
	_mbOnJsQuery.LoadOnce()
	_mbOnJsQuery.Call(webView, callback, param)
}
func ResponseQuery(webView WebView, queryId int64, customMsg int, response string) {
	_mbResponseQuery.LoadOnce()
	_mbResponseQuery.Call(webView, queryId, customMsg, response)
}
func RunJs(webView WebView, frameId WebFrameHandle, script string, isInClosure bool, callback RunJsCallback, param, unused uintptr) {
	_mbRunJs.LoadOnce()
	_mbRunJs.Call(webView, frameId, script, isInClosure, callback, param, unused)
}
func RunJsSync(webView WebView, frameId WebFrameHandle, script string, isInClosure bool) JsValue {
	_mbRunJsSync.LoadOnce()
	return _mbRunJsSync.Call(webView, frameId, script, isInClosure)
}
func WebFrameGetMainFrame(webView WebView) WebFrameHandle {
	_mbWebFrameGetMainFrame.LoadOnce()
	return _mbWebFrameGetMainFrame.Call(webView)
}
func IsMainFrame(webView WebView, frameId WebFrameHandle) bool {
	_mbIsMainFrame.LoadOnce()
	return _mbIsMainFrame.Call(webView, frameId)
}
func SetNodeJsEnable(webView WebView, enable bool) {
	_mbSetNodeJsEnable.LoadOnce()
	_mbSetNodeJsEnable.Call(webView, enable)
}
func SetDeviceParameter(webView WebView, device, paramStr string, paramInt int, paramFloat float32) {
	_mbSetDeviceParameter.LoadOnce()
	_mbSetDeviceParameter.Call(webView, device, paramStr, paramInt, paramFloat)
}
func GetContentAsMarkup(webView WebView, callback GetContentAsMarkupCallback, param uintptr, frameId WebFrameHandle) {
	_mbGetContentAsMarkup.LoadOnce()
	_mbGetContentAsMarkup.Call(webView, callback, param, frameId)
}
func GetSource(webView WebView, callback GetSourceCallback, param uintptr) {
	_mbGetSource.LoadOnce()
	_mbGetSource.Call(webView, callback, param)
}
func UtilSerializeToMHTML(webView WebView, callback GetSourceCallback, param uintptr) {
	_mbUtilSerializeToMHTML.LoadOnce()
	_mbUtilSerializeToMHTML.Call(webView, callback, param)
}
func UtilCreateRequestCode(registerInfo string) string {
	_mbUtilCreateRequestCode.LoadOnce()
	return _mbUtilCreateRequestCode.Call(registerInfo)
}
func UtilIsRegistered(defaultPath string) bool {
	_mbUtilIsRegistered.LoadOnce()
	return _mbUtilIsRegistered.Call(defaultPath)
}
func UtilPrint(webView WebView, frameId WebFrameHandle, printParams *PrintSettings) bool {
	_mbUtilPrint.LoadOnce()
	return _mbUtilPrint.Call(webView, frameId, printParams)
}
func UtilBase64Encode(str string) string {
	_mbUtilBase64Encode.LoadOnce()
	return _mbUtilBase64Encode.Call(str)
}
func UtilBase64Decode(str string) string {
	_mbUtilBase64Decode.LoadOnce()
	return _mbUtilBase64Decode.Call(str)
}
func UtilDecodeURLEscape(url string) string {
	_mbUtilDecodeURLEscape.LoadOnce()
	return _mbUtilDecodeURLEscape.Call(url)
}
func UtilEncodeURLEscape(url string) string {
	_mbUtilEncodeURLEscape.LoadOnce()
	return _mbUtilEncodeURLEscape.Call(url)
}
func UtilCreateV8Snapshot(str string) *MemBuf {
	_mbUtilCreateV8Snapshot.LoadOnce()
	return _mbUtilCreateV8Snapshot.Call(str)
}
func UtilPrintToPdf(webView WebView, frameId WebFrameHandle, settings *PrintSettings, callback PrintPdfDataCallback, param uintptr) {
	_mbUtilPrintToPdf.LoadOnce()
	_mbUtilPrintToPdf.Call(webView, frameId, settings, callback, param)
}
func UtilPrintToBitmap(webView WebView, frameId WebFrameHandle, settings *ScreenshotSettings, callback PrintBitmapCallback, param uintptr) {
	_mbUtilPrintToBitmap.LoadOnce()
	_mbUtilPrintToBitmap.Call(webView, frameId, settings, callback, param)
}
func UtilScreenshot(webView WebView, settings *ScreenshotSettings, callback OnScreenshotCallback, param uintptr) {
	_mbUtilScreenshot.LoadOnce()
	_mbUtilScreenshot.Call(webView, settings, callback, param)
}
func UtilsSilentPrint(webView WebView, settings string) bool {
	_mbUtilsSilentPrint.LoadOnce()
	return _mbUtilsSilentPrint.Call(webView, settings)
}
func UtilSetDefaultPrinterSettings(webView WebView, setting *DefaultPrinterSettings) {
	_mbUtilSetDefaultPrinterSettings.LoadOnce()
	_mbUtilSetDefaultPrinterSettings.Call(webView, setting)
}
func PopupDownloadMgr(webView WebView, url string, downloadJob uintptr) bool {
	_mbPopupDownloadMgr.LoadOnce()
	return _mbPopupDownloadMgr.Call(webView, url, downloadJob)
}
func PopupDialogAndDownload(webView WebView, dialogOpt *DialogOptions, contentLength uint64, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	_mbPopupDialogAndDownload.LoadOnce()
	return _mbPopupDialogAndDownload.Call(webView, dialogOpt, contentLength, url, mime, disposition, job, dataBind, callbackBind)
}
func DownloadByPath(webView WebView, downloadOptions *DownloadOptions, path string, contentLength uint64, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	_mbDownloadByPath.LoadOnce()
	return _mbDownloadByPath.Call(webView, downloadOptions, path, contentLength, url, mime, disposition, job, dataBind, callbackBind)
}
func GetPdfPageData(webView WebView, callback OnGetPdfPageDataCallback, param uintptr) {
	_mbGetPdfPageData.LoadOnce()
	_mbGetPdfPageData.Call(webView, callback, param)
}
func CreateMemBuf(webView WebView, buf *byte, length uint64) *MemBuf {
	_mbCreateMemBuf.LoadOnce()
	return _mbCreateMemBuf.Call(webView, buf, length)
}
func FreeMemBuf(buf *MemBuf) {
	_mbFreeMemBuf.LoadOnce()
	_mbFreeMemBuf.Call(buf)
}
func SetUserKeyValue(webView WebView, key string, value uintptr) {
	_mbSetUserKeyValue.LoadOnce()
	_mbSetUserKeyValue.Call(webView, key, value)
}
func GetUserKeyValue(webView WebView, key string) uintptr {
	_mbGetUserKeyValue.LoadOnce()
	return _mbGetUserKeyValue.Call(webView, key)
}
func PluginListBuilderAddPlugin(builder uintptr, name, description, fileName string) {
	_mbPluginListBuilderAddPlugin.LoadOnce()
	_mbPluginListBuilderAddPlugin.Call(builder, name, description, fileName)
}
func PluginListBuilderAddMediaTypeToLastPlugin(builder uintptr, name, description string) {
	_mbPluginListBuilderAddMediaTypeToLastPlugin.LoadOnce()
	_mbPluginListBuilderAddMediaTypeToLastPlugin.Call(builder, name, description)
}
func PluginListBuilderAddFileExtensionToLastMediaType(builder uintptr, fileExtension string) {
	_mbPluginListBuilderAddFileExtensionToLastMediaType.LoadOnce()
	_mbPluginListBuilderAddFileExtensionToLastMediaType.Call(builder, fileExtension)
}
func GetBlinkMainThreadIsolate() V8Isolate {
	_mbGetBlinkMainThreadIsolate.LoadOnce()
	return _mbGetBlinkMainThreadIsolate.Call()
}
func WebFrameGetMainWorldScriptContext(webView WebView, frameId WebFrameHandle, contextOut *V8ContextPtr) {
	_mbWebFrameGetMainWorldScriptContext.LoadOnce()
	_mbWebFrameGetMainWorldScriptContext.Call(webView, frameId, contextOut)
}
func EnableHighDPISupport() {
	_mbEnableHighDPISupport.LoadOnce()
	_mbEnableHighDPISupport.Call()
}
func RunMessageLoop() {
	_mbRunMessageLoop.LoadOnce()
	_mbRunMessageLoop.Call()
}
func GetContentWidth(webView WebView) int {
	_mbGetContentWidth.LoadOnce()
	return _mbGetContentWidth.Call(webView)
}
func GetContentHeight(webView WebView) int {
	_mbGetContentHeight.LoadOnce()
	return _mbGetContentHeight.Call(webView)
}
func GetWebViewForCurrentContext() WebView {
	_mbGetWebViewForCurrentContext.LoadOnce()
	return _mbGetWebViewForCurrentContext.Call()
}
func RegisterEmbedderCustomElement(webviewHandle WebView, frameId WebFrameHandle, name string, options, outResult uintptr) bool {
	_mbRegisterEmbedderCustomElement.LoadOnce()
	return _mbRegisterEmbedderCustomElement.Call(webviewHandle, frameId, name, options, outResult)
}
func OnNodeCreateProcess(webviewHandle WebView, callback NodeOnCreateProcessCallback, param uintptr) {
	_mbOnNodeCreateProcess.LoadOnce()
	_mbOnNodeCreateProcess.Call(webviewHandle, callback, param)
}
func OnThreadIdle(callback ThreadCallback, param1, param2 uintptr) {
	_mbOnThreadIdle.LoadOnce()
	_mbOnThreadIdle.Call(callback, param1, param2)
}
func GetProcAddr(name string) uintptr {
	_mbGetProcAddr.LoadOnce()
	return _mbGetProcAddr.Call(name)
}
func SetMbDllPath(dllPath string) {
	_mbSetMbDllPath.LoadOnce()
	_mbSetMbDllPath.Call(dllPath)
}
func SetMbMainDllPath(dllPath string) {
	_mbSetMbMainDllPath.LoadOnce()
	_mbSetMbMainDllPath.Call(dllPath)
}
func FillFuncPtr() {
	_mbFillFuncPtr.LoadOnce()
	_mbFillFuncPtr.Call()
}
