package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spknetwork/ecency_starting-point_be/helper"
)

type CommunityStruct struct {
  HiveId string `json:"hive_id"`
  Tags string `json:"tags"`
  ServerIP string `json:"server_ip"`
  ServerUsername string `json:"server_username"`
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

	conn, err := helper.Connect(fmt.Sprintf("%s:22", payload.ServerIP), payload.ServerUsername, payload.Password)
	if err != nil {
		c.JSON(fiber.Map{
			"error": "Can't connect to given server",
		})
		c.Status(500)
		return err
	}

	if _, err := conn.SendCommands("sudo apt-get install docker"); err != nil {
		c.JSON(fiber.Map{
			"error": "Can't send command to given server",
		})
		c.Status(500)
		return err
	}

	return nil;
}