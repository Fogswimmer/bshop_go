package fileservice

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kashifkhan0771/utils/slugger"
)

func GetUploadRootDir() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("UPLOADS_PATH")
}

func GetRelUploadsSubDir(subdir string, title string) string {
	s := slugger.New(map[string]string{}, false)
	slug := s.Slug(title, "-")
	return subdir + "/" + slug
}

func GetAbsUploadsSubDir(subdir string, title string) string {
	return GetUploadRootDir() + "/" + GetRelUploadsSubDir(subdir, title)
}

func UploadFile(c *gin.Context, file *multipart.FileHeader, absPath string) (string, error) {
	_, err := os.Stat(absPath)
	if err == nil {
		os.Remove(absPath)
	}
	fp := absPath + "/" + file.Filename
	err = c.SaveUploadedFile(file, fp)
	return fp, err
}

func GetStaticFileURL(path string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("STATIC_URL") + "/" + path
}
