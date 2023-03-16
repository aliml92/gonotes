package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


type createWalletRequest struct {
	Wallet struct {
		UserID        string `json:"user_id" binding:"required"`
		WalletAddress string `json:"wallet_address" binding:"required,eth_addr"`
		IpfsCidV1     string `json:"ipfs_cid" binding:"required,ipfs_cid_v1"`
	} `json:"wallet"`
}


func saveWallet(c *gin.Context) {
	var input createWalletRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}


func setupRouter() *gin.Engine {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ipfs_cid_v1", isIpfsCidV1)
	}

	r.POST("/api/v1/wallets", saveWallet)
	return r
}

var isIpfsCidV1 validator.Func = func(fl validator.FieldLevel) bool {
	cid, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	regexp := regexp.MustCompile(`^Qm[a-zA-Z0-9]{44}$`)
	return regexp.MatchString(cid)
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}