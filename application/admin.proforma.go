package application

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getProformaHandler(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma

	err = fetchProforma(ctx, pid, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jp)
}

func getProformaByCompanyHandler(ctx *gin.Context) {
	cid, err := util.ParseUint(ctx.Param("cid"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var jps []Proforma

	err = fetchProformaByCompanyRC(ctx, cid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jps)
}

func putProforma(ctx *gin.Context) {
	var jp Proforma

	err := ctx.BindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateProforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserID(ctx)

	logrus.Infof("%v edited a proforma with id %d", user, jp.ID)

	ctx.JSON(200, jp)
}

func putProformaByCompanyID(ctx *gin.Context) {
	cid_string := ctx.Param("cid")
	cid, err := util.ParseUint(cid_string)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma
	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Company ID mismatch"})
		return
	}

	err = ctx.BindJSON(&jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateProforma(ctx, &jp)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": "edited proforma"})
}

func deleteProformaByCompanyID(ctx *gin.Context) {
	pid_string := ctx.Param("pid")
	pid, err := util.ParseUint(pid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	cid_string := ctx.Param("cid")
	cid, err := util.ParseUint(cid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jp Proforma
	err = fetchProforma(ctx, pid, &jp)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if jp.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Company ID mismatch"})
		return
	}

	err = deleteProforma(ctx, pid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": "deleted proforma"})
}
