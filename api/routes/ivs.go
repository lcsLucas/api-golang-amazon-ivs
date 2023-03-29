package routes

import (
	"encoding/json"
	"golang-ivs/service"
	ivs "golang-ivs/service/ivs/channel"
	"golang-ivs/utils"

	"github.com/gofiber/fiber/v2"
)

var svcchannel ivs.ServiceIVSChannel

func InicializeRoutesIVS() {

	App.Get("/ivs/channels/list", func(c *fiber.Ctx) error {
		var err error

		/* validação jwt */
		/*
			tokenID, err := jwt.ExtractTokenID(c)
			if err != nil {
				c.Status(fiber.StatusUnprocessableEntity)

				return c.JSON(utils.ResponseError{
					Message: err.Error(),
				})
			}

			fmt.Println(tokenID)
		*/

		ctx := c.UserContext()

		channels, err := svcchannel.ListChannels(ctx)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(channels)
	})

	App.Get("/ivs/channels/get", func(c *fiber.Ctx) error {
		var err error

		/* validação jwt */
		/*
			tokenID, err := jwt.ExtractTokenID(c)
			if err != nil {
				c.Status(fiber.StatusUnprocessableEntity)

				return c.JSON(utils.ResponseError{
					Message: err.Error(),
				})
			}

			fmt.Println(tokenID)
		*/

		var paramBody struct {
			Arn string `json:"arn"`
		}

		err = json.Unmarshal(c.Body(), &paramBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: err.Error(),
			})
		}

		if len(paramBody.Arn) < 1 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: "Parametro inválido",
			})
		}

		ctx := c.UserContext()

		cn, err := svcchannel.GetChannel(ctx, paramBody.Arn)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(cn)
	}) //get channel

	/*
		App.Post("/ivs/channels/") // create channel
		App.Put("/ivs/channels/1") // update channel
		App.Delete("/ivs/channels/1") // delete channel
		App.Get("/ivs/status") // status service
	*/

	App.Get("/ivs/stream-key/list", func(c *fiber.Ctx) error {
		var err error

		/* validação jwt */
		/*
			tokenID, err := jwt.ExtractTokenID(c)
			if err != nil {
				c.Status(fiber.StatusUnprocessableEntity)

				return c.JSON(utils.ResponseError{
					Message: err.Error(),
				})
			}

			fmt.Println(tokenID)
		*/

		var paramBody struct {
			Arn string `json:"arn"`
		}

		err = json.Unmarshal(c.Body(), &paramBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: err.Error(),
			})
		}

		if len(paramBody.Arn) < 1 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: "Parametro inválido",
			})
		}

		ctx := c.UserContext()

		keys, err := svcchannel.ListStreamKey(ctx, paramBody.Arn)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(keys)

	}) // lista stream key de um canal

	App.Get("/ivs/stream-key/get", func(c *fiber.Ctx) error {
		var err error

		/* validação jwt */
		/*
			tokenID, err := jwt.ExtractTokenID(c)
			if err != nil {
				c.Status(fiber.StatusUnprocessableEntity)

				return c.JSON(utils.ResponseError{
					Message: err.Error(),
				})
			}

			fmt.Println(tokenID)
		*/

		var paramBody struct {
			Arn string `json:"arn"`
		}

		err = json.Unmarshal(c.Body(), &paramBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: err.Error(),
			})
		}

		if len(paramBody.Arn) < 1 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: "Parametro inválido",
			})
		}

		ctx := c.UserContext()

		key, err := svcchannel.GetStreamKey(ctx, paramBody.Arn)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ResponseError{
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(key)
	}) // get stream key de um canal

	App.Put("/new-stream", func(c *fiber.Ctx) error {

		ctx := c.UserContext()

		dataResponse, err := service.NewStream(ctx)

		if err != nil {
			c.Status(fiber.StatusUnprocessableEntity)

			return c.JSON(utils.ResponseError{
				Message: err.Error(),
			})

		}

		c.Status(fiber.StatusOK)
		return c.JSON(dataResponse)

	})

}

func init() {
	svcchannel = ivs.NewServiceChannel()
}
