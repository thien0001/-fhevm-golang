package handlers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/thien0001/fhevm-go-demo/pkg/crypto"
	"github.com/thien0001/fhevm-go-demo/pkg/fhevm"
)

type ApiHandler struct {
	FheClient *fhevm.Client
}

func NewApiHandler(c *fhevm.Client) *ApiHandler {
	return &ApiHandler{FheClient: c}
}

type EncryptRequest struct {
	Data string `json:"data"`
}

type SendRequest struct {
	Cipher string `json:"cipher"`
	To     string `json:"to"`
}

func (h *ApiHandler) Encrypt(c *fiber.Ctx) error {
	req := new(EncryptRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	cipher, err := crypto.Encrypt(req.Data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"cipher": cipher})
}

func (h *ApiHandler) Send(c *fiber.Ctx) error {
	req := new(SendRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	toAddr := common.HexToAddress(req.To)
	txHash, err := h.FheClient.SendRawDataToContract(toAddr, []byte(req.Cipher))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"txHash": txHash})
}

func (h *ApiHandler) Result(c *fiber.Ctx) error {
	txHash := c.Params("txHash")

	fakeCipher := "REPLACE_WITH_RETURNED_CIPHER_IF_APPLICABLE"
	plain, _ := crypto.Decrypt(fakeCipher)

	return c.JSON(fiber.Map{
		"cipherResult": fakeCipher,
		"plain":        plain,
		"txHash":       txHash,
	})
}
