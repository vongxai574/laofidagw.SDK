package laofida

import (
	"github.com/shopspring/decimal"
)

type (
	ReqFilter struct {
		DateStart string `json:"dateStart"`
		DateEnd   string `json:"dateEnd"`
		TIN       string `json:"tin"`
		Type      string `json:"type"`
	}

	SmartTaxRecord struct {
		InstanceID   string          `json:"instanceid"`
		IdeCuoCod    string          `json:"ide_cuo_cod"`
		IdeCuoNam    string          `json:"ide_cuo_nam"`
		IdeTypSad    string          `json:"ide_typ_sad"`
		IdeRegSer    string          `json:"ide_reg_ser"`
		IdeRegNbr    string          `json:"ide_reg_nbr"`
		IdeRegDat    string          `json:"ide_reg_dat"`
		IdeRcpNbr    *string         `json:"ide_rcp_nbr"`
		IdeRcpDat    *string         `json:"ide_rcp_dat"`
		IdeAstNbr    string          `json:"ide_ast_nbr"`
		CmpConCod    *string         `json:"cmp_con_cod"`
		CmpNam       *string         `json:"cmp_nam"`
		DecCod       string          `json:"dec_cod"`
		DecNam       string          `json:"decnam"`
		TptLopCod    string          `json:"tpt_lop_cod"`
		TptMotBrdCod string          `json:"tpt_mot_brd_cod"`
		TptMotDpaNam string          `json:"tpt_mot_dpa_nam"`
		TptMotBrdNam string          `json:"tpt_mot_brd_nam"`
		GenCtyEptCod string          `json:"gen_cty_ept_cod"`
		GenCtyDesCod string          `json:"gen_cty_des_cod"`
		GdsOrgCty    string          `json:"gds_org_cty"`
		TarPrf       *string         `json:"tar_prf"`
		TarPrcExt    string          `json:"tar_prc_ext"`
		TarPrcNat    string          `json:"tar_prc_nat"`
		KeyItmNbr    string          `json:"key_itm_nbr"`
		TarHscNb1    string          `json:"tar_hsc_nb1"`
		TarHscNb2    string          `json:"tar_hsc_nb2"`
		GdsDscV      string          `json:"gds_dsc_v"`
		PckMrk2      string          `json:"pck_mrk2"`
		TarSupQty    string          `json:"tar_sup_qty"`
		TarSupCod    string          `json:"tar_sup_cod"`
		VitWgtGrs    string          `json:"vit_wgt_grs"`
		VitWgtNet    string          `json:"vit_wgt_net"`
		CdRat        string          `json:"cdrat"`
		Cd           decimal.Decimal `json:"cd"`
		ExcRat       string          `json:"excrat"`
		Exc          decimal.Decimal `json:"exc"`
		VatRat       string          `json:"vatrat"`
		Vat          decimal.Decimal `json:"vat"`
		BtRat        string          `json:"btrat"`
		Bt           string          `json:"bt"`
		ExdRat       string          `json:"exdrat"`
		Exd          string          `json:"exd"`
		EdrRat       string          `json:"edrrat"`
		Edr          string          `json:"edr"`
		TotalTax     decimal.Decimal `json:"totaltax"`
		Total        decimal.Decimal `json:"total"`
		VitInvAmtFcx string          `json:"vit_inv_amt_fcx"`
		VitInvAmtNmu string          `json:"vit_inv_amt_nmu"`
		Status       string          `json:"status"`
		IdeAstDat    string          `json:"ide_ast_dat"`
		CmpExpCod    string          `json:"cmp_exp_cod"`
		CmpExpNam    string          `json:"cmp_exp_nam"`
		VgsInvCurCod string          `json:"vgs_inv_cur_cod"`
		VgsInvCurNam string          `json:"vgs_inv_cur_nam"`
		VgsInvCurRat string          `json:"vgs_inv_cur_rat"`
		VgsInvAmtNmu string          `json:"vgs_inv_amt_nmu"`
		VgsInvAmtFcx string          `json:"vgs_inv_amt_fcx"`
		FinTopCod    string          `json:"fin_top_cod"`
		FinTopNam    string          `json:"fin_top_nam"`
		FinBnkCod    string          `json:"fin_bnk_cod"`
		FinBnkNam    string          `json:"fin_bnk_nam"`
		FinBnkFre    string          `json:"fin_bnk_fre"`
		CmpTel       string          `json:"cmp_tel"`
		IdeMan       *string         `json:"ide_man"`
		LnkPrvDoc    *string         `json:"lnk_prv_doc"`
	}

	SmartTaxRecords []*SmartTaxRecord
)
