// generated by stringer -type=TokenType token.go; DO NOT EDIT

package ast

import "fmt"

const _TokenType_name = "T_SPACET_COMMENT_LINET_COMMENT_BLOCKT_SEMICOLONT_COMMAT_IDENTT_URLT_MEDIAT_TRUET_FALSET_NULLT_ONLYT_MS_PARAM_NAMET_FUNCTION_NAMET_ID_SELECTORT_CLASS_SELECTORT_TYPE_SELECTORT_UNIVERSAL_SELECTORT_PARENT_SELECTORT_PSEUDO_SELECTORT_FUNCTIONAL_PSEUDOT_INTERPOLATION_SELECTORT_LITERAL_CONCATT_CONCATT_MS_PROGIDT_AND_SELECTORT_DESCENDANT_COMBINATORT_CHILD_COMBINATORT_ADJACENT_SIBLING_COMBINATORT_GENERAL_SIBLING_COMBINATORT_UNICODE_RANGET_IFT_ELSET_ELSE_IFT_INCLUDET_MIXINT_FUNCTIONT_FORT_FOR_FROMT_FOR_THROUGHT_FOR_TOT_FOR_INT_WHILET_RETURNT_RANGET_GLOBALT_DEFAULTT_IMPORTANTT_OPTIONALT_FONT_FACET_LOGICAL_NOTT_LOGICAL_ORT_LOGICAL_ANDT_LOGICAL_XORT_NOPT_PLUST_DIVT_MULT_MINUST_MODT_BRACE_STARTT_BRACE_ENDT_LANG_CODET_BRACKET_LEFTT_ATTRIBUTE_NAMET_BRACKET_RIGHTT_EQUALT_UNEQUALT_GTT_LTT_GET_LET_ASSIGNT_ATTR_EQUALT_ATTR_TILDE_EQUALT_ATTR_HYPHEN_EQUALT_VARIABLET_IMPORTT_AT_RULET_CHARSETT_QQ_STRINGT_Q_STRINGT_UNQUOTE_STRINGT_PAREN_STARTT_PAREN_ENDT_CONSTANTT_INTEGERT_FLOATT_UNIT_NONET_UNIT_PERCENTT_UNIT_SECONDT_UNIT_MILLISECONDT_UNIT_EMT_UNIT_EXT_UNIT_CHT_UNIT_REMT_UNIT_CMT_UNIT_INT_UNIT_MMT_UNIT_PCT_UNIT_PTT_UNIT_PXT_UNIT_VHT_UNIT_VWT_UNIT_VMINT_UNIT_VMAXT_UNIT_HZT_UNIT_KHZT_UNIT_DPIT_UNIT_DPCMT_UNIT_DPPXT_UNIT_DEGT_UNIT_GRADT_UNIT_RADT_UNIT_TURNT_PROPERTY_NAME_TOKENT_PROPERTY_VALUET_HEX_COLORT_COLONT_INTERPOLATION_STARTT_INTERPOLATION_INNERT_INTERPOLATION_END"

var _TokenType_index = [...]uint16{0, 7, 21, 36, 47, 54, 61, 66, 73, 79, 86, 92, 98, 113, 128, 141, 157, 172, 192, 209, 226, 245, 269, 285, 293, 304, 318, 341, 359, 388, 416, 431, 435, 441, 450, 459, 466, 476, 481, 491, 504, 512, 520, 527, 535, 542, 550, 559, 570, 580, 591, 604, 616, 629, 642, 647, 653, 658, 663, 670, 675, 688, 699, 710, 724, 740, 755, 762, 771, 775, 779, 783, 787, 795, 807, 825, 844, 854, 862, 871, 880, 891, 901, 917, 930, 941, 951, 960, 967, 978, 992, 1005, 1023, 1032, 1041, 1050, 1060, 1069, 1078, 1087, 1096, 1105, 1114, 1123, 1132, 1143, 1154, 1163, 1173, 1183, 1194, 1205, 1215, 1226, 1236, 1247, 1268, 1284, 1295, 1302, 1323, 1344, 1363}

func (i TokenType) String() string {
	if i < 0 || i+1 >= TokenType(len(_TokenType_index)) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
