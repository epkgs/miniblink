package wke

import (
	"unsafe"

	"github.com/epkgs/miniblink/internal/lib"
)

var LoadLibrary = lib.LoadLibrary
var LoadEmbedLibrary = lib.LoadEmbedLibrary

var _wkeJsNativeFunction = lib.LazyFunc[wkeJsNativeFunction]{Name: "wkeJsNativeFunction"}
var _wkeShutdown = lib.LazyFunc[wkeShutdown]{Name: "wkeShutdown"}
var _wkeShutdownForDebug = lib.LazyFunc[wkeShutdownForDebug]{Name: "wkeShutdownForDebug"}
var _wkeVersion = lib.LazyFunc[wkeVersion]{Name: "wkeVersion"}
var _wkeVersionString = lib.LazyFunc[wkeVersionString]{Name: "wkeVersionString"}
var _wkeGC = lib.LazyFunc[wkeGC]{Name: "wkeGC"}
var _wkeSetResourceGc = lib.LazyFunc[wkeSetResourceGc]{Name: "wkeSetResourceGc"}
var _wkeSetFileSystem = lib.LazyFunc[wkeSetFileSystem]{Name: "wkeSetFileSystem"}
var _wkeWebViewName = lib.LazyFunc[wkeWebViewName]{Name: "wkeWebViewName"}
var _wkeSetWebViewName = lib.LazyFunc[wkeSetWebViewName]{Name: "wkeSetWebViewName"}
var _wkeIsLoaded = lib.LazyFunc[wkeIsLoaded]{Name: "wkeIsLoaded"}
var _wkeIsLoadFailed = lib.LazyFunc[wkeIsLoadFailed]{Name: "wkeIsLoadFailed"}
var _wkeIsLoadComplete = lib.LazyFunc[wkeIsLoadComplete]{Name: "wkeIsLoadComplete"}
var _wkeGetSource = lib.LazyFunc[wkeGetSource]{Name: "wkeGetSource"}
var _wkeTitle = lib.LazyFunc[wkeTitle]{Name: "wkeTitle"}
var _wkeTitleW = lib.LazyFunc[wkeTitleW]{Name: "wkeTitleW"}
var _wkeWidth = lib.LazyFunc[wkeWidth]{Name: "wkeWidth"}
var _wkeHeight = lib.LazyFunc[wkeHeight]{Name: "wkeHeight"}
var _wkeContentsWidth = lib.LazyFunc[wkeContentsWidth]{Name: "wkeContentsWidth"}
var _wkeContentsHeight = lib.LazyFunc[wkeContentsHeight]{Name: "wkeContentsHeight"}
var _wkeSelectAll = lib.LazyFunc[wkeSelectAll]{Name: "wkeSelectAll"}
var _wkeCopy = lib.LazyFunc[wkeCopy]{Name: "wkeCopy"}
var _wkeCut = lib.LazyFunc[wkeCut]{Name: "wkeCut"}
var _wkePaste = lib.LazyFunc[wkePaste]{Name: "wkePaste"}
var _wkeDelete = lib.LazyFunc[wkeDelete]{Name: "wkeDelete"}
var _wkeCookieEnabled = lib.LazyFunc[wkeCookieEnabled]{Name: "wkeCookieEnabled"}
var _wkeMediaVolume = lib.LazyFunc[wkeMediaVolume]{Name: "wkeMediaVolume"}
var _wkeMouseEvent = lib.LazyFunc[wkeMouseEvent]{Name: "wkeMouseEvent"}
var _wkeContextMenuEvent = lib.LazyFunc[wkeContextMenuEvent]{Name: "wkeContextMenuEvent"}
var _wkeMouseWheel = lib.LazyFunc[wkeMouseWheel]{Name: "wkeMouseWheel"}
var _wkeKeyUp = lib.LazyFunc[wkeKeyUp]{Name: "wkeKeyUp"}
var _wkeKeyDown = lib.LazyFunc[wkeKeyDown]{Name: "wkeKeyDown"}
var _wkeKeyPress = lib.LazyFunc[wkeKeyPress]{Name: "wkeKeyPress"}
var _wkeFocus = lib.LazyFunc[wkeFocus]{Name: "wkeFocus"}
var _wkeUnfocus = lib.LazyFunc[wkeUnfocus]{Name: "wkeUnfocus"}
var _wkeGetCaret = lib.LazyFunc[wkeGetCaret]{Name: "wkeGetCaret"}
var _wkeAwaken = lib.LazyFunc[wkeAwaken]{Name: "wkeAwaken"}
var _wkeZoomFactor = lib.LazyFunc[wkeZoomFactor]{Name: "wkeZoomFactor"}
var _wkeSetClientHandler = lib.LazyFunc[wkeSetClientHandler]{Name: "wkeSetClientHandler"}
var _wkeGetClientHandler = lib.LazyFunc[wkeGetClientHandler]{Name: "wkeGetClientHandler"}
var _jsToString = lib.LazyFunc[jsToString]{Name: "jsToString"}
var _jsToStringW = lib.LazyFunc[jsToStringW]{Name: "jsToStringW"}
var _wkeSetViewSettings = lib.LazyFunc[wkeSetViewSettings]{Name: "wkeSetViewSettings"}
var _wkeSetDebugConfig = lib.LazyFunc[wkeSetDebugConfig]{Name: "wkeSetDebugConfig"}
var _wkeToString = lib.LazyFunc[wkeToString]{Name: "wkeToString"}
var _wkeToStringW = lib.LazyFunc[wkeToStringW]{Name: "wkeToStringW"}
var _wkeGetDebugConfig = lib.LazyFunc[wkeGetDebugConfig]{Name: "wkeGetDebugConfig"}
var _wkeIsInitialize = lib.LazyFunc[wkeIsInitialize]{Name: "wkeIsInitialize"}
var _wkeFinalize = lib.LazyFunc[wkeFinalize]{Name: "wkeFinalize"}
var _wkeUpdate = lib.LazyFunc[wkeUpdate]{Name: "wkeUpdate"}
var _wkeGetVersion = lib.LazyFunc[wkeGetVersion]{Name: "wkeGetVersion"}
var _wkeGetVersionString = lib.LazyFunc[wkeGetVersionString]{Name: "wkeGetVersionString"}
var _wkeCreateWebView = lib.LazyFunc[wkeCreateWebView]{Name: "wkeCreateWebView"}
var _wkeDestroyWebView = lib.LazyFunc[wkeDestroyWebView]{Name: "wkeDestroyWebView"}
var _wkeSetMemoryCacheEnable = lib.LazyFunc[wkeSetMemoryCacheEnable]{Name: "wkeSetMemoryCacheEnable"}
var _wkeSetMouseEnabled = lib.LazyFunc[wkeSetMouseEnabled]{Name: "wkeSetMouseEnabled"}
var _wkeSetTouchEnabled = lib.LazyFunc[wkeSetTouchEnabled]{Name: "wkeSetTouchEnabled"}
var _wkeSetSystemTouchEnabled = lib.LazyFunc[wkeSetSystemTouchEnabled]{Name: "wkeSetSystemTouchEnabled"}
var _wkeSetContextMenuEnabled = lib.LazyFunc[wkeSetContextMenuEnabled]{Name: "wkeSetContextMenuEnabled"}
var _wkeSetNavigationToNewWindowEnable = lib.LazyFunc[wkeSetNavigationToNewWindowEnable]{Name: "wkeSetNavigationToNewWindowEnable"}
var _wkeSetCspCheckEnable = lib.LazyFunc[wkeSetCspCheckEnable]{Name: "wkeSetCspCheckEnable"}
var _wkeSetNpapiPluginsEnabled = lib.LazyFunc[wkeSetNpapiPluginsEnabled]{Name: "wkeSetNpapiPluginsEnabled"}
var _wkeSetHeadlessEnabled = lib.LazyFunc[wkeSetHeadlessEnabled]{Name: "wkeSetHeadlessEnabled"}
var _wkeSetDragEnable = lib.LazyFunc[wkeSetDragEnable]{Name: "wkeSetDragEnable"}
var _wkeSetDragDropEnable = lib.LazyFunc[wkeSetDragDropEnable]{Name: "wkeSetDragDropEnable"}
var _wkeSetLanguage = lib.LazyFunc[wkeSetLanguage]{Name: "wkeSetLanguage"}
var _wkeSetViewNetInterface = lib.LazyFunc[wkeSetViewNetInterface]{Name: "wkeSetViewNetInterface"}
var _wkeSetContextMenuItemShow = lib.LazyFunc[wkeSetContextMenuItemShow]{Name: "wkeSetContextMenuItemShow"}
var _wkeSetProxy = lib.LazyFunc[wkeSetProxy]{Name: "wkeSetProxy"}
var _wkeGetName = lib.LazyFunc[wkeGetName]{Name: "wkeGetName"}
var _wkeIsTransparent = lib.LazyFunc[wkeIsTransparent]{Name: "wkeIsTransparent"}
var _wkeGetUserAgent = lib.LazyFunc[wkeGetUserAgent]{Name: "wkeGetUserAgent"}
var _wkeSetViewProxy = lib.LazyFunc[wkeSetViewProxy]{Name: "wkeSetViewProxy"}
var _wkeSetName = lib.LazyFunc[wkeSetName]{Name: "wkeSetName"}
var _wkeSetHandle = lib.LazyFunc[wkeSetHandle]{Name: "wkeSetHandle"}
var _wkeSetTransparent = lib.LazyFunc[wkeSetTransparent]{Name: "wkeSetTransparent"}
var _wkeSetUserAgent = lib.LazyFunc[wkeSetUserAgent]{Name: "wkeSetUserAgent"}
var _wkeSetUserAgentW = lib.LazyFunc[wkeSetUserAgentW]{Name: "wkeSetUserAgentW"}
var _wkeSetHandleOffset = lib.LazyFunc[wkeSetHandleOffset]{Name: "wkeSetHandleOffset"}
var _wkeShowDevtools = lib.LazyFunc[wkeShowDevtools]{Name: "wkeShowDevtools"}
var _wkeLoadW = lib.LazyFunc[wkeLoadW]{Name: "wkeLoadW"}
var _wkeLoadURL = lib.LazyFunc[wkeLoadURL]{Name: "wkeLoadURL"}
var _wkeLoadURLW = lib.LazyFunc[wkeLoadURLW]{Name: "wkeLoadURLW"}
var _wkePostURL = lib.LazyFunc[wkePostURL]{Name: "wkePostURL"}
var _wkePostURLW = lib.LazyFunc[wkePostURLW]{Name: "wkePostURLW"}
var _wkeLoadHTML = lib.LazyFunc[wkeLoadHTML]{Name: "wkeLoadHTML"}
var _wkeLoadHtmlWithBaseUrl = lib.LazyFunc[wkeLoadHtmlWithBaseUrl]{Name: "wkeLoadHtmlWithBaseUrl"}
var _wkeLoadHTMLW = lib.LazyFunc[wkeLoadHTMLW]{Name: "wkeLoadHTMLW"}
var _wkeLoadFile = lib.LazyFunc[wkeLoadFile]{Name: "wkeLoadFile"}
var _wkeLoadFileW = lib.LazyFunc[wkeLoadFileW]{Name: "wkeLoadFileW"}
var _wkeGetURL = lib.LazyFunc[wkeGetURL]{Name: "wkeGetURL"}
var _wkeGetFrameUrl = lib.LazyFunc[wkeGetFrameUrl]{Name: "wkeGetFrameUrl"}
var _wkeIsLoading = lib.LazyFunc[wkeIsLoading]{Name: "wkeIsLoading"}
var _wkeIsLoadingSucceeded = lib.LazyFunc[wkeIsLoadingSucceeded]{Name: "wkeIsLoadingSucceeded"}
var _wkeIsLoadingFailed = lib.LazyFunc[wkeIsLoadingFailed]{Name: "wkeIsLoadingFailed"}
var _wkeIsLoadingCompleted = lib.LazyFunc[wkeIsLoadingCompleted]{Name: "wkeIsLoadingCompleted"}
var _wkeIsDocumentReady = lib.LazyFunc[wkeIsDocumentReady]{Name: "wkeIsDocumentReady"}
var _wkeStopLoading = lib.LazyFunc[wkeStopLoading]{Name: "wkeStopLoading"}
var _wkeReload = lib.LazyFunc[wkeReload]{Name: "wkeReload"}
var _wkeGoToOffset = lib.LazyFunc[wkeGoToOffset]{Name: "wkeGoToOffset"}
var _wkeGoToIndex = lib.LazyFunc[wkeGoToIndex]{Name: "wkeGoToIndex"}
var _wkeGetWebviewId = lib.LazyFunc[wkeGetWebviewId]{Name: "wkeGetWebviewId"}
var _wkeIsWebviewAlive = lib.LazyFunc[wkeIsWebviewAlive]{Name: "wkeIsWebviewAlive"}
var _wkeIsWebviewValid = lib.LazyFunc[wkeIsWebviewValid]{Name: "wkeIsWebviewValid"}
var _wkeGetDocumentCompleteURL = lib.LazyFunc[wkeGetDocumentCompleteURL]{Name: "wkeGetDocumentCompleteURL"}
var _wkeCreateMemBuf = lib.LazyFunc[wkeCreateMemBuf]{Name: "wkeCreateMemBuf"}
var _wkeFreeMemBuf = lib.LazyFunc[wkeFreeMemBuf]{Name: "wkeFreeMemBuf"}
var _wkeGetTitle = lib.LazyFunc[wkeGetTitle]{Name: "wkeGetTitle"}
var _wkeGetTitleW = lib.LazyFunc[wkeGetTitleW]{Name: "wkeGetTitleW"}
var _wkeResize = lib.LazyFunc[wkeResize]{Name: "wkeResize"}
var _wkeGetWidth = lib.LazyFunc[wkeGetWidth]{Name: "wkeGetWidth"}
var _wkeGetHeight = lib.LazyFunc[wkeGetHeight]{Name: "wkeGetHeight"}
var _wkeGetContentWidth = lib.LazyFunc[wkeGetContentWidth]{Name: "wkeGetContentWidth"}
var _wkeGetContentHeight = lib.LazyFunc[wkeGetContentHeight]{Name: "wkeGetContentHeight"}
var _wkeSetDirty = lib.LazyFunc[wkeSetDirty]{Name: "wkeSetDirty"}
var _wkeIsDirty = lib.LazyFunc[wkeIsDirty]{Name: "wkeIsDirty"}
var _wkeAddDirtyArea = lib.LazyFunc[wkeAddDirtyArea]{Name: "wkeAddDirtyArea"}
var _wkeLayoutIfNeeded = lib.LazyFunc[wkeLayoutIfNeeded]{Name: "wkeLayoutIfNeeded"}
var _wkePaint2 = lib.LazyFunc[wkePaint2]{Name: "wkePaint2"}
var _wkePaint = lib.LazyFunc[wkePaint]{Name: "wkePaint"}
var _wkeRepaintIfNeeded = lib.LazyFunc[wkeRepaintIfNeeded]{Name: "wkeRepaintIfNeeded"}
var _wkeGetViewDC = lib.LazyFunc[wkeGetViewDC]{Name: "wkeGetViewDC"}
var _wkeUnlockViewDC = lib.LazyFunc[wkeUnlockViewDC]{Name: "wkeUnlockViewDC"}
var _wkeGetHostHWND = lib.LazyFunc[wkeGetHostHWND]{Name: "wkeGetHostHWND"}
var _wkeCanGoBack = lib.LazyFunc[wkeCanGoBack]{Name: "wkeCanGoBack"}
var _wkeGoBack = lib.LazyFunc[wkeGoBack]{Name: "wkeGoBack"}
var _wkeCanGoForward = lib.LazyFunc[wkeCanGoForward]{Name: "wkeCanGoForward"}
var _wkeGoForward = lib.LazyFunc[wkeGoForward]{Name: "wkeGoForward"}
var _wkeNavigateAtIndex = lib.LazyFunc[wkeNavigateAtIndex]{Name: "wkeNavigateAtIndex"}
var _wkeGetNavigateIndex = lib.LazyFunc[wkeGetNavigateIndex]{Name: "wkeGetNavigateIndex"}
var _wkeEditorSelectAll = lib.LazyFunc[wkeEditorSelectAll]{Name: "wkeEditorSelectAll"}
var _wkeEditorUnSelect = lib.LazyFunc[wkeEditorUnSelect]{Name: "wkeEditorUnSelect"}
var _wkeEditorCopy = lib.LazyFunc[wkeEditorCopy]{Name: "wkeEditorCopy"}
var _wkeEditorCut = lib.LazyFunc[wkeEditorCut]{Name: "wkeEditorCut"}
var _wkeEditorPaste = lib.LazyFunc[wkeEditorPaste]{Name: "wkeEditorPaste"}
var _wkeEditorDelete = lib.LazyFunc[wkeEditorDelete]{Name: "wkeEditorDelete"}
var _wkeEditorUndo = lib.LazyFunc[wkeEditorUndo]{Name: "wkeEditorUndo"}
var _wkeEditorRedo = lib.LazyFunc[wkeEditorRedo]{Name: "wkeEditorRedo"}
var _wkeGetCookieW = lib.LazyFunc[wkeGetCookieW]{Name: "wkeGetCookieW"}
var _wkeGetCookie = lib.LazyFunc[wkeGetCookie]{Name: "wkeGetCookie"}
var _wkeSetCookie = lib.LazyFunc[wkeSetCookie]{Name: "wkeSetCookie"}
var _wkeVisitAllCookie = lib.LazyFunc[wkeVisitAllCookie]{Name: "wkeVisitAllCookie"}
var _wkePerformCookieCommand = lib.LazyFunc[wkePerformCookieCommand]{Name: "wkePerformCookieCommand"}
var _wkeSetCookieEnabled = lib.LazyFunc[wkeSetCookieEnabled]{Name: "wkeSetCookieEnabled"}
var _wkeIsCookieEnabled = lib.LazyFunc[wkeIsCookieEnabled]{Name: "wkeIsCookieEnabled"}
var _wkeSetCookieJarPath = lib.LazyFunc[wkeSetCookieJarPath]{Name: "wkeSetCookieJarPath"}
var _wkeSetCookieJarFullPath = lib.LazyFunc[wkeSetCookieJarFullPath]{Name: "wkeSetCookieJarFullPath"}
var _wkeClearCookie = lib.LazyFunc[wkeClearCookie]{Name: "wkeClearCookie"}
var _wkeSetLocalStorageFullPath = lib.LazyFunc[wkeSetLocalStorageFullPath]{Name: "wkeSetLocalStorageFullPath"}
var _wkeAddPluginDirectory = lib.LazyFunc[wkeAddPluginDirectory]{Name: "wkeAddPluginDirectory"}
var _wkeSetMediaVolume = lib.LazyFunc[wkeSetMediaVolume]{Name: "wkeSetMediaVolume"}
var _wkeGetMediaVolume = lib.LazyFunc[wkeGetMediaVolume]{Name: "wkeGetMediaVolume"}
var _wkeFireMouseEvent = lib.LazyFunc[wkeFireMouseEvent]{Name: "wkeFireMouseEvent"}
var _wkeFireContextMenuEvent = lib.LazyFunc[wkeFireContextMenuEvent]{Name: "wkeFireContextMenuEvent"}
var _wkeFireMouseWheelEvent = lib.LazyFunc[wkeFireMouseWheelEvent]{Name: "wkeFireMouseWheelEvent"}
var _wkeFireKeyUpEvent = lib.LazyFunc[wkeFireKeyUpEvent]{Name: "wkeFireKeyUpEvent"}
var _wkeFireKeyDownEvent = lib.LazyFunc[wkeFireKeyDownEvent]{Name: "wkeFireKeyDownEvent"}
var _wkeFireKeyPressEvent = lib.LazyFunc[wkeFireKeyPressEvent]{Name: "wkeFireKeyPressEvent"}
var _wkeFireWindowsMessage = lib.LazyFunc[wkeFireWindowsMessage]{Name: "wkeFireWindowsMessage"}
var _wkeSetFocus = lib.LazyFunc[wkeSetFocus]{Name: "wkeSetFocus"}
var _wkeKillFocus = lib.LazyFunc[wkeKillFocus]{Name: "wkeKillFocus"}
var _wkeGetCaretRect = lib.LazyFunc[wkeGetCaretRect]{Name: "wkeGetCaretRect"}
var _wkeGetCaretRect2 = lib.LazyFunc[wkeGetCaretRect2]{Name: "wkeGetCaretRect2"}
var _wkeRunJS = lib.LazyFunc[wkeRunJS]{Name: "wkeRunJS"}
var _wkeRunJSW = lib.LazyFunc[wkeRunJSW]{Name: "wkeRunJSW"}
var _wkeGlobalExec = lib.LazyFunc[wkeGlobalExec]{Name: "wkeGlobalExec"}
var _wkeGetGlobalExecByFrame = lib.LazyFunc[wkeGetGlobalExecByFrame]{Name: "wkeGetGlobalExecByFrame"}
var _wkeSleep = lib.LazyFunc[wkeSleep]{Name: "wkeSleep"}
var _wkeWake = lib.LazyFunc[wkeWake]{Name: "wkeWake"}
var _wkeIsAwake = lib.LazyFunc[wkeIsAwake]{Name: "wkeIsAwake"}
var _wkeSetZoomFactor = lib.LazyFunc[wkeSetZoomFactor]{Name: "wkeSetZoomFactor"}
var _wkeGetZoomFactor = lib.LazyFunc[wkeGetZoomFactor]{Name: "wkeGetZoomFactor"}
var _wkeEnableHighDPISupport = lib.LazyFunc[wkeEnableHighDPISupport]{Name: "wkeEnableHighDPISupport"}
var _wkeSetEditable = lib.LazyFunc[wkeSetEditable]{Name: "wkeSetEditable"}
var _wkeGetString = lib.LazyFunc[wkeGetString]{Name: "wkeGetString"}
var _wkeGetStringW = lib.LazyFunc[wkeGetStringW]{Name: "wkeGetStringW"}
var _wkeSetString = lib.LazyFunc[wkeSetString]{Name: "wkeSetString"}
var _wkeSetStringWithoutNullTermination = lib.LazyFunc[wkeSetStringWithoutNullTermination]{Name: "wkeSetStringWithoutNullTermination"}
var _wkeSetStringW = lib.LazyFunc[wkeSetStringW]{Name: "wkeSetStringW"}
var _wkeCreateString = lib.LazyFunc[wkeCreateString]{Name: "wkeCreateString"}
var _wkeCreateStringW = lib.LazyFunc[wkeCreateStringW]{Name: "wkeCreateStringW"}
var _wkeCreateStringWithoutNullTermination = lib.LazyFunc[wkeCreateStringWithoutNullTermination]{Name: "wkeCreateStringWithoutNullTermination"}
var _wkeGetStringLen = lib.LazyFunc[wkeGetStringLen]{Name: "wkeGetStringLen"}
var _wkeDeleteString = lib.LazyFunc[wkeDeleteString]{Name: "wkeDeleteString"}
var _wkeGetWebViewForCurrentContext = lib.LazyFunc[wkeGetWebViewForCurrentContext]{Name: "wkeGetWebViewForCurrentContext"}
var _wkeSetUserKeyValue = lib.LazyFunc[wkeSetUserKeyValue]{Name: "wkeSetUserKeyValue"}
var _wkeGetUserKeyValue = lib.LazyFunc[wkeGetUserKeyValue]{Name: "wkeGetUserKeyValue"}
var _wkeGetCursorInfoType = lib.LazyFunc[wkeGetCursorInfoType]{Name: "wkeGetCursorInfoType"}
var _wkeSetCursorInfoType = lib.LazyFunc[wkeSetCursorInfoType]{Name: "wkeSetCursorInfoType"}
var _wkeSetDragFiles = lib.LazyFunc[wkeSetDragFiles]{Name: "wkeSetDragFiles"}
var _wkeSetDeviceParameter = lib.LazyFunc[wkeSetDeviceParameter]{Name: "wkeSetDeviceParameter"}
var _wkeGetTempCallbackInfo = lib.LazyFunc[wkeGetTempCallbackInfo]{Name: "wkeGetTempCallbackInfo"}
var _wkeOnCaretChanged = lib.LazyFunc[wkeOnCaretChanged]{Name: "wkeOnCaretChanged"}
var _wkeOnMouseOverUrlChanged = lib.LazyFunc[wkeOnMouseOverUrlChanged]{Name: "wkeOnMouseOverUrlChanged"}
var _wkeOnTitleChanged = lib.LazyFunc[wkeOnTitleChanged]{Name: "wkeOnTitleChanged"}
var _wkeOnURLChanged = lib.LazyFunc[wkeOnURLChanged]{Name: "wkeOnURLChanged"}
var _wkeOnURLChanged2 = lib.LazyFunc[wkeOnURLChanged2]{Name: "wkeOnURLChanged2"}
var _wkeOnPaintUpdated = lib.LazyFunc[wkeOnPaintUpdated]{Name: "wkeOnPaintUpdated"}
var _wkeOnPaintBitUpdated = lib.LazyFunc[wkeOnPaintBitUpdated]{Name: "wkeOnPaintBitUpdated"}
var _wkeOnAlertBox = lib.LazyFunc[wkeOnAlertBox]{Name: "wkeOnAlertBox"}
var _wkeOnConfirmBox = lib.LazyFunc[wkeOnConfirmBox]{Name: "wkeOnConfirmBox"}
var _wkeOnPromptBox = lib.LazyFunc[wkeOnPromptBox]{Name: "wkeOnPromptBox"}
var _wkeOnNavigation = lib.LazyFunc[wkeOnNavigation]{Name: "wkeOnNavigation"}
var _wkeOnCreateView = lib.LazyFunc[wkeOnCreateView]{Name: "wkeOnCreateView"}
var _wkeOnDocumentReady = lib.LazyFunc[wkeOnDocumentReady]{Name: "wkeOnDocumentReady"}
var _wkeOnDocumentReady2 = lib.LazyFunc[wkeOnDocumentReady2]{Name: "wkeOnDocumentReady2"}
var _wkeOnLoadingFinish = lib.LazyFunc[wkeOnLoadingFinish]{Name: "wkeOnLoadingFinish"}
var _wkeOnDownload = lib.LazyFunc[wkeOnDownload]{Name: "wkeOnDownload"}
var _wkeOnDownload2 = lib.LazyFunc[wkeOnDownload2]{Name: "wkeOnDownload2"}
var _wkeOnConsole = lib.LazyFunc[wkeOnConsole]{Name: "wkeOnConsole"}
var _wkeSetUIThreadCallback = lib.LazyFunc[wkeSetUIThreadCallback]{Name: "wkeSetUIThreadCallback"}
var _wkeOnLoadUrlBegin = lib.LazyFunc[wkeOnLoadUrlBegin]{Name: "wkeOnLoadUrlBegin"}
var _wkeOnLoadUrlEnd = lib.LazyFunc[wkeOnLoadUrlEnd]{Name: "wkeOnLoadUrlEnd"}
var _wkeOnLoadUrlHeadersReceived = lib.LazyFunc[wkeOnLoadUrlHeadersReceived]{Name: "wkeOnLoadUrlHeadersReceived"}
var _wkeOnLoadUrlFinish = lib.LazyFunc[wkeOnLoadUrlFinish]{Name: "wkeOnLoadUrlFinish"}
var _wkeOnLoadUrlFail = lib.LazyFunc[wkeOnLoadUrlFail]{Name: "wkeOnLoadUrlFail"}
var _wkeOnDidCreateScriptContext = lib.LazyFunc[wkeOnDidCreateScriptContext]{Name: "wkeOnDidCreateScriptContext"}
var _wkeOnWillReleaseScriptContext = lib.LazyFunc[wkeOnWillReleaseScriptContext]{Name: "wkeOnWillReleaseScriptContext"}
var _wkeOnWindowClosing = lib.LazyFunc[wkeOnWindowClosing]{Name: "wkeOnWindowClosing"}
var _wkeOnWindowDestroy = lib.LazyFunc[wkeOnWindowDestroy]{Name: "wkeOnWindowDestroy"}
var _wkeOnDraggableRegionsChanged = lib.LazyFunc[wkeOnDraggableRegionsChanged]{Name: "wkeOnDraggableRegionsChanged"}
var _wkeOnWillMediaLoad = lib.LazyFunc[wkeOnWillMediaLoad]{Name: "wkeOnWillMediaLoad"}
var _wkeOnStartDragging = lib.LazyFunc[wkeOnStartDragging]{Name: "wkeOnStartDragging"}
var _wkeOnPrint = lib.LazyFunc[wkeOnPrint]{Name: "wkeOnPrint"}
var _wkeScreenshot = lib.LazyFunc[wkeScreenshot]{Name: "wkeScreenshot"}
var _wkeOnOtherLoad = lib.LazyFunc[wkeOnOtherLoad]{Name: "wkeOnOtherLoad"}
var _wkeOnContextMenuItemClick = lib.LazyFunc[wkeOnContextMenuItemClick]{Name: "wkeOnContextMenuItemClick"}
var _wkeIsProcessingUserGesture = lib.LazyFunc[wkeIsProcessingUserGesture]{Name: "wkeIsProcessingUserGesture"}
var _wkeNetSetMIMEType = lib.LazyFunc[wkeNetSetMIMEType]{Name: "wkeNetSetMIMEType"}
var _wkeNetGetMIMEType = lib.LazyFunc[wkeNetGetMIMEType]{Name: "wkeNetGetMIMEType"}
var _wkeNetGetReferrer = lib.LazyFunc[wkeNetGetReferrer]{Name: "wkeNetGetReferrer"}
var _wkeNetSetHTTPHeaderField = lib.LazyFunc[wkeNetSetHTTPHeaderField]{Name: "wkeNetSetHTTPHeaderField"}
var _wkeNetGetHTTPHeaderField = lib.LazyFunc[wkeNetGetHTTPHeaderField]{Name: "wkeNetGetHTTPHeaderField"}
var _wkeNetGetHTTPHeaderFieldFromResponse = lib.LazyFunc[wkeNetGetHTTPHeaderFieldFromResponse]{Name: "wkeNetGetHTTPHeaderFieldFromResponse"}
var _wkeNetSetData = lib.LazyFunc[wkeNetSetData]{Name: "wkeNetSetData"}
var _wkeNetHookRequest = lib.LazyFunc[wkeNetHookRequest]{Name: "wkeNetHookRequest"}
var _wkeNetGetRequestMethod = lib.LazyFunc[wkeNetGetRequestMethod]{Name: "wkeNetGetRequestMethod"}
var _wkeNetContinueJob = lib.LazyFunc[wkeNetContinueJob]{Name: "wkeNetContinueJob"}
var _wkeNetGetUrlByJob = lib.LazyFunc[wkeNetGetUrlByJob]{Name: "wkeNetGetUrlByJob"}
var _wkeNetGetRawHttpHead = lib.LazyFunc[wkeNetGetRawHttpHead]{Name: "wkeNetGetRawHttpHead"}
var _wkeNetGetRawResponseHead = lib.LazyFunc[wkeNetGetRawResponseHead]{Name: "wkeNetGetRawResponseHead"}
var _wkeNetCancelRequest = lib.LazyFunc[wkeNetCancelRequest]{Name: "wkeNetCancelRequest"}
var _wkeNetHoldJobToAsynCommit = lib.LazyFunc[wkeNetHoldJobToAsynCommit]{Name: "wkeNetHoldJobToAsynCommit"}
var _wkeNetOnResponse = lib.LazyFunc[wkeNetOnResponse]{Name: "wkeNetOnResponse"}
var _wkeNetGetFavicon = lib.LazyFunc[wkeNetGetFavicon]{Name: "wkeNetGetFavicon"}
var _wkeNetCreateWebUrlRequest = lib.LazyFunc[wkeNetCreateWebUrlRequest]{Name: "wkeNetCreateWebUrlRequest"}
var _wkeNetCreateWebUrlRequest2 = lib.LazyFunc[wkeNetCreateWebUrlRequest2]{Name: "wkeNetCreateWebUrlRequest2"}
var _wkeNetCopyWebUrlRequest = lib.LazyFunc[wkeNetCopyWebUrlRequest]{Name: "wkeNetCopyWebUrlRequest"}
var _wkeNetDeleteBlinkWebURLRequestPtr = lib.LazyFunc[wkeNetDeleteBlinkWebURLRequestPtr]{Name: "wkeNetDeleteBlinkWebURLRequestPtr"}
var _wkeNetAddHTTPHeaderFieldToUrlRequest = lib.LazyFunc[wkeNetAddHTTPHeaderFieldToUrlRequest]{Name: "wkeNetAddHTTPHeaderFieldToUrlRequest"}
var _wkeNetStartUrlRequest = lib.LazyFunc[wkeNetStartUrlRequest]{Name: "wkeNetStartUrlRequest"}
var _wkeNetGetHttpStatusCode = lib.LazyFunc[wkeNetGetHttpStatusCode]{Name: "wkeNetGetHttpStatusCode"}
var _wkeNetGetExpectedContentLength = lib.LazyFunc[wkeNetGetExpectedContentLength]{Name: "wkeNetGetExpectedContentLength"}
var _wkeNetGetResponseUrl = lib.LazyFunc[wkeNetGetResponseUrl]{Name: "wkeNetGetResponseUrl"}
var _wkeNetCancelWebUrlRequest = lib.LazyFunc[wkeNetCancelWebUrlRequest]{Name: "wkeNetCancelWebUrlRequest"}
var _wkeNetGetPostBody = lib.LazyFunc[wkeNetGetPostBody]{Name: "wkeNetGetPostBody"}
var _wkeNetFreePostBodyElements = lib.LazyFunc[wkeNetFreePostBodyElements]{Name: "wkeNetFreePostBodyElements"}
var _wkeNetCreatePostBodyElements = lib.LazyFunc[wkeNetCreatePostBodyElements]{Name: "wkeNetCreatePostBodyElements"}
var _wkeNetCreatePostBodyElement = lib.LazyFunc[wkeNetCreatePostBodyElement]{Name: "wkeNetCreatePostBodyElement"}
var _wkeNetFreePostBodyElement = lib.LazyFunc[wkeNetFreePostBodyElement]{Name: "wkeNetFreePostBodyElement"}
var _wkePopupDialogAndDownload = lib.LazyFunc[wkePopupDialogAndDownload]{Name: "wkePopupDialogAndDownload"}
var _wkeDownloadByPath = lib.LazyFunc[wkeDownloadByPath]{Name: "wkeDownloadByPath"}
var _wkeIsMainFrame = lib.LazyFunc[wkeIsMainFrame]{Name: "wkeIsMainFrame"}
var _wkeIsWebRemoteFrame = lib.LazyFunc[wkeIsWebRemoteFrame]{Name: "wkeIsWebRemoteFrame"}
var _wkeWebFrameGetMainFrame = lib.LazyFunc[wkeWebFrameGetMainFrame]{Name: "wkeWebFrameGetMainFrame"}
var _wkeRunJsByFrame = lib.LazyFunc[wkeRunJsByFrame]{Name: "wkeRunJsByFrame"}
var _wkeInsertCSSByFrame = lib.LazyFunc[wkeInsertCSSByFrame]{Name: "wkeInsertCSSByFrame"}
var _wkeWebFrameGetMainWorldScriptContext = lib.LazyFunc[wkeWebFrameGetMainWorldScriptContext]{Name: "wkeWebFrameGetMainWorldScriptContext"}
var _wkeGetBlinkMainThreadIsolate = lib.LazyFunc[wkeGetBlinkMainThreadIsolate]{Name: "wkeGetBlinkMainThreadIsolate"}
var _wkeCreateWebWindow = lib.LazyFunc[wkeCreateWebWindow]{Name: "wkeCreateWebWindow"}
var _wkeCreateWebCustomWindow = lib.LazyFunc[wkeCreateWebCustomWindow]{Name: "wkeCreateWebCustomWindow"}
var _wkeDestroyWebWindow = lib.LazyFunc[wkeDestroyWebWindow]{Name: "wkeDestroyWebWindow"}
var _wkeGetWindowHandle = lib.LazyFunc[wkeGetWindowHandle]{Name: "wkeGetWindowHandle"}
var _wkeShowWindow = lib.LazyFunc[wkeShowWindow]{Name: "wkeShowWindow"}
var _wkeEnableWindow = lib.LazyFunc[wkeEnableWindow]{Name: "wkeEnableWindow"}
var _wkeMoveWindow = lib.LazyFunc[wkeMoveWindow]{Name: "wkeMoveWindow"}
var _wkeMoveToCenter = lib.LazyFunc[wkeMoveToCenter]{Name: "wkeMoveToCenter"}
var _wkeResizeWindow = lib.LazyFunc[wkeResizeWindow]{Name: "wkeResizeWindow"}
var _wkeDragTargetDragEnter = lib.LazyFunc[wkeDragTargetDragEnter]{Name: "wkeDragTargetDragEnter"}
var _wkeDragTargetDragOver = lib.LazyFunc[wkeDragTargetDragOver]{Name: "wkeDragTargetDragOver"}
var _wkeDragTargetDragLeave = lib.LazyFunc[wkeDragTargetDragLeave]{Name: "wkeDragTargetDragLeave"}
var _wkeDragTargetDrop = lib.LazyFunc[wkeDragTargetDrop]{Name: "wkeDragTargetDrop"}
var _wkeDragTargetEnd = lib.LazyFunc[wkeDragTargetEnd]{Name: "wkeDragTargetEnd"}
var _wkeUtilSetUiCallback = lib.LazyFunc[wkeUtilSetUiCallback]{Name: "wkeUtilSetUiCallback"}
var _wkeUtilSerializeToMHTML = lib.LazyFunc[wkeUtilSerializeToMHTML]{Name: "wkeUtilSerializeToMHTML"}
var _wkeUtilPrintToPdf = lib.LazyFunc[wkeUtilPrintToPdf]{Name: "wkeUtilPrintToPdf"}
var _wkePrintToBitmap = lib.LazyFunc[wkePrintToBitmap]{Name: "wkePrintToBitmap"}
var _wkeUtilRelasePrintPdfDatas = lib.LazyFunc[wkeUtilRelasePrintPdfDatas]{Name: "wkeUtilRelasePrintPdfDatas"}
var _wkeSetWindowTitle = lib.LazyFunc[wkeSetWindowTitle]{Name: "wkeSetWindowTitle"}
var _wkeSetWindowTitleW = lib.LazyFunc[wkeSetWindowTitleW]{Name: "wkeSetWindowTitleW"}
var _wkeNodeOnCreateProcess = lib.LazyFunc[wkeNodeOnCreateProcess]{Name: "wkeNodeOnCreateProcess"}
var _wkeOnPluginFind = lib.LazyFunc[wkeOnPluginFind]{Name: "wkeOnPluginFind"}
var _wkeAddNpapiPlugin = lib.LazyFunc[wkeAddNpapiPlugin]{Name: "wkeAddNpapiPlugin"}
var _wkePluginListBuilderAddPlugin = lib.LazyFunc[wkePluginListBuilderAddPlugin]{Name: "wkePluginListBuilderAddPlugin"}
var _wkePluginListBuilderAddMediaTypeToLastPlugin = lib.LazyFunc[wkePluginListBuilderAddMediaTypeToLastPlugin]{Name: "wkePluginListBuilderAddMediaTypeToLastPlugin"}
var _wkePluginListBuilderAddFileExtensionToLastMediaType = lib.LazyFunc[wkePluginListBuilderAddFileExtensionToLastMediaType]{Name: "wkePluginListBuilderAddFileExtensionToLastMediaType"}
var _wkeGetWebViewByNData = lib.LazyFunc[wkeGetWebViewByNData]{Name: "wkeGetWebViewByNData"}
var _wkeRegisterEmbedderCustomElement = lib.LazyFunc[wkeRegisterEmbedderCustomElement]{Name: "wkeRegisterEmbedderCustomElement"}
var _wkeSetMediaPlayerFactory = lib.LazyFunc[wkeSetMediaPlayerFactory]{Name: "wkeSetMediaPlayerFactory"}
var _wkeGetContentAsMarkup = lib.LazyFunc[wkeGetContentAsMarkup]{Name: "wkeGetContentAsMarkup"}
var _wkeUtilDecodeURLEscape = lib.LazyFunc[wkeUtilDecodeURLEscape]{Name: "wkeUtilDecodeURLEscape"}
var _wkeUtilEncodeURLEscape = lib.LazyFunc[wkeUtilEncodeURLEscape]{Name: "wkeUtilEncodeURLEscape"}
var _wkeUtilBase64Encode = lib.LazyFunc[wkeUtilBase64Encode]{Name: "wkeUtilBase64Encode"}
var _wkeUtilBase64Decode = lib.LazyFunc[wkeUtilBase64Decode]{Name: "wkeUtilBase64Decode"}
var _wkeUtilCreateV8Snapshot = lib.LazyFunc[wkeUtilCreateV8Snapshot]{Name: "wkeUtilCreateV8Snapshot"}
var _wkeRunMessageLoop = lib.LazyFunc[wkeRunMessageLoop]{Name: "wkeRunMessageLoop"}
var _wkeSaveMemoryCache = lib.LazyFunc[wkeSaveMemoryCache]{Name: "wkeSaveMemoryCache"}
var _wkeJsBindFunction = lib.LazyFunc[wkeJsBindFunction]{Name: "wkeJsBindFunction"}
var _wkeJsBindGetter = lib.LazyFunc[wkeJsBindGetter]{Name: "wkeJsBindGetter"}
var _wkeJsBindSetter = lib.LazyFunc[wkeJsBindSetter]{Name: "wkeJsBindSetter"}
var _jsArgCount = lib.LazyFunc[jsArgCount]{Name: "jsArgCount"}
var _jsArgType = lib.LazyFunc[jsArgType]{Name: "jsArgType"}
var _jsArg = lib.LazyFunc[jsArg]{Name: "jsArg"}
var _jsTypeOf = lib.LazyFunc[jsTypeOf]{Name: "jsTypeOf"}
var _jsIsNumber = lib.LazyFunc[jsIsNumber]{Name: "jsIsNumber"}
var _jsIsString = lib.LazyFunc[jsIsString]{Name: "jsIsString"}
var _jsIsBoolean = lib.LazyFunc[jsIsBoolean]{Name: "jsIsBoolean"}
var _jsIsObject = lib.LazyFunc[jsIsObject]{Name: "jsIsObject"}
var _jsIsFunction = lib.LazyFunc[jsIsFunction]{Name: "jsIsFunction"}
var _jsIsUndefined = lib.LazyFunc[jsIsUndefined]{Name: "jsIsUndefined"}
var _jsIsNull = lib.LazyFunc[jsIsNull]{Name: "jsIsNull"}
var _jsIsArray = lib.LazyFunc[jsIsArray]{Name: "jsIsArray"}
var _jsIsTrue = lib.LazyFunc[jsIsTrue]{Name: "jsIsTrue"}
var _jsIsFalse = lib.LazyFunc[jsIsFalse]{Name: "jsIsFalse"}
var _jsToInt = lib.LazyFunc[jsToInt]{Name: "jsToInt"}
var _jsToFloat = lib.LazyFunc[jsToFloat]{Name: "jsToFloat"}
var _jsToDouble = lib.LazyFunc[jsToDouble]{Name: "jsToDouble"}
var _jsToDoubleString = lib.LazyFunc[jsToDoubleString]{Name: "jsToDoubleString"}
var _jsToBoolean = lib.LazyFunc[jsToBoolean]{Name: "jsToBoolean"}
var _jsArrayBuffer = lib.LazyFunc[jsArrayBuffer]{Name: "jsArrayBuffer"}
var _jsGetArrayBuffer = lib.LazyFunc[jsGetArrayBuffer]{Name: "jsGetArrayBuffer"}
var _jsToTempString = lib.LazyFunc[jsToTempString]{Name: "jsToTempString"}
var _jsToTempStringW = lib.LazyFunc[jsToTempStringW]{Name: "jsToTempStringW"}
var _jsToV8Value = lib.LazyFunc[jsToV8Value]{Name: "jsToV8Value"}
var _jsInt = lib.LazyFunc[jsInt]{Name: "jsInt"}
var _jsFloat = lib.LazyFunc[jsFloat]{Name: "jsFloat"}
var _jsDouble = lib.LazyFunc[jsDouble]{Name: "jsDouble"}
var _jsDoubleString = lib.LazyFunc[jsDoubleString]{Name: "jsDoubleString"}
var _jsBoolean = lib.LazyFunc[jsBoolean]{Name: "jsBoolean"}
var _jsUndefined = lib.LazyFunc[jsUndefined]{Name: "jsUndefined"}
var _jsNull = lib.LazyFunc[jsNull]{Name: "jsNull"}
var _jsTrue = lib.LazyFunc[jsTrue]{Name: "jsTrue"}
var _jsFalse = lib.LazyFunc[jsFalse]{Name: "jsFalse"}
var _jsString = lib.LazyFunc[jsString]{Name: "jsString"}
var _jsStringW = lib.LazyFunc[jsStringW]{Name: "jsStringW"}
var _jsEmptyObject = lib.LazyFunc[jsEmptyObject]{Name: "jsEmptyObject"}
var _jsEmptyArray = lib.LazyFunc[jsEmptyArray]{Name: "jsEmptyArray"}
var _jsObject = lib.LazyFunc[jsObject]{Name: "jsObject"}
var _jsFunction = lib.LazyFunc[jsFunction]{Name: "jsFunction"}
var _jsGetData = lib.LazyFunc[jsGetData]{Name: "jsGetData"}
var _jsGet = lib.LazyFunc[jsGet]{Name: "jsGet"}
var _jsSet = lib.LazyFunc[jsSet]{Name: "jsSet"}
var _jsGetAt = lib.LazyFunc[jsGetAt]{Name: "jsGetAt"}
var _jsSetAt = lib.LazyFunc[jsSetAt]{Name: "jsSetAt"}
var _jsGetKeys = lib.LazyFunc[jsGetKeys]{Name: "jsGetKeys"}
var _jsIsJsValueValid = lib.LazyFunc[jsIsJsValueValid]{Name: "jsIsJsValueValid"}
var _jsIsValidExecState = lib.LazyFunc[jsIsValidExecState]{Name: "jsIsValidExecState"}
var _jsDeleteObjectProp = lib.LazyFunc[jsDeleteObjectProp]{Name: "jsDeleteObjectProp"}
var _jsGetLength = lib.LazyFunc[jsGetLength]{Name: "jsGetLength"}
var _jsSetLength = lib.LazyFunc[jsSetLength]{Name: "jsSetLength"}
var _jsGlobalObject = lib.LazyFunc[jsGlobalObject]{Name: "jsGlobalObject"}
var _jsGetWebView = lib.LazyFunc[jsGetWebView]{Name: "jsGetWebView"}
var _jsEval = lib.LazyFunc[jsEval]{Name: "jsEval"}
var _jsEvalW = lib.LazyFunc[jsEvalW]{Name: "jsEvalW"}
var _jsEvalExW = lib.LazyFunc[jsEvalExW]{Name: "jsEvalExW"}
var _jsCall = lib.LazyFunc[jsCall]{Name: "jsCall"}
var _jsCallGlobal = lib.LazyFunc[jsCallGlobal]{Name: "jsCallGlobal"}
var _jsGetGlobal = lib.LazyFunc[jsGetGlobal]{Name: "jsGetGlobal"}
var _jsSetGlobal = lib.LazyFunc[jsSetGlobal]{Name: "jsSetGlobal"}
var _jsGC = lib.LazyFunc[jsGC]{Name: "jsGC"}
var _jsAddRef = lib.LazyFunc[jsAddRef]{Name: "jsAddRef"}
var _jsReleaseRef = lib.LazyFunc[jsReleaseRef]{Name: "jsReleaseRef"}
var _jsGetLastErrorIfException = lib.LazyFunc[jsGetLastErrorIfException]{Name: "jsGetLastErrorIfException"}
var _jsThrowException = lib.LazyFunc[jsThrowException]{Name: "jsThrowException"}
var _jsGetCallstack = lib.LazyFunc[jsGetCallstack]{Name: "jsGetCallstack"}
var _wkeInit = lib.LazyFunc[wkeInit]{Name: "wkeInit"}
var _wkeInitialize = lib.LazyFunc[wkeInitialize]{Name: "wkeInitialize"}
var _wkeInitializeEx = lib.LazyFunc[wkeInitializeEx]{Name: "wkeInitializeEx"}

func JsNativeFunction(es JsExecState, param unsafe.Pointer) JsValue {
	_wkeJsNativeFunction.LoadOnce()
	return _wkeJsNativeFunction.Call(es, param)
}

func Shutdown() {
	_wkeShutdown.LoadOnce()
	_wkeShutdown.Call()
}

func ShutdownForDebug() {
	_wkeShutdownForDebug.LoadOnce()
	_wkeShutdownForDebug.Call()
}

func Version() uint {
	_wkeVersion.LoadOnce()
	return _wkeVersion.Call()
}

func VersionString() string {
	_wkeVersionString.LoadOnce()
	return _wkeVersionString.Call()
}

func GC(webView WebView, intervalSec int64) {
	_wkeGC.LoadOnce()
	_wkeGC.Call(webView, intervalSec)
}

func SetResourceGc(webView WebView, intervalSec int64) {
	_wkeSetResourceGc.LoadOnce()
	_wkeSetResourceGc.Call(webView, intervalSec)
}

func SetFileSystem(pfnOpen FILE_OPEN, pfnClose FILE_CLOSE, pfnSize FILE_SIZE, pfnRead FILE_READ, pfnSeek FILE_SEEK) {
	_wkeSetFileSystem.LoadOnce()
	_wkeSetFileSystem.Call(pfnOpen, pfnClose, pfnSize, pfnRead, pfnSeek)
}

func WebViewName(webView WebView) string {
	_wkeWebViewName.LoadOnce()
	return _wkeWebViewName.Call(webView)
}

func SetWebViewName(webView WebView, name string) {
	_wkeSetWebViewName.LoadOnce()
	_wkeSetWebViewName.Call(webView, name)
}

func IsLoaded(webView WebView) bool {
	_wkeIsLoaded.LoadOnce()
	return _wkeIsLoaded.Call(webView)
}

func IsLoadFailed(webView WebView) bool {
	_wkeIsLoadFailed.LoadOnce()
	return _wkeIsLoadFailed.Call(webView)
}

func IsLoadComplete(webView WebView) bool {
	_wkeIsLoadComplete.LoadOnce()
	return _wkeIsLoadComplete.Call(webView)
}

func GetSource(webView WebView) string {
	_wkeGetSource.LoadOnce()
	return _wkeGetSource.Call(webView)
}

func Title(webView WebView) string {
	_wkeTitle.LoadOnce()
	return _wkeTitle.Call(webView)
}

func TitleW(webView WebView) []uint16 {
	_wkeTitleW.LoadOnce()
	return _wkeTitleW.Call(webView)
}

func Width(webView WebView) int {
	_wkeWidth.LoadOnce()
	return _wkeWidth.Call(webView)
}

func Height(webView WebView) int {
	_wkeHeight.LoadOnce()
	return _wkeHeight.Call(webView)
}

func ContentsWidth(webView WebView) int {
	_wkeContentsWidth.LoadOnce()
	return _wkeContentsWidth.Call(webView)
}

func ContentsHeight(webView WebView) int {
	_wkeContentsHeight.LoadOnce()
	return _wkeContentsHeight.Call(webView)
}

func SelectAll(webView WebView) {
	_wkeSelectAll.LoadOnce()
	_wkeSelectAll.Call(webView)
}

func Copy(webView WebView) {
	_wkeCopy.LoadOnce()
	_wkeCopy.Call(webView)
}

func Cut(webView WebView) {
	_wkeCut.LoadOnce()
	_wkeCut.Call(webView)
}

func Paste(webView WebView) {
	_wkePaste.LoadOnce()
	_wkePaste.Call(webView)
}

func Delete(webView WebView) {
	_wkeDelete.LoadOnce()
	_wkeDelete.Call(webView)
}

func CookieEnabled(webView WebView) bool {
	_wkeCookieEnabled.LoadOnce()
	return _wkeCookieEnabled.Call(webView)
}

func MediaVolume(webView WebView) float32 {
	_wkeMediaVolume.LoadOnce()
	return _wkeMediaVolume.Call(webView)
}

func MouseEvent(webView WebView, message uint, x, y int, flags uint) bool {
	_wkeMouseEvent.LoadOnce()
	return _wkeMouseEvent.Call(webView, message, x, y, flags)
}

func ContextMenuEvent(webView WebView, x, y int, flags uint) bool {
	_wkeContextMenuEvent.LoadOnce()
	return _wkeContextMenuEvent.Call(webView, x, y, flags)
}

func MouseWheel(webView WebView, x, y, delta int, flags uint) bool {
	_wkeMouseWheel.LoadOnce()
	return _wkeMouseWheel.Call(webView, x, y, delta, flags)
}

func KeyUp(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	_wkeKeyUp.LoadOnce()
	return _wkeKeyUp.Call(webView, virtualKeyCode, flags, systemKey)
}

func KeyDown(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	_wkeKeyDown.LoadOnce()
	return _wkeKeyDown.Call(webView, virtualKeyCode, flags, systemKey)
}

func KeyPress(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	_wkeKeyPress.LoadOnce()
	return _wkeKeyPress.Call(webView, virtualKeyCode, flags, systemKey)
}

func Focus(webView WebView) {
	_wkeFocus.LoadOnce()
	_wkeFocus.Call(webView)
}

func Unfocus(webView WebView) {
	_wkeUnfocus.LoadOnce()
	_wkeUnfocus.Call(webView)
}

func GetCaret(webView WebView) Rect {
	_wkeGetCaret.LoadOnce()
	return _wkeGetCaret.Call(webView)
}

func Awaken(webView WebView) {
	_wkeAwaken.LoadOnce()
	_wkeAwaken.Call(webView)
}

func ZoomFactor(webView WebView) float32 {
	_wkeZoomFactor.LoadOnce()
	return _wkeZoomFactor.Call(webView)
}

func SetClientHandler(webView WebView, handler *ClientHandler) {
	_wkeSetClientHandler.LoadOnce()
	_wkeSetClientHandler.Call(webView, handler)
}

func GetClientHandler(webView WebView) *ClientHandler {
	_wkeGetClientHandler.LoadOnce()
	return _wkeGetClientHandler.Call(webView)
}

func JsToString(es JsExecState, v JsValue) string {
	_jsToString.LoadOnce()
	return _jsToString.Call(es, v)
}

func JsToStringW(es JsExecState, v JsValue) []uint16 {
	_jsToStringW.LoadOnce()
	return _jsToStringW.Call(es, v)
}

func SetViewSettings(webView WebView, settings *ViewSettings) {
	_wkeSetViewSettings.LoadOnce()
	_wkeSetViewSettings.Call(webView, settings)
}

func SetDebugConfig(webView WebView, debugString, param string) {
	_wkeSetDebugConfig.LoadOnce()
	_wkeSetDebugConfig.Call(webView, debugString, param)
}

func ToString(wkeStr String) string {
	_wkeToString.LoadOnce()
	return _wkeToString.Call(wkeStr)
}

func ToStringW(wkeStr String) []uint16 {
	_wkeToStringW.LoadOnce()
	return _wkeToStringW.Call(wkeStr)
}

func GetDebugConfig(webView WebView, debugString string) unsafe.Pointer {
	_wkeGetDebugConfig.LoadOnce()
	return _wkeGetDebugConfig.Call(webView, debugString)
}

func IsInitialize() bool {
	_wkeIsInitialize.LoadOnce()
	return _wkeIsInitialize.Call()
}

func Finalize() {
	_wkeFinalize.LoadOnce()
	_wkeFinalize.Call()
}

func Update() {
	_wkeUpdate.LoadOnce()
	_wkeUpdate.Call()
}

func GetVersion() uint {
	_wkeGetVersion.LoadOnce()
	return _wkeGetVersion.Call()
}

func GetVersionString() string {
	_wkeGetVersionString.LoadOnce()
	return _wkeGetVersionString.Call()
}

func CreateWebView() WebView {
	_wkeCreateWebView.LoadOnce()
	return _wkeCreateWebView.Call()
}

func DestroyWebView(webView WebView) {
	_wkeDestroyWebView.LoadOnce()
	_wkeDestroyWebView.Call(webView)
}

func SetMemoryCacheEnable(webView WebView, b bool) {
	_wkeSetMemoryCacheEnable.LoadOnce()
	_wkeSetMemoryCacheEnable.Call(webView, b)
}

func SetMouseEnabled(webView WebView, b bool) {
	_wkeSetMouseEnabled.LoadOnce()
	_wkeSetMouseEnabled.Call(webView, b)
}

func SetTouchEnabled(webView WebView, b bool) {
	_wkeSetTouchEnabled.LoadOnce()
	_wkeSetTouchEnabled.Call(webView, b)
}

func SetSystemTouchEnabled(webView WebView, b bool) {
	_wkeSetSystemTouchEnabled.LoadOnce()
	_wkeSetSystemTouchEnabled.Call(webView, b)
}

func SetContextMenuEnabled(webView WebView, b bool) {
	_wkeSetContextMenuEnabled.LoadOnce()
	_wkeSetContextMenuEnabled.Call(webView, b)
}

func SetNavigationToNewWindowEnable(webView WebView, b bool) {
	_wkeSetNavigationToNewWindowEnable.LoadOnce()
	_wkeSetNavigationToNewWindowEnable.Call(webView, b)
}

func SetCspCheckEnable(webView WebView, b bool) {
	_wkeSetCspCheckEnable.LoadOnce()
	_wkeSetCspCheckEnable.Call(webView, b)
}

func SetNpapiPluginsEnabled(webView WebView, b bool) {
	_wkeSetNpapiPluginsEnabled.LoadOnce()
	_wkeSetNpapiPluginsEnabled.Call(webView, b)
}

func SetHeadlessEnabled(webView WebView, b bool) {
	_wkeSetHeadlessEnabled.LoadOnce()
	_wkeSetHeadlessEnabled.Call(webView, b)
}

func SetDragEnable(webView WebView, b bool) {
	_wkeSetDragEnable.LoadOnce()
	_wkeSetDragEnable.Call(webView, b)
}

func SetDragDropEnable(webView WebView, b bool) {
	_wkeSetDragDropEnable.LoadOnce()
	_wkeSetDragDropEnable.Call(webView, b)
}

func SetLanguage(webView WebView, language string) {
	_wkeSetLanguage.LoadOnce()
	_wkeSetLanguage.Call(webView, language)
}

func SetViewNetInterface(webView WebView, netInterface string) {
	_wkeSetViewNetInterface.LoadOnce()
	_wkeSetViewNetInterface.Call(webView, netInterface)
}

func SetContextMenuItemShow(webView WebView, item MenuItemId, isShow bool) {
	_wkeSetContextMenuItemShow.LoadOnce()
	_wkeSetContextMenuItemShow.Call(webView, item, isShow)
}

func SetProxy(proxy *Proxy) {
	_wkeSetProxy.LoadOnce()
	_wkeSetProxy.Call(proxy)
}

func GetName(webView WebView) string {
	_wkeGetName.LoadOnce()
	return _wkeGetName.Call(webView)
}

func IsTransparent(webView WebView) bool {
	_wkeIsTransparent.LoadOnce()
	return _wkeIsTransparent.Call(webView)
}

func GetUserAgent(webView WebView) string {
	_wkeGetUserAgent.LoadOnce()
	return _wkeGetUserAgent.Call(webView)
}

func SetViewProxy(webView WebView, proxy *Proxy) {
	_wkeSetViewProxy.LoadOnce()
	_wkeSetViewProxy.Call(webView, proxy)
}

func SetName(webView WebView, name string) {
	_wkeSetName.LoadOnce()
	_wkeSetName.Call(webView, name)
}

func SetHandle(webView WebView, wnd HWND) {
	_wkeSetHandle.LoadOnce()
	_wkeSetHandle.Call(webView, wnd)
}

func SetTransparent(webView WebView, transparent bool) {
	_wkeSetTransparent.LoadOnce()
	_wkeSetTransparent.Call(webView, transparent)
}

func SetUserAgent(webView WebView, userAgent string) {
	_wkeSetUserAgent.LoadOnce()
	_wkeSetUserAgent.Call(webView, userAgent)
}

func SetUserAgentW(webView WebView, userAgent []uint16) {
	_wkeSetUserAgentW.LoadOnce()
	_wkeSetUserAgentW.Call(webView, userAgent)
}

func SetHandleOffset(webView WebView, x, y int) {
	_wkeSetHandleOffset.LoadOnce()
	_wkeSetHandleOffset.Call(webView, x, y)
}

func ShowDevtools(webView WebView, path []uint16, callback OnShowDevtoolsCallback, param unsafe.Pointer) {
	_wkeShowDevtools.LoadOnce()
	_wkeShowDevtools.Call(webView, path, callback, param)
}

func LoadW(webView WebView, url []uint16) {
	_wkeLoadW.LoadOnce()
	_wkeLoadW.Call(webView, url)
}

func LoadURL(webView WebView, url string) {
	_wkeLoadURL.LoadOnce()
	_wkeLoadURL.Call(webView, url)
}

func LoadURLW(webView WebView, url []uint16) {
	_wkeLoadURLW.LoadOnce()
	_wkeLoadURLW.Call(webView, url)
}

func PostURL(wkeView WebView, url string, postData []byte, postLen int) {
	_wkePostURL.LoadOnce()
	_wkePostURL.Call(wkeView, url, postData, postLen)
}

func PostURLW(wkeView WebView, url []uint16, postData []byte, postLen int) {
	_wkePostURLW.LoadOnce()
	_wkePostURLW.Call(wkeView, url, postData, postLen)
}

func LoadHTML(webView WebView, html string) {
	_wkeLoadHTML.LoadOnce()
	_wkeLoadHTML.Call(webView, html)
}

func LoadHtmlWithBaseUrl(webView WebView, html, baseUrl string) {
	_wkeLoadHtmlWithBaseUrl.LoadOnce()
	_wkeLoadHtmlWithBaseUrl.Call(webView, html, baseUrl)
}

func LoadHTMLW(webView WebView, html []uint16) {
	_wkeLoadHTMLW.LoadOnce()
	_wkeLoadHTMLW.Call(webView, html)
}

func LoadFile(webView WebView, filename string) {
	_wkeLoadFile.LoadOnce()
	_wkeLoadFile.Call(webView, filename)
}

func LoadFileW(webView WebView, filename []uint16) {
	_wkeLoadFileW.LoadOnce()
	_wkeLoadFileW.Call(webView, filename)
}

func GetURL(webView WebView) string {
	_wkeGetURL.LoadOnce()
	return _wkeGetURL.Call(webView)
}

func GetFrameUrl(webView WebView, frameId WebFrameHandle) string {
	_wkeGetFrameUrl.LoadOnce()
	return _wkeGetFrameUrl.Call(webView, frameId)
}

func IsLoading(webView WebView) bool {
	_wkeIsLoading.LoadOnce()
	return _wkeIsLoading.Call(webView)
}

func IsLoadingSucceeded(webView WebView) bool {
	_wkeIsLoadingSucceeded.LoadOnce()
	return _wkeIsLoadingSucceeded.Call(webView)
}

func IsLoadingFailed(webView WebView) bool {
	_wkeIsLoadingFailed.LoadOnce()
	return _wkeIsLoadingFailed.Call(webView)
}

func IsLoadingCompleted(webView WebView) bool {
	_wkeIsLoadingCompleted.LoadOnce()
	return _wkeIsLoadingCompleted.Call(webView)
}

func IsDocumentReady(webView WebView) bool {
	_wkeIsDocumentReady.LoadOnce()
	return _wkeIsDocumentReady.Call(webView)
}

func StopLoading(webView WebView) {
	_wkeStopLoading.LoadOnce()
	_wkeStopLoading.Call(webView)
}

func Reload(webView WebView) {
	_wkeReload.LoadOnce()
	_wkeReload.Call(webView)
}

func GoToOffset(webView WebView, offset int) {
	_wkeGoToOffset.LoadOnce()
	_wkeGoToOffset.Call(webView, offset)
}

func GoToIndex(webView WebView, index int) {
	_wkeGoToIndex.LoadOnce()
	_wkeGoToIndex.Call(webView, index)
}

func GetWebviewId(webView WebView) int {
	_wkeGetWebviewId.LoadOnce()
	return _wkeGetWebviewId.Call(webView)
}

func IsWebviewAlive(id int) bool {
	_wkeIsWebviewAlive.LoadOnce()
	return _wkeIsWebviewAlive.Call(id)
}

func IsWebviewValid(webView WebView) bool {
	_wkeIsWebviewValid.LoadOnce()
	return _wkeIsWebviewValid.Call(webView)
}

func GetDocumentCompleteURL(webView WebView, frameId WebFrameHandle, partialURL string) string {
	_wkeGetDocumentCompleteURL.LoadOnce()
	return _wkeGetDocumentCompleteURL.Call(webView, frameId, partialURL)
}

func CreateMemBuf(webView WebView, buf unsafe.Pointer, length uintptr) *MemBuf {
	_wkeCreateMemBuf.LoadOnce()
	return _wkeCreateMemBuf.Call(webView, buf, length)
}

func FreeMemBuf(buf *MemBuf) {
	_wkeFreeMemBuf.LoadOnce()
	_wkeFreeMemBuf.Call(buf)
}

func GetTitle(webView WebView) string {
	_wkeGetTitle.LoadOnce()
	return _wkeGetTitle.Call(webView)
}

func GetTitleW(webView WebView) []uint16 {
	_wkeGetTitleW.LoadOnce()
	return _wkeGetTitleW.Call(webView)
}

func Resize(webView WebView, w, h int) {
	_wkeResize.LoadOnce()
	_wkeResize.Call(webView, w, h)
}

func GetWidth(webView WebView) int {
	_wkeGetWidth.LoadOnce()
	return _wkeGetWidth.Call(webView)
}

func GetHeight(webView WebView) int {
	_wkeGetHeight.LoadOnce()
	return _wkeGetHeight.Call(webView)
}

func GetContentWidth(webView WebView) int {
	_wkeGetContentWidth.LoadOnce()
	return _wkeGetContentWidth.Call(webView)
}

func GetContentHeight(webView WebView) int {
	_wkeGetContentHeight.LoadOnce()
	return _wkeGetContentHeight.Call(webView)
}

func SetDirty(webView WebView, dirty bool) {
	_wkeSetDirty.LoadOnce()
	_wkeSetDirty.Call(webView, dirty)
}

func IsDirty(webView WebView) bool {
	_wkeIsDirty.LoadOnce()
	return _wkeIsDirty.Call(webView)
}

func AddDirtyArea(webView WebView, x, y, w, h int) {
	_wkeAddDirtyArea.LoadOnce()
	_wkeAddDirtyArea.Call(webView, x, y, w, h)
}

func LayoutIfNeeded(webView WebView) {
	_wkeLayoutIfNeeded.LoadOnce()
	_wkeLayoutIfNeeded.Call(webView)
}

func Paint2(webView WebView, bits unsafe.Pointer, bufWid, bufHei, xDst, yDst, w, h, xSrc, ySrc int, bCopyAlpha bool) {
	_wkePaint2.LoadOnce()
	_wkePaint2.Call(webView, bits, bufWid, bufHei, xDst, yDst, w, h, xSrc, ySrc, bCopyAlpha)
}

func Paint(webView WebView, bits unsafe.Pointer, pitch int) {
	_wkePaint.LoadOnce()
	_wkePaint.Call(webView, bits, pitch)
}

func RepaintIfNeeded(webView WebView) {
	_wkeRepaintIfNeeded.LoadOnce()
	_wkeRepaintIfNeeded.Call(webView)
}

func GetViewDC(webView WebView) HDC {
	_wkeGetViewDC.LoadOnce()
	return _wkeGetViewDC.Call(webView)
}

func UnlockViewDC(webView WebView) {
	_wkeUnlockViewDC.LoadOnce()
	_wkeUnlockViewDC.Call(webView)
}

func GetHostHWND(webView WebView) HWND {
	_wkeGetHostHWND.LoadOnce()
	return _wkeGetHostHWND.Call(webView)
}

func CanGoBack(webView WebView) bool {
	_wkeCanGoBack.LoadOnce()
	return _wkeCanGoBack.Call(webView)
}

func GoBack(webView WebView) bool {
	_wkeGoBack.LoadOnce()
	return _wkeGoBack.Call(webView)
}

func CanGoForward(webView WebView) bool {
	_wkeCanGoForward.LoadOnce()
	return _wkeCanGoForward.Call(webView)
}

func GoForward(webView WebView) bool {
	_wkeGoForward.LoadOnce()
	return _wkeGoForward.Call(webView)
}

func NavigateAtIndex(webView WebView, index int) bool {
	_wkeNavigateAtIndex.LoadOnce()
	return _wkeNavigateAtIndex.Call(webView, index)
}

func GetNavigateIndex(webView WebView) int {
	_wkeGetNavigateIndex.LoadOnce()
	return _wkeGetNavigateIndex.Call(webView)
}

func EditorSelectAll(webView WebView) {
	_wkeEditorSelectAll.LoadOnce()
	_wkeEditorSelectAll.Call(webView)
}

func EditorUnSelect(webView WebView) {
	_wkeEditorUnSelect.LoadOnce()
	_wkeEditorUnSelect.Call(webView)
}

func EditorCopy(webView WebView) {
	_wkeEditorCopy.LoadOnce()
	_wkeEditorCopy.Call(webView)
}

func EditorCut(webView WebView) {
	_wkeEditorCut.LoadOnce()
	_wkeEditorCut.Call(webView)
}

func EditorPaste(webView WebView) {
	_wkeEditorPaste.LoadOnce()
	_wkeEditorPaste.Call(webView)
}

func EditorDelete(webView WebView) {
	_wkeEditorDelete.LoadOnce()
	_wkeEditorDelete.Call(webView)
}

func EditorUndo(webView WebView) {
	_wkeEditorUndo.LoadOnce()
	_wkeEditorUndo.Call(webView)
}

func EditorRedo(webView WebView) {
	_wkeEditorRedo.LoadOnce()
	_wkeEditorRedo.Call(webView)
}

func GetCookieW(webView WebView) []uint16 {
	_wkeGetCookieW.LoadOnce()
	return _wkeGetCookieW.Call(webView)
}

func GetCookie(webView WebView) string {
	_wkeGetCookie.LoadOnce()
	return _wkeGetCookie.Call(webView)
}

func SetCookie(webView WebView, url, cookie string) {
	_wkeSetCookie.LoadOnce()
	_wkeSetCookie.Call(webView, url, cookie)
}

func VisitAllCookie(webView WebView, params unsafe.Pointer, visitor CookieVisitor) {
	_wkeVisitAllCookie.LoadOnce()
	_wkeVisitAllCookie.Call(webView, params, visitor)
}

func PerformCookieCommand(webView WebView, command CookieCommand) {
	_wkePerformCookieCommand.LoadOnce()
	_wkePerformCookieCommand.Call(webView, command)
}

func SetCookieEnabled(webView WebView, enable bool) {
	_wkeSetCookieEnabled.LoadOnce()
	_wkeSetCookieEnabled.Call(webView, enable)
}

func IsCookieEnabled(webView WebView) bool {
	_wkeIsCookieEnabled.LoadOnce()
	return _wkeIsCookieEnabled.Call(webView)
}

func SetCookieJarPath(webView WebView, path []uint16) {
	_wkeSetCookieJarPath.LoadOnce()
	_wkeSetCookieJarPath.Call(webView, path)
}

func SetCookieJarFullPath(webView WebView, path []uint16) {
	_wkeSetCookieJarFullPath.LoadOnce()
	_wkeSetCookieJarFullPath.Call(webView, path)
}

func ClearCookie(webView WebView) {
	_wkeClearCookie.LoadOnce()
	_wkeClearCookie.Call(webView)
}

func SetLocalStorageFullPath(webView WebView, path []uint16) {
	_wkeSetLocalStorageFullPath.LoadOnce()
	_wkeSetLocalStorageFullPath.Call(webView, path)
}

func AddPluginDirectory(webView WebView, path []uint16) {
	_wkeAddPluginDirectory.LoadOnce()
	_wkeAddPluginDirectory.Call(webView, path)
}

func SetMediaVolume(webView WebView, volume float32) {
	_wkeSetMediaVolume.LoadOnce()
	_wkeSetMediaVolume.Call(webView, volume)
}

func GetMediaVolume(webView WebView) float32 {
	_wkeGetMediaVolume.LoadOnce()
	return _wkeGetMediaVolume.Call(webView)
}

func FireMouseEvent(webView WebView, message uint, x, y int, flags uint) bool {
	_wkeFireMouseEvent.LoadOnce()
	return _wkeFireMouseEvent.Call(webView, message, x, y, flags)
}

func FireContextMenuEvent(webView WebView, x, y int, flags uint) bool {
	_wkeFireContextMenuEvent.LoadOnce()
	return _wkeFireContextMenuEvent.Call(webView, x, y, flags)
}

func FireMouseWheelEvent(webView WebView, x, y, delta int, flags uint) bool {
	_wkeFireMouseWheelEvent.LoadOnce()
	return _wkeFireMouseWheelEvent.Call(webView, x, y, delta, flags)
}

func FireKeyUpEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	_wkeFireKeyUpEvent.LoadOnce()
	return _wkeFireKeyUpEvent.Call(webView, virtualKeyCode, flags, systemKey)
}

func FireKeyDownEvent(webView WebView, virtualKeyCode, flags uint, systemKey bool) bool {
	_wkeFireKeyDownEvent.LoadOnce()
	return _wkeFireKeyDownEvent.Call(webView, virtualKeyCode, flags, systemKey)
}

func FireKeyPressEvent(webView WebView, charCode, flags uint, systemKey bool) bool {
	_wkeFireKeyPressEvent.LoadOnce()
	return _wkeFireKeyPressEvent.Call(webView, charCode, flags, systemKey)
}

func FireWindowsMessage(webView WebView, hWnd HWND, message uint, wParam WPARAM, lParam LPARAM, result *LRESULT) bool {
	_wkeFireWindowsMessage.LoadOnce()
	return _wkeFireWindowsMessage.Call(webView, hWnd, message, wParam, lParam, result)
}

func SetFocus(webView WebView) {
	_wkeSetFocus.LoadOnce()
	_wkeSetFocus.Call(webView)
}

func KillFocus(webView WebView) {
	_wkeKillFocus.LoadOnce()
	_wkeKillFocus.Call(webView)
}

func GetCaretRect(webView WebView) Rect {
	_wkeGetCaretRect.LoadOnce()
	return _wkeGetCaretRect.Call(webView)
}

func GetCaretRect2(webView WebView) *Rect {
	_wkeGetCaretRect2.LoadOnce()
	return _wkeGetCaretRect2.Call(webView)
}

func RunJS(webView WebView, script string) JsValue {
	_wkeRunJS.LoadOnce()
	return _wkeRunJS.Call(webView, script)
}

func RunJSW(webView WebView, script []uint16) JsValue {
	_wkeRunJSW.LoadOnce()
	return _wkeRunJSW.Call(webView, script)
}

func GlobalExec(webView WebView) JsExecState {
	_wkeGlobalExec.LoadOnce()
	return _wkeGlobalExec.Call(webView)
}

func GetGlobalExecByFrame(webView WebView, frameId WebFrameHandle) JsExecState {
	_wkeGetGlobalExecByFrame.LoadOnce()
	return _wkeGetGlobalExecByFrame.Call(webView, frameId)
}

func Sleep(webView WebView) {
	_wkeSleep.LoadOnce()
	_wkeSleep.Call(webView)
}

func Wake(webView WebView) {
	_wkeWake.LoadOnce()
	_wkeWake.Call(webView)
}

func IsAwake(webView WebView) bool {
	_wkeIsAwake.LoadOnce()
	return _wkeIsAwake.Call(webView)
}

func SetZoomFactor(webView WebView, factor float32) {
	_wkeSetZoomFactor.LoadOnce()
	_wkeSetZoomFactor.Call(webView, factor)
}

func GetZoomFactor(webView WebView) float32 {
	_wkeGetZoomFactor.LoadOnce()
	return _wkeGetZoomFactor.Call(webView)
}

func EnableHighDPISupport() {
	_wkeEnableHighDPISupport.LoadOnce()
	_wkeEnableHighDPISupport.Call()
}

func SetEditable(webView WebView, editable bool) {
	_wkeSetEditable.LoadOnce()
	_wkeSetEditable.Call(webView, editable)
}

func GetString(wkeStr String) string {
	_wkeGetString.LoadOnce()
	return _wkeGetString.Call(wkeStr)
}

func GetStringW(wkeStr String) []uint16 {
	_wkeGetStringW.LoadOnce()
	return _wkeGetStringW.Call(wkeStr)
}

func SetString(wkeStr String, str *byte, len uintptr) {
	_wkeSetString.LoadOnce()
	_wkeSetString.Call(wkeStr, str, len)
}

func SetStringWithoutNullTermination(wkeStr String, str *byte, len uintptr) {
	_wkeSetStringWithoutNullTermination.LoadOnce()
	_wkeSetStringWithoutNullTermination.Call(wkeStr, str, len)
}

func SetStringW(wkeStr String, str *uint16, len uintptr) {
	_wkeSetStringW.LoadOnce()
	_wkeSetStringW.Call(wkeStr, str, len)
}

func CreateString(str *byte, len uintptr) String {
	_wkeCreateString.LoadOnce()
	return _wkeCreateString.Call(str, len)
}

func CreateStringW(str *uint16, len uintptr) String {
	_wkeCreateStringW.LoadOnce()
	return _wkeCreateStringW.Call(str, len)
}

func CreateStringWithoutNullTermination(str *byte, len uintptr) String {
	_wkeCreateStringWithoutNullTermination.LoadOnce()
	return _wkeCreateStringWithoutNullTermination.Call(str, len)
}

func GetStringLen(wkeStr String) uintptr {
	_wkeGetStringLen.LoadOnce()
	return _wkeGetStringLen.Call(wkeStr)
}

func DeleteString(wkeStr String) {
	_wkeDeleteString.LoadOnce()
	_wkeDeleteString.Call(wkeStr)
}

func GetWebViewForCurrentContext() WebView {
	_wkeGetWebViewForCurrentContext.LoadOnce()
	return _wkeGetWebViewForCurrentContext.Call()
}

func SetUserKeyValue(webView WebView, key string, value unsafe.Pointer) {
	_wkeSetUserKeyValue.LoadOnce()
	_wkeSetUserKeyValue.Call(webView, key, value)
}

func GetUserKeyValue(webView WebView, key string) unsafe.Pointer {
	_wkeGetUserKeyValue.LoadOnce()
	return _wkeGetUserKeyValue.Call(webView, key)
}

func GetCursorInfoType(webView WebView) int {
	_wkeGetCursorInfoType.LoadOnce()
	return _wkeGetCursorInfoType.Call(webView)
}

func SetCursorInfoType(webView WebView, type_ int) {
	_wkeSetCursorInfoType.LoadOnce()
	_wkeSetCursorInfoType.Call(webView, type_)
}

func SetDragFiles(webView WebView, clintPos, screenPos *Point, files *String, filesCount int) {
	_wkeSetDragFiles.LoadOnce()
	_wkeSetDragFiles.Call(webView, clintPos, screenPos, files, filesCount)
}

func SetDeviceParameter(webView WebView, device, paramStr string, paramInt int, paramFloat float32) {
	_wkeSetDeviceParameter.LoadOnce()
	_wkeSetDeviceParameter.Call(webView, device, paramStr, paramInt, paramFloat)
}

func GetTempCallbackInfo(webView WebView) *TempCallbackInfo {
	_wkeGetTempCallbackInfo.LoadOnce()
	return _wkeGetTempCallbackInfo.Call(webView)
}

func OnCaretChanged(webView WebView, callback CaretChangedCallback, callbackParam unsafe.Pointer) {
	_wkeOnCaretChanged.LoadOnce()
	_wkeOnCaretChanged.Call(webView, callback, callbackParam)
}

func OnMouseOverUrlChanged(webView WebView, callback TitleChangedCallback, callbackParam unsafe.Pointer) {
	_wkeOnMouseOverUrlChanged.LoadOnce()
	_wkeOnMouseOverUrlChanged.Call(webView, callback, callbackParam)
}

func OnTitleChanged(webView WebView, callback TitleChangedCallback, callbackParam unsafe.Pointer) {
	_wkeOnTitleChanged.LoadOnce()
	_wkeOnTitleChanged.Call(webView, callback, callbackParam)
}

func OnURLChanged(webView WebView, callback URLChangedCallback, callbackParam unsafe.Pointer) {
	_wkeOnURLChanged.LoadOnce()
	_wkeOnURLChanged.Call(webView, callback, callbackParam)
}

func OnURLChanged2(webView WebView, callback URLChangedCallback2, callbackParam unsafe.Pointer) {
	_wkeOnURLChanged2.LoadOnce()
	_wkeOnURLChanged2.Call(webView, callback, callbackParam)
}

func OnPaintUpdated(webView WebView, callback PaintUpdatedCallback, callbackParam unsafe.Pointer) {
	_wkeOnPaintUpdated.LoadOnce()
	_wkeOnPaintUpdated.Call(webView, callback, callbackParam)
}

func OnPaintBitUpdated(webView WebView, callback PaintBitUpdatedCallback, callbackParam unsafe.Pointer) {
	_wkeOnPaintBitUpdated.LoadOnce()
	_wkeOnPaintBitUpdated.Call(webView, callback, callbackParam)
}

func OnAlertBox(webView WebView, callback AlertBoxCallback, callbackParam unsafe.Pointer) {
	_wkeOnAlertBox.LoadOnce()
	_wkeOnAlertBox.Call(webView, callback, callbackParam)
}

func OnConfirmBox(webView WebView, callback ConfirmBoxCallback, callbackParam unsafe.Pointer) {
	_wkeOnConfirmBox.LoadOnce()
	_wkeOnConfirmBox.Call(webView, callback, callbackParam)
}

func OnPromptBox(webView WebView, callback PromptBoxCallback, callbackParam unsafe.Pointer) {
	_wkeOnPromptBox.LoadOnce()
	_wkeOnPromptBox.Call(webView, callback, callbackParam)
}

func OnNavigation(webView WebView, callback NavigationCallback, param unsafe.Pointer) {
	_wkeOnNavigation.LoadOnce()
	_wkeOnNavigation.Call(webView, callback, param)
}

func OnCreateView(webView WebView, callback CreateViewCallback, param unsafe.Pointer) {
	_wkeOnCreateView.LoadOnce()
	_wkeOnCreateView.Call(webView, callback, param)
}

func OnDocumentReady(webView WebView, callback DocumentReadyCallback, param unsafe.Pointer) {
	_wkeOnDocumentReady.LoadOnce()
	_wkeOnDocumentReady.Call(webView, callback, param)
}

func OnDocumentReady2(webView WebView, callback DocumentReady2Callback, param unsafe.Pointer) {
	_wkeOnDocumentReady2.LoadOnce()
	_wkeOnDocumentReady2.Call(webView, callback, param)
}

func OnLoadingFinish(webView WebView, callback LoadingFinishCallback, param unsafe.Pointer) {
	_wkeOnLoadingFinish.LoadOnce()
	_wkeOnLoadingFinish.Call(webView, callback, param)
}

func OnDownload(webView WebView, callback DownloadCallback, param unsafe.Pointer) {
	_wkeOnDownload.LoadOnce()
	_wkeOnDownload.Call(webView, callback, param)
}

func OnDownload2(webView WebView, callback Download2Callback, param unsafe.Pointer) {
	_wkeOnDownload2.LoadOnce()
	_wkeOnDownload2.Call(webView, callback, param)
}

func OnConsole(webView WebView, callback ConsoleCallback, param unsafe.Pointer) {
	_wkeOnConsole.LoadOnce()
	_wkeOnConsole.Call(webView, callback, param)
}

func SetUIThreadCallback(webView WebView, callback CallUiThread, param unsafe.Pointer) {
	_wkeSetUIThreadCallback.LoadOnce()
	_wkeSetUIThreadCallback.Call(webView, callback, param)
}

func OnLoadUrlBegin(webView WebView, callback LoadUrlBeginCallback, callbackParam unsafe.Pointer) {
	_wkeOnLoadUrlBegin.LoadOnce()
	_wkeOnLoadUrlBegin.Call(webView, callback, callbackParam)
}

func OnLoadUrlEnd(webView WebView, callback LoadUrlEndCallback, callbackParam unsafe.Pointer) {
	_wkeOnLoadUrlEnd.LoadOnce()
	_wkeOnLoadUrlEnd.Call(webView, callback, callbackParam)
}

func OnLoadUrlHeadersReceived(webView WebView, callback LoadUrlHeadersReceivedCallback, callbackParam unsafe.Pointer) {
	_wkeOnLoadUrlHeadersReceived.LoadOnce()
	_wkeOnLoadUrlHeadersReceived.Call(webView, callback, callbackParam)
}

func OnLoadUrlFinish(webView WebView, callback LoadUrlFinishCallback, callbackParam unsafe.Pointer) {
	_wkeOnLoadUrlFinish.LoadOnce()
	_wkeOnLoadUrlFinish.Call(webView, callback, callbackParam)
}

func OnLoadUrlFail(webView WebView, callback LoadUrlFailCallback, callbackParam unsafe.Pointer) {
	_wkeOnLoadUrlFail.LoadOnce()
	_wkeOnLoadUrlFail.Call(webView, callback, callbackParam)
}

func OnDidCreateScriptContext(webView WebView, callback DidCreateScriptContextCallback, callbackParam unsafe.Pointer) {
	_wkeOnDidCreateScriptContext.LoadOnce()
	_wkeOnDidCreateScriptContext.Call(webView, callback, callbackParam)
}

func OnWillReleaseScriptContext(webView WebView, callback WillReleaseScriptContextCallback, callbackParam unsafe.Pointer) {
	_wkeOnWillReleaseScriptContext.LoadOnce()
	_wkeOnWillReleaseScriptContext.Call(webView, callback, callbackParam)
}

func OnWindowClosing(webWindow WebView, callback WindowClosingCallback, param unsafe.Pointer) {
	_wkeOnWindowClosing.LoadOnce()
	_wkeOnWindowClosing.Call(webWindow, callback, param)
}

func OnWindowDestroy(webWindow WebView, callback WindowDestroyCallback, param unsafe.Pointer) {
	_wkeOnWindowDestroy.LoadOnce()
	_wkeOnWindowDestroy.Call(webWindow, callback, param)
}

func OnDraggableRegionsChanged(webView WebView, callback DraggableRegionsChangedCallback, param unsafe.Pointer) {
	_wkeOnDraggableRegionsChanged.LoadOnce()
	_wkeOnDraggableRegionsChanged.Call(webView, callback, param)
}

func OnWillMediaLoad(webView WebView, callback WillMediaLoadCallback, param unsafe.Pointer) {
	_wkeOnWillMediaLoad.LoadOnce()
	_wkeOnWillMediaLoad.Call(webView, callback, param)
}

func OnStartDragging(webView WebView, callback StartDraggingCallback, param unsafe.Pointer) {
	_wkeOnStartDragging.LoadOnce()
	_wkeOnStartDragging.Call(webView, callback, param)
}

func OnPrint(webView WebView, callback OnPrintCallback, param unsafe.Pointer) {
	_wkeOnPrint.LoadOnce()
	_wkeOnPrint.Call(webView, callback, param)
}

func Screenshot(webView WebView, settings *ScreenshotSettings, callback OnScreenshotCallback, param unsafe.Pointer) {
	_wkeScreenshot.LoadOnce()
	_wkeScreenshot.Call(webView, settings, callback, param)
}

func OnOtherLoad(webView WebView, callback OnOtherLoadCallback, param unsafe.Pointer) {
	_wkeOnOtherLoad.LoadOnce()
	_wkeOnOtherLoad.Call(webView, callback, param)
}

func OnContextMenuItemClick(webView WebView, callback OnContextMenuItemClickCallback, param unsafe.Pointer) {
	_wkeOnContextMenuItemClick.LoadOnce()
	_wkeOnContextMenuItemClick.Call(webView, callback, param)
}

func IsProcessingUserGesture(webView WebView) bool {
	_wkeIsProcessingUserGesture.LoadOnce()
	return _wkeIsProcessingUserGesture.Call(webView)
}

func NetSetMIMEType(jobPtr NetJob, type_ string) {
	_wkeNetSetMIMEType.LoadOnce()
	_wkeNetSetMIMEType.Call(jobPtr, type_)
}

func NetGetMIMEType(jobPtr NetJob, mime *String) string {
	_wkeNetGetMIMEType.LoadOnce()
	return _wkeNetGetMIMEType.Call(jobPtr, mime)
}

func NetGetReferrer(jobPtr NetJob) string {
	_wkeNetGetReferrer.LoadOnce()
	return _wkeNetGetReferrer.Call(jobPtr)
}

func NetSetHTTPHeaderField(jobPtr NetJob, key, value uint16, response bool) {
	_wkeNetSetHTTPHeaderField.LoadOnce()
	_wkeNetSetHTTPHeaderField.Call(jobPtr, key, value, response)
}

func NetGetHTTPHeaderField(jobPtr NetJob, key string) string {
	_wkeNetGetHTTPHeaderField.LoadOnce()
	return _wkeNetGetHTTPHeaderField.Call(jobPtr, key)
}

func NetGetHTTPHeaderFieldFromResponse(jobPtr NetJob, key string) string {
	_wkeNetGetHTTPHeaderFieldFromResponse.LoadOnce()
	return _wkeNetGetHTTPHeaderFieldFromResponse.Call(jobPtr, key)
}

func NetSetData(jobPtr NetJob, buf *byte, len int) {
	_wkeNetSetData.LoadOnce()
	_wkeNetSetData.Call(jobPtr, buf, len)
}

func NetHookRequest(jobPtr NetJob) {
	_wkeNetHookRequest.LoadOnce()
	_wkeNetHookRequest.Call(jobPtr)
}

func NetGetRequestMethod(jobPtr NetJob) RequestType {
	_wkeNetGetRequestMethod.LoadOnce()
	return _wkeNetGetRequestMethod.Call(jobPtr)
}

func NetContinueJob(jobPtr NetJob) {
	_wkeNetContinueJob.LoadOnce()
	_wkeNetContinueJob.Call(jobPtr)
}

func NetGetUrlByJob(jobPtr NetJob) string {
	_wkeNetGetUrlByJob.LoadOnce()
	return _wkeNetGetUrlByJob.Call(jobPtr)
}

func NetGetRawHttpHead(jobPtr NetJob) Slist {
	_wkeNetGetRawHttpHead.LoadOnce()
	return _wkeNetGetRawHttpHead.Call(jobPtr)
}

func NetGetRawResponseHead(jobPtr NetJob) Slist {
	_wkeNetGetRawResponseHead.LoadOnce()
	return _wkeNetGetRawResponseHead.Call(jobPtr)
}

func NetCancelRequest(jobPtr NetJob) {
	_wkeNetCancelRequest.LoadOnce()
	_wkeNetCancelRequest.Call(jobPtr)
}

func NetHoldJobToAsynCommit(jobPtr NetJob) bool {
	_wkeNetHoldJobToAsynCommit.LoadOnce()
	return _wkeNetHoldJobToAsynCommit.Call(jobPtr)
}

func NetOnResponse(webView WebView, callback NetResponseCallback, param unsafe.Pointer) {
	_wkeNetOnResponse.LoadOnce()
	_wkeNetOnResponse.Call(webView, callback, param)
}

func NetGetFavicon(webView WebView, callback OnNetGetFaviconCallback, param unsafe.Pointer) int {
	_wkeNetGetFavicon.LoadOnce()
	return _wkeNetGetFavicon.Call(webView, callback, param)
}

func NetCreateWebUrlRequest(url, method, mime string) WebUrlRequestPtr {
	_wkeNetCreateWebUrlRequest.LoadOnce()
	return _wkeNetCreateWebUrlRequest.Call(url, method, mime)
}

func NetCreateWebUrlRequest2(request blinkWebURLRequestPtr) WebUrlRequestPtr {
	_wkeNetCreateWebUrlRequest2.LoadOnce()
	return _wkeNetCreateWebUrlRequest2.Call(request)
}

func NetCopyWebUrlRequest(jobPtr NetJob, needExtraData bool) blinkWebURLRequestPtr {
	_wkeNetCopyWebUrlRequest.LoadOnce()
	return _wkeNetCopyWebUrlRequest.Call(jobPtr, needExtraData)
}

func NetDeleteBlinkWebURLRequestPtr(request blinkWebURLRequestPtr) {
	_wkeNetDeleteBlinkWebURLRequestPtr.LoadOnce()
	_wkeNetDeleteBlinkWebURLRequestPtr.Call(request)
}

func NetAddHTTPHeaderFieldToUrlRequest(request WebUrlRequestPtr, name, value string) {
	_wkeNetAddHTTPHeaderFieldToUrlRequest.LoadOnce()
	_wkeNetAddHTTPHeaderFieldToUrlRequest.Call(request, name, value)
}

func NetStartUrlRequest(webView WebView, request WebUrlRequestPtr, param unsafe.Pointer, callbacks *UrlRequestCallbacks) int {
	_wkeNetStartUrlRequest.LoadOnce()
	return _wkeNetStartUrlRequest.Call(webView, request, param, callbacks)
}

func NetGetHttpStatusCode(response WebUrlResponsePtr) int {
	_wkeNetGetHttpStatusCode.LoadOnce()
	return _wkeNetGetHttpStatusCode.Call(response)
}

func NetGetExpectedContentLength(response WebUrlResponsePtr) int64 {
	_wkeNetGetExpectedContentLength.LoadOnce()
	return _wkeNetGetExpectedContentLength.Call(response)
}

func NetGetResponseUrl(response WebUrlResponsePtr) string {
	_wkeNetGetResponseUrl.LoadOnce()
	return _wkeNetGetResponseUrl.Call(response)
}

func NetCancelWebUrlRequest(requestId int) {
	_wkeNetCancelWebUrlRequest.LoadOnce()
	_wkeNetCancelWebUrlRequest.Call(requestId)
}

func NetGetPostBody(jobPtr NetJob) *PostBodyElements {
	_wkeNetGetPostBody.LoadOnce()
	return _wkeNetGetPostBody.Call(jobPtr)
}

func NetFreePostBodyElements(elements *PostBodyElements) {
	_wkeNetFreePostBodyElements.LoadOnce()
	_wkeNetFreePostBodyElements.Call(elements)
}

func NetCreatePostBodyElements(webView WebView, length uintptr) *PostBodyElements {
	_wkeNetCreatePostBodyElements.LoadOnce()
	return _wkeNetCreatePostBodyElements.Call(webView, length)
}

func NetCreatePostBodyElement(webView WebView) *PostBodyElement {
	_wkeNetCreatePostBodyElement.LoadOnce()
	return _wkeNetCreatePostBodyElement.Call(webView)
}

func NetFreePostBodyElement(element *PostBodyElement) {
	_wkeNetFreePostBodyElement.LoadOnce()
	_wkeNetFreePostBodyElement.Call(element)
}

func PopupDialogAndDownload(webviewHandle WebView, dialogOpt *DialogOptions, expectedContentLength uintptr, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	_wkePopupDialogAndDownload.LoadOnce()
	return _wkePopupDialogAndDownload.Call(webviewHandle, dialogOpt, expectedContentLength, url, mime, disposition, job, dataBind, callbackBind)
}

func DownloadByPath(webviewHandle WebView, dialogOpt *DialogOptions, path []uint16, expectedContentLength uintptr, url, mime, disposition string, job NetJob, dataBind *NetJobDataBind, callbackBind *DownloadBind) DownloadOpt {
	_wkeDownloadByPath.LoadOnce()
	return _wkeDownloadByPath.Call(webviewHandle, dialogOpt, path, expectedContentLength, url, mime, disposition, job, dataBind, callbackBind)
}

func IsMainFrame(webView WebView, frameId WebFrameHandle) bool {
	_wkeIsMainFrame.LoadOnce()
	return _wkeIsMainFrame.Call(webView, frameId)
}

func IsWebRemoteFrame(webView WebView, frameId WebFrameHandle) bool {
	_wkeIsWebRemoteFrame.LoadOnce()
	return _wkeIsWebRemoteFrame.Call(webView, frameId)
}

func WebFrameGetMainFrame(webView WebView) WebFrameHandle {
	_wkeWebFrameGetMainFrame.LoadOnce()
	return _wkeWebFrameGetMainFrame.Call(webView)
}

func RunJsByFrame(webView WebView, frameId WebFrameHandle, script string, isInClosure bool) JsValue {
	_wkeRunJsByFrame.LoadOnce()
	return _wkeRunJsByFrame.Call(webView, frameId, script, isInClosure)
}

func InsertCSSByFrame(webView WebView, frameId WebFrameHandle, cssText string) {
	_wkeInsertCSSByFrame.LoadOnce()
	_wkeInsertCSSByFrame.Call(webView, frameId, cssText)
}

func WebFrameGetMainWorldScriptContext(webView WebView, webFrameId WebFrameHandle, contextOut *V8ContextPtr) {
	_wkeWebFrameGetMainWorldScriptContext.LoadOnce()
	_wkeWebFrameGetMainWorldScriptContext.Call(webView, webFrameId, contextOut)
}

func GetBlinkMainThreadIsolate() V8Isolate {
	_wkeGetBlinkMainThreadIsolate.LoadOnce()
	return _wkeGetBlinkMainThreadIsolate.Call()
}

func CreateWebWindow(typ WindowType, parent HWND, x, y, width, height int) WebView {
	_wkeCreateWebWindow.LoadOnce()
	return _wkeCreateWebWindow.Call(typ, parent, x, y, width, height)
}

func CreateWebCustomWindow(info *WindowCreateInfo) WebView {
	_wkeCreateWebCustomWindow.LoadOnce()
	return _wkeCreateWebCustomWindow.Call(info)
}

func DestroyWebWindow(webWindow WebView) {
	_wkeDestroyWebWindow.LoadOnce()
	_wkeDestroyWebWindow.Call(webWindow)
}

func GetWindowHandle(webWindow WebView) HWND {
	_wkeGetWindowHandle.LoadOnce()
	return _wkeGetWindowHandle.Call(webWindow)
}

func ShowWindow(webWindow WebView, show bool) {
	_wkeShowWindow.LoadOnce()
	_wkeShowWindow.Call(webWindow, show)
}

func EnableWindow(webWindow WebView, enable bool) {
	_wkeEnableWindow.LoadOnce()
	_wkeEnableWindow.Call(webWindow, enable)
}

func MoveWindow(webWindow WebView, x, y, width, height int) {
	_wkeMoveWindow.LoadOnce()
	_wkeMoveWindow.Call(webWindow, x, y, width, height)
}

func MoveToCenter(webWindow WebView) {
	_wkeMoveToCenter.LoadOnce()
	_wkeMoveToCenter.Call(webWindow)
}

func ResizeWindow(webWindow WebView, width, height int) {
	_wkeResizeWindow.LoadOnce()
	_wkeResizeWindow.Call(webWindow, width, height)
}

func DragTargetDragEnter(webView WebView, webDragData *WebDragData, clientPoint, screenPoint *Point, operationsAllowed WebDragOperationsMask, modifiers int) WebDragOperationsMask {
	_wkeDragTargetDragEnter.LoadOnce()
	return _wkeDragTargetDragEnter.Call(webView, webDragData, clientPoint, screenPoint, operationsAllowed, modifiers)
}

func DragTargetDragOver(webView WebView, clientPoint, screenPoint *Point, operationsAllowed WebDragOperationsMask, modifiers int) WebDragOperationsMask {
	_wkeDragTargetDragOver.LoadOnce()
	return _wkeDragTargetDragOver.Call(webView, clientPoint, screenPoint, operationsAllowed, modifiers)
}

func DragTargetDragLeave(webView WebView) {
	_wkeDragTargetDragLeave.LoadOnce()
	_wkeDragTargetDragLeave.Call(webView)
}

func DragTargetDrop(webView WebView, clientPoint, screenPoint *Point, modifiers int) {
	_wkeDragTargetDrop.LoadOnce()
	_wkeDragTargetDrop.Call(webView, clientPoint, screenPoint, modifiers)
}

func DragTargetEnd(webView WebView, clientPoint, screenPoint *Point, operation WebDragOperationsMask) {
	_wkeDragTargetEnd.LoadOnce()
	_wkeDragTargetEnd.Call(webView, clientPoint, screenPoint, operation)
}

func UtilSetUiCallback(callback UiThreadPostTaskCallback) {
	_wkeUtilSetUiCallback.LoadOnce()
	_wkeUtilSetUiCallback.Call(callback)
}

func UtilSerializeToMHTML(webView WebView) string {
	_wkeUtilSerializeToMHTML.LoadOnce()
	return _wkeUtilSerializeToMHTML.Call(webView)
}

func UtilPrintToPdf(webView WebView, frameId WebFrameHandle, settings *PrintSettings) *PdfDatas {
	_wkeUtilPrintToPdf.LoadOnce()
	return _wkeUtilPrintToPdf.Call(webView, frameId, settings)
}

func PrintToBitmap(webView WebView, frameId WebFrameHandle, settings *ScreenshotSettings) *MemBuf {
	_wkePrintToBitmap.LoadOnce()
	return _wkePrintToBitmap.Call(webView, frameId, settings)
}

func UtilRelasePrintPdfDatas(datas *PdfDatas) {
	_wkeUtilRelasePrintPdfDatas.LoadOnce()
	_wkeUtilRelasePrintPdfDatas.Call(datas)
}

func SetWindowTitle(webWindow WebView, title string) {
	_wkeSetWindowTitle.LoadOnce()
	_wkeSetWindowTitle.Call(webWindow, title)
}

func SetWindowTitleW(webWindow WebView, title uint16) {
	_wkeSetWindowTitleW.LoadOnce()
	_wkeSetWindowTitleW.Call(webWindow, title)
}

func NodeOnCreateProcess(webView WebView, callback NodeOnCreateProcessCallback, param unsafe.Pointer) {
	_wkeNodeOnCreateProcess.LoadOnce()
	_wkeNodeOnCreateProcess.Call(webView, callback, param)
}

func OnPluginFind(webView WebView, mime string, callback OnPluginFindCallback, param unsafe.Pointer) {
	_wkeOnPluginFind.LoadOnce()
	_wkeOnPluginFind.Call(webView, mime, callback, param)
}

func AddNpapiPlugin(webView WebView, initializeFunc, getEntryPointsFunc, shutdownFunc unsafe.Pointer) {
	_wkeAddNpapiPlugin.LoadOnce()
	_wkeAddNpapiPlugin.Call(webView, initializeFunc, getEntryPointsFunc, shutdownFunc)
}

func PluginListBuilderAddPlugin(builder unsafe.Pointer, name, description, fileName string) {
	_wkePluginListBuilderAddPlugin.LoadOnce()
	_wkePluginListBuilderAddPlugin.Call(builder, name, description, fileName)
}

func PluginListBuilderAddMediaTypeToLastPlugin(builder unsafe.Pointer, name, description string) {
	_wkePluginListBuilderAddMediaTypeToLastPlugin.LoadOnce()
	_wkePluginListBuilderAddMediaTypeToLastPlugin.Call(builder, name, description)
}

func PluginListBuilderAddFileExtensionToLastMediaType(builder unsafe.Pointer, fileExtension string) {
	_wkePluginListBuilderAddFileExtensionToLastMediaType.LoadOnce()
	_wkePluginListBuilderAddFileExtensionToLastMediaType.Call(builder, fileExtension)
}

func GetWebViewByNData(ndata unsafe.Pointer) WebView {
	_wkeGetWebViewByNData.LoadOnce()
	return _wkeGetWebViewByNData.Call(ndata)
}

func RegisterEmbedderCustomElement(webView WebView, frameId WebFrameHandle, name string, options, outResult unsafe.Pointer) bool {
	_wkeRegisterEmbedderCustomElement.LoadOnce()
	return _wkeRegisterEmbedderCustomElement.Call(webView, frameId, name, options, outResult)
}

func SetMediaPlayerFactory(webView WebView, factory MediaPlayerFactory, callback OnIsMediaPlayerSupportsMIMEType) {
	_wkeSetMediaPlayerFactory.LoadOnce()
	_wkeSetMediaPlayerFactory.Call(webView, factory, callback)
}

func GetContentAsMarkup(webView WebView, frame WebFrameHandle, size *uintptr) string {
	_wkeGetContentAsMarkup.LoadOnce()
	return _wkeGetContentAsMarkup.Call(webView, frame, size)
}

func UtilDecodeURLEscape(url string) string {
	_wkeUtilDecodeURLEscape.LoadOnce()
	return _wkeUtilDecodeURLEscape.Call(url)
}

func UtilEncodeURLEscape(url string) string {
	_wkeUtilEncodeURLEscape.LoadOnce()
	return _wkeUtilEncodeURLEscape.Call(url)
}

func UtilBase64Encode(str string) string {
	_wkeUtilBase64Encode.LoadOnce()
	return _wkeUtilBase64Encode.Call(str)
}

func UtilBase64Decode(str string) string {
	_wkeUtilBase64Decode.LoadOnce()
	return _wkeUtilBase64Decode.Call(str)
}

func UtilCreateV8Snapshot(str string) *MemBuf {
	_wkeUtilCreateV8Snapshot.LoadOnce()
	return _wkeUtilCreateV8Snapshot.Call(str)
}

func RunMessageLoop() {
	_wkeRunMessageLoop.LoadOnce()
	_wkeRunMessageLoop.Call()
}

func SaveMemoryCache(webView WebView) {
	_wkeSaveMemoryCache.LoadOnce()
	_wkeSaveMemoryCache.Call(webView)
}

func JsBindFunction(name string, fn wkeJsNativeFunction, param unsafe.Pointer, argCount uint) {
	_wkeJsBindFunction.LoadOnce()
	_wkeJsBindFunction.Call(name, fn, param, argCount)
}

func JsBindGetter(name string, fn wkeJsNativeFunction, param unsafe.Pointer) {
	_wkeJsBindGetter.LoadOnce()
	_wkeJsBindGetter.Call(name, fn, param)
}

func JsBindSetter(name string, fn wkeJsNativeFunction, param unsafe.Pointer) {
	_wkeJsBindSetter.LoadOnce()
	_wkeJsBindSetter.Call(name, fn, param)
}

func JsArgCount(es JsExecState) int {
	_jsArgCount.LoadOnce()
	return _jsArgCount.Call(es)
}

func JsArgType(es JsExecState, argIdx int) JsType {
	_jsArgType.LoadOnce()
	return _jsArgType.Call(es, argIdx)
}

func JsArg(es JsExecState, argIdx int) JsValue {
	_jsArg.LoadOnce()
	return _jsArg.Call(es, argIdx)
}

func JsTypeOf(v JsValue) JsType {
	_jsTypeOf.LoadOnce()
	return _jsTypeOf.Call(v)
}

func JsIsNumber(v JsValue) bool {
	_jsIsNumber.LoadOnce()
	return _jsIsNumber.Call(v)
}

func JsIsString(v JsValue) bool {
	_jsIsString.LoadOnce()
	return _jsIsString.Call(v)
}

func JsIsBoolean(v JsValue) bool {
	_jsIsBoolean.LoadOnce()
	return _jsIsBoolean.Call(v)
}

func JsIsObject(v JsValue) bool {
	_jsIsObject.LoadOnce()
	return _jsIsObject.Call(v)
}

func JsIsFunction(v JsValue) bool {
	_jsIsFunction.LoadOnce()
	return _jsIsFunction.Call(v)
}

func JsIsUndefined(v JsValue) bool {
	_jsIsUndefined.LoadOnce()
	return _jsIsUndefined.Call(v)
}

func JsIsNull(v JsValue) bool {
	_jsIsNull.LoadOnce()
	return _jsIsNull.Call(v)
}

func JsIsArray(v JsValue) bool {
	_jsIsArray.LoadOnce()
	return _jsIsArray.Call(v)
}

func JsIsTrue(v JsValue) bool {
	_jsIsTrue.LoadOnce()
	return _jsIsTrue.Call(v)
}

func JsIsFalse(v JsValue) bool {
	_jsIsFalse.LoadOnce()
	return _jsIsFalse.Call(v)
}

func JsToInt(es JsExecState, v JsValue) int {
	_jsToInt.LoadOnce()
	return _jsToInt.Call(es, v)
}

func JsToFloat(es JsExecState, v JsValue) float32 {
	_jsToFloat.LoadOnce()
	return _jsToFloat.Call(es, v)
}

func JsToDouble(es JsExecState, v JsValue) float64 {
	_jsToDouble.LoadOnce()
	return _jsToDouble.Call(es, v)
}

func JsToDoubleString(es JsExecState, v JsValue) string {
	_jsToDoubleString.LoadOnce()
	return _jsToDoubleString.Call(es, v)
}

func JsToBoolean(es JsExecState, v JsValue) bool {
	_jsToBoolean.LoadOnce()
	return _jsToBoolean.Call(es, v)
}

func JsArrayBuffer(es JsExecState, buffer *byte, size uintptr) JsValue {
	_jsArrayBuffer.LoadOnce()
	return _jsArrayBuffer.Call(es, buffer, size)
}

func JsGetArrayBuffer(es JsExecState, value JsValue) *MemBuf {
	_jsGetArrayBuffer.LoadOnce()
	return _jsGetArrayBuffer.Call(es, value)
}

func JsToTempString(es JsExecState, v JsValue) string {
	_jsToTempString.LoadOnce()
	return _jsToTempString.Call(es, v)
}

func JsToTempStringW(es JsExecState, v JsValue) uint16 {
	_jsToTempStringW.LoadOnce()
	return _jsToTempStringW.Call(es, v)
}

func JsToV8Value(es JsExecState, v JsValue) v8ValuePtr {
	_jsToV8Value.LoadOnce()
	return _jsToV8Value.Call(es, v)
}

func JsInt(n int) JsValue {
	_jsInt.LoadOnce()
	return _jsInt.Call(n)
}

func JsFloat(f float32) JsValue {
	_jsFloat.LoadOnce()
	return _jsFloat.Call(f)
}

func JsDouble(d float64) JsValue {
	_jsDouble.LoadOnce()
	return _jsDouble.Call(d)
}

func JsDoubleString(str string) JsValue {
	_jsDoubleString.LoadOnce()
	return _jsDoubleString.Call(str)
}

func JsBoolean(b bool) JsValue {
	_jsBoolean.LoadOnce()
	return _jsBoolean.Call(b)
}

func JsUndefined() JsValue {
	_jsUndefined.LoadOnce()
	return _jsUndefined.Call()
}

func JsNull() JsValue {
	_jsNull.LoadOnce()
	return _jsNull.Call()
}

func JsTrue() JsValue {
	_jsTrue.LoadOnce()
	return _jsTrue.Call()
}

func JsFalse() JsValue {
	_jsFalse.LoadOnce()
	return _jsFalse.Call()
}

func JsString(es JsExecState, str string) JsValue {
	_jsString.LoadOnce()
	return _jsString.Call(es, str)
}

func JsStringW(es JsExecState, str uint16) JsValue {
	_jsStringW.LoadOnce()
	return _jsStringW.Call(es, str)
}

func JsEmptyObject(es JsExecState) JsValue {
	_jsEmptyObject.LoadOnce()
	return _jsEmptyObject.Call(es)
}

func JsEmptyArray(es JsExecState) JsValue {
	_jsEmptyArray.LoadOnce()
	return _jsEmptyArray.Call(es)
}

func JsObject(es JsExecState, obj JsData) JsValue {
	_jsObject.LoadOnce()
	return _jsObject.Call(es, obj)
}

func JsFunction(es JsExecState, obj JsData) JsValue {
	_jsFunction.LoadOnce()
	return _jsFunction.Call(es, obj)
}

func JsGetData(es JsExecState, object JsValue) JsData {
	_jsGetData.LoadOnce()
	return _jsGetData.Call(es, object)
}

func JsGet(es JsExecState, object JsValue, prop string) JsValue {
	_jsGet.LoadOnce()
	return _jsGet.Call(es, object, prop)
}

func JsSet(es JsExecState, object JsValue, prop string, v JsValue) {
	_jsSet.LoadOnce()
	_jsSet.Call(es, object, prop, v)
}

func JsGetAt(es JsExecState, object JsValue, index int) JsValue {
	_jsGetAt.LoadOnce()
	return _jsGetAt.Call(es, object, index)
}

func JsSetAt(es JsExecState, object JsValue, index int, v JsValue) {
	_jsSetAt.LoadOnce()
	_jsSetAt.Call(es, object, index, v)
}

func JsGetKeys(es JsExecState, object JsValue) JsKeys {
	_jsGetKeys.LoadOnce()
	return _jsGetKeys.Call(es, object)
}

func JsIsJsValueValid(es JsExecState, object JsValue) bool {
	_jsIsJsValueValid.LoadOnce()
	return _jsIsJsValueValid.Call(es, object)
}

func JsIsValidExecState(es JsExecState) bool {
	_jsIsValidExecState.LoadOnce()
	return _jsIsValidExecState.Call(es)
}

func JsDeleteObjectProp(es JsExecState, object JsValue, prop string) {
	_jsDeleteObjectProp.LoadOnce()
	_jsDeleteObjectProp.Call(es, object, prop)
}

func JsGetLength(es JsExecState, object JsValue) int {
	_jsGetLength.LoadOnce()
	return _jsGetLength.Call(es, object)
}

func JsSetLength(es JsExecState, object JsValue, length int) {
	_jsSetLength.LoadOnce()
	_jsSetLength.Call(es, object, length)
}

func JsGlobalObject(es JsExecState) JsValue {
	_jsGlobalObject.LoadOnce()
	return _jsGlobalObject.Call(es)
}

func JsGetWebView(es JsExecState) WebView {
	_jsGetWebView.LoadOnce()
	return _jsGetWebView.Call(es)
}

func JsEval(es JsExecState, str string) JsValue {
	_jsEval.LoadOnce()
	return _jsEval.Call(es, str)
}

func JsEvalW(es JsExecState, str uint16) JsValue {
	_jsEvalW.LoadOnce()
	return _jsEvalW.Call(es, str)
}

func JsEvalExW(es JsExecState, str uint16, isInClosure bool) JsValue {
	_jsEvalExW.LoadOnce()
	return _jsEvalExW.Call(es, str, isInClosure)
}

func JsCall(es JsExecState, funcObj, thisObject JsValue, args []JsValue, argCount int) JsValue {
	_jsCall.LoadOnce()
	return _jsCall.Call(es, funcObj, thisObject, args, argCount)
}

func JsCallGlobal(es JsExecState, funcObj JsValue, args []JsValue, argCount int) JsValue {
	_jsCallGlobal.LoadOnce()
	return _jsCallGlobal.Call(es, funcObj, args, argCount)
}

func JsGetGlobal(es JsExecState, prop string) JsValue {
	_jsGetGlobal.LoadOnce()
	return _jsGetGlobal.Call(es, prop)
}

func JsSetGlobal(es JsExecState, prop string, v JsValue) {
	_jsSetGlobal.LoadOnce()
	_jsSetGlobal.Call(es, prop, v)
}

func JsGC() {
	_jsGC.LoadOnce()
	_jsGC.Call()
}

func JsAddRef(es JsExecState, val JsValue) bool {
	_jsAddRef.LoadOnce()
	return _jsAddRef.Call(es, val)
}

func JsReleaseRef(es JsExecState, val JsValue) bool {
	_jsReleaseRef.LoadOnce()
	return _jsReleaseRef.Call(es, val)
}

func JsGetLastErrorIfException(es JsExecState) *JsExceptionInfo {
	_jsGetLastErrorIfException.LoadOnce()
	return _jsGetLastErrorIfException.Call(es)
}

func JsThrowException(es JsExecState, exception string) JsValue {
	_jsThrowException.LoadOnce()
	return _jsThrowException.Call(es, exception)
}

func JsGetCallstack(es JsExecState) string {
	_jsGetCallstack.LoadOnce()
	return _jsGetCallstack.Call(es)
}

func Init() {
	_wkeInit.LoadOnce()
	_wkeInit.Call()
}

func Initialize() {
	_wkeInitialize.LoadOnce()
	_wkeInitialize.Call()
}

func InitializeEx(settings *Settings) {
	_wkeInitializeEx.LoadOnce()
	_wkeInitializeEx.Call(settings)
}
