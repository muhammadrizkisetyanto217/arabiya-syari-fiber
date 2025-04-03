package user_progress

import (
	"arabiya-syari-fiber/internals/models/progress_user"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RankLevelRequirementController struct {
	DB *gorm.DB
}

func NewRankLevelRequirementController(db *gorm.DB) *RankLevelRequirementController {
	return &RankLevelRequirementController{DB: db}
}

// GET /api/rank_level_requirements
func (ctrl *RankLevelRequirementController) GetAll(c *fiber.Ctx) error {
	var ranks []progress_user.RankLevelRequirement
	if err := ctrl.DB.Find(&ranks).Error; err != nil {
		log.Println("[ERROR] GetAll RankLevelRequirement:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data",
		})
	}
	return c.JSON(ranks)
}

// GET /api/rank_level_requirements/:id
func (ctrl *RankLevelRequirementController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var rank progress_user.RankLevelRequirement
	if err := ctrl.DB.First(&rank, id).Error; err != nil {
		log.Println("[ERROR] GetByID RankLevelRequirement:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data tidak ditemukan",
		})
	}
	return c.JSON(rank)
}

// POST /api/rank_level_requirements
func (ctrl *RankLevelRequirementController) Create(c *fiber.Ctx) error {
	var multiple []progress_user.RankLevelRequirement
	var single progress_user.RankLevelRequirement

	// Coba parse sebagai array dulu
	if err := c.BodyParser(&multiple); err == nil && len(multiple) > 0 {
		// Input adalah array
		if err := ctrl.DB.Create(&multiple).Error; err != nil {
			log.Println("[ERROR] Gagal menyimpan array data:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menyimpan data array",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Berhasil simpan banyak data",
			"data":    multiple,
		})
	}

	// Jika bukan array, parse sebagai objek tunggal
	if err := c.BodyParser(&single); err != nil {
		log.Println("[ERROR] BodyParser gagal (bukan object/array):", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Request tidak valid",
		})
	}

	if err := ctrl.DB.Create(&single).Error; err != nil {
		log.Println("[ERROR] Gagal menyimpan satu data:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menyimpan data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil simpan satu data",
		"data":    single,
	})
}

// PUT /api/rank_level_requirements/:id
func (ctrl *RankLevelRequirementController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var rank progress_user.RankLevelRequirement

	if err := ctrl.DB.First(&rank, id).Error; err != nil {
		log.Println("[ERROR] Update First RankLevelRequirement:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data tidak ditemukan",
		})
	}

	if err := c.BodyParser(&rank); err != nil {
		log.Println("[ERROR] BodyParser Update:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Request tidak valid",
		})
	}

	if err := ctrl.DB.Save(&rank).Error; err != nil {
		log.Println("[ERROR] Save Update RankLevelRequirement:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update data",
		})
	}
	return c.JSON(rank)
}

// DELETE /api/rank_level_requirements/:id
func (ctrl *RankLevelRequirementController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := ctrl.DB.Delete(&progress_user.RankLevelRequirement{}, id).Error; err != nil {
		log.Println("[ERROR] Delete RankLevelRequirement:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus data",
		})
	}
	return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
}
