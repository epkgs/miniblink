//go:build windows
// +build windows

package wke

import (
	"fmt"
	"unsafe"

	"github.com/epkgs/miniblink/internal/lib"
	"golang.org/x/sys/windows"
)

var wkeDLL *windows.LazyDLL

var _wkeJsNativeFunction *windows.LazyProc
var _wkeShutdown *windows.LazyProc
var _wkeShutdownForDebug *windows.LazyProc
var _wkeVersion *windows.LazyProc
var _wkeVersionString *windows.LazyProc
var _wkeGC *windows.LazyProc
var _wkeSetResourceGc *windows.LazyProc
var _wkeSetFileSystem *windows.LazyProc
var _wkeWebViewName *windows.LazyProc
var _wkeSetWebViewName *windows.LazyProc
var _wkeIsLoaded *windows.LazyProc
var _wkeIsLoadFailed *windows.LazyProc
var _wkeIsLoadComplete *windows.LazyProc
var _wkeGetSource *windows.LazyProc
var _wkeTitle *windows.LazyProc
var _wkeTitleW *windows.LazyProc
var _wkeWidth *windows.LazyProc
var _wkeHeight *windows.LazyProc
var _wkeContentsWidth *windows.LazyProc
var _wkeContentsHeight *windows.LazyProc
var _wkeSelectAll *windows.LazyProc
var _wkeCopy *windows.LazyProc
var _wkeCut *windows.LazyProc
var _wkePaste *windows.LazyProc
var _wkeDelete *windows.LazyProc
var _wkeCookieEnabled *windows.LazyProc
var _wkeMediaVolume *windows.LazyProc
var _wkeMouseEvent *windows.LazyProc
var _wkeContextMenuEvent *windows.LazyProc
var _wkeMouseWheel *windows.LazyProc
var _wkeKeyUp *windows.LazyProc
var _wkeKeyDown *windows.LazyProc
var _wkeKeyPress *windows.LazyProc
var _wkeFocus *windows.LazyProc
var _wkeUnfocus *windows.LazyProc
var _wkeGetCaret *windows.LazyProc
var _wkeAwaken *windows.LazyProc
var _wkeZoomFactor *windows.LazyProc
var _wkeSetClientHandler *windows.LazyProc
var _wkeGetClientHandler *windows.LazyProc
var _jsToString *windows.LazyProc
var _jsToStringW *windows.LazyProc
var _wkeSetViewSettings *windows.LazyProc
var _wkeSetDebugConfig *windows.LazyProc
var _wkeToString *windows.LazyProc
var _wkeToStringW *windows.LazyProc
var _wkeGetDebugConfig *windows.LazyProc
var _wkeIsInitialize *windows.LazyProc
var _wkeFinalize *windows.LazyProc
var _wkeUpdate *windows.LazyProc
var _wkeGetVersion *windows.LazyProc
var _wkeGetVersionString *windows.LazyProc
var _wkeCreateWebView *windows.LazyProc
var _wkeDestroyWebView *windows.LazyProc
var _wkeSetMemoryCacheEnable *windows.LazyProc
var _wkeSetMouseEnabled *windows.LazyProc
var _wkeSetTouchEnabled *windows.LazyProc
var _wkeSetSystemTouchEnabled *windows.LazyProc
var _wkeSetContextMenuEnabled *windows.LazyProc
var _wkeSetNavigationToNewWindowEnable *windows.LazyProc
var _wkeSetCspCheckEnable *windows.LazyProc
var _wkeSetNpapiPluginsEnabled *windows.LazyProc
var _wkeSetHeadlessEnabled *windows.LazyProc
var _wkeSetDragEnable *windows.LazyProc
var _wkeSetDragDropEnable *windows.LazyProc
var _wkeSetLanguage *windows.LazyProc
var _wkeSetViewNetInterface *windows.LazyProc
var _wkeSetContextMenuItemShow *windows.LazyProc
var _wkeSetProxy *windows.LazyProc
var _wkeGetName *windows.LazyProc
var _wkeIsTransparent *windows.LazyProc
var _wkeGetUserAgent *windows.LazyProc
var _wkeSetViewProxy *windows.LazyProc
var _wkeSetName *windows.LazyProc
var _wkeSetHandle *windows.LazyProc
var _wkeSetTransparent *windows.LazyProc
var _wkeSetUserAgent *windows.LazyProc
var _wkeSetUserAgentW *windows.LazyProc
var _wkeSetHandleOffset *windows.LazyProc
var _wkeShowDevtools *windows.LazyProc
var _wkeLoadW *windows.LazyProc
var _wkeLoadURL *windows.LazyProc
var _wkeLoadURLW *windows.LazyProc
var _wkePostURL *windows.LazyProc
var _wkePostURLW *windows.LazyProc
var _wkeLoadHTML *windows.LazyProc
var _wkeLoadHtmlWithBaseUrl *windows.LazyProc
var _wkeLoadHTMLW *windows.LazyProc
var _wkeLoadFile *windows.LazyProc
var _wkeLoadFileW *windows.LazyProc
var _wkeGetURL *windows.LazyProc
var _wkeGetFrameUrl *windows.LazyProc
var _wkeIsLoading *windows.LazyProc
var _wkeIsLoadingSucceeded *windows.LazyProc
var _wkeIsLoadingFailed *windows.LazyProc
var _wkeIsLoadingCompleted *windows.LazyProc
var _wkeIsDocumentReady *windows.LazyProc
var _wkeStopLoading *windows.LazyProc
var _wkeReload *windows.LazyProc
var _wkeGoToOffset *windows.LazyProc
var _wkeGoToIndex *windows.LazyProc
var _wkeGetWebviewId *windows.LazyProc
var _wkeIsWebviewAlive *windows.LazyProc
var _wkeIsWebviewValid *windows.LazyProc
var _wkeGetDocumentCompleteURL *windows.LazyProc
var _wkeCreateMemBuf *windows.LazyProc
var _wkeFreeMemBuf *windows.LazyProc
var _wkeGetTitle *windows.LazyProc
var _wkeGetTitleW *windows.LazyProc
var _wkeResize *windows.LazyProc
var _wkeGetWidth *windows.LazyProc
var _wkeGetHeight *windows.LazyProc
var _wkeGetContentWidth *windows.LazyProc
var _wkeGetContentHeight *windows.LazyProc
var _wkeSetDirty *windows.LazyProc
var _wkeIsDirty *windows.LazyProc
var _wkeAddDirtyArea *windows.LazyProc
var _wkeLayoutIfNeeded *windows.LazyProc
var _wkePaint2 *windows.LazyProc
var _wkePaint *windows.LazyProc
var _wkeRepaintIfNeeded *windows.LazyProc
var _wkeGetViewDC *windows.LazyProc
var _wkeUnlockViewDC *windows.LazyProc
var _wkeGetHostHWND *windows.LazyProc
var _wkeCanGoBack *windows.LazyProc
var _wkeGoBack *windows.LazyProc
var _wkeCanGoForward *windows.LazyProc
var _wkeGoForward *windows.LazyProc
var _wkeNavigateAtIndex *windows.LazyProc
var _wkeGetNavigateIndex *windows.LazyProc
var _wkeEditorSelectAll *windows.LazyProc
var _wkeEditorUnSelect *windows.LazyProc
var _wkeEditorCopy *windows.LazyProc
var _wkeEditorCut *windows.LazyProc
var _wkeEditorPaste *windows.LazyProc
var _wkeEditorDelete *windows.LazyProc
var _wkeEditorUndo *windows.LazyProc
var _wkeEditorRedo *windows.LazyProc
var _wkeGetCookieW *windows.LazyProc
var _wkeGetCookie *windows.LazyProc
var _wkeSetCookie *windows.LazyProc
var _wkeVisitAllCookie *windows.LazyProc
var _wkePerformCookieCommand *windows.LazyProc
var _wkeSetCookieEnabled *windows.LazyProc
var _wkeIsCookieEnabled *windows.LazyProc
var _wkeSetCookieJarPath *windows.LazyProc
var _wkeSetCookieJarFullPath *windows.LazyProc
var _wkeClearCookie *windows.LazyProc
var _wkeSetLocalStorageFullPath *windows.LazyProc
var _wkeAddPluginDirectory *windows.LazyProc
var _wkeSetMediaVolume *windows.LazyProc
var _wkeGetMediaVolume *windows.LazyProc
var _wkeFireMouseEvent *windows.LazyProc
var _wkeFireContextMenuEvent *windows.LazyProc
var _wkeFireMouseWheelEvent *windows.LazyProc
var _wkeFireKeyUpEvent *windows.LazyProc
var _wkeFireKeyDownEvent *windows.LazyProc
var _wkeFireKeyPressEvent *windows.LazyProc
var _wkeFireWindowsMessage *windows.LazyProc
var _wkeSetFocus *windows.LazyProc
var _wkeKillFocus *windows.LazyProc
var _wkeGetCaretRect *windows.LazyProc
var _wkeGetCaretRect2 *windows.LazyProc
var _wkeRunJS *windows.LazyProc
var _wkeRunJSW *windows.LazyProc
var _wkeGlobalExec *windows.LazyProc
var _wkeGetGlobalExecByFrame *windows.LazyProc
var _wkeSleep *windows.LazyProc
var _wkeWake *windows.LazyProc
var _wkeIsAwake *windows.LazyProc
var _wkeSetZoomFactor *windows.LazyProc
var _wkeGetZoomFactor *windows.LazyProc
var _wkeEnableHighDPISupport *windows.LazyProc
var _wkeSetEditable *windows.LazyProc
var _wkeGetString *windows.LazyProc
var _wkeGetStringW *windows.LazyProc
var _wkeSetString *windows.LazyProc
var _wkeSetStringWithoutNullTermination *windows.LazyProc
var _wkeSetStringW *windows.LazyProc
var _wkeCreateString *windows.LazyProc
var _wkeCreateStringW *windows.LazyProc
var _wkeCreateStringWithoutNullTermination *windows.LazyProc
var _wkeGetStringLen *windows.LazyProc
var _wkeDeleteString *windows.LazyProc
var _wkeGetWebViewForCurrentContext *windows.LazyProc
var _wkeSetUserKeyValue *windows.LazyProc
var _wkeGetUserKeyValue *windows.LazyProc
var _wkeGetCursorInfoType *windows.LazyProc
var _wkeSetCursorInfoType *windows.LazyProc
var _wkeSetDragFiles *windows.LazyProc
var _wkeSetDeviceParameter *windows.LazyProc
var _wkeGetTempCallbackInfo *windows.LazyProc
var _wkeOnCaretChanged *windows.LazyProc
var _wkeOnMouseOverUrlChanged *windows.LazyProc
var _wkeOnTitleChanged *windows.LazyProc
var _wkeOnURLChanged *windows.LazyProc
var _wkeOnURLChanged2 *windows.LazyProc
var _wkeOnPaintUpdated *windows.LazyProc
var _wkeOnPaintBitUpdated *windows.LazyProc
var _wkeOnAlertBox *windows.LazyProc
var _wkeOnConfirmBox *windows.LazyProc
var _wkeOnPromptBox *windows.LazyProc
var _wkeOnNavigation *windows.LazyProc
var _wkeOnCreateView *windows.LazyProc
var _wkeOnDocumentReady *windows.LazyProc
var _wkeOnDocumentReady2 *windows.LazyProc
var _wkeOnLoadingFinish *windows.LazyProc
var _wkeOnDownload *windows.LazyProc
var _wkeOnDownload2 *windows.LazyProc
var _wkeOnConsole *windows.LazyProc
var _wkeSetUIThreadCallback *windows.LazyProc
var _wkeOnLoadUrlBegin *windows.LazyProc
var _wkeOnLoadUrlEnd *windows.LazyProc
var _wkeOnLoadUrlHeadersReceived *windows.LazyProc
var _wkeOnLoadUrlFinish *windows.LazyProc
var _wkeOnLoadUrlFail *windows.LazyProc
var _wkeOnDidCreateScriptContext *windows.LazyProc
var _wkeOnWillReleaseScriptContext *windows.LazyProc
var _wkeOnWindowClosing *windows.LazyProc
var _wkeOnWindowDestroy *windows.LazyProc
var _wkeOnDraggableRegionsChanged *windows.LazyProc
var _wkeOnWillMediaLoad *windows.LazyProc
var _wkeOnStartDragging *windows.LazyProc
var _wkeOnPrint *windows.LazyProc
var _wkeScreenshot *windows.LazyProc
var _wkeOnOtherLoad *windows.LazyProc
var _wkeOnContextMenuItemClick *windows.LazyProc
var _wkeIsProcessingUserGesture *windows.LazyProc
var _wkeNetSetMIMEType *windows.LazyProc
var _wkeNetGetMIMEType *windows.LazyProc
var _wkeNetGetReferrer *windows.LazyProc
var _wkeNetSetHTTPHeaderField *windows.LazyProc
var _wkeNetGetHTTPHeaderField *windows.LazyProc
var _wkeNetGetHTTPHeaderFieldFromResponse *windows.LazyProc
var _wkeNetSetData *windows.LazyProc
var _wkeNetHookRequest *windows.LazyProc
var _wkeNetGetRequestMethod *windows.LazyProc
var _wkeNetContinueJob *windows.LazyProc
var _wkeNetGetUrlByJob *windows.LazyProc
var _wkeNetGetRawHttpHead *windows.LazyProc
var _wkeNetGetRawResponseHead *windows.LazyProc
var _wkeNetCancelRequest *windows.LazyProc
var _wkeNetHoldJobToAsynCommit *windows.LazyProc
var _wkeNetOnResponse *windows.LazyProc
var _wkeNetGetFavicon *windows.LazyProc
var _wkeNetCreateWebUrlRequest *windows.LazyProc
var _wkeNetCreateWebUrlRequest2 *windows.LazyProc
var _wkeNetCopyWebUrlRequest *windows.LazyProc
var _wkeNetDeleteBlinkWebURLRequestPtr *windows.LazyProc
var _wkeNetAddHTTPHeaderFieldToUrlRequest *windows.LazyProc
var _wkeNetStartUrlRequest *windows.LazyProc
var _wkeNetGetHttpStatusCode *windows.LazyProc
var _wkeNetGetExpectedContentLength *windows.LazyProc
var _wkeNetGetResponseUrl *windows.LazyProc
var _wkeNetCancelWebUrlRequest *windows.LazyProc
var _wkeNetGetPostBody *windows.LazyProc
var _wkeNetFreePostBodyElements *windows.LazyProc
var _wkeNetCreatePostBodyElements *windows.LazyProc
var _wkeNetCreatePostBodyElement *windows.LazyProc
var _wkeNetFreePostBodyElement *windows.LazyProc
var _wkePopupDialogAndDownload *windows.LazyProc
var _wkeDownloadByPath *windows.LazyProc
var _wkeIsMainFrame *windows.LazyProc
var _wkeIsWebRemoteFrame *windows.LazyProc
var _wkeWebFrameGetMainFrame *windows.LazyProc
var _wkeRunJsByFrame *windows.LazyProc
var _wkeInsertCSSByFrame *windows.LazyProc
var _wkeWebFrameGetMainWorldScriptContext *windows.LazyProc
var _wkeGetBlinkMainThreadIsolate *windows.LazyProc
var _wkeCreateWebWindow *windows.LazyProc
var _wkeCreateWebCustomWindow *windows.LazyProc
var _wkeDestroyWebWindow *windows.LazyProc
var _wkeGetWindowHandle *windows.LazyProc
var _wkeShowWindow *windows.LazyProc
var _wkeEnableWindow *windows.LazyProc
var _wkeMoveWindow *windows.LazyProc
var _wkeMoveToCenter *windows.LazyProc
var _wkeResizeWindow *windows.LazyProc
var _wkeDragTargetDragEnter *windows.LazyProc
var _wkeDragTargetDragOver *windows.LazyProc
var _wkeDragTargetDragLeave *windows.LazyProc
var _wkeDragTargetDrop *windows.LazyProc
var _wkeDragTargetEnd *windows.LazyProc
var _wkeUtilSetUiCallback *windows.LazyProc
var _wkeUtilSerializeToMHTML *windows.LazyProc
var _wkeUtilPrintToPdf *windows.LazyProc
var _wkePrintToBitmap *windows.LazyProc
var _wkeUtilRelasePrintPdfDatas *windows.LazyProc
var _wkeSetWindowTitle *windows.LazyProc
var _wkeSetWindowTitleW *windows.LazyProc
var _wkeNodeOnCreateProcess *windows.LazyProc
var _wkeOnPluginFind *windows.LazyProc
var _wkeAddNpapiPlugin *windows.LazyProc
var _wkePluginListBuilderAddPlugin *windows.LazyProc
var _wkePluginListBuilderAddMediaTypeToLastPlugin *windows.LazyProc
var _wkePluginListBuilderAddFileExtensionToLastMediaType *windows.LazyProc
var _wkeGetWebViewByNData *windows.LazyProc
var _wkeRegisterEmbedderCustomElement *windows.LazyProc
var _wkeSetMediaPlayerFactory *windows.LazyProc
var _wkeGetContentAsMarkup *windows.LazyProc
var _wkeUtilDecodeURLEscape *windows.LazyProc
var _wkeUtilEncodeURLEscape *windows.LazyProc
var _wkeUtilBase64Encode *windows.LazyProc
var _wkeUtilBase64Decode *windows.LazyProc
var _wkeUtilCreateV8Snapshot *windows.LazyProc
var _wkeRunMessageLoop *windows.LazyProc
var _wkeSaveMemoryCache *windows.LazyProc
var _wkeJsBindFunction *windows.LazyProc
var _wkeJsBindGetter *windows.LazyProc
var _wkeJsBindSetter *windows.LazyProc
var _jsArgCount *windows.LazyProc
var _jsArgType *windows.LazyProc
var _jsArg *windows.LazyProc
var _jsTypeOf *windows.LazyProc
var _jsIsNumber *windows.LazyProc
var _jsIsString *windows.LazyProc
var _jsIsBoolean *windows.LazyProc
var _jsIsObject *windows.LazyProc
var _jsIsFunction *windows.LazyProc
var _jsIsUndefined *windows.LazyProc
var _jsIsNull *windows.LazyProc
var _jsIsArray *windows.LazyProc
var _jsIsTrue *windows.LazyProc
var _jsIsFalse *windows.LazyProc
var _jsToInt *windows.LazyProc
var _jsToFloat *windows.LazyProc
var _jsToDouble *windows.LazyProc
var _jsToDoubleString *windows.LazyProc
var _jsToBoolean *windows.LazyProc
var _jsArrayBuffer *windows.LazyProc
var _jsGetArrayBuffer *windows.LazyProc
var _jsToTempString *windows.LazyProc
var _jsToTempStringW *windows.LazyProc
var _jsToV8Value *windows.LazyProc
var _jsInt *windows.LazyProc
var _jsFloat *windows.LazyProc
var _jsDouble *windows.LazyProc
var _jsDoubleString *windows.LazyProc
var _jsBoolean *windows.LazyProc
var _jsUndefined *windows.LazyProc
var _jsNull *windows.LazyProc
var _jsTrue *windows.LazyProc
var _jsFalse *windows.LazyProc
var _jsString *windows.LazyProc
var _jsStringW *windows.LazyProc
var _jsEmptyObject *windows.LazyProc
var _jsEmptyArray *windows.LazyProc
var _jsObject *windows.LazyProc
var _jsFunction *windows.LazyProc
var _jsGetData *windows.LazyProc
var _jsGet *windows.LazyProc
var _jsSet *windows.LazyProc
var _jsGetAt *windows.LazyProc
var _jsSetAt *windows.LazyProc
var _jsGetKeys *windows.LazyProc
var _jsIsJsValueValid *windows.LazyProc
var _jsIsValidExecState *windows.LazyProc
var _jsDeleteObjectProp *windows.LazyProc
var _jsGetLength *windows.LazyProc
var _jsSetLength *windows.LazyProc
var _jsGlobalObject *windows.LazyProc
var _jsGetWebView *windows.LazyProc
var _jsEval *windows.LazyProc
var _jsEvalW *windows.LazyProc
var _jsEvalExW *windows.LazyProc
var _jsCall *windows.LazyProc
var _jsCallGlobal *windows.LazyProc
var _jsGetGlobal *windows.LazyProc
var _jsSetGlobal *windows.LazyProc
var _jsGC *windows.LazyProc
var _jsAddRef *windows.LazyProc
var _jsReleaseRef *windows.LazyProc
var _jsGetLastErrorIfException *windows.LazyProc
var _jsThrowException *windows.LazyProc
var _jsGetCallstack *windows.LazyProc
var _wkeInit *windows.LazyProc
var _wkeInitialize *windows.LazyProc
var _wkeInitializeEx *windows.LazyProc

func LoadLibrary(name string) error {
	dll, err := lib.LoadLibrary(name)

	if err != nil {
		return err
	}

	wkeDLL = windows.NewLazyDLL(dll.Name)

	fmt.Printf("DLL Name: %s\n", wkeDLL.Name)

	initProcs()

	return nil
}

func initProcs() {

	_wkeJsNativeFunction = wkeDLL.NewProc("wkeJsNativeFunction")
	_wkeShutdown = wkeDLL.NewProc("wkeShutdown")
	_wkeShutdownForDebug = wkeDLL.NewProc("wkeShutdownForDebug")
	_wkeVersion = wkeDLL.NewProc("wkeVersion")
	_wkeVersionString = wkeDLL.NewProc("wkeVersionString")
	_wkeGC = wkeDLL.NewProc("wkeGC")
	_wkeSetResourceGc = wkeDLL.NewProc("wkeSetResourceGc")
	_wkeSetFileSystem = wkeDLL.NewProc("wkeSetFileSystem")
	_wkeWebViewName = wkeDLL.NewProc("wkeWebViewName")
	_wkeSetWebViewName = wkeDLL.NewProc("wkeSetWebViewName")
	_wkeIsLoaded = wkeDLL.NewProc("wkeIsLoaded")
	_wkeIsLoadFailed = wkeDLL.NewProc("wkeIsLoadFailed")
	_wkeIsLoadComplete = wkeDLL.NewProc("wkeIsLoadComplete")
	_wkeGetSource = wkeDLL.NewProc("wkeGetSource")
	_wkeTitle = wkeDLL.NewProc("wkeTitle")
	_wkeTitleW = wkeDLL.NewProc("wkeTitleW")
	_wkeWidth = wkeDLL.NewProc("wkeWidth")
	_wkeHeight = wkeDLL.NewProc("wkeHeight")
	_wkeContentsWidth = wkeDLL.NewProc("wkeContentsWidth")
	_wkeContentsHeight = wkeDLL.NewProc("wkeContentsHeight")
	_wkeSelectAll = wkeDLL.NewProc("wkeSelectAll")
	_wkeCopy = wkeDLL.NewProc("wkeCopy")
	_wkeCut = wkeDLL.NewProc("wkeCut")
	_wkePaste = wkeDLL.NewProc("wkePaste")
	_wkeDelete = wkeDLL.NewProc("wkeDelete")
	_wkeCookieEnabled = wkeDLL.NewProc("wkeCookieEnabled")
	_wkeMediaVolume = wkeDLL.NewProc("wkeMediaVolume")
	_wkeMouseEvent = wkeDLL.NewProc("wkeMouseEvent")
	_wkeContextMenuEvent = wkeDLL.NewProc("wkeContextMenuEvent")
	_wkeMouseWheel = wkeDLL.NewProc("wkeMouseWheel")
	_wkeKeyUp = wkeDLL.NewProc("wkeKeyUp")
	_wkeKeyDown = wkeDLL.NewProc("wkeKeyDown")
	_wkeKeyPress = wkeDLL.NewProc("wkeKeyPress")
	_wkeFocus = wkeDLL.NewProc("wkeFocus")
	_wkeUnfocus = wkeDLL.NewProc("wkeUnfocus")
	_wkeGetCaret = wkeDLL.NewProc("wkeGetCaret")
	_wkeAwaken = wkeDLL.NewProc("wkeAwaken")
	_wkeZoomFactor = wkeDLL.NewProc("wkeZoomFactor")
	_wkeSetClientHandler = wkeDLL.NewProc("wkeSetClientHandler")
	_wkeGetClientHandler = wkeDLL.NewProc("wkeGetClientHandler")
	_jsToString = wkeDLL.NewProc("jsToString")
	_jsToStringW = wkeDLL.NewProc("jsToStringW")
	_wkeSetViewSettings = wkeDLL.NewProc("wkeSetViewSettings")
	_wkeSetDebugConfig = wkeDLL.NewProc("wkeSetDebugConfig")
	_wkeToString = wkeDLL.NewProc("wkeToString")
	_wkeToStringW = wkeDLL.NewProc("wkeToStringW")
	_wkeGetDebugConfig = wkeDLL.NewProc("wkeGetDebugConfig")
	_wkeIsInitialize = wkeDLL.NewProc("wkeIsInitialize")
	_wkeFinalize = wkeDLL.NewProc("wkeFinalize")
	_wkeUpdate = wkeDLL.NewProc("wkeUpdate")
	_wkeGetVersion = wkeDLL.NewProc("wkeGetVersion")
	_wkeGetVersionString = wkeDLL.NewProc("wkeGetVersionString")
	_wkeCreateWebView = wkeDLL.NewProc("wkeCreateWebView")
	_wkeDestroyWebView = wkeDLL.NewProc("wkeDestroyWebView")
	_wkeSetMemoryCacheEnable = wkeDLL.NewProc("wkeSetMemoryCacheEnable")
	_wkeSetMouseEnabled = wkeDLL.NewProc("wkeSetMouseEnabled")
	_wkeSetTouchEnabled = wkeDLL.NewProc("wkeSetTouchEnabled")
	_wkeSetSystemTouchEnabled = wkeDLL.NewProc("wkeSetSystemTouchEnabled")
	_wkeSetContextMenuEnabled = wkeDLL.NewProc("wkeSetContextMenuEnabled")
	_wkeSetNavigationToNewWindowEnable = wkeDLL.NewProc("wkeSetNavigationToNewWindowEnable")
	_wkeSetCspCheckEnable = wkeDLL.NewProc("wkeSetCspCheckEnable")
	_wkeSetNpapiPluginsEnabled = wkeDLL.NewProc("wkeSetNpapiPluginsEnabled")
	_wkeSetHeadlessEnabled = wkeDLL.NewProc("wkeSetHeadlessEnabled")
	_wkeSetDragEnable = wkeDLL.NewProc("wkeSetDragEnable")
	_wkeSetDragDropEnable = wkeDLL.NewProc("wkeSetDragDropEnable")
	_wkeSetLanguage = wkeDLL.NewProc("wkeSetLanguage")
	_wkeSetViewNetInterface = wkeDLL.NewProc("wkeSetViewNetInterface")
	_wkeSetContextMenuItemShow = wkeDLL.NewProc("wkeSetContextMenuItemShow")
	_wkeSetProxy = wkeDLL.NewProc("wkeSetProxy")
	_wkeGetName = wkeDLL.NewProc("wkeGetName")
	_wkeIsTransparent = wkeDLL.NewProc("wkeIsTransparent")
	_wkeGetUserAgent = wkeDLL.NewProc("wkeGetUserAgent")
	_wkeSetViewProxy = wkeDLL.NewProc("wkeSetViewProxy")
	_wkeSetName = wkeDLL.NewProc("wkeSetName")
	_wkeSetHandle = wkeDLL.NewProc("wkeSetHandle")
	_wkeSetTransparent = wkeDLL.NewProc("wkeSetTransparent")
	_wkeSetUserAgent = wkeDLL.NewProc("wkeSetUserAgent")
	_wkeSetUserAgentW = wkeDLL.NewProc("wkeSetUserAgentW")
	_wkeSetHandleOffset = wkeDLL.NewProc("wkeSetHandleOffset")
	_wkeShowDevtools = wkeDLL.NewProc("wkeShowDevtools")
	_wkeLoadW = wkeDLL.NewProc("wkeLoadW")
	_wkeLoadURL = wkeDLL.NewProc("wkeLoadURL")
	_wkeLoadURLW = wkeDLL.NewProc("wkeLoadURLW")
	_wkePostURL = wkeDLL.NewProc("wkePostURL")
	_wkePostURLW = wkeDLL.NewProc("wkePostURLW")
	_wkeLoadHTML = wkeDLL.NewProc("wkeLoadHTML")
	_wkeLoadHtmlWithBaseUrl = wkeDLL.NewProc("wkeLoadHtmlWithBaseUrl")
	_wkeLoadHTMLW = wkeDLL.NewProc("wkeLoadHTMLW")
	_wkeLoadFile = wkeDLL.NewProc("wkeLoadFile")
	_wkeLoadFileW = wkeDLL.NewProc("wkeLoadFileW")
	_wkeGetURL = wkeDLL.NewProc("wkeGetURL")
	_wkeGetFrameUrl = wkeDLL.NewProc("wkeGetFrameUrl")
	_wkeIsLoading = wkeDLL.NewProc("wkeIsLoading")
	_wkeIsLoadingSucceeded = wkeDLL.NewProc("wkeIsLoadingSucceeded")
	_wkeIsLoadingFailed = wkeDLL.NewProc("wkeIsLoadingFailed")
	_wkeIsLoadingCompleted = wkeDLL.NewProc("wkeIsLoadingCompleted")
	_wkeIsDocumentReady = wkeDLL.NewProc("wkeIsDocumentReady")
	_wkeStopLoading = wkeDLL.NewProc("wkeStopLoading")
	_wkeReload = wkeDLL.NewProc("wkeReload")
	_wkeGoToOffset = wkeDLL.NewProc("wkeGoToOffset")
	_wkeGoToIndex = wkeDLL.NewProc("wkeGoToIndex")
	_wkeGetWebviewId = wkeDLL.NewProc("wkeGetWebviewId")
	_wkeIsWebviewAlive = wkeDLL.NewProc("wkeIsWebviewAlive")
	_wkeIsWebviewValid = wkeDLL.NewProc("wkeIsWebviewValid")
	_wkeGetDocumentCompleteURL = wkeDLL.NewProc("wkeGetDocumentCompleteURL")
	_wkeCreateMemBuf = wkeDLL.NewProc("wkeCreateMemBuf")
	_wkeFreeMemBuf = wkeDLL.NewProc("wkeFreeMemBuf")
	_wkeGetTitle = wkeDLL.NewProc("wkeGetTitle")
	_wkeGetTitleW = wkeDLL.NewProc("wkeGetTitleW")
	_wkeResize = wkeDLL.NewProc("wkeResize")
	_wkeGetWidth = wkeDLL.NewProc("wkeGetWidth")
	_wkeGetHeight = wkeDLL.NewProc("wkeGetHeight")
	_wkeGetContentWidth = wkeDLL.NewProc("wkeGetContentWidth")
	_wkeGetContentHeight = wkeDLL.NewProc("wkeGetContentHeight")
	_wkeSetDirty = wkeDLL.NewProc("wkeSetDirty")
	_wkeIsDirty = wkeDLL.NewProc("wkeIsDirty")
	_wkeAddDirtyArea = wkeDLL.NewProc("wkeAddDirtyArea")
	_wkeLayoutIfNeeded = wkeDLL.NewProc("wkeLayoutIfNeeded")
	_wkePaint2 = wkeDLL.NewProc("wkePaint2")
	_wkePaint = wkeDLL.NewProc("wkePaint")
	_wkeRepaintIfNeeded = wkeDLL.NewProc("wkeRepaintIfNeeded")
	_wkeGetViewDC = wkeDLL.NewProc("wkeGetViewDC")
	_wkeUnlockViewDC = wkeDLL.NewProc("wkeUnlockViewDC")
	_wkeGetHostHWND = wkeDLL.NewProc("wkeGetHostHWND")
	_wkeCanGoBack = wkeDLL.NewProc("wkeCanGoBack")
	_wkeGoBack = wkeDLL.NewProc("wkeGoBack")
	_wkeCanGoForward = wkeDLL.NewProc("wkeCanGoForward")
	_wkeGoForward = wkeDLL.NewProc("wkeGoForward")
	_wkeNavigateAtIndex = wkeDLL.NewProc("wkeNavigateAtIndex")
	_wkeGetNavigateIndex = wkeDLL.NewProc("wkeGetNavigateIndex")
	_wkeEditorSelectAll = wkeDLL.NewProc("wkeEditorSelectAll")
	_wkeEditorUnSelect = wkeDLL.NewProc("wkeEditorUnSelect")
	_wkeEditorCopy = wkeDLL.NewProc("wkeEditorCopy")
	_wkeEditorCut = wkeDLL.NewProc("wkeEditorCut")
	_wkeEditorPaste = wkeDLL.NewProc("wkeEditorPaste")
	_wkeEditorDelete = wkeDLL.NewProc("wkeEditorDelete")
	_wkeEditorUndo = wkeDLL.NewProc("wkeEditorUndo")
	_wkeEditorRedo = wkeDLL.NewProc("wkeEditorRedo")
	_wkeGetCookieW = wkeDLL.NewProc("wkeGetCookieW")
	_wkeGetCookie = wkeDLL.NewProc("wkeGetCookie")
	_wkeSetCookie = wkeDLL.NewProc("wkeSetCookie")
	_wkeVisitAllCookie = wkeDLL.NewProc("wkeVisitAllCookie")
	_wkePerformCookieCommand = wkeDLL.NewProc("wkePerformCookieCommand")
	_wkeSetCookieEnabled = wkeDLL.NewProc("wkeSetCookieEnabled")
	_wkeIsCookieEnabled = wkeDLL.NewProc("wkeIsCookieEnabled")
	_wkeSetCookieJarPath = wkeDLL.NewProc("wkeSetCookieJarPath")
	_wkeSetCookieJarFullPath = wkeDLL.NewProc("wkeSetCookieJarFullPath")
	_wkeClearCookie = wkeDLL.NewProc("wkeClearCookie")
	_wkeSetLocalStorageFullPath = wkeDLL.NewProc("wkeSetLocalStorageFullPath")
	_wkeAddPluginDirectory = wkeDLL.NewProc("wkeAddPluginDirectory")
	_wkeSetMediaVolume = wkeDLL.NewProc("wkeSetMediaVolume")
	_wkeGetMediaVolume = wkeDLL.NewProc("wkeGetMediaVolume")
	_wkeFireMouseEvent = wkeDLL.NewProc("wkeFireMouseEvent")
	_wkeFireContextMenuEvent = wkeDLL.NewProc("wkeFireContextMenuEvent")
	_wkeFireMouseWheelEvent = wkeDLL.NewProc("wkeFireMouseWheelEvent")
	_wkeFireKeyUpEvent = wkeDLL.NewProc("wkeFireKeyUpEvent")
	_wkeFireKeyDownEvent = wkeDLL.NewProc("wkeFireKeyDownEvent")
	_wkeFireKeyPressEvent = wkeDLL.NewProc("wkeFireKeyPressEvent")
	_wkeFireWindowsMessage = wkeDLL.NewProc("wkeFireWindowsMessage")
	_wkeSetFocus = wkeDLL.NewProc("wkeSetFocus")
	_wkeKillFocus = wkeDLL.NewProc("wkeKillFocus")
	_wkeGetCaretRect = wkeDLL.NewProc("wkeGetCaretRect")
	_wkeGetCaretRect2 = wkeDLL.NewProc("wkeGetCaretRect2")
	_wkeRunJS = wkeDLL.NewProc("wkeRunJS")
	_wkeRunJSW = wkeDLL.NewProc("wkeRunJSW")
	_wkeGlobalExec = wkeDLL.NewProc("wkeGlobalExec")
	_wkeGetGlobalExecByFrame = wkeDLL.NewProc("wkeGetGlobalExecByFrame")
	_wkeSleep = wkeDLL.NewProc("wkeSleep")
	_wkeWake = wkeDLL.NewProc("wkeWake")
	_wkeIsAwake = wkeDLL.NewProc("wkeIsAwake")
	_wkeSetZoomFactor = wkeDLL.NewProc("wkeSetZoomFactor")
	_wkeGetZoomFactor = wkeDLL.NewProc("wkeGetZoomFactor")
	_wkeEnableHighDPISupport = wkeDLL.NewProc("wkeEnableHighDPISupport")
	_wkeSetEditable = wkeDLL.NewProc("wkeSetEditable")
	_wkeGetString = wkeDLL.NewProc("wkeGetString")
	_wkeGetStringW = wkeDLL.NewProc("wkeGetStringW")
	_wkeSetString = wkeDLL.NewProc("wkeSetString")
	_wkeSetStringWithoutNullTermination = wkeDLL.NewProc("wkeSetStringWithoutNullTermination")
	_wkeSetStringW = wkeDLL.NewProc("wkeSetStringW")
	_wkeCreateString = wkeDLL.NewProc("wkeCreateString")
	_wkeCreateStringW = wkeDLL.NewProc("wkeCreateStringW")
	_wkeCreateStringWithoutNullTermination = wkeDLL.NewProc("wkeCreateStringWithoutNullTermination")
	_wkeGetStringLen = wkeDLL.NewProc("wkeGetStringLen")
	_wkeDeleteString = wkeDLL.NewProc("wkeDeleteString")
	_wkeGetWebViewForCurrentContext = wkeDLL.NewProc("wkeGetWebViewForCurrentContext")
	_wkeSetUserKeyValue = wkeDLL.NewProc("wkeSetUserKeyValue")
	_wkeGetUserKeyValue = wkeDLL.NewProc("wkeGetUserKeyValue")
	_wkeGetCursorInfoType = wkeDLL.NewProc("wkeGetCursorInfoType")
	_wkeSetCursorInfoType = wkeDLL.NewProc("wkeSetCursorInfoType")
	_wkeSetDragFiles = wkeDLL.NewProc("wkeSetDragFiles")
	_wkeSetDeviceParameter = wkeDLL.NewProc("wkeSetDeviceParameter")
	_wkeGetTempCallbackInfo = wkeDLL.NewProc("wkeGetTempCallbackInfo")
	_wkeOnCaretChanged = wkeDLL.NewProc("wkeOnCaretChanged")
	_wkeOnMouseOverUrlChanged = wkeDLL.NewProc("wkeOnMouseOverUrlChanged")
	_wkeOnTitleChanged = wkeDLL.NewProc("wkeOnTitleChanged")
	_wkeOnURLChanged = wkeDLL.NewProc("wkeOnURLChanged")
	_wkeOnURLChanged2 = wkeDLL.NewProc("wkeOnURLChanged2")
	_wkeOnPaintUpdated = wkeDLL.NewProc("wkeOnPaintUpdated")
	_wkeOnPaintBitUpdated = wkeDLL.NewProc("wkeOnPaintBitUpdated")
	_wkeOnAlertBox = wkeDLL.NewProc("wkeOnAlertBox")
	_wkeOnConfirmBox = wkeDLL.NewProc("wkeOnConfirmBox")
	_wkeOnPromptBox = wkeDLL.NewProc("wkeOnPromptBox")
	_wkeOnNavigation = wkeDLL.NewProc("wkeOnNavigation")
	_wkeOnCreateView = wkeDLL.NewProc("wkeOnCreateView")
	_wkeOnDocumentReady = wkeDLL.NewProc("wkeOnDocumentReady")
	_wkeOnDocumentReady2 = wkeDLL.NewProc("wkeOnDocumentReady2")
	_wkeOnLoadingFinish = wkeDLL.NewProc("wkeOnLoadingFinish")
	_wkeOnDownload = wkeDLL.NewProc("wkeOnDownload")
	_wkeOnDownload2 = wkeDLL.NewProc("wkeOnDownload2")
	_wkeOnConsole = wkeDLL.NewProc("wkeOnConsole")
	_wkeSetUIThreadCallback = wkeDLL.NewProc("wkeSetUIThreadCallback")
	_wkeOnLoadUrlBegin = wkeDLL.NewProc("wkeOnLoadUrlBegin")
	_wkeOnLoadUrlEnd = wkeDLL.NewProc("wkeOnLoadUrlEnd")
	_wkeOnLoadUrlHeadersReceived = wkeDLL.NewProc("wkeOnLoadUrlHeadersReceived")
	_wkeOnLoadUrlFinish = wkeDLL.NewProc("wkeOnLoadUrlFinish")
	_wkeOnLoadUrlFail = wkeDLL.NewProc("wkeOnLoadUrlFail")
	_wkeOnDidCreateScriptContext = wkeDLL.NewProc("wkeOnDidCreateScriptContext")
	_wkeOnWillReleaseScriptContext = wkeDLL.NewProc("wkeOnWillReleaseScriptContext")
	_wkeOnWindowClosing = wkeDLL.NewProc("wkeOnWindowClosing")
	_wkeOnWindowDestroy = wkeDLL.NewProc("wkeOnWindowDestroy")
	_wkeOnDraggableRegionsChanged = wkeDLL.NewProc("wkeOnDraggableRegionsChanged")
	_wkeOnWillMediaLoad = wkeDLL.NewProc("wkeOnWillMediaLoad")
	_wkeOnStartDragging = wkeDLL.NewProc("wkeOnStartDragging")
	_wkeOnPrint = wkeDLL.NewProc("wkeOnPrint")
	_wkeScreenshot = wkeDLL.NewProc("wkeScreenshot")
	_wkeOnOtherLoad = wkeDLL.NewProc("wkeOnOtherLoad")
	_wkeOnContextMenuItemClick = wkeDLL.NewProc("wkeOnContextMenuItemClick")
	_wkeIsProcessingUserGesture = wkeDLL.NewProc("wkeIsProcessingUserGesture")
	_wkeNetSetMIMEType = wkeDLL.NewProc("wkeNetSetMIMEType")
	_wkeNetGetMIMEType = wkeDLL.NewProc("wkeNetGetMIMEType")
	_wkeNetGetReferrer = wkeDLL.NewProc("wkeNetGetReferrer")
	_wkeNetSetHTTPHeaderField = wkeDLL.NewProc("wkeNetSetHTTPHeaderField")
	_wkeNetGetHTTPHeaderField = wkeDLL.NewProc("wkeNetGetHTTPHeaderField")
	_wkeNetGetHTTPHeaderFieldFromResponse = wkeDLL.NewProc("wkeNetGetHTTPHeaderFieldFromResponse")
	_wkeNetSetData = wkeDLL.NewProc("wkeNetSetData")
	_wkeNetHookRequest = wkeDLL.NewProc("wkeNetHookRequest")
	_wkeNetGetRequestMethod = wkeDLL.NewProc("wkeNetGetRequestMethod")
	_wkeNetContinueJob = wkeDLL.NewProc("wkeNetContinueJob")
	_wkeNetGetUrlByJob = wkeDLL.NewProc("wkeNetGetUrlByJob")
	_wkeNetGetRawHttpHead = wkeDLL.NewProc("wkeNetGetRawHttpHead")
	_wkeNetGetRawResponseHead = wkeDLL.NewProc("wkeNetGetRawResponseHead")
	_wkeNetCancelRequest = wkeDLL.NewProc("wkeNetCancelRequest")
	_wkeNetHoldJobToAsynCommit = wkeDLL.NewProc("wkeNetHoldJobToAsynCommit")
	_wkeNetOnResponse = wkeDLL.NewProc("wkeNetOnResponse")
	_wkeNetGetFavicon = wkeDLL.NewProc("wkeNetGetFavicon")
	_wkeNetCreateWebUrlRequest = wkeDLL.NewProc("wkeNetCreateWebUrlRequest")
	_wkeNetCreateWebUrlRequest2 = wkeDLL.NewProc("wkeNetCreateWebUrlRequest2")
	_wkeNetCopyWebUrlRequest = wkeDLL.NewProc("wkeNetCopyWebUrlRequest")
	_wkeNetDeleteBlinkWebURLRequestPtr = wkeDLL.NewProc("wkeNetDeleteBlinkWebURLRequestPtr")
	_wkeNetAddHTTPHeaderFieldToUrlRequest = wkeDLL.NewProc("wkeNetAddHTTPHeaderFieldToUrlRequest")
	_wkeNetStartUrlRequest = wkeDLL.NewProc("wkeNetStartUrlRequest")
	_wkeNetGetHttpStatusCode = wkeDLL.NewProc("wkeNetGetHttpStatusCode")
	_wkeNetGetExpectedContentLength = wkeDLL.NewProc("wkeNetGetExpectedContentLength")
	_wkeNetGetResponseUrl = wkeDLL.NewProc("wkeNetGetResponseUrl")
	_wkeNetCancelWebUrlRequest = wkeDLL.NewProc("wkeNetCancelWebUrlRequest")
	_wkeNetGetPostBody = wkeDLL.NewProc("wkeNetGetPostBody")
	_wkeNetFreePostBodyElements = wkeDLL.NewProc("wkeNetFreePostBodyElements")
	_wkeNetCreatePostBodyElements = wkeDLL.NewProc("wkeNetCreatePostBodyElements")
	_wkeNetCreatePostBodyElement = wkeDLL.NewProc("wkeNetCreatePostBodyElement")
	_wkeNetFreePostBodyElement = wkeDLL.NewProc("wkeNetFreePostBodyElement")
	_wkePopupDialogAndDownload = wkeDLL.NewProc("wkePopupDialogAndDownload")
	_wkeDownloadByPath = wkeDLL.NewProc("wkeDownloadByPath")
	_wkeIsMainFrame = wkeDLL.NewProc("wkeIsMainFrame")
	_wkeIsWebRemoteFrame = wkeDLL.NewProc("wkeIsWebRemoteFrame")
	_wkeWebFrameGetMainFrame = wkeDLL.NewProc("wkeWebFrameGetMainFrame")
	_wkeRunJsByFrame = wkeDLL.NewProc("wkeRunJsByFrame")
	_wkeInsertCSSByFrame = wkeDLL.NewProc("wkeInsertCSSByFrame")
	_wkeWebFrameGetMainWorldScriptContext = wkeDLL.NewProc("wkeWebFrameGetMainWorldScriptContext")
	_wkeGetBlinkMainThreadIsolate = wkeDLL.NewProc("wkeGetBlinkMainThreadIsolate")
	_wkeCreateWebWindow = wkeDLL.NewProc("wkeCreateWebWindow")
	_wkeCreateWebCustomWindow = wkeDLL.NewProc("wkeCreateWebCustomWindow")
	_wkeDestroyWebWindow = wkeDLL.NewProc("wkeDestroyWebWindow")
	_wkeGetWindowHandle = wkeDLL.NewProc("wkeGetWindowHandle")
	_wkeShowWindow = wkeDLL.NewProc("wkeShowWindow")
	_wkeEnableWindow = wkeDLL.NewProc("wkeEnableWindow")
	_wkeMoveWindow = wkeDLL.NewProc("wkeMoveWindow")
	_wkeMoveToCenter = wkeDLL.NewProc("wkeMoveToCenter")
	_wkeResizeWindow = wkeDLL.NewProc("wkeResizeWindow")
	_wkeDragTargetDragEnter = wkeDLL.NewProc("wkeDragTargetDragEnter")
	_wkeDragTargetDragOver = wkeDLL.NewProc("wkeDragTargetDragOver")
	_wkeDragTargetDragLeave = wkeDLL.NewProc("wkeDragTargetDragLeave")
	_wkeDragTargetDrop = wkeDLL.NewProc("wkeDragTargetDrop")
	_wkeDragTargetEnd = wkeDLL.NewProc("wkeDragTargetEnd")
	_wkeUtilSetUiCallback = wkeDLL.NewProc("wkeUtilSetUiCallback")
	_wkeUtilSerializeToMHTML = wkeDLL.NewProc("wkeUtilSerializeToMHTML")
	_wkeUtilPrintToPdf = wkeDLL.NewProc("wkeUtilPrintToPdf")
	_wkePrintToBitmap = wkeDLL.NewProc("wkePrintToBitmap")
	_wkeUtilRelasePrintPdfDatas = wkeDLL.NewProc("wkeUtilRelasePrintPdfDatas")
	_wkeSetWindowTitle = wkeDLL.NewProc("wkeSetWindowTitle")
	_wkeSetWindowTitleW = wkeDLL.NewProc("wkeSetWindowTitleW")
	_wkeNodeOnCreateProcess = wkeDLL.NewProc("wkeNodeOnCreateProcess")
	_wkeOnPluginFind = wkeDLL.NewProc("wkeOnPluginFind")
	_wkeAddNpapiPlugin = wkeDLL.NewProc("wkeAddNpapiPlugin")
	_wkePluginListBuilderAddPlugin = wkeDLL.NewProc("wkePluginListBuilderAddPlugin")
	_wkePluginListBuilderAddMediaTypeToLastPlugin = wkeDLL.NewProc("wkePluginListBuilderAddMediaTypeToLastPlugin")
	_wkePluginListBuilderAddFileExtensionToLastMediaType = wkeDLL.NewProc("wkePluginListBuilderAddFileExtensionToLastMediaType")
	_wkeGetWebViewByNData = wkeDLL.NewProc("wkeGetWebViewByNData")
	_wkeRegisterEmbedderCustomElement = wkeDLL.NewProc("wkeRegisterEmbedderCustomElement")
	_wkeSetMediaPlayerFactory = wkeDLL.NewProc("wkeSetMediaPlayerFactory")
	_wkeGetContentAsMarkup = wkeDLL.NewProc("wkeGetContentAsMarkup")
	_wkeUtilDecodeURLEscape = wkeDLL.NewProc("wkeUtilDecodeURLEscape")
	_wkeUtilEncodeURLEscape = wkeDLL.NewProc("wkeUtilEncodeURLEscape")
	_wkeUtilBase64Encode = wkeDLL.NewProc("wkeUtilBase64Encode")
	_wkeUtilBase64Decode = wkeDLL.NewProc("wkeUtilBase64Decode")
	_wkeUtilCreateV8Snapshot = wkeDLL.NewProc("wkeUtilCreateV8Snapshot")
	_wkeRunMessageLoop = wkeDLL.NewProc("wkeRunMessageLoop")
	_wkeSaveMemoryCache = wkeDLL.NewProc("wkeSaveMemoryCache")
	_wkeJsBindFunction = wkeDLL.NewProc("wkeJsBindFunction")
	_wkeJsBindGetter = wkeDLL.NewProc("wkeJsBindGetter")
	_wkeJsBindSetter = wkeDLL.NewProc("wkeJsBindSetter")
	_jsArgCount = wkeDLL.NewProc("jsArgCount")
	_jsArgType = wkeDLL.NewProc("jsArgType")
	_jsArg = wkeDLL.NewProc("jsArg")
	_jsTypeOf = wkeDLL.NewProc("jsTypeOf")
	_jsIsNumber = wkeDLL.NewProc("jsIsNumber")
	_jsIsString = wkeDLL.NewProc("jsIsString")
	_jsIsBoolean = wkeDLL.NewProc("jsIsBoolean")
	_jsIsObject = wkeDLL.NewProc("jsIsObject")
	_jsIsFunction = wkeDLL.NewProc("jsIsFunction")
	_jsIsUndefined = wkeDLL.NewProc("jsIsUndefined")
	_jsIsNull = wkeDLL.NewProc("jsIsNull")
	_jsIsArray = wkeDLL.NewProc("jsIsArray")
	_jsIsTrue = wkeDLL.NewProc("jsIsTrue")
	_jsIsFalse = wkeDLL.NewProc("jsIsFalse")
	_jsToInt = wkeDLL.NewProc("jsToInt")
	_jsToFloat = wkeDLL.NewProc("jsToFloat")
	_jsToDouble = wkeDLL.NewProc("jsToDouble")
	_jsToDoubleString = wkeDLL.NewProc("jsToDoubleString")
	_jsToBoolean = wkeDLL.NewProc("jsToBoolean")
	_jsArrayBuffer = wkeDLL.NewProc("jsArrayBuffer")
	_jsGetArrayBuffer = wkeDLL.NewProc("jsGetArrayBuffer")
	_jsToTempString = wkeDLL.NewProc("jsToTempString")
	_jsToTempStringW = wkeDLL.NewProc("jsToTempStringW")
	_jsToV8Value = wkeDLL.NewProc("jsToV8Value")
	_jsInt = wkeDLL.NewProc("jsInt")
	_jsFloat = wkeDLL.NewProc("jsFloat")
	_jsDouble = wkeDLL.NewProc("jsDouble")
	_jsDoubleString = wkeDLL.NewProc("jsDoubleString")
	_jsBoolean = wkeDLL.NewProc("jsBoolean")
	_jsUndefined = wkeDLL.NewProc("jsUndefined")
	_jsNull = wkeDLL.NewProc("jsNull")
	_jsTrue = wkeDLL.NewProc("jsTrue")
	_jsFalse = wkeDLL.NewProc("jsFalse")
	_jsString = wkeDLL.NewProc("jsString")
	_jsStringW = wkeDLL.NewProc("jsStringW")
	_jsEmptyObject = wkeDLL.NewProc("jsEmptyObject")
	_jsEmptyArray = wkeDLL.NewProc("jsEmptyArray")
	_jsObject = wkeDLL.NewProc("jsObject")
	_jsFunction = wkeDLL.NewProc("jsFunction")
	_jsGetData = wkeDLL.NewProc("jsGetData")
	_jsGet = wkeDLL.NewProc("jsGet")
	_jsSet = wkeDLL.NewProc("jsSet")
	_jsGetAt = wkeDLL.NewProc("jsGetAt")
	_jsSetAt = wkeDLL.NewProc("jsSetAt")
	_jsGetKeys = wkeDLL.NewProc("jsGetKeys")
	_jsIsJsValueValid = wkeDLL.NewProc("jsIsJsValueValid")
	_jsIsValidExecState = wkeDLL.NewProc("jsIsValidExecState")
	_jsDeleteObjectProp = wkeDLL.NewProc("jsDeleteObjectProp")
	_jsGetLength = wkeDLL.NewProc("jsGetLength")
	_jsSetLength = wkeDLL.NewProc("jsSetLength")
	_jsGlobalObject = wkeDLL.NewProc("jsGlobalObject")
	_jsGetWebView = wkeDLL.NewProc("jsGetWebView")
	_jsEval = wkeDLL.NewProc("jsEval")
	_jsEvalW = wkeDLL.NewProc("jsEvalW")
	_jsEvalExW = wkeDLL.NewProc("jsEvalExW")
	_jsCall = wkeDLL.NewProc("jsCall")
	_jsCallGlobal = wkeDLL.NewProc("jsCallGlobal")
	_jsGetGlobal = wkeDLL.NewProc("jsGetGlobal")
	_jsSetGlobal = wkeDLL.NewProc("jsSetGlobal")
	_jsGC = wkeDLL.NewProc("jsGC")
	_jsAddRef = wkeDLL.NewProc("jsAddRef")
	_jsReleaseRef = wkeDLL.NewProc("jsReleaseRef")
	_jsGetLastErrorIfException = wkeDLL.NewProc("jsGetLastErrorIfException")
	_jsThrowException = wkeDLL.NewProc("jsThrowException")
	_jsGetCallstack = wkeDLL.NewProc("jsGetCallstack")
	_wkeInit = wkeDLL.NewProc("wkeInit")
	_wkeInitialize = wkeDLL.NewProc("wkeInitialize")
	_wkeInitializeEx = wkeDLL.NewProc("wkeInitializeEx")
}

func Release() error {
	return lib.Release()
}

func JsNativeFunction(es JsExecState, param uintptr) JsValue {
	r, _, _ := _wkeJsNativeFunction.Call(uintptr(es), param)
	return JsValue(r)
}

func Shutdown() {
	_wkeShutdown.Call()
}

func ShutdownForDebug() {
	_wkeShutdownForDebug.Call()
}

func Version() uint {
	r, _, _ := _wkeVersion.Call()
	return uint(r)
}

func VersionString() string {
	r, _, _ := _wkeVersionString.Call()
	return StrFromPtr(r)
}

func GC(webView WebView, intervalSec int64) {
	_wkeGC.Call(uintptr(webView), uintptr(intervalSec))
}

func SetResourceGc(webView WebView, intervalSec int64) {
	_wkeSetResourceGc.Call(uintptr(webView), uintptr(intervalSec))
}

func SetFileSystem(pfnOpen FILE_OPEN, pfnClose FILE_CLOSE, pfnSize FILE_SIZE, pfnRead FILE_READ, pfnSeek FILE_SEEK) {
	_wkeSetFileSystem.Call(CallbackToPtr(pfnOpen), CallbackToPtr(pfnClose), CallbackToPtr(pfnSize), CallbackToPtr(pfnRead), CallbackToPtr(pfnSeek))
}

func WebViewName(webView WebView) string {
	r, _, _ := _wkeWebViewName.Call(uintptr(webView))
	return StrFromPtr(r)
}

func SetWebViewName(webView WebView, name string) {
	_wkeSetWebViewName.Call(uintptr(webView), StrToPtr(name))
}

func IsLoaded(webView WebView) bool {
	r, _, _ := _wkeIsLoaded.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func IsLoadFailed(webView WebView) bool {
	r, _, _ := _wkeIsLoadFailed.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func IsLoadComplete(webView WebView) bool {
	r, _, _ := _wkeIsLoadComplete.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func GetSource(webView WebView) string {
	r, _, _ := _wkeGetSource.Call(uintptr(webView))
	return StrFromPtr(r)
}

func Title(webView WebView) string {
	r, _, _ := _wkeTitle.Call(uintptr(webView))
	return StrFromPtr(r)
}

func TitleW(webView WebView) string {
	r, _, _ := _wkeTitleW.Call(uintptr(webView))
	return StrFromWPtr(r)
}

func Width(webView WebView) int {
	r, _, _ := _wkeWidth.Call(uintptr(webView))
	return int(r)
}

func Height(webView WebView) int {
	r, _, _ := _wkeHeight.Call(uintptr(webView))
	return int(r)
}

func ContentsWidth(webView WebView) int {
	r, _, _ := _wkeContentsWidth.Call(uintptr(webView))
	return int(r)
}

func ContentsHeight(webView WebView) int {
	r, _, _ := _wkeContentsHeight.Call(uintptr(webView))
	return int(r)
}

func SelectAll(webView WebView) {
	_wkeSelectAll.Call(uintptr(webView))
}

func Copy(webView WebView) {
	_wkeCopy.Call(uintptr(webView))
}

func Cut(webView WebView) {
	_wkeCut.Call(uintptr(webView))
}

func Paste(webView WebView) {
	_wkePaste.Call(uintptr(webView))
}

func Delete(webView WebView) {
	_wkeDelete.Call(uintptr(webView))
}

func CookieEnabled(webView WebView) bool {
	r, _, _ := _wkeCookieEnabled.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func MediaVolume(webView WebView) float32 {
	_, r2, _ := _wkeMediaVolume.Call(uintptr(webView))
	return F32FromPtr(r2)
}

func MouseEvent(webView WebView, message uint, x, y int, flags uint) bool {
	r, _, _ := _wkeMouseEvent.Call(uintptr(webView), uintptr(message), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r)
}

func ContextMenuEvent(webView WebView, x, y int, flags uint) bool {
	r, _, _ := _wkeContextMenuEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r)
}

func MouseWheel(webView WebView, x, y, delta int, flags uint) bool {
	r, _, _ := _wkeMouseWheel.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
	return BoolFromPtr(r)
}

func KeyUp(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r, _, _ := _wkeKeyUp.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r)
}

func KeyDown(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r, _, _ := _wkeKeyDown.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r)
}

func KeyPress(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r, _, _ := _wkeKeyPress.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r)
}

func Focus(webView WebView) {
	_wkeFocus.Call(uintptr(webView))
}

func Unfocus(webView WebView) {
	_wkeUnfocus.Call(uintptr(webView))
}

func GetCaret(webView WebView) Rect {
	// 下面是来自 GPT 的解释
	// var rect Rect 是声明了一个 Rect 类型的变量 rect，但它本身并没有在声明时完成赋值。rect 变量的赋值发生在调用 DLL 函数之后，但这个过程是隐式的。
	// 因为 Go 的调用约定和 C 的调用约定在这里是兼容的（假设你使用的是 cdecl 调用约定，并且你的 DLL 也是用 cdecl 编译的）。
	// 当你调用 proc.Call(uintptr(webView)) 时，DLL 中的 wkeGetCaretRect 函数会被执行，并且它的返回值（一个 Rect 结构体）会被直接拷贝到 Go 的调用栈中。
	// 由于 Go 调用栈上的 rect 变量在内存布局上与 C 的 Rect 结构体兼容，因此这个返回值会被自动“存储”到 rect 变量中。
	// 这个过程是自动的，你不需要显式地将返回值赋给 rect。一旦 proc.Call 返回，你就可以直接读取 rect 变量的值，它会包含从 DLL 函数返回的结构体数据。
	// 简单来说，rect 变量的赋值是在 proc.Call 执行期间隐式完成的，作为 C DLL 函数返回值的一部分。

	var rect Rect
	_wkeGetCaret.Call(uintptr(webView))
	return rect
}

func Awaken(webView WebView) {
	_wkeAwaken.Call(uintptr(webView))
}

func ZoomFactor(webView WebView) float32 {
	_, r2, _ := _wkeZoomFactor.Call(uintptr(webView))
	return F32FromPtr(r2)
}

func SetClientHandler(webView WebView, handler *ClientHandler) {
	_wkeSetClientHandler.Call(uintptr(webView), uintptr(unsafe.Pointer(handler)))
}

func GetClientHandler(webView WebView) *ClientHandler {
	r, _, _ := _wkeGetClientHandler.Call(uintptr(webView))
	return (*ClientHandler)(unsafe.Pointer(r))
}

func JsToString(es JsExecState, v JsValue) string {
	r, _, _ := _jsToString.Call(uintptr(es), uintptr(v))
	return StrFromPtr(r)
}

func JsToStringW(es JsExecState, v JsValue) string {
	r, _, _ := _jsToStringW.Call(uintptr(es), uintptr(v))
	return StrFromWPtr(r)
}

func SetViewSettings(webView WebView, settings *ViewSettings) {
	_wkeSetViewSettings.Call(uintptr(webView), uintptr(unsafe.Pointer(settings)))
}

func SetDebugConfig(webView WebView, debugString, param string) {
	_wkeSetDebugConfig.Call(uintptr(webView), StrToPtr(debugString), StrToPtr(param))
}

func ToString(wkeStr WkeString) string {
	r, _, _ := _wkeToString.Call(uintptr(unsafe.Pointer(&wkeStr)))
	return StrFromPtr(r)
}

func ToStringW(wkeStr WkeString) string {
	r, _, _ := _wkeToStringW.Call(uintptr(unsafe.Pointer(&wkeStr)))
	return StrFromWPtr(r)
}

func GetDebugConfig(webView WebView, debugString string) unsafe.Pointer {
	r, _, _ := _wkeGetDebugConfig.Call(uintptr(webView), StrToPtr(debugString))
	return unsafe.Pointer(r)
}

func IsInitialize() bool {
	r, _, _ := _wkeIsInitialize.Call()
	return BoolFromPtr(r)
}

func Finalize() {
	_wkeFinalize.Call()
}

func Update() {
	_wkeUpdate.Call()
}

func GetVersion() uint {
	r, _, _ := _wkeGetVersion.Call()
	return uint(r)
}

func GetVersionString() string {
	r, _, _ := _wkeGetVersionString.Call()
	return StrFromPtr(r)
}

func CreateWebView() WebView {
	r, _, _ := _wkeCreateWebView.Call()
	return WebView(r)
}

func DestroyWebView(webView WebView) {
	_wkeDestroyWebView.Call(uintptr(webView))
}

func SetMemoryCacheEnable(webView WebView, b bool) {
	_wkeSetMemoryCacheEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetMouseEnabled(webView WebView, b bool) {
	_wkeSetMouseEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetTouchEnabled(webView WebView, b bool) {
	_wkeSetTouchEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetSystemTouchEnabled(webView WebView, b bool) {
	_wkeSetSystemTouchEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetContextMenuEnabled(webView WebView, b bool) {
	_wkeSetContextMenuEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetNavigationToNewWindowEnable(webView WebView, b bool) {
	_wkeSetNavigationToNewWindowEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetCspCheckEnable(webView WebView, b bool) {
	_wkeSetCspCheckEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetNpapiPluginsEnabled(webView WebView, b bool) {
	_wkeSetNpapiPluginsEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetHeadlessEnabled(webView WebView, b bool) {
	_wkeSetHeadlessEnabled.Call(uintptr(webView), BoolToPtr(b))
}

func SetDragEnable(webView WebView, b bool) {
	_wkeSetDragEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetDragDropEnable(webView WebView, b bool) {
	_wkeSetDragDropEnable.Call(uintptr(webView), BoolToPtr(b))
}

func SetLanguage(webView WebView, language string) {
	_wkeSetLanguage.Call(uintptr(webView), StrToPtr(language))
}

func SetViewNetInterface(webView WebView, netInterface string) {
	_wkeSetViewNetInterface.Call(uintptr(webView), StrToPtr(netInterface))
}

func SetContextMenuItemShow(webView WebView, item MenuItemId, isShow bool) {
	_wkeSetContextMenuItemShow.Call(uintptr(webView), uintptr(item), BoolToPtr(isShow))
}

func SetProxy(proxy *Proxy) {
	_wkeSetProxy.Call(uintptr(unsafe.Pointer(proxy)))
}

func GetName(webView WebView) string {
	r, _, _ := _wkeGetName.Call(uintptr(webView))
	return StrFromPtr(r)
}

func IsTransparent(webView WebView) bool {
	r, _, _ := _wkeIsTransparent.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func GetUserAgent(webView WebView) string {
	r, _, _ := _wkeGetUserAgent.Call(uintptr(webView))
	return StrFromPtr(r)
}

func SetViewProxy(webView WebView, proxy *Proxy) {
	_wkeSetViewProxy.Call(uintptr(webView), uintptr(unsafe.Pointer(proxy)))
}

func SetName(webView WebView, name string) {
	_wkeSetName.Call(uintptr(webView), StrToPtr(name))
}

func SetHandle(webView WebView, wnd HWND) {
	_wkeSetHandle.Call(uintptr(webView), uintptr(wnd))
}

func SetTransparent(webView WebView, transparent bool) {
	_wkeSetTransparent.Call(uintptr(webView), BoolToPtr(transparent))
}

func SetUserAgent(webView WebView, userAgent string) {
	_wkeSetUserAgent.Call(uintptr(webView), StrToPtr(userAgent))
}

func SetUserAgentW(webView WebView, userAgent string) {
	_wkeSetUserAgentW.Call(uintptr(webView), StrToWPtr(userAgent))
}

func SetHandleOffset(webView WebView, x, y int) {
	_wkeSetHandleOffset.Call(uintptr(webView), uintptr(x), uintptr(y))
}

func ShowDevtools(webView WebView, path string, callback OnShowDevtoolsCallback) {
	_wkeShowDevtools.Call(uintptr(webView), StrToWPtr(path), CallbackToPtr(callback), 0)
}

func LoadW(webView WebView, url string) {
	_wkeLoadW.Call(uintptr(webView), StrToWPtr(url))
}

func LoadURL(webView WebView, url string) {
	_wkeLoadURL.Call(uintptr(webView), StrToPtr(url))
}

func LoadURLW(webView WebView, url string) {
	_wkeLoadURLW.Call(uintptr(webView), StrToWPtr(url))
}

func PostURL(wkeView WebView, url string, postData []byte) {
	_wkePostURL.Call(uintptr(wkeView), StrToPtr(url), uintptr(unsafe.Pointer(&postData[0])), uintptr(len(postData)))
}

func PostURLW(wkeView WebView, url string, postData []byte) {
	_wkePostURLW.Call(uintptr(wkeView), StrToWPtr(url), uintptr(unsafe.Pointer(&postData[0])), uintptr(len(postData)))
}

func LoadHTML(webView WebView, html string) {
	_wkeLoadHTML.Call(uintptr(webView), StrToPtr(html))
}

func LoadHtmlWithBaseUrl(webView WebView, html, baseUrl string) {
	_wkeLoadHtmlWithBaseUrl.Call(uintptr(webView), StrToPtr(html), StrToPtr(baseUrl))
}

func LoadHTMLW(webView WebView, html string) {
	_wkeLoadHTMLW.Call(uintptr(webView), StrToWPtr(html))
}

func LoadFile(webView WebView, filename string) {
	_wkeLoadFile.Call(uintptr(webView), StrToPtr(filename))
}

func LoadFileW(webView WebView, filename string) {
	_wkeLoadFileW.Call(uintptr(webView), StrToWPtr(filename))
}

func GetURL(webView WebView) string {
	r, _, _ := _wkeGetURL.Call(uintptr(webView))
	return StrFromPtr(r)
}

func GetFrameUrl(webView WebView, frameId WebFrameHandle) string {
	r, _, _ := _wkeGetFrameUrl.Call(uintptr(webView), uintptr(frameId))
	return StrFromPtr(r)
}

func IsLoading(webView WebView) bool {
	r, _, _ := _wkeIsLoading.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func IsLoadingSucceeded(webView WebView) bool {
	r, _, _ := _wkeIsLoadingSucceeded.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func IsLoadingFailed(webView WebView) bool {
	r, _, _ := _wkeIsLoadingFailed.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func IsLoadingCompleted(webView WebView) bool {
	r, _, _ := _wkeIsLoadingCompleted.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func IsDocumentReady(webView WebView) bool {
	r, _, _ := _wkeIsDocumentReady.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func StopLoading(webView WebView) {
	_wkeStopLoading.Call(uintptr(webView))
}

func Reload(webView WebView) {
	_wkeReload.Call(uintptr(webView))
}

func GoToOffset(webView WebView, offset int) {
	_wkeGoToOffset.Call(uintptr(webView), uintptr(offset))
}

func GoToIndex(webView WebView, index int) {
	_wkeGoToIndex.Call(uintptr(webView), uintptr(index))
}

func GetWebviewId(webView WebView) int {
	r, _, _ := _wkeGetWebviewId.Call(uintptr(webView))
	return int(r)
}

func IsWebviewAlive(id int) bool {
	r, _, _ := _wkeIsWebviewAlive.Call(uintptr(id))
	return BoolFromPtr(r)
}

func IsWebviewValid(webView WebView) bool {
	r, _, _ := _wkeIsWebviewValid.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func GetDocumentCompleteURL(webView WebView, frameId WebFrameHandle, partialURL string) string {
	r, _, _ := _wkeGetDocumentCompleteURL.Call(uintptr(webView), uintptr(frameId), StrToPtr(partialURL))
	return StrFromPtr(r)
}

func CreateMemBuf(webView WebView, bufPtr uintptr, length uintptr) *MemBuf {
	r, _, _ := _wkeCreateMemBuf.Call(uintptr(webView), bufPtr, length)
	return (*MemBuf)(unsafe.Pointer(r))
}

func FreeMemBuf(buf *MemBuf) {
	_wkeFreeMemBuf.Call(uintptr(unsafe.Pointer(buf)))
}

func GetTitle(webView WebView) string {
	r, _, _ := _wkeGetTitle.Call(uintptr(webView))
	return StrFromPtr(r)
}

func GetTitleW(webView WebView) string {
	r, _, _ := _wkeGetTitleW.Call(uintptr(webView))
	return StrFromPtr(r)
}

func Resize(webView WebView, w, h int) {
	_wkeResize.Call(uintptr(webView), uintptr(w), uintptr(h))
}

func GetWidth(webView WebView) int {
	r, _, _ := _wkeGetWidth.Call(uintptr(webView))
	return int(r)
}

func GetHeight(webView WebView) int {
	r, _, _ := _wkeGetHeight.Call(uintptr(webView))
	return int(r)
}

func GetContentWidth(webView WebView) int {
	r, _, _ := _wkeGetContentWidth.Call(uintptr(webView))
	return int(r)
}

func GetContentHeight(webView WebView) int {
	r, _, _ := _wkeGetContentHeight.Call(uintptr(webView))
	return int(r)
}

func SetDirty(webView WebView, dirty bool) {
	_wkeSetDirty.Call(uintptr(webView), BoolToPtr(dirty))
}

func IsDirty(webView WebView) bool {
	r, _, _ := _wkeIsDirty.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func AddDirtyArea(webView WebView, x, y, w, h int) {
	_wkeAddDirtyArea.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(w), uintptr(h))
}

func LayoutIfNeeded(webView WebView) {
	_wkeLayoutIfNeeded.Call(uintptr(webView))
}

func Paint2(webView WebView, bits uintptr, bufWid, bufHei, xDst, yDst, w, h, xSrc, ySrc int, bCopyAlpha bool) {
	_wkePaint2.Call(uintptr(webView), bits, uintptr(bufWid), uintptr(bufHei), uintptr(xDst), uintptr(yDst), uintptr(w), uintptr(h), uintptr(xSrc), uintptr(ySrc), BoolToPtr(bCopyAlpha))
}

func Paint(webView WebView, bits uintptr, pitch int) {
	_wkePaint.Call(uintptr(webView), bits, uintptr(pitch))
}

func RepaintIfNeeded(webView WebView) {
	_wkeRepaintIfNeeded.Call(uintptr(webView))
}

func GetViewDC(webView WebView) HDC {
	r, _, _ := _wkeGetViewDC.Call(uintptr(webView))
	return HDC(r)
}

func UnlockViewDC(webView WebView) {
	_wkeUnlockViewDC.Call(uintptr(webView))
}

func GetHostHWND(webView WebView) HWND {
	r, _, _ := _wkeGetHostHWND.Call(uintptr(webView))
	return HWND(r)
}

func CanGoBack(webView WebView) bool {
	r, _, _ := _wkeCanGoBack.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func GoBack(webView WebView) bool {
	r, _, _ := _wkeGoBack.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func CanGoForward(webView WebView) bool {
	r, _, _ := _wkeCanGoForward.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func GoForward(webView WebView) bool {
	r, _, _ := _wkeGoForward.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func NavigateAtIndex(webView WebView, index int) bool {
	r, _, _ := _wkeNavigateAtIndex.Call(uintptr(webView), uintptr(index))
	return BoolFromPtr(r)
}

func GetNavigateIndex(webView WebView) int {
	r, _, _ := _wkeGetNavigateIndex.Call(uintptr(webView))
	return int(r)
}

func EditorSelectAll(webView WebView) {
	_wkeEditorSelectAll.Call(uintptr(webView))
}

func EditorUnSelect(webView WebView) {
	_wkeEditorUnSelect.Call(uintptr(webView))
}

func EditorCopy(webView WebView) {
	_wkeEditorCopy.Call(uintptr(webView))
}

func EditorCut(webView WebView) {
	_wkeEditorCut.Call(uintptr(webView))
}

func EditorPaste(webView WebView) {
	_wkeEditorPaste.Call(uintptr(webView))
}

func EditorDelete(webView WebView) {
	_wkeEditorDelete.Call(uintptr(webView))
}

func EditorUndo(webView WebView) {
	_wkeEditorUndo.Call(uintptr(webView))
}

func EditorRedo(webView WebView) {
	_wkeEditorRedo.Call(uintptr(webView))
}

func GetCookieW(webView WebView) string {
	r, _, _ := _wkeGetCookieW.Call(uintptr(webView))
	return StrFromWPtr(r)
}

func GetCookie(webView WebView) string {
	r, _, _ := _wkeGetCookie.Call(uintptr(webView))
	return StrFromPtr(r)
}

func SetCookie(webView WebView, url, cookie string) {
	_wkeSetCookie.Call(uintptr(webView), StrToPtr(url), StrToPtr(cookie))
}

func VisitAllCookie(webView WebView, params uintptr, visitor CookieVisitor) {
	_wkeVisitAllCookie.Call(uintptr(webView), params, CallbackToPtr(visitor))
}

func PerformCookieCommand(webView WebView, command CookieCommand) {
	_wkePerformCookieCommand.Call(uintptr(webView), uintptr(command))
}

func SetCookieEnabled(webView WebView, enable bool) {
	_wkeSetCookieEnabled.Call(uintptr(webView), BoolToPtr(enable))
}

func IsCookieEnabled(webView WebView) bool {
	r, _, _ := _wkeIsCookieEnabled.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func SetCookieJarPath(webView WebView, path string) {
	_wkeSetCookieJarPath.Call(uintptr(webView), StrToWPtr(path))
}

func SetCookieJarFullPath(webView WebView, path string) {
	_wkeSetCookieJarFullPath.Call(uintptr(webView), StrToWPtr(path))
}

func ClearCookie(webView WebView) {
	_wkeClearCookie.Call(uintptr(webView))
}

func SetLocalStorageFullPath(webView WebView, path string) {
	_wkeSetLocalStorageFullPath.Call(uintptr(webView), StrToWPtr(path))
}

func AddPluginDirectory(webView WebView, path string) {
	_wkeAddPluginDirectory.Call(uintptr(webView), StrToWPtr(path))
}

func SetMediaVolume(webView WebView, volume float32) {
	_wkeSetMediaVolume.Call(uintptr(webView), F32ToPtr(volume))
}

func GetMediaVolume(webView WebView) float32 {
	_, r, _ := _wkeGetMediaVolume.Call(uintptr(webView))
	return F32FromPtr(r)
}

func FireMouseEvent(webView WebView, message uint, x, y int, flags uint) bool {
	r, _, _ := _wkeFireMouseEvent.Call(uintptr(webView), uintptr(message), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r)
}

func FireContextMenuEvent(webView WebView, x, y int, flags uint) bool {
	r, _, _ := _wkeFireContextMenuEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(flags))
	return BoolFromPtr(r)
}

func FireMouseWheelEvent(webView WebView, x, y, delta int, flags uint) bool {
	r, _, _ := _wkeFireMouseWheelEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
	return BoolFromPtr(r)
}

func FireKeyUpEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r, _, _ := _wkeFireKeyUpEvent.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r)
}

func FireKeyDownEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	r, _, _ := _wkeFireKeyDownEvent.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r)
}

func FireKeyPressEvent(webView WebView, charCode, flags uint, systemKey bool) bool {
	r, _, _ := _wkeFireKeyPressEvent.Call(uintptr(webView), uintptr(charCode), uintptr(flags), BoolToPtr(systemKey))
	return BoolFromPtr(r)
}

func FireWindowsMessage(webView WebView, hWnd HWND, message uint, wParam WPARAM, lParam LPARAM, result *LRESULT) bool {
	r, _, _ := _wkeFireWindowsMessage.Call(uintptr(webView), uintptr(hWnd), uintptr(message), uintptr(wParam), uintptr(lParam), uintptr(unsafe.Pointer(result)))
	return BoolFromPtr(r)
}

func SetFocus(webView WebView) {
	_wkeSetFocus.Call(uintptr(webView))
}

func KillFocus(webView WebView) {
	_wkeKillFocus.Call(uintptr(webView))
}

// 获取编辑框的那个游标的位置
func GetCaretRect(webView WebView) Rect {
	// 下面是来自 GPT 的解释
	// var rect Rect 是声明了一个 Rect 类型的变量 rect，但它本身并没有在声明时完成赋值。rect 变量的赋值发生在调用 DLL 函数之后，但这个过程是隐式的。
	// 因为 Go 的调用约定和 C 的调用约定在这里是兼容的（假设你使用的是 cdecl 调用约定，并且你的 DLL 也是用 cdecl 编译的）。
	// 当你调用 proc.Call(uintptr(webView)) 时，DLL 中的 wkeGetCaretRect 函数会被执行，并且它的返回值（一个 Rect 结构体）会被直接拷贝到 Go 的调用栈中。
	// 由于 Go 调用栈上的 rect 变量在内存布局上与 C 的 Rect 结构体兼容，因此这个返回值会被自动“存储”到 rect 变量中。
	// 这个过程是自动的，你不需要显式地将返回值赋给 rect。一旦 proc.Call 返回，你就可以直接读取 rect 变量的值，它会包含从 DLL 函数返回的结构体数据。
	// 简单来说，rect 变量的赋值是在 proc.Call 执行期间隐式完成的，作为 C DLL 函数返回值的一部分。

	var rect Rect
	_wkeGetCaretRect.Call(uintptr(webView))
	return rect
}

func GetCaretRect2(webView WebView) *Rect {
	r, _, _ := _wkeGetCaretRect2.Call(uintptr(webView))
	return (*Rect)(unsafe.Pointer(r))
}

func RunJS(webView WebView, script string) JsValue {
	r, _, _ := _wkeRunJS.Call(uintptr(webView), StrToPtr(script))
	return JsValue(r)
}

func RunJSW(webView WebView, script string) JsValue {
	r, _, _ := _wkeRunJSW.Call(uintptr(webView), StrToWPtr(script))
	return JsValue(r)
}

func GlobalExec(webView WebView) JsExecState {
	r, _, _ := _wkeGlobalExec.Call(uintptr(webView))
	return JsExecState(r)
}

func GetGlobalExecByFrame(webView WebView, frameId WebFrameHandle) JsExecState {
	r, _, _ := _wkeGetGlobalExecByFrame.Call(uintptr(webView), uintptr(frameId))
	return JsExecState(r)
}

func Sleep(webView WebView) {
	_wkeSleep.Call(uintptr(webView))
}

func Wake(webView WebView) {
	_wkeWake.Call(uintptr(webView))
}

func IsAwake(webView WebView) bool {
	r, _, _ := _wkeIsAwake.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func SetZoomFactor(webView WebView, factor float32) {
	_wkeSetZoomFactor.Call(uintptr(webView), F32ToPtr(factor))
}

func GetZoomFactor(webView WebView) float32 {
	_, r, _ := _wkeGetZoomFactor.Call(uintptr(webView))
	return F32FromPtr(r)
}

func EnableHighDPISupport() {
	_wkeEnableHighDPISupport.Call()
}

func SetEditable(webView WebView, editable bool) {
	_wkeSetEditable.Call(uintptr(webView), BoolToPtr(editable))
}

func GetString(wkeStr WkeString) string {
	r, _, _ := _wkeGetString.Call(uintptr(wkeStr))
	return StrFromPtr(r)
}

func GetStringW(wkeStr WkeString) string {
	r, _, _ := _wkeGetStringW.Call(uintptr(wkeStr))
	return StrFromWPtr(r)
}

func SetString(wkeStr WkeString, str string) {
	byts := []byte(str)
	_wkeSetString.Call(uintptr(wkeStr), uintptr(unsafe.Pointer(&byts[0])), uintptr(len(byts)))
}

func SetStringWithoutNullTermination(wkeStr WkeString, str string) {
	byts := []byte(str)
	_wkeSetStringWithoutNullTermination.Call(uintptr(wkeStr), uintptr(unsafe.Pointer(&byts[0])), uintptr(len(byts)))
}

func SetStringW(wkeStr WkeString, str string) {
	wchar := StrToWChar(str)
	_wkeSetStringW.Call(uintptr(wkeStr), uintptr(unsafe.Pointer(&wchar[0])), uintptr(len(wchar)))
}

func CreateString(str string) WkeString {
	byts := []byte(str)
	r, _, _ := _wkeCreateString.Call(uintptr(unsafe.Pointer(&byts[0])), uintptr(len(byts)))
	return WkeString(r)
}

func CreateStringW(str string) WkeString {
	wchar := StrToWChar(str)
	r, _, _ := _wkeCreateStringW.Call(uintptr(unsafe.Pointer(&wchar[0])), uintptr(len(wchar)))
	return WkeString(r)
}

func CreateStringWithoutNullTermination(str string) WkeString {
	byts := []byte(str)
	r, _, _ := _wkeCreateStringWithoutNullTermination.Call(uintptr(unsafe.Pointer(&byts[0])), uintptr(len(byts)))
	return WkeString(r)
}

func GetStringLen(wkeStr WkeString) uintptr {
	r, _, _ := _wkeGetStringLen.Call(uintptr(wkeStr))
	return r
}

func DeleteString(wkeStr WkeString) {
	_wkeDeleteString.Call(uintptr(wkeStr))
}

func GetWebViewForCurrentContext() WebView {
	r, _, _ := _wkeGetWebViewForCurrentContext.Call()
	return WebView(r)
}

func SetUserKeyValue(webView WebView, key string, value any) {
	_wkeSetUserKeyValue.Call(uintptr(webView), StrToPtr(key), uintptr(unsafe.Pointer(&value)))
}

func GetUserKeyValue(webView WebView, key string) any {
	r, _, _ := _wkeGetUserKeyValue.Call(uintptr(webView), StrToPtr(key))
	return any(unsafe.Pointer(r))
}

func GetCursorInfoType(webView WebView) int {
	r, _, _ := _wkeGetCursorInfoType.Call(uintptr(webView))
	return int(r)
}

func SetCursorInfoType(webView WebView, type_ int) {
	_wkeSetCursorInfoType.Call(uintptr(webView), uintptr(type_))
}

func SetDragFiles(webView WebView, clintPos, screenPos *Point, files *WkeString, filesCount int) {
	_wkeSetDragFiles.Call(uintptr(webView), uintptr(unsafe.Pointer(clintPos)), uintptr(unsafe.Pointer(screenPos)), uintptr(unsafe.Pointer(files)), uintptr(filesCount))
}

func SetDeviceParameter(webView WebView, device, paramStr string, paramInt int, paramFloat float32) {
	_wkeSetDeviceParameter.Call(uintptr(webView), StrToPtr(device), StrToPtr(paramStr), uintptr(paramInt), F32ToPtr(paramFloat))
}

func GetTempCallbackInfo(webView WebView) *TempCallbackInfo {
	r, _, _ := _wkeGetTempCallbackInfo.Call(uintptr(webView))
	return (*TempCallbackInfo)(unsafe.Pointer(r))
}

func OnCaretChanged(webView WebView, callback CaretChangedCallback) {
	_wkeOnCaretChanged.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnMouseOverUrlChanged(webView WebView, callback TitleChangedCallback) {
	_wkeOnMouseOverUrlChanged.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnTitleChanged(webView WebView, callback TitleChangedCallback) {
	_wkeOnTitleChanged.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnURLChanged(webView WebView, callback URLChangedCallback) {
	_wkeOnURLChanged.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnURLChanged2(webView WebView, callback URLChangedCallback2) {
	_wkeOnURLChanged2.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnPaintUpdated(webView WebView, callback PaintUpdatedCallback) {
	_wkeOnPaintUpdated.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnPaintBitUpdated(webView WebView, callback PaintBitUpdatedCallback) {
	_wkeOnPaintBitUpdated.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnAlertBox(webView WebView, callback AlertBoxCallback) {
	_wkeOnAlertBox.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnConfirmBox(webView WebView, callback ConfirmBoxCallback) {
	_wkeOnConfirmBox.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnPromptBox(webView WebView, callback PromptBoxCallback) {
	_wkeOnPromptBox.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnNavigation(webView WebView, callback NavigationCallback) {
	_wkeOnNavigation.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnCreateView(webView WebView, callback CreateViewCallback) {
	_wkeOnCreateView.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnDocumentReady(webView WebView, callback DocumentReadyCallback) {
	_wkeOnDocumentReady.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnDocumentReady2(webView WebView, callback DocumentReady2Callback) {
	_wkeOnDocumentReady2.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadingFinish(webView WebView, callback LoadingFinishCallback) {
	_wkeOnLoadingFinish.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnDownload(webView WebView, callback DownloadCallback) {
	_wkeOnDownload.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnDownload2(webView WebView, callback Download2Callback) {
	_wkeOnDownload2.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnConsole(webView WebView, callback ConsoleCallback) {
	_wkeOnConsole.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func SetUIThreadCallback(webView WebView, callback CallUiThread) {
	_wkeSetUIThreadCallback.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlBegin(webView WebView, callback LoadUrlBeginCallback) {
	_wkeOnLoadUrlBegin.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlEnd(webView WebView, callback LoadUrlEndCallback) {
	_wkeOnLoadUrlEnd.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlHeadersReceived(webView WebView, callback LoadUrlHeadersReceivedCallback) {
	_wkeOnLoadUrlHeadersReceived.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlFinish(webView WebView, callback LoadUrlFinishCallback) {
	_wkeOnLoadUrlFinish.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnLoadUrlFail(webView WebView, callback LoadUrlFailCallback) {
	_wkeOnLoadUrlFail.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnDidCreateScriptContext(webView WebView, callback DidCreateScriptContextCallback) {
	_wkeOnDidCreateScriptContext.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnWillReleaseScriptContext(webView WebView, callback WillReleaseScriptContextCallback) {
	_wkeOnWillReleaseScriptContext.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnWindowClosing(webWindow WebView, callback WindowClosingCallback) {
	_wkeOnWindowClosing.Call(uintptr(webWindow), CallbackToPtr(callback), 0)
}

func OnWindowDestroy(webWindow WebView, callback WindowDestroyCallback) {
	_wkeOnWindowDestroy.Call(uintptr(webWindow), CallbackToPtr(callback), 0)
}

func OnDraggableRegionsChanged(webView WebView, callback DraggableRegionsChangedCallback) {
	_wkeOnDraggableRegionsChanged.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnWillMediaLoad(webView WebView, callback WillMediaLoadCallback) {
	_wkeOnWillMediaLoad.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnStartDragging(webView WebView, callback StartDraggingCallback) {
	_wkeOnStartDragging.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnPrint(webView WebView, callback OnPrintCallback) {
	_wkeOnPrint.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func Screenshot(webView WebView, settings *ScreenshotSettings, callback OnScreenshotCallback) {
	_wkeScreenshot.Call(uintptr(webView), uintptr(unsafe.Pointer(settings)), CallbackToPtr(callback))
}

func OnOtherLoad(webView WebView, callback OnOtherLoadCallback) {
	_wkeOnOtherLoad.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnContextMenuItemClick(webView WebView, callback OnContextMenuItemClickCallback) {
	_wkeOnContextMenuItemClick.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func IsProcessingUserGesture(webView WebView) bool {
	r, _, _ := _wkeIsProcessingUserGesture.Call(uintptr(webView))
	return BoolFromPtr(r)
}

func NetSetMIMEType(jobPtr NetJob, type_ string) {
	_wkeNetSetMIMEType.Call(uintptr(jobPtr), StrToPtr(type_))
}

func NetGetMIMEType(jobPtr NetJob, mime *WkeString) string {
	r, _, _ := _wkeNetGetMIMEType.Call(uintptr(jobPtr), uintptr(unsafe.Pointer(mime)))
	return StrFromPtr(r)
}

func NetGetReferrer(jobPtr NetJob) string {
	r, _, _ := _wkeNetGetReferrer.Call(uintptr(jobPtr))
	return StrFromPtr(r)
}

func NetSetHTTPHeaderField(jobPtr NetJob, key, value string, response bool) {
	_wkeNetSetHTTPHeaderField.Call(uintptr(jobPtr), StrToWPtr(key), StrToWPtr(value), BoolToPtr(response))
}

func NetGetHTTPHeaderField(jobPtr NetJob, key string) string {
	r, _, _ := _wkeNetGetHTTPHeaderField.Call(uintptr(jobPtr), StrToPtr(key))
	return StrFromPtr(r)
}

func NetGetHTTPHeaderFieldFromResponse(jobPtr NetJob, key string) string {
	r, _, _ := _wkeNetGetHTTPHeaderFieldFromResponse.Call(uintptr(jobPtr), StrToPtr(key))
	return StrFromPtr(r)
}

func NetSetData(jobPtr NetJob, buf []byte) {
	_wkeNetSetData.Call(uintptr(jobPtr), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func NetHookRequest(jobPtr NetJob) {
	_wkeNetHookRequest.Call(uintptr(jobPtr))
}

func NetGetRequestMethod(jobPtr NetJob) RequestType {
	r, _, _ := _wkeNetGetRequestMethod.Call(uintptr(jobPtr))
	return RequestType(r)
}

func NetContinueJob(jobPtr NetJob) {
	_wkeNetContinueJob.Call(uintptr(jobPtr))
}

func NetGetUrlByJob(jobPtr NetJob) string {
	r, _, _ := _wkeNetGetUrlByJob.Call(uintptr(jobPtr))
	return StrFromPtr(r)
}

func NetGetRawHttpHead(jobPtr NetJob) *Slist {
	r, _, _ := _wkeNetGetRawHttpHead.Call(uintptr(jobPtr))
	return (*Slist)(unsafe.Pointer(r))
}

func NetGetRawResponseHead(jobPtr NetJob) *Slist {
	r, _, _ := _wkeNetGetRawResponseHead.Call(uintptr(jobPtr))
	return (*Slist)(unsafe.Pointer(r))
}

func NetCancelRequest(jobPtr NetJob) {
	_wkeNetCancelRequest.Call(uintptr(jobPtr))
}

func NetHoldJobToAsynCommit(jobPtr NetJob) bool {
	r, _, _ := _wkeNetHoldJobToAsynCommit.Call(uintptr(jobPtr))
	return BoolFromPtr(r)
}

func NetOnResponse(webView WebView, callback NetResponseCallback) {
	_wkeNetOnResponse.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func NetGetFavicon(webView WebView, callback OnNetGetFaviconCallback) int {
	r, _, _ := _wkeNetGetFavicon.Call(uintptr(webView), CallbackToPtr(callback), 0)
	return int(r)
}

func NetCreateWebUrlRequest(url, method, mime string) WebUrlRequestPtr {
	r, _, _ := _wkeNetCreateWebUrlRequest.Call(StrToPtr(url), StrToPtr(method), StrToPtr(mime))
	return WebUrlRequestPtr(unsafe.Pointer(r))
}

func NetCreateWebUrlRequest2(request blinkWebURLRequestPtr) WebUrlRequestPtr {
	r, _, _ := _wkeNetCreateWebUrlRequest2.Call(uintptr(request))
	return WebUrlRequestPtr(unsafe.Pointer(r))
}

func NetCopyWebUrlRequest(jobPtr NetJob, needExtraData bool) blinkWebURLRequestPtr {
	r, _, _ := _wkeNetCopyWebUrlRequest.Call(uintptr(jobPtr), BoolToPtr(needExtraData))
	return blinkWebURLRequestPtr(unsafe.Pointer(r))
}

func NetDeleteBlinkWebURLRequestPtr(request blinkWebURLRequestPtr) {
	_wkeNetDeleteBlinkWebURLRequestPtr.Call(uintptr(request))
}

func NetAddHTTPHeaderFieldToUrlRequest(request WebUrlRequestPtr, name, value string) {
	_wkeNetAddHTTPHeaderFieldToUrlRequest.Call(uintptr(unsafe.Pointer(request)), StrToPtr(name), StrToPtr(value))
}

func NetStartUrlRequest(webView WebView, request WebUrlRequestPtr, callbacks *UrlRequestCallbacks) int {

	type Callbacks struct {
		WillRedirectCallback       uintptr
		DidReceiveResponseCallback uintptr
		DidReceiveDataCallback     uintptr
		DidFailCallback            uintptr
		DidFinishLoadingCallback   uintptr
	}

	cbs := Callbacks{
		WillRedirectCallback:       CallbackToPtr(callbacks.WillRedirectCallback),
		DidReceiveResponseCallback: CallbackToPtr(callbacks.DidReceiveResponseCallback),
		DidReceiveDataCallback:     CallbackToPtr(callbacks.DidReceiveDataCallback),
		DidFailCallback:            CallbackToPtr(callbacks.DidFailCallback),
		DidFinishLoadingCallback:   CallbackToPtr(callbacks.DidFinishLoadingCallback),
	}

	r, _, _ := _wkeNetStartUrlRequest.Call(uintptr(webView), uintptr(unsafe.Pointer(request)), 0, uintptr(unsafe.Pointer(&cbs)))
	return int(r)
}

func NetGetHttpStatusCode(response WebUrlResponsePtr) int {
	r, _, _ := _wkeNetGetHttpStatusCode.Call(uintptr(unsafe.Pointer(response)))
	return int(r)
}

func NetGetExpectedContentLength(response WebUrlResponsePtr) int64 {
	r, _, _ := _wkeNetGetExpectedContentLength.Call(uintptr(unsafe.Pointer(response)))
	return int64(r)
}

func NetGetResponseUrl(response WebUrlResponsePtr) string {
	r, _, _ := _wkeNetGetResponseUrl.Call(uintptr(unsafe.Pointer(response)))
	return StrFromPtr(r)
}

func NetCancelWebUrlRequest(requestId int) {
	_wkeNetCancelWebUrlRequest.Call(uintptr(requestId))
}

func NetGetPostBody(jobPtr NetJob) *PostBodyElements {
	r, _, _ := _wkeNetGetPostBody.Call(uintptr(jobPtr))
	return (*PostBodyElements)(unsafe.Pointer(r))
}

func NetFreePostBodyElements(elements *PostBodyElements) {
	_wkeNetFreePostBodyElements.Call(uintptr(unsafe.Pointer(elements)))
}

func NetCreatePostBodyElements(webView WebView, length uintptr) *PostBodyElements {
	r, _, _ := _wkeNetCreatePostBodyElements.Call(uintptr(webView), length)
	return (*PostBodyElements)(unsafe.Pointer(r))
}

func NetCreatePostBodyElement(webView WebView) *PostBodyElement {
	r, _, _ := _wkeNetCreatePostBodyElement.Call(uintptr(webView))
	return (*PostBodyElement)(unsafe.Pointer(r))
}

func NetFreePostBodyElement(element *PostBodyElement) {
	_wkeNetFreePostBodyElement.Call(uintptr(unsafe.Pointer(element)))
}

func PopupDialogAndDownload(webviewHandle WebView, dialogOpt *DialogOptions, expectedContentLength uintptr, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	// TODO: 需要验证 struct 内的callback是否可不转uintptr
	r, _, _ := _wkePopupDialogAndDownload.Call(uintptr(webviewHandle), uintptr(unsafe.Pointer(dialogOpt)), expectedContentLength, StrToPtr(url), StrToPtr(mime), StrToPtr(disposition), uintptr(job), uintptr(unsafe.Pointer(dataBind)), uintptr(unsafe.Pointer(callbackBind)))
	return DownloadOpt(r)
}

func DownloadByPath(webviewHandle WebView, dialogOpt *DialogOptions, path string, expectedContentLength uintptr, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	r, _, _ := _wkeDownloadByPath.Call(uintptr(webviewHandle), uintptr(unsafe.Pointer(dialogOpt)), StrToWPtr(path), expectedContentLength, StrToPtr(url), StrToPtr(mime), StrToPtr(disposition), uintptr(job), uintptr(unsafe.Pointer(dataBind)), uintptr(unsafe.Pointer(callbackBind)))
	return DownloadOpt(r)
}

func IsMainFrame(webView WebView, frameId WebFrameHandle) bool {
	r, _, _ := _wkeIsMainFrame.Call(uintptr(webView), uintptr(frameId))
	return BoolFromPtr(r)
}

func IsWebRemoteFrame(webView WebView, frameId WebFrameHandle) bool {
	r, _, _ := _wkeIsWebRemoteFrame.Call(uintptr(webView), uintptr(frameId))
	return BoolFromPtr(r)
}

func WebFrameGetMainFrame(webView WebView) WebFrameHandle {
	r, _, _ := _wkeWebFrameGetMainFrame.Call(uintptr(webView))
	return WebFrameHandle(r)
}

func RunJsByFrame(webView WebView, frameId WebFrameHandle, script string, isInClosure bool) JsValue {
	r, _, _ := _wkeRunJsByFrame.Call(uintptr(webView), uintptr(frameId), StrToPtr(script), BoolToPtr(isInClosure))
	return JsValue(r)
}

func InsertCSSByFrame(webView WebView, frameId WebFrameHandle, cssText string) {
	_wkeInsertCSSByFrame.Call(uintptr(webView), uintptr(frameId), StrToPtr(cssText))
}

func WebFrameGetMainWorldScriptContext(webView WebView, webFrameId WebFrameHandle, contextOut *V8ContextPtr) {
	_wkeWebFrameGetMainWorldScriptContext.Call(uintptr(webView), uintptr(webFrameId), uintptr(unsafe.Pointer(contextOut)))
}

func GetBlinkMainThreadIsolate() V8Isolate {
	r, _, _ := _wkeGetBlinkMainThreadIsolate.Call()
	return V8Isolate(r)
}

func CreateWebWindow(typ WindowType, parent HWND, x, y, width, height int) WebView {
	r, _, _ := _wkeCreateWebWindow.Call(uintptr(typ), uintptr(parent), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return WebView(r)
}

func CreateWebCustomWindow(info *WindowCreateInfo) WebView {
	r, _, _ := _wkeCreateWebCustomWindow.Call(uintptr(unsafe.Pointer(info)))
	return WebView(r)
}

func DestroyWebWindow(webWindow WebView) {
	_wkeDestroyWebWindow.Call(uintptr(webWindow))
}

func GetWindowHandle(webWindow WebView) HWND {
	r, _, _ := _wkeGetWindowHandle.Call(uintptr(webWindow))
	return HWND(r)
}

func ShowWindow(webWindow WebView, show bool) {
	_wkeShowWindow.Call(uintptr(webWindow), BoolToPtr(show))
}

func EnableWindow(webWindow WebView, enable bool) {
	_wkeEnableWindow.Call(uintptr(webWindow), BoolToPtr(enable))
}

func MoveWindow(webWindow WebView, x, y, width, height int) {
	_wkeMoveWindow.Call(uintptr(webWindow), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func MoveToCenter(webWindow WebView) {
	_wkeMoveToCenter.Call(uintptr(webWindow))
}

func ResizeWindow(webWindow WebView, width, height int) {
	_wkeResizeWindow.Call(uintptr(webWindow), uintptr(width), uintptr(height))
}

func DragTargetDragEnter(webView WebView, webDragData *WebDragData, clientPoint, screenPoint *Point, operationsAllowed WebDragOperationsMask, modifiers int) WebDragOperationsMask {
	r, _, _ := _wkeDragTargetDragEnter.Call(uintptr(webView), uintptr(unsafe.Pointer(webDragData)), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(operationsAllowed), uintptr(modifiers))
	return WebDragOperationsMask(r)
}

func DragTargetDragOver(webView WebView, clientPoint, screenPoint *Point, operationsAllowed WebDragOperationsMask, modifiers int) WebDragOperationsMask {
	r, _, _ := _wkeDragTargetDragOver.Call(uintptr(webView), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(operationsAllowed), uintptr(modifiers))
	return WebDragOperationsMask(r)
}

func DragTargetDragLeave(webView WebView) {
	_wkeDragTargetDragLeave.Call(uintptr(webView))
}

func DragTargetDrop(webView WebView, clientPoint, screenPoint *Point, modifiers int) {
	_wkeDragTargetDrop.Call(uintptr(webView), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(modifiers))
}

func DragTargetEnd(webView WebView, clientPoint, screenPoint *Point, operation WebDragOperationsMask) {
	_wkeDragTargetEnd.Call(uintptr(webView), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(operation))
}

func UtilSetUiCallback(callback UiThreadPostTaskCallback) {
	_wkeUtilSetUiCallback.Call(CallbackToPtr(callback))
}

func UtilSerializeToMHTML(webView WebView) string {
	r, _, _ := _wkeUtilSerializeToMHTML.Call(uintptr(webView))
	return StrFromPtr(r)
}

func UtilPrintToPdf(webView WebView, frameId WebFrameHandle, settings *PrintSettings) *PdfDatas {
	r, _, _ := _wkeUtilPrintToPdf.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)))
	return (*PdfDatas)(unsafe.Pointer(r))
}

func PrintToBitmap(webView WebView, frameId WebFrameHandle, settings *ScreenshotSettings) *MemBuf {
	r, _, _ := _wkePrintToBitmap.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)))
	return (*MemBuf)(unsafe.Pointer(r))
}

func UtilRelasePrintPdfDatas(datas *PdfDatas) {
	_wkeUtilRelasePrintPdfDatas.Call(uintptr(unsafe.Pointer(datas)))
}

func SetWindowTitle(webWindow WebView, title string) {
	_wkeSetWindowTitle.Call(uintptr(webWindow), StrToPtr(title))
}

func SetWindowTitleW(webWindow WebView, title string) {
	_wkeSetWindowTitleW.Call(uintptr(webWindow), StrToWPtr(title))
}

func NodeOnCreateProcess(webView WebView, callback NodeOnCreateProcessCallback) {
	_wkeNodeOnCreateProcess.Call(uintptr(webView), CallbackToPtr(callback), 0)
}

func OnPluginFind(webView WebView, mime string, callback OnPluginFindCallback) {
	_wkeOnPluginFind.Call(uintptr(webView), StrToPtr(mime), CallbackToPtr(callback), 0)
}

func AddNpapiPlugin(webView WebView, initializeFunc, getEntryPointsFunc, shutdownFunc any) {
	_wkeAddNpapiPlugin.Call(uintptr(webView), CallbackToPtr(initializeFunc), CallbackToPtr(getEntryPointsFunc), CallbackToPtr(shutdownFunc))
}

func PluginListBuilderAddPlugin(builder uintptr, name, description, fileName string) {
	_wkePluginListBuilderAddPlugin.Call(builder, StrToPtr(name), StrToPtr(description), StrToPtr(fileName))
}

func PluginListBuilderAddMediaTypeToLastPlugin(builder uintptr, name, description string) {
	_wkePluginListBuilderAddMediaTypeToLastPlugin.Call(builder, StrToPtr(name), StrToPtr(description))
}

func PluginListBuilderAddFileExtensionToLastMediaType(builder uintptr, fileExtension string) {
	_wkePluginListBuilderAddFileExtensionToLastMediaType.Call(builder, StrToPtr(fileExtension))
}

func GetWebViewByNData(ndata uintptr) WebView {
	r, _, _ := _wkeGetWebViewByNData.Call(ndata)
	return WebView(r)
}

func RegisterEmbedderCustomElement(webView WebView, frameId WebFrameHandle, name string, options, outResult uintptr) bool {
	r, _, _ := _wkeRegisterEmbedderCustomElement.Call(uintptr(webView), uintptr(frameId), StrToPtr(name), options, outResult)
	return BoolFromPtr(r)
}

func SetMediaPlayerFactory(webView WebView, factory MediaPlayerFactory, callback OnIsMediaPlayerSupportsMIMEType) {
	_wkeSetMediaPlayerFactory.Call(uintptr(webView), CallbackToPtr(factory), CallbackToPtr(callback))
}

func GetContentAsMarkup(webView WebView, frame WebFrameHandle, size *uintptr) string {
	r, _, _ := _wkeGetContentAsMarkup.Call(uintptr(webView), uintptr(frame), uintptr(unsafe.Pointer(size)))
	return StrFromPtr(r)
}

func UtilDecodeURLEscape(url string) string {
	r, _, _ := _wkeUtilDecodeURLEscape.Call(StrToPtr(url))
	return StrFromPtr(r)
}

func UtilEncodeURLEscape(url string) string {
	r, _, _ := _wkeUtilEncodeURLEscape.Call(StrToPtr(url))
	return StrFromPtr(r)
}

func UtilBase64Encode(str string) string {
	r, _, _ := _wkeUtilBase64Encode.Call(StrToPtr(str))
	return StrFromPtr(r)
}

func UtilBase64Decode(str string) string {
	r, _, _ := _wkeUtilBase64Decode.Call(StrToPtr(str))
	return StrFromPtr(r)
}

func UtilCreateV8Snapshot(str string) *MemBuf {
	r, _, _ := _wkeUtilCreateV8Snapshot.Call(StrToPtr(str))
	return (*MemBuf)(unsafe.Pointer(r))
}

func RunMessageLoop() {
	_wkeRunMessageLoop.Call()
}

func SaveMemoryCache(webView WebView) {
	_wkeSaveMemoryCache.Call(uintptr(webView))
}

func JsBindFunction(name string, fn wkeJsNativeFunction, argCount uint) {
	_wkeJsBindFunction.Call(StrToPtr(name), CallbackToPtr(fn), 0, uintptr(argCount))
}

func JsBindGetter(name string, fn wkeJsNativeFunction) {
	_wkeJsBindGetter.Call(StrToPtr(name), CallbackToPtr(fn), 0)
}

func JsBindSetter(name string, fn wkeJsNativeFunction) {
	_wkeJsBindSetter.Call(StrToPtr(name), CallbackToPtr(fn), 0)
}

func JsArgCount(es JsExecState) int {
	r, _, _ := _jsArgCount.Call(uintptr(es))
	return int(r)
}

func JsArgType(es JsExecState, argIdx int) JsType {
	r, _, _ := _jsArgType.Call(uintptr(es), uintptr(argIdx))
	return JsType(r)
}

func JsArg(es JsExecState, argIdx int) JsValue {
	r, _, _ := _jsArg.Call(uintptr(es), uintptr(argIdx))
	return JsValue(r)
}

func JsTypeOf(v JsValue) JsType {
	r, _, _ := _jsTypeOf.Call(uintptr(v))
	return JsType(r)
}

func JsIsNumber(v JsValue) bool {
	r, _, _ := _jsIsNumber.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsString(v JsValue) bool {
	r, _, _ := _jsIsString.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsBoolean(v JsValue) bool {
	r, _, _ := _jsIsBoolean.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsObject(v JsValue) bool {
	r, _, _ := _jsIsObject.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsFunction(v JsValue) bool {
	r, _, _ := _jsIsFunction.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsUndefined(v JsValue) bool {
	r, _, _ := _jsIsUndefined.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsNull(v JsValue) bool {
	r, _, _ := _jsIsNull.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsArray(v JsValue) bool {
	r, _, _ := _jsIsArray.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsTrue(v JsValue) bool {
	r, _, _ := _jsIsTrue.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsIsFalse(v JsValue) bool {
	r, _, _ := _jsIsFalse.Call(uintptr(v))
	return BoolFromPtr(r)
}

func JsToInt(es JsExecState, v JsValue) int {
	r, _, _ := _jsToInt.Call(uintptr(es), uintptr(v))
	return int(r)
}

func JsToFloat(es JsExecState, v JsValue) float32 {
	_, r, _ := _jsToFloat.Call(uintptr(es), uintptr(v))
	return F32FromPtr(r)
}

func JsToDouble(es JsExecState, v JsValue) float64 {
	_, r, _ := _jsToDouble.Call(uintptr(es), uintptr(v))
	return F64FromPtr(r)
}

func JsToDoubleString(es JsExecState, v JsValue) string {
	r, _, _ := _jsToDoubleString.Call(uintptr(es), uintptr(v))
	return StrFromPtr(r)
}

func JsToBoolean(es JsExecState, v JsValue) bool {
	r, _, _ := _jsToBoolean.Call(uintptr(es), uintptr(v))
	return BoolFromPtr(r)
}

func JsArrayBuffer(es JsExecState, buffer []byte) JsValue {
	r, _, _ := _jsArrayBuffer.Call(uintptr(es), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
	return JsValue(r)
}

func JsGetArrayBuffer(es JsExecState, value JsValue) *MemBuf {
	r, _, _ := _jsGetArrayBuffer.Call(uintptr(es), uintptr(value))
	return (*MemBuf)(unsafe.Pointer(r))
}

func JsToTempString(es JsExecState, v JsValue) string {
	r, _, _ := _jsToTempString.Call(uintptr(es), uintptr(v))
	return StrFromPtr(r)
}

func JsToTempStringW(es JsExecState, v JsValue) string {
	r, _, _ := _jsToTempStringW.Call(uintptr(es), uintptr(v))
	return StrFromWPtr(r)
}

func JsToV8Value(es JsExecState, v JsValue) v8ValuePtr {
	r, _, _ := _jsToV8Value.Call(uintptr(es), uintptr(v))
	return v8ValuePtr(r)
}

func JsInt(n int) JsValue {
	r, _, _ := _jsInt.Call(uintptr(n))
	return JsValue(r)
}

func JsFloat(f float32) JsValue {
	r, _, _ := _jsFloat.Call(F32ToPtr(f))
	return JsValue(r)
}

func JsDouble(d float64) JsValue {
	r, _, _ := _jsDouble.Call(F64ToPtr(d))
	return JsValue(r)
}

func JsDoubleString(str string) JsValue {
	r, _, _ := _jsDoubleString.Call(StrToPtr(str))
	return JsValue(r)
}

func JsBoolean(b bool) JsValue {
	r, _, _ := _jsBoolean.Call(BoolToPtr(b))
	return JsValue(r)
}

func JsUndefined() JsValue {
	r, _, _ := _jsUndefined.Call()
	return JsValue(r)
}

func JsNull() JsValue {
	r, _, _ := _jsNull.Call()
	return JsValue(r)
}

func JsTrue() JsValue {
	r, _, _ := _jsTrue.Call()
	return JsValue(r)
}

func JsFalse() JsValue {
	r, _, _ := _jsFalse.Call()
	return JsValue(r)
}

func JsString(es JsExecState, str string) JsValue {
	r, _, _ := _jsString.Call(uintptr(es), StrToPtr(str))
	return JsValue(r)
}

func JsStringW(es JsExecState, str string) JsValue {
	r, _, _ := _jsStringW.Call(uintptr(es), StrToWPtr(str))
	return JsValue(r)
}

func JsEmptyObject(es JsExecState) JsValue {
	r, _, _ := _jsEmptyObject.Call(uintptr(es))
	return JsValue(r)
}

func JsEmptyArray(es JsExecState) JsValue {
	r, _, _ := _jsEmptyArray.Call(uintptr(es))
	return JsValue(r)
}

func JsObject(es JsExecState, obj *JsData) JsValue {
	// TODO: 检测 struct 内的 callback 是否可以不用转换
	r, _, _ := _jsObject.Call(uintptr(es), uintptr(unsafe.Pointer(obj)))
	return JsValue(r)
}

func JsFunction(es JsExecState, obj *JsData) JsValue {
	r, _, _ := _jsFunction.Call(uintptr(es), uintptr(unsafe.Pointer(obj)))
	return JsValue(r)
}

func JsGetData(es JsExecState, object JsValue) *JsData {
	r, _, _ := _jsGetData.Call(uintptr(es), uintptr(object))
	return (*JsData)(unsafe.Pointer(r))
}

func JsGet(es JsExecState, object JsValue, prop string) JsValue {
	r, _, _ := _jsGet.Call(uintptr(es), uintptr(object), StrToPtr(prop))
	return JsValue(r)
}

func JsSet(es JsExecState, object JsValue, prop string, v JsValue) {
	_jsSet.Call(uintptr(es), uintptr(object), StrToPtr(prop), uintptr(v))
}

func JsGetAt(es JsExecState, object JsValue, index int) JsValue {
	r, _, _ := _jsGetAt.Call(uintptr(es), uintptr(object), uintptr(index))
	return JsValue(r)
}

func JsSetAt(es JsExecState, object JsValue, index int, v JsValue) {
	_jsSetAt.Call(uintptr(es), uintptr(object), uintptr(index), uintptr(v))
}

func JsGetKeys(es JsExecState, object JsValue) *JsKeys {
	r, _, _ := _jsGetKeys.Call(uintptr(es), uintptr(object))
	return (*JsKeys)(unsafe.Pointer(r))
}

func JsIsJsValueValid(es JsExecState, object JsValue) bool {
	r, _, _ := _jsIsJsValueValid.Call(uintptr(es), uintptr(object))
	return BoolFromPtr(r)
}

func JsIsValidExecState(es JsExecState) bool {
	r, _, _ := _jsIsValidExecState.Call(uintptr(es))
	return BoolFromPtr(r)
}

func JsDeleteObjectProp(es JsExecState, object JsValue, prop string) {
	_jsDeleteObjectProp.Call(uintptr(es), uintptr(object), StrToPtr(prop))
}

func JsGetLength(es JsExecState, object JsValue) int {
	r, _, _ := _jsGetLength.Call(uintptr(es), uintptr(object))
	return int(r)
}

func JsSetLength(es JsExecState, object JsValue, length int) {
	_jsSetLength.Call(uintptr(es), uintptr(object), uintptr(length))
}

func JsGlobalObject(es JsExecState) JsValue {
	r, _, _ := _jsGlobalObject.Call(uintptr(es))
	return JsValue(r)
}

func JsGetWebView(es JsExecState) WebView {
	r, _, _ := _jsGetWebView.Call(uintptr(es))
	return WebView(r)
}

func JsEval(es JsExecState, str string) JsValue {
	r, _, _ := _jsEval.Call(uintptr(es), StrToPtr(str))
	return JsValue(r)
}

func JsEvalW(es JsExecState, str string) JsValue {
	r, _, _ := _jsEvalW.Call(uintptr(es), StrToWPtr(str))
	return JsValue(r)
}

func JsEvalExW(es JsExecState, str string, isInClosure bool) JsValue {
	r, _, _ := _jsEvalExW.Call(uintptr(es), StrToWPtr(str), BoolToPtr(isInClosure))
	return JsValue(r)
}

func JsCall(es JsExecState, funcObj, thisObject JsValue, args []JsValue) JsValue {
	r, _, _ := _jsCall.Call(uintptr(es), uintptr(funcObj), uintptr(thisObject), uintptr(unsafe.Pointer(&args[0])), uintptr(len(args)))
	return JsValue(r)
}

func JsCallGlobal(es JsExecState, funcObj JsValue, args []JsValue) JsValue {
	r, _, _ := _jsCallGlobal.Call(uintptr(es), uintptr(funcObj), uintptr(unsafe.Pointer(&args[0])), uintptr(len(args)))
	return JsValue(r)
}

func JsGetGlobal(es JsExecState, prop string) JsValue {
	r, _, _ := _jsGetGlobal.Call(uintptr(es), StrToPtr(prop))
	return JsValue(r)
}

func JsSetGlobal(es JsExecState, prop string, v JsValue) {
	_jsSetGlobal.Call(uintptr(es), StrToPtr(prop), uintptr(v))
}

func JsGC() {
	_jsGC.Call()
}

func JsAddRef(es JsExecState, val JsValue) bool {
	r, _, _ := _jsAddRef.Call(uintptr(es), uintptr(val))
	return BoolFromPtr(r)
}

func JsReleaseRef(es JsExecState, val JsValue) bool {
	r, _, _ := _jsReleaseRef.Call(uintptr(es), uintptr(val))
	return BoolFromPtr(r)
}

func JsGetLastErrorIfException(es JsExecState) *JsExceptionInfo {
	r, _, _ := _jsGetLastErrorIfException.Call(uintptr(es))
	return (*JsExceptionInfo)(unsafe.Pointer(r))
}

func JsThrowException(es JsExecState, exception string) JsValue {
	r, _, _ := _jsThrowException.Call(uintptr(es), StrToPtr(exception))
	return JsValue(r)
}

func JsGetCallstack(es JsExecState) string {
	r, _, _ := _jsGetCallstack.Call(uintptr(es))
	return StrFromPtr(r)
}

func Init() {
	_wkeInit.Call()
}

func Initialize() {
	_wkeInitialize.Call()
}

func InitializeEx(settings *Settings) {
	_wkeInitializeEx.Call(uintptr(unsafe.Pointer(settings)))
}
