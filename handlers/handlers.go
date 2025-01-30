package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"test_Task_New_server/models"
	"test_Task_New_server/repository"
	"test_Task_New_server/logger" 
)

type NewsHandler struct {
	repo repository.NewsRepository
}

func NewNewsHandler(repo repository.NewsRepository) *NewsHandler {
	return &NewsHandler{repo: repo}
}

func (h *NewsHandler) UpdateNews(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Log.Errorf("Invalid ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var news models.News
	if err := c.BodyParser(&news); err != nil {
		logger.Log.Errorf("Invalid request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	news.Id = id

	logger.Log.Infof("Updating news with ID: %d", id)

	if err := h.repo.UpdateNews(&news); err != nil {
		logger.Log.Errorf("Error updating news with ID %d: %v", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Infof("Successfully updated news with ID: %d", id)
	return c.JSON(fiber.Map{"message": "News updated successfully"})
}

func (h *NewsHandler) GetAllNews(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	logger.Log.Infof("Fetching news with limit: %d and offset: %d", limit, offset)

	newsList, err := h.repo.GetAllNews(limit, offset)
	if err != nil {
		logger.Log.Errorf("Error fetching news: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Infof("Successfully fetched %d news items", len(newsList))
	return c.JSON(newsList)
}

func (h *NewsHandler) GetNewsById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	news, err := h.repo.GetNewsWithCategories(int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if news == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	return c.JSON(news)
}