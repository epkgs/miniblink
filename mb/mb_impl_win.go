//go:build windows
// +build windows

package mb

import (
	"fmt"
	"unsafe"

	"github.com/epkgs/miniblink/internal/lib"
	"golang.org/x/sys/windows"
)

var mbDLL *windows.LazyDLL

var _mbInit *windows.LazyProc
var _mbUninit *windows.LazyProc
var _mbCreateInitSettings *windows.LazyProc
var _mbSetInitSettings *windows.LazyProc
var _mbCreateWebView *windows.LazyProc
var _mbDestroyWebView *windows.LazyProc
var _mbCreateWebWindow *windows.LazyProc
var _mbCreateWebCustomWindow *windows.LazyProc
var _mbMoveWindow *windows.LazyProc
var _mbMoveToCenter *windows.LazyProc
var _mbSetAutoDrawToHwnd *windows.LazyProc
var _mbGetCaretRect *windows.LazyProc
var _mbSetAudioMuted *windows.LazyProc
var _mbIsAudioMuted *windows.LazyProc
var _mbCreateString *windows.LazyProc
var _mbCreateStringWithoutNullTermination *windows.LazyProc
var _mbDeleteString *windows.LazyProc
var _mbGetStringLen *windows.LazyProc
var _mbGetString *windows.LazyProc
var _mbSetProxy *windows.LazyProc
var _mbSetDebugConfig *windows.LazyProc
var _mbNetSetData *windows.LazyProc
var _mbNetHookRequest *windows.LazyProc
var _mbNetChangeRequestUrl *windows.LazyProc
var _mbNetContinueJob *windows.LazyProc
var _mbNetGetRawHttpHeadInBlinkThread *windows.LazyProc
var _mbNetGetRawResponseHeadInBlinkThread *windows.LazyProc
var _mbNetHoldJobToAsynCommit *windows.LazyProc
var _mbNetCancelRequest *windows.LazyProc
var _mbNetOnResponse *windows.LazyProc
var _mbNetSetWebsocketCallback *windows.LazyProc
var _mbNetSendWsText *windows.LazyProc
var _mbNetSendWsBlob *windows.LazyProc
var _mbNetEnableResPacket *windows.LazyProc
var _mbNetGetPostBody *windows.LazyProc
var _mbNetCreatePostBodyElements *windows.LazyProc
var _mbNetFreePostBodyElements *windows.LazyProc
var _mbNetCreatePostBodyElement *windows.LazyProc
var _mbNetFreePostBodyElement *windows.LazyProc
var _mbNetCreateWebUrlRequest *windows.LazyProc
var _mbNetAddHTTPHeaderFieldToUrlRequest *windows.LazyProc
var _mbNetStartUrlRequest *windows.LazyProc
var _mbNetGetHttpStatusCode *windows.LazyProc
var _mbNetGetRequestMethod *windows.LazyProc
var _mbNetGetExpectedContentLength *windows.LazyProc
var _mbNetGetResponseUrl *windows.LazyProc
var _mbNetCancelWebUrlRequest *windows.LazyProc
var _mbSetViewProxy *windows.LazyProc
var _mbNetSetMIMEType *windows.LazyProc
var _mbNetGetMIMEType *windows.LazyProc
var _mbNetGetHTTPHeaderField *windows.LazyProc
var _mbNetGetReferrer *windows.LazyProc
var _mbNetSetHTTPHeaderField *windows.LazyProc
var _mbSetMouseEnabled *windows.LazyProc
var _mbSetTouchEnabled *windows.LazyProc
var _mbSetSystemTouchEnabled *windows.LazyProc
var _mbSetContextMenuEnabled *windows.LazyProc
var _mbSetNavigationToNewWindowEnabled *windows.LazyProc
var _mbSetHeadlessEnabled *windows.LazyProc
var _mbSetDragDropEnable *windows.LazyProc
var _mbSetDragEnable *windows.LazyProc
var _mbSetContextMenuItemShow *windows.LazyProc
var _mbSetHandle *windows.LazyProc
var _mbSetHandleOffset *windows.LazyProc
var _mbGetHostHWND *windows.LazyProc
var _mbSetTransparent *windows.LazyProc
var _mbSetViewSettings *windows.LazyProc
var _mbSetCspCheckEnable *windows.LazyProc
var _mbSetNpapiPluginsEnabled *windows.LazyProc
var _mbSetMemoryCacheEnable *windows.LazyProc
var _mbSetCookie *windows.LazyProc
var _mbSetCookieEnabled *windows.LazyProc
var _mbSetCookieJarPath *windows.LazyProc
var _mbSetCookieJarFullPath *windows.LazyProc
var _mbSetLocalStorageFullPath *windows.LazyProc
var _mbGetTitle *windows.LazyProc
var _mbSetWindowTitle *windows.LazyProc
var _mbSetWindowTitleW *windows.LazyProc
var _mbGetUrl *windows.LazyProc
var _mbGetCursorInfoType *windows.LazyProc
var _mbAddPluginDirectory *windows.LazyProc
var _mbSetUserAgent *windows.LazyProc
var _mbSetZoomFactor *windows.LazyProc
var _mbGetZoomFactor *windows.LazyProc
var _mbSetDiskCacheEnabled *windows.LazyProc
var _mbSetDiskCachePath *windows.LazyProc
var _mbSetDiskCacheLimit *windows.LazyProc
var _mbSetDiskCacheLimitDisk *windows.LazyProc
var _mbSetDiskCacheLevel *windows.LazyProc
var _mbSetResourceGc *windows.LazyProc
var _mbCanGoForward *windows.LazyProc
var _mbCanGoBack *windows.LazyProc
var _mbGetCookie *windows.LazyProc
var _mbGetCookieOnBlinkThread *windows.LazyProc
var _mbClearCookie *windows.LazyProc
var _mbResize *windows.LazyProc
var _mbOnNavigation *windows.LazyProc
var _mbOnNavigationSync *windows.LazyProc
var _mbOnCreateView *windows.LazyProc
var _mbOnDocumentReady *windows.LazyProc
var _mbOnDocumentReadyInBlinkThread *windows.LazyProc
var _mbOnPaintUpdated *windows.LazyProc
var _mbOnPaintBitUpdated *windows.LazyProc
var _mbOnAcceleratedPaint *windows.LazyProc
var _mbOnLoadUrlBegin *windows.LazyProc
var _mbOnLoadUrlEnd *windows.LazyProc
var _mbOnLoadUrlFail *windows.LazyProc
var _mbOnLoadUrlHeadersReceived *windows.LazyProc
var _mbOnLoadUrlFinish *windows.LazyProc
var _mbOnTitleChanged *windows.LazyProc
var _mbOnURLChanged *windows.LazyProc
var _mbOnLoadingFinish *windows.LazyProc
var _mbOnDownload *windows.LazyProc
var _mbOnDownloadInBlinkThread *windows.LazyProc
var _mbOnAlertBox *windows.LazyProc
var _mbOnConfirmBox *windows.LazyProc
var _mbOnPromptBox *windows.LazyProc
var _mbOnNetGetFavicon *windows.LazyProc
var _mbOnConsole *windows.LazyProc
var _mbOnClose *windows.LazyProc
var _mbOnDestroy *windows.LazyProc
var _mbOnPrinting *windows.LazyProc
var _mbOnPluginList *windows.LazyProc
var _mbOnImageBufferToDataURL *windows.LazyProc
var _mbOnDidCreateScriptContext *windows.LazyProc
var _mbOnWillReleaseScriptContext *windows.LazyProc
var _mbGoBack *windows.LazyProc
var _mbGoForward *windows.LazyProc
var _mbGoToOffset *windows.LazyProc
var _mbGoToIndex *windows.LazyProc
var _mbNavigateAtIndex *windows.LazyProc
var _mbGetNavigateIndex *windows.LazyProc
var _mbStopLoading *windows.LazyProc
var _mbReload *windows.LazyProc
var _mbPerformCookieCommand *windows.LazyProc
var _mbInsertCSSByFrame *windows.LazyProc
var _mbEditorSelectAll *windows.LazyProc
var _mbEditorUnSelect *windows.LazyProc
var _mbEditorCopy *windows.LazyProc
var _mbEditorCut *windows.LazyProc
var _mbEditorPaste *windows.LazyProc
var _mbEditorDelete *windows.LazyProc
var _mbEditorUndo *windows.LazyProc
var _mbEditorRedo *windows.LazyProc
var _mbSetEditable *windows.LazyProc
var _mbFireMouseEvent *windows.LazyProc
var _mbFireContextMenuEvent *windows.LazyProc
var _mbFireMouseWheelEvent *windows.LazyProc
var _mbFireKeyUpEvent *windows.LazyProc
var _mbFireKeyDownEvent *windows.LazyProc
var _mbFireKeyPressEvent *windows.LazyProc
var _mbFireWindowsMessage *windows.LazyProc
var _mbSetFocus *windows.LazyProc
var _mbKillFocus *windows.LazyProc
var _mbShowWindow *windows.LazyProc
var _mbLoadURL *windows.LazyProc
var _mbLoadHtmlWithBaseUrl *windows.LazyProc
var _mbPostURL *windows.LazyProc
var _mbGetLockedViewDC *windows.LazyProc
var _mbUnlockViewDC *windows.LazyProc
var _mbWake *windows.LazyProc
var _mbJsToV8Value *windows.LazyProc
var _mbGetGlobalExecByFrame *windows.LazyProc
var _mbJsToDouble *windows.LazyProc
var _mbJsToBoolean *windows.LazyProc
var _mbJsToString *windows.LazyProc
var _mbGetJsValueType *windows.LazyProc
var _mbOnJsQuery *windows.LazyProc
var _mbResponseQuery *windows.LazyProc
var _mbRunJs *windows.LazyProc
var _mbRunJsSync *windows.LazyProc
var _mbWebFrameGetMainFrame *windows.LazyProc
var _mbIsMainFrame *windows.LazyProc
var _mbSetNodeJsEnable *windows.LazyProc
var _mbSetDeviceParameter *windows.LazyProc
var _mbGetContentAsMarkup *windows.LazyProc
var _mbGetSource *windows.LazyProc
var _mbUtilSerializeToMHTML *windows.LazyProc
var _mbUtilCreateRequestCode *windows.LazyProc
var _mbUtilIsRegistered *windows.LazyProc
var _mbUtilPrint *windows.LazyProc
var _mbUtilBase64Encode *windows.LazyProc
var _mbUtilBase64Decode *windows.LazyProc
var _mbUtilDecodeURLEscape *windows.LazyProc
var _mbUtilEncodeURLEscape *windows.LazyProc
var _mbUtilCreateV8Snapshot *windows.LazyProc
var _mbUtilPrintToPdf *windows.LazyProc
var _mbUtilPrintToBitmap *windows.LazyProc
var _mbUtilScreenshot *windows.LazyProc
var _mbUtilsSilentPrint *windows.LazyProc
var _mbUtilSetDefaultPrinterSettings *windows.LazyProc
var _mbPopupDownloadMgr *windows.LazyProc
var _mbPopupDialogAndDownload *windows.LazyProc
var _mbDownloadByPath *windows.LazyProc
var _mbGetPdfPageData *windows.LazyProc
var _mbCreateMemBuf *windows.LazyProc
var _mbFreeMemBuf *windows.LazyProc
var _mbSetUserKeyValue *windows.LazyProc
var _mbGetUserKeyValue *windows.LazyProc
var _mbPluginListBuilderAddPlugin *windows.LazyProc
var _mbPluginListBuilderAddMediaTypeToLastPlugin *windows.LazyProc
var _mbPluginListBuilderAddFileExtensionToLastMediaType *windows.LazyProc
var _mbGetBlinkMainThreadIsolate *windows.LazyProc
var _mbWebFrameGetMainWorldScriptContext *windows.LazyProc
var _mbEnableHighDPISupport *windows.LazyProc
var _mbRunMessageLoop *windows.LazyProc
var _mbGetContentWidth *windows.LazyProc
var _mbGetContentHeight *windows.LazyProc
var _mbGetWebViewForCurrentContext *windows.LazyProc
var _mbRegisterEmbedderCustomElement *windows.LazyProc
var _mbOnNodeCreateProcess *windows.LazyProc
var _mbOnThreadIdle *windows.LazyProc
var _mbGetProcAddr *windows.LazyProc
var _mbSetMbDllPath *windows.LazyProc
var _mbSetMbMainDllPath *windows.LazyProc
var _mbFillFuncPtr *windows.LazyProc

func LoadLibrary(name string) error {
	dll, err := lib.LoadLibrary(name)

	if err != nil {
		return err
	}

	mbDLL = windows.NewLazyDLL(dll.Name)

	fmt.Printf("DLL Name: %s\n", mbDLL.Name)

	initProcs()

	return nil
}

func initProcs() {

	_mbInit = mbDLL.NewProc("mbInit")
	_mbUninit = mbDLL.NewProc("mbUninit")
	_mbCreateInitSettings = mbDLL.NewProc("mbCreateInitSettings")
	_mbSetInitSettings = mbDLL.NewProc("mbSetInitSettings")
	_mbCreateWebView = mbDLL.NewProc("mbCreateWebView")
	_mbDestroyWebView = mbDLL.NewProc("mbDestroyWebView")
	_mbCreateWebWindow = mbDLL.NewProc("mbCreateWebWindow")
	_mbCreateWebCustomWindow = mbDLL.NewProc("mbCreateWebCustomWindow")
	_mbMoveWindow = mbDLL.NewProc("mbMoveWindow")
	_mbMoveToCenter = mbDLL.NewProc("mbMoveToCenter")
	_mbSetAutoDrawToHwnd = mbDLL.NewProc("mbSetAutoDrawToHwnd")
	_mbGetCaretRect = mbDLL.NewProc("mbGetCaretRect")
	_mbSetAudioMuted = mbDLL.NewProc("mbSetAudioMuted")
	_mbIsAudioMuted = mbDLL.NewProc("mbIsAudioMuted")
	_mbCreateString = mbDLL.NewProc("mbCreateString")
	_mbCreateStringWithoutNullTermination = mbDLL.NewProc("mbCreateStringWithoutNullTermination")
	_mbDeleteString = mbDLL.NewProc("mbDeleteString")
	_mbGetStringLen = mbDLL.NewProc("mbGetStringLen")
	_mbGetString = mbDLL.NewProc("mbGetString")
	_mbSetProxy = mbDLL.NewProc("mbSetProxy")
	_mbSetDebugConfig = mbDLL.NewProc("mbSetDebugConfig")
	_mbNetSetData = mbDLL.NewProc("mbNetSetData")
	_mbNetHookRequest = mbDLL.NewProc("mbNetHookRequest")
	_mbNetChangeRequestUrl = mbDLL.NewProc("mbNetChangeRequestUrl")
	_mbNetContinueJob = mbDLL.NewProc("mbNetContinueJob")
	_mbNetGetRawHttpHeadInBlinkThread = mbDLL.NewProc("mbNetGetRawHttpHeadInBlinkThread")
	_mbNetGetRawResponseHeadInBlinkThread = mbDLL.NewProc("mbNetGetRawResponseHeadInBlinkThread")
	_mbNetHoldJobToAsynCommit = mbDLL.NewProc("mbNetHoldJobToAsynCommit")
	_mbNetCancelRequest = mbDLL.NewProc("mbNetCancelRequest")
	_mbNetOnResponse = mbDLL.NewProc("mbNetOnResponse")
	_mbNetSetWebsocketCallback = mbDLL.NewProc("mbNetSetWebsocketCallback")
	_mbNetSendWsText = mbDLL.NewProc("mbNetSendWsText")
	_mbNetSendWsBlob = mbDLL.NewProc("mbNetSendWsBlob")
	_mbNetEnableResPacket = mbDLL.NewProc("mbNetEnableResPacket")
	_mbNetGetPostBody = mbDLL.NewProc("mbNetGetPostBody")
	_mbNetCreatePostBodyElements = mbDLL.NewProc("mbNetCreatePostBodyElements")
	_mbNetFreePostBodyElements = mbDLL.NewProc("mbNetFreePostBodyElements")
	_mbNetCreatePostBodyElement = mbDLL.NewProc("mbNetCreatePostBodyElement")
	_mbNetFreePostBodyElement = mbDLL.NewProc("mbNetFreePostBodyElement")
	_mbNetCreateWebUrlRequest = mbDLL.NewProc("mbNetCreateWebUrlRequest")
	_mbNetAddHTTPHeaderFieldToUrlRequest = mbDLL.NewProc("mbNetAddHTTPHeaderFieldToUrlRequest")
	_mbNetStartUrlRequest = mbDLL.NewProc("mbNetStartUrlRequest")
	_mbNetGetHttpStatusCode = mbDLL.NewProc("mbNetGetHttpStatusCode")
	_mbNetGetRequestMethod = mbDLL.NewProc("mbNetGetRequestMethod")
	_mbNetGetExpectedContentLength = mbDLL.NewProc("mbNetGetExpectedContentLength")
	_mbNetGetResponseUrl = mbDLL.NewProc("mbNetGetResponseUrl")
	_mbNetCancelWebUrlRequest = mbDLL.NewProc("mbNetCancelWebUrlRequest")
	_mbSetViewProxy = mbDLL.NewProc("mbSetViewProxy")
	_mbNetSetMIMEType = mbDLL.NewProc("mbNetSetMIMEType")
	_mbNetGetMIMEType = mbDLL.NewProc("mbNetGetMIMEType")
	_mbNetGetHTTPHeaderField = mbDLL.NewProc("mbNetGetHTTPHeaderField")
	_mbNetGetReferrer = mbDLL.NewProc("mbNetGetReferrer")
	_mbNetSetHTTPHeaderField = mbDLL.NewProc("mbNetSetHTTPHeaderField")
	_mbSetMouseEnabled = mbDLL.NewProc("mbSetMouseEnabled")
	_mbSetTouchEnabled = mbDLL.NewProc("mbSetTouchEnabled")
	_mbSetSystemTouchEnabled = mbDLL.NewProc("mbSetSystemTouchEnabled")
	_mbSetContextMenuEnabled = mbDLL.NewProc("mbSetContextMenuEnabled")
	_mbSetNavigationToNewWindowEnabled = mbDLL.NewProc("mbSetNavigationToNewWindowEnabled")
	_mbSetHeadlessEnabled = mbDLL.NewProc("mbSetHeadlessEnabled")
	_mbSetDragDropEnable = mbDLL.NewProc("mbSetDragDropEnable")
	_mbSetDragEnable = mbDLL.NewProc("mbSetDragEnable")
	_mbSetContextMenuItemShow = mbDLL.NewProc("mbSetContextMenuItemShow")
	_mbSetHandle = mbDLL.NewProc("mbSetHandle")
	_mbSetHandleOffset = mbDLL.NewProc("mbSetHandleOffset")
	_mbGetHostHWND = mbDLL.NewProc("mbGetHostHWND")
	_mbSetTransparent = mbDLL.NewProc("mbSetTransparent")
	_mbSetViewSettings = mbDLL.NewProc("mbSetViewSettings")
	_mbSetCspCheckEnable = mbDLL.NewProc("mbSetCspCheckEnable")
	_mbSetNpapiPluginsEnabled = mbDLL.NewProc("mbSetNpapiPluginsEnabled")
	_mbSetMemoryCacheEnable = mbDLL.NewProc("mbSetMemoryCacheEnable")
	_mbSetCookie = mbDLL.NewProc("mbSetCookie")
	_mbSetCookieEnabled = mbDLL.NewProc("mbSetCookieEnabled")
	_mbSetCookieJarPath = mbDLL.NewProc("mbSetCookieJarPath")
	_mbSetCookieJarFullPath = mbDLL.NewProc("mbSetCookieJarFullPath")
	_mbSetLocalStorageFullPath = mbDLL.NewProc("mbSetLocalStorageFullPath")
	_mbGetTitle = mbDLL.NewProc("mbGetTitle")
	_mbSetWindowTitle = mbDLL.NewProc("mbSetWindowTitle")
	_mbSetWindowTitleW = mbDLL.NewProc("mbSetWindowTitleW")
	_mbGetUrl = mbDLL.NewProc("mbGetUrl")
	_mbGetCursorInfoType = mbDLL.NewProc("mbGetCursorInfoType")
	_mbAddPluginDirectory = mbDLL.NewProc("mbAddPluginDirectory")
	_mbSetUserAgent = mbDLL.NewProc("mbSetUserAgent")
	_mbSetZoomFactor = mbDLL.NewProc("mbSetZoomFactor")
	_mbGetZoomFactor = mbDLL.NewProc("mbGetZoomFactor")
	_mbSetDiskCacheEnabled = mbDLL.NewProc("mbSetDiskCacheEnabled")
	_mbSetDiskCachePath = mbDLL.NewProc("mbSetDiskCachePath")
	_mbSetDiskCacheLimit = mbDLL.NewProc("mbSetDiskCacheLimit")
	_mbSetDiskCacheLimitDisk = mbDLL.NewProc("mbSetDiskCacheLimitDisk")
	_mbSetDiskCacheLevel = mbDLL.NewProc("mbSetDiskCacheLevel")
	_mbSetResourceGc = mbDLL.NewProc("mbSetResourceGc")
	_mbCanGoForward = mbDLL.NewProc("mbCanGoForward")
	_mbCanGoBack = mbDLL.NewProc("mbCanGoBack")
	_mbGetCookie = mbDLL.NewProc("mbGetCookie")
	_mbGetCookieOnBlinkThread = mbDLL.NewProc("mbGetCookieOnBlinkThread")
	_mbClearCookie = mbDLL.NewProc("mbClearCookie")
	_mbResize = mbDLL.NewProc("mbResize")
	_mbOnNavigation = mbDLL.NewProc("mbOnNavigation")
	_mbOnNavigationSync = mbDLL.NewProc("mbOnNavigationSync")
	_mbOnCreateView = mbDLL.NewProc("mbOnCreateView")
	_mbOnDocumentReady = mbDLL.NewProc("mbOnDocumentReady")
	_mbOnDocumentReadyInBlinkThread = mbDLL.NewProc("mbOnDocumentReadyInBlinkThread")
	_mbOnPaintUpdated = mbDLL.NewProc("mbOnPaintUpdated")
	_mbOnPaintBitUpdated = mbDLL.NewProc("mbOnPaintBitUpdated")
	_mbOnAcceleratedPaint = mbDLL.NewProc("mbOnAcceleratedPaint")
	_mbOnLoadUrlBegin = mbDLL.NewProc("mbOnLoadUrlBegin")
	_mbOnLoadUrlEnd = mbDLL.NewProc("mbOnLoadUrlEnd")
	_mbOnLoadUrlFail = mbDLL.NewProc("mbOnLoadUrlFail")
	_mbOnLoadUrlHeadersReceived = mbDLL.NewProc("mbOnLoadUrlHeadersReceived")
	_mbOnLoadUrlFinish = mbDLL.NewProc("mbOnLoadUrlFinish")
	_mbOnTitleChanged = mbDLL.NewProc("mbOnTitleChanged")
	_mbOnURLChanged = mbDLL.NewProc("mbOnURLChanged")
	_mbOnLoadingFinish = mbDLL.NewProc("mbOnLoadingFinish")
	_mbOnDownload = mbDLL.NewProc("mbOnDownload")
	_mbOnDownloadInBlinkThread = mbDLL.NewProc("mbOnDownloadInBlinkThread")
	_mbOnAlertBox = mbDLL.NewProc("mbOnAlertBox")
	_mbOnConfirmBox = mbDLL.NewProc("mbOnConfirmBox")
	_mbOnPromptBox = mbDLL.NewProc("mbOnPromptBox")
	_mbOnNetGetFavicon = mbDLL.NewProc("mbOnNetGetFavicon")
	_mbOnConsole = mbDLL.NewProc("mbOnConsole")
	_mbOnClose = mbDLL.NewProc("mbOnClose")
	_mbOnDestroy = mbDLL.NewProc("mbOnDestroy")
	_mbOnPrinting = mbDLL.NewProc("mbOnPrinting")
	_mbOnPluginList = mbDLL.NewProc("mbOnPluginList")
	_mbOnImageBufferToDataURL = mbDLL.NewProc("mbOnImageBufferToDataURL")
	_mbOnDidCreateScriptContext = mbDLL.NewProc("mbOnDidCreateScriptContext")
	_mbOnWillReleaseScriptContext = mbDLL.NewProc("mbOnWillReleaseScriptContext")
	_mbGoBack = mbDLL.NewProc("mbGoBack")
	_mbGoForward = mbDLL.NewProc("mbGoForward")
	_mbGoToOffset = mbDLL.NewProc("mbGoToOffset")
	_mbGoToIndex = mbDLL.NewProc("mbGoToIndex")
	_mbNavigateAtIndex = mbDLL.NewProc("mbNavigateAtIndex")
	_mbGetNavigateIndex = mbDLL.NewProc("mbGetNavigateIndex")
	_mbStopLoading = mbDLL.NewProc("mbStopLoading")
	_mbReload = mbDLL.NewProc("mbReload")
	_mbPerformCookieCommand = mbDLL.NewProc("mbPerformCookieCommand")
	_mbInsertCSSByFrame = mbDLL.NewProc("mbInsertCSSByFrame")
	_mbEditorSelectAll = mbDLL.NewProc("mbEditorSelectAll")
	_mbEditorUnSelect = mbDLL.NewProc("mbEditorUnSelect")
	_mbEditorCopy = mbDLL.NewProc("mbEditorCopy")
	_mbEditorCut = mbDLL.NewProc("mbEditorCut")
	_mbEditorPaste = mbDLL.NewProc("mbEditorPaste")
	_mbEditorDelete = mbDLL.NewProc("mbEditorDelete")
	_mbEditorUndo = mbDLL.NewProc("mbEditorUndo")
	_mbEditorRedo = mbDLL.NewProc("mbEditorRedo")
	_mbSetEditable = mbDLL.NewProc("mbSetEditable")
	_mbFireMouseEvent = mbDLL.NewProc("mbFireMouseEvent")
	_mbFireContextMenuEvent = mbDLL.NewProc("mbFireContextMenuEvent")
	_mbFireMouseWheelEvent = mbDLL.NewProc("mbFireMouseWheelEvent")
	_mbFireKeyUpEvent = mbDLL.NewProc("mbFireKeyUpEvent")
	_mbFireKeyDownEvent = mbDLL.NewProc("mbFireKeyDownEvent")
	_mbFireKeyPressEvent = mbDLL.NewProc("mbFireKeyPressEvent")
	_mbFireWindowsMessage = mbDLL.NewProc("mbFireWindowsMessage")
	_mbSetFocus = mbDLL.NewProc("mbSetFocus")
	_mbKillFocus = mbDLL.NewProc("mbKillFocus")
	_mbShowWindow = mbDLL.NewProc("mbShowWindow")
	_mbLoadURL = mbDLL.NewProc("mbLoadURL")
	_mbLoadHtmlWithBaseUrl = mbDLL.NewProc("mbLoadHtmlWithBaseUrl")
	_mbPostURL = mbDLL.NewProc("mbPostURL")
	_mbGetLockedViewDC = mbDLL.NewProc("mbGetLockedViewDC")
	_mbUnlockViewDC = mbDLL.NewProc("mbUnlockViewDC")
	_mbWake = mbDLL.NewProc("mbWake")
	_mbJsToV8Value = mbDLL.NewProc("mbJsToV8Value")
	_mbGetGlobalExecByFrame = mbDLL.NewProc("mbGetGlobalExecByFrame")
	_mbJsToDouble = mbDLL.NewProc("mbJsToDouble")
	_mbJsToBoolean = mbDLL.NewProc("mbJsToBoolean")
	_mbJsToString = mbDLL.NewProc("mbJsToString")
	_mbGetJsValueType = mbDLL.NewProc("mbGetJsValueType")
	_mbOnJsQuery = mbDLL.NewProc("mbOnJsQuery")
	_mbResponseQuery = mbDLL.NewProc("mbResponseQuery")
	_mbRunJs = mbDLL.NewProc("mbRunJs")
	_mbRunJsSync = mbDLL.NewProc("mbRunJsSync")
	_mbWebFrameGetMainFrame = mbDLL.NewProc("mbWebFrameGetMainFrame")
	_mbIsMainFrame = mbDLL.NewProc("mbIsMainFrame")
	_mbSetNodeJsEnable = mbDLL.NewProc("mbSetNodeJsEnable")
	_mbSetDeviceParameter = mbDLL.NewProc("mbSetDeviceParameter")
	_mbGetContentAsMarkup = mbDLL.NewProc("mbGetContentAsMarkup")
	_mbGetSource = mbDLL.NewProc("mbGetSource")
	_mbUtilSerializeToMHTML = mbDLL.NewProc("mbUtilSerializeToMHTML")
	_mbUtilCreateRequestCode = mbDLL.NewProc("mbUtilCreateRequestCode")
	_mbUtilIsRegistered = mbDLL.NewProc("mbUtilIsRegistered")
	_mbUtilPrint = mbDLL.NewProc("mbUtilPrint")
	_mbUtilBase64Encode = mbDLL.NewProc("mbUtilBase64Encode")
	_mbUtilBase64Decode = mbDLL.NewProc("mbUtilBase64Decode")
	_mbUtilDecodeURLEscape = mbDLL.NewProc("mbUtilDecodeURLEscape")
	_mbUtilEncodeURLEscape = mbDLL.NewProc("mbUtilEncodeURLEscape")
	_mbUtilCreateV8Snapshot = mbDLL.NewProc("mbUtilCreateV8Snapshot")
	_mbUtilPrintToPdf = mbDLL.NewProc("mbUtilPrintToPdf")
	_mbUtilPrintToBitmap = mbDLL.NewProc("mbUtilPrintToBitmap")
	_mbUtilScreenshot = mbDLL.NewProc("mbUtilScreenshot")
	_mbUtilsSilentPrint = mbDLL.NewProc("mbUtilsSilentPrint")
	_mbUtilSetDefaultPrinterSettings = mbDLL.NewProc("mbUtilSetDefaultPrinterSettings")
	_mbPopupDownloadMgr = mbDLL.NewProc("mbPopupDownloadMgr")
	_mbPopupDialogAndDownload = mbDLL.NewProc("mbPopupDialogAndDownload")
	_mbDownloadByPath = mbDLL.NewProc("mbDownloadByPath")
	_mbGetPdfPageData = mbDLL.NewProc("mbGetPdfPageData")
	_mbCreateMemBuf = mbDLL.NewProc("mbCreateMemBuf")
	_mbFreeMemBuf = mbDLL.NewProc("mbFreeMemBuf")
	_mbSetUserKeyValue = mbDLL.NewProc("mbSetUserKeyValue")
	_mbGetUserKeyValue = mbDLL.NewProc("mbGetUserKeyValue")
	_mbPluginListBuilderAddPlugin = mbDLL.NewProc("mbPluginListBuilderAddPlugin")
	_mbPluginListBuilderAddMediaTypeToLastPlugin = mbDLL.NewProc("mbPluginListBuilderAddMediaTypeToLastPlugin")
	_mbPluginListBuilderAddFileExtensionToLastMediaType = mbDLL.NewProc("mbPluginListBuilderAddFileExtensionToLastMediaType")
	_mbGetBlinkMainThreadIsolate = mbDLL.NewProc("mbGetBlinkMainThreadIsolate")
	_mbWebFrameGetMainWorldScriptContext = mbDLL.NewProc("mbWebFrameGetMainWorldScriptContext")
	_mbEnableHighDPISupport = mbDLL.NewProc("mbEnableHighDPISupport")
	_mbRunMessageLoop = mbDLL.NewProc("mbRunMessageLoop")
	_mbGetContentWidth = mbDLL.NewProc("mbGetContentWidth")
	_mbGetContentHeight = mbDLL.NewProc("mbGetContentHeight")
	_mbGetWebViewForCurrentContext = mbDLL.NewProc("mbGetWebViewForCurrentContext")
	_mbRegisterEmbedderCustomElement = mbDLL.NewProc("mbRegisterEmbedderCustomElement")
	_mbOnNodeCreateProcess = mbDLL.NewProc("mbOnNodeCreateProcess")
	_mbOnThreadIdle = mbDLL.NewProc("mbOnThreadIdle")
	_mbGetProcAddr = mbDLL.NewProc("mbGetProcAddr")
	_mbSetMbDllPath = mbDLL.NewProc("mbSetMbDllPath")
	_mbSetMbMainDllPath = mbDLL.NewProc("mbSetMbMainDllPath")
	_mbFillFuncPtr = mbDLL.NewProc("mbFillFuncPtr")
}

func Init(settings *Settings) {
	_mbInit.Call(uintptr(unsafe.Pointer(settings)))
}

func Uninit() {
	_mbUninit.Call()
}

func CreateInitSettings() *Settings {
	r1, _, _ := _mbCreateInitSettings.Call()
	return (*Settings)(unsafe.Pointer(r1))
}

func SetInitSettings(settings *Settings, name, value string) {
	_mbSetInitSettings.Call(uintptr(unsafe.Pointer(settings)), StrToPtr(name), StrToPtr(value))
}

func CreateWebView() WebView {
	r1, _, _ := _mbCreateWebView.Call()
	return WebView(r1)
}
func DestroyWebView(webview WebView) {
	_mbDestroyWebView.Call(uintptr(webview))
}

func CreateWebWindow(typ WindowType, parent Handle, x, y, width, height int) WebView {
	r1, _, _ := _mbCreateWebWindow.Call(uintptr(typ), uintptr(parent), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return WebView(r1)
}

func CreateWebCustomWindow(parent Handle, style, styleEx uint32, x, y, width, height int) WebView {
	r1, _, _ := _mbCreateWebCustomWindow.Call(uintptr(parent), uintptr(style), uintptr(styleEx), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return WebView(r1)
}

func MoveWindow(webview WebView, x, y, width, height int) {
	_mbMoveWindow.Call(uintptr(webview), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func MoveToCenter(webview WebView) {
	_mbMoveToCenter.Call(uintptr(webview))
}

func SetAutoDrawToHwnd(webview WebView, b bool) {
	_mbSetAutoDrawToHwnd.Call(uintptr(webview), BoolToPtr(b))
}

func GetCaretRect(webview WebView, r *Rect) {
	_mbGetCaretRect.Call(uintptr(webview), uintptr(unsafe.Pointer(r)))
}

func SetAudioMuted(webview WebView, b bool) {
	_mbSetAudioMuted.Call(uintptr(webview), BoolToPtr(b))
}

func IsAudioMuted(webview WebView) bool {
	r1, _, _ := _mbIsAudioMuted.Call(uintptr(webview))
	return 0 != r1
}

func CreateString(str string, length uintptr) StringPtr {
	r1, _, _ := _mbCreateString.Call(StrToPtr(str), length)
	return StringPtr(r1)
}

func CreateStringWithoutNullTermination(str string, length uintptr) StringPtr {
	r1, _, _ := _mbCreateStringWithoutNullTermination.Call(StrToPtr(str), length)
	return StringPtr(r1)
}

func DeleteString(str StringPtr) {
	_mbDeleteString.Call(uintptr(str))
}

func GetStringLen(str StringPtr) uintptr {
	r1, _, _ := _mbGetStringLen.Call(uintptr(str))
	return r1
}

func GetString(str StringPtr) string {
	r1, _, _ := _mbGetString.Call(uintptr(str))
	return StrFromPtr(r1)
}
func SetProxy(webView WebView, proxy *Proxy) {
	_mbSetProxy.Call(uintptr(webView), uintptr(unsafe.Pointer(proxy)))
}

func SetDebugConfig(webView WebView, debugString, param string) {
	_mbSetDebugConfig.Call(uintptr(webView), StrToPtr(debugString), StrToPtr(param))
}

func NetSetData(jobPtr NetJob, buf []byte) {
	_mbNetSetData.Call(uintptr(jobPtr), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func NetHookRequest(jobPtr NetJob) {
	_mbNetHookRequest.Call(uintptr(jobPtr))
}

func NetChangeRequestUrl(jobPtr NetJob, url string) {
	_mbNetChangeRequestUrl.Call(uintptr(jobPtr), StrToPtr(url))
}

func NetContinueJob(jobPtr NetJob) {
	_mbNetContinueJob.Call(uintptr(jobPtr))
}

func NetGetRawHttpHeadInBlinkThread(jobPtr NetJob) *Slist {
	r1, _, _ := _mbNetGetRawHttpHeadInBlinkThread.Call(uintptr(jobPtr))
	return (*Slist)(unsafe.Pointer(r1))
}

func NetGetRawResponseHeadInBlinkThread(jobPtr NetJob) *Slist {
	r1, _, _ := _mbNetGetRawResponseHeadInBlinkThread.Call(uintptr(jobPtr))
	return (*Slist)(unsafe.Pointer(r1))
}

func NetHoldJobToAsynCommit(jobPtr NetJob) bool {
	r1, _, _ := _mbNetHoldJobToAsynCommit.Call(uintptr(jobPtr))
	return BoolFromPtr(r1)
}

func NetCancelRequest(jobPtr NetJob) {
	_mbNetCancelRequest.Call(uintptr(jobPtr))
}

func NetOnResponse(webviewHandle WebView, callback NetResponseCallback) {
	_mbNetOnResponse.Call(uintptr(webviewHandle), CallbackToPtr(callback), 0)
}

func NetSetWebsocketCallback(webview WebView, callbacks *WebsocketHookCallbacks) {

	type mbWebsocketHookCallbacks struct {
		onWillConnect uintptr
		onConnected   uintptr
		onReceive     uintptr
		onSend        uintptr
		onError       uintptr
	}

	cbs := mbWebsocketHookCallbacks{
		onWillConnect: CallbackToPtr(callbacks.OnWillConnect),
		onConnected:   CallbackToPtr(callbacks.OnConnected),
		onReceive:     CallbackToPtr(callbacks.OnReceive),
		onSend:        CallbackToPtr(callbacks.OnSend),
		onError:       CallbackToPtr(callbacks.OnError),
	}

	_mbNetSetWebsocketCallback.Call(uintptr(webview), uintptr(unsafe.Pointer(&cbs)), 0)
}

func NetSendWsText(channel WebSocketChannel, buf []byte) {
	_mbNetSendWsText.Call(uintptr(channel), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func NetSendWsBlob(channel WebSocketChannel, buf []byte) {
	_mbNetSendWsBlob.Call(uintptr(channel), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func NetEnableResPacket(webviewHandle WebView, pathName string) {
	_mbNetEnableResPacket.Call(uintptr(webviewHandle), StrToWPtr(pathName))
}

func NetGetPostBody(jobPtr NetJob) *PostBodyElements {
	r1, _, _ := _mbNetGetPostBody.Call(uintptr(jobPtr))
	return (*PostBodyElements)(unsafe.Pointer(r1))
}

func NetCreatePostBodyElements(webView WebView, length uintptr) *PostBodyElements {
	r1, _, _ := _mbNetCreatePostBodyElements.Call(uintptr(webView), length)
	return (*PostBodyElements)(unsafe.Pointer(r1))
}

func NetFreePostBodyElements(elements *PostBodyElements) {
	_mbNetFreePostBodyElements.Call(uintptr(unsafe.Pointer(elements)))
}

func NetCreatePostBodyElement(webView WebView) *PostBodyElement {
	r1, _, _ := _mbNetCreatePostBodyElement.Call(uintptr(webView))
	return (*PostBodyElement)(unsafe.Pointer(r1))
}

func NetFreePostBodyElement(element *PostBodyElement) {
	_mbNetFreePostBodyElement.Call(uintptr(unsafe.Pointer(element)))
}

func NetCreateWebUrlRequest(url, method, mime string) *WebUrlRequest {
	r1, _, _ := _mbNetCreateWebUrlRequest.Call(StrToPtr(url), StrToPtr(method), StrToPtr(mime))
	return (*WebUrlRequest)(unsafe.Pointer(r1))
}

func NetAddHTTPHeaderFieldToUrlRequest(request *WebUrlRequest, name, value string) {
	_mbNetAddHTTPHeaderFieldToUrlRequest.Call(uintptr(unsafe.Pointer(request)), StrToPtr(name), StrToPtr(value))
}

func NetStartUrlRequest(webView WebView, request *WebUrlRequest, callbacks *UrlRequestCallbacks) int {
	r1, _, _ := _mbNetStartUrlRequest.Call(uintptr(webView), uintptr(unsafe.Pointer(request)), 0, CallbackToPtr(callbacks))
	return int(r1)
}

func NetGetHttpStatusCode(response *WebUrlResponse) int {
	r1, _, _ := _mbNetGetHttpStatusCode.Call(uintptr(unsafe.Pointer(response)))
	return int(r1)
}

func NetGetRequestMethod(jobPtr NetJob) RequestType {
	r1, _, _ := _mbNetGetRequestMethod.Call(uintptr(jobPtr))
	return RequestType(r1)
}

func NetGetExpectedContentLength(response *WebUrlResponse) int64 {
	r1, _, _ := _mbNetGetExpectedContentLength.Call(uintptr(unsafe.Pointer(response)))
	return int64(r1)
}

func NetGetResponseUrl(response *WebUrlResponse) string {
	r1, _, _ := _mbNetGetResponseUrl.Call(uintptr(unsafe.Pointer(response)))
	return StrFromPtr(r1)
}

func NetCancelWebUrlRequest(requestId int) {
	_mbNetCancelWebUrlRequest.Call(uintptr(requestId))
}

func SetViewProxy(webView WebView, proxy *Proxy) {
	_mbSetViewProxy.Call(uintptr(webView), uintptr(unsafe.Pointer(proxy)))
}

func NetSetMIMEType(jobPtr NetJob, typeStr string) {
	_mbNetSetMIMEType.Call(uintptr(jobPtr), StrToPtr(typeStr))
}

func NetGetMIMEType(jobPtr NetJob) string {
	r1, _, _ := _mbNetGetMIMEType.Call(uintptr(jobPtr))
	return StrFromPtr(r1)
}

func NetGetHTTPHeaderField(job NetJob, key string, fromRequestOrResponse bool) string {
	r1, _, _ := _mbNetGetHTTPHeaderField.Call(uintptr(job), StrToPtr(key), BoolToPtr(fromRequestOrResponse))
	return StrFromPtr(r1)
}

func NetGetReferrer(jobPtr NetJob) string {
	r1, _, _ := _mbNetGetReferrer.Call(uintptr(jobPtr))
	return StrFromPtr(r1)
}
func NetSetHTTPHeaderField(jobPtr NetJob, key, value string, response bool) {
	_mbNetSetHTTPHeaderField.Call(uintptr(jobPtr), StrToPtr(key), StrToPtr(value), BoolToPtr(response))
}

func SetMouseEnabled(webView WebView, b bool) {
	_mbSetMouseEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetTouchEnabled(webView WebView, b bool) {
	_mbSetTouchEnabled.Call(uintptr(webView), BoolToPtr(b))
}
func SetSystemTouchEnabled(webView WebView, b bool) {
	_mbSetSystemTouchEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetContextMenuEnabled(webView WebView, b bool) {
	_mbSetContextMenuEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetNavigationToNewWindowEnabled(webView WebView, b bool) {
	_mbSetNavigationToNewWindowEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetHeadlessEnabled(webView WebView, b bool) {
	_mbSetHeadlessEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetDragDropEnable(webView WebView, b bool) {
	_mbSetDragDropEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetDragEnable(webView WebView, b bool) {
	_mbSetDragEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetContextMenuItemShow(webView WebView, item MenuItemId, isShow bool) {
	_mbSetContextMenuItemShow.Call(uintptr(webView), uintptr(item), BoolToPtr(isShow))
}

func SetHandle(webView WebView, wnd HWND) {
	_mbSetHandle.Call(uintptr(webView), uintptr(wnd))
}

func SetHandleOffset(webView WebView, x, y int) {
	_mbSetHandleOffset.Call(uintptr(webView), uintptr(x), uintptr(y))
}

func GetHostHWND(webView WebView) HWND {
	r1, _, _ := _mbGetHostHWND.Call(uintptr(webView))
	return HWND(r1)
}

func SetTransparent(webviewHandle WebView, transparent bool) {
	_mbSetTransparent.Call(uintptr(webviewHandle), BoolToPtr(transparent))
}

func SetViewSettings(webviewHandle WebView, settings *ViewSettings) {
	_mbSetViewSettings.Call(uintptr(webviewHandle), uintptr(unsafe.Pointer(settings)))
}

func SetCspCheckEnable(webView WebView, b bool) {
	_mbSetCspCheckEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetNpapiPluginsEnabled(webView WebView, b bool) {
	_mbSetNpapiPluginsEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetMemoryCacheEnable(webView WebView, b bool) {
	_mbSetMemoryCacheEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetCookie(webView WebView, url, cookie string) {
	_mbSetCookie.Call(uintptr(webView), StrToPtr(url), StrToPtr(cookie))
}

func SetCookieEnabled(webView WebView, enable bool) {
	_mbSetCookieEnabled.Call(uintptr(webView), BoolToPtr(enable))
}

func SetCookieJarPath(webView WebView, path string) {
	_mbSetCookieJarPath.Call(uintptr(webView), StrToPtr(path))
}

func SetCookieJarFullPath(webView WebView, path string) {
	_mbSetCookieJarFullPath.Call(uintptr(webView), StrToPtr(path))
}

func SetLocalStorageFullPath(webView WebView, path string) {
	_mbSetLocalStorageFullPath.Call(uintptr(webView), StrToWPtr(path))
}

func GetTitle(webView WebView) string {
	r1, _, _ := _mbGetTitle.Call(uintptr(webView))
	return StrFromPtr(r1)
}

func SetWindowTitle(webView WebView, title string) {
	_mbSetWindowTitle.Call(uintptr(webView), StrToPtr(title))
}

func SetWindowTitleW(webView WebView, title string) {
	_mbSetWindowTitleW.Call(uintptr(webView), StrToPtr(title))
}

func GetUrl(webView WebView) string {
	r1, _, _ := _mbGetUrl.Call(uintptr(webView))
	return StrFromPtr(r1)
}

func GetCursorInfoType(webView WebView) int {
	r1, _, _ := _mbGetCursorInfoType.Call(uintptr(webView))
	return int(r1)
}

func AddPluginDirectory(webView WebView, path string) {
	_mbAddPluginDirectory.Call(uintptr(webView), StrToPtr(path))
}

func SetUserAgent(webView WebView, userAgent string) {
	_mbSetUserAgent.Call(uintptr(webView), StrToPtr(userAgent))
}

func SetZoomFactor(webView WebView, factor float32) {
	_mbSetZoomFactor.Call(uintptr(webView), F32ToPtr(factor))
}

func GetZoomFactor(webView WebView) float32 {
	_, r2, _ := _mbGetZoomFactor.Call(uintptr(webView))
	return F32FromPtr(r2)
}
func SetDiskCacheEnabled(webView WebView, enable bool) {
	_mbSetDiskCacheEnabled.Call(uintptr(webView), BoolToPtr(enable))
}

func SetDiskCachePath(webView WebView, path string) {
	_mbSetDiskCachePath.Call(uintptr(webView), StrToPtr(path))
}

func SetDiskCacheLimit(webView WebView, limit uint64) {
	_mbSetDiskCacheLimit.Call(uintptr(webView), uintptr(limit))
}

func SetDiskCacheLimitDisk(webView WebView, limit uint64) {
	_mbSetDiskCacheLimitDisk.Call(uintptr(webView), uintptr(limit))
}

func SetDiskCacheLevel(webView WebView, level int) {
	_mbSetDiskCacheLevel.Call(uintptr(webView), uintptr(level))
}

func SetResourceGc(webView WebView, intervalSec int) {
	_mbSetResourceGc.Call(uintptr(webView), uintptr(intervalSec))
}

func CanGoForward(webView WebView, callback CanGoBackForwardCallback) {
	_mbCanGoForward.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func CanGoBack(webView WebView, callback CanGoBackForwardCallback) {
	_mbCanGoBack.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func GetCookie(webView WebView, callback GetCookieCallback) {
	_mbGetCookie.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func GetCookieOnBlinkThread(webView WebView) string {
	r1, _, _ := _mbGetCookieOnBlinkThread.Call(uintptr(webView))
	return StrFromPtr(r1)
}

func ClearCookie(webView WebView) {
	_mbClearCookie.Call(uintptr(webView))
}

func Resize(webView WebView, w, h int) {
	_mbResize.Call(uintptr(webView), uintptr(w), uintptr(h))
}

func OnNavigation(webView WebView, callback NavigationCallback) {
	_mbOnNavigation.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnNavigationSync(webView WebView, callback NavigationCallback) {
	_mbOnNavigationSync.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnCreateView(webView WebView, callback CreateViewCallback) {
	_mbOnCreateView.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnDocumentReady(webView WebView, callback DocumentReadyCallback) {
	_mbOnDocumentReady.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnDocumentReadyInBlinkThread(webView WebView, callback DocumentReadyCallback) {
	_mbOnDocumentReadyInBlinkThread.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnPaintUpdated(webView WebView, callback PaintUpdatedCallback) {
	_mbOnPaintUpdated.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnPaintBitUpdated(webView WebView, callback PaintBitUpdatedCallback) {
	_mbOnPaintBitUpdated.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnAcceleratedPaint(webView WebView, callback AcceleratedPaintCallback) {
	_mbOnAcceleratedPaint.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlBegin(webView WebView, callback LoadUrlBeginCallback) {
	_mbOnLoadUrlBegin.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlEnd(webView WebView, callback LoadUrlEndCallback) {
	_mbOnLoadUrlEnd.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlFail(webView WebView, callback LoadUrlFailCallback) {
	_mbOnLoadUrlFail.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlHeadersReceived(webView WebView, callback LoadUrlHeadersReceivedCallback) {
	_mbOnLoadUrlHeadersReceived.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlFinish(webView WebView, callback LoadUrlFinishCallback) {
	_mbOnLoadUrlFinish.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnTitleChanged(webView WebView, callback TitleChangedCallback) {
	_mbOnTitleChanged.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnURLChanged(webView WebView, callback URLChangedCallback) {
	_mbOnURLChanged.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnLoadingFinish(webView WebView, callback LoadingFinishCallback) {
	_mbOnLoadingFinish.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnDownload(webView WebView, callback DownloadCallback) {
	_mbOnDownload.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnDownloadInBlinkThread(webView WebView, callback DownloadInBlinkThreadCallback) {
	_mbOnDownloadInBlinkThread.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnAlertBox(webView WebView, callback AlertBoxCallback) {
	_mbOnAlertBox.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnConfirmBox(webView WebView, callback ConfirmBoxCallback) {
	_mbOnConfirmBox.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnPromptBox(webView WebView, callback PromptBoxCallback) {
	_mbOnPromptBox.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnNetGetFavicon(webView WebView, callback NetGetFaviconCallback) {
	_mbOnNetGetFavicon.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnConsole(webView WebView, callback ConsoleCallback) {
	_mbOnConsole.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnClose(webView WebView, callback CloseCallback) {
	_mbOnClose.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnDestroy(webView WebView, callback DestroyCallback) {
	_mbOnDestroy.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnPrinting(webView WebView, callback PrintingCallback) {
	_mbOnPrinting.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnPluginList(webView WebView, callback GetPluginListCallback) {
	_mbOnPluginList.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnImageBufferToDataURL(webView WebView, callback ImageBufferToDataURLCallback) {
	_mbOnImageBufferToDataURL.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnDidCreateScriptContext(webView WebView, callback DidCreateScriptContextCallback) {
	_mbOnDidCreateScriptContext.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func OnWillReleaseScriptContext(webView WebView, callback WillReleaseScriptContextCallback) {
	_mbOnWillReleaseScriptContext.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func GoBack(webView WebView) {
	_mbGoBack.Call(uintptr(webView))
}
func GoForward(webView WebView) {
	_mbGoForward.Call(uintptr(webView))
}
func GoToOffset(webView WebView, offset int) {
	_mbGoToOffset.Call(uintptr(webView), uintptr(offset))
}
func GoToIndex(webView WebView, index int) {
	_mbGoToIndex.Call(uintptr(webView), uintptr(index))
}
func NavigateAtIndex(webView WebView, index int) {
	_mbNavigateAtIndex.Call(uintptr(webView), uintptr(index))
}
func GetNavigateIndex(webView WebView) int {
	r1, _, _ := _mbGetNavigateIndex.Call(uintptr(webView))
	return int(r1)
}
func StopLoading(webView WebView) {
	_mbStopLoading.Call(uintptr(webView))
}
func Reload(webView WebView) {
	_mbReload.Call(uintptr(webView))
}
func PerformCookieCommand(webView WebView, command CookieCommand) {
	_mbPerformCookieCommand.Call(uintptr(webView), uintptr(command))
}
func InsertCSSByFrame(webView WebView, frameId WebFrameHandle, cssText string) {
	_mbInsertCSSByFrame.Call(uintptr(webView), uintptr(frameId), StrToPtr(cssText))
}
func EditorSelectAll(webView WebView) {
	_mbEditorSelectAll.Call(uintptr(webView))
}
func EditorUnSelect(webView WebView) {
	_mbEditorUnSelect.Call(uintptr(webView))
}
func EditorCopy(webView WebView) {
	_mbEditorCopy.Call(uintptr(webView))
}
func EditorCut(webView WebView) {
	_mbEditorCut.Call(uintptr(webView))
}
func EditorPaste(webView WebView) {
	_mbEditorPaste.Call(uintptr(webView))
}
func EditorDelete(webView WebView) {
	_mbEditorDelete.Call(uintptr(webView))
}
func EditorUndo(webView WebView) {
	_mbEditorUndo.Call(uintptr(webView))
}
func EditorRedo(webView WebView) {
	_mbEditorRedo.Call(uintptr(webView))
}
func SetEditable(webView WebView, editable bool) {
	_mbSetEditable.Call(uintptr(webView), BoolToPtr(editable))
}
func FireMouseEvent(webView WebView, message uint, x, y int, flags uint) bool {
	r1, _, _ := _mbFireMouseEvent.Call(uintptr(webView), uintptr(message), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r1)
}
func FireContextMenuEvent(webView WebView, x, y int, flags uint) bool {
	r1, _, _ := _mbFireContextMenuEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r1)
}
func FireMouseWheelEvent(webView WebView, x, y, delta int, flags uint) bool {
	r1, _, _ := _mbFireMouseWheelEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
	return BoolFromPtr(r1)
}
func FireKeyUpEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r1, _, _ := _mbFireKeyUpEvent.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r1)
}
func FireKeyDownEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r1, _, _ := _mbFireKeyDownEvent.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r1)
}
func FireKeyPressEvent(webView WebView, charCode, flags uint, systemKey bool) bool {
	r1, _, _ := _mbFireKeyPressEvent.Call(uintptr(webView), uintptr(charCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r1)
}
func FireWindowsMessage(webView WebView, hWnd HWND, message uint, wParam WPARAM, lParam LPARAM, result *LRESULT) bool {
	r1, _, _ := _mbFireWindowsMessage.Call(uintptr(webView), uintptr(hWnd), uintptr(message), uintptr(wParam), uintptr(lParam), uintptr(unsafe.Pointer(result)))
	return BoolFromPtr(r1)
}
func SetFocus(webView WebView) {
	_mbSetFocus.Call(uintptr(webView))
}
func KillFocus(webView WebView) {
	_mbKillFocus.Call(uintptr(webView))
}
func ShowWindow(webView WebView, show bool) {
	_mbShowWindow.Call(uintptr(webView), BoolToPtr(show))
}

func LoadURL(webView WebView, url string) {
	_mbLoadURL.Call(uintptr(webView), StrToPtr(url))
}

func LoadHtmlWithBaseUrl(webView WebView, html string, baseUrl string) {
	_mbLoadHtmlWithBaseUrl.Call(uintptr(webView), StrToPtr(html), StrToPtr(baseUrl))
}
func PostURL(webView WebView, url string, postData []byte) {
	_mbPostURL.Call(uintptr(webView), StrToPtr(url), uintptr(unsafe.Pointer(&postData[0])), uintptr(len(postData)))
}
func GetLockedViewDC(webView WebView) HDC {
	r1, _, _ := _mbGetLockedViewDC.Call(uintptr(webView))
	return HDC(r1)
}
func UnlockViewDC(webView WebView) {
	_mbUnlockViewDC.Call(uintptr(webView))
}
func Wake(webView WebView) {
	_mbWake.Call(uintptr(webView))
}
func JsToV8Value(es JsExecState, v JsValue) uintptr {
	r1, _, _ := _mbJsToV8Value.Call(uintptr(es), uintptr(v))
	return r1
}
func GetGlobalExecByFrame(webView WebView, frameId WebFrameHandle) JsExecState {
	r1, _, _ := _mbGetGlobalExecByFrame.Call(uintptr(webView), uintptr(frameId))
	return JsExecState(r1)
}
func JsToDouble(es JsExecState, v JsValue) float64 {
	_, r2, _ := _mbJsToDouble.Call(uintptr(es), uintptr(v))
	return F64FromPtr(r2)
}
func JsToBoolean(es JsExecState, v JsValue) bool {
	r1, _, _ := _mbJsToBoolean.Call(uintptr(es), uintptr(v))
	return BoolFromPtr(r1)
}
func JsToString(es JsExecState, v JsValue) string {
	r1, _, _ := _mbJsToString.Call(uintptr(es), uintptr(v))
	return StrFromPtr(r1)
}
func GetJsValueType(es JsExecState, v JsValue) JsType {
	r1, _, _ := _mbGetJsValueType.Call(uintptr(es), uintptr(v))
	return JsType(r1)
}
func OnJsQuery(webView WebView, callback JsQueryCallback) {
	_mbOnJsQuery.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func ResponseQuery(webView WebView, queryId int64, customMsg int, response string) {
	_mbResponseQuery.Call(uintptr(webView), uintptr(queryId), uintptr(customMsg), StrToPtr(response))
}
func RunJs(webView WebView, frameId WebFrameHandle, script string, isInClosure bool, callback RunJsCallback, unuse uintptr) {
	_mbRunJs.Call(uintptr(webView), uintptr(frameId), StrToPtr(script), BoolToPtr(isInClosure), CallbackToPtr(callback), 0, unuse)
}
func RunJsSync(webView WebView, frameId WebFrameHandle, script string, isInClosure bool) JsValue {
	r1, _, _ := _mbRunJsSync.Call(uintptr(webView), uintptr(frameId), StrToPtr(script), BoolToPtr(isInClosure))
	return JsValue(r1)
}
func WebFrameGetMainFrame(webView WebView) WebFrameHandle {
	r1, _, _ := _mbWebFrameGetMainFrame.Call(uintptr(webView))
	return WebFrameHandle(r1)
}
func IsMainFrame(webView WebView, frameId WebFrameHandle) bool {
	r1, _, _ := _mbIsMainFrame.Call(uintptr(webView), uintptr(frameId))
	return BoolFromPtr(r1)
}
func SetNodeJsEnable(webView WebView, enable bool) {
	_mbSetNodeJsEnable.Call(uintptr(webView), BoolToPtr(enable))
}
func SetDeviceParameter(webView WebView, device, paramStr string, paramInt int, paramFloat float32) {
	_mbSetDeviceParameter.Call(uintptr(webView), StrToPtr(device), StrToPtr(paramStr), uintptr(paramInt), F32ToPtr(paramFloat))
}
func GetContentAsMarkup(webView WebView, callback GetContentAsMarkupCallback, frameId WebFrameHandle) {
	_mbGetContentAsMarkup.Call(uintptr(webView), CallbackToPtr(callback), 0, uintptr(frameId))
}
func GetSource(webView WebView, callback GetSourceCallback) {
	_mbGetSource.Call(uintptr(webView), CallbackToPtr(callback))
}
func UtilSerializeToMHTML(webView WebView, callback GetSourceCallback) {
	_mbUtilSerializeToMHTML.Call(uintptr(webView), CallbackToPtr(callback))
}
func UtilCreateRequestCode(registerInfo string) string {
	r1, _, _ := _mbUtilCreateRequestCode.Call(StrToPtr(registerInfo))
	return StrFromPtr(r1)
}
func UtilIsRegistered(defaultPath string) bool {
	r1, _, _ := _mbUtilIsRegistered.Call(StrToPtr(defaultPath))
	return BoolFromPtr(r1)
}
func UtilPrint(webView WebView, frameId WebFrameHandle, printParams *PrintSettings) bool {
	r1, _, _ := _mbUtilPrint.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(printParams)))
	return BoolFromPtr(r1)
}
func UtilBase64Encode(str string) string {
	r1, _, _ := _mbUtilBase64Encode.Call(StrToPtr(str))
	return StrFromPtr(r1)
}
func UtilBase64Decode(str string) string {
	r1, _, _ := _mbUtilBase64Decode.Call(StrToPtr(str))
	return StrFromPtr(r1)
}
func UtilDecodeURLEscape(url string) string {
	r1, _, _ := _mbUtilDecodeURLEscape.Call(StrToPtr(url))
	return StrFromPtr(r1)
}
func UtilEncodeURLEscape(url string) string {
	r1, _, _ := _mbUtilEncodeURLEscape.Call(StrToPtr(url))
	return StrFromPtr(r1)
}
func UtilCreateV8Snapshot(str string) *MemBuf {
	r1, _, _ := _mbUtilCreateV8Snapshot.Call(StrToPtr(str))
	return (*MemBuf)(unsafe.Pointer(r1))
}
func UtilPrintToPdf(webView WebView, frameId WebFrameHandle, settings *PrintSettings, callback PrintPdfDataCallback) {
	_mbUtilPrintToPdf.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)), CallbackToPtr(callback), 0)
}
func UtilPrintToBitmap(webView WebView, frameId WebFrameHandle, settings *ScreenshotSettings, callback PrintBitmapCallback) {
	_mbUtilPrintToBitmap.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)), CallbackToPtr(callback), 0)
}
func UtilScreenshot(webView WebView, settings *ScreenshotSettings, callback OnScreenshotCallback) {
	_mbUtilScreenshot.Call(uintptr(webView), uintptr(unsafe.Pointer(settings)), CallbackToPtr(callback), 0)
}
func UtilsSilentPrint(webView WebView, settings string) bool {
	r1, _, _ := _mbUtilsSilentPrint.Call(uintptr(webView), StrToPtr(settings))
	return BoolFromPtr(r1)
}
func UtilSetDefaultPrinterSettings(webView WebView, setting *DefaultPrinterSettings) {
	_mbUtilSetDefaultPrinterSettings.Call(uintptr(webView), uintptr(unsafe.Pointer(setting)))
}
func PopupDownloadMgr(webView WebView, url string, downloadJob uintptr) bool {
	r1, _, _ := _mbPopupDownloadMgr.Call(uintptr(webView), StrToPtr(url), downloadJob)
	return BoolFromPtr(r1)
}
func PopupDialogAndDownload(webView WebView, dialogOpt *DialogOptions, contentLength uint64, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	r1, _, _ := _mbPopupDialogAndDownload.Call(uintptr(webView), uintptr(unsafe.Pointer(dialogOpt)), uintptr(contentLength), StrToPtr(url), StrToPtr(mime), StrToPtr(disposition), uintptr(job), uintptr(unsafe.Pointer(dataBind)), uintptr(unsafe.Pointer(callbackBind)))
	return DownloadOpt(r1)
}
func DownloadByPath(webView WebView, downloadOptions *DownloadOptions, path string, contentLength uint64, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	r1, _, _ := _mbDownloadByPath.Call(uintptr(webView), uintptr(unsafe.Pointer(downloadOptions)), StrToPtr(path), uintptr(contentLength), StrToPtr(url), StrToPtr(mime), StrToPtr(disposition), uintptr(job), uintptr(unsafe.Pointer(dataBind)), uintptr(unsafe.Pointer(callbackBind)))
	return DownloadOpt(r1)
}
func GetPdfPageData(webView WebView, callback OnGetPdfPageDataCallback) {
	_mbGetPdfPageData.Call(uintptr(webView), CallbackToPtr(callback), 0)
}
func CreateMemBuf(webView WebView, buf []byte) *MemBuf {
	r1, _, _ := _mbCreateMemBuf.Call(uintptr(webView), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return (*MemBuf)(unsafe.Pointer(r1))
}
func FreeMemBuf(buf *MemBuf) {
	_mbFreeMemBuf.Call(uintptr(unsafe.Pointer(buf)))
}
func SetUserKeyValue(webView WebView, key string, value uintptr) {
	_mbSetUserKeyValue.Call(uintptr(webView), StrToPtr(key), value)
}
func GetUserKeyValue(webView WebView, key string) uintptr {
	r1, _, _ := _mbGetUserKeyValue.Call(uintptr(webView), StrToPtr(key))
	return r1
}
func PluginListBuilderAddPlugin(builder uintptr, name, description, fileName string) {
	_mbPluginListBuilderAddPlugin.Call(builder, StrToPtr(name), StrToPtr(description), StrToPtr(fileName))
}
func PluginListBuilderAddMediaTypeToLastPlugin(builder uintptr, name, description string) {
	_mbPluginListBuilderAddMediaTypeToLastPlugin.Call(builder, StrToPtr(name), StrToPtr(description))
}
func PluginListBuilderAddFileExtensionToLastMediaType(builder uintptr, fileExtension string) {
	_mbPluginListBuilderAddFileExtensionToLastMediaType.Call(builder, StrToPtr(fileExtension))
}
func GetBlinkMainThreadIsolate() V8Isolate {
	r1, _, _ := _mbGetBlinkMainThreadIsolate.Call()
	return V8Isolate(r1)
}
func WebFrameGetMainWorldScriptContext(webView WebView, frameId WebFrameHandle, contextOut *V8ContextPtr) {
	_mbWebFrameGetMainWorldScriptContext.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(contextOut)))
}
func EnableHighDPISupport() {
	_mbEnableHighDPISupport.Call()
}
func RunMessageLoop() {
	_mbRunMessageLoop.Call()
}
func GetContentWidth(webView WebView) int {
	r1, _, _ := _mbGetContentWidth.Call(uintptr(webView))
	return int(r1)
}
func GetContentHeight(webView WebView) int {
	r1, _, _ := _mbGetContentHeight.Call(uintptr(webView))
	return int(r1)
}
func GetWebViewForCurrentContext() WebView {
	r1, _, _ := _mbGetWebViewForCurrentContext.Call()
	return WebView(r1)
}
func RegisterEmbedderCustomElement(webviewHandle WebView, frameId WebFrameHandle, name string, options, outResult uintptr) bool {
	r1, _, _ := _mbRegisterEmbedderCustomElement.Call(uintptr(webviewHandle), uintptr(frameId), StrToPtr(name), options, outResult)
	return BoolFromPtr(r1)
}
func OnNodeCreateProcess(webviewHandle WebView, callback NodeOnCreateProcessCallback) {
	_mbOnNodeCreateProcess.Call(uintptr(webviewHandle), CallbackToPtr(callback), 0)
}
func OnThreadIdle(callback ThreadCallback, param1, param2 uintptr) {
	_mbOnThreadIdle.Call(CallbackToPtr(callback), param1, param2)
}
func GetProcAddr(name string) uintptr {
	r1, _, _ := _mbGetProcAddr.Call(StrToPtr(name))
	return r1
}
func SetMbDllPath(dllPath string) {
	_mbSetMbDllPath.Call(StrToPtr(dllPath))
}
func SetMbMainDllPath(dllPath string) {
	_mbSetMbMainDllPath.Call(StrToPtr(dllPath))
}
func FillFuncPtr() {
	_mbFillFuncPtr.Call()
}
