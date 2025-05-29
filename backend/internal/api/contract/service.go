package contract

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// API структура для работы с договорами
type API struct {
	contractService service.ContractService
}

// NewAPI возвращает новый объект API с зависимостями
func NewAPI(contractService service.ContractService) *API {
	return &API{
		contractService: contractService,
	}
}

// RegisterRoutes регистрирует все маршруты API для работы с платежами
func (API *API) RegisterRoutes(router *gin.Engine) {
	contractGroup := router.Group("/contracts")
	{
		// contractGroup.POST("", API.CreateContract)
		// contractGroup.GET("", API.GetAllContracts)
		contractGroup.GET("/:contract_id", API.DownloadContract)

	}
}

// DownloadContract отдает файл договора по contract_id
// @Summary Скачать договор аренды
// @Description Возвращает PDF-файл договора, связанного с арендуемой записью
// @Tags contracts
// @Produce application/pdf
// @Security BearerAuth
// @Param contract_id path int64 true "ID аренды"
// @Success 200 {file} file
// @Failure 400 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse "Договор не найден"
// @Failure 500 {object} sys.ErrorResponse
// @Router /contracts/{contract_id} [get]
func (a *API) DownloadContract(c *gin.Context) {
	contractIDParam := c.Param("contract_id")
	contractID, err := strconv.ParseInt(contractIDParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	contract, err := a.contractService.GetByContractID(c.Request.Context(), contractID)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	file, err := os.Open(contract.FilePath)
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	defer file.Close()
	fmt.Println(contract)

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=contract-%d.pdf", contract.ContractID))
	c.Header("Content-Type", "application/pdf")
	c.File(contract.FilePath)
}

// type CreateContractInput struct {
// 	RentalID  int    `json:"rentalId" binding:"required"`
// 	FilePath  string `json:"filePath" binding:"required"`
// 	Signature string `json:"signature" binding:"required"` // base64
// 	PublicKey string `json:"publicKey" binding:"required"`
// }

// func (api *API) CreateContract(c *gin.Context) {
// 	var input CreateContractInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	sigBytes, err := decodeBase64(input.Signature)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
// 		return
// 	}

// 	contract := model.Contract{
// 		RentalID:  input.RentalID,
// 		FilePath:  input.FilePath,
// 		Signature: sigBytes,
// 		PublicKey: input.PublicKey,
// 		IsActive:  true,
// 		SignedAt:  time.Now(),
// 	}

// 	if err := api.service.CreateContract(c.Request.Context(), contract); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "contract created"})
// }

// func (api *API) GetContractsByRental(c *gin.Context) {
// 	rentalID, err := parseIntParam(c, "rentalId")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental id"})
// 		return
// 	}

// 	contracts, err := api.service.GetContractsByRental(c.Request.Context(), rentalID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, contracts)
// }

// func (api *API) GetAllContracts(c *gin.Context) {
// 	contracts, err := api.service.GetAllContracts(c.Request.Context())
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, contracts)
// }
