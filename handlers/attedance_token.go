package handlers

import (
	"time"

	"github.com/KicauOrgspark/BE-Absensi-Siswa/database"
	"github.com/KicauOrgspark/BE-Absensi-Siswa/dto/requests"
	"github.com/KicauOrgspark/BE-Absensi-Siswa/mappers"
	"github.com/KicauOrgspark/BE-Absensi-Siswa/models"
	"github.com/KicauOrgspark/BE-Absensi-Siswa/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateToken(c *fiber.Ctx) error {
	adminID := c.Locals("user_id").(int64)

	var req requests.TokenReq
	if err := c.BodyParser(&req);err != nil {
		return c.Status(400).JSON(fiber.Map{"error" : "invalid payload"})
	}

	token, err := utils.CreateToken(adminID, req.Duration, req.LateAfter)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error" : err})
	}

	return c.Status(201).JSON(fiber.Map{"Message" : "Success to create Token !", "data" : mappers.ToTokenResponse(token)})
}


func SubmitToken(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int64)

	var req requests.SubmitToken
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error" : "invalid payload"})
	}

	token, err := utils.VerifyTokenCode(req.TokenCode)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error" : "token invalid or expired !"})
	}

	status := "present"

	if time.Now().After(token.LateAfter) {
		status = "late"
	}

	log := models.AttedanceLogs{
		UserID:      userID,
		TokenID:     token.ID,
		Status:      status,
		ClockInTime: time.Now(),
	}

	if err := database.DB.Create(&log).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error"  : err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"Message" : "Success To Absen"})
}