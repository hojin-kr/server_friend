package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hojin-kr/friend"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/friend/list/send/:uuid/:status", func(c *fiber.Ctx) error {
		_friend := &friend.Friend{}
		friends := _friend.FilteredList("UUID_Send", "=", c.Params("uuid"), "Status", "=", c.Params("status"), 100)
		return c.JSON(friends)
	})

	app.Get("/friend/list/receive/:uuid/:status", func(c *fiber.Ctx) error {
		_friend := &friend.Friend{}
		friends := _friend.FilteredList("UUID_Receive", "=", c.Params("uuid"), "Status", "=", c.Params("status"), 100)
		return c.JSON(friends)
	})

	app.Get("/friend/new/:uuid_send/:uuid_receive", func(c *fiber.Ctx) error {
		_friend := &friend.Friend{
			UUID_Send:    c.Params("uuid_send"),
			UUID_Receive: c.Params("uuid_receive"),
		}
		_friend.PendingFriend()
		return c.JSON(_friend)
	})

	app.Get("/friend/accept/:friend_id", func(c *fiber.Ctx) error {
		_friend := &friend.Friend{
			Friend_ID: c.Params("friend_id"),
		}
		_friend.GetFriend()
		_friend.AcceptFriend()
		return c.JSON(_friend)
	})

	app.Get("/friend/reject/:friend_id", func(c *fiber.Ctx) error {
		_friend := &friend.Friend{
			Friend_ID: c.Params("friend_id"),
		}
		_friend.GetFriend()
		_friend.RejectFriend()
		return c.JSON(_friend)
	})

	app.Get("/friend/block/:friend_id", func(c *fiber.Ctx) error {
		_friend := &friend.Friend{
			Friend_ID: c.Params("friend_id"),
		}
		_friend.GetFriend()
		_friend.BlockFriend()
		return c.JSON(_friend)
	})

	app.Listen(":3000")
}
