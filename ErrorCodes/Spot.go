package errorcodes

var Spot = struct {
	Names SpotNames
	Codes SpotCodes
}{
	Names: SpotNames{
		UNKNOWN:                             &SpotErr_UNKNOWN,
		DISCONNECTED:                        &SpotErr_DISCONNECTED,
		UNAUTHORIZED:                        &SpotErr_UNAUTHORIZED,
		TOO_MANY_REQUESTS:                   &SpotErr_TOO_MANY_REQUESTS,
		UNEXPECTED_RESP:                     &SpotErr_UNEXPECTED_RESP,
		TIMEOUT:                             &SpotErr_TIMEOUT,
		SERVER_BUSY:                         &SpotErr_SERVER_BUSY,
		INVALID_MESSAGE:                     &SpotErr_INVALID_MESSAGE,
		UNKNOWN_ORDER_COMPOSITION:           &SpotErr_UNKNOWN_ORDER_COMPOSITION,
		TOO_MANY_ORDERS:                     &SpotErr_TOO_MANY_ORDERS,
		SERVICE_SHUTTING_DOWN:               &SpotErr_SERVICE_SHUTTING_DOWN,
		UNSUPPORTED_OPERATION:               &SpotErr_UNSUPPORTED_OPERATION,
		INVALID_TIMESTAMP:                   &SpotErr_INVALID_TIMESTAMP,
		INVALID_SIGNATURE:                   &SpotErr_INVALID_SIGNATURE,
		COMP_ID_IN_USE:                      &SpotErr_COMP_ID_IN_USE,
		TOO_MANY_CONNECTIONS:                &SpotErr_TOO_MANY_CONNECTIONS,
		LOGGED_OUT:                          &SpotErr_LOGGED_OUT,
		ILLEGAL_CHARS:                       &SpotErr_ILLEGAL_CHARS,
		TOO_MANY_PARAMETERS:                 &SpotErr_TOO_MANY_PARAMETERS,
		MANDATORY_PARAM_EMPTY_OR_MALFORMED:  &SpotErr_MANDATORY_PARAM_EMPTY_OR_MALFORMED,
		UNKNOWN_PARAM:                       &SpotErr_UNKNOWN_PARAM,
		UNREAD_PARAMETERS:                   &SpotErr_UNREAD_PARAMETERS,
		PARAM_EMPTY:                         &SpotErr_PARAM_EMPTY,
		PARAM_NOT_REQUIRED:                  &SpotErr_PARAM_NOT_REQUIRED,
		PARAM_OVERFLOW:                      &SpotErr_PARAM_OVERFLOW,
		BAD_PRECISION:                       &SpotErr_BAD_PRECISION,
		NO_DEPTH:                            &SpotErr_NO_DEPTH,
		TIF_NOT_REQUIRED:                    &SpotErr_TIF_NOT_REQUIRED,
		INVALID_TIF:                         &SpotErr_INVALID_TIF,
		INVALID_ORDER_TYPE:                  &SpotErr_INVALID_ORDER_TYPE,
		INVALID_SIDE:                        &SpotErr_INVALID_SIDE,
		EMPTY_NEW_CL_ORD_ID:                 &SpotErr_EMPTY_NEW_CL_ORD_ID,
		EMPTY_ORG_CL_ORD_ID:                 &SpotErr_EMPTY_ORG_CL_ORD_ID,
		BAD_INTERVAL:                        &SpotErr_BAD_INTERVAL,
		BAD_SYMBOL:                          &SpotErr_BAD_SYMBOL,
		INVALID_SYMBOLSTATUS:                &SpotErr_INVALID_SYMBOLSTATUS,
		INVALID_LISTEN_KEY:                  &SpotErr_INVALID_LISTEN_KEY,
		MORE_THAN_XX_HOURS:                  &SpotErr_MORE_THAN_XX_HOURS,
		OPTIONAL_PARAMS_BAD_COMBO:           &SpotErr_OPTIONAL_PARAMS_BAD_COMBO,
		INVALID_PARAMETER:                   &SpotErr_INVALID_PARAMETER,
		BAD_STRATEGY_TYPE:                   &SpotErr_BAD_STRATEGY_TYPE,
		INVALID_JSON:                        &SpotErr_INVALID_JSON,
		INVALID_TICKER_TYPE:                 &SpotErr_INVALID_TICKER_TYPE,
		INVALID_CANCEL_RESTRICTIONS:         &SpotErr_INVALID_CANCEL_RESTRICTIONS,
		DUPLICATE_SYMBOLS:                   &SpotErr_DUPLICATE_SYMBOLS,
		INVALID_SBE_HEADER:                  &SpotErr_INVALID_SBE_HEADER,
		UNSUPPORTED_SCHEMA_ID:               &SpotErr_UNSUPPORTED_SCHEMA_ID,
		SBE_DISABLED:                        &SpotErr_SBE_DISABLED,
		OCO_ORDER_TYPE_REJECTED:             &SpotErr_OCO_ORDER_TYPE_REJECTED,
		OCO_ICEBERGQTY_TIMEINFORCE:          &SpotErr_OCO_ICEBERGQTY_TIMEINFORCE,
		DEPRECATED_SCHEMA:                   &SpotErr_DEPRECATED_SCHEMA,
		BUY_OCO_LIMIT_MUST_BE_BELOW:         &SpotErr_BUY_OCO_LIMIT_MUST_BE_BELOW,
		SELL_OCO_LIMIT_MUST_BE_ABOVE:        &SpotErr_SELL_OCO_LIMIT_MUST_BE_ABOVE,
		BOTH_OCO_ORDERS_CANNOT_BE_LIMIT:     &SpotErr_BOTH_OCO_ORDERS_CANNOT_BE_LIMIT,
		INVALID_TAG_NUMBER:                  &SpotErr_INVALID_TAG_NUMBER,
		TAG_NOT_DEFINED_IN_MESSAGE:          &SpotErr_TAG_NOT_DEFINED_IN_MESSAGE,
		TAG_APPEARS_MORE_THAN_ONCE:          &SpotErr_TAG_APPEARS_MORE_THAN_ONCE,
		TAG_OUT_OF_ORDER:                    &SpotErr_TAG_OUT_OF_ORDER,
		GROUP_FIELDS_OUT_OF_ORDER:           &SpotErr_GROUP_FIELDS_OUT_OF_ORDER,
		INVALID_COMPONENT:                   &SpotErr_INVALID_COMPONENT,
		RESET_SEQ_NUM_SUPPORT:               &SpotErr_RESET_SEQ_NUM_SUPPORT,
		ALREADY_LOGGED_IN:                   &SpotErr_ALREADY_LOGGED_IN,
		GARBLED_MESSAGE:                     &SpotErr_GARBLED_MESSAGE,
		BAD_SENDER_COMPID:                   &SpotErr_BAD_SENDER_COMPID,
		BAD_SEQ_NUM:                         &SpotErr_BAD_SEQ_NUM,
		EXPECTED_LOGON:                      &SpotErr_EXPECTED_LOGON,
		TOO_MANY_MESSAGES:                   &SpotErr_TOO_MANY_MESSAGES,
		PARAMS_BAD_COMBO:                    &SpotErr_PARAMS_BAD_COMBO,
		NOT_ALLOWED_IN_DROP_COPY_SESSIONS:   &SpotErr_NOT_ALLOWED_IN_DROP_COPY_SESSIONS,
		DROP_COPY_SESSION_NOT_ALLOWED:       &SpotErr_DROP_COPY_SESSION_NOT_ALLOWED,
		DROP_COPY_SESSION_REQUIRED:          &SpotErr_DROP_COPY_SESSION_REQUIRED,
		NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS: &SpotErr_NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS,
		NOT_ALLOWED_IN_MARKET_DATA_SESSIONS: &SpotErr_NOT_ALLOWED_IN_MARKET_DATA_SESSIONS,
		INCORRECT_NUM_IN_GROUP_COUNT:        &SpotErr_INCORRECT_NUM_IN_GROUP_COUNT,
		DUPLICATE_ENTRIES_IN_A_GROUP:        &SpotErr_DUPLICATE_ENTRIES_IN_A_GROUP,
		INVALID_REQUEST_ID:                  &SpotErr_INVALID_REQUEST_ID,
		TOO_MANY_SUBSCRIPTIONS:              &SpotErr_TOO_MANY_SUBSCRIPTIONS,
		INVALID_TIME_UNIT:                   &SpotErr_INVALID_TIME_UNIT,
		BUY_OCO_STOP_LOSS_MUST_BE_ABOVE:     &SpotErr_BUY_OCO_STOP_LOSS_MUST_BE_ABOVE,
		SELL_OCO_STOP_LOSS_MUST_BE_BELOW:    &SpotErr_SELL_OCO_STOP_LOSS_MUST_BE_BELOW,
		BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW:   &SpotErr_BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW,
		SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE:  &SpotErr_SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE,
		NEW_ORDER_REJECTED:                  &SpotErr_NEW_ORDER_REJECTED,
		CANCEL_REJECTED:                     &SpotErr_CANCEL_REJECTED,
		NO_SUCH_ORDER:                       &SpotErr_NO_SUCH_ORDER,
		BAD_API_KEY_FMT:                     &SpotErr_BAD_API_KEY_FMT,
		REJECTED_MBX_KEY:                    &SpotErr_REJECTED_MBX_KEY,
		NO_TRADING_WINDOW:                   &SpotErr_NO_TRADING_WINDOW,
		ORDER_ARCHIVED:                      &SpotErr_ORDER_ARCHIVED,
		SUBSCRIPTION_ACTIVE:                 &SpotErr_SUBSCRIPTION_ACTIVE,
		SUBSCRIPTION_INACTIVE:               &SpotErr_SUBSCRIPTION_INACTIVE,
	},
	Codes: SpotCodes{
		N1000: &SpotErr_UNKNOWN,
		N1001: &SpotErr_DISCONNECTED,
		N1002: &SpotErr_UNAUTHORIZED,
		N1003: &SpotErr_TOO_MANY_REQUESTS,
		N1006: &SpotErr_UNEXPECTED_RESP,
		N1007: &SpotErr_TIMEOUT,
		N1008: &SpotErr_SERVER_BUSY,
		N1013: &SpotErr_INVALID_MESSAGE,
		N1014: &SpotErr_UNKNOWN_ORDER_COMPOSITION,
		N1015: &SpotErr_TOO_MANY_ORDERS,
		N1016: &SpotErr_SERVICE_SHUTTING_DOWN,
		N1020: &SpotErr_UNSUPPORTED_OPERATION,
		N1021: &SpotErr_INVALID_TIMESTAMP,
		N1022: &SpotErr_INVALID_SIGNATURE,
		N1033: &SpotErr_COMP_ID_IN_USE,
		N1034: &SpotErr_TOO_MANY_CONNECTIONS,
		N1035: &SpotErr_LOGGED_OUT,
		N1100: &SpotErr_ILLEGAL_CHARS,
		N1101: &SpotErr_TOO_MANY_PARAMETERS,
		N1102: &SpotErr_MANDATORY_PARAM_EMPTY_OR_MALFORMED,
		N1103: &SpotErr_UNKNOWN_PARAM,
		N1104: &SpotErr_UNREAD_PARAMETERS,
		N1105: &SpotErr_PARAM_EMPTY,
		N1106: &SpotErr_PARAM_NOT_REQUIRED,
		N1108: &SpotErr_PARAM_OVERFLOW,
		N1111: &SpotErr_BAD_PRECISION,
		N1112: &SpotErr_NO_DEPTH,
		N1114: &SpotErr_TIF_NOT_REQUIRED,
		N1115: &SpotErr_INVALID_TIF,
		N1116: &SpotErr_INVALID_ORDER_TYPE,
		N1117: &SpotErr_INVALID_SIDE,
		N1118: &SpotErr_EMPTY_NEW_CL_ORD_ID,
		N1119: &SpotErr_EMPTY_ORG_CL_ORD_ID,
		N1120: &SpotErr_BAD_INTERVAL,
		N1121: &SpotErr_BAD_SYMBOL,
		N1122: &SpotErr_INVALID_SYMBOLSTATUS,
		N1125: &SpotErr_INVALID_LISTEN_KEY,
		N1127: &SpotErr_MORE_THAN_XX_HOURS,
		N1128: &SpotErr_OPTIONAL_PARAMS_BAD_COMBO,
		N1130: &SpotErr_INVALID_PARAMETER,
		N1134: &SpotErr_BAD_STRATEGY_TYPE,
		N1135: &SpotErr_INVALID_JSON,
		N1139: &SpotErr_INVALID_TICKER_TYPE,
		N1145: &SpotErr_INVALID_CANCEL_RESTRICTIONS,
		N1151: &SpotErr_DUPLICATE_SYMBOLS,
		N1152: &SpotErr_INVALID_SBE_HEADER,
		N1153: &SpotErr_UNSUPPORTED_SCHEMA_ID,
		N1155: &SpotErr_SBE_DISABLED,
		N1158: &SpotErr_OCO_ORDER_TYPE_REJECTED,
		N1160: &SpotErr_OCO_ICEBERGQTY_TIMEINFORCE,
		N1161: &SpotErr_DEPRECATED_SCHEMA,
		N1165: &SpotErr_BUY_OCO_LIMIT_MUST_BE_BELOW,
		N1166: &SpotErr_SELL_OCO_LIMIT_MUST_BE_ABOVE,
		N1168: &SpotErr_BOTH_OCO_ORDERS_CANNOT_BE_LIMIT,
		N1169: &SpotErr_INVALID_TAG_NUMBER,
		N1170: &SpotErr_TAG_NOT_DEFINED_IN_MESSAGE,
		N1171: &SpotErr_TAG_APPEARS_MORE_THAN_ONCE,
		N1172: &SpotErr_TAG_OUT_OF_ORDER,
		N1173: &SpotErr_GROUP_FIELDS_OUT_OF_ORDER,
		N1174: &SpotErr_INVALID_COMPONENT,
		N1175: &SpotErr_RESET_SEQ_NUM_SUPPORT,
		N1176: &SpotErr_ALREADY_LOGGED_IN,
		N1177: &SpotErr_GARBLED_MESSAGE,
		N1178: &SpotErr_BAD_SENDER_COMPID,
		N1179: &SpotErr_BAD_SEQ_NUM,
		N1180: &SpotErr_EXPECTED_LOGON,
		N1181: &SpotErr_TOO_MANY_MESSAGES,
		N1182: &SpotErr_PARAMS_BAD_COMBO,
		N1183: &SpotErr_NOT_ALLOWED_IN_DROP_COPY_SESSIONS,
		N1184: &SpotErr_DROP_COPY_SESSION_NOT_ALLOWED,
		N1185: &SpotErr_DROP_COPY_SESSION_REQUIRED,
		N1186: &SpotErr_NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS,
		N1187: &SpotErr_NOT_ALLOWED_IN_MARKET_DATA_SESSIONS,
		N1188: &SpotErr_INCORRECT_NUM_IN_GROUP_COUNT,
		N1189: &SpotErr_DUPLICATE_ENTRIES_IN_A_GROUP,
		N1190: &SpotErr_INVALID_REQUEST_ID,
		N1191: &SpotErr_TOO_MANY_SUBSCRIPTIONS,
		N1194: &SpotErr_INVALID_TIME_UNIT,
		N1196: &SpotErr_BUY_OCO_STOP_LOSS_MUST_BE_ABOVE,
		N1197: &SpotErr_SELL_OCO_STOP_LOSS_MUST_BE_BELOW,
		N1198: &SpotErr_BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW,
		N1199: &SpotErr_SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE,
		N2010: &SpotErr_NEW_ORDER_REJECTED,
		N2011: &SpotErr_CANCEL_REJECTED,
		N2013: &SpotErr_NO_SUCH_ORDER,
		N2014: &SpotErr_BAD_API_KEY_FMT,
		N2015: &SpotErr_REJECTED_MBX_KEY,
		N2016: &SpotErr_NO_TRADING_WINDOW,
		N2026: &SpotErr_ORDER_ARCHIVED,
		N2035: &SpotErr_SUBSCRIPTION_ACTIVE,
		N2036: &SpotErr_SUBSCRIPTION_INACTIVE,
	},
}

type SpotNames struct {
	UNKNOWN                             *BinanceErrorCode
	DISCONNECTED                        *BinanceErrorCode
	UNAUTHORIZED                        *BinanceErrorCode
	TOO_MANY_REQUESTS                   *BinanceErrorCode
	UNEXPECTED_RESP                     *BinanceErrorCode
	TIMEOUT                             *BinanceErrorCode
	SERVER_BUSY                         *BinanceErrorCode
	INVALID_MESSAGE                     *BinanceErrorCode
	UNKNOWN_ORDER_COMPOSITION           *BinanceErrorCode
	TOO_MANY_ORDERS                     *BinanceErrorCode
	SERVICE_SHUTTING_DOWN               *BinanceErrorCode
	UNSUPPORTED_OPERATION               *BinanceErrorCode
	INVALID_TIMESTAMP                   *BinanceErrorCode
	INVALID_SIGNATURE                   *BinanceErrorCode
	COMP_ID_IN_USE                      *BinanceErrorCode
	TOO_MANY_CONNECTIONS                *BinanceErrorCode
	LOGGED_OUT                          *BinanceErrorCode
	ILLEGAL_CHARS                       *BinanceErrorCode
	TOO_MANY_PARAMETERS                 *BinanceErrorCode
	MANDATORY_PARAM_EMPTY_OR_MALFORMED  *BinanceErrorCode
	UNKNOWN_PARAM                       *BinanceErrorCode
	UNREAD_PARAMETERS                   *BinanceErrorCode
	PARAM_EMPTY                         *BinanceErrorCode
	PARAM_NOT_REQUIRED                  *BinanceErrorCode
	PARAM_OVERFLOW                      *BinanceErrorCode
	BAD_PRECISION                       *BinanceErrorCode
	NO_DEPTH                            *BinanceErrorCode
	TIF_NOT_REQUIRED                    *BinanceErrorCode
	INVALID_TIF                         *BinanceErrorCode
	INVALID_ORDER_TYPE                  *BinanceErrorCode
	INVALID_SIDE                        *BinanceErrorCode
	EMPTY_NEW_CL_ORD_ID                 *BinanceErrorCode
	EMPTY_ORG_CL_ORD_ID                 *BinanceErrorCode
	BAD_INTERVAL                        *BinanceErrorCode
	BAD_SYMBOL                          *BinanceErrorCode
	INVALID_SYMBOLSTATUS                *BinanceErrorCode
	INVALID_LISTEN_KEY                  *BinanceErrorCode
	MORE_THAN_XX_HOURS                  *BinanceErrorCode
	OPTIONAL_PARAMS_BAD_COMBO           *BinanceErrorCode
	INVALID_PARAMETER                   *BinanceErrorCode
	BAD_STRATEGY_TYPE                   *BinanceErrorCode
	INVALID_JSON                        *BinanceErrorCode
	INVALID_TICKER_TYPE                 *BinanceErrorCode
	INVALID_CANCEL_RESTRICTIONS         *BinanceErrorCode
	DUPLICATE_SYMBOLS                   *BinanceErrorCode
	INVALID_SBE_HEADER                  *BinanceErrorCode
	UNSUPPORTED_SCHEMA_ID               *BinanceErrorCode
	SBE_DISABLED                        *BinanceErrorCode
	OCO_ORDER_TYPE_REJECTED             *BinanceErrorCode
	OCO_ICEBERGQTY_TIMEINFORCE          *BinanceErrorCode
	DEPRECATED_SCHEMA                   *BinanceErrorCode
	BUY_OCO_LIMIT_MUST_BE_BELOW         *BinanceErrorCode
	SELL_OCO_LIMIT_MUST_BE_ABOVE        *BinanceErrorCode
	BOTH_OCO_ORDERS_CANNOT_BE_LIMIT     *BinanceErrorCode
	INVALID_TAG_NUMBER                  *BinanceErrorCode
	TAG_NOT_DEFINED_IN_MESSAGE          *BinanceErrorCode
	TAG_APPEARS_MORE_THAN_ONCE          *BinanceErrorCode
	TAG_OUT_OF_ORDER                    *BinanceErrorCode
	GROUP_FIELDS_OUT_OF_ORDER           *BinanceErrorCode
	INVALID_COMPONENT                   *BinanceErrorCode
	RESET_SEQ_NUM_SUPPORT               *BinanceErrorCode
	ALREADY_LOGGED_IN                   *BinanceErrorCode
	GARBLED_MESSAGE                     *BinanceErrorCode
	BAD_SENDER_COMPID                   *BinanceErrorCode
	BAD_SEQ_NUM                         *BinanceErrorCode
	EXPECTED_LOGON                      *BinanceErrorCode
	TOO_MANY_MESSAGES                   *BinanceErrorCode
	PARAMS_BAD_COMBO                    *BinanceErrorCode
	NOT_ALLOWED_IN_DROP_COPY_SESSIONS   *BinanceErrorCode
	DROP_COPY_SESSION_NOT_ALLOWED       *BinanceErrorCode
	DROP_COPY_SESSION_REQUIRED          *BinanceErrorCode
	NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS *BinanceErrorCode
	NOT_ALLOWED_IN_MARKET_DATA_SESSIONS *BinanceErrorCode
	INCORRECT_NUM_IN_GROUP_COUNT        *BinanceErrorCode
	DUPLICATE_ENTRIES_IN_A_GROUP        *BinanceErrorCode
	INVALID_REQUEST_ID                  *BinanceErrorCode
	TOO_MANY_SUBSCRIPTIONS              *BinanceErrorCode
	INVALID_TIME_UNIT                   *BinanceErrorCode
	BUY_OCO_STOP_LOSS_MUST_BE_ABOVE     *BinanceErrorCode
	SELL_OCO_STOP_LOSS_MUST_BE_BELOW    *BinanceErrorCode
	BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW   *BinanceErrorCode
	SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE  *BinanceErrorCode
	NEW_ORDER_REJECTED                  *BinanceErrorCode
	CANCEL_REJECTED                     *BinanceErrorCode
	NO_SUCH_ORDER                       *BinanceErrorCode
	BAD_API_KEY_FMT                     *BinanceErrorCode
	REJECTED_MBX_KEY                    *BinanceErrorCode
	NO_TRADING_WINDOW                   *BinanceErrorCode
	ORDER_ARCHIVED                      *BinanceErrorCode
	SUBSCRIPTION_ACTIVE                 *BinanceErrorCode
	SUBSCRIPTION_INACTIVE               *BinanceErrorCode
}
type SpotCodes struct {
	N1000 *BinanceErrorCode
	N1001 *BinanceErrorCode
	N1002 *BinanceErrorCode
	N1003 *BinanceErrorCode
	N1006 *BinanceErrorCode
	N1007 *BinanceErrorCode
	N1008 *BinanceErrorCode
	N1013 *BinanceErrorCode
	N1014 *BinanceErrorCode
	N1015 *BinanceErrorCode
	N1016 *BinanceErrorCode
	N1020 *BinanceErrorCode
	N1021 *BinanceErrorCode
	N1022 *BinanceErrorCode
	N1033 *BinanceErrorCode
	N1034 *BinanceErrorCode
	N1035 *BinanceErrorCode
	N1100 *BinanceErrorCode
	N1101 *BinanceErrorCode
	N1102 *BinanceErrorCode
	N1103 *BinanceErrorCode
	N1104 *BinanceErrorCode
	N1105 *BinanceErrorCode
	N1106 *BinanceErrorCode
	N1108 *BinanceErrorCode
	N1111 *BinanceErrorCode
	N1112 *BinanceErrorCode
	N1114 *BinanceErrorCode
	N1115 *BinanceErrorCode
	N1116 *BinanceErrorCode
	N1117 *BinanceErrorCode
	N1118 *BinanceErrorCode
	N1119 *BinanceErrorCode
	N1120 *BinanceErrorCode
	N1121 *BinanceErrorCode
	N1122 *BinanceErrorCode
	N1125 *BinanceErrorCode
	N1127 *BinanceErrorCode
	N1128 *BinanceErrorCode
	N1130 *BinanceErrorCode
	N1134 *BinanceErrorCode
	N1135 *BinanceErrorCode
	N1139 *BinanceErrorCode
	N1145 *BinanceErrorCode
	N1151 *BinanceErrorCode
	N1152 *BinanceErrorCode
	N1153 *BinanceErrorCode
	N1155 *BinanceErrorCode
	N1158 *BinanceErrorCode
	N1160 *BinanceErrorCode
	N1161 *BinanceErrorCode
	N1165 *BinanceErrorCode
	N1166 *BinanceErrorCode
	N1168 *BinanceErrorCode
	N1169 *BinanceErrorCode
	N1170 *BinanceErrorCode
	N1171 *BinanceErrorCode
	N1172 *BinanceErrorCode
	N1173 *BinanceErrorCode
	N1174 *BinanceErrorCode
	N1175 *BinanceErrorCode
	N1176 *BinanceErrorCode
	N1177 *BinanceErrorCode
	N1178 *BinanceErrorCode
	N1179 *BinanceErrorCode
	N1180 *BinanceErrorCode
	N1181 *BinanceErrorCode
	N1182 *BinanceErrorCode
	N1183 *BinanceErrorCode
	N1184 *BinanceErrorCode
	N1185 *BinanceErrorCode
	N1186 *BinanceErrorCode
	N1187 *BinanceErrorCode
	N1188 *BinanceErrorCode
	N1189 *BinanceErrorCode
	N1190 *BinanceErrorCode
	N1191 *BinanceErrorCode
	N1194 *BinanceErrorCode
	N1196 *BinanceErrorCode
	N1197 *BinanceErrorCode
	N1198 *BinanceErrorCode
	N1199 *BinanceErrorCode
	N2010 *BinanceErrorCode
	N2011 *BinanceErrorCode
	N2013 *BinanceErrorCode
	N2014 *BinanceErrorCode
	N2015 *BinanceErrorCode
	N2016 *BinanceErrorCode
	N2026 *BinanceErrorCode
	N2035 *BinanceErrorCode
	N2036 *BinanceErrorCode
}

var (
	SpotErr_UNKNOWN                             = BinanceErrorCode{Code: -1000, Name: "UNKNOWN", Descriptions: []string{"An unknown error occurred while processing the request."}}
	SpotErr_DISCONNECTED                        = BinanceErrorCode{Code: -1001, Name: "DISCONNECTED", Descriptions: []string{"Internal error; unable to process your request. Please try again."}}
	SpotErr_UNAUTHORIZED                        = BinanceErrorCode{Code: -1002, Name: "UNAUTHORIZED", Descriptions: []string{"You are not authorized to execute this request."}}
	SpotErr_TOO_MANY_REQUESTS                   = BinanceErrorCode{Code: -1003, Name: "TOO_MANY_REQUESTS", Descriptions: []string{"Too many requests queued.", "Too much request weight used; current limit is %s request weight per %s. Please use WebSocket Streams for live updates to avoid polling the API.", "Way too much request weight used; IP banned until %s. Please use WebSocket Streams for live updates to avoid bans."}}
	SpotErr_UNEXPECTED_RESP                     = BinanceErrorCode{Code: -1006, Name: "UNEXPECTED_RESP", Descriptions: []string{"An unexpected response was received from the message bus. Execution status unknown."}}
	SpotErr_TIMEOUT                             = BinanceErrorCode{Code: -1007, Name: "TIMEOUT", Descriptions: []string{"Timeout waiting for response from backend server. Send status unknown; execution status unknown."}}
	SpotErr_SERVER_BUSY                         = BinanceErrorCode{Code: -1008, Name: "SERVER_BUSY", Descriptions: []string{"Server is currently overloaded with other requests. Please try again in a few minutes."}}
	SpotErr_INVALID_MESSAGE                     = BinanceErrorCode{Code: -1013, Name: "INVALID_MESSAGE", Descriptions: []string{"The request is rejected by the API. (i.e. The request didn't reach the Matching Engine.)", "Potential error messages can be found in Filter Failures or Failures during order placement."}}
	SpotErr_UNKNOWN_ORDER_COMPOSITION           = BinanceErrorCode{Code: -1014, Name: "UNKNOWN_ORDER_COMPOSITION", Descriptions: []string{"Unsupported order combination."}}
	SpotErr_TOO_MANY_ORDERS                     = BinanceErrorCode{Code: -1015, Name: "TOO_MANY_ORDERS", Descriptions: []string{"Too many new orders.", "Too many new orders; current limit is %s orders per %s."}}
	SpotErr_SERVICE_SHUTTING_DOWN               = BinanceErrorCode{Code: -1016, Name: "SERVICE_SHUTTING_DOWN", Descriptions: []string{"This service is no longer available."}}
	SpotErr_UNSUPPORTED_OPERATION               = BinanceErrorCode{Code: -1020, Name: "UNSUPPORTED_OPERATION", Descriptions: []string{"This operation is not supported."}}
	SpotErr_INVALID_TIMESTAMP                   = BinanceErrorCode{Code: -1021, Name: "INVALID_TIMESTAMP", Descriptions: []string{"Timestamp for this request is outside of the recvWindow.", "Timestamp for this request was 1000ms ahead of the server's time."}}
	SpotErr_INVALID_SIGNATURE                   = BinanceErrorCode{Code: -1022, Name: "INVALID_SIGNATURE", Descriptions: []string{"Signature for this request is not valid."}}
	SpotErr_COMP_ID_IN_USE                      = BinanceErrorCode{Code: -1033, Name: "COMP_ID_IN_USE", Descriptions: []string{"SenderCompId(49) is currently in use. Concurrent use of the same SenderCompId within one account is not allowed."}}
	SpotErr_TOO_MANY_CONNECTIONS                = BinanceErrorCode{Code: -1034, Name: "TOO_MANY_CONNECTIONS", Descriptions: []string{"Too many concurrent connections; current limit is '%s'.", "Too many connection attempts for account; current limit is %s per '%s'.", "Too many connection attempts from IP; current limit is %s per '%s'."}}
	SpotErr_LOGGED_OUT                          = BinanceErrorCode{Code: -1035, Name: "LOGGED_OUT", Descriptions: []string{"Please send Logout<5> message to close the session."}}
	SpotErr_ILLEGAL_CHARS                       = BinanceErrorCode{Code: -1100, Name: "ILLEGAL_CHARS", Descriptions: []string{"Illegal characters found in a parameter.", "Illegal characters found in parameter '%s'; legal range is '%s'."}}
	SpotErr_TOO_MANY_PARAMETERS                 = BinanceErrorCode{Code: -1101, Name: "TOO_MANY_PARAMETERS", Descriptions: []string{"Too many parameters sent for this endpoint.", "Too many parameters; expected '%s' and received '%s'.", "Duplicate values for a parameter detected."}}
	SpotErr_MANDATORY_PARAM_EMPTY_OR_MALFORMED  = BinanceErrorCode{Code: -1102, Name: "MANDATORY_PARAM_EMPTY_OR_MALFORMED", Descriptions: []string{"A mandatory parameter was not sent, was empty/null, or malformed.", "Mandatory parameter '%s' was not sent, was empty/null, or malformed.", "Param '%s' or '%s' must be sent, but both were empty/null!", "Required tag '%s' missing.", "Field value was empty or malformed.", "'%s' contains unexpected value. Cannot be greater than %s."}}
	SpotErr_UNKNOWN_PARAM                       = BinanceErrorCode{Code: -1103, Name: "UNKNOWN_PARAM", Descriptions: []string{"An unknown parameter was sent.", "Undefined Tag."}}
	SpotErr_UNREAD_PARAMETERS                   = BinanceErrorCode{Code: -1104, Name: "UNREAD_PARAMETERS", Descriptions: []string{"Not all sent parameters were read.", "Not all sent parameters were read; read '%s' parameter(s) but was sent '%s'."}}
	SpotErr_PARAM_EMPTY                         = BinanceErrorCode{Code: -1105, Name: "PARAM_EMPTY", Descriptions: []string{"A parameter was empty.", "Parameter '%s' was empty."}}
	SpotErr_PARAM_NOT_REQUIRED                  = BinanceErrorCode{Code: -1106, Name: "PARAM_NOT_REQUIRED", Descriptions: []string{"A parameter was sent when not required.", "Parameter '%s' sent when not required.", "A tag '%s' was sent when not required."}}
	SpotErr_PARAM_OVERFLOW                      = BinanceErrorCode{Code: -1108, Name: "PARAM_OVERFLOW", Descriptions: []string{"Parameter '%s' overflowed."}}
	SpotErr_BAD_PRECISION                       = BinanceErrorCode{Code: -1111, Name: "BAD_PRECISION", Descriptions: []string{"Parameter '%s' has too much precision."}}
	SpotErr_NO_DEPTH                            = BinanceErrorCode{Code: -1112, Name: "NO_DEPTH", Descriptions: []string{"No orders on book for symbol."}}
	SpotErr_TIF_NOT_REQUIRED                    = BinanceErrorCode{Code: -1114, Name: "TIF_NOT_REQUIRED", Descriptions: []string{"TimeInForce parameter sent when not required."}}
	SpotErr_INVALID_TIF                         = BinanceErrorCode{Code: -1115, Name: "INVALID_TIF", Descriptions: []string{"Invalid timeInForce."}}
	SpotErr_INVALID_ORDER_TYPE                  = BinanceErrorCode{Code: -1116, Name: "INVALID_ORDER_TYPE", Descriptions: []string{"Invalid orderType."}}
	SpotErr_INVALID_SIDE                        = BinanceErrorCode{Code: -1117, Name: "INVALID_SIDE", Descriptions: []string{"Invalid side."}}
	SpotErr_EMPTY_NEW_CL_ORD_ID                 = BinanceErrorCode{Code: -1118, Name: "EMPTY_NEW_CL_ORD_ID", Descriptions: []string{"New client order ID was empty."}}
	SpotErr_EMPTY_ORG_CL_ORD_ID                 = BinanceErrorCode{Code: -1119, Name: "EMPTY_ORG_CL_ORD_ID", Descriptions: []string{"Original client order ID was empty."}}
	SpotErr_BAD_INTERVAL                        = BinanceErrorCode{Code: -1120, Name: "BAD_INTERVAL", Descriptions: []string{"Invalid interval."}}
	SpotErr_BAD_SYMBOL                          = BinanceErrorCode{Code: -1121, Name: "BAD_SYMBOL", Descriptions: []string{"Invalid symbol."}}
	SpotErr_INVALID_SYMBOLSTATUS                = BinanceErrorCode{Code: -1122, Name: "INVALID_SYMBOLSTATUS", Descriptions: []string{"Invalid symbolStatus."}}
	SpotErr_INVALID_LISTEN_KEY                  = BinanceErrorCode{Code: -1125, Name: "INVALID_LISTEN_KEY", Descriptions: []string{"This listenKey does not exist."}}
	SpotErr_MORE_THAN_XX_HOURS                  = BinanceErrorCode{Code: -1127, Name: "MORE_THAN_XX_HOURS", Descriptions: []string{"Lookup interval is too big.", "More than %s hours between startTime and endTime."}}
	SpotErr_OPTIONAL_PARAMS_BAD_COMBO           = BinanceErrorCode{Code: -1128, Name: "OPTIONAL_PARAMS_BAD_COMBO", Descriptions: []string{"Combination of optional parameters invalid.", "Combination of optional fields invalid. Recommendation: '%s' and '%s' must both be sent.", "Fields [%s] must be sent together or omitted entirely.", "Invalid MDEntryType (269) combination. BID and OFFER must be requested together."}}
	SpotErr_INVALID_PARAMETER                   = BinanceErrorCode{Code: -1130, Name: "INVALID_PARAMETER", Descriptions: []string{"Invalid data sent for a parameter.", "Data sent for parameter '%s' is not valid."}}
	SpotErr_BAD_STRATEGY_TYPE                   = BinanceErrorCode{Code: -1134, Name: "BAD_STRATEGY_TYPE", Descriptions: []string{"strategyType was less than 1000000.", "TargetStrategy (847) was less than 1000000."}}
	SpotErr_INVALID_JSON                        = BinanceErrorCode{Code: -1135, Name: "INVALID_JSON", Descriptions: []string{"Invalid JSON Request", "JSON sent for parameter '%s' is not valid"}}
	SpotErr_INVALID_TICKER_TYPE                 = BinanceErrorCode{Code: -1139, Name: "INVALID_TICKER_TYPE", Descriptions: []string{"Invalid ticker type."}}
	SpotErr_INVALID_CANCEL_RESTRICTIONS         = BinanceErrorCode{Code: -1145, Name: "INVALID_CANCEL_RESTRICTIONS", Descriptions: []string{"cancelRestrictions has to be either ONLY_NEW or ONLY_PARTIALLY_FILLED."}}
	SpotErr_DUPLICATE_SYMBOLS                   = BinanceErrorCode{Code: -1151, Name: "DUPLICATE_SYMBOLS", Descriptions: []string{"Symbol is present multiple times in the list."}}
	SpotErr_INVALID_SBE_HEADER                  = BinanceErrorCode{Code: -1152, Name: "INVALID_SBE_HEADER", Descriptions: []string{"Invalid X-MBX-SBE header; expected <SCHEMA_ID>:<VERSION>."}}
	SpotErr_UNSUPPORTED_SCHEMA_ID               = BinanceErrorCode{Code: -1153, Name: "UNSUPPORTED_SCHEMA_ID", Descriptions: []string{"Unsupported SBE schema ID or version specified in the X-MBX-SBE header."}}
	SpotErr_SBE_DISABLED                        = BinanceErrorCode{Code: -1155, Name: "SBE_DISABLED", Descriptions: []string{"SBE is not enabled."}}
	SpotErr_OCO_ORDER_TYPE_REJECTED             = BinanceErrorCode{Code: -1158, Name: "OCO_ORDER_TYPE_REJECTED", Descriptions: []string{"Order type not supported in OCO.", "If the order type provided in the aboveType and/or belowType is not supported."}}
	SpotErr_OCO_ICEBERGQTY_TIMEINFORCE          = BinanceErrorCode{Code: -1160, Name: "OCO_ICEBERGQTY_TIMEINFORCE", Descriptions: []string{"Parameter '%s' is not supported if aboveTimeInForce/belowTimeInForce is not GTC.", "If the order type for the above or below leg is STOP_LOSS_LIMIT, and icebergQty is provided for that leg, the timeInForce has to be GTC else it will throw an error.", "TimeInForce (59) must be GTC (1) when MaxFloor (111) is used."}}
	SpotErr_DEPRECATED_SCHEMA                   = BinanceErrorCode{Code: -1161, Name: "DEPRECATED_SCHEMA", Descriptions: []string{"Unable to encode the response in SBE schema 'x'. Please use schema 'y' or higher."}}
	SpotErr_BUY_OCO_LIMIT_MUST_BE_BELOW         = BinanceErrorCode{Code: -1165, Name: "BUY_OCO_LIMIT_MUST_BE_BELOW", Descriptions: []string{"A limit order in a buy OCO must be below."}}
	SpotErr_SELL_OCO_LIMIT_MUST_BE_ABOVE        = BinanceErrorCode{Code: -1166, Name: "SELL_OCO_LIMIT_MUST_BE_ABOVE", Descriptions: []string{"A limit order in a sell OCO must be above."}}
	SpotErr_BOTH_OCO_ORDERS_CANNOT_BE_LIMIT     = BinanceErrorCode{Code: -1168, Name: "BOTH_OCO_ORDERS_CANNOT_BE_LIMIT", Descriptions: []string{"At least one OCO order must be contingent."}}
	SpotErr_INVALID_TAG_NUMBER                  = BinanceErrorCode{Code: -1169, Name: "INVALID_TAG_NUMBER", Descriptions: []string{"Invalid tag number."}}
	SpotErr_TAG_NOT_DEFINED_IN_MESSAGE          = BinanceErrorCode{Code: -1170, Name: "TAG_NOT_DEFINED_IN_MESSAGE", Descriptions: []string{"Tag '%s' not defined for this message type."}}
	SpotErr_TAG_APPEARS_MORE_THAN_ONCE          = BinanceErrorCode{Code: -1171, Name: "TAG_APPEARS_MORE_THAN_ONCE", Descriptions: []string{"Tag '%s' appears more than once."}}
	SpotErr_TAG_OUT_OF_ORDER                    = BinanceErrorCode{Code: -1172, Name: "TAG_OUT_OF_ORDER", Descriptions: []string{"Tag '%s' specified out of required order."}}
	SpotErr_GROUP_FIELDS_OUT_OF_ORDER           = BinanceErrorCode{Code: -1173, Name: "GROUP_FIELDS_OUT_OF_ORDER", Descriptions: []string{"Repeating group '%s' fields out of order."}}
	SpotErr_INVALID_COMPONENT                   = BinanceErrorCode{Code: -1174, Name: "INVALID_COMPONENT", Descriptions: []string{"Component '%s' is incorrectly populated on '%s' order. Recommendation: '%s'"}}
	SpotErr_RESET_SEQ_NUM_SUPPORT               = BinanceErrorCode{Code: -1175, Name: "RESET_SEQ_NUM_SUPPORT", Descriptions: []string{"Continuation of sequence numbers to new session is currently unsupported. Sequence numbers must be reset for each new session."}}
	SpotErr_ALREADY_LOGGED_IN                   = BinanceErrorCode{Code: -1176, Name: "ALREADY_LOGGED_IN", Descriptions: []string{"Logon<A> should only be sent once."}}
	SpotErr_GARBLED_MESSAGE                     = BinanceErrorCode{Code: -1177, Name: "GARBLED_MESSAGE", Descriptions: []string{"CheckSum(10) contains an incorrect value.", "BeginString (8) is not the first tag in a message.", "MsgType (35) is not the third tag in a message.", "BodyLength (9) does not contain the correct byte count.", "Only printable ASCII characters and SOH (Start of Header) are allowed."}}
	SpotErr_BAD_SENDER_COMPID                   = BinanceErrorCode{Code: -1178, Name: "BAD_SENDER_COMPID", Descriptions: []string{"SenderCompId(49) contains an incorrect value. The SenderCompID value should not change throughout the lifetime of a session."}}
	SpotErr_BAD_SEQ_NUM                         = BinanceErrorCode{Code: -1179, Name: "BAD_SEQ_NUM", Descriptions: []string{"MsgSeqNum(34) contains an unexpected value. Expected: '%d'."}}
	SpotErr_EXPECTED_LOGON                      = BinanceErrorCode{Code: -1180, Name: "EXPECTED_LOGON", Descriptions: []string{"Logon<A> must be the first message in the session."}}
	SpotErr_TOO_MANY_MESSAGES                   = BinanceErrorCode{Code: -1181, Name: "TOO_MANY_MESSAGES", Descriptions: []string{"Too many messages; current limit is '%d' messages per '%s'."}}
	SpotErr_PARAMS_BAD_COMBO                    = BinanceErrorCode{Code: -1182, Name: "PARAMS_BAD_COMBO", Descriptions: []string{"Conflicting fields: [%s]"}}
	SpotErr_NOT_ALLOWED_IN_DROP_COPY_SESSIONS   = BinanceErrorCode{Code: -1183, Name: "NOT_ALLOWED_IN_DROP_COPY_SESSIONS", Descriptions: []string{"Requested operation is not allowed in DropCopy sessions."}}
	SpotErr_DROP_COPY_SESSION_NOT_ALLOWED       = BinanceErrorCode{Code: -1184, Name: "DROP_COPY_SESSION_NOT_ALLOWED", Descriptions: []string{"DropCopy sessions are not supported on this server. Please reconnect to a drop copy server."}}
	SpotErr_DROP_COPY_SESSION_REQUIRED          = BinanceErrorCode{Code: -1185, Name: "DROP_COPY_SESSION_REQUIRED", Descriptions: []string{"Only DropCopy sessions are supported on this server. Either reconnect to order entry server or send DropCopyFlag (9406) field."}}
	SpotErr_NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS = BinanceErrorCode{Code: -1186, Name: "NOT_ALLOWED_IN_ORDER_ENTRY_SESSIONS", Descriptions: []string{"Requested operation is not allowed in order entry sessions."}}
	SpotErr_NOT_ALLOWED_IN_MARKET_DATA_SESSIONS = BinanceErrorCode{Code: -1187, Name: "NOT_ALLOWED_IN_MARKET_DATA_SESSIONS", Descriptions: []string{"Requested operation is not allowed in market data sessions."}}
	SpotErr_INCORRECT_NUM_IN_GROUP_COUNT        = BinanceErrorCode{Code: -1188, Name: "INCORRECT_NUM_IN_GROUP_COUNT", Descriptions: []string{"Incorrect NumInGroup count for repeating group '%s'."}}
	SpotErr_DUPLICATE_ENTRIES_IN_A_GROUP        = BinanceErrorCode{Code: -1189, Name: "DUPLICATE_ENTRIES_IN_A_GROUP", Descriptions: []string{"Group '%s' contains duplicate entries."}}
	SpotErr_INVALID_REQUEST_ID                  = BinanceErrorCode{Code: -1190, Name: "INVALID_REQUEST_ID", Descriptions: []string{"MDReqID (262) contains a subscription request id that is already in use on this connection.", "MDReqID (262) contains an unsubscription request id that does not match any active subscription."}}
	SpotErr_TOO_MANY_SUBSCRIPTIONS              = BinanceErrorCode{Code: -1191, Name: "TOO_MANY_SUBSCRIPTIONS", Descriptions: []string{"Too many subscriptions. Connection may create up to '%s' subscriptions at a time.", "Similar subscription is already active on this connection. Symbol='%s', active subscription id: '%s'."}}
	SpotErr_INVALID_TIME_UNIT                   = BinanceErrorCode{Code: -1194, Name: "INVALID_TIME_UNIT", Descriptions: []string{"Invalid value for time unit; expected either MICROSECOND or MILLISECOND."}}
	SpotErr_BUY_OCO_STOP_LOSS_MUST_BE_ABOVE     = BinanceErrorCode{Code: -1196, Name: "BUY_OCO_STOP_LOSS_MUST_BE_ABOVE", Descriptions: []string{"A stop loss order in a buy OCO must be above."}}
	SpotErr_SELL_OCO_STOP_LOSS_MUST_BE_BELOW    = BinanceErrorCode{Code: -1197, Name: "SELL_OCO_STOP_LOSS_MUST_BE_BELOW", Descriptions: []string{"A stop loss order in a sell OCO must be below."}}
	SpotErr_BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW   = BinanceErrorCode{Code: -1198, Name: "BUY_OCO_TAKE_PROFIT_MUST_BE_BELOW", Descriptions: []string{"A take profit order in a buy OCO must be below."}}
	SpotErr_SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE  = BinanceErrorCode{Code: -1199, Name: "SELL_OCO_TAKE_PROFIT_MUST_BE_ABOVE", Descriptions: []string{"A take profit order in a sell OCO must be above."}}
	SpotErr_NEW_ORDER_REJECTED                  = BinanceErrorCode{Code: -2010, Name: "NEW_ORDER_REJECTED", Descriptions: []string{"NEW_ORDER_REJECTED"}}
	SpotErr_CANCEL_REJECTED                     = BinanceErrorCode{Code: -2011, Name: "CANCEL_REJECTED", Descriptions: []string{"CANCEL_REJECTED"}}
	SpotErr_NO_SUCH_ORDER                       = BinanceErrorCode{Code: -2013, Name: "NO_SUCH_ORDER", Descriptions: []string{"Order does not exist."}}
	SpotErr_BAD_API_KEY_FMT                     = BinanceErrorCode{Code: -2014, Name: "BAD_API_KEY_FMT", Descriptions: []string{"API-key format invalid."}}
	SpotErr_REJECTED_MBX_KEY                    = BinanceErrorCode{Code: -2015, Name: "REJECTED_MBX_KEY", Descriptions: []string{"Invalid API-key, IP, or permissions for action."}}
	SpotErr_NO_TRADING_WINDOW                   = BinanceErrorCode{Code: -2016, Name: "NO_TRADING_WINDOW", Descriptions: []string{"No trading window could be found for the symbol. Try ticker/24hrs instead."}}
	SpotErr_ORDER_ARCHIVED                      = BinanceErrorCode{Code: -2026, Name: "ORDER_ARCHIVED", Descriptions: []string{"Order was canceled or expired with no executed qty over 90 days ago and has been archived."}}
	SpotErr_SUBSCRIPTION_ACTIVE                 = BinanceErrorCode{Code: -2035, Name: "SUBSCRIPTION_ACTIVE", Descriptions: []string{"User Data Stream subscription already active."}}
	SpotErr_SUBSCRIPTION_INACTIVE               = BinanceErrorCode{Code: -2036, Name: "SUBSCRIPTION_INACTIVE", Descriptions: []string{"User Data Stream subscription not active."}}
)
