// generated by stringer -type=TokenType; DO NOT EDIT

package c6

import "fmt"

const _TokenType_name = "T_SPACET_COMMENT_LINET_COMMENT_BLOCKT_SEMICOLONT_COMMAT_IDENTT_URLT_ID_SELECTORT_CLASS_SELECTORT_TYPE_SELECTORT_UNIVERSAL_SELECTORT_PARENT_SELECTORT_PSEUDO_SELECTORT_INTERPOLATION_SELECTORT_CONCATT_AND_SELECTORT_DESCENDANT_SELECTORT_CHILD_SELECTORT_PLUST_GTT_BRACE_STARTT_BRACE_ENDT_LANG_CODET_BRACKET_LEFTT_ATTRIBUTE_NAMET_BRACKET_RIGHTT_EQUALT_TILDE_EQUALT_PIPE_EQUALT_VARIABLET_IMPORTT_CHARSETT_QQ_STRINGT_Q_STRINGT_UNQUOTE_STRINGT_PAREN_STARTT_PAREN_ENDT_CONSTANTT_INTEGERT_FLOATT_UNIT_PXT_UNIT_PTT_UNIT_EMT_UNIT_REMT_UNIT_DEGT_UNIT_PERCENTT_PROPERTY_NAMET_PROPERTY_VALUET_HEX_COLORT_COLONT_INTERPOLATION_STARTT_INTERPOLATION_INNERT_INTERPOLATION_ENDT_DIVT_MULT_MINUS"

var _TokenType_index = [...]uint16{0, 7, 21, 36, 47, 54, 61, 66, 79, 95, 110, 130, 147, 164, 188, 196, 210, 231, 247, 253, 257, 270, 281, 292, 306, 322, 337, 344, 357, 369, 379, 387, 396, 407, 417, 433, 446, 457, 467, 476, 483, 492, 501, 510, 520, 530, 544, 559, 575, 586, 593, 614, 635, 654, 659, 664, 671}

func (i TokenType) String() string {
	if i < 0 || i+1 >= TokenType(len(_TokenType_index)) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
