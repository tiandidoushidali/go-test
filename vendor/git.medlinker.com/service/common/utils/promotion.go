package utils

import (
	"fmt"
	"strings"
)

type PromotionInfos struct {
	// 药品spu_pack_id
	SpuPackId int64 `json:"SpuPackId"`
	// 满减
	FullCut *InfosFullCutEntity `json:"fullCut"`
	// 折扣
	Discount *InfosDiscountEntity `json:"discount"`
	// 换购
	Redeem []*InfosRedeemEntity `json:"redeem"`
}

type InfosEntity struct {
	// 活动pid
	Pid string `json:"pid"`
	// 活动名称
	Name string `json:"name"`
	// 开始时间
	StartAt int64 `json:"startAt"`
	// 结束时间
	EndAt int64 `json:"endAt"`
}

type InfosFullCutEntity struct {
	InfosEntity
	// 具体规则
	Rules []*InfosFullCutRuleEntity `json:"rules"`
}

type InfosFullCutRuleEntity struct {
	// 满多少盒
	Full int64 `json:"full"`
	// 减多少钱，单位分
	Cut int64 `json:"cut"`
	// 满减类型，1-金额直减，2-件数直减
	PcutType int64 `json:"pcutType"`
	//满减额外规则
	ExtEntity *InfosFullCutRuleExtEntity `json:"extEntity"`
}

type InfosFullCutRuleExtEntity struct {
	BaseFull   int64 `json:"baseFull"`
	BaseCut    int64 `json:"baseCut"`
	LimitLower int64 `json:"LimitLower"`
	LimitUpper int64 `json:"LimitUpper"`
}

type InfosDiscountEntity struct {
	InfosEntity
	// 每人限制购买数量, -1代表不限量
	PerUserLimit int64 `json:"perUserLimit"`
	// 药品总限量
	LimitAmount int64 `json:"limitAmount"`
	// 折扣率, 50代表5折
	Rate int64 `json:"rate"`
}

type InfosRedeemEntity struct {
	InfosEntity
	// 换购具体规则
	Rules []*InfosRedeemRuleEntity `json:"rules"`
}

type InfosRedeemRuleEntity struct {
	// 第几次换购，0代表统一换购
	RedeemTimes int64 `json:"redeemTimes"`
	// 满多少, 如果是金额单位是分
	Full int64 `json:"full"`
	// 满减类型，1-金额直减，2-件数直减
	PcutType int64 `json:"pcutType"`
	// 加多少钱，单位分
	Add int64 `json:"add"`
	// 换购的药品列表
	RedeemGoods []*InfosRedeemCutEntity `json:"redeemGoods"`
}

type InfosRedeemCutEntity struct {
	// 换购品数量
	Count int64 `json:"count"`
	// 换购品spu_pack_id
	SpuPackId int64 `json:"SpuPackId"`
	// 换购品商品名
	TradeName string `json:"tradeName"`
	// 换购品规格
	Pack string `json:"pack"`
}

// PromotionInfoEntity ...
type PromotionInfoEntity struct {
	// the pid
	Pid string
	// 优惠标签
	Tips string
	// 优惠展示标题
	Title string
	// 优惠规则
	List []string
}

// tips
const (
	FullCutTips  = "满减"
	DiscountTips = "折扣"
	RedeemTips   = "换购"
)

// title
const (
	FullCutTitle             = "满%d盒减%s元"
	DiscountTitle            = "%s折"
	DiscountTitleWithLimit   = "%s折，限购%d件"
	RedeemTitle              = "查看换购规则"
	FullCutMuIncludeLowerUp  = "每满%d盒优惠%s元（%d盒起，上限%d盒）"
	FullCutMuIncludeLower    = "每满%d盒优惠%s元（%d盒起）"
	FullCutMuIncludeUp       = "每满%d盒优惠%s元（上限%d盒）"
	FullCutExcludeLowerUpper = "每满%d盒优惠%s元"
)

const (
	RedeemRule          = "满%s%s，加%s元可换购，%s" // 满a(盒/元)，加b元可换购，+RedeemGoods1, +RedeemGoods2...
	RedeemRuleWithTimes = "第%d次%s"           //第n次 + RedeemRule
	RedeemGoods         = "%d盒%s %s"         //n盒a药品b规格
)

// pcut_type
const (
	PcutTypeYuan = iota + 1
	PcutTypeHe
)

var pcutmap = map[int64]string{
	PcutTypeYuan: "元",
	PcutTypeHe:   "盒",
}

type PromotionFullType int32

const (
	// ignore
	PromotionFullType_PcutIgnore PromotionFullType = 0
	// 金额
	PromotionFullType_Amount PromotionFullType = 1
	// 个数
	PromotionFullType_Count PromotionFullType = 2
	// 指定价格
	PromotionFullType_FixPrice PromotionFullType = 3
	// 折扣
	PromotionFullType_Discount PromotionFullType = 4
	//倍数扣减
	PromotionFullType_Multiple PromotionFullType = 5
)

// BuildPromotionInfo 构建优惠信息
// 医生端和患者端都使用
func BuildPromotionInfo(in *PromotionInfos) []*PromotionInfoEntity {
	if in == nil {
		return nil
	}

	ret := make([]*PromotionInfoEntity, 0, 4)
	if fullCut := buildFullCut(in.FullCut); fullCut != nil {
		ret = append(ret, fullCut)
	}

	if discount := buildDiscount(in.Discount); discount != nil {
		ret = append(ret, discount)
	}

	if redeem := buildRedeem(in.Redeem); redeem != nil {
		ret = append(ret, redeem)
	}

	return ret
}

// 构建满减
func buildFullCut(fullCut *InfosFullCutEntity) *PromotionInfoEntity {
	if fullCut == nil || len(fullCut.Rules) == 0 || len(fullCut.Pid) == 0 {
		return nil
	}

	format := func(rule *InfosFullCutRuleEntity) string {
		if rule.PcutType == int64(PromotionFullType_Multiple) {
			if rule.ExtEntity == nil {
				return ""
			}
			if rule.ExtEntity.LimitUpper != 0 && rule.ExtEntity.LimitLower != 0 {
				return fmt.Sprintf(FullCutMuIncludeLowerUp, rule.ExtEntity.BaseFull, formatMoney(rule.ExtEntity.BaseCut), rule.ExtEntity.LimitLower, rule.ExtEntity.LimitUpper)
			}
			if rule.ExtEntity.LimitUpper != 0 {
				return fmt.Sprintf(FullCutMuIncludeUp, rule.ExtEntity.BaseFull, formatMoney(rule.ExtEntity.BaseCut), rule.ExtEntity.LimitUpper)
			}
			if rule.ExtEntity.LimitLower != 0 {
				return fmt.Sprintf(FullCutMuIncludeLower, rule.ExtEntity.BaseFull, formatMoney(rule.ExtEntity.BaseCut), rule.ExtEntity.LimitLower)
			}
			return fmt.Sprintf(FullCutExcludeLowerUpper, rule.ExtEntity.BaseFull, formatMoney(rule.ExtEntity.BaseCut))
		} else {
			return fmt.Sprintf(FullCutTitle, rule.Full, formatMoney(rule.Cut))
		}
	}

	ret := &PromotionInfoEntity{
		Pid:   fullCut.Pid,
		Tips:  FullCutTips,
		Title: format(fullCut.Rules[0]),
	}

	for _, v := range fullCut.Rules {
		ret.List = append(ret.List, format(v))
	}

	return ret
}

// 构建折扣
func buildDiscount(discount *InfosDiscountEntity) *PromotionInfoEntity {
	// 0折和10折不展示
	if discount == nil || discount.Rate <= 0 || discount.Rate >= 100 {
		return nil
	}

	ret := &PromotionInfoEntity{
		Pid:  discount.Pid,
		Tips: DiscountTips,
	}

	formatRate := func(rate int64) string {
		if rate%10 == 0 {
			return fmt.Sprintf("%d", rate/10)
		} else {
			return fmt.Sprintf("%.1f", float64(rate)/10)
		}
	}

	if discount.PerUserLimit <= 0 {
		ret.Title = fmt.Sprintf(DiscountTitle, formatRate(discount.Rate))
	} else {
		ret.Title = fmt.Sprintf(DiscountTitleWithLimit, formatRate(discount.Rate), discount.PerUserLimit)
	}

	return ret
}

// 构建换购
func buildRedeem(redeems []*InfosRedeemEntity) *PromotionInfoEntity {
	if len(redeems) == 0 {
		return nil
	}

	ret := &PromotionInfoEntity{
		Pid:   redeems[0].Pid,
		Tips:  RedeemTips,
		Title: RedeemTitle,
	}

	getRules := func(in *InfosRedeemEntity) []*InfosRedeemRuleEntity {
		if in == nil {
			return nil
		}
		return in.Rules
	}

	getGoods := func(in *InfosRedeemRuleEntity) []*InfosRedeemCutEntity {
		if in == nil {
			return nil
		}
		return in.RedeemGoods
	}

	format := func(rule *InfosRedeemRuleEntity) string {
		switch rule.PcutType {
		case PcutTypeHe:
			return fmt.Sprintf("%d", rule.Full)
		case PcutTypeYuan:
			return formatMoney(rule.Full)
		}
		return ""
	}

	for _, redeem := range redeems {
		if redeem == nil || len(redeem.Pid) == 0 {
			continue
		}

		for _, rule := range getRules(redeem) {
			var goods []string
			for _, v := range getGoods(rule) {
				goods = append(goods, fmt.Sprintf(RedeemGoods, v.Count, v.TradeName, v.Pack))
			}
			buildRule := fmt.Sprintf(RedeemRule, format(rule), pcutmap[rule.PcutType], formatMoney(rule.Add), strings.Join(goods, "，"))
			if rule.RedeemTimes > 0 {
				buildRule = fmt.Sprintf(RedeemRuleWithTimes, rule.RedeemTimes, buildRule)
			}
			ret.List = append(ret.List, buildRule)
		}
	}

	return ret
}

// formatMoney money单位分
func formatMoney(money int64) string {
	var yuan string
	if money%100 == 0 {
		yuan = fmt.Sprintf("%d", money/100)
	} else if money%10 == 0 {
		yuan = fmt.Sprintf("%.1f", float64(money)/100)
	} else {
		yuan = fmt.Sprintf("%.2f", float64(money)/100)
	}
	return yuan
}
