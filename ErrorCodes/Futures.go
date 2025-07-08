package errorcodes

var Futures = struct {
	Names FuturesNames
	Codes FuturesCodes
}{
	Names: FuturesNames{
		UNKNOWN:                                       &FuturesErr_UNKNOWN,
		DISCONNECTED:                                  &FuturesErr_DISCONNECTED,
		UNAUTHORIZED:                                  &FuturesErr_UNAUTHORIZED,
		TOO_MANY_REQUESTS:                             &FuturesErr_TOO_MANY_REQUESTS,
		DUPLICATE_IP:                                  &FuturesErr_DUPLICATE_IP,
		NO_SUCH_IP:                                    &FuturesErr_NO_SUCH_IP,
		UNEXPECTED_RESP:                               &FuturesErr_UNEXPECTED_RESP,
		TIMEOUT:                                       &FuturesErr_TIMEOUT,
		Service:                                       &FuturesErr_Service,
		ERROR_MSG_RECEIVED:                            &FuturesErr_ERROR_MSG_RECEIVED,
		NON_WHITE_LIST:                                &FuturesErr_NON_WHITE_LIST,
		INVALID_MESSAGE:                               &FuturesErr_INVALID_MESSAGE,
		UNKNOWN_ORDER_COMPOSITION:                     &FuturesErr_UNKNOWN_ORDER_COMPOSITION,
		TOO_MANY_ORDERS:                               &FuturesErr_TOO_MANY_ORDERS,
		SERVICE_SHUTTING_DOWN:                         &FuturesErr_SERVICE_SHUTTING_DOWN,
		UNSUPPORTED_OPERATION:                         &FuturesErr_UNSUPPORTED_OPERATION,
		INVALID_TIMESTAMP:                             &FuturesErr_INVALID_TIMESTAMP,
		INVALID_SIGNATURE:                             &FuturesErr_INVALID_SIGNATURE,
		START_TIME_GREATER_THAN_END_TIME:              &FuturesErr_START_TIME_GREATER_THAN_END_TIME,
		NOT_FOUND:                                     &FuturesErr_NOT_FOUND,
		ILLEGAL_CHARS:                                 &FuturesErr_ILLEGAL_CHARS,
		TOO_MANY_PARAMETERS:                           &FuturesErr_TOO_MANY_PARAMETERS,
		MANDATORY_PARAM_EMPTY_OR_MALFORMED:            &FuturesErr_MANDATORY_PARAM_EMPTY_OR_MALFORMED,
		UNKNOWN_PARAM:                                 &FuturesErr_UNKNOWN_PARAM,
		UNREAD_PARAMETERS:                             &FuturesErr_UNREAD_PARAMETERS,
		PARAM_EMPTY:                                   &FuturesErr_PARAM_EMPTY,
		PARAM_NOT_REQUIRED:                            &FuturesErr_PARAM_NOT_REQUIRED,
		BAD_ASSET:                                     &FuturesErr_BAD_ASSET,
		BAD_ACCOUNT:                                   &FuturesErr_BAD_ACCOUNT,
		BAD_INSTRUMENT_TYPE:                           &FuturesErr_BAD_INSTRUMENT_TYPE,
		BAD_PRECISION:                                 &FuturesErr_BAD_PRECISION,
		NO_DEPTH:                                      &FuturesErr_NO_DEPTH,
		WITHDRAW_NOT_NEGATIVE:                         &FuturesErr_WITHDRAW_NOT_NEGATIVE,
		TIF_NOT_REQUIRED:                              &FuturesErr_TIF_NOT_REQUIRED,
		INVALID_TIF:                                   &FuturesErr_INVALID_TIF,
		INVALID_ORDER_TYPE:                            &FuturesErr_INVALID_ORDER_TYPE,
		INVALID_SIDE:                                  &FuturesErr_INVALID_SIDE,
		EMPTY_NEW_CL_ORD_ID:                           &FuturesErr_EMPTY_NEW_CL_ORD_ID,
		EMPTY_ORG_CL_ORD_ID:                           &FuturesErr_EMPTY_ORG_CL_ORD_ID,
		BAD_INTERVAL:                                  &FuturesErr_BAD_INTERVAL,
		BAD_SYMBOL:                                    &FuturesErr_BAD_SYMBOL,
		INVALID_SYMBOL_STATUS:                         &FuturesErr_INVALID_SYMBOL_STATUS,
		INVALID_LISTEN_KEY:                            &FuturesErr_INVALID_LISTEN_KEY,
		ASSET_NOT_SUPPORTED:                           &FuturesErr_ASSET_NOT_SUPPORTED,
		MORE_THAN_XX_HOURS:                            &FuturesErr_MORE_THAN_XX_HOURS,
		OPTIONAL_PARAMS_BAD_COMBO:                     &FuturesErr_OPTIONAL_PARAMS_BAD_COMBO,
		INVALID_PARAMETER:                             &FuturesErr_INVALID_PARAMETER,
		INVALID_NEW_ORDER_RESP_TYPE:                   &FuturesErr_INVALID_NEW_ORDER_RESP_TYPE,
		NEW_ORDER_REJECTED:                            &FuturesErr_NEW_ORDER_REJECTED,
		CANCEL_REJECTED:                               &FuturesErr_CANCEL_REJECTED,
		CANCEL_ALL_FAIL:                               &FuturesErr_CANCEL_ALL_FAIL,
		NO_SUCH_ORDER:                                 &FuturesErr_NO_SUCH_ORDER,
		BAD_API_KEY_FMT:                               &FuturesErr_BAD_API_KEY_FMT,
		REJECTED_MBX_KEY:                              &FuturesErr_REJECTED_MBX_KEY,
		NO_TRADING_WINDOW:                             &FuturesErr_NO_TRADING_WINDOW,
		API_KEYS_LOCKED:                               &FuturesErr_API_KEYS_LOCKED,
		BALANCE_NOT_SUFFICIENT:                        &FuturesErr_BALANCE_NOT_SUFFICIENT,
		MARGIN_NOT_SUFFICIEN:                          &FuturesErr_MARGIN_NOT_SUFFICIEN,
		UNABLE_TO_FILL:                                &FuturesErr_UNABLE_TO_FILL,
		ORDER_WOULD_IMMEDIATELY_TRIGGER:               &FuturesErr_ORDER_WOULD_IMMEDIATELY_TRIGGER,
		REDUCE_ONLY_REJECT:                            &FuturesErr_REDUCE_ONLY_REJECT,
		USER_IN_LIQUIDATION:                           &FuturesErr_USER_IN_LIQUIDATION,
		POSITION_NOT_SUFFICIENT:                       &FuturesErr_POSITION_NOT_SUFFICIENT,
		MAX_OPEN_ORDER_EXCEEDED:                       &FuturesErr_MAX_OPEN_ORDER_EXCEEDED,
		REDUCE_ONLY_ORDER_TYPE_NOT_SUPPORTED:          &FuturesErr_REDUCE_ONLY_ORDER_TYPE_NOT_SUPPORTED,
		MAX_LEVERAGE_RATIO:                            &FuturesErr_MAX_LEVERAGE_RATIO,
		MIN_LEVERAGE_RATIO:                            &FuturesErr_MIN_LEVERAGE_RATIO,
		INVALID_ORDER_STATUS:                          &FuturesErr_INVALID_ORDER_STATUS,
		PRICE_LESS_THAN_ZERO:                          &FuturesErr_PRICE_LESS_THAN_ZERO,
		PRICE_GREATER_THAN_MAX_PRICE:                  &FuturesErr_PRICE_GREATER_THAN_MAX_PRICE,
		QTY_LESS_THAN_ZERO:                            &FuturesErr_QTY_LESS_THAN_ZERO,
		QTY_LESS_THAN_MIN_QTY:                         &FuturesErr_QTY_LESS_THAN_MIN_QTY,
		QTY_GREATER_THAN_MAX_QTY:                      &FuturesErr_QTY_GREATER_THAN_MAX_QTY,
		STOP_PRICE_LESS_THAN_ZERO:                     &FuturesErr_STOP_PRICE_LESS_THAN_ZERO,
		STOP_PRICE_GREATER_THAN_MAX_PRICE:             &FuturesErr_STOP_PRICE_GREATER_THAN_MAX_PRICE,
		TICK_SIZE_LESS_THAN_ZERO:                      &FuturesErr_TICK_SIZE_LESS_THAN_ZERO,
		MAX_PRICE_LESS_THAN_MIN_PRICE:                 &FuturesErr_MAX_PRICE_LESS_THAN_MIN_PRICE,
		MAX_QTY_LESS_THAN_MIN_QTY:                     &FuturesErr_MAX_QTY_LESS_THAN_MIN_QTY,
		STEP_SIZE_LESS_THAN_ZERO:                      &FuturesErr_STEP_SIZE_LESS_THAN_ZERO,
		MAX_NUM_ORDERS_LESS_THAN_ZERO:                 &FuturesErr_MAX_NUM_ORDERS_LESS_THAN_ZERO,
		PRICE_LESS_THAN_MIN_PRICE:                     &FuturesErr_PRICE_LESS_THAN_MIN_PRICE,
		PRICE_NOT_INCREASED_BY_TICK_SIZE:              &FuturesErr_PRICE_NOT_INCREASED_BY_TICK_SIZE,
		INVALID_CL_ORD_ID_LEN:                         &FuturesErr_INVALID_CL_ORD_ID_LEN,
		PRICE_HIGHTER_THAN_MULTIPLIER_UP:              &FuturesErr_PRICE_HIGHTER_THAN_MULTIPLIER_UP,
		MULTIPLIER_UP_LESS_THAN_ZERO:                  &FuturesErr_MULTIPLIER_UP_LESS_THAN_ZERO,
		MULTIPLIER_DOWN_LESS_THAN_ZERO:                &FuturesErr_MULTIPLIER_DOWN_LESS_THAN_ZERO,
		COMPOSITE_SCALE_OVERFLOW:                      &FuturesErr_COMPOSITE_SCALE_OVERFLOW,
		TARGET_STRATEGY_INVALID:                       &FuturesErr_TARGET_STRATEGY_INVALID,
		INVALID_DEPTH_LIMIT:                           &FuturesErr_INVALID_DEPTH_LIMIT,
		WRONG_MARKET_STATUS:                           &FuturesErr_WRONG_MARKET_STATUS,
		QTY_NOT_INCREASED_BY_STEP_SIZE:                &FuturesErr_QTY_NOT_INCREASED_BY_STEP_SIZE,
		PRICE_LOWER_THAN_MULTIPLIER_DOWN:              &FuturesErr_PRICE_LOWER_THAN_MULTIPLIER_DOWN,
		MULTIPLIER_DECIMAL_LESS_THAN_ZERO:             &FuturesErr_MULTIPLIER_DECIMAL_LESS_THAN_ZERO,
		COMMISSION_INVALID:                            &FuturesErr_COMMISSION_INVALID,
		INVALID_ACCOUNT_TYPE:                          &FuturesErr_INVALID_ACCOUNT_TYPE,
		INVALID_LEVERAGE:                              &FuturesErr_INVALID_LEVERAGE,
		INVALID_TICK_SIZE_PRECISION:                   &FuturesErr_INVALID_TICK_SIZE_PRECISION,
		INVALID_STEP_SIZE_PRECISION:                   &FuturesErr_INVALID_STEP_SIZE_PRECISION,
		INVALID_WORKING_TYPE:                          &FuturesErr_INVALID_WORKING_TYPE,
		EXCEED_MAX_CANCEL_ORDER_SIZE:                  &FuturesErr_EXCEED_MAX_CANCEL_ORDER_SIZE,
		INSURANCE_ACCOUNT_NOT_FOUND:                   &FuturesErr_INSURANCE_ACCOUNT_NOT_FOUND,
		INVALID_BALANCE_TYPE:                          &FuturesErr_INVALID_BALANCE_TYPE,
		MAX_STOP_ORDER_EXCEEDED:                       &FuturesErr_MAX_STOP_ORDER_EXCEEDED,
		NO_NEED_TO_CHANGE_MARGIN_TYPE:                 &FuturesErr_NO_NEED_TO_CHANGE_MARGIN_TYPE,
		THERE_EXISTS_OPEN_ORDERS:                      &FuturesErr_THERE_EXISTS_OPEN_ORDERS,
		THERE_EXISTS_QUANTITY:                         &FuturesErr_THERE_EXISTS_QUANTITY,
		ADD_ISOLATED_MARGIN_REJECT:                    &FuturesErr_ADD_ISOLATED_MARGIN_REJECT,
		CROSS_BALANCE_INSUFFICIENT:                    &FuturesErr_CROSS_BALANCE_INSUFFICIENT,
		ISOLATED_BALANCE_INSUFFICIENT:                 &FuturesErr_ISOLATED_BALANCE_INSUFFICIENT,
		NO_NEED_TO_CHANGE_AUTO_ADD_MARGIN:             &FuturesErr_NO_NEED_TO_CHANGE_AUTO_ADD_MARGIN,
		AUTO_ADD_CROSSED_MARGIN_REJECT:                &FuturesErr_AUTO_ADD_CROSSED_MARGIN_REJECT,
		ADD_ISOLATED_MARGIN_NO_POSITION_REJECT:        &FuturesErr_ADD_ISOLATED_MARGIN_NO_POSITION_REJECT,
		AMOUNT_MUST_BE_POSITIVE:                       &FuturesErr_AMOUNT_MUST_BE_POSITIVE,
		INVALID_API_KEY_TYPE:                          &FuturesErr_INVALID_API_KEY_TYPE,
		INVALID_RSA_PUBLIC_KEY:                        &FuturesErr_INVALID_RSA_PUBLIC_KEY,
		MAX_PRICE_TOO_LARGE:                           &FuturesErr_MAX_PRICE_TOO_LARGE,
		NO_NEED_TO_CHANGE_POSITION_SIDE:               &FuturesErr_NO_NEED_TO_CHANGE_POSITION_SIDE,
		INVALID_POSITION_SIDE:                         &FuturesErr_INVALID_POSITION_SIDE,
		POSITION_SIDE_NOT_MATCH:                       &FuturesErr_POSITION_SIDE_NOT_MATCH,
		REDUCE_ONLY_CONFLICT:                          &FuturesErr_REDUCE_ONLY_CONFLICT,
		INVALID_OPTIONS_REQUEST_TYPE:                  &FuturesErr_INVALID_OPTIONS_REQUEST_TYPE,
		INVALID_OPTIONS_TIME_FRAME:                    &FuturesErr_INVALID_OPTIONS_TIME_FRAME,
		INVALID_OPTIONS_AMOUNT:                        &FuturesErr_INVALID_OPTIONS_AMOUNT,
		INVALID_OPTIONS_EVENT_TYPE:                    &FuturesErr_INVALID_OPTIONS_EVENT_TYPE,
		POSITION_SIDE_CHANGE_EXISTS_OPEN_ORDERS:       &FuturesErr_POSITION_SIDE_CHANGE_EXISTS_OPEN_ORDERS,
		POSITION_SIDE_CHANGE_EXISTS_QUANTITY:          &FuturesErr_POSITION_SIDE_CHANGE_EXISTS_QUANTITY,
		INVALID_OPTIONS_PREMIUM_FEE:                   &FuturesErr_INVALID_OPTIONS_PREMIUM_FEE,
		INVALID_CL_OPTIONS_ID_LEN:                     &FuturesErr_INVALID_CL_OPTIONS_ID_LEN,
		INVALID_OPTIONS_DIRECTION:                     &FuturesErr_INVALID_OPTIONS_DIRECTION,
		OPTIONS_PREMIUM_NOT_UPDATE:                    &FuturesErr_OPTIONS_PREMIUM_NOT_UPDATE,
		OPTIONS_PREMIUM_INPUT_LESS_THAN_ZERO:          &FuturesErr_OPTIONS_PREMIUM_INPUT_LESS_THAN_ZERO,
		OPTIONS_AMOUNT_BIGGER_THAN_UPPER:              &FuturesErr_OPTIONS_AMOUNT_BIGGER_THAN_UPPER,
		OPTIONS_PREMIUM_OUTPUT_ZERO:                   &FuturesErr_OPTIONS_PREMIUM_OUTPUT_ZERO,
		OPTIONS_PREMIUM_TOO_DIFF:                      &FuturesErr_OPTIONS_PREMIUM_TOO_DIFF,
		OPTIONS_PREMIUM_REACH_LIMIT:                   &FuturesErr_OPTIONS_PREMIUM_REACH_LIMIT,
		OPTIONS_COMMON_ERROR:                          &FuturesErr_OPTIONS_COMMON_ERROR,
		INVALID_OPTIONS_ID:                            &FuturesErr_INVALID_OPTIONS_ID,
		OPTIONS_USER_NOT_FOUND:                        &FuturesErr_OPTIONS_USER_NOT_FOUND,
		OPTIONS_NOT_FOUND:                             &FuturesErr_OPTIONS_NOT_FOUND,
		INVALID_BATCH_PLACE_ORDER_SIZE:                &FuturesErr_INVALID_BATCH_PLACE_ORDER_SIZE,
		PLACE_BATCH_ORDERS_FAIL:                       &FuturesErr_PLACE_BATCH_ORDERS_FAIL,
		UPCOMING_METHOD:                               &FuturesErr_UPCOMING_METHOD,
		INVALID_NOTIONAL_LIMIT_COEF:                   &FuturesErr_INVALID_NOTIONAL_LIMIT_COEF,
		INVALID_PRICE_SPREAD_THRESHOLD:                &FuturesErr_INVALID_PRICE_SPREAD_THRESHOLD,
		REDUCE_ONLY_ORDER_PERMISSION:                  &FuturesErr_REDUCE_ONLY_ORDER_PERMISSION,
		NO_PLACE_ORDER_PERMISSION:                     &FuturesErr_NO_PLACE_ORDER_PERMISSION,
		INVALID_CONTRACT_TYPE:                         &FuturesErr_INVALID_CONTRACT_TYPE,
		INVALID_CLIENT_TRAN_ID_LEN:                    &FuturesErr_INVALID_CLIENT_TRAN_ID_LEN,
		DUPLICATED_CLIENT_TRAN_ID:                     &FuturesErr_DUPLICATED_CLIENT_TRAN_ID,
		DUPLICATED_CLIENT_ORDER_ID:                    &FuturesErr_DUPLICATED_CLIENT_ORDER_ID,
		STOP_ORDER_TRIGGERING:                         &FuturesErr_STOP_ORDER_TRIGGERING,
		REDUCE_ONLY_MARGIN_CHECK_FAILED:               &FuturesErr_REDUCE_ONLY_MARGIN_CHECK_FAILED,
		MARKET_ORDER_REJECT:                           &FuturesErr_MARKET_ORDER_REJECT,
		INVALID_ACTIVATION_PRICE:                      &FuturesErr_INVALID_ACTIVATION_PRICE,
		QUANTITY_EXISTS_WITH_CLOSE_POSITION:           &FuturesErr_QUANTITY_EXISTS_WITH_CLOSE_POSITION,
		REDUCE_ONLY_MUST_BE_TRUE:                      &FuturesErr_REDUCE_ONLY_MUST_BE_TRUE,
		ORDER_TYPE_CANNOT_BE_MKT:                      &FuturesErr_ORDER_TYPE_CANNOT_BE_MKT,
		INVALID_OPENING_POSITION_STATUS:               &FuturesErr_INVALID_OPENING_POSITION_STATUS,
		SYMBOL_ALREADY_CLOSED:                         &FuturesErr_SYMBOL_ALREADY_CLOSED,
		STRATEGY_INVALID_TRIGGER_PRICE:                &FuturesErr_STRATEGY_INVALID_TRIGGER_PRICE,
		INVALID_PAIR:                                  &FuturesErr_INVALID_PAIR,
		ISOLATED_LEVERAGE_REJECT_WITH_POSITION:        &FuturesErr_ISOLATED_LEVERAGE_REJECT_WITH_POSITION,
		MIN_NOTIONAL:                                  &FuturesErr_MIN_NOTIONAL,
		INVALID_TIME_INTERVAL:                         &FuturesErr_INVALID_TIME_INTERVAL,
		ISOLATED_REJECT_WITH_JOINT_MARGIN:             &FuturesErr_ISOLATED_REJECT_WITH_JOINT_MARGIN,
		JOINT_MARGIN_REJECT_WITH_ISOLATED:             &FuturesErr_JOINT_MARGIN_REJECT_WITH_ISOLATED,
		JOINT_MARGIN_REJECT_WITH_MB:                   &FuturesErr_JOINT_MARGIN_REJECT_WITH_MB,
		JOINT_MARGIN_REJECT_WITH_OPEN_ORDER:           &FuturesErr_JOINT_MARGIN_REJECT_WITH_OPEN_ORDER,
		NO_NEED_TO_CHANGE_JOINT_MARGIN:                &FuturesErr_NO_NEED_TO_CHANGE_JOINT_MARGIN,
		JOINT_MARGIN_REJECT_WITH_NEGATIVE_BALANCE:     &FuturesErr_JOINT_MARGIN_REJECT_WITH_NEGATIVE_BALANCE,
		ISOLATED_REJECT_WITH_JOINT_MARGIN_PRICE:       &FuturesErr_ISOLATED_REJECT_WITH_JOINT_MARGIN_PRICE,
		PRICE_LOWER_THAN_STOP_MULTIPLIER_DOWN:         &FuturesErr_PRICE_LOWER_THAN_STOP_MULTIPLIER_DOWN,
		COOLING_OFF_PERIOD:                            &FuturesErr_COOLING_OFF_PERIOD,
		ADJUST_LEVERAGE_KYC_FAILED:                    &FuturesErr_ADJUST_LEVERAGE_KYC_FAILED,
		ADJUST_LEVERAGE_ONE_MONTH_FAILED:              &FuturesErr_ADJUST_LEVERAGE_ONE_MONTH_FAILED,
		ADJUST_LEVERAGE_X_DAYS_FAILED:                 &FuturesErr_ADJUST_LEVERAGE_X_DAYS_FAILED,
		ADJUST_LEVERAGE_KYC_LIMIT:                     &FuturesErr_ADJUST_LEVERAGE_KYC_LIMIT,
		ADJUST_LEVERAGE_ACCOUNT_SYMBOL_FAILED:         &FuturesErr_ADJUST_LEVERAGE_ACCOUNT_SYMBOL_FAILED,
		ADJUST_LEVERAGE_SYMBOL_FAILED:                 &FuturesErr_ADJUST_LEVERAGE_SYMBOL_FAILED,
		STOP_PRICE_HIGHER_THAN_PRICE_MULTIPLIER_LIMIT: &FuturesErr_STOP_PRICE_HIGHER_THAN_PRICE_MULTIPLIER_LIMIT,
		STOP_PRICE_LOWER_THAN_PRICE_MULTIPLIER_LIMIT:  &FuturesErr_STOP_PRICE_LOWER_THAN_PRICE_MULTIPLIER_LIMIT,
		TRADING_QUANTITATIVE_RULE:                     &FuturesErr_TRADING_QUANTITATIVE_RULE,
		LARGE_POSITION_SYM_RULE:                       &FuturesErr_LARGE_POSITION_SYM_RULE,
		COMPLIANCE_BLACK_SYMBOL_RESTRICTION:           &FuturesErr_COMPLIANCE_BLACK_SYMBOL_RESTRICTION,
		ADJUST_LEVERAGE_COMPLIANCE_FAILED:             &FuturesErr_ADJUST_LEVERAGE_COMPLIANCE_FAILED,
		FOK_ORDER_REJECT:                              &FuturesErr_FOK_ORDER_REJECT,
		GTX_ORDER_REJECT:                              &FuturesErr_GTX_ORDER_REJECT,
		MOVE_ORDER_NOT_ALLOWED_SYMBOL_REASON:          &FuturesErr_MOVE_ORDER_NOT_ALLOWED_SYMBOL_REASON,
		LIMIT_ORDER_ONLY:                              &FuturesErr_LIMIT_ORDER_ONLY,
		Exceed_Maximum_Modify_Order_Limit:             &FuturesErr_Exceed_Maximum_Modify_Order_Limit,
		SAME_ORDER:                                    &FuturesErr_SAME_ORDER,
		ME_RECVWINDOW_REJECT:                          &FuturesErr_ME_RECVWINDOW_REJECT,
		MODIFICATION_MIN_NOTIONAL:                     &FuturesErr_MODIFICATION_MIN_NOTIONAL,
		INVALID_PRICE_MATCH:                           &FuturesErr_INVALID_PRICE_MATCH,
		UNSUPPORTED_ORDER_TYPE_PRICE_MATCH:            &FuturesErr_UNSUPPORTED_ORDER_TYPE_PRICE_MATCH,
		INVALID_SELF_TRADE_PREVENTION_MODE:            &FuturesErr_INVALID_SELF_TRADE_PREVENTION_MODE,
		FUTURE_GOOD_TILL_DATE:                         &FuturesErr_FUTURE_GOOD_TILL_DATE,
		BBO_ORDER_REJECT:                              &FuturesErr_BBO_ORDER_REJECT,
	},
	Codes: FuturesCodes{
		N1000: &FuturesErr_UNKNOWN,
		N1001: &FuturesErr_DISCONNECTED,
		N1002: &FuturesErr_UNAUTHORIZED,
		N1003: &FuturesErr_TOO_MANY_REQUESTS,
		N1004: &FuturesErr_DUPLICATE_IP,
		N1005: &FuturesErr_NO_SUCH_IP,
		N1006: &FuturesErr_UNEXPECTED_RESP,
		N1007: &FuturesErr_TIMEOUT,
		N1008: &FuturesErr_Service,
		N1010: &FuturesErr_ERROR_MSG_RECEIVED,
		N1011: &FuturesErr_NON_WHITE_LIST,
		N1013: &FuturesErr_INVALID_MESSAGE,
		N1014: &FuturesErr_UNKNOWN_ORDER_COMPOSITION,
		N1015: &FuturesErr_TOO_MANY_ORDERS,
		N1016: &FuturesErr_SERVICE_SHUTTING_DOWN,
		N1020: &FuturesErr_UNSUPPORTED_OPERATION,
		N1021: &FuturesErr_INVALID_TIMESTAMP,
		N1022: &FuturesErr_INVALID_SIGNATURE,
		N1023: &FuturesErr_START_TIME_GREATER_THAN_END_TIME,
		N1099: &FuturesErr_NOT_FOUND,
		N1100: &FuturesErr_ILLEGAL_CHARS,
		N1101: &FuturesErr_TOO_MANY_PARAMETERS,
		N1102: &FuturesErr_MANDATORY_PARAM_EMPTY_OR_MALFORMED,
		N1103: &FuturesErr_UNKNOWN_PARAM,
		N1104: &FuturesErr_UNREAD_PARAMETERS,
		N1105: &FuturesErr_PARAM_EMPTY,
		N1106: &FuturesErr_PARAM_NOT_REQUIRED,
		N1108: &FuturesErr_BAD_ASSET,
		N1109: &FuturesErr_BAD_ACCOUNT,
		N1110: &FuturesErr_BAD_INSTRUMENT_TYPE,
		N1111: &FuturesErr_BAD_PRECISION,
		N1112: &FuturesErr_NO_DEPTH,
		N1113: &FuturesErr_WITHDRAW_NOT_NEGATIVE,
		N1114: &FuturesErr_TIF_NOT_REQUIRED,
		N1115: &FuturesErr_INVALID_TIF,
		N1116: &FuturesErr_INVALID_ORDER_TYPE,
		N1117: &FuturesErr_INVALID_SIDE,
		N1118: &FuturesErr_EMPTY_NEW_CL_ORD_ID,
		N1119: &FuturesErr_EMPTY_ORG_CL_ORD_ID,
		N1120: &FuturesErr_BAD_INTERVAL,
		N1121: &FuturesErr_BAD_SYMBOL,
		N1122: &FuturesErr_INVALID_SYMBOL_STATUS,
		N1125: &FuturesErr_INVALID_LISTEN_KEY,
		N1126: &FuturesErr_ASSET_NOT_SUPPORTED,
		N1127: &FuturesErr_MORE_THAN_XX_HOURS,
		N1128: &FuturesErr_OPTIONAL_PARAMS_BAD_COMBO,
		N1130: &FuturesErr_INVALID_PARAMETER,
		N1136: &FuturesErr_INVALID_NEW_ORDER_RESP_TYPE,
		N2010: &FuturesErr_NEW_ORDER_REJECTED,
		N2011: &FuturesErr_CANCEL_REJECTED,
		N2012: &FuturesErr_CANCEL_ALL_FAIL,
		N2013: &FuturesErr_NO_SUCH_ORDER,
		N2014: &FuturesErr_BAD_API_KEY_FMT,
		N2015: &FuturesErr_REJECTED_MBX_KEY,
		N2016: &FuturesErr_NO_TRADING_WINDOW,
		N2017: &FuturesErr_API_KEYS_LOCKED,
		N2018: &FuturesErr_BALANCE_NOT_SUFFICIENT,
		N2019: &FuturesErr_MARGIN_NOT_SUFFICIEN,
		N2020: &FuturesErr_UNABLE_TO_FILL,
		N2021: &FuturesErr_ORDER_WOULD_IMMEDIATELY_TRIGGER,
		N2022: &FuturesErr_REDUCE_ONLY_REJECT,
		N2023: &FuturesErr_USER_IN_LIQUIDATION,
		N2024: &FuturesErr_POSITION_NOT_SUFFICIENT,
		N2025: &FuturesErr_MAX_OPEN_ORDER_EXCEEDED,
		N2026: &FuturesErr_REDUCE_ONLY_ORDER_TYPE_NOT_SUPPORTED,
		N2027: &FuturesErr_MAX_LEVERAGE_RATIO,
		N2028: &FuturesErr_MIN_LEVERAGE_RATIO,
		N4000: &FuturesErr_INVALID_ORDER_STATUS,
		N4001: &FuturesErr_PRICE_LESS_THAN_ZERO,
		N4002: &FuturesErr_PRICE_GREATER_THAN_MAX_PRICE,
		N4003: &FuturesErr_QTY_LESS_THAN_ZERO,
		N4004: &FuturesErr_QTY_LESS_THAN_MIN_QTY,
		N4005: &FuturesErr_QTY_GREATER_THAN_MAX_QTY,
		N4006: &FuturesErr_STOP_PRICE_LESS_THAN_ZERO,
		N4007: &FuturesErr_STOP_PRICE_GREATER_THAN_MAX_PRICE,
		N4008: &FuturesErr_TICK_SIZE_LESS_THAN_ZERO,
		N4009: &FuturesErr_MAX_PRICE_LESS_THAN_MIN_PRICE,
		N4010: &FuturesErr_MAX_QTY_LESS_THAN_MIN_QTY,
		N4011: &FuturesErr_STEP_SIZE_LESS_THAN_ZERO,
		N4012: &FuturesErr_MAX_NUM_ORDERS_LESS_THAN_ZERO,
		N4013: &FuturesErr_PRICE_LESS_THAN_MIN_PRICE,
		N4014: &FuturesErr_PRICE_NOT_INCREASED_BY_TICK_SIZE,
		N4015: &FuturesErr_INVALID_CL_ORD_ID_LEN,
		N4016: &FuturesErr_PRICE_HIGHTER_THAN_MULTIPLIER_UP,
		N4017: &FuturesErr_MULTIPLIER_UP_LESS_THAN_ZERO,
		N4018: &FuturesErr_MULTIPLIER_DOWN_LESS_THAN_ZERO,
		N4019: &FuturesErr_COMPOSITE_SCALE_OVERFLOW,
		N4020: &FuturesErr_TARGET_STRATEGY_INVALID,
		N4021: &FuturesErr_INVALID_DEPTH_LIMIT,
		N4022: &FuturesErr_WRONG_MARKET_STATUS,
		N4023: &FuturesErr_QTY_NOT_INCREASED_BY_STEP_SIZE,
		N4024: &FuturesErr_PRICE_LOWER_THAN_MULTIPLIER_DOWN,
		N4025: &FuturesErr_MULTIPLIER_DECIMAL_LESS_THAN_ZERO,
		N4026: &FuturesErr_COMMISSION_INVALID,
		N4027: &FuturesErr_INVALID_ACCOUNT_TYPE,
		N4028: &FuturesErr_INVALID_LEVERAGE,
		N4029: &FuturesErr_INVALID_TICK_SIZE_PRECISION,
		N4030: &FuturesErr_INVALID_STEP_SIZE_PRECISION,
		N4031: &FuturesErr_INVALID_WORKING_TYPE,
		N4032: &FuturesErr_EXCEED_MAX_CANCEL_ORDER_SIZE,
		N4033: &FuturesErr_INSURANCE_ACCOUNT_NOT_FOUND,
		N4044: &FuturesErr_INVALID_BALANCE_TYPE,
		N4045: &FuturesErr_MAX_STOP_ORDER_EXCEEDED,
		N4046: &FuturesErr_NO_NEED_TO_CHANGE_MARGIN_TYPE,
		N4047: &FuturesErr_THERE_EXISTS_OPEN_ORDERS,
		N4048: &FuturesErr_THERE_EXISTS_QUANTITY,
		N4049: &FuturesErr_ADD_ISOLATED_MARGIN_REJECT,
		N4050: &FuturesErr_CROSS_BALANCE_INSUFFICIENT,
		N4051: &FuturesErr_ISOLATED_BALANCE_INSUFFICIENT,
		N4052: &FuturesErr_NO_NEED_TO_CHANGE_AUTO_ADD_MARGIN,
		N4053: &FuturesErr_AUTO_ADD_CROSSED_MARGIN_REJECT,
		N4054: &FuturesErr_ADD_ISOLATED_MARGIN_NO_POSITION_REJECT,
		N4055: &FuturesErr_AMOUNT_MUST_BE_POSITIVE,
		N4056: &FuturesErr_INVALID_API_KEY_TYPE,
		N4057: &FuturesErr_INVALID_RSA_PUBLIC_KEY,
		N4058: &FuturesErr_MAX_PRICE_TOO_LARGE,
		N4059: &FuturesErr_NO_NEED_TO_CHANGE_POSITION_SIDE,
		N4060: &FuturesErr_INVALID_POSITION_SIDE,
		N4061: &FuturesErr_POSITION_SIDE_NOT_MATCH,
		N4062: &FuturesErr_REDUCE_ONLY_CONFLICT,
		N4063: &FuturesErr_INVALID_OPTIONS_REQUEST_TYPE,
		N4064: &FuturesErr_INVALID_OPTIONS_TIME_FRAME,
		N4065: &FuturesErr_INVALID_OPTIONS_AMOUNT,
		N4066: &FuturesErr_INVALID_OPTIONS_EVENT_TYPE,
		N4067: &FuturesErr_POSITION_SIDE_CHANGE_EXISTS_OPEN_ORDERS,
		N4068: &FuturesErr_POSITION_SIDE_CHANGE_EXISTS_QUANTITY,
		N4069: &FuturesErr_INVALID_OPTIONS_PREMIUM_FEE,
		N4070: &FuturesErr_INVALID_CL_OPTIONS_ID_LEN,
		N4071: &FuturesErr_INVALID_OPTIONS_DIRECTION,
		N4072: &FuturesErr_OPTIONS_PREMIUM_NOT_UPDATE,
		N4073: &FuturesErr_OPTIONS_PREMIUM_INPUT_LESS_THAN_ZERO,
		N4074: &FuturesErr_OPTIONS_AMOUNT_BIGGER_THAN_UPPER,
		N4075: &FuturesErr_OPTIONS_PREMIUM_OUTPUT_ZERO,
		N4076: &FuturesErr_OPTIONS_PREMIUM_TOO_DIFF,
		N4077: &FuturesErr_OPTIONS_PREMIUM_REACH_LIMIT,
		N4078: &FuturesErr_OPTIONS_COMMON_ERROR,
		N4079: &FuturesErr_INVALID_OPTIONS_ID,
		N4080: &FuturesErr_OPTIONS_USER_NOT_FOUND,
		N4081: &FuturesErr_OPTIONS_NOT_FOUND,
		N4082: &FuturesErr_INVALID_BATCH_PLACE_ORDER_SIZE,
		N4083: &FuturesErr_PLACE_BATCH_ORDERS_FAIL,
		N4084: &FuturesErr_UPCOMING_METHOD,
		N4085: &FuturesErr_INVALID_NOTIONAL_LIMIT_COEF,
		N4086: &FuturesErr_INVALID_PRICE_SPREAD_THRESHOLD,
		N4087: &FuturesErr_REDUCE_ONLY_ORDER_PERMISSION,
		N4088: &FuturesErr_NO_PLACE_ORDER_PERMISSION,
		N4104: &FuturesErr_INVALID_CONTRACT_TYPE,
		N4114: &FuturesErr_INVALID_CLIENT_TRAN_ID_LEN,
		N4115: &FuturesErr_DUPLICATED_CLIENT_TRAN_ID,
		N4116: &FuturesErr_DUPLICATED_CLIENT_ORDER_ID,
		N4117: &FuturesErr_STOP_ORDER_TRIGGERING,
		N4118: &FuturesErr_REDUCE_ONLY_MARGIN_CHECK_FAILED,
		N4131: &FuturesErr_MARKET_ORDER_REJECT,
		N4135: &FuturesErr_INVALID_ACTIVATION_PRICE,
		N4137: &FuturesErr_QUANTITY_EXISTS_WITH_CLOSE_POSITION,
		N4138: &FuturesErr_REDUCE_ONLY_MUST_BE_TRUE,
		N4139: &FuturesErr_ORDER_TYPE_CANNOT_BE_MKT,
		N4140: &FuturesErr_INVALID_OPENING_POSITION_STATUS,
		N4141: &FuturesErr_SYMBOL_ALREADY_CLOSED,
		N4142: &FuturesErr_STRATEGY_INVALID_TRIGGER_PRICE,
		N4144: &FuturesErr_INVALID_PAIR,
		N4161: &FuturesErr_ISOLATED_LEVERAGE_REJECT_WITH_POSITION,
		N4164: &FuturesErr_MIN_NOTIONAL,
		N4165: &FuturesErr_INVALID_TIME_INTERVAL,
		N4167: &FuturesErr_ISOLATED_REJECT_WITH_JOINT_MARGIN,
		N4168: &FuturesErr_JOINT_MARGIN_REJECT_WITH_ISOLATED,
		N4169: &FuturesErr_JOINT_MARGIN_REJECT_WITH_MB,
		N4170: &FuturesErr_JOINT_MARGIN_REJECT_WITH_OPEN_ORDER,
		N4171: &FuturesErr_NO_NEED_TO_CHANGE_JOINT_MARGIN,
		N4172: &FuturesErr_JOINT_MARGIN_REJECT_WITH_NEGATIVE_BALANCE,
		N4183: &FuturesErr_ISOLATED_REJECT_WITH_JOINT_MARGIN,
		N4184: &FuturesErr_PRICE_LOWER_THAN_STOP_MULTIPLIER_DOWN,
		N4192: &FuturesErr_COOLING_OFF_PERIOD,
		N4202: &FuturesErr_ADJUST_LEVERAGE_KYC_FAILED,
		N4203: &FuturesErr_ADJUST_LEVERAGE_ONE_MONTH_FAILED,
		N4205: &FuturesErr_ADJUST_LEVERAGE_X_DAYS_FAILED,
		N4206: &FuturesErr_ADJUST_LEVERAGE_KYC_LIMIT,
		N4208: &FuturesErr_ADJUST_LEVERAGE_ACCOUNT_SYMBOL_FAILED,
		N4209: &FuturesErr_ADJUST_LEVERAGE_SYMBOL_FAILED,
		N4210: &FuturesErr_STOP_PRICE_HIGHER_THAN_PRICE_MULTIPLIER_LIMIT,
		N4211: &FuturesErr_STOP_PRICE_LOWER_THAN_PRICE_MULTIPLIER_LIMIT,
		N4400: &FuturesErr_TRADING_QUANTITATIVE_RULE,
		N4401: &FuturesErr_LARGE_POSITION_SYM_RULE,
		N4402: &FuturesErr_COMPLIANCE_BLACK_SYMBOL_RESTRICTION,
		N4403: &FuturesErr_ADJUST_LEVERAGE_COMPLIANCE_FAILED,
		N5021: &FuturesErr_FOK_ORDER_REJECT,
		N5022: &FuturesErr_GTX_ORDER_REJECT,
		N5024: &FuturesErr_MOVE_ORDER_NOT_ALLOWED_SYMBOL_REASON,
		N5025: &FuturesErr_LIMIT_ORDER_ONLY,
		N5026: &FuturesErr_Exceed_Maximum_Modify_Order_Limit,
		N5027: &FuturesErr_SAME_ORDER,
		N5028: &FuturesErr_ME_RECVWINDOW_REJECT,
		N5029: &FuturesErr_MODIFICATION_MIN_NOTIONAL,
		N5037: &FuturesErr_INVALID_PRICE_MATCH,
		N5038: &FuturesErr_UNSUPPORTED_ORDER_TYPE_PRICE_MATCH,
		N5039: &FuturesErr_INVALID_SELF_TRADE_PREVENTION_MODE,
		N5040: &FuturesErr_FUTURE_GOOD_TILL_DATE,
		N5041: &FuturesErr_BBO_ORDER_REJECT,
	},
}

type FuturesNames struct {
	UNKNOWN                                       *BinanceErrorCode
	DISCONNECTED                                  *BinanceErrorCode
	UNAUTHORIZED                                  *BinanceErrorCode
	TOO_MANY_REQUESTS                             *BinanceErrorCode
	DUPLICATE_IP                                  *BinanceErrorCode
	NO_SUCH_IP                                    *BinanceErrorCode
	UNEXPECTED_RESP                               *BinanceErrorCode
	TIMEOUT                                       *BinanceErrorCode
	Service                                       *BinanceErrorCode
	ERROR_MSG_RECEIVED                            *BinanceErrorCode
	NON_WHITE_LIST                                *BinanceErrorCode
	INVALID_MESSAGE                               *BinanceErrorCode
	UNKNOWN_ORDER_COMPOSITION                     *BinanceErrorCode
	TOO_MANY_ORDERS                               *BinanceErrorCode
	SERVICE_SHUTTING_DOWN                         *BinanceErrorCode
	UNSUPPORTED_OPERATION                         *BinanceErrorCode
	INVALID_TIMESTAMP                             *BinanceErrorCode
	INVALID_SIGNATURE                             *BinanceErrorCode
	START_TIME_GREATER_THAN_END_TIME              *BinanceErrorCode
	NOT_FOUND                                     *BinanceErrorCode
	ILLEGAL_CHARS                                 *BinanceErrorCode
	TOO_MANY_PARAMETERS                           *BinanceErrorCode
	MANDATORY_PARAM_EMPTY_OR_MALFORMED            *BinanceErrorCode
	UNKNOWN_PARAM                                 *BinanceErrorCode
	UNREAD_PARAMETERS                             *BinanceErrorCode
	PARAM_EMPTY                                   *BinanceErrorCode
	PARAM_NOT_REQUIRED                            *BinanceErrorCode
	BAD_ASSET                                     *BinanceErrorCode
	BAD_ACCOUNT                                   *BinanceErrorCode
	BAD_INSTRUMENT_TYPE                           *BinanceErrorCode
	BAD_PRECISION                                 *BinanceErrorCode
	NO_DEPTH                                      *BinanceErrorCode
	WITHDRAW_NOT_NEGATIVE                         *BinanceErrorCode
	TIF_NOT_REQUIRED                              *BinanceErrorCode
	INVALID_TIF                                   *BinanceErrorCode
	INVALID_ORDER_TYPE                            *BinanceErrorCode
	INVALID_SIDE                                  *BinanceErrorCode
	EMPTY_NEW_CL_ORD_ID                           *BinanceErrorCode
	EMPTY_ORG_CL_ORD_ID                           *BinanceErrorCode
	BAD_INTERVAL                                  *BinanceErrorCode
	BAD_SYMBOL                                    *BinanceErrorCode
	INVALID_SYMBOL_STATUS                         *BinanceErrorCode
	INVALID_LISTEN_KEY                            *BinanceErrorCode
	ASSET_NOT_SUPPORTED                           *BinanceErrorCode
	MORE_THAN_XX_HOURS                            *BinanceErrorCode
	OPTIONAL_PARAMS_BAD_COMBO                     *BinanceErrorCode
	INVALID_PARAMETER                             *BinanceErrorCode
	INVALID_NEW_ORDER_RESP_TYPE                   *BinanceErrorCode
	NEW_ORDER_REJECTED                            *BinanceErrorCode
	CANCEL_REJECTED                               *BinanceErrorCode
	CANCEL_ALL_FAIL                               *BinanceErrorCode
	NO_SUCH_ORDER                                 *BinanceErrorCode
	BAD_API_KEY_FMT                               *BinanceErrorCode
	REJECTED_MBX_KEY                              *BinanceErrorCode
	NO_TRADING_WINDOW                             *BinanceErrorCode
	API_KEYS_LOCKED                               *BinanceErrorCode
	BALANCE_NOT_SUFFICIENT                        *BinanceErrorCode
	MARGIN_NOT_SUFFICIEN                          *BinanceErrorCode
	UNABLE_TO_FILL                                *BinanceErrorCode
	ORDER_WOULD_IMMEDIATELY_TRIGGER               *BinanceErrorCode
	REDUCE_ONLY_REJECT                            *BinanceErrorCode
	USER_IN_LIQUIDATION                           *BinanceErrorCode
	POSITION_NOT_SUFFICIENT                       *BinanceErrorCode
	MAX_OPEN_ORDER_EXCEEDED                       *BinanceErrorCode
	REDUCE_ONLY_ORDER_TYPE_NOT_SUPPORTED          *BinanceErrorCode
	MAX_LEVERAGE_RATIO                            *BinanceErrorCode
	MIN_LEVERAGE_RATIO                            *BinanceErrorCode
	INVALID_ORDER_STATUS                          *BinanceErrorCode
	PRICE_LESS_THAN_ZERO                          *BinanceErrorCode
	PRICE_GREATER_THAN_MAX_PRICE                  *BinanceErrorCode
	QTY_LESS_THAN_ZERO                            *BinanceErrorCode
	QTY_LESS_THAN_MIN_QTY                         *BinanceErrorCode
	QTY_GREATER_THAN_MAX_QTY                      *BinanceErrorCode
	STOP_PRICE_LESS_THAN_ZERO                     *BinanceErrorCode
	STOP_PRICE_GREATER_THAN_MAX_PRICE             *BinanceErrorCode
	TICK_SIZE_LESS_THAN_ZERO                      *BinanceErrorCode
	MAX_PRICE_LESS_THAN_MIN_PRICE                 *BinanceErrorCode
	MAX_QTY_LESS_THAN_MIN_QTY                     *BinanceErrorCode
	STEP_SIZE_LESS_THAN_ZERO                      *BinanceErrorCode
	MAX_NUM_ORDERS_LESS_THAN_ZERO                 *BinanceErrorCode
	PRICE_LESS_THAN_MIN_PRICE                     *BinanceErrorCode
	PRICE_NOT_INCREASED_BY_TICK_SIZE              *BinanceErrorCode
	INVALID_CL_ORD_ID_LEN                         *BinanceErrorCode
	PRICE_HIGHTER_THAN_MULTIPLIER_UP              *BinanceErrorCode
	MULTIPLIER_UP_LESS_THAN_ZERO                  *BinanceErrorCode
	MULTIPLIER_DOWN_LESS_THAN_ZERO                *BinanceErrorCode
	COMPOSITE_SCALE_OVERFLOW                      *BinanceErrorCode
	TARGET_STRATEGY_INVALID                       *BinanceErrorCode
	INVALID_DEPTH_LIMIT                           *BinanceErrorCode
	WRONG_MARKET_STATUS                           *BinanceErrorCode
	QTY_NOT_INCREASED_BY_STEP_SIZE                *BinanceErrorCode
	PRICE_LOWER_THAN_MULTIPLIER_DOWN              *BinanceErrorCode
	MULTIPLIER_DECIMAL_LESS_THAN_ZERO             *BinanceErrorCode
	COMMISSION_INVALID                            *BinanceErrorCode
	INVALID_ACCOUNT_TYPE                          *BinanceErrorCode
	INVALID_LEVERAGE                              *BinanceErrorCode
	INVALID_TICK_SIZE_PRECISION                   *BinanceErrorCode
	INVALID_STEP_SIZE_PRECISION                   *BinanceErrorCode
	INVALID_WORKING_TYPE                          *BinanceErrorCode
	EXCEED_MAX_CANCEL_ORDER_SIZE                  *BinanceErrorCode
	INSURANCE_ACCOUNT_NOT_FOUND                   *BinanceErrorCode
	INVALID_BALANCE_TYPE                          *BinanceErrorCode
	MAX_STOP_ORDER_EXCEEDED                       *BinanceErrorCode
	NO_NEED_TO_CHANGE_MARGIN_TYPE                 *BinanceErrorCode
	THERE_EXISTS_OPEN_ORDERS                      *BinanceErrorCode
	THERE_EXISTS_QUANTITY                         *BinanceErrorCode
	ADD_ISOLATED_MARGIN_REJECT                    *BinanceErrorCode
	CROSS_BALANCE_INSUFFICIENT                    *BinanceErrorCode
	ISOLATED_BALANCE_INSUFFICIENT                 *BinanceErrorCode
	NO_NEED_TO_CHANGE_AUTO_ADD_MARGIN             *BinanceErrorCode
	AUTO_ADD_CROSSED_MARGIN_REJECT                *BinanceErrorCode
	ADD_ISOLATED_MARGIN_NO_POSITION_REJECT        *BinanceErrorCode
	AMOUNT_MUST_BE_POSITIVE                       *BinanceErrorCode
	INVALID_API_KEY_TYPE                          *BinanceErrorCode
	INVALID_RSA_PUBLIC_KEY                        *BinanceErrorCode
	MAX_PRICE_TOO_LARGE                           *BinanceErrorCode
	NO_NEED_TO_CHANGE_POSITION_SIDE               *BinanceErrorCode
	INVALID_POSITION_SIDE                         *BinanceErrorCode
	POSITION_SIDE_NOT_MATCH                       *BinanceErrorCode
	REDUCE_ONLY_CONFLICT                          *BinanceErrorCode
	INVALID_OPTIONS_REQUEST_TYPE                  *BinanceErrorCode
	INVALID_OPTIONS_TIME_FRAME                    *BinanceErrorCode
	INVALID_OPTIONS_AMOUNT                        *BinanceErrorCode
	INVALID_OPTIONS_EVENT_TYPE                    *BinanceErrorCode
	POSITION_SIDE_CHANGE_EXISTS_OPEN_ORDERS       *BinanceErrorCode
	POSITION_SIDE_CHANGE_EXISTS_QUANTITY          *BinanceErrorCode
	INVALID_OPTIONS_PREMIUM_FEE                   *BinanceErrorCode
	INVALID_CL_OPTIONS_ID_LEN                     *BinanceErrorCode
	INVALID_OPTIONS_DIRECTION                     *BinanceErrorCode
	OPTIONS_PREMIUM_NOT_UPDATE                    *BinanceErrorCode
	OPTIONS_PREMIUM_INPUT_LESS_THAN_ZERO          *BinanceErrorCode
	OPTIONS_AMOUNT_BIGGER_THAN_UPPER              *BinanceErrorCode
	OPTIONS_PREMIUM_OUTPUT_ZERO                   *BinanceErrorCode
	OPTIONS_PREMIUM_TOO_DIFF                      *BinanceErrorCode
	OPTIONS_PREMIUM_REACH_LIMIT                   *BinanceErrorCode
	OPTIONS_COMMON_ERROR                          *BinanceErrorCode
	INVALID_OPTIONS_ID                            *BinanceErrorCode
	OPTIONS_USER_NOT_FOUND                        *BinanceErrorCode
	OPTIONS_NOT_FOUND                             *BinanceErrorCode
	INVALID_BATCH_PLACE_ORDER_SIZE                *BinanceErrorCode
	PLACE_BATCH_ORDERS_FAIL                       *BinanceErrorCode
	UPCOMING_METHOD                               *BinanceErrorCode
	INVALID_NOTIONAL_LIMIT_COEF                   *BinanceErrorCode
	INVALID_PRICE_SPREAD_THRESHOLD                *BinanceErrorCode
	REDUCE_ONLY_ORDER_PERMISSION                  *BinanceErrorCode
	NO_PLACE_ORDER_PERMISSION                     *BinanceErrorCode
	INVALID_CONTRACT_TYPE                         *BinanceErrorCode
	INVALID_CLIENT_TRAN_ID_LEN                    *BinanceErrorCode
	DUPLICATED_CLIENT_TRAN_ID                     *BinanceErrorCode
	DUPLICATED_CLIENT_ORDER_ID                    *BinanceErrorCode
	STOP_ORDER_TRIGGERING                         *BinanceErrorCode
	REDUCE_ONLY_MARGIN_CHECK_FAILED               *BinanceErrorCode
	MARKET_ORDER_REJECT                           *BinanceErrorCode
	INVALID_ACTIVATION_PRICE                      *BinanceErrorCode
	QUANTITY_EXISTS_WITH_CLOSE_POSITION           *BinanceErrorCode
	REDUCE_ONLY_MUST_BE_TRUE                      *BinanceErrorCode
	ORDER_TYPE_CANNOT_BE_MKT                      *BinanceErrorCode
	INVALID_OPENING_POSITION_STATUS               *BinanceErrorCode
	SYMBOL_ALREADY_CLOSED                         *BinanceErrorCode
	STRATEGY_INVALID_TRIGGER_PRICE                *BinanceErrorCode
	INVALID_PAIR                                  *BinanceErrorCode
	ISOLATED_LEVERAGE_REJECT_WITH_POSITION        *BinanceErrorCode
	MIN_NOTIONAL                                  *BinanceErrorCode
	INVALID_TIME_INTERVAL                         *BinanceErrorCode
	ISOLATED_REJECT_WITH_JOINT_MARGIN             *BinanceErrorCode
	JOINT_MARGIN_REJECT_WITH_ISOLATED             *BinanceErrorCode
	JOINT_MARGIN_REJECT_WITH_MB                   *BinanceErrorCode
	JOINT_MARGIN_REJECT_WITH_OPEN_ORDER           *BinanceErrorCode
	NO_NEED_TO_CHANGE_JOINT_MARGIN                *BinanceErrorCode
	JOINT_MARGIN_REJECT_WITH_NEGATIVE_BALANCE     *BinanceErrorCode
	ISOLATED_REJECT_WITH_JOINT_MARGIN_PRICE       *BinanceErrorCode
	PRICE_LOWER_THAN_STOP_MULTIPLIER_DOWN         *BinanceErrorCode
	COOLING_OFF_PERIOD                            *BinanceErrorCode
	ADJUST_LEVERAGE_KYC_FAILED                    *BinanceErrorCode
	ADJUST_LEVERAGE_ONE_MONTH_FAILED              *BinanceErrorCode
	ADJUST_LEVERAGE_X_DAYS_FAILED                 *BinanceErrorCode
	ADJUST_LEVERAGE_KYC_LIMIT                     *BinanceErrorCode
	ADJUST_LEVERAGE_ACCOUNT_SYMBOL_FAILED         *BinanceErrorCode
	ADJUST_LEVERAGE_SYMBOL_FAILED                 *BinanceErrorCode
	STOP_PRICE_HIGHER_THAN_PRICE_MULTIPLIER_LIMIT *BinanceErrorCode
	STOP_PRICE_LOWER_THAN_PRICE_MULTIPLIER_LIMIT  *BinanceErrorCode
	TRADING_QUANTITATIVE_RULE                     *BinanceErrorCode
	LARGE_POSITION_SYM_RULE                       *BinanceErrorCode
	COMPLIANCE_BLACK_SYMBOL_RESTRICTION           *BinanceErrorCode
	ADJUST_LEVERAGE_COMPLIANCE_FAILED             *BinanceErrorCode
	FOK_ORDER_REJECT                              *BinanceErrorCode
	GTX_ORDER_REJECT                              *BinanceErrorCode
	MOVE_ORDER_NOT_ALLOWED_SYMBOL_REASON          *BinanceErrorCode
	LIMIT_ORDER_ONLY                              *BinanceErrorCode
	Exceed_Maximum_Modify_Order_Limit             *BinanceErrorCode
	SAME_ORDER                                    *BinanceErrorCode
	ME_RECVWINDOW_REJECT                          *BinanceErrorCode
	MODIFICATION_MIN_NOTIONAL                     *BinanceErrorCode
	INVALID_PRICE_MATCH                           *BinanceErrorCode
	UNSUPPORTED_ORDER_TYPE_PRICE_MATCH            *BinanceErrorCode
	INVALID_SELF_TRADE_PREVENTION_MODE            *BinanceErrorCode
	FUTURE_GOOD_TILL_DATE                         *BinanceErrorCode
	BBO_ORDER_REJECT                              *BinanceErrorCode
}
type FuturesCodes struct {
	N1000 *BinanceErrorCode
	N1001 *BinanceErrorCode
	N1002 *BinanceErrorCode
	N1003 *BinanceErrorCode
	N1004 *BinanceErrorCode
	N1005 *BinanceErrorCode
	N1006 *BinanceErrorCode
	N1007 *BinanceErrorCode
	N1008 *BinanceErrorCode
	N1010 *BinanceErrorCode
	N1011 *BinanceErrorCode
	N1013 *BinanceErrorCode
	N1014 *BinanceErrorCode
	N1015 *BinanceErrorCode
	N1016 *BinanceErrorCode
	N1020 *BinanceErrorCode
	N1021 *BinanceErrorCode
	N1022 *BinanceErrorCode
	N1023 *BinanceErrorCode
	N1099 *BinanceErrorCode
	N1100 *BinanceErrorCode
	N1101 *BinanceErrorCode
	N1102 *BinanceErrorCode
	N1103 *BinanceErrorCode
	N1104 *BinanceErrorCode
	N1105 *BinanceErrorCode
	N1106 *BinanceErrorCode
	N1108 *BinanceErrorCode
	N1109 *BinanceErrorCode
	N1110 *BinanceErrorCode
	N1111 *BinanceErrorCode
	N1112 *BinanceErrorCode
	N1113 *BinanceErrorCode
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
	N1126 *BinanceErrorCode
	N1127 *BinanceErrorCode
	N1128 *BinanceErrorCode
	N1130 *BinanceErrorCode
	N1136 *BinanceErrorCode
	N2010 *BinanceErrorCode
	N2011 *BinanceErrorCode
	N2012 *BinanceErrorCode
	N2013 *BinanceErrorCode
	N2014 *BinanceErrorCode
	N2015 *BinanceErrorCode
	N2016 *BinanceErrorCode
	N2017 *BinanceErrorCode
	N2018 *BinanceErrorCode
	N2019 *BinanceErrorCode
	N2020 *BinanceErrorCode
	N2021 *BinanceErrorCode
	N2022 *BinanceErrorCode
	N2023 *BinanceErrorCode
	N2024 *BinanceErrorCode
	N2025 *BinanceErrorCode
	N2026 *BinanceErrorCode
	N2027 *BinanceErrorCode
	N2028 *BinanceErrorCode
	N4000 *BinanceErrorCode
	N4001 *BinanceErrorCode
	N4002 *BinanceErrorCode
	N4003 *BinanceErrorCode
	N4004 *BinanceErrorCode
	N4005 *BinanceErrorCode
	N4006 *BinanceErrorCode
	N4007 *BinanceErrorCode
	N4008 *BinanceErrorCode
	N4009 *BinanceErrorCode
	N4010 *BinanceErrorCode
	N4011 *BinanceErrorCode
	N4012 *BinanceErrorCode
	N4013 *BinanceErrorCode
	N4014 *BinanceErrorCode
	N4015 *BinanceErrorCode
	N4016 *BinanceErrorCode
	N4017 *BinanceErrorCode
	N4018 *BinanceErrorCode
	N4019 *BinanceErrorCode
	N4020 *BinanceErrorCode
	N4021 *BinanceErrorCode
	N4022 *BinanceErrorCode
	N4023 *BinanceErrorCode
	N4024 *BinanceErrorCode
	N4025 *BinanceErrorCode
	N4026 *BinanceErrorCode
	N4027 *BinanceErrorCode
	N4028 *BinanceErrorCode
	N4029 *BinanceErrorCode
	N4030 *BinanceErrorCode
	N4031 *BinanceErrorCode
	N4032 *BinanceErrorCode
	N4033 *BinanceErrorCode
	N4044 *BinanceErrorCode
	N4045 *BinanceErrorCode
	N4046 *BinanceErrorCode
	N4047 *BinanceErrorCode
	N4048 *BinanceErrorCode
	N4049 *BinanceErrorCode
	N4050 *BinanceErrorCode
	N4051 *BinanceErrorCode
	N4052 *BinanceErrorCode
	N4053 *BinanceErrorCode
	N4054 *BinanceErrorCode
	N4055 *BinanceErrorCode
	N4056 *BinanceErrorCode
	N4057 *BinanceErrorCode
	N4058 *BinanceErrorCode
	N4059 *BinanceErrorCode
	N4060 *BinanceErrorCode
	N4061 *BinanceErrorCode
	N4062 *BinanceErrorCode
	N4063 *BinanceErrorCode
	N4064 *BinanceErrorCode
	N4065 *BinanceErrorCode
	N4066 *BinanceErrorCode
	N4067 *BinanceErrorCode
	N4068 *BinanceErrorCode
	N4069 *BinanceErrorCode
	N4070 *BinanceErrorCode
	N4071 *BinanceErrorCode
	N4072 *BinanceErrorCode
	N4073 *BinanceErrorCode
	N4074 *BinanceErrorCode
	N4075 *BinanceErrorCode
	N4076 *BinanceErrorCode
	N4077 *BinanceErrorCode
	N4078 *BinanceErrorCode
	N4079 *BinanceErrorCode
	N4080 *BinanceErrorCode
	N4081 *BinanceErrorCode
	N4082 *BinanceErrorCode
	N4083 *BinanceErrorCode
	N4084 *BinanceErrorCode
	N4085 *BinanceErrorCode
	N4086 *BinanceErrorCode
	N4087 *BinanceErrorCode
	N4088 *BinanceErrorCode
	N4104 *BinanceErrorCode
	N4114 *BinanceErrorCode
	N4115 *BinanceErrorCode
	N4116 *BinanceErrorCode
	N4117 *BinanceErrorCode
	N4118 *BinanceErrorCode
	N4131 *BinanceErrorCode
	N4135 *BinanceErrorCode
	N4137 *BinanceErrorCode
	N4138 *BinanceErrorCode
	N4139 *BinanceErrorCode
	N4140 *BinanceErrorCode
	N4141 *BinanceErrorCode
	N4142 *BinanceErrorCode
	N4144 *BinanceErrorCode
	N4161 *BinanceErrorCode
	N4164 *BinanceErrorCode
	N4165 *BinanceErrorCode
	N4167 *BinanceErrorCode
	N4168 *BinanceErrorCode
	N4169 *BinanceErrorCode
	N4170 *BinanceErrorCode
	N4171 *BinanceErrorCode
	N4172 *BinanceErrorCode
	N4183 *BinanceErrorCode
	N4184 *BinanceErrorCode
	N4192 *BinanceErrorCode
	N4202 *BinanceErrorCode
	N4203 *BinanceErrorCode
	N4205 *BinanceErrorCode
	N4206 *BinanceErrorCode
	N4208 *BinanceErrorCode
	N4209 *BinanceErrorCode
	N4210 *BinanceErrorCode
	N4211 *BinanceErrorCode
	N4400 *BinanceErrorCode
	N4401 *BinanceErrorCode
	N4402 *BinanceErrorCode
	N4403 *BinanceErrorCode
	N5021 *BinanceErrorCode
	N5022 *BinanceErrorCode
	N5024 *BinanceErrorCode
	N5025 *BinanceErrorCode
	N5026 *BinanceErrorCode
	N5027 *BinanceErrorCode
	N5028 *BinanceErrorCode
	N5029 *BinanceErrorCode
	N5037 *BinanceErrorCode
	N5038 *BinanceErrorCode
	N5039 *BinanceErrorCode
	N5040 *BinanceErrorCode
	N5041 *BinanceErrorCode
}

var (
	FuturesErr_UNKNOWN                                       = BinanceErrorCode{Code: -1000, Name: "UNKNOWN", Descriptions: []string{"An unknown error occured while processing the request."}}
	FuturesErr_DISCONNECTED                                  = BinanceErrorCode{Code: -1001, Name: "DISCONNECTED", Descriptions: []string{"Internal error; unable to process your request. Please try again."}}
	FuturesErr_UNAUTHORIZED                                  = BinanceErrorCode{Code: -1002, Name: "UNAUTHORIZED", Descriptions: []string{"You are not authorized to execute this request."}}
	FuturesErr_TOO_MANY_REQUESTS                             = BinanceErrorCode{Code: -1003, Name: "TOO_MANY_REQUESTS", Descriptions: []string{"Too many requests; current limit is %s requests per minute. Please use the websocket for live updates to avoid polling the API.", "Way too many requests; IP banned until %s. Please use the websocket for live updates to avoid bans."}}
	FuturesErr_DUPLICATE_IP                                  = BinanceErrorCode{Code: -1004, Name: "DUPLICATE_IP", Descriptions: []string{"This IP is already on the white list"}}
	FuturesErr_NO_SUCH_IP                                    = BinanceErrorCode{Code: -1005, Name: "NO_SUCH_IP", Descriptions: []string{"No such IP has been white listed"}}
	FuturesErr_UNEXPECTED_RESP                               = BinanceErrorCode{Code: -1006, Name: "UNEXPECTED_RESP", Descriptions: []string{"An unexpected response was received from the message bus. Execution status unknown."}}
	FuturesErr_TIMEOUT                                       = BinanceErrorCode{Code: -1007, Name: "TIMEOUT", Descriptions: []string{"Timeout waiting for response from backend server. Send status unknown; execution status unknown."}}
	FuturesErr_Service                                       = BinanceErrorCode{Code: -1008, Name: "Service", Descriptions: []string{"Server is currently overloaded with other requests. Please try again in a few minutes."}}
	FuturesErr_ERROR_MSG_RECEIVED                            = BinanceErrorCode{Code: -1010, Name: "ERROR_MSG_RECEIVED", Descriptions: []string{"ERROR_MSG_RECEIVED."}}
	FuturesErr_NON_WHITE_LIST                                = BinanceErrorCode{Code: -1011, Name: "NON_WHITE_LIST", Descriptions: []string{"This IP cannot access this route."}}
	FuturesErr_INVALID_MESSAGE                               = BinanceErrorCode{Code: -1013, Name: "INVALID_MESSAGE", Descriptions: []string{"INVALID_MESSAGE."}}
	FuturesErr_UNKNOWN_ORDER_COMPOSITION                     = BinanceErrorCode{Code: -1014, Name: "UNKNOWN_ORDER_COMPOSITION", Descriptions: []string{"Unsupported order combination."}}
	FuturesErr_TOO_MANY_ORDERS                               = BinanceErrorCode{Code: -1015, Name: "TOO_MANY_ORDERS", Descriptions: []string{"Too many new orders.", "Too many new orders; current limit is %s orders per %s."}}
	FuturesErr_SERVICE_SHUTTING_DOWN                         = BinanceErrorCode{Code: -1016, Name: "SERVICE_SHUTTING_DOWN", Descriptions: []string{"This service is no longer available."}}
	FuturesErr_UNSUPPORTED_OPERATION                         = BinanceErrorCode{Code: -1020, Name: "UNSUPPORTED_OPERATION", Descriptions: []string{"This operation is not supported."}}
	FuturesErr_INVALID_TIMESTAMP                             = BinanceErrorCode{Code: -1021, Name: "INVALID_TIMESTAMP", Descriptions: []string{"Timestamp for this request is outside of the recvWindow.", "Timestamp for this request was 1000ms ahead of the server's time."}}
	FuturesErr_INVALID_SIGNATURE                             = BinanceErrorCode{Code: -1022, Name: "INVALID_SIGNATURE", Descriptions: []string{"Signature for this request is not valid."}}
	FuturesErr_START_TIME_GREATER_THAN_END_TIME              = BinanceErrorCode{Code: -1023, Name: "START_TIME_GREATER_THAN_END_TIME", Descriptions: []string{"Start time is greater than end time."}}
	FuturesErr_NOT_FOUND                                     = BinanceErrorCode{Code: -1099, Name: "NOT_FOUND", Descriptions: []string{"Not found, unauthenticated, or unauthorized."}}
	FuturesErr_ILLEGAL_CHARS                                 = BinanceErrorCode{Code: -1100, Name: "ILLEGAL_CHARS", Descriptions: []string{"Illegal characters found in a parameter.", "Illegal characters found in parameter '%s'; legal range is '%s'."}}
	FuturesErr_TOO_MANY_PARAMETERS                           = BinanceErrorCode{Code: -1101, Name: "TOO_MANY_PARAMETERS", Descriptions: []string{"Too many parameters sent for this endpoint.", "Too many parameters; expected '%s' and received '%s'.", "Duplicate values for a parameter detected."}}
	FuturesErr_MANDATORY_PARAM_EMPTY_OR_MALFORMED            = BinanceErrorCode{Code: -1102, Name: "MANDATORY_PARAM_EMPTY_OR_MALFORMED", Descriptions: []string{"A mandatory parameter was not sent, was empty/null, or malformed.", "Mandatory parameter '%s' was not sent, was empty/null, or malformed.", "Param '%s' or '%s' must be sent, but both were empty/null!"}}
	FuturesErr_UNKNOWN_PARAM                                 = BinanceErrorCode{Code: -1103, Name: "UNKNOWN_PARAM", Descriptions: []string{"An unknown parameter was sent."}}
	FuturesErr_UNREAD_PARAMETERS                             = BinanceErrorCode{Code: -1104, Name: "UNREAD_PARAMETERS", Descriptions: []string{"Not all sent parameters were read.", "Not all sent parameters were read; read '%s' parameter(s) but was sent '%s'."}}
	FuturesErr_PARAM_EMPTY                                   = BinanceErrorCode{Code: -1105, Name: "PARAM_EMPTY", Descriptions: []string{"A parameter was empty.", "Parameter '%s' was empty."}}
	FuturesErr_PARAM_NOT_REQUIRED                            = BinanceErrorCode{Code: -1106, Name: "PARAM_NOT_REQUIRED", Descriptions: []string{"A parameter was sent when not required.", "Parameter '%s' sent when not required."}}
	FuturesErr_BAD_ASSET                                     = BinanceErrorCode{Code: -1108, Name: "BAD_ASSET", Descriptions: []string{"Invalid asset."}}
	FuturesErr_BAD_ACCOUNT                                   = BinanceErrorCode{Code: -1109, Name: "BAD_ACCOUNT", Descriptions: []string{"Invalid account."}}
	FuturesErr_BAD_INSTRUMENT_TYPE                           = BinanceErrorCode{Code: -1110, Name: "BAD_INSTRUMENT_TYPE", Descriptions: []string{"Invalid symbolType."}}
	FuturesErr_BAD_PRECISION                                 = BinanceErrorCode{Code: -1111, Name: "BAD_PRECISION", Descriptions: []string{"Precision is over the maximum defined for this asset."}}
	FuturesErr_NO_DEPTH                                      = BinanceErrorCode{Code: -1112, Name: "NO_DEPTH", Descriptions: []string{"No orders on book for symbol."}}
	FuturesErr_WITHDRAW_NOT_NEGATIVE                         = BinanceErrorCode{Code: -1113, Name: "WITHDRAW_NOT_NEGATIVE", Descriptions: []string{"Withdrawal amount must be negative."}}
	FuturesErr_TIF_NOT_REQUIRED                              = BinanceErrorCode{Code: -1114, Name: "TIF_NOT_REQUIRED", Descriptions: []string{"TimeInForce parameter sent when not required."}}
	FuturesErr_INVALID_TIF                                   = BinanceErrorCode{Code: -1115, Name: "INVALID_TIF", Descriptions: []string{"Invalid timeInForce."}}
	FuturesErr_INVALID_ORDER_TYPE                            = BinanceErrorCode{Code: -1116, Name: "INVALID_ORDER_TYPE", Descriptions: []string{"Invalid orderType."}}
	FuturesErr_INVALID_SIDE                                  = BinanceErrorCode{Code: -1117, Name: "INVALID_SIDE", Descriptions: []string{"Invalid side."}}
	FuturesErr_EMPTY_NEW_CL_ORD_ID                           = BinanceErrorCode{Code: -1118, Name: "EMPTY_NEW_CL_ORD_ID", Descriptions: []string{"New client order ID was empty."}}
	FuturesErr_EMPTY_ORG_CL_ORD_ID                           = BinanceErrorCode{Code: -1119, Name: "EMPTY_ORG_CL_ORD_ID", Descriptions: []string{"Original client order ID was empty."}}
	FuturesErr_BAD_INTERVAL                                  = BinanceErrorCode{Code: -1120, Name: "BAD_INTERVAL", Descriptions: []string{"Invalid interval."}}
	FuturesErr_BAD_SYMBOL                                    = BinanceErrorCode{Code: -1121, Name: "BAD_SYMBOL", Descriptions: []string{"Invalid symbol."}}
	FuturesErr_INVALID_SYMBOL_STATUS                         = BinanceErrorCode{Code: -1122, Name: "INVALID_SYMBOL_STATUS", Descriptions: []string{"Invalid symbol status."}}
	FuturesErr_INVALID_LISTEN_KEY                            = BinanceErrorCode{Code: -1125, Name: "INVALID_LISTEN_KEY", Descriptions: []string{"This listenKey does not exist. Please use POST /fapi/v1/listenKey to recreate listenKey"}}
	FuturesErr_ASSET_NOT_SUPPORTED                           = BinanceErrorCode{Code: -1126, Name: "ASSET_NOT_SUPPORTED", Descriptions: []string{"This asset is not supported."}}
	FuturesErr_MORE_THAN_XX_HOURS                            = BinanceErrorCode{Code: -1127, Name: "MORE_THAN_XX_HOURS", Descriptions: []string{"Lookup interval is too big.", "More than %s hours between startTime and endTime."}}
	FuturesErr_OPTIONAL_PARAMS_BAD_COMBO                     = BinanceErrorCode{Code: -1128, Name: "OPTIONAL_PARAMS_BAD_COMBO", Descriptions: []string{"Combination of optional parameters invalid."}}
	FuturesErr_INVALID_PARAMETER                             = BinanceErrorCode{Code: -1130, Name: "INVALID_PARAMETER", Descriptions: []string{"Invalid data sent for a parameter.", "Data sent for parameter '%s' is not valid."}}
	FuturesErr_INVALID_NEW_ORDER_RESP_TYPE                   = BinanceErrorCode{Code: -1136, Name: "INVALID_NEW_ORDER_RESP_TYPE", Descriptions: []string{"Invalid newOrderRespType."}}
	FuturesErr_NEW_ORDER_REJECTED                            = BinanceErrorCode{Code: -2010, Name: "NEW_ORDER_REJECTED", Descriptions: []string{"NEW_ORDER_REJECTED"}}
	FuturesErr_CANCEL_REJECTED                               = BinanceErrorCode{Code: -2011, Name: "CANCEL_REJECTED", Descriptions: []string{"CANCEL_REJECTED"}}
	FuturesErr_CANCEL_ALL_FAIL                               = BinanceErrorCode{Code: -2012, Name: "CANCEL_ALL_FAIL", Descriptions: []string{"Batch cancel failure."}}
	FuturesErr_NO_SUCH_ORDER                                 = BinanceErrorCode{Code: -2013, Name: "NO_SUCH_ORDER", Descriptions: []string{"Order does not exist."}}
	FuturesErr_BAD_API_KEY_FMT                               = BinanceErrorCode{Code: -2014, Name: "BAD_API_KEY_FMT", Descriptions: []string{"API-key format invalid."}}
	FuturesErr_REJECTED_MBX_KEY                              = BinanceErrorCode{Code: -2015, Name: "REJECTED_MBX_KEY", Descriptions: []string{"Invalid API-key, IP, or permissions for action."}}
	FuturesErr_NO_TRADING_WINDOW                             = BinanceErrorCode{Code: -2016, Name: "NO_TRADING_WINDOW", Descriptions: []string{"No trading window could be found for the symbol. Try ticker/24hrs instead."}}
	FuturesErr_API_KEYS_LOCKED                               = BinanceErrorCode{Code: -2017, Name: "API_KEYS_LOCKED", Descriptions: []string{"API Keys are locked on this account."}}
	FuturesErr_BALANCE_NOT_SUFFICIENT                        = BinanceErrorCode{Code: -2018, Name: "BALANCE_NOT_SUFFICIENT", Descriptions: []string{"Balance is insufficient."}}
	FuturesErr_MARGIN_NOT_SUFFICIEN                          = BinanceErrorCode{Code: -2019, Name: "MARGIN_NOT_SUFFICIEN", Descriptions: []string{"Margin is insufficient."}}
	FuturesErr_UNABLE_TO_FILL                                = BinanceErrorCode{Code: -2020, Name: "UNABLE_TO_FILL", Descriptions: []string{"Unable to fill."}}
	FuturesErr_ORDER_WOULD_IMMEDIATELY_TRIGGER               = BinanceErrorCode{Code: -2021, Name: "ORDER_WOULD_IMMEDIATELY_TRIGGER", Descriptions: []string{"Order would immediately trigger."}}
	FuturesErr_REDUCE_ONLY_REJECT                            = BinanceErrorCode{Code: -2022, Name: "REDUCE_ONLY_REJECT", Descriptions: []string{"ReduceOnly Order is rejected."}}
	FuturesErr_USER_IN_LIQUIDATION                           = BinanceErrorCode{Code: -2023, Name: "USER_IN_LIQUIDATION", Descriptions: []string{"User in liquidation mode now."}}
	FuturesErr_POSITION_NOT_SUFFICIENT                       = BinanceErrorCode{Code: -2024, Name: "POSITION_NOT_SUFFICIENT", Descriptions: []string{"Position is not sufficient."}}
	FuturesErr_MAX_OPEN_ORDER_EXCEEDED                       = BinanceErrorCode{Code: -2025, Name: "MAX_OPEN_ORDER_EXCEEDED", Descriptions: []string{"Reach max open order limit."}}
	FuturesErr_REDUCE_ONLY_ORDER_TYPE_NOT_SUPPORTED          = BinanceErrorCode{Code: -2026, Name: "REDUCE_ONLY_ORDER_TYPE_NOT_SUPPORTED", Descriptions: []string{"This OrderType is not supported when reduceOnly."}}
	FuturesErr_MAX_LEVERAGE_RATIO                            = BinanceErrorCode{Code: -2027, Name: "MAX_LEVERAGE_RATIO", Descriptions: []string{"Exceeded the maximum allowable position at current leverage."}}
	FuturesErr_MIN_LEVERAGE_RATIO                            = BinanceErrorCode{Code: -2028, Name: "MIN_LEVERAGE_RATIO", Descriptions: []string{"Leverage is smaller than permitted: insufficient margin balance."}}
	FuturesErr_INVALID_ORDER_STATUS                          = BinanceErrorCode{Code: -4000, Name: "INVALID_ORDER_STATUS", Descriptions: []string{"Invalid order status."}}
	FuturesErr_PRICE_LESS_THAN_ZERO                          = BinanceErrorCode{Code: -4001, Name: "PRICE_LESS_THAN_ZERO", Descriptions: []string{"Price less than 0."}}
	FuturesErr_PRICE_GREATER_THAN_MAX_PRICE                  = BinanceErrorCode{Code: -4002, Name: "PRICE_GREATER_THAN_MAX_PRICE", Descriptions: []string{"Price greater than max price."}}
	FuturesErr_QTY_LESS_THAN_ZERO                            = BinanceErrorCode{Code: -4003, Name: "QTY_LESS_THAN_ZERO", Descriptions: []string{"Quantity less than zero."}}
	FuturesErr_QTY_LESS_THAN_MIN_QTY                         = BinanceErrorCode{Code: -4004, Name: "QTY_LESS_THAN_MIN_QTY", Descriptions: []string{"Quantity less than min quantity."}}
	FuturesErr_QTY_GREATER_THAN_MAX_QTY                      = BinanceErrorCode{Code: -4005, Name: "QTY_GREATER_THAN_MAX_QTY", Descriptions: []string{"Quantity greater than max quantity."}}
	FuturesErr_STOP_PRICE_LESS_THAN_ZERO                     = BinanceErrorCode{Code: -4006, Name: "STOP_PRICE_LESS_THAN_ZERO", Descriptions: []string{"Stop price less than zero."}}
	FuturesErr_STOP_PRICE_GREATER_THAN_MAX_PRICE             = BinanceErrorCode{Code: -4007, Name: "STOP_PRICE_GREATER_THAN_MAX_PRICE", Descriptions: []string{"Stop price greater than max price."}}
	FuturesErr_TICK_SIZE_LESS_THAN_ZERO                      = BinanceErrorCode{Code: -4008, Name: "TICK_SIZE_LESS_THAN_ZERO", Descriptions: []string{"Tick size less than zero."}}
	FuturesErr_MAX_PRICE_LESS_THAN_MIN_PRICE                 = BinanceErrorCode{Code: -4009, Name: "MAX_PRICE_LESS_THAN_MIN_PRICE", Descriptions: []string{"Max price less than min price."}}
	FuturesErr_MAX_QTY_LESS_THAN_MIN_QTY                     = BinanceErrorCode{Code: -4010, Name: "MAX_QTY_LESS_THAN_MIN_QTY", Descriptions: []string{"Max qty less than min qty."}}
	FuturesErr_STEP_SIZE_LESS_THAN_ZERO                      = BinanceErrorCode{Code: -4011, Name: "STEP_SIZE_LESS_THAN_ZERO", Descriptions: []string{"Step size less than zero."}}
	FuturesErr_MAX_NUM_ORDERS_LESS_THAN_ZERO                 = BinanceErrorCode{Code: -4012, Name: "MAX_NUM_ORDERS_LESS_THAN_ZERO", Descriptions: []string{"Max mum orders less than zero."}}
	FuturesErr_PRICE_LESS_THAN_MIN_PRICE                     = BinanceErrorCode{Code: -4013, Name: "PRICE_LESS_THAN_MIN_PRICE", Descriptions: []string{"Price less than min price."}}
	FuturesErr_PRICE_NOT_INCREASED_BY_TICK_SIZE              = BinanceErrorCode{Code: -4014, Name: "PRICE_NOT_INCREASED_BY_TICK_SIZE", Descriptions: []string{"Price not increased by tick size."}}
	FuturesErr_INVALID_CL_ORD_ID_LEN                         = BinanceErrorCode{Code: -4015, Name: "INVALID_CL_ORD_ID_LEN", Descriptions: []string{"Client order id is not valid.", "Client order id length should not be more than 36 chars"}}
	FuturesErr_PRICE_HIGHTER_THAN_MULTIPLIER_UP              = BinanceErrorCode{Code: -4016, Name: "PRICE_HIGHTER_THAN_MULTIPLIER_UP", Descriptions: []string{"Price is higher than mark price multiplier cap."}}
	FuturesErr_MULTIPLIER_UP_LESS_THAN_ZERO                  = BinanceErrorCode{Code: -4017, Name: "MULTIPLIER_UP_LESS_THAN_ZERO", Descriptions: []string{"Multiplier up less than zero."}}
	FuturesErr_MULTIPLIER_DOWN_LESS_THAN_ZERO                = BinanceErrorCode{Code: -4018, Name: "MULTIPLIER_DOWN_LESS_THAN_ZERO", Descriptions: []string{"Multiplier down less than zero."}}
	FuturesErr_COMPOSITE_SCALE_OVERFLOW                      = BinanceErrorCode{Code: -4019, Name: "COMPOSITE_SCALE_OVERFLOW", Descriptions: []string{"Composite scale too large."}}
	FuturesErr_TARGET_STRATEGY_INVALID                       = BinanceErrorCode{Code: -4020, Name: "TARGET_STRATEGY_INVALID", Descriptions: []string{"Target strategy invalid for orderType '%s',reduceOnly '%b'."}}
	FuturesErr_INVALID_DEPTH_LIMIT                           = BinanceErrorCode{Code: -4021, Name: "INVALID_DEPTH_LIMIT", Descriptions: []string{"Invalid depth limit.", "'%s' is not valid depth limit."}}
	FuturesErr_WRONG_MARKET_STATUS                           = BinanceErrorCode{Code: -4022, Name: "WRONG_MARKET_STATUS", Descriptions: []string{"market status sent is not valid."}}
	FuturesErr_QTY_NOT_INCREASED_BY_STEP_SIZE                = BinanceErrorCode{Code: -4023, Name: "QTY_NOT_INCREASED_BY_STEP_SIZE", Descriptions: []string{"Qty not increased by step size."}}
	FuturesErr_PRICE_LOWER_THAN_MULTIPLIER_DOWN              = BinanceErrorCode{Code: -4024, Name: "PRICE_LOWER_THAN_MULTIPLIER_DOWN", Descriptions: []string{"Price is lower than mark price multiplier floor."}}
	FuturesErr_MULTIPLIER_DECIMAL_LESS_THAN_ZERO             = BinanceErrorCode{Code: -4025, Name: "MULTIPLIER_DECIMAL_LESS_THAN_ZERO", Descriptions: []string{"Multiplier decimal less than zero."}}
	FuturesErr_COMMISSION_INVALID                            = BinanceErrorCode{Code: -4026, Name: "COMMISSION_INVALID", Descriptions: []string{"Commission invalid.", "%s less than zero.", "%s absolute value greater than %s"}}
	FuturesErr_INVALID_ACCOUNT_TYPE                          = BinanceErrorCode{Code: -4027, Name: "INVALID_ACCOUNT_TYPE", Descriptions: []string{"Invalid account type."}}
	FuturesErr_INVALID_LEVERAGE                              = BinanceErrorCode{Code: -4028, Name: "INVALID_LEVERAGE", Descriptions: []string{"Invalid leverage", "Leverage %s is not valid", "Leverage %s already exist with %s"}}
	FuturesErr_INVALID_TICK_SIZE_PRECISION                   = BinanceErrorCode{Code: -4029, Name: "INVALID_TICK_SIZE_PRECISION", Descriptions: []string{"Tick size precision is invalid."}}
	FuturesErr_INVALID_STEP_SIZE_PRECISION                   = BinanceErrorCode{Code: -4030, Name: "INVALID_STEP_SIZE_PRECISION", Descriptions: []string{"Step size precision is invalid."}}
	FuturesErr_INVALID_WORKING_TYPE                          = BinanceErrorCode{Code: -4031, Name: "INVALID_WORKING_TYPE", Descriptions: []string{"Invalid parameter working type", "Invalid parameter working type: %s"}}
	FuturesErr_EXCEED_MAX_CANCEL_ORDER_SIZE                  = BinanceErrorCode{Code: -4032, Name: "EXCEED_MAX_CANCEL_ORDER_SIZE", Descriptions: []string{"Exceed maximum cancel order size.", "Invalid parameter working type: %s"}}
	FuturesErr_INSURANCE_ACCOUNT_NOT_FOUND                   = BinanceErrorCode{Code: -4033, Name: "INSURANCE_ACCOUNT_NOT_FOUND", Descriptions: []string{"Insurance account not found."}}
	FuturesErr_INVALID_BALANCE_TYPE                          = BinanceErrorCode{Code: -4044, Name: "INVALID_BALANCE_TYPE", Descriptions: []string{"Balance Type is invalid."}}
	FuturesErr_MAX_STOP_ORDER_EXCEEDED                       = BinanceErrorCode{Code: -4045, Name: "MAX_STOP_ORDER_EXCEEDED", Descriptions: []string{"Reach max stop order limit."}}
	FuturesErr_NO_NEED_TO_CHANGE_MARGIN_TYPE                 = BinanceErrorCode{Code: -4046, Name: "NO_NEED_TO_CHANGE_MARGIN_TYPE", Descriptions: []string{"No need to change margin type."}}
	FuturesErr_THERE_EXISTS_OPEN_ORDERS                      = BinanceErrorCode{Code: -4047, Name: "THERE_EXISTS_OPEN_ORDERS", Descriptions: []string{"Margin type cannot be changed if there exists open orders."}}
	FuturesErr_THERE_EXISTS_QUANTITY                         = BinanceErrorCode{Code: -4048, Name: "THERE_EXISTS_QUANTITY", Descriptions: []string{"Margin type cannot be changed if there exists position."}}
	FuturesErr_ADD_ISOLATED_MARGIN_REJECT                    = BinanceErrorCode{Code: -4049, Name: "ADD_ISOLATED_MARGIN_REJECT", Descriptions: []string{"Add margin only support for isolated position."}}
	FuturesErr_CROSS_BALANCE_INSUFFICIENT                    = BinanceErrorCode{Code: -4050, Name: "CROSS_BALANCE_INSUFFICIENT", Descriptions: []string{"Cross balance insufficient."}}
	FuturesErr_ISOLATED_BALANCE_INSUFFICIENT                 = BinanceErrorCode{Code: -4051, Name: "ISOLATED_BALANCE_INSUFFICIENT", Descriptions: []string{"Isolated balance insufficient."}}
	FuturesErr_NO_NEED_TO_CHANGE_AUTO_ADD_MARGIN             = BinanceErrorCode{Code: -4052, Name: "NO_NEED_TO_CHANGE_AUTO_ADD_MARGIN", Descriptions: []string{"No need to change auto add margin."}}
	FuturesErr_AUTO_ADD_CROSSED_MARGIN_REJECT                = BinanceErrorCode{Code: -4053, Name: "AUTO_ADD_CROSSED_MARGIN_REJECT", Descriptions: []string{"Auto add margin only support for isolated position."}}
	FuturesErr_ADD_ISOLATED_MARGIN_NO_POSITION_REJECT        = BinanceErrorCode{Code: -4054, Name: "ADD_ISOLATED_MARGIN_NO_POSITION_REJECT", Descriptions: []string{"Cannot add position margin: position is 0."}}
	FuturesErr_AMOUNT_MUST_BE_POSITIVE                       = BinanceErrorCode{Code: -4055, Name: "AMOUNT_MUST_BE_POSITIVE", Descriptions: []string{"Amount must be positive."}}
	FuturesErr_INVALID_API_KEY_TYPE                          = BinanceErrorCode{Code: -4056, Name: "INVALID_API_KEY_TYPE", Descriptions: []string{"Invalid api key type."}}
	FuturesErr_INVALID_RSA_PUBLIC_KEY                        = BinanceErrorCode{Code: -4057, Name: "INVALID_RSA_PUBLIC_KEY", Descriptions: []string{"Invalid api public key"}}
	FuturesErr_MAX_PRICE_TOO_LARGE                           = BinanceErrorCode{Code: -4058, Name: "MAX_PRICE_TOO_LARGE", Descriptions: []string{"maxPrice and priceDecimal too large,please check."}}
	FuturesErr_NO_NEED_TO_CHANGE_POSITION_SIDE               = BinanceErrorCode{Code: -4059, Name: "NO_NEED_TO_CHANGE_POSITION_SIDE", Descriptions: []string{"No need to change position side."}}
	FuturesErr_INVALID_POSITION_SIDE                         = BinanceErrorCode{Code: -4060, Name: "INVALID_POSITION_SIDE", Descriptions: []string{"Invalid position side."}}
	FuturesErr_POSITION_SIDE_NOT_MATCH                       = BinanceErrorCode{Code: -4061, Name: "POSITION_SIDE_NOT_MATCH", Descriptions: []string{"Order's position side does not match user's setting."}}
	FuturesErr_REDUCE_ONLY_CONFLICT                          = BinanceErrorCode{Code: -4062, Name: "REDUCE_ONLY_CONFLICT", Descriptions: []string{"Invalid or improper reduceOnly value."}}
	FuturesErr_INVALID_OPTIONS_REQUEST_TYPE                  = BinanceErrorCode{Code: -4063, Name: "INVALID_OPTIONS_REQUEST_TYPE", Descriptions: []string{"Invalid options request type"}}
	FuturesErr_INVALID_OPTIONS_TIME_FRAME                    = BinanceErrorCode{Code: -4064, Name: "INVALID_OPTIONS_TIME_FRAME", Descriptions: []string{"Invalid options time frame"}}
	FuturesErr_INVALID_OPTIONS_AMOUNT                        = BinanceErrorCode{Code: -4065, Name: "INVALID_OPTIONS_AMOUNT", Descriptions: []string{"Invalid options amount"}}
	FuturesErr_INVALID_OPTIONS_EVENT_TYPE                    = BinanceErrorCode{Code: -4066, Name: "INVALID_OPTIONS_EVENT_TYPE", Descriptions: []string{"Invalid options event type"}}
	FuturesErr_POSITION_SIDE_CHANGE_EXISTS_OPEN_ORDERS       = BinanceErrorCode{Code: -4067, Name: "POSITION_SIDE_CHANGE_EXISTS_OPEN_ORDERS", Descriptions: []string{"Position side cannot be changed if there exists open orders."}}
	FuturesErr_POSITION_SIDE_CHANGE_EXISTS_QUANTITY          = BinanceErrorCode{Code: -4068, Name: "POSITION_SIDE_CHANGE_EXISTS_QUANTITY", Descriptions: []string{"Position side cannot be changed if there exists position."}}
	FuturesErr_INVALID_OPTIONS_PREMIUM_FEE                   = BinanceErrorCode{Code: -4069, Name: "INVALID_OPTIONS_PREMIUM_FEE", Descriptions: []string{"Invalid options premium fee"}}
	FuturesErr_INVALID_CL_OPTIONS_ID_LEN                     = BinanceErrorCode{Code: -4070, Name: "INVALID_CL_OPTIONS_ID_LEN", Descriptions: []string{"Client options id is not valid.", "Client options id length should be less than 32 chars"}}
	FuturesErr_INVALID_OPTIONS_DIRECTION                     = BinanceErrorCode{Code: -4071, Name: "INVALID_OPTIONS_DIRECTION", Descriptions: []string{"Invalid options direction"}}
	FuturesErr_OPTIONS_PREMIUM_NOT_UPDATE                    = BinanceErrorCode{Code: -4072, Name: "OPTIONS_PREMIUM_NOT_UPDATE", Descriptions: []string{"premium fee is not updated, reject order"}}
	FuturesErr_OPTIONS_PREMIUM_INPUT_LESS_THAN_ZERO          = BinanceErrorCode{Code: -4073, Name: "OPTIONS_PREMIUM_INPUT_LESS_THAN_ZERO", Descriptions: []string{"input premium fee is less than 0, reject order"}}
	FuturesErr_OPTIONS_AMOUNT_BIGGER_THAN_UPPER              = BinanceErrorCode{Code: -4074, Name: "OPTIONS_AMOUNT_BIGGER_THAN_UPPER", Descriptions: []string{"Order amount is bigger than upper boundary or less than 0, reject order"}}
	FuturesErr_OPTIONS_PREMIUM_OUTPUT_ZERO                   = BinanceErrorCode{Code: -4075, Name: "OPTIONS_PREMIUM_OUTPUT_ZERO", Descriptions: []string{"output premium fee is less than 0, reject order"}}
	FuturesErr_OPTIONS_PREMIUM_TOO_DIFF                      = BinanceErrorCode{Code: -4076, Name: "OPTIONS_PREMIUM_TOO_DIFF", Descriptions: []string{"original fee is too much higher than last fee"}}
	FuturesErr_OPTIONS_PREMIUM_REACH_LIMIT                   = BinanceErrorCode{Code: -4077, Name: "OPTIONS_PREMIUM_REACH_LIMIT", Descriptions: []string{"place order amount has reached to limit, reject order"}}
	FuturesErr_OPTIONS_COMMON_ERROR                          = BinanceErrorCode{Code: -4078, Name: "OPTIONS_COMMON_ERROR", Descriptions: []string{"options internal error"}}
	FuturesErr_INVALID_OPTIONS_ID                            = BinanceErrorCode{Code: -4079, Name: "INVALID_OPTIONS_ID", Descriptions: []string{"invalid options id", "invalid options id: %s", "duplicate options id %d for user %d"}}
	FuturesErr_OPTIONS_USER_NOT_FOUND                        = BinanceErrorCode{Code: -4080, Name: "OPTIONS_USER_NOT_FOUND", Descriptions: []string{"user not found", "user not found with id: %s"}}
	FuturesErr_OPTIONS_NOT_FOUND                             = BinanceErrorCode{Code: -4081, Name: "OPTIONS_NOT_FOUND", Descriptions: []string{"options not found", "options not found with id: %s"}}
	FuturesErr_INVALID_BATCH_PLACE_ORDER_SIZE                = BinanceErrorCode{Code: -4082, Name: "INVALID_BATCH_PLACE_ORDER_SIZE", Descriptions: []string{"Invalid number of batch place orders.", "Invalid number of batch place orders: %s"}}
	FuturesErr_PLACE_BATCH_ORDERS_FAIL                       = BinanceErrorCode{Code: -4083, Name: "PLACE_BATCH_ORDERS_FAIL", Descriptions: []string{"Fail to place batch orders."}}
	FuturesErr_UPCOMING_METHOD                               = BinanceErrorCode{Code: -4084, Name: "UPCOMING_METHOD", Descriptions: []string{"Method is not allowed currently. Upcoming soon."}}
	FuturesErr_INVALID_NOTIONAL_LIMIT_COEF                   = BinanceErrorCode{Code: -4085, Name: "INVALID_NOTIONAL_LIMIT_COEF", Descriptions: []string{"Invalid notional limit coefficient"}}
	FuturesErr_INVALID_PRICE_SPREAD_THRESHOLD                = BinanceErrorCode{Code: -4086, Name: "INVALID_PRICE_SPREAD_THRESHOLD", Descriptions: []string{"Invalid price spread threshold"}}
	FuturesErr_REDUCE_ONLY_ORDER_PERMISSION                  = BinanceErrorCode{Code: -4087, Name: "REDUCE_ONLY_ORDER_PERMISSION", Descriptions: []string{"User can only place reduce only order"}}
	FuturesErr_NO_PLACE_ORDER_PERMISSION                     = BinanceErrorCode{Code: -4088, Name: "NO_PLACE_ORDER_PERMISSION", Descriptions: []string{"User can not place order currently"}}
	FuturesErr_INVALID_CONTRACT_TYPE                         = BinanceErrorCode{Code: -4104, Name: "INVALID_CONTRACT_TYPE", Descriptions: []string{"Invalid contract type"}}
	FuturesErr_INVALID_CLIENT_TRAN_ID_LEN                    = BinanceErrorCode{Code: -4114, Name: "INVALID_CLIENT_TRAN_ID_LEN", Descriptions: []string{"clientTranId is not valid", "Client tran id length should be less than 64 chars"}}
	FuturesErr_DUPLICATED_CLIENT_TRAN_ID                     = BinanceErrorCode{Code: -4115, Name: "DUPLICATED_CLIENT_TRAN_ID", Descriptions: []string{"clientTranId is duplicated", "Client tran id should be unique within 7 days"}}
	FuturesErr_DUPLICATED_CLIENT_ORDER_ID                    = BinanceErrorCode{Code: -4116, Name: "DUPLICATED_CLIENT_ORDER_ID", Descriptions: []string{"clientOrderId is duplicated"}}
	FuturesErr_STOP_ORDER_TRIGGERING                         = BinanceErrorCode{Code: -4117, Name: "STOP_ORDER_TRIGGERING", Descriptions: []string{"stop order is triggering"}}
	FuturesErr_REDUCE_ONLY_MARGIN_CHECK_FAILED               = BinanceErrorCode{Code: -4118, Name: "REDUCE_ONLY_MARGIN_CHECK_FAILED", Descriptions: []string{"ReduceOnly Order Failed. Please check your existing position and open orders"}}
	FuturesErr_MARKET_ORDER_REJECT                           = BinanceErrorCode{Code: -4131, Name: "MARKET_ORDER_REJECT", Descriptions: []string{"The counterparty's best price does not meet the PERCENT_PRICE filter limit"}}
	FuturesErr_INVALID_ACTIVATION_PRICE                      = BinanceErrorCode{Code: -4135, Name: "INVALID_ACTIVATION_PRICE", Descriptions: []string{"Invalid activation price"}}
	FuturesErr_QUANTITY_EXISTS_WITH_CLOSE_POSITION           = BinanceErrorCode{Code: -4137, Name: "QUANTITY_EXISTS_WITH_CLOSE_POSITION", Descriptions: []string{"Quantity must be zero with closePosition equals true"}}
	FuturesErr_REDUCE_ONLY_MUST_BE_TRUE                      = BinanceErrorCode{Code: -4138, Name: "REDUCE_ONLY_MUST_BE_TRUE", Descriptions: []string{"Reduce only must be true with closePosition equals true"}}
	FuturesErr_ORDER_TYPE_CANNOT_BE_MKT                      = BinanceErrorCode{Code: -4139, Name: "ORDER_TYPE_CANNOT_BE_MKT", Descriptions: []string{"Order type can not be market if it's unable to cancel"}}
	FuturesErr_INVALID_OPENING_POSITION_STATUS               = BinanceErrorCode{Code: -4140, Name: "INVALID_OPENING_POSITION_STATUS", Descriptions: []string{"Invalid symbol status for opening position"}}
	FuturesErr_SYMBOL_ALREADY_CLOSED                         = BinanceErrorCode{Code: -4141, Name: "SYMBOL_ALREADY_CLOSED", Descriptions: []string{"Symbol is closed"}}
	FuturesErr_STRATEGY_INVALID_TRIGGER_PRICE                = BinanceErrorCode{Code: -4142, Name: "STRATEGY_INVALID_TRIGGER_PRICE", Descriptions: []string{"REJECT: take profit or stop order will be triggered immediately"}}
	FuturesErr_INVALID_PAIR                                  = BinanceErrorCode{Code: -4144, Name: "INVALID_PAIR", Descriptions: []string{"Invalid pair"}}
	FuturesErr_ISOLATED_LEVERAGE_REJECT_WITH_POSITION        = BinanceErrorCode{Code: -4161, Name: "ISOLATED_LEVERAGE_REJECT_WITH_POSITION", Descriptions: []string{"Leverage reduction is not supported in Isolated Margin Mode with open positions"}}
	FuturesErr_MIN_NOTIONAL                                  = BinanceErrorCode{Code: -4164, Name: "MIN_NOTIONAL", Descriptions: []string{"Order's notional must be no smaller than 5.0 (unless you choose reduce only)", "Order's notional must be no smaller than %s (unless you choose reduce only)"}}
	FuturesErr_INVALID_TIME_INTERVAL                         = BinanceErrorCode{Code: -4165, Name: "INVALID_TIME_INTERVAL", Descriptions: []string{"Invalid time interval", "Maximum time interval is %s days"}}
	FuturesErr_ISOLATED_REJECT_WITH_JOINT_MARGIN             = BinanceErrorCode{Code: -4167, Name: "ISOLATED_REJECT_WITH_JOINT_MARGIN", Descriptions: []string{"Unable to adjust to Multi-Assets mode with symbols of USDⓈ-M Futures under isolated-margin mode."}}
	FuturesErr_JOINT_MARGIN_REJECT_WITH_ISOLATED             = BinanceErrorCode{Code: -4168, Name: "JOINT_MARGIN_REJECT_WITH_ISOLATED", Descriptions: []string{"Unable to adjust to isolated-margin mode under the Multi-Assets mode."}}
	FuturesErr_JOINT_MARGIN_REJECT_WITH_MB                   = BinanceErrorCode{Code: -4169, Name: "JOINT_MARGIN_REJECT_WITH_MB", Descriptions: []string{"Unable to adjust Multi-Assets Mode with insufficient margin balance in USDⓈ-M Futures."}}
	FuturesErr_JOINT_MARGIN_REJECT_WITH_OPEN_ORDER           = BinanceErrorCode{Code: -4170, Name: "JOINT_MARGIN_REJECT_WITH_OPEN_ORDER", Descriptions: []string{"Unable to adjust Multi-Assets Mode with open orders in USDⓈ-M Futures."}}
	FuturesErr_NO_NEED_TO_CHANGE_JOINT_MARGIN                = BinanceErrorCode{Code: -4171, Name: "NO_NEED_TO_CHANGE_JOINT_MARGIN", Descriptions: []string{"Adjusted asset mode is currently set and does not need to be adjusted repeatedly."}}
	FuturesErr_JOINT_MARGIN_REJECT_WITH_NEGATIVE_BALANCE     = BinanceErrorCode{Code: -4172, Name: "JOINT_MARGIN_REJECT_WITH_NEGATIVE_BALANCE", Descriptions: []string{"Unable to adjust Multi-Assets Mode with a negative wallet balance of margin available asset in USDⓈ-M Futures account."}}
	FuturesErr_ISOLATED_REJECT_WITH_JOINT_MARGIN_PRICE       = BinanceErrorCode{Code: -4183, Name: "ISOLATED_REJECT_WITH_JOINT_MARGIN", Descriptions: []string{"Price is higher than stop price multiplier cap.", "Limit price can't be higher than %s."}}
	FuturesErr_PRICE_LOWER_THAN_STOP_MULTIPLIER_DOWN         = BinanceErrorCode{Code: -4184, Name: "PRICE_LOWER_THAN_STOP_MULTIPLIER_DOWN", Descriptions: []string{"Price is lower than stop price multiplier floor.", "Limit price can't be lower than %s."}}
	FuturesErr_COOLING_OFF_PERIOD                            = BinanceErrorCode{Code: -4192, Name: "COOLING_OFF_PERIOD", Descriptions: []string{"Trade forbidden due to Cooling-off Period."}}
	FuturesErr_ADJUST_LEVERAGE_KYC_FAILED                    = BinanceErrorCode{Code: -4202, Name: "ADJUST_LEVERAGE_KYC_FAILED", Descriptions: []string{"Intermediate Personal Verification is required for adjusting leverage over 20x"}}
	FuturesErr_ADJUST_LEVERAGE_ONE_MONTH_FAILED              = BinanceErrorCode{Code: -4203, Name: "ADJUST_LEVERAGE_ONE_MONTH_FAILED", Descriptions: []string{"More than 20x leverage is available one month after account registration."}}
	FuturesErr_ADJUST_LEVERAGE_X_DAYS_FAILED                 = BinanceErrorCode{Code: -4205, Name: "ADJUST_LEVERAGE_X_DAYS_FAILED", Descriptions: []string{"More than 20x leverage is available %s days after Futures account registration."}}
	FuturesErr_ADJUST_LEVERAGE_KYC_LIMIT                     = BinanceErrorCode{Code: -4206, Name: "ADJUST_LEVERAGE_KYC_LIMIT", Descriptions: []string{"Users in this country has limited adjust leverage.", "Users in your location/country can only access a maximum leverage of %s"}}
	FuturesErr_ADJUST_LEVERAGE_ACCOUNT_SYMBOL_FAILED         = BinanceErrorCode{Code: -4208, Name: "ADJUST_LEVERAGE_ACCOUNT_SYMBOL_FAILED", Descriptions: []string{"Current symbol leverage cannot exceed 20 when using position limit adjustment service."}}
	FuturesErr_ADJUST_LEVERAGE_SYMBOL_FAILED                 = BinanceErrorCode{Code: -4209, Name: "ADJUST_LEVERAGE_SYMBOL_FAILED", Descriptions: []string{"The max leverage of Symbol is 20x", "Leverage adjustment failed. Current symbol max leverage limit is %sx"}}
	FuturesErr_STOP_PRICE_HIGHER_THAN_PRICE_MULTIPLIER_LIMIT = BinanceErrorCode{Code: -4210, Name: "STOP_PRICE_HIGHER_THAN_PRICE_MULTIPLIER_LIMIT", Descriptions: []string{"Stop price is higher than price multiplier cap.", "Stop price can't be higher than %s"}}
	FuturesErr_STOP_PRICE_LOWER_THAN_PRICE_MULTIPLIER_LIMIT  = BinanceErrorCode{Code: -4211, Name: "STOP_PRICE_LOWER_THAN_PRICE_MULTIPLIER_LIMIT", Descriptions: []string{"Stop price is lower than price multiplier floor.", "Stop price can't be lower than %s"}}
	FuturesErr_TRADING_QUANTITATIVE_RULE                     = BinanceErrorCode{Code: -4400, Name: "TRADING_QUANTITATIVE_RULE", Descriptions: []string{"Futures Trading Quantitative Rules violated, only reduceOnly order is allowed, please try again later."}}
	FuturesErr_LARGE_POSITION_SYM_RULE                       = BinanceErrorCode{Code: -4401, Name: "LARGE_POSITION_SYM_RULE", Descriptions: []string{"Futures Trading Risk Control Rules of large position holding violated, only reduceOnly order is allowed, please reduce the position. ."}}
	FuturesErr_COMPLIANCE_BLACK_SYMBOL_RESTRICTION           = BinanceErrorCode{Code: -4402, Name: "COMPLIANCE_BLACK_SYMBOL_RESTRICTION", Descriptions: []string{"Dear user, as per our Terms of Use and compliance with local regulations, this feature is currently not available in your region."}}
	FuturesErr_ADJUST_LEVERAGE_COMPLIANCE_FAILED             = BinanceErrorCode{Code: -4403, Name: "ADJUST_LEVERAGE_COMPLIANCE_FAILED", Descriptions: []string{"Dear user, as per our Terms of Use and compliance with local regulations, the leverage can only up to 10x in your region", "Dear user, as per our Terms of Use and compliance with local regulations, the leverage can only up to %sx in your region"}}
	FuturesErr_FOK_ORDER_REJECT                              = BinanceErrorCode{Code: -5021, Name: "FOK_ORDER_REJECT", Descriptions: []string{"Due to the order could not be filled immediately, the FOK order has been rejected."}}
	FuturesErr_GTX_ORDER_REJECT                              = BinanceErrorCode{Code: -5022, Name: "GTX_ORDER_REJECT", Descriptions: []string{"Due to the order could not be executed as maker, the Post Only order will be rejected."}}
	FuturesErr_MOVE_ORDER_NOT_ALLOWED_SYMBOL_REASON          = BinanceErrorCode{Code: -5024, Name: "MOVE_ORDER_NOT_ALLOWED_SYMBOL_REASON", Descriptions: []string{"Symbol is not in trading status. Order amendment is not permitted."}}
	FuturesErr_LIMIT_ORDER_ONLY                              = BinanceErrorCode{Code: -5025, Name: "LIMIT_ORDER_ONLY", Descriptions: []string{"Only limit order is supported."}}
	FuturesErr_Exceed_Maximum_Modify_Order_Limit             = BinanceErrorCode{Code: -5026, Name: "Exceed_Maximum_Modify_Order_Limit", Descriptions: []string{"Exceed maximum modify order limit."}}
	FuturesErr_SAME_ORDER                                    = BinanceErrorCode{Code: -5027, Name: "SAME_ORDER", Descriptions: []string{"No need to modify the order."}}
	FuturesErr_ME_RECVWINDOW_REJECT                          = BinanceErrorCode{Code: -5028, Name: "ME_RECVWINDOW_REJECT", Descriptions: []string{"Timestamp for this request is outside of the ME recvWindow."}}
	FuturesErr_MODIFICATION_MIN_NOTIONAL                     = BinanceErrorCode{Code: -5029, Name: "MODIFICATION_MIN_NOTIONAL", Descriptions: []string{"Order's notional must be no smaller than %s"}}
	FuturesErr_INVALID_PRICE_MATCH                           = BinanceErrorCode{Code: -5037, Name: "INVALID_PRICE_MATCH", Descriptions: []string{"Invalid price match"}}
	FuturesErr_UNSUPPORTED_ORDER_TYPE_PRICE_MATCH            = BinanceErrorCode{Code: -5038, Name: "UNSUPPORTED_ORDER_TYPE_PRICE_MATCH", Descriptions: []string{"Price match only supports order type: LIMIT, STOP AND TAKE_PROFIT"}}
	FuturesErr_INVALID_SELF_TRADE_PREVENTION_MODE            = BinanceErrorCode{Code: -5039, Name: "INVALID_SELF_TRADE_PREVENTION_MODE", Descriptions: []string{"Invalid self trade prevention mode"}}
	FuturesErr_FUTURE_GOOD_TILL_DATE                         = BinanceErrorCode{Code: -5040, Name: "FUTURE_GOOD_TILL_DATE", Descriptions: []string{"The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (UTC 9999-12-31 23:59:59)"}}
	FuturesErr_BBO_ORDER_REJECT                              = BinanceErrorCode{Code: -5041, Name: "BBO_ORDER_REJECT", Descriptions: []string{"No depth matches this BBO order"}}
)
