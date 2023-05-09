package utils

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/thumbnails/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	//	HANDLE FILE
	file, errFile := ctx.FormFile("thumbnail")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	var filename *string
	if file != nil {
		errCheckContentType := checkContentType(file, "image/jpg", "image/png", "image/gif")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}

		filename = &file.Filename
		extenstionFile := filepath.Ext(*filename)
		newFilename := fmt.Sprintf("%s%s", RandStringRunes(10), extenstionFile)

		filename = &newFilename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/thumbnails/%s", newFilename))
		if errSaveFile != nil {
			log.Println("Fail to store file into public/thumbnails directory.")
		}
	} else {
		log.Println("Nothing file to uploading.")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error Read Multipart Form Request, Error = ", errForm)
	}

	files := form.File["photos"]

	var filenames []string

	for i, file := range files {
		var filename string
		if file != nil {
			extenstionFile := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%d-%s%s", i, "gambar", extenstionFile)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/thumbnails/%s", filename))
			if errSaveFile != nil {
				log.Println("Fail to store file into public/thumbnails directory.")
			}
		} else {
			log.Println("Nothing file to uploading.")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	ctx.Locals("filenames", filenames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	}

	return nil
}

func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}

		return errors.New("not allowed file type")
	} else {
		return errors.New("not found content type to be checking")
	}
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
