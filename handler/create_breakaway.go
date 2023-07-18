package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CommunityStruct struct {
  HiveId string `json:"hive_id"`
  Tags string `json:"tags"`
  ServerIP string `json:"server_ip"`
  Password string `json:"server_password"`
}


func CreateBreakaway(c *fiber.Ctx) error {
	payload := new(CommunityStruct)

	if err := c.BodyParser(&payload); err != nil {
		c.JSON(fiber.Map{
			"error": "Can't parse your given paylaod",
		})
		c.Status(400)
		return err
    }

	c.JSON(fiber.Map{
		"message": fmt.Sprintf("created community with id %s", payload.HiveId),
	})

	return nil;
}