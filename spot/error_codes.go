package spot

import "github.com/GTedZ/binancego/internal/berror"

var ErrorCodes = errorCodeNames{
	UNKNOWN:                             UNKNOWN,
	DISCONNECTED:                        DISCONNECTED,
	UNAUTHORIZED:                        UNAUTHORIZED,
	TOO_MANY_REQUESTS:                   TOO_MANY_REQUESTS,
	UNEXPECTED_RESP:                     UNEXPECTED_RESP,
	TIMEOUT:                             TIMEOUT,
	SERVER_BUSY:                         SERVER_BUSY,
	INVALID_MESSAGE:                     INVALID_MESSAGE,
	UNKNOWN_ORDER_COMPOSITION:           UNKNOWN_ORDER_COMPOSITION,
	TOO_MANY_ORDERS:                     TOO_MANY_ORDERS,
	SERVICE_SHUTTING_DOWN:               SERVICE_SHUTTING_DOWN,
	UNSUPPORTED_OPERATION:               UNSUPPORTED_OPERATION,
	INVALID_TIMESTAMP:                   INVALID_TIMESTAMP,
	INVALID_SIGNATURE:                   INVALID_SIGNATURE,
	COMP_ID_IN_USE:                      COMP_ID_IN_USE,
	TOO_MANY_CONNECTIONS:                TOO_MANY_CONNECTIONS,
	LOGGED_OUT:                          LOGGED_OUT,
	ILLEGAL_CHARS:                       ILLEGAL_CHARS,
	TOO_MANY_PARAMETERS:                 TOO_MANY_PARAMETERS,
	MANDATORY_PARAM_EMPTY_OR_MALFORMED:  MANDATORY_PARAM_EMPTY_OR_MALFORMED,
	UNKNOWN_PARAM:                       UNKNOWN_PARAM,
	UNREAD_PARAMETERS:                   UNREAD_PARAMETERS,
	PARAM_EMPTY:                         PARAM_EMPTY,
	PARAM_NOT_REQUIRED:                  PARAM_NOT_REQUIRED,
	PARAM_OVERFLOW:                      PARAM_OVERFLOW,
	BAD_PRECISION:                       BAD_PRECISION,
	NO_DEPTH:                            NO_DEPTH,
	TIF_NOT_REQUIRED:                    TIF_NOT_REQUIRED,
	INVALID_TIF:                         INVALID_TIF,
	INVALID_ORDER_TYPE:                  INVALID_ORDER_TYPE,
	INVALID_SIDE:                        INVALID_SIDE,
	EMPTY_NEW_CL_ORD_ID:                 EMPTY_NEW_CL_ORD_ID,
	EMPTY_ORG_CL_ORD_ID:                 EMPTY_ORG_CL_ORD_ID,
	BAD_INTERVAL:                        BAD_INTERVAL,
	BAD_SYMBOL:                          BAD_SYMBOL,
	INVALID_SYMBOLSTATUS:                INVALID_SYMBOLSTATUS,
	INVALID_LISTEN_KEY:                  INVALID_LISTEN_KEY,
	MORE_THAN_XX_HOURS:                  MORE_THAN_XX_HOURS,
	OPTIONAL_PARAMS_BAD_COMBO:           OPTIONAL_PARAMS_BAD_COMBO,
	INVALID_PARAMETER:                   INVALID_PARAMETER,
	BAD_STRATEGY_TYPE:                   BAD_STRATEGY_TYPE,
	INVALID_JSON:                        INVALID_JSON,
	INVALID_TICKER_TYPE:                 INVALID_TICKER_TYPE,
	INVALID_CANCEL_RESTRICTIONS:         INVALID_CANCEL_RESTRICTIONS,
	DUPLICATE_SYMBOLS:                   DUPLICATE_SYMBOLS,
	INVALID_SBE_HEADER:                  INVALID_SBE_HEADER,
	UNSUPPORTED_SCHEMA_ID:               UNSUPPORTED_SCHEMA_ID,
	SBE_DISABLED:                        SBE_DISABLED,
	OCO_ORDER_TYPE_REJECTED:             OCO_ORDER_TYPE_REJECTED,
	OCO_ICEBERGQTY_TIMEINFORCE:          OCO_ICEBERGQTY_TIMEINFORCE,
	DEPRECATED_SCHEMA:                   DEPRECATED_SCHEMA,
	BUY_OCO_LIMIT_MUST_BE_BELOW:         BUY_OCO_LIMIT_MUST_BE_BELOW,
	SELL_OCO_LIMIT_MUST_BE_ABOVE:        SELL_OCO_LIMIT_MUST_BE_ABOVE,
	BOTH_OCO_ORDERS_CANNOT_BE_LIMIT:     BOTH_OCO_ORDERS_CANNOT_BE_LIMIT,
	INVALID_TAG_NUMBER:                  INVALID_TAG_NUMBER,
	TAG_NOT_DEFINED_IN_MESSAGE:          TAG_NOT_DEFINED_IN_MESSAGE,
	TAG_APPEARS_MORE_THAN_ONCE:          TAG_APPEARS_MORE_THAN_ONCE,
	TAG_OUT_OF_ORDER:                    TAG_OUT_OF_ORDER,
	GROUP_FIELDS_OUT_OF_ORDER:           GROUP_FIELDS_OUT_OF_ORDER,
	INVALID_COMPONENT:                   INVALID_COMPONENT,
	RESET_SEQ_NUM_SUPPORT:               RESET_SEQ_NUM_SUPPORT,
	ALREADY_LOGGED_IN:                   ALREADY_LOGGED_IN,
	GARBLED_MESSAGE:                     GARBLED_MESSAGE,
	BAD_SENDER_COMPID:                   BAD_SENDER_COMPID,
	BAD_SEQ_NUM:                         BAD_SEQ_NUM,
	EXPECTED_LOGON:                      EXPECTED_LOGON,
	TOO_MANY_MESSAGES:                   TOO_MANY_MESSAGES,
	PARAMS_BAD_COMBO:                    PARAMS_BAD_COMBO,
	NOT_ALLOWED_IN_DROP_COPY_SESSIONS:   NOT_ALLOWED_IN_DROP_COPY_SESSIONS,
	DROP_COPY_SESSION_NOT_ALLOWED:       DROP_COPY_SESSION_NOT_ALLOWED,
	DROP_COPY_SESSION_REQUIRED:          DROP_COPY_SESSION_REQUIRED,
	NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS: NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS,
	NOT_ALLOWED_IN_MARKET_DATA_SESSIONS: NOT_ALLOWED_IN_MARKET_DATA_SESSIONS,
	INCORRECT_NUM_IN_GROUP_COUNT:        INCORRECT_NUM_IN_GROUP_COUNT,
	DUPLICATE_ENTRIES_IN_A_GROUP:        DUPLICATE_ENTRIES_IN_A_GROUP,
	INVALID_REQUEST_ID:                  INVALID_REQUEST_ID,
	TOO_MANY_SUBSCRIPTIONS:              TOO_MANY_SUBSCRIPTIONS,
	INVALID_TIME_UNIT:                   INVALID_TIME_UNIT,
	BUY_OCO_STOP_LOSS_MUST_BE_ABOVE:     BUY_OCO_STOP_LOSS_MUST_BE_ABOVE,
	SELL_OCO_STOP_LOSS_MUST_BE_BELOW:    SELL_OCO_STOP_LOSS_MUST_BE_BELOW,
	BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW:   BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW,
	SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE:  SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE,
	NEW_ORDER_REJECTED:                  NEW_ORDER_REJECTED,
	CANCEL_REJECTED:                     CANCEL_REJECTED,
	NO_SUCH_ORDER:                       NO_SUCH_ORDER,
	BAD_API_KEY_FMT:                     BAD_API_KEY_FMT,
	REJECTED_MBX_KEY:                    REJECTED_MBX_KEY,
	NO_TRADING_WINDOW:                   NO_TRADING_WINDOW,
	ORDER_ARCHIVED:                      ORDER_ARCHIVED,
	SUBSCRIPTION_ACTIVE:                 SUBSCRIPTION_ACTIVE,
	SUBSCRIPTION_INACTIVE:               SUBSCRIPTION_INACTIVE,
}

type errorCodeNames struct {
	UNKNOWN                             berror.ErrorCode
	DISCONNECTED                        berror.ErrorCode
	UNAUTHORIZED                        berror.ErrorCode
	TOO_MANY_REQUESTS                   berror.ErrorCode
	UNEXPECTED_RESP                     berror.ErrorCode
	TIMEOUT                             berror.ErrorCode
	SERVER_BUSY                         berror.ErrorCode
	INVALID_MESSAGE                     berror.ErrorCode
	UNKNOWN_ORDER_COMPOSITION           berror.ErrorCode
	TOO_MANY_ORDERS                     berror.ErrorCode
	SERVICE_SHUTTING_DOWN               berror.ErrorCode
	UNSUPPORTED_OPERATION               berror.ErrorCode
	INVALID_TIMESTAMP                   berror.ErrorCode
	INVALID_SIGNATURE                   berror.ErrorCode
	COMP_ID_IN_USE                      berror.ErrorCode
	TOO_MANY_CONNECTIONS                berror.ErrorCode
	LOGGED_OUT                          berror.ErrorCode
	ILLEGAL_CHARS                       berror.ErrorCode
	TOO_MANY_PARAMETERS                 berror.ErrorCode
	MANDATORY_PARAM_EMPTY_OR_MALFORMED  berror.ErrorCode
	UNKNOWN_PARAM                       berror.ErrorCode
	UNREAD_PARAMETERS                   berror.ErrorCode
	PARAM_EMPTY                         berror.ErrorCode
	PARAM_NOT_REQUIRED                  berror.ErrorCode
	PARAM_OVERFLOW                      berror.ErrorCode
	BAD_PRECISION                       berror.ErrorCode
	NO_DEPTH                            berror.ErrorCode
	TIF_NOT_REQUIRED                    berror.ErrorCode
	INVALID_TIF                         berror.ErrorCode
	INVALID_ORDER_TYPE                  berror.ErrorCode
	INVALID_SIDE                        berror.ErrorCode
	EMPTY_NEW_CL_ORD_ID                 berror.ErrorCode
	EMPTY_ORG_CL_ORD_ID                 berror.ErrorCode
	BAD_INTERVAL                        berror.ErrorCode
	BAD_SYMBOL                          berror.ErrorCode
	INVALID_SYMBOLSTATUS                berror.ErrorCode
	INVALID_LISTEN_KEY                  berror.ErrorCode
	MORE_THAN_XX_HOURS                  berror.ErrorCode
	OPTIONAL_PARAMS_BAD_COMBO           berror.ErrorCode
	INVALID_PARAMETER                   berror.ErrorCode
	BAD_STRATEGY_TYPE                   berror.ErrorCode
	INVALID_JSON                        berror.ErrorCode
	INVALID_TICKER_TYPE                 berror.ErrorCode
	INVALID_CANCEL_RESTRICTIONS         berror.ErrorCode
	DUPLICATE_SYMBOLS                   berror.ErrorCode
	INVALID_SBE_HEADER                  berror.ErrorCode
	UNSUPPORTED_SCHEMA_ID               berror.ErrorCode
	SBE_DISABLED                        berror.ErrorCode
	OCO_ORDER_TYPE_REJECTED             berror.ErrorCode
	OCO_ICEBERGQTY_TIMEINFORCE          berror.ErrorCode
	DEPRECATED_SCHEMA                   berror.ErrorCode
	BUY_OCO_LIMIT_MUST_BE_BELOW         berror.ErrorCode
	SELL_OCO_LIMIT_MUST_BE_ABOVE        berror.ErrorCode
	BOTH_OCO_ORDERS_CANNOT_BE_LIMIT     berror.ErrorCode
	INVALID_TAG_NUMBER                  berror.ErrorCode
	TAG_NOT_DEFINED_IN_MESSAGE          berror.ErrorCode
	TAG_APPEARS_MORE_THAN_ONCE          berror.ErrorCode
	TAG_OUT_OF_ORDER                    berror.ErrorCode
	GROUP_FIELDS_OUT_OF_ORDER           berror.ErrorCode
	INVALID_COMPONENT                   berror.ErrorCode
	RESET_SEQ_NUM_SUPPORT               berror.ErrorCode
	ALREADY_LOGGED_IN                   berror.ErrorCode
	GARBLED_MESSAGE                     berror.ErrorCode
	BAD_SENDER_COMPID                   berror.ErrorCode
	BAD_SEQ_NUM                         berror.ErrorCode
	EXPECTED_LOGON                      berror.ErrorCode
	TOO_MANY_MESSAGES                   berror.ErrorCode
	PARAMS_BAD_COMBO                    berror.ErrorCode
	NOT_ALLOWED_IN_DROP_COPY_SESSIONS   berror.ErrorCode
	DROP_COPY_SESSION_NOT_ALLOWED       berror.ErrorCode
	DROP_COPY_SESSION_REQUIRED          berror.ErrorCode
	NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS berror.ErrorCode
	NOT_ALLOWED_IN_MARKET_DATA_SESSIONS berror.ErrorCode
	INCORRECT_NUM_IN_GROUP_COUNT        berror.ErrorCode
	DUPLICATE_ENTRIES_IN_A_GROUP        berror.ErrorCode
	INVALID_REQUEST_ID                  berror.ErrorCode
	TOO_MANY_SUBSCRIPTIONS              berror.ErrorCode
	INVALID_TIME_UNIT                   berror.ErrorCode
	BUY_OCO_STOP_LOSS_MUST_BE_ABOVE     berror.ErrorCode
	SELL_OCO_STOP_LOSS_MUST_BE_BELOW    berror.ErrorCode
	BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW   berror.ErrorCode
	SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE  berror.ErrorCode
	NEW_ORDER_REJECTED                  berror.ErrorCode
	CANCEL_REJECTED                     berror.ErrorCode
	NO_SUCH_ORDER                       berror.ErrorCode
	BAD_API_KEY_FMT                     berror.ErrorCode
	REJECTED_MBX_KEY                    berror.ErrorCode
	NO_TRADING_WINDOW                   berror.ErrorCode
	ORDER_ARCHIVED                      berror.ErrorCode
	SUBSCRIPTION_ACTIVE                 berror.ErrorCode
	SUBSCRIPTION_INACTIVE               berror.ErrorCode
}

var (
	UNKNOWN                             = newErrorCode(-1000, []string{"An unknown error occurred while processing the request."})
	DISCONNECTED                        = newErrorCode(-1001, []string{"Internal error; unable to process your request. Please try again."})
	UNAUTHORIZED                        = newErrorCode(-1002, []string{"You are not authorized to execute this request."})
	TOO_MANY_REQUESTS                   = newErrorCode(-1003, []string{"Too many requests queued.", "Too much request weight used; current limit is %s request weight per %s. Please use WebSocket Streams for live updates to avoid polling the API.", "Way too much request weight used; IP banned until %s. Please use WebSocket Streams for live updates to avoid bans."})
	UNEXPECTED_RESP                     = newErrorCode(-1006, []string{"An unexpected response was received from the message bus. Execution status unknown."})
	TIMEOUT                             = newErrorCode(-1007, []string{"Timeout waiting for response from backend server. Send status unknown; execution status unknown."})
	SERVER_BUSY                         = newErrorCode(-1008, []string{"Server is currently overloaded with other requests. Please try again in a few minutes."})
	INVALID_MESSAGE                     = newErrorCode(-1013, []string{"The request is rejected by the API. (i.e. The request didn't reach the Matching Engine.)", "Potential error messages can be found in Filter Failures or Failures during order placement."})
	UNKNOWN_ORDER_COMPOSITION           = newErrorCode(-1014, []string{"Unsupported order combination."})
	TOO_MANY_ORDERS                     = newErrorCode(-1015, []string{"Too many new orders.", "Too many new orders; current limit is %s orders per %s."})
	SERVICE_SHUTTING_DOWN               = newErrorCode(-1016, []string{"This service is no longer available."})
	UNSUPPORTED_OPERATION               = newErrorCode(-1020, []string{"This operation is not supported."})
	INVALID_TIMESTAMP                   = newErrorCode(-1021, []string{"Timestamp for this request is outside of the recvWindow.", "Timestamp for this request was 1000ms ahead of the server's time."})
	INVALID_SIGNATURE                   = newErrorCode(-1022, []string{"Signature for this request is not valid."})
	COMP_ID_IN_USE                      = newErrorCode(-1033, []string{"SenderCompId(49) is currently in use. Concurrent use of the same SenderCompId within one account is not allowed."})
	TOO_MANY_CONNECTIONS                = newErrorCode(-1034, []string{"Too many concurrent connections; current limit is '%s'.", "Too many connection attempts for account; current limit is %s per '%s'.", "Too many connection attempts from IP; current limit is %s per '%s'."})
	LOGGED_OUT                          = newErrorCode(-1035, []string{"Please send Logout<5> message to close the session."})
	ILLEGAL_CHARS                       = newErrorCode(-1100, []string{"Illegal characters found in a parameter.", "Illegal characters found in parameter '%s'; legal range is '%s'."})
	TOO_MANY_PARAMETERS                 = newErrorCode(-1101, []string{"Too many parameters sent for this endpoint.", "Too many parameters; expected '%s' and received '%s'.", "Duplicate values for a parameter detected."})
	MANDATORY_PARAM_EMPTY_OR_MALFORMED  = newErrorCode(-1102, []string{"A mandatory parameter was not sent, was empty/null, or malformed.", "Mandatory parameter '%s' was not sent, was empty/null, or malformed.", "Param '%s' or '%s' must be sent, but both were empty/null!", "Required tag '%s' missing.", "Field value was empty or malformed.", "'%s' contains unexpected value. Cannot be greater than %s."})
	UNKNOWN_PARAM                       = newErrorCode(-1103, []string{"An unknown parameter was sent.", "Undefined Tag."})
	UNREAD_PARAMETERS                   = newErrorCode(-1104, []string{"Not all sent parameters were read.", "Not all sent parameters were read; read '%s' parameter(s) but was sent '%s'."})
	PARAM_EMPTY                         = newErrorCode(-1105, []string{"A parameter was empty.", "Parameter '%s' was empty."})
	PARAM_NOT_REQUIRED                  = newErrorCode(-1106, []string{"A parameter was sent when not required.", "Parameter '%s' sent when not required.", "A tag '%s' was sent when not required."})
	PARAM_OVERFLOW                      = newErrorCode(-1108, []string{"Parameter '%s' overflowed."})
	BAD_PRECISION                       = newErrorCode(-1111, []string{"Parameter '%s' has too much precision."})
	NO_DEPTH                            = newErrorCode(-1112, []string{"No orders on book for symbol."})
	TIF_NOT_REQUIRED                    = newErrorCode(-1114, []string{"TimeInForce parameter sent when not required."})
	INVALID_TIF                         = newErrorCode(-1115, []string{"Invalid timeInForce."})
	INVALID_ORDER_TYPE                  = newErrorCode(-1116, []string{"Invalid orderType."})
	INVALID_SIDE                        = newErrorCode(-1117, []string{"Invalid side."})
	EMPTY_NEW_CL_ORD_ID                 = newErrorCode(-1118, []string{"New client order ID was empty."})
	EMPTY_ORG_CL_ORD_ID                 = newErrorCode(-1119, []string{"Original client order ID was empty."})
	BAD_INTERVAL                        = newErrorCode(-1120, []string{"Invalid interval."})
	BAD_SYMBOL                          = newErrorCode(-1121, []string{"Invalid symbol."})
	INVALID_SYMBOLSTATUS                = newErrorCode(-1122, []string{"Invalid symbolStatus."})
	INVALID_LISTEN_KEY                  = newErrorCode(-1125, []string{"This listenKey does not exist."})
	MORE_THAN_XX_HOURS                  = newErrorCode(-1127, []string{"Lookup interval is too big.", "More than %s hours between startTime and endTime."})
	OPTIONAL_PARAMS_BAD_COMBO           = newErrorCode(-1128, []string{"Combination of optional parameters invalid.", "Combination of optional fields invalid. Recommendation: '%s' and '%s' must both be sent.", "Fields [%s] must be sent together or omitted entirely.", "Invalid MDEntryType (269) combination. BID and OFFER must be requested together."})
	INVALID_PARAMETER                   = newErrorCode(-1130, []string{"Invalid data sent for a parameter.", "Data sent for parameter '%s' is not valid."})
	BAD_STRATEGY_TYPE                   = newErrorCode(-1134, []string{"strategyType was less than 1000000.", "TargetStrategy (847) was less than 1000000."})
	INVALID_JSON                        = newErrorCode(-1135, []string{"Invalid JSON Request", "JSON sent for parameter '%s' is not valid"})
	INVALID_TICKER_TYPE                 = newErrorCode(-1139, []string{"Invalid ticker type."})
	INVALID_CANCEL_RESTRICTIONS         = newErrorCode(-1145, []string{"cancelRestrictions has to be either ONLY_NEW or ONLY_PARTIALLY_FILLED."})
	DUPLICATE_SYMBOLS                   = newErrorCode(-1151, []string{"Symbol is present multiple times in the list."})
	INVALID_SBE_HEADER                  = newErrorCode(-1152, []string{"Invalid X-MBX-SBE header; expected <SCHEMA_ID>:<VERSION>."})
	UNSUPPORTED_SCHEMA_ID               = newErrorCode(-1153, []string{"Unsupported SBE schema ID or version specified in the X-MBX-SBE header."})
	SBE_DISABLED                        = newErrorCode(-1155, []string{"SBE is not enabled."})
	OCO_ORDER_TYPE_REJECTED             = newErrorCode(-1158, []string{"Order type not supported in OCO.", "If the order type provided in the aboveType and/or belowType is not supported."})
	OCO_ICEBERGQTY_TIMEINFORCE          = newErrorCode(-1160, []string{"Parameter '%s' is not supported if aboveTimeInForce/belowTimeInForce is not GTC.", "If the order type for the above or below leg is STOP_LOSS_LIMIT, and icebergQty is provided for that leg, the timeInForce has to be GTC else it will throw an error.", "TimeInForce (59) must be GTC (1) when MaxFloor (111) is used."})
	DEPRECATED_SCHEMA                   = newErrorCode(-1161, []string{"Unable to encode the response in SBE schema 'x'. Please use schema 'y' or higher."})
	BUY_OCO_LIMIT_MUST_BE_BELOW         = newErrorCode(-1165, []string{"A limit order in a buy OCO must be below."})
	SELL_OCO_LIMIT_MUST_BE_ABOVE        = newErrorCode(-1166, []string{"A limit order in a sell OCO must be above."})
	BOTH_OCO_ORDERS_CANNOT_BE_LIMIT     = newErrorCode(-1168, []string{"At least one OCO order must be contingent."})
	INVALID_TAG_NUMBER                  = newErrorCode(-1169, []string{"Invalid tag number."})
	TAG_NOT_DEFINED_IN_MESSAGE          = newErrorCode(-1170, []string{"Tag '%s' not defined for this message type."})
	TAG_APPEARS_MORE_THAN_ONCE          = newErrorCode(-1171, []string{"Tag '%s' appears more than once."})
	TAG_OUT_OF_ORDER                    = newErrorCode(-1172, []string{"Tag '%s' specified out of required order."})
	GROUP_FIELDS_OUT_OF_ORDER           = newErrorCode(-1173, []string{"Repeating group '%s' fields out of order."})
	INVALID_COMPONENT                   = newErrorCode(-1174, []string{"Component '%s' is incorrectly populated on '%s' order. Recommendation: '%s'"})
	RESET_SEQ_NUM_SUPPORT               = newErrorCode(-1175, []string{"Continuation of sequence numbers to new session is currently unsupported. Sequence numbers must be reset for each new session."})
	ALREADY_LOGGED_IN                   = newErrorCode(-1176, []string{"Logon<A> should only be sent once."})
	GARBLED_MESSAGE                     = newErrorCode(-1177, []string{"CheckSum(10) contains an incorrect value.", "BeginString (8) is not the first tag in a message.", "MsgType (35) is not the third tag in a message.", "BodyLength (9) does not contain the correct byte count.", "Only printable ASCII characters and SOH (Start of Header) are allowed."})
	BAD_SENDER_COMPID                   = newErrorCode(-1178, []string{"SenderCompId(49) contains an incorrect value. The SenderCompID value should not change throughout the lifetime of a session."})
	BAD_SEQ_NUM                         = newErrorCode(-1179, []string{"MsgSeqNum(34) contains an unexpected value. Expected: '%d'."})
	EXPECTED_LOGON                      = newErrorCode(-1180, []string{"Logon<A> must be the first message in the session."})
	TOO_MANY_MESSAGES                   = newErrorCode(-1181, []string{"Too many messages; current limit is '%d' messages per '%s'."})
	PARAMS_BAD_COMBO                    = newErrorCode(-1182, []string{"Conflicting fields: [%s]"})
	NOT_ALLOWED_IN_DROP_COPY_SESSIONS   = newErrorCode(-1183, []string{"Requested operation is not allowed in DropCopy sessions."})
	DROP_COPY_SESSION_NOT_ALLOWED       = newErrorCode(-1184, []string{"DropCopy sessions are not supported on this server. Please reconnect to a drop copy server."})
	DROP_COPY_SESSION_REQUIRED          = newErrorCode(-1185, []string{"Only DropCopy sessions are supported on this server. Either reconnect to order entry server or send DropCopyFlag (9406) field."})
	NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS = newErrorCode(-1186, []string{"Requested operation is not allowed in order entry sessions."})
	NOT_ALLOWED_IN_MARKET_DATA_SESSIONS = newErrorCode(-1187, []string{"Requested operation is not allowed in market data sessions."})
	INCORRECT_NUM_IN_GROUP_COUNT        = newErrorCode(-1188, []string{"Incorrect NumInGroup count for repeating group '%s'."})
	DUPLICATE_ENTRIES_IN_A_GROUP        = newErrorCode(-1189, []string{"Group '%s' contains duplicate entries."})
	INVALID_REQUEST_ID                  = newErrorCode(-1190, []string{"MDReqID (262) contains a subscription request id that is already in use on this connection.", "MDReqID (262) contains an unsubscription request id that does not match any active subscription."})
	TOO_MANY_SUBSCRIPTIONS              = newErrorCode(-1191, []string{"Too many subscriptions. Connection may create up to '%s' subscriptions at a time.", "Similar subscription is already active on this connection. Symbol='%s', active subscription id: '%s'."})
	INVALID_TIME_UNIT                   = newErrorCode(-1194, []string{"Invalid value for time unit; expected either MICROSECOND or MILLISECOND."})
	BUY_OCO_STOP_LOSS_MUST_BE_ABOVE     = newErrorCode(-1196, []string{"A stop loss order in a buy OCO must be above."})
	SELL_OCO_STOP_LOSS_MUST_BE_BELOW    = newErrorCode(-1197, []string{"A stop loss order in a sell OCO must be below."})
	BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW   = newErrorCode(-1198, []string{"A take profit order in a buy OCO must be below."})
	SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE  = newErrorCode(-1199, []string{"A take profit order in a sell OCO must be above."})
	NEW_ORDER_REJECTED                  = newErrorCode(-2010, []string{"NEW_ORDER_REJECTED"})
	CANCEL_REJECTED                     = newErrorCode(-2011, []string{"CANCEL_REJECTED"})
	NO_SUCH_ORDER                       = newErrorCode(-2013, []string{"Order does not exist."})
	BAD_API_KEY_FMT                     = newErrorCode(-2014, []string{"API-key format invalid."})
	REJECTED_MBX_KEY                    = newErrorCode(-2015, []string{"Invalid API-key, IP, or permissions for action."})
	NO_TRADING_WINDOW                   = newErrorCode(-2016, []string{"No trading window could be found for the symbol. Try ticker/24hrs instead."})
	ORDER_ARCHIVED                      = newErrorCode(-2026, []string{"Order was canceled or expired with no executed qty over 90 days ago and has been archived."})
	SUBSCRIPTION_ACTIVE                 = newErrorCode(-2035, []string{"User Data Stream subscription already active."})
	SUBSCRIPTION_INACTIVE               = newErrorCode(-2036, []string{"User Data Stream subscription not active."})
)

// //
// Helper function
// //
func newErrorCode(code int, descriptions []string) berror.ErrorCode {
	return berror.ErrorCode{
		Code:         code,
		Descriptions: descriptions,
	}
}
