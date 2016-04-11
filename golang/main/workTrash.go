
import (
	"ssp_mbv_web/models/adslot"
	"ssp_web/models/notify"
	"ssp_web/utils"
	"sunteng/commons/constant"

	"gopkg.in/mgo.v2/bson"
)

func (this *AdSlotController) ArchiveRecover(doType int) {
	ids := this.intArrayQuery("ids")
	extraQuery := bson.M{"ChildStatus.Archived": bson.M{"$ne": doType}}
	adSlots, err := adslot.GetAdSlotByIds(ids, extraQuery, this.userBson())
	if err != nil {
		logger.Noticef("广告位查询失败: %s", err.Error())
		this.writeError(this.locale("query_fail"))
		return
	}
	if len(adSlots) < 1 {
		this.writeError(this.locale("recoder_not_found"))
		return
	}

	var appIds []int
	var illegalIds []int
	var normalIds []int
	for _, v := range adSlots {
		appIds = append(appIds, v.AppId)
		if v.ChildStatus == nil {
			logger.Noticef("广告位: %d数据错误，检查!", v.Id)
			this.writeError(this.locale("verify_data_load_fail"))
			return
		}
		v.ChildStatus.Archived = doType
		v.SumAndSetStatus()
		if doType == constant.StatusArchived {
			normalIds = append(normalIds, v.Id)
		}
	}

	illegalIds = append(illegalIds, utils.IntArraySubtract(ids, normalIds)...)

	err = adslot.ArchiveRecoverByIds(normalIds, adSlots[0].Status, adSlots[0].ChildStatus.Archived, this.userBson())
	if err != nil {
		if doType == constant.StatusOrigin {
			logger.Noticef("广告位还原失败: %s", err.Error())
			this.writeError(this.locale("recover_fail"))
			return
		}
		logger.Noticef("广告位归档失败: %s", err.Error())
		this.writeError(this.locale("archived_fail"))
		return
	}

	if doType == constant.StatusArchived {
		notify.Do(constant.NotifyTypeAdSlot, constant.ActionDel, normalIds...)

		adslot.ArchiveAdSlotOrderByAdSlotIds(normalIds)
		if len(illegalIds) > 0 {
			this.writeError(this.locale("archived_part", normalIds, illegalIds))
			return
		}
		this.writeSuccess(this.locale("archived_success"))
		return
	}

	var notifyIds []int
	for _, adSlot := range adSlots {
		if adSlot.Status != constant.StatusDraft {
			notifyIds = append(notifyIds, adSlot.Id)
		}
	}
	notify.Do(constant.NotifyTypeAdSlot, constant.ActionAdd, notifyIds...)
	adslot.RecoverAdSlotOrderByAdSlotIds(ids)
	if len(illegalIds) > 0 {
		this.writeError(this.locale("recover_part", normalIds, illegalIds))
		return
	}
	this.writeSuccess(this.locale("recover_success"))
	return
}


println("------ad archvie recover  pause \n========================\n")

func (this *AdController) Recover() {
	if !this.HasPermission("material.recover") {
		this.writeError(this.locale("no_permission"))
		return
	}

	this.ArchiveRecover(constant.StatusOrigin)
	// var normalIds []int
	// doType := constant.StatusOrigin
	// extraQuery := bson.M{"ChildStatus.Archived": bson.M{"$ne": doType}}
	// ids := this.intArrayQuery("ids")
	// ads, err := ad.GetAdByIds(ids, this.userBson())
	// if err != nil {
	// 	logger.Noticef("广告查询失败: %s", err.Error())
	// 	this.writeError(this.locale("archived_fail"))
	// 	return
	// }
	//
	// for _, v := range ads {
	// 	normalIds = append(normalIds, v.Id)
	// 	if v.ChildStatus == nil {
	// 		logger.Noticef("广告: %d数据错误，检查!", v.Id)
	// 		this.writeError(this.locale("verify_data_load_fail"))
	// 		return
	// 	}
	// 	v.ChildStatus.Archived = doType
	// 	v.SumAndSetStatus()
	// }
	//
	// var illegalIds []int
	// illegalIds = utils.IntArraySubtract(ids, normalIds)
	// err := ad.ArchiveRecoverByIds(normalIds, this.userBson())
	// if err != nil {
	// 	logger.Noticef("广告还原失败: %s", err.Error())
	// 	this.writeError(this.locale("recover_fail"))
	// 	return
	// }
	//
	// notify.Do(constant.NotifyTypeAd, constant.ActionAdd, normalIds...)
	// order.RecoverOrderAdByAdIds(normalIds)
	// if len(illegalIds) > 0 {
	// 	this.writeError(this.locale("recover_part", normalIds, illegalIds))
	// 	return
	// }
	// this.writeSuccess(this.locale("recover_success"))
	// return
}

func (this *AdController) Archive() {
	if !this.HasPermission("material.archive") {
		this.writeError(this.locale("no_permission"))
		return
	}

	this.ArchiveRecover(constant.StatusArchived)
	// ids := this.intArrayQuery("ids")
	// err := ad.GetAdByIds(ids, this.userBson())
	// if err != nil {
	// 	logger.Noticef("广告归档失败: %s", err.Error())
	// 	this.writeError(this.locale("archived_fail"))
	// 	return
	// }
	//
	// notify.Do(constant.NotifyTypeAd, constant.ActionDel, ids...)
	// err = order.ArchiveOrderAdByAdIds(ids)
	// if err != nil {
	// 	logger.Noticef("广告归档处理关联表失败: %s", err.Error())
	// 	this.writeError(this.locale("archive_fail"))
	// 	return
	// }
	// this.writeSuccess(this.locale("archived_success"))
}
